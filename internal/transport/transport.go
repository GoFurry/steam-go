package transport

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/time/rate"
)

// ProxySelector chooses a proxy for a request.
type ProxySelector interface {
	Next(req *http.Request) (*url.URL, error)
}

// RateLimiterConfig defines one token-bucket limiter.
type RateLimiterConfig struct {
	Limit rate.Limit
	Burst int
}

// Client applies retry and rate limiting on top of an http.Client.
type Client struct {
	httpClient *http.Client
	limiter    *rate.Limiter
}

// New creates a transport client.
func New(httpClient *http.Client, cfg RateLimiterConfig) *Client {
	var limiter *rate.Limiter
	if cfg.Limit > 0 && cfg.Burst > 0 {
		limiter = rate.NewLimiter(cfg.Limit, cfg.Burst)
	}
	return &Client{
		httpClient: httpClient,
		limiter:    limiter,
	}
}

// Do executes a single request with optional rate limiting.
func (c *Client) Do(ctx context.Context, req *http.Request) (*http.Response, error) {
	if c.limiter != nil {
		if err := c.limiter.Wait(ctx); err != nil {
			return nil, err
		}
	}
	return c.httpClient.Do(req.Clone(ctx))
}

// WrapRoundTripper installs proxy selection on top of an existing transport.
func WrapRoundTripper(rt http.RoundTripper, selector ProxySelector) (http.RoundTripper, error) {
	if selector == nil {
		if rt != nil {
			return rt, nil
		}
		return defaultTransport(), nil
	}

	base, err := baseTransport(rt)
	if err != nil {
		return nil, err
	}
	base.Proxy = func(req *http.Request) (*url.URL, error) {
		return selector.Next(req)
	}
	return base, nil
}

func baseTransport(rt http.RoundTripper) (*http.Transport, error) {
	if rt == nil {
		return defaultTransport(), nil
	}
	if transport, ok := rt.(*http.Transport); ok {
		return transport.Clone(), nil
	}
	return nil, fmt.Errorf("proxy selector requires an *http.Transport or nil transport, got %T", rt)
}

func defaultTransport() *http.Transport {
	if base, ok := http.DefaultTransport.(*http.Transport); ok {
		return base.Clone()
	}
	return &http.Transport{
		Proxy:                 http.ProxyFromEnvironment,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
}
