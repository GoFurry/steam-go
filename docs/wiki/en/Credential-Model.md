# Credential Model

`steam-go` separates Steam credentials instead of treating them as one generic token.

This makes the SDK safer and easier to reason about.

## API Key

A Steam Web API key is used for endpoints that require a `key` parameter.

```go
client, err := steam.NewClient(
    steam.WithAPIKey("your-key"),
)
```

## Rotating API Keys

```go
client, err := steam.NewClient(
    steam.WithAPIKeys("key-a", "key-b", "key-c"),
)
```

Use this to spread traffic across keys that your application legitimately owns.

## Health-Checked API Keys

```go
client, err := steam.NewClient(
    steam.WithHealthCheckedAPIKeys(
        steam.DefaultAPIKeyHealthConfig(),
        "key-a",
        "key-b",
    ),
)
```

When a key repeatedly gets `401` or `429`, it can be temporarily cooled down.

## Access Token

An access token is separate from a Web API key.

```go
client, err := steam.NewClient(
    steam.WithAccessToken("access-token"),
)
```

Use it only for endpoints or workflows that require token-based access.

## OpenID

OpenID is used for browser-based Steam sign-in.

It verifies a user's Steam identity and returns a SteamID64.

OpenID does not replace:

- Steam Web API keys
- access tokens
- application sessions
- profile-fetching APIs

## Backend-Only Credentials

Do not expose API keys or access tokens to frontend JavaScript.

Recommended pattern:

```text
Browser -> Your Backend -> steam-go -> Steam
```

Avoid:

```text
Browser -> Steam with raw API key
```

## Logging Safety

Steam credentials are often passed through query parameters.

Before logging URLs:

```go
safeURL := steam.RedactSensitiveURL(rawURL)
```

## Practical Rule

Use the weakest credential that can complete the job:

| Need | Credential |
|---|---|
| Public Web API data requiring key | API Key |
| User-authorized protected action | Access Token |
| Browser sign-in identity | OpenID |
| Server-side publisher operations | Publisher Web API Key |

## References

- [Steam Web API Overview](https://partner.steamgames.com/doc/webapi_overview?language=english)
- [Steam Web API Terms of Use](https://steamcommunity.com/dev/apiterms)
