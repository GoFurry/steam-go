# Why Safe Defaults

`WithSafeDefaults()` exists because external API clients should not start with aggressive behavior.

## What It Does

`WithSafeDefaults()` applies a conservative profile:

```text
retry = 2
rate limit = 3 requests/second
burst = 3
```

## Why It Matters

A new SDK user may not immediately know:

- how often their code calls Steam APIs
- whether calls happen concurrently
- whether the service retries on failure
- whether the deployment has multiple instances
- whether the same API key is shared by other jobs

Safe defaults reduce accidental abuse and make first usage more stable.

## When to Use It

Use it for:

- examples
- small services
- CLI tools
- early integration
- tests that touch real external endpoints

```go
client, err := steam.NewClient(
    steam.WithSafeDefaults(),
)
```

## When to Override It

Override it when you have:

- a queue
- metrics
- production traffic estimates
- explicit concurrency controls
- service-level rate budgets

```go
client, err := steam.NewClient(
    steam.WithRetry(2),
    steam.WithRateLimiter(rate.Limit(10), 10),
)
```

## Design Principle

The SDK should make the safe path easy, while still allowing advanced users to tune behavior precisely.

