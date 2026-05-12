# Rate Limiting Strategy

`steam-go` uses conservative request controls because Steam Web API usage has both documented and practical limits.

## Why Rate Limiting Exists

The Steam Web API Terms of Use limit applications to 100,000 calls per day.

In addition to documented daily limits, production callers should avoid aggressive bursts. External services may return `429`, degrade temporarily, or block abusive-looking traffic.

## SDK Strategy

`steam-go` provides multiple layers instead of one hardcoded limit:

- `WithSafeDefaults()`
- `WithRateLimit(...)`
- `WithRateLimiter(...)`
- `WithRetry(...)`
- `WithRetryBackoff(...)`
- `WithRetryRespectRetryAfter(...)`
- `WithHealthCheckedAPIKeys(...)`
- `TrafficRateLimiterPolicy`
- `HostControl`
- `SessionControl`

## Safe Default

For normal external traffic:

```go
client, err := steam.NewClient(
    steam.WithSafeDefaults(),
)
```

This is intentionally conservative and easy to override later.

## Custom Rate Limiter

For heavier workloads:

```go
client, err := steam.NewClient(
    steam.WithRetry(2),
    steam.WithRateLimiter(rate.Limit(5), 5),
    steam.WithRetryRespectRetryAfter(true),
)
```

## Traffic-Class Rate Limit

For public Store-page-like traffic:

```go
client, err := steam.NewClient(
    steam.WithTrafficPolicy(
        steam.TrafficClassPublicStorePage,
        steam.TrafficPolicy{
            RateLimiter: &steam.TrafficRateLimiterPolicy{
                Limit: 2,
                Burst: 2,
            },
        },
    ),
)
```

## Practical Advice

- Do not run unbounded goroutines against Steam APIs.
- Treat `429` as a signal to slow down.
- Prefer cache for repeated reads.
- Use lower concurrency for public Store pages.
- Use key rotation for resilience, not for bypassing policy.
- Avoid documenting unofficial fixed per-second limits as facts.
- Put limits near the SDK boundary, not deep inside business logic.

## Suggested Production Defaults

| Scenario | Suggested Strategy |
|---|---|
| Small tool / CLI | `WithSafeDefaults()` |
| Backend service | `WithRetry(2)` + explicit rate limiter |
| Public Store page access | Low RPS + cache + block detection |
| Large batch job | Queue + worker pool + global limiter |
| Multi-key setup | Health-checked key provider + retry |

## References

- [Steam Web API Overview](https://partner.steamgames.com/doc/webapi_overview?language=english)
- [Steam Web API Terms of Use](https://steamcommunity.com/dev/apiterms)
