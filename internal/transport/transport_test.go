package transport

import (
	"context"
	"errors"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"sync/atomic"
	"testing"
)

func TestWrapRoundTripperSupportsHTTPSProxy(t *testing.T) {
	t.Parallel()

	target := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`ok`))
	}))
	defer target.Close()

	var proxyCalls atomic.Int32
	proxy := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodConnect {
			t.Fatalf("unexpected proxy method: %s", r.Method)
		}
		proxyCalls.Add(1)

		hj, ok := w.(http.Hijacker)
		if !ok {
			t.Fatal("proxy response writer does not support hijacking")
		}
		clientConn, _, err := hj.Hijack()
		if err != nil {
			t.Fatalf("Hijack returned error: %v", err)
		}

		targetConn, err := net.Dial("tcp", target.Listener.Addr().String())
		if err != nil {
			_ = clientConn.Close()
			t.Fatalf("Dial target returned error: %v", err)
		}

		_, _ = clientConn.Write([]byte("HTTP/1.1 200 Connection Established\r\n\r\n"))
		go proxyCopy(clientConn, targetConn)
		go proxyCopy(targetConn, clientConn)
	}))
	defer proxy.Close()

	proxyURL, err := url.Parse(proxy.URL)
	if err != nil {
		t.Fatalf("Parse proxy URL returned error: %v", err)
	}

	base := target.Client().Transport.(*http.Transport)
	rt, err := WrapRoundTripper(base, staticProxySelector{url: proxyURL})
	if err != nil {
		t.Fatalf("WrapRoundTripper returned error: %v", err)
	}
	client := &http.Client{
		Transport: rt,
	}

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, target.URL, nil)
	if err != nil {
		t.Fatalf("NewRequestWithContext returned error: %v", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("client.Do returned error: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("ReadAll returned error: %v", err)
	}
	if got, want := string(body), "ok"; got != want {
		t.Fatalf("body = %q, want %q", got, want)
	}
	if proxyCalls.Load() == 0 {
		t.Fatal("expected proxy CONNECT to be used")
	}
}

func TestWrapRoundTripperPreservesCustomRoundTripperWithoutSelector(t *testing.T) {
	t.Parallel()

	custom := roundTripperFunc(func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusNoContent,
			Body:       io.NopCloser(strings.NewReader("")),
			Header:     make(http.Header),
		}, nil
	})

	rt, err := WrapRoundTripper(custom, nil)
	if err != nil {
		t.Fatalf("WrapRoundTripper returned error: %v", err)
	}

	resp, err := rt.RoundTrip(httptest.NewRequest(http.MethodGet, "https://example.com", nil))
	if err != nil {
		t.Fatalf("RoundTrip returned error: %v", err)
	}
	if got, want := resp.StatusCode, http.StatusNoContent; got != want {
		t.Fatalf("status = %d, want %d", got, want)
	}
}

func TestWrapRoundTripperRejectsCustomRoundTripperWithProxySelector(t *testing.T) {
	t.Parallel()

	custom := roundTripperFunc(func(*http.Request) (*http.Response, error) {
		return nil, errors.New("should not be called")
	})

	_, err := WrapRoundTripper(custom, staticProxySelector{url: &url.URL{Scheme: "http", Host: "127.0.0.1:7897"}})
	if err == nil {
		t.Fatal("expected error")
	}
}

func proxyCopy(dst net.Conn, src net.Conn) {
	_, _ = io.Copy(dst, src)
	_ = dst.Close()
	_ = src.Close()
}

type staticProxySelector struct {
	url *url.URL
}

func (s staticProxySelector) Next(*http.Request) (*url.URL, error) {
	return s.url, nil
}

type roundTripperFunc func(*http.Request) (*http.Response, error)

func (f roundTripperFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}
