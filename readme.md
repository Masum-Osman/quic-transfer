# README.md
# QUIC File Transfer

This project implements a file transfer service using the QUIC protocol in Go. It consists of a client and server application that can securely transfer files over a network.

## Features

- QUIC protocol for efficient and secure file transfers
- Configurable client and server
- Support for sending and receiving files of any size
- TLS encryption for secure communication

## Prerequisites

- Go 1.20 or later
- `github.com/quic-go/quic-go` library
- `github.com/spf13/viper` for configuration management

## Installation

1. Clone the repository:
   ```
   git clone https://github.com/masum-osman/quic-transfer.git
   cd quic-transfer
   ```

2. Install dependencies:
   ```
   go mod tidy
   ```

## Configuration

Create configuration files for the client and server in the `config` directory:

### config/server.yaml
```yaml
port: 8080
filePath: "/path/to/save/received/files"
```

### config/client.yaml
```yaml
host: "localhost"
port: 8080
filePath: "/path/to/file/to/send"
destination: "optional_new_filename"
```

## Usage

### Running the server

```
go run cmd/server/main.go
```

### Running the client

```
go run cmd/client/main.go
```

## Project Structure

```
.
├── cmd
│   ├── client
│   │   └── main.go
│   └── server
│       └── main.go
├── internal
│   ├── client
│   │   └── client.go
│   ├── config
│   │   └── config.go
│   ├── protocol
│   │   └── protocol.go
│   ├── server
│   │   └── server.go
│   └── utils
│       └── utils.go
├── go.mod
└── README.md
```

## License

This project is licensed under the MIT License - see the LICENSE file for details.