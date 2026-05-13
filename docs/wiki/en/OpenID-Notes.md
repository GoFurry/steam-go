# OpenID Notes

`addons/openid` is used for browser-based Steam sign-in verification.

It confirms the user's Steam identity. It is not a general Web API credential system.

## What It Does

The OpenID addon helps with:

- building the Steam OpenID login URL
- verifying the callback against Steam
- calling `check_authentication`
- recovering `SteamID64`
- preserving and checking `state`

## What It Does Not Do

It does not:

- replace a Steam Web API key
- replace an access token
- fetch profile data automatically
- manage your application session
- store users in your database
- handle frontend UI for you

## Typical Flow

```text
1. Generate random state
2. Store state in a secure cookie or server-side session
3. Redirect user to Steam OpenID login URL
4. Steam redirects back to your callback URL
5. Verify callback with addons/openid
6. Compare returned state with stored state
7. Create your own application session
```

## Example

```bash
go run ./examples/openid
```

With proxy:

```bash
go run ./examples/openid --proxy http://127.0.0.1:7897
```

## Security Notes

- Always verify `state`.
- Do not ask users for Steam username or password.
- OpenID only proves identity; authorization decisions still belong to your application.
- Use HTTPS in production.
- Keep your application session separate from Steam OpenID verification.

