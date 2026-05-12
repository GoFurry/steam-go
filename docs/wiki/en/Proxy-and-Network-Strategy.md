# Proxy and Network Strategy

`steam-go` models proxy usage as a selector instead of a single static proxy URL.

This makes proxy behavior composable, testable, and suitable for different Steam traffic types.

## Why ProxySelector Exists

Different Steam traffic may need different network routes:

- Official Web API traffic
- Steam Community OpenID verification
- Public Store page traffic
- Region-sensitive requests
- Session-like browser flows
- Internal testing or debugging traffic

## Proxy Types

| Selector | Use Case |
|---|---|
| Static Proxy | Send all requests through one proxy |
| Round-Robin Proxy | Rotate across multiple proxies |
| Health-Checked Proxy | Temporarily avoid failing proxies |
| Sticky Proxy | Keep one session on the same proxy |
| Routing Proxy | Route by host and path |

## Static Proxy

```go
selector, err := steam.NewStaticProxySelector("http://127.0.0.1:7897")

client, err := steam.NewClient(
    steam.WithProxySelector(selector),
)
```

## Round-Robin Proxy

```go
selector, err := steam.NewRoundRobinProxySelector(
    "http://127.0.0.1:7897",
    "http://127.0.0.1:7898",
)
```

## Health-Checked Proxy

```go
selector, err := steam.NewHealthCheckedRoundRobinProxySelector(
    steam.DefaultProxyHealthConfig(),
    "http://127.0.0.1:7897",
    "http://127.0.0.1:7898",
)

metrics := selector.(steam.ProxyMetricsProvider).ProxyMetricsSnapshot()
```

Use this when proxies may fail independently.

## Sticky Proxy

```go
base, _ := steam.NewRoundRobinProxySelector(
    "http://127.0.0.1:7897",
    "http://127.0.0.1:7898",
)

selector := steam.NewStickyProxySelector(base)

ctx := steam.WithProxySessionKey(context.Background(), "session-1")
```

Use this when a browser-like or login-like flow should stay on one proxy.

## Routing Proxy

```go
selector, err := steam.NewRoutingProxySelector(
    steam.ProxyRoute{
        Host:       "api.steampowered.com",
        PathPrefix: "/ISteamUser/",
        ProxyURL:   "http://127.0.0.1:7897",
    },
    steam.ProxyRoute{
        Host:       "steamcommunity.com",
        PathPrefix: "/openid/",
        ProxyURL:   "",
    },
)
```

An empty `ProxyURL` means direct connection.

## Practical Advice

- Prefer one global proxy policy instead of scattered custom transports.
- Use sticky proxy only when session affinity matters.
- Use health checking when proxy availability is unstable.
- Use routing when official API traffic and OpenID / Store traffic need different network paths.
- Do not hide network failures by retrying forever.

