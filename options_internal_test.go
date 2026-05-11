package steam

import (
	"testing"

	"github.com/GoFurry/steam-go/internal/transport"
	"golang.org/x/time/rate"
)

func TestWithRateLimitSetsRateLimiterConfig(t *testing.T) {
	t.Parallel()

	cfg := defaultClientConfig()
	if err := WithRateLimit(3)(&cfg); err != nil {
		t.Fatalf("WithRateLimit returned error: %v", err)
	}

	want := transport.RateLimiterConfig{
		Limit: rate.Limit(3),
		Burst: 3,
	}
	if cfg.rateLimiter != want {
		t.Fatalf("rateLimiter = %#v, want %#v", cfg.rateLimiter, want)
	}
}

func TestWithRateLimiterDisablesLimiterWithZeroValues(t *testing.T) {
	t.Parallel()

	cfg := defaultClientConfig()
	cfg.rateLimiter = transport.RateLimiterConfig{
		Limit: rate.Limit(5),
		Burst: 5,
	}
	if err := WithRateLimiter(0, 0)(&cfg); err != nil {
		t.Fatalf("WithRateLimiter returned error: %v", err)
	}
	if cfg.rateLimiter != (transport.RateLimiterConfig{}) {
		t.Fatalf("expected limiter to be disabled, got %#v", cfg.rateLimiter)
	}
}

func TestWithRateLimiterRejectsPartialZeroValues(t *testing.T) {
	t.Parallel()

	cfg := defaultClientConfig()
	if err := WithRateLimiter(rate.Limit(1), 0)(&cfg); err == nil {
		t.Fatal("expected error for partial zero limiter config")
	}
	if err := WithRateLimiter(0, 1)(&cfg); err == nil {
		t.Fatal("expected error for partial zero limiter config")
	}
}

func TestLastRateLimiterOptionWins(t *testing.T) {
	t.Parallel()

	cfg := defaultClientConfig()
	if err := WithRateLimit(5)(&cfg); err != nil {
		t.Fatalf("WithRateLimit returned error: %v", err)
	}
	if err := WithRateLimiter(rate.Limit(2), 7)(&cfg); err != nil {
		t.Fatalf("WithRateLimiter returned error: %v", err)
	}

	want := transport.RateLimiterConfig{
		Limit: rate.Limit(2),
		Burst: 7,
	}
	if cfg.rateLimiter != want {
		t.Fatalf("rateLimiter = %#v, want %#v", cfg.rateLimiter, want)
	}
}

func TestWithRateLimitCanDisableExplicitLimiter(t *testing.T) {
	t.Parallel()

	cfg := defaultClientConfig()
	if err := WithRateLimiter(rate.Limit(2), 7)(&cfg); err != nil {
		t.Fatalf("WithRateLimiter returned error: %v", err)
	}
	if err := WithRateLimit(0)(&cfg); err != nil {
		t.Fatalf("WithRateLimit returned error: %v", err)
	}
	if cfg.rateLimiter != (transport.RateLimiterConfig{}) {
		t.Fatalf("expected limiter to be disabled, got %#v", cfg.rateLimiter)
	}
}
