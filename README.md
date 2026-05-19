# Vassago SDK

[![Go Report Card](https://goreportcard.com/badge/github.com/aj-nt/vassago-sdk)](https://goreportcard.com/report/github.com/aj-nt/vassago-sdk)
[![License: Apache 2.0](https://img.shields.io/badge/License-Apache%202.0-green.svg)](LICENSE)

Shared Go SDK for the Vassago ecosystem — the gRPC protocol definitions, client libraries, and cross-cutting packages used by vassago, vagent, and gaap.

## What's Inside

| Package | Purpose |
|---------|---------|
| `client/` | gRPC client with connection lifecycle, reconnection, heartbeat |
| `messaging/` | Platform adapters (Discord, Telegram, iMessage, Email) |
| `messaging/format/` | Message formatting and content handling |
| `proto/` | Protobuf definitions + generated Go stubs for all 42 Vassago RPCs |

## Installation

```go
// go.mod
require github.com/aj-nt/vassago-sdk v0.3.0
```

```bash
go get github.com/aj-nt/vassago-sdk@latest
```

## Usage

```go
import "github.com/aj-nt/vassago-sdk/client"

conn, err := client.Dial("localhost:50051")
if err != nil {
    log.Fatal(err)
}
defer conn.Close()

memoryClient := pb.NewMemoryClient(conn)
resp, err := memoryClient.AddMemory(ctx, &pb.AddMemoryRequest{
    Target:   "memory",
    Category: "fact",
    Content:  "Hello, Vassago!",
    Priority: 3,
})
```

## License

Apache 2.0 — see [LICENSE](LICENSE).
