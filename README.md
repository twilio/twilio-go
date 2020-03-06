# twilio-go

## Documentation

The documentation for the Twilio API can be found [here][apidocs].

### Supported Go Versions

This library supports the following Go implementations:

* 1.13.8 

## Installation

To use twilio-go in your project initialize go modules then run: 

```bash
go get github.com/twilio/twilio-go@latest
``` 



## Getting Started

Getting started with the Twilio API couldn't be easier. Create a
`Client` and you're ready to go.

### API Credentials

The Twilio `Client` needs your Twilio credentials. You should pass these
directly to the constructor (see the code below).

```go
import "github.com/twilio/twilio-go"

sid := "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
at := "YYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYY"
c := twilio.NewClient(s, a)
```

### Buy a phone number

```go
package main
import (
	"fmt"
	"github.com/twilio/twilio-go"
)

func main() {
    sid := "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
    at := "YYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYY"
    c := twilio.NewClient(sid, at)
    params := &twilio.IncomingPhoneNumberParams{PhoneNumber:twilio.String("+15017122661")}
    pn, err := c.IncomingPhoneNumbers.Create(params)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(pn)
}
```

### Building
To build *twilio-go* run:

```bash
go build
```

### Testing

To execute the test suite run:

```bash
go test [-v]
```
 
 ### Code Organization
 In general each service's client is implemented in a namesake file (e.g. [Chat Service](https://www.twilio.com/docs/chat/rest/service-resource) is in **chat_service.go**)
 Large services are split between files and will share the same prefix (e.g. **taskrouter_activity.go** and **taskrouter_taskqueue.go**).
 Files for testing are appended with `_test` (e.g. **sync_service_test.go**).
 **twilio.go** ties the library together by defining the `Twilio` struct and the `Client` constructor `NewClient`.  
The `Client` structure promotes memory reuse between service clients, and provides the `baseURL` to service clients to allow redirecting requests to non-production domains.

Under the hood the Twilio Client relies on the private functionality within **internal/twilio.go**.
The main functionality is captured by `SendRequest` and `doWithErr`. 
`SendRequest` and `doWithErr` facilitate network requests and allow the library to provide an `http.Response` and `error` object.
`SendRequest` performs the required encoding and request configuration to comply with Twilio's HTTP standards.
Go's built in struct marshalling is not compatible with Twilio's requirements for form encoding so encoding is performed by the `form` package forked from https://github.com/ajg/form.
`doWithErr` executes the request from `SendRequest`, parses the error (if any), and returns the result.
The `error` returned by `doWithErr` may be the `Error` object, which is meant to parse Twilio-specific error messages from the body of responses to failed requests.

### Behavioral Notes
One of the main behavioral considerations of *twilio-go* is that there is no built-in parameter validation.
This includes the lack of parameter enumeration types that certain services expect, where instead we allow the user to provide any string.
There is a 10 second timeout for requests and there is no retry functionality.


### Code Style
The code is styled with `go fmt` and adheres to [Go's Style Guide](https://github.com/golang/go/wiki/CodeReviewComments) wherever possible.
We use `golangci-lint` for linting locally and as part of the PR process (this will catch most violations of the above guidelines).


[apidocs]: https://www.twilio.com/docs/api
