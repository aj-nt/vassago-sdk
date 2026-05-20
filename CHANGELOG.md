# Changelog

All notable changes to the Vassago SDK.

## [Unreleased]

Shared protocol, client, and messaging packages for the Vassago ecosystem. Not yet released.

### Packages

- **gRPC client** — connection lifecycle with reconnection, heartbeat, and optional TLS
- **Proto definitions** — 42 Vassago RPCs across memory, search, sessions, todos, skills, saved tools, cron, replication, pub/sub
- **Messaging adapters** — Discord, Telegram, iMessage (macOS), Email (IMAP/SMTP) with shared message formatting
- **Apache 2.0** — dual-licensed alongside vassago (AGPLv3) and vagent (AGPLv3+Apache 2.0)
- **207 tests** — all passing with `-race` (race detector)
- `suppress_categories` support for `get_hot_block()` client method
- Public repository for cross-repo CI access
