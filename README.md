# Package Tracker

## Overview

The `tracker` package is a lightweight web tracking service that logs and tracks user visits to specified paths on a web server. It uses Redis for storing tracking data and allows tracking users based on a generated fingerprint, which is stored in a cookie. The package provides an easy way to track page visits, the number of times a user has visited a specific path, and the user’s IP address and User-Agent.

## Features

- **User Tracking:** Track user visits based on a fingerprint generated using the user's IP address and User-Agent.
- **Path-Specific Visit Counting:** Track the number of visits a user makes to specific paths.
- **Redis Integration:** Store visit data in Redis for efficient retrieval and scalability.
- **Cookie-Based Fingerprinting:** Use HTTP cookies to persist user fingerprints across sessions.

## Installation

To use this package in your project, you need to have Go installed. You can get the package via:

```bash
go get github.com/exhibit-io/tracker
```

Ensure you have Redis installed and running as it is a required dependency.

## Configuration

The `tracker` package requires a configuration object (`config.Config`) to initialize. The configuration should include Redis connection details and cookie settings.

Example configuration structure (`config.Config`):

```go
type Config struct {
    Redis struct {
        Addr     string
        Password string
    }
    Tracker struct {
        CookieConfig TrackerCookieConfig
    }
}

type TrackerCookieConfig struct {
    Name     string
    Domain   string
    Secure   bool
    HttpOnly bool
}
```

## Usage

### Initialization

Before using the tracker, initialize it by calling the `Init` function with the configuration object:

```go
import "github.com/exhibit-io/tracker/config"

func main() {
    config := &config.Config{
        Redis: config.Redis{
            Addr:     "localhost:6379",
            Password: "",
        },
        Tracker: config.Tracker{
            CookieConfig: config.TrackerCookieConfig{
                Name:     "fingerprint",
                Domain:   "example.com",
                Secure:   true,
                HttpOnly: true,
            },
        },
    }

    tracker.Init(config)
}
```

### Handler

The `TrackerHandler` is an HTTP handler function that tracks user visits. You can register this handler with an HTTP router to handle specific routes:

```go
import (
    "github.com/exhibit-io/tracker"
    "github.com/julienschmidt/httprouter"
    "net/http"
)

func main() {
    router := httprouter.New()
    router.GET("/track", tracker.TrackerHandler)
    http.ListenAndServe(":8080", router)
}
```

### Example Request

When a user visits the `/track` endpoint with a query parameter `path`, the tracker logs their visit and responds with details:

```http
GET /track?path=homepage HTTP/1.1
Host: example.com
```

The response will include information such as the user's IP, User-Agent, fingerprint, and the number of visits:

```json
{
    "ip": "192.168.1.1",
    "userAgent": "Mozilla/5.0",
    "fingerprint": "a1b2c3d4e5f6...",
    "path": "homepage",
    "visits": "3"
}
```

## Error Handling

If Redis is not available or there are issues connecting, the `Init` function will log the error and terminate the application.

## Logging

The package logs each tracked visit, including the first seven characters of the fingerprint, the number of visits, and the accessed path.

## License

This project is licensed under the MIT License. See the LICENSE file for details.