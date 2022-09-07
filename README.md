# smocker-client

[![GoDoc](https://godoc.org/github.com/churmd/smockerclient?status.svg)](https://pkg.go.dev/github.com/churmd/smockerclient)

A go client to manage [Smocker](https://smocker.dev/) servers.

Example:

```go
package main

import (
	"net/http"

	"github.com/churmd/smockerclient"
	"github.com/churmd/smockerclient/mock"
)

func main() {
  instance := smockerclient.DefaultInstance()

  // Clear any old sessions and mocks
  _ = instance.ResetAllSessionsAndMocks()

  // Start a new session for your new mocks
  _ = instance.StartSession("SmockerClientSession")

  // Add a healthcheck mock
  request := mock.NewRequestBuilder(http.MethodGet, "/healthcheck").
    AddHeader("Accept", "application/json").
    Build()

  response := mock.NewResponseBuilder(http.StatusOK).AddBody(`{"status": "OK"}`).Build()

  mockDefinition := mock.NewDefinition(request, response)

  _ = instance.AddMock(mockDefinition)
}
```

## Functions

* `ResetAllSessionsAndMocks` - Clears the Smocker server of all sessions and mocks. Leaving it in a clean state.
* `StartSession` - Starts a new session on the Smocker server with the given name. New mocks will be added to the latest
  session started.
* `AddMock` - Adds a new mock to the latest session on the Smocker server. Mocks can be made using the provided builders
  or raw json option detailed below.

## Mock Definitions

### Builders

Builders for the request and response part of the mock definition are provided in the `mock` package. These allow mocks
to be setup in a more programmatic way.

```go
request := mock.NewRequestBuilder(http.MethodPut, "/foo/bar").
    AddQueryParam("limit", "10").
    AddQueryParam("filters", "red", "green").
    AddHeader("Content-Type", "application/json", "application/vnd.api+json").
    AddBearerAuthToken("sv2361fr1o8ph3oin").
    AddJsonBody(`{"example": "body"`).
    Build()

response := mock.NewResponseBuilder(http.StatusOK).
    AddBody(`{"status": "OK"}`).
    AddHeader("Content-Type", "application/json").
    Build()

mockDefinition := mock.NewDefinition(request, response)
```

### Raw Json

Not all features of the Smocker mocks have been captured in the builders and new features may be added in the future. To
help deal with these issues, a `RawJsonDefinition` also exists in the `mock` package. Json conforming to
the [Smocker Mock Definition](https://smocker.dev/technical-documentation/mock-definition.html) can be passed directly
to this, meaning more complex mocks can be created and new mock capabilities can be used immediately.

```go
mockJson := `
{
    "request": {
        "method": "GET",
        "path": "/example"
    },
    "context": {
        "times": 5
    },
    "response": {
        "status": 200,
        "body": "{\"status\": \"OK\"}"
    }
}`
mockDefinition := mock.NewRawJsonDefinition(mockJson)
```