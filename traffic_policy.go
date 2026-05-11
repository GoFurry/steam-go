package steam

import (
	"context"
	"fmt"
	"net/http"

	"github.com/GoFurry/steam-go/internal/request"
	itraffic "github.com/GoFurry/steam-go/internal/traffic"
	"golang.org/x/time/rate"
)

// TrafficClass identifies one request traffic category.
type TrafficClass = itraffic.Class

const (
	TrafficClassOfficialAPI     TrafficClass = itraffic.ClassOfficialAPI
	TrafficClassPublicStorePage TrafficClass = itraffic.ClassPublicStorePage
)

// RetryBackoffConfig exposes the SDK retry backoff shape for policy overrides.
type RetryBackoffConfig = request.RetryBackoffConfig

// DefaultRetryBackoffConfig returns the SDK retry backoff defaults.
func DefaultRetryBackoffConfig() RetryBackoffConfig {
	return request.DefaultRetryBackoffConfig()
}

// TrafficRateLimiterPolicy overrides per-class token-bucket settings.
type TrafficRateLimiterPolicy struct {
	Limit rate.Limit
	Burst int
}

// TrafficRetryPolicy overrides per-class retry behavior.
type TrafficRetryPolicy struct {
	Retry   int
	Backoff RetryBackoffConfig
}

// TrafficPolicy overrides selected request behavior for one traffic class.
type TrafficPolicy struct {
	ProxySelector ProxySelector
	CookieJar     http.CookieJar
	RateLimiter   *TrafficRateLimiterPolicy
	Retry         *TrafficRetryPolicy
}

// WithTrafficClass attaches one traffic class to a request context.
func WithTrafficClass(ctx context.Context, class TrafficClass) context.Context {
	return itraffic.WithClass(ctx, class)
}

type trafficPolicyConfig struct {
	proxySelector     ProxySelector
	cookieJar         http.CookieJar
	rateLimiter       *TrafficRateLimiterPolicy
	retry             *TrafficRetryPolicy
	cookieJarProvided bool
}

// WithTrafficPolicy configures one per-class request policy override.
func WithTrafficPolicy(class TrafficClass, policy TrafficPolicy) Option {
	return func(cfg *clientConfig) error {
		if !supportedTrafficClass(class) {
			return fmt.Errorf("unsupported traffic class")
		}
		class = normalizeTrafficClass(class)
		if policy.RateLimiter != nil {
			if policy.RateLimiter.Limit < 0 {
				return fmt.Errorf("traffic policy rate limit must not be negative")
			}
			if policy.RateLimiter.Burst < 0 {
				return fmt.Errorf("traffic policy rate limiter burst must not be negative")
			}
			if policy.RateLimiter.Limit == 0 || policy.RateLimiter.Burst == 0 {
				if policy.RateLimiter.Limit != 0 || policy.RateLimiter.Burst != 0 {
					return fmt.Errorf("traffic policy rate limiter limit and burst must both be zero to disable")
				}
			}
		}
		if policy.Retry != nil {
			if policy.Retry.Retry < 0 {
				return fmt.Errorf("traffic policy retry must not be negative")
			}
			if policy.Retry.Backoff.BaseDelay <= 0 {
				return fmt.Errorf("traffic policy retry base delay must be greater than zero")
			}
			if policy.Retry.Backoff.MaxDelay <= 0 {
				return fmt.Errorf("traffic policy retry max delay must be greater than zero")
			}
			if policy.Retry.Backoff.MaxDelay < policy.Retry.Backoff.BaseDelay {
				return fmt.Errorf("traffic policy retry max delay must be greater than or equal to base delay")
			}
		}
		if cfg.trafficPolicies == nil {
			cfg.trafficPolicies = make(map[TrafficClass]trafficPolicyConfig)
		}
		cfg.trafficPolicies[class] = trafficPolicyConfig{
			proxySelector:     policy.ProxySelector,
			cookieJar:         policy.CookieJar,
			rateLimiter:       policy.RateLimiter,
			retry:             policy.Retry,
			cookieJarProvided: policy.CookieJar != nil,
		}
		return nil
	}
}

func normalizeTrafficClass(class TrafficClass) TrafficClass {
	return itraffic.NormalizeClass(class)
}

func supportedTrafficClass(class TrafficClass) bool {
	switch class {
	case TrafficClassOfficialAPI, TrafficClassPublicStorePage:
		return true
	default:
		return false
	}
}
