# Error Handling

`steam-go` uses `*steam.APIError` to classify SDK errors.

The goal is to make failures easy to inspect without parsing error strings.

## Error Kinds

| Kind | Meaning |
|---|---|
| `request_build` | Failed to build request |
| `transport` | Network or transport failure |
| `http_status` | Non-success HTTP status |
| `decode` | Failed to decode response |
| `api_response` | Steam API returned an application-level error |

## Example

```go
var apiErr *steam.APIError
if errors.As(err, &apiErr) {
    fmt.Println(apiErr.Kind)
    fmt.Println(apiErr.StatusCode)
    fmt.Println(apiErr.BodyPreview)
}
```

## What to Log

Useful fields:

- error kind
- HTTP status code
- request method
- redacted URL
- short body preview
- retry count if available
- traffic class if available

## What Not to Log

Avoid logging:

- raw API key
- raw access token
- full request URL before redaction
- full response body when it may contain sensitive or huge content

## URL Redaction

```go
safeURL := steam.RedactSensitiveURL(rawURL)
```

## Recommended Handling

| Error Kind | Suggested Response |
|---|---|
| `request_build` | Fix caller input or SDK bug |
| `transport` | Retry if safe |
| `http_status` | Check status, rate limit, auth, or block |
| `decode` | Inspect raw response or update model |
| `api_response` | Handle Steam-level error semantics |

