# Why Traffic Classes

Traffic classes exist because not all Steam-related HTTP requests are the same.

## The Problem

A simple SDK may treat all requests as one category:

```text
one client
one retry policy
one proxy
one rate limit
one cookie jar
```

That is easy at first, but it becomes limiting when the SDK grows.

## Different Traffic Has Different Needs

| Traffic | Needs |
|---|---|
| Official Web API | keys, typed JSON, retry, rate limit |
| Public Store page | browser-like headers, Referer, cache, block detection |
| OpenID verification | callback safety, possible proxy, strict state validation |
| Future session-like flows | cookies, sticky proxy, low concurrency |

## SDK Design

`steam-go` uses explicit traffic classes:

```go
steam.TrafficClassOfficialAPI
steam.TrafficClassPublicStorePage
```

Then `TrafficPolicy` can configure behavior per class.

## Benefits

- Avoids mixing API and browser-like request behavior.
- Keeps business API methods clean.
- Makes future Store-page support easier.
- Allows safer defaults per traffic type.
- Makes proxy, cookie, and cache behavior explicit.

## Example

```go
client, err := steam.NewClient(
    steam.WithTrafficPolicy(
        steam.TrafficClassPublicStorePage,
        steam.TrafficPolicy{
            Cache: &steam.TrafficCachePolicy{
                TTL: time.Minute,
            },
        },
    ),
)
```

## Design Principle

Classify traffic first, then apply the right policy.

