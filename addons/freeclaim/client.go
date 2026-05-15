package freeclaim

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gofurry/steam-go/web/storefront"
)

// Client is safe for concurrent use.
type Client struct {
	storefront           *storefront.Service
	httpClient           *http.Client
	timeout              time.Duration
	maxResponseBodyBytes int64
	storeBaseURL         *url.URL
}

type CookieJarProvider interface {
	CookieJar(ctx context.Context) (http.CookieJar, error)
}

type StaticCookieJarProvider struct {
	Jar http.CookieJar
}

func NewStaticCookieJarProvider(jar http.CookieJar) StaticCookieJarProvider {
	return StaticCookieJarProvider{Jar: jar}
}

func (p StaticCookieJarProvider) CookieJar(_ context.Context) (http.CookieJar, error) {
	if p.Jar == nil {
		return nil, &Error{Code: ErrorCodeConfig, Op: "cookie_jar", Message: "cookie jar must not be nil"}
	}
	return p.Jar, nil
}

type httpResult struct {
	StatusCode int
	FinalURL   *url.URL
	Header     http.Header
	Body       []byte
}

func NewClient(storefrontService *storefront.Service, opts ...Option) (*Client, error) {
	if storefrontService == nil {
		return nil, configError("new_client", "storefront service must not be nil", nil)
	}
	options, err := defaultClientOptions()
	if err != nil {
		return nil, configError("new_client", "invalid default options", err)
	}
	for _, opt := range opts {
		if opt == nil {
			continue
		}
		if err := opt(&options); err != nil {
			return nil, err
		}
	}
	return &Client{
		storefront:           storefrontService,
		httpClient:           options.httpClient,
		timeout:              options.timeout,
		maxResponseBodyBytes: options.maxResponseBodyBytes,
		storeBaseURL:         cloneURL(options.storeBaseURL),
	}, nil
}

func (c *Client) withTimeout(ctx context.Context) (context.Context, context.CancelFunc) {
	if ctx == nil {
		ctx = context.Background()
	}
	if c.timeout <= 0 {
		return context.WithCancel(ctx)
	}
	return context.WithTimeout(ctx, c.timeout)
}

func (c *Client) doRequestWithJar(ctx context.Context, jar http.CookieJar, method, rawURL string, body io.Reader, contentType string, headers map[string]string, op string) (httpResult, error) {
	reqCtx, cancel := c.withTimeout(ctx)
	defer cancel()

	client := cloneHTTPClient(c.httpClient)
	client.Jar = jar

	req, err := http.NewRequestWithContext(reqCtx, method, rawURL, body)
	if err != nil {
		return httpResult{}, &Error{Code: ErrorCodeRequestBuild, Op: op, Message: "build request failed", Err: err}
	}
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return httpResult{}, &Error{Code: ErrorCodeTransport, Op: op, Message: "request failed", Err: err}
	}
	defer resp.Body.Close()

	responseBody, readErr := readBodyLimited(resp.Body, c.maxResponseBodyBytes)
	if readErr != nil {
		return httpResult{}, &Error{Code: ErrorCodeTransport, Op: op, Message: "read response failed", Err: readErr}
	}

	result := httpResult{
		StatusCode: resp.StatusCode,
		Header:     resp.Header.Clone(),
		Body:       responseBody,
	}
	if resp.Request != nil && resp.Request.URL != nil {
		result.FinalURL = cloneURL(resp.Request.URL)
	}
	return result, nil
}

func (c *Client) resolveProviderJar(ctx context.Context, provider CookieJarProvider, op string) (http.CookieJar, error) {
	if provider == nil {
		return nil, &Error{Code: ErrorCodeRequestBuild, Op: op, Message: "cookie jar provider must not be nil"}
	}
	jar, err := provider.CookieJar(ctx)
	if err != nil {
		return nil, &Error{Code: ErrorCodeConfig, Op: op, Message: "resolve cookie jar failed", Err: err}
	}
	if jar == nil {
		return nil, &Error{Code: ErrorCodeConfig, Op: op, Message: "cookie jar must not be nil"}
	}
	return jar, nil
}

func readBodyLimited(r io.Reader, maxBytes int64) ([]byte, error) {
	if maxBytes <= 0 {
		return io.ReadAll(r)
	}
	reader := &io.LimitedReader{R: r, N: maxBytes + 1}
	body, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	if int64(len(body)) > maxBytes {
		return nil, fmt.Errorf("response body exceeds limit of %d bytes", maxBytes)
	}
	return body, nil
}

func parseAbsoluteURL(raw string) (*url.URL, error) {
	if strings.TrimSpace(raw) == "" {
		return nil, fmt.Errorf("url must not be empty")
	}
	parsed, err := url.Parse(raw)
	if err != nil {
		return nil, err
	}
	if parsed.Scheme == "" || parsed.Host == "" {
		return nil, fmt.Errorf("url must be absolute")
	}
	return parsed, nil
}

func cloneURL(u *url.URL) *url.URL {
	if u == nil {
		return nil
	}
	cloned := *u
	return &cloned
}

func resolveURL(base *url.URL, path string) *url.URL {
	if base == nil {
		return &url.URL{Path: path}
	}
	ref := &url.URL{Path: path}
	return cloneURL(base).ResolveReference(ref)
}

func cloneHTTPClient(base *http.Client) *http.Client {
	if base == nil {
		return &http.Client{}
	}
	cloned := *base
	return &cloned
}
