# Contributing to Vassago SDK

The shared SDK for the Vassago ecosystem — used by vassago, vagent, and gaap.

## Quick Start

```bash
go test -race ./...
go vet ./...
```

## Packages

- `client/` — gRPC client with reconnection and heartbeat
- `messaging/` — platform adapters (Discord, Telegram, iMessage, Email)
- `messaging/format/` — message formatting
- `proto/` — protobuf definitions + generated stubs

## Proto Generation

```bash
make proto
```

Generated `.pb.go` files are tracked in git. The proto source of truth lives here — vassago and all consumers pull from this repo.

## Code Style

Standard Go conventions. `gofmt` before committing.

## License

Apache 2.0. Contributions under the same license.
