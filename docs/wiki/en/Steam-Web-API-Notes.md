# Steam Web API Notes

This page records practical notes about using the Steam Web API in real-world applications.

It is not a full copy of the official Steamworks documentation. Instead, it highlights the behaviors that matter when designing an SDK or production integration.

## Public and Protected APIs

The Steam Web API contains both public methods and protected methods.

Public methods can be called by applications that can make HTTP requests. Protected methods require authentication and are intended for trusted backend applications.

## Public API Host

Public Steam Web API requests are normally sent to:

```text
api.steampowered.com
```

The usual request shape is:

```text
https://api.steampowered.com/<interface>/<method>/v<version>/
```

## Request Format

Steam Web API methods commonly accept GET or POST parameters.

For POST requests, the official documentation expects form URL encoded bodies and UTF-8 text.

## Service Interfaces and input_json

Some Steam Web API interfaces are service interfaces.

If an interface name ends with `Service`, such as `IPlayerService`, it may support passing arguments as a single JSON blob using `input_json`.

Example shape:

```text
?key=XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX&input_json={...}
```

The JSON value must be URL encoded.

## SteamID64

Steam Web API identifies users through a 64-bit Steam ID.

For browser-based login flows, use Steam OpenID verification to obtain the user's Steam identity. Do not ask users for Steam credentials directly.

## Practical SDK Advice

- Keep typed API methods focused and predictable.
- Keep raw response methods available when Steam payloads are large or unstable.
- Prefer backend-side credential injection.
- Treat public Store pages as different from official Web API JSON endpoints.
- Avoid assuming undocumented per-second limits.
- Use conservative rate limiting and caching in production.

## References

- [Steam Web API Overview](https://partner.steamgames.com/doc/webapi_overview?language=english)
- [Steam Web API Terms of Use](https://steamcommunity.com/dev/apiterms)
