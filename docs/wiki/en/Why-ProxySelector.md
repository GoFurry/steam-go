# Why ProxySelector

`ProxySelector` exists because a single proxy URL is not enough for real-world Steam integrations.

## The Simple Approach

A basic client might support only:

```go
http.ProxyURL(proxyURL)
```

This works for small cases, but it cannot express:

- rotating proxies
- health checks
- cooldown after failures
- session affinity
- routing by host and path
- direct connection fallback
- per-traffic-class proxy policy

## The steam-go Approach

`steam-go` uses:

```go
type ProxySelector interface {
    Next(req *http.Request) (*url.URL, error)
}
```

The selector receives the request, so it can make decisions based on:

- request host
- request path
- request context
- session key
- traffic class
- proxy health state

## What This Enables

| Feature | Why It Needs a Selector |
|---|---|
| Round-robin | Needs internal state |
| Health check | Needs failure memory |
| Sticky proxy | Needs session key |
| Routing proxy | Needs request host/path |
| Direct fallback | Needs route-level decision |
| Metrics | Needs selection and result tracking |

## Example

```go
base, _ := steam.NewRoundRobinProxySelector(
    "http://127.0.0.1:7897",
    "http://127.0.0.1:7898",
)

selector := steam.NewStickyProxySelector(base)

client, err := steam.NewClient(
    steam.WithProxySelector(selector),
)
```

## Design Principle

Proxy behavior is request policy, not just static configuration.

