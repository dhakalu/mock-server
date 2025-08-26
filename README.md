# Mock Server

git branch -M main
git remote add origin https://github.com/dhakalu/mock-server.git
git push -u origin main. Perfect for local development (UI, or integeration with Third Party API) and quick testing when you need to simulate API responses without setting up a full backend.

## Features

- 🚀 Simple HTTP server with graceful shutdown
- 📁 Serves mock responses from static JSON files
- 🔄 Hot-reloadable mock data (just update the JSON files)
- ⚡ Lightweight and fast startup
- 🧪 Ideal for development and testing environments

## Quick Start

### Prerequisites

- Go 1.21+ installed on your system

### Installation

1. Clone or download this repository
2. Navigate to the project directory:
   ```bash
   cd mock-server
   ```

### Running the Server

```bash
go run cmd/mock-server/main.go
```

The server will start on port `8080` by default. If you need to change the port use `MOCK_SERVER_PORT` environment variable.

### Testing

Open your browser or use curl to test:

```bash
curl http://localhost:8080/api/users
```

## Project Structure

```
mock-server/
├── cmd/mock-server/
│   └── main.go              # Server entry point
├── mock-responses/
│   ├── index.json           # Root response
│   └── Components/
│       └── Handlers/
│           └── Get_Visit_Info.ashx.json  # Example API response
├── go.mod                   # Go module definition
├── go.sum                   # Go dependencies
└── README.md               # This file
```

## Mock Responses

Place your JSON mock files in the `mock-responses/` directory. The server will serve these files based on the request path.

### Example Mock File

Create a file at `mock-responses/api/users.json`:

```json
{
  "users": [
    {
      "id": 1,
      "name": "John Doe",
      "email": "john@example.com"
    },
    {
      "id": 2,
      "name": "Jane Smith",
      "email": "jane@example.com"
    }
  ]
}
```

Then access it via: `http://localhost:8080/api/users`

## Configuration

### Port Configuration

The server runs on port `8080` by default. You can change this by modifying `MOCK_SERVER_PORT` environment variable.

### Adding New Mock Endpoints

1. Create a new JSON file in the `mock-responses/` directory
2. Structure the path to match your desired API endpoint
3. The server will automatically serve the file content

## Development

### Building

```bash
go build -o bin/mock-server cmd/mock-server/main.go
```

### Running the Binary

```bash
./bin/mock-server
```

## Use Cases

- **Frontend Development**: Mock backend APIs while developing UI components
- **Integration Testing**: Provide consistent responses for automated tests
- **API Design**: Prototype API responses before implementing the actual backend
- **Demos**: Quick setup for demonstrations without complex infrastructure
- **Offline Development**: Work on projects without internet connectivity to real APIs

## Graceful Shutdown

The server handles `CTRL+C` (SIGINT) gracefully:
- Stops accepting new connections
- Allows existing requests to complete (with 120-second timeout)
- Logs shutdown process

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test thoroughly
5. Submit a pull request

## License

This project is open source. Feel free to use it for your development needs.

## Roadmap

- [ ] Implement request/response logging
- [ ] Add configuration file support
- [ ] Support for dynamic responses with templating
- [ ] Add support for different HTTP methods (POST, PUT, DELETE)
- [ ] Request matching based on headers or query parameters
