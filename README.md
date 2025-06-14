# MultiRobot

A gRPC-based telemetry system for managing multiple robots in a fleet. This project implements the HerdTelemetry service for GOAT Robotics' fleet management system.

## Features

- gRPC server implementation for HerdTelemetry service
- Client library for sending telemetry data
- Support for fleet states and dispatch states telemetry
- Example test client implementation

## Project Structure

```
.
├── client/             # Client library implementation
├── cmd/               # Command-line applications
│   └── test_client/   # Test client implementation
├── proto/             # Protocol buffer definitions
├── server/            # Server implementation
└── main.go           # Main server entry point
```

## Setup

1. Install Go (1.16 or later)
2. Install Protocol Buffers compiler (protoc)
3. Install Go plugins for protoc:
   ```bash
   go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
   go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
   ```

## Building and Running

1. Generate protobuf code:
   ```bash
   protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/v1/fleet_manager/fleet_manager.proto
   ```

2. Run the server:
   ```bash
   go run main.go
   ```

3. Run the test client:
   ```bash
   go run cmd/test_client/main.go
   ```

## License

MIT License 