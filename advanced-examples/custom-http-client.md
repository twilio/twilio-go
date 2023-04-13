# Custom HTTP Clients for the Twilio Go Helper Library

If you are working with the [Twilio Go Helper Library](../README.md), and you need to be able to modify the HTTP requests that the library makes to the Twilio servers, you’re in the right place. The most common reason for altering the HTTP request is to connect and authenticate with an enterprise’s proxy server. We’ll provide sample code that you can drop into your app to handle this use case.

## Connect and authenticate with a proxy server

To connect and provide credentials to a proxy server that may be between your app and Twilio, you need a way to modify the HTTP requests that the Twilio helper library makes on your behalf when invoking Twilio's REST API.

In Go, the Twilio helper library uses the native [net/http package](https://pkg.go.dev/net/http) under the hood to make the HTTP requests. The Twilio Helper Library allows you to provide your own `Client` for making API requests.

The following example shows a typical request without a custom `Client`.

```go
twilioClient := twilio.NewRestClient()

params := &twilioApi.CreateMessageParams{}
params.SetTo("+15558675309")
params.SetFrom("+15017250604")
params.SetBody("Hey there!")

resp, err := twilioClient.Api.CreateMessage(params)
```

Out of the box, the helper library creates a default `Client` for you, using the Twilio credentials from your environment variables or that you pass in directly. However, there’s nothing stopping you from creating your own client and using that.

Once you have your own `Client`, you can pass it to any Twilio REST API resource action you want.

## Create and use your custom Client

When you take a closer look at the input parameters for `twilio.RestClient`, you see that the `Client` parameter is actually of type `client.BaseClient`.

`client.BaseClient` is an abstraction that allows plugging in any implementation of an HTTP client you want (or even creating a mocking layer for unit testing).

Now that you understand how all the components fit together, you can create your own `Client` that can connect through a proxy server. To make this reusable, here’s a class that you can use to create this `HttpClient` whenever you need one.

Here’s an example of sending an SMS message with a custom client:

```go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/twilio/twilio-go"
	"github.com/twilio/twilio-go/client"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

func main() {
	accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")

	// Add proxy settings to a http Transport object
	transport := &http.Transport{
		// https://pkg.go.dev/net/http#ProxyFromEnvironment
		Proxy: http.ProxyFromEnvironment,
	}

	// Add the Transport to an http Client
	httpClient := &http.Client{
		Transport: transport,
	}

	// Create your custom Twilio client using the http client and your credentials
	twilioHttpClient := client.Client{
		Credentials: client.NewCredentials(accountSid, authToken),
		HTTPClient:  httpClient,
	}
	twilioHttpClient.SetAccountSid(accountSid)
	twilioClient := twilio.NewRestClientWithParams(twilio.ClientParams{Client: &twilioHttpClient})

	params := &twilioApi.CreateMessageParams{}
	params.SetTo("+15558675310")
	params.SetFrom("+15017122661")
	params.SetBody("Hey there!")

	resp, err := twilioClient.Api.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		response, _ := json.Marshal(*resp)
		fmt.Println("Response: " + string(response))
	}
}
```

In this example, you use environment variables loaded at the program startup to retrieve various configuration settings:

- Your Twilio Account Sid and Auth Token ([found here, in the Twilio console](https://www.twilio.com/console))
- A proxy address in the form of `http://127.0.0.1:8888`

These settings are either exported manually by yourself in the terminal, or located in a file such as `.env`, like so:

```text
ACCOUNT_SID=ACxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
AUTH_TOKEN=your_auth_token

HTTPS_PROXY=https://127.0.0.1:8888
HTTP_PROXY=http://127.0.0.1:8888
```

## What else can this technique be used for?

Now that you know how to inject your own `Client` into the Twilio API request pipeline, you could use this technique to add custom HTTP headers and authorization to the requests (perhaps as required by an upstream proxy server). You could do so by overriding the `SendRequest` method, and adding any desired pre-processing to your requests and responses, like so:

```go
package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/twilio/twilio-go"
	"github.com/twilio/twilio-go/client"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

type MyClient struct {
	client.Client
}

func (c *MyClient) SendRequest(method string, rawURL string, data url.Values, headers map[string]interface{}) (*http.Response, error) {
	// Custom code to pre-process request here
	resp, err := c.Client.SendRequest(method, rawURL, data, headers)
	// Custom code to pre-process response here
	fmt.Println(resp.StatusCode)
	return resp, err
}

func main() {
	accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")

	customClient := &MyClient{
		Client: client.Client{
			Credentials: client.NewCredentials(accountSid, authToken),
		},
	}
	customClient.SetAccountSid(accountSid)

	twilioClient := twilio.NewRestClientWithParams(twilio.ClientParams{Client: customClient})

	// You may also use custom clients with standalone product services
	twilioApiV2010 := openapi.NewApiServiceWithClient(customClient)
}
```

You could also implement your own `Client` to mock the Twilio API responses so your unit and integration tests can run quickly without needing to make a connection to Twilio. In fact, there’s already an example online showing [how to do exactly that with Node.js and Prism](https://www.twilio.com/docs/openapi/mock-api-generation-with-twilio-openapi-spec).

We can’t wait to see what you build!
