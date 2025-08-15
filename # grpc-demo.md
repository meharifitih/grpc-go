# grpc-demo

A Go-based demonstration of gRPC with unary, server streaming, client streaming, and bidirectional streaming RPCs.

## Project Structure

```
go.mod
go.sum
client/
    bi_stream.go
    client_stream.go
    main.go
    servedr_steam.go
    unary.go
proto/
    greet_grpc.pb.go
    greet.pb.go
    greet.proto
server/
    bi_stream.go
    client_steam.go
    main.go
    server_stream.go
    unary.go
```

## Prerequisites

- Go 1.24+
- [protoc](https://grpc.io/docs/protoc-installation/) (Protocol Buffers compiler)
- `protoc-gen-go` and `protoc-gen-go-grpc` plugins

## Setup

1. **Install dependencies:**

   ```sh
   go mod tidy
   ```

2. **(Optional) Regenerate gRPC code:**

   If you modify `proto/greet.proto`, regenerate Go code:

   ```sh
   protoc --go_out=. --go-grpc_out=. proto/greet.proto
   ```

## Running the Server

```sh
cd server
go run main.go unary.go server_stream.go client_steam.go bi_stream.go
```

The server listens on `:8080`.

## Running the Client

Open a new terminal:

```sh
cd client
go run main.go unary.go servedr_steam.go client_stream.go bi_stream.go
```

Edit `client/main.go` to uncomment the function call for the RPC you want to test (unary, server streaming, client streaming, or bidirectional streaming).

## Features

- **Unary RPC:** Simple request/response (`SayHello`)
- **Server Streaming RPC:** Server sends a stream of responses (`SayHelloServerStreaming`)
- **Client Streaming RPC:** Client sends a stream of requests (`SayHelloClientStreaming`)
- **Bidirectional Streaming RPC:** Both client and server stream messages (`SatHelloBiDirectional`)

## Proto Definition

See [`proto/greet.proto`](proto/greet.proto) for service and message definitions.