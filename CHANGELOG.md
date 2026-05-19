# Changelog

All notable changes to the Vassago SDK.

## [Unreleased]

Shared protocol, client, and messaging packages for the Vassago ecosystem. Not yet released.

### Packages

- **gRPC client** — connection lifecycle with reconnection, heartbeat, and optional TLS
- **Proto definitions** — 42 Vassago RPCs across memory, search, sessions, todos, skills, saved tools, cron, replication, pub/sub
- **Messaging adapters** — Discord, Telegram, iMessage (macOS), Email (IMAP/SMTP) with shared message formatting
- **Apache 2.0** — dual-licensed alongside vassago (AGPLv3) and vagent (AGPLv3+Apache 2.0)

## [0.3.0] - 2026-05-19

- Added `suppress_categories` to `get_hot_block()` client method

## [0.2.0] - 2026-05-19

- Public release — repository made public for cross-repo CI access

## [0.1.0] - 2026-05-14

- Initial release: proto definitions, gRPC client, messaging adapters
