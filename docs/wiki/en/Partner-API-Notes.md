# Partner API Notes

Steam provides a partner-only Web API host for secure publisher-server use cases.

This page is mainly for future `steam-go` design and advanced users.

## Public Host vs Partner Host

| Host | Use Case |
|---|---|
| `api.steampowered.com` | Public Steam Web API |
| `partner.steam-api.com` | Partner-only secure server requests |

## Partner Host Properties

The partner host has different operational rules:

- It is only accessible over HTTPS.
- It is intended for secure publisher servers.
- Every request must include a valid publisher Web API key.
- Requests without a valid publisher key return `403`.
- Repeated `403` responses may trigger strict rate limits for the connecting IP.
- It is not the same as a normal public Web API key workflow.

## IP Whitelisting

Steam supports IP whitelisting for Web API keys.

Once a whitelist is configured, requests from non-whitelisted addresses are blocked with `403`.

This is an extra security layer, not a replacement for key protection.

## SDK Design Implications

Future partner-oriented features should consider:

- dedicated base URL support
- explicit publisher key handling
- stricter error handling for `403`
- safer logs and redaction
- configuration examples for server-side use only

## Practical Advice

- Never expose publisher keys to clients.
- Prefer backend-only access.
- Be careful when testing with the wrong key type.
- Treat repeated `403` as a configuration error, not something to retry aggressively.

## References

- [Steam Web API Overview](https://partner.steamgames.com/doc/webapi_overview?language=english)
- [Steam Web API Terms of Use](https://steamcommunity.com/dev/apiterms)
