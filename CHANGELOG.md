# Changelog

All notable changes to `steam-go` will be documented in this file.

The format is intentionally simple during the pre-v1 release phase.

## Unreleased

- no new API coverage is planned before `v1.1.0`

## v1.0.0

### Stable release summary

- first stable `steam-go` release
- stable root `Client` and grouped `client.API.*` access pattern
- stable functional option system for credentials, timeout, retry, rate limit, proxy, and traffic policy configuration
- production-oriented request controls including safe defaults, retry hardening, proxy routing, sticky proxy support, health-checked proxy pools, and traffic-class isolation
- stable addon entrypoints for `addons/openid` and `addons/a2s`
- published compatibility, endpoint stability, endpoint coverage, and release-governance documentation
- live smoke validation consolidated under `examples/live/`

## v1.0.0-alpha-3

### Added

- per-traffic-class request policy routing
- public store-page request profiles and Referer strategies
- short in-memory caching with conditional revalidation
- block detection for public store-page traffic
- host and session request controls
- sticky proxy selection, proxy health checks, and proxy metrics snapshots
- per-class transport hook foundations

### Improved

- safer retry and rate-limit configuration
- cookie jar based session persistence
- proxy and traffic policy documentation
- OpenID hardening, proxy support, and test coverage
- CI checks including race, vet, staticcheck, and govulncheck

## v1.0.0-alpha-2

### Added

- `WishlistService`
- expanded `PlayerService` coverage
- proxy selection helpers and request-layer hardening

### Improved

- retry behavior
- verifier body size limits
- safer default timeouts
- request safety around large or volatile responses

## v1.0.0-alpha-1

### Added

- initial root `Client` and grouped `client.API.*` structure
- functional option based client configuration
- initial typed Steam Web API support
- addon foundations for `openid` and `a2s`
