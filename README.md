# Bread Information API

This is a simple Go server that provides an API for retrieving and managing information about different types of bread. The API is defined according to the OpenAPI specification in [`openapi.spec.json`](openapi.spec.json).

## Getting Started

### Prerequisites

- Go 1.21.6 or later
- [Wiretap](https://github.com/pb33f/wiretap)
  - demo'd with v0.0.50: in order to install this version, you'll need to build from source
    - ```sh
      git clone git@github.com:pb33f/wiretap.git --branch v0.0.50
      make build
      cp ./bin/wiretap <SOMEWHERE_IN_YOUR_PATH>
      ```
  - you can install the latest with brew, but sometimes it's a little flaky:
    - `brew install pb33f/taps/wiretap`

### Running

1. Run the server:

```sh
go run .
```

2. Start Wiretap by running the following command:

```sh
wiretap
```

If successful, the following service should be running on localhost:

| Service       | Port |
| ------------- | ---- |
| Bread API     | 8081 |
| Wiretap Proxy | 9090 |
| Wiretap UI    | 9091 |

## API Endpoints

The server provides the following endpoint:

- `GET /api/breads`: Get a list of bread types
