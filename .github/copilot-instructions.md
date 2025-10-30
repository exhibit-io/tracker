# Copilot Instructions for Package Tracker

## Overview

This repository contains a lightweight web tracking service written in Go that logs and tracks user visits. It uses Redis for storage and provides cookie-based user fingerprinting.

## Project Structure

```
.
├── main.go                 # Main application entry point
├── tracker/
│   ├── tracker.go          # Core tracking logic
│   ├── tracker_test.go     # Unit tests
│   └── config/             # Configuration management
│       ├── config.go       # Main config struct
│       ├── env.go          # Environment variable loading
│       ├── redis.go        # Redis configuration
│       ├── tracker.go      # Tracker configuration
│       └── cors.go         # CORS configuration
├── .github/workflows/      # GitHub Actions workflows
└── docker-compose.yml      # Docker setup with Redis
```

## Technology Stack

- **Language**: Go 1.22+
- **Database**: Redis (for tracking data storage)
- **HTTP Router**: julienschmidt/httprouter
- **CORS**: rs/cors
- **Environment**: godotenv for .env file support

## Development Guidelines

### Code Style

- Follow standard Go conventions and formatting (use `gofmt`)
- Write idiomatic Go code
- Keep functions focused and single-purpose
- Use meaningful variable and function names
- Add comments for exported functions and complex logic

### Building and Testing

- **Build**: `go build -v ./...`
- **Test**: `go test ./...`
- **Dependencies**: `go mod tidy`
- **Run locally**: Use `docker-compose up` to start Redis, then run the application

### Architecture Patterns

1. **Configuration**: All configuration is loaded from environment variables via the `config` package
   - Use `.env.example` as a template for local development
   - Configuration is structured and type-safe

2. **HTTP Handlers**: Use httprouter for routing
   - Handler signature: `func(w http.ResponseWriter, r *http.Request, _ httprouter.Params)`
   - Main handler: `GetFingerprintHandler` in `tracker/tracker.go`

3. **Redis Integration**:
   - Redis client is initialized in `tracker.Init()`
   - Keys use format: `tracker:{fingerprint}:{path}`
   - Connection is validated on startup

4. **User Tracking**:
   - Fingerprints are SHA-256 hashes of IP + User-Agent
   - Stored in cookies using configurable cookie settings
   - Visit counts are incremented atomically in Redis

### Adding New Features

When adding new features:

1. Add configuration options to the appropriate config struct (in `tracker/config/`)
2. Update environment variable loading if needed
3. Add corresponding fields to `.env.example`
4. Write unit tests following existing patterns in `tracker_test.go`
5. Update README.md with usage examples if the feature is user-facing
6. Ensure Redis operations use the global `ctx` context

### Testing Guidelines

- Write table-driven tests where appropriate
- Test both success and error cases
- Use descriptive test names (e.g., `TestCreateFingerprintConsistency`)
- Mock Redis interactions for unit tests when necessary
- Ensure tests can run without external dependencies when possible

### Dependencies

- Only add dependencies when necessary
- Run `go mod tidy` after adding/removing dependencies
- Prefer standard library solutions when available
- Keep dependencies up to date but ensure compatibility

### Common Tasks

- **Add new tracking metric**: Modify `GetFingerprintHandler` response and Redis storage
- **Add new configuration**: Add to appropriate config struct and environment loading
- **Add new endpoint**: Register route in `main.go` using the router
- **Update CORS settings**: Modify `config/cors.go` and environment variables

### Error Handling

- Log errors appropriately using the standard `log` package
- Return appropriate HTTP status codes
- Fatal errors (e.g., Redis connection failure) should use `log.Fatal`
- Non-fatal errors should be logged and handled gracefully

### Security Considerations

- Cookie security flags (Secure, HttpOnly) are configurable
- IP address extraction handles proxy headers (X-Forwarded-For, X-Real-IP)
- User-Agent and IP are hashed for fingerprinting
- Validate and sanitize all user inputs

### CI/CD

- All pull requests must pass the `go test ./...` check
- Docker images are built and published automatically
- Tests run on every push to master and on all pull requests

## Getting Started for Contributors

1. Clone the repository
2. Copy `.env.example` to `.env` and configure as needed
3. Start Redis: `docker-compose up -d`
4. Run tests: `go test ./...`
5. Build and run: `go build && ./main` or `go run main.go`

## Questions or Issues?

- Check existing issues and pull requests
- Review the README.md for usage documentation
- Consult Go documentation for standard library usage
