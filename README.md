# twilio-go

[![Tests](https://github.com/twilio/twilio-go/actions/workflows/test-and-deploy.yml/badge.svg)](https://github.com/twilio/twilio-go/actions/workflows/test-and-deploy.yml)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/twilio/twilio-go)](https://pkg.go.dev/github.com/twilio/twilio-go)
[![Release](https://img.shields.io/github/release/twilio/twilio-go.svg)](https://github.com/twilio/twilio-go/releases/latest)

All the code [here](./rest) was generated by [twilio-oai-generator](https://github.com/twilio/twilio-oai-generator) by
leveraging [openapi-generator](https://github.com/OpenAPITools/openapi-generator)
and [twilio-oai](https://github.com/twilio/twilio-oai). If you find an issue with the generation or the OpenAPI specs,
please go ahead and open an issue or a PR against the relevant repositories.

## 🚀 Feature Update
Twilio Go Helper Library's version 1.20.0 adds support for the application/json content type in the request body. See example [here][content-api-example].
Behind the scenes Go Helper is now auto-generated via OpenAPI with this release.
This enables us to rapidly add new features and enhance consistency across versions and languages.

## Documentation

The documentation for the Twilio API can be found [here][apidocs].

The Go library documentation can be found [here][libdocs].

### Supported Go Versions

This library supports the following Go implementations:

- Go 1.18
- Go 1.19
- Go 1.20
- Go 1.21
- Go 1.22
- Go 1.23
- Go 1.24

## Installation

The recommended way to install `twilio-go` is by using [Go modules](https://go.dev/ref/mod#go-get).

If you already have an initialized project, you can run the command below from your terminal in the project directory to install the library:

```shell
go get github.com/twilio/twilio-go
```

If you are starting from scratch in a new directory, you will first need to create a go.mod file for tracking dependencies such as twilio-go. This is similar to using package.json in a Node.js project or requirements.txt in a Python project. [You can read more about mod files in the Go documentation](https://golang.org/doc/modules/managing-dependencies). To create the file, run the following command in your terminal:

```shell
go mod init twilio-example
```

Once the module is initialized, you may run the installation command from above, which will update your go.mod file to include twilio-go.

### Test your installation

Try sending yourself an SMS message by pasting the following code example into a sendsms.go file in the same directory where you installed twilio-go. Be sure to update the accountSid, authToken, and from phone number with values from your Twilio account. The to phone number can be your own mobile phone number.

```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

func main() {
	accountSid := "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
	authToken := "f2xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	params := &twilioApi.CreateMessageParams{}
	params.SetTo("+15558675309")
	params.SetFrom("+15017250604")
	params.SetBody("Hello from Go!")

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println("Error sending SMS message: " + err.Error())
	} else {
		response, _ := json.Marshal(*resp)
		fmt.Println("Response: " + string(response))
	}
}
```

Save `sendsms.go`. In your terminal from the same directory, run:

```shell
go run sendsms.go
```

After a brief delay, you will receive the text message on your phone.

> **Warning**
> It's okay to hardcode your credentials when testing locally, but you should use environment variables to keep them secret before committing any code or deploying to production. Check out [How to Set Environment Variables](https://www.twilio.com/blog/2017/01/how-to-set-environment-variables.html) for more information.

## Use the helper library

### API credentials

The Twilio `RestClient` needs your Twilio credentials. We recommend storing them as environment variables, so that you don't have to worry about committing and accidentally posting them somewhere public. See http://twil.io/secure for more details on how to store environment variables.

```go
package main

import "github.com/twilio/twilio-go"

func main() {
	// This will look for `TWILIO_ACCOUNT_SID` and `TWILIO_AUTH_TOKEN` variables inside the current environment to initialize the constructor
	// You can find your Account SID and Auth Token at twilio.com/console
	// `TWILIO_ACCOUNT_SID` should be in format "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
	client := twilio.NewRestClient()
}
```

If you don't want to use environment variables, you can also pass the credentials directly to the constructor as below.

```go
package main

import "github.com/twilio/twilio-go"

func main() {
	accountSid := "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
	authToken := "YYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYY"
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})
}
```

### Use a Subaccount

Subaccounts in Twilio are just accounts that are "owned" by your account. Twilio users can create subaccounts to help separate Twilio account usage into different buckets.

If you wish to make API calls with a Subaccount, you can do so by setting the `AccountSid` field in the `twilio.ClientParams`:

```go
package main

import "github.com/twilio/twilio-go"

func main() {
	// subaccountSid should also be in format "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
	subaccountSid := os.Getenv("TWILIO_SUBACCOUNT_SID")
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		AccountSid: subaccountSid,
	})
}
```

### Use API Keys

Lastly, if you want to follow best practices and initialize your client using an [API Key and Secret](https://www.twilio.com/docs/iam/keys/api-key), instead of potentially exposing your account's AuthToken, use the following pattern:

```go
package main

import (
	"os"

	"github.com/twilio/twilio-go"
)

func main() {
	accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	apiKey := os.Getenv("TWILIO_API_KEY")
	apiSecret := os.Getenv("TWILIO_API_SECRET")

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username:   apiKey,
		Password:   apiSecret,
		AccountSid: accountSid,
	})
}

```

### Specify a Region and/or Edge

```go
package main

import (
	"github.com/twilio/twilio-go"
)

func main() {
	client := twilio.NewRestClient()
	client.SetRegion("au1")
	client.SetEdge("sydney")
}
```

This will result in the `hostname` transforming from `api.twilio.com` to `api.sydney.au1.twilio.com`.

A Twilio client constructed without these parameters will also look for `TWILIO_REGION` and `TWILIO_EDGE` variables inside the current environment.

### Buy a phone number

```go
package main

import (
	"fmt"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

func main() {
	phoneNumber := "AVAILABLE_TWILIO_PHONE_NUMBER"

	client := twilio.NewRestClient()

	params := &twilioApi.CreateIncomingPhoneNumberParams{}
	params.SetPhoneNumber(phoneNumber)

	resp, err := client.Api.CreateIncomingPhoneNumber(params)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Phone Number Status: " + *resp.Status)
	}
}
```

### Make a call

```go
package main

import (
	"fmt"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
	"os"
)

func main() {
	from := os.Getenv("TWILIO_FROM_PHONE_NUMBER")
	to := os.Getenv("TWILIO_TO_PHONE_NUMBER")

	client := twilio.NewRestClient()

	params := &twilioApi.CreateCallParams{}
	params.SetTo(to)
	params.SetFrom(from)
	params.SetUrl("http://twimlets.com/holdmusic?Bucket=com.twilio.music.ambient")

	resp, err := client.Api.CreateCall(params)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Call Status: " + *resp.Status)
		fmt.Println("Call Sid: " + *resp.Sid)
		fmt.Println("Call Direction: " + *resp.Direction)
	}
}
```

### Get data about an existing Call

```go
package main

import (
	"fmt"
	"os"

	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

func main() {
	accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	apiKey := os.Getenv("TWILIO_API_KEY")
	apiSecret := os.Getenv("TWILIO_API_SECRET")

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username:   apiKey,
		Password:   apiSecret,
		AccountSid: accountSid,
	})

	params := &twilioApi.FetchCallParams{}

	resp, err := client.Api.FetchCall("CA42ed11f93dc08b952027ffbc406d0868", params)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Call Status: " + *resp.Status)
		fmt.Println("Call Sid: " + *resp.Sid)
		fmt.Println("Call Direction: " + *resp.Direction)
	}
}
```

### Iterate through records

This library also offers paging functionality. Collections such as calls and messages have `ListXxx` and `StreamXxx`
functions that page under the hood. With both list and stream, you can specify the number of records you want to
receive (limit) and the maximum size you want each page fetch to be (pageSize). The library will then handle the task
for you.

`List` eagerly fetches all records and returns them as a list, whereas `Stream` streams the records and lazily retrieves
the pages as you iterate over the collection. Also, `List` returns no records if any errors are encountered while paging,
whereas `Stream` returns all records up until encountering an error. You can also page manually using the `PageXxx`
function in each of the apis.

#### Use `ListXxx` or `StreamXxx`

```go
package main

import (
	"fmt"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
	"os"
)

func main() {
	from := os.Getenv("TWILIO_FROM_PHONE_NUMBER")

	client := twilio.NewRestClient()

	params := &twilioApi.ListMessageParams{}
	params.SetFrom(from)
	params.SetPageSize(20)
	params.SetLimit(100)

	resp, _ := client.Api.ListMessage(params)
	for record := range resp {
		fmt.Println("Body: ", *resp[record].Body)
	}

	channel, _ := client.Api.StreamMessage(params)
	for record := range channel {
		fmt.Println("Body: ", *record.Body)
	}
}
```

#### Use `PageXxx`

```go
package main

import (
	"fmt"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
	"net/url"
	"os"
)

func main() {
	from := os.Getenv("TWILIO_FROM_PHONE_NUMBER")

	client := twilio.NewRestClient()

	params := &twilioApi.ListMessageParams{}
	params.SetFrom(from)
	params.SetPageSize(20)

	var pageToken string
	var pageNumber string
	resp, err = client.Api.PageMessage(params, "", "")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp.NextPageUri)
		u, _ := url.Parse(resp.NextPageUri)
		q := u.Query()
		pageToken = q.Get("PageToken")
		pageNumber = q.Get("Page")
	}

	resp, err := client.Api.PageMessage(params, pageToken, pageNumber)
	if err != nil {
		fmt.Println(err)
	} else {
		if resp != nil {
			fmt.Println(*resp.Messages[0].Body)
		}
	}
}
```

### Handle Exceptions

If the Twilio API returns a 400 or a 500 level HTTP response, the twilio-go library will include information in the returned err value. 400-level errors are [normal during API operation](https://www.twilio.com/docs/usage/requests-to-twilio) ("Invalid number", "Cannot deliver SMS to that number", for example) and should be handled appropriately.

```go
package main

import (
	"fmt"
	"os"

	"github.com/twilio/twilio-go"
	twilioclient "github.com/twilio/twilio-go/client"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

func main() {
	phoneNumber := os.Getenv("TWILIO_PHONE_NUMBER")

	client := twilio.NewRestClient()

	params := &twilioApi.CreateIncomingPhoneNumberParams{}
	params.SetPhoneNumber(phoneNumber)

	resp, err := client.Api.CreateIncomingPhoneNumber(params)
	if err != nil {
		twilioError := err.(*twilioclient.TwilioRestError)
		fmt.Println(twilioError.Error())
	}
}
```

### Generate TwiML

To control phone calls, your application needs to output [TwiML](https://www.twilio.com/docs/voice/twiml).

Use the `twiml` package to easily create such responses.

```go
package main

import (
	"fmt"
	"github.com/twilio/twilio-go/twiml"
)

func main() {
	//Construct Verbs
	dial := &twiml.VoiceDial{}
	say := &twiml.VoiceSay{
		Message:            "Welcome to Twilio!",
		Voice:              "woman",
		Language:           "en-gb",
		OptionalAttributes: map[string]string{"input": "test"},
	}
	pause := &twiml.VoicePause{
		Length: "10",
	}
	//Construct Noun
	queue := &twiml.VoiceQueue{
		Url: "www.twilio.com",
	}
	//Adding Queue to Dial
	dial.InnerElements = []twiml.Element{queue}

	//Adding all Verbs to twiml.Voice
	verbList := []twiml.Element{dial, say, pause}
	twimlResult, err := twiml.Voice(verbList)
	if err == nil {
		fmt.Println(twimlResult)
	} else {
		fmt.Println(err)
	}
}
```

This will print the following:

```xml
<?xml version="1.0" encoding="UTF-8"?>
<Response>
    <Dial>
        <Queue url="www.twilio.com"/>
    </Dial>
    <Say voice="woman" language="en-gb" input="test">Welcome to Twilio!</Say>
    <Pause length="10"/>
</Response>
```

## Advanced Usage

### Use the request validator

Validate that GET/POST Requests are coming from Twilio:

```go
package main

import (
	"fmt"
	"os"

	"github.com/twilio/twilio-go/client"
)

func main() {
	// You can find your Auth Token at twilio.com/console
	// For this example: authToken := "12345"
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")

	requestValidator := client.NewRequestValidator(authToken)

	// Twilio's request URL
	url := "https://mycompany.com/myapp.php?foo=1&bar=2"

	// Post variables in Twilio's request
	params := map[string]string{
		"CallSid": "CA1234567890ABCDE",
		"Caller":  "+12349013030",
		"Digits":  "1234",
		"From":    "+12349013030",
		"To":      "+18005551212",
	}

	// X-Twilio-Signature header attached to the request
	signature := "0/KCTR6DLpKmkAf8muzZqo1nDgQ="

	// Validate GET request
	fmt.Println(requestValidator.Validate(url, params, signature))

	// Example of the POST request
	Body := []byte(`{"property": "value", "boolean": true}`)
	theUrl := "https://mycompany.com/myapp.php?bodySHA256=0a1ff7634d9ab3b95db5c9a2dfe9416e41502b283a80c7cf19632632f96e6620"
	theSignature := "y77kIzt2vzLz71DgmJGsen2scGs="

	// Validate POST request
	fmt.Println(requestValidator.ValidateBody(theUrl, Body, theSignature))
}
```

### Use standalone products

Don't want to import the top-level Twilio RestClient with access to the full suite of Twilio products? Use standalone product services instead:

```go
package main

import (
	"github.com/twilio/twilio-go/client"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
	serverless "github.com/twilio/twilio-go/rest/serverless/v1"
	"os"
)

func main() {
	accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")

	// Create an instance of our default BaseClient implementation
	// You will need to provide your API credentials to the Client manually
	defaultClient := &client.Client{
		Credentials: client.NewCredentials(accountSid, authToken),
	}
	defaultClient.SetAccountSid(accountSid)

	coreApiService := twilioApi.NewApiServiceWithClient(defaultClient)
	serverlessApiService := serverless.NewApiServiceWithClient(defaultClient)
}
```

### Other advanced examples

- [Learn how to create your own custom HTTP client](./advanced-examples/custom-http-client.md)

## Build Access Tokens

This library supports [access token](https://www.twilio.com/docs/iam/access-tokens) generation for use in the Twilio Client SDKs.

Here's how you would generate a token for the Voice SDK:

```go
package main

import (
	"os"
	"github.com/twilio/twilio-go/client/jwt"
)

accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
applicationSid := os.Getenv("TWILIO_TWIML_APP_SID")
apiKey := os.Getenv("TWILIO_API_KEY")
apiSecret := os.Getenv("TWILIO_API_SECRET")
identity := "fake123"

params := jwt.AccessTokenParams{
	AccountSid: accountSid,
	SigningKeySid: apiKey,
	Secret: apiSecret,
	Identity: identity,
}

jwtToken := jwt.CreateAccessToken(params)
voiceGrant := &jwt.VoiceGrant{
	Incoming: jwt.Incoming{Allow: true},
	Outgoing: jwt.Outgoing{
		ApplicationSid: applicationSid,
	},
}

jwtToken.AddGrant(voiceGrant)
token, err := jwtToken.ToJwt()
```

Create Capability Tokens for TaskRouter v1:

```go
package main

import (
	"os"
	"github.com/twilio/twilio-go/client/jwt/taskrouter"
)

AccountSid := os.Getenv("TWILIO_ACCOUNT_SID")
AuthToken := os.Getenv("TWILIO_AUTH_TOKEN")
WorkspaceSid := os.Getenv("TWILIO_WORKSPACE_SID")
ChannelID := os.Getenv("TWILIO_CHANNEL_ID")

Params = taskrouter.CapabilityTokenParams{
	AccountSid: AccountSid,
	AuthToken: AuthToken,
	WorkspaceSid: WorkspaceSid,
	ChannelID: ChannelID,
}

capabilityToken := taskrouter.CreateCapabilityToken(Params)
token, err := capabilityToken.ToJwt()
```

## Local Usage

### Building

To build _twilio-go_ run:

```shell
go build ./...
```

### Testing

To execute the test suite run:

```shell
go test ./...
```

### Generating Local Documentation

To generate documentation, from the root directory:

```shell
godoc -http=localhost:{port number}
```

Then, navigate to `http://localhost:{port number}/pkg/github.com/twilio/twilio-go` in your local browser.

Example:

```shell
godoc -http=localhost:6060
```
http://localhost:6060/pkg/github.com/twilio/twilio-go

## OAuth Feature for twilio-go
The Twilio Go library now includes an OAuth feature that allows you to authenticate API requests using OAuth 2.0 tokens. This feature is particularly useful for applications that require enhanced security and token-based authentication.

To use the OAuth feature, you can refer to [public-oauth.md](https://github.com/twilio/twilio-go/blob/main/advanced-examples/public-oauth.md) for an example of how to implement it in your application.


## Docker Image

The `Dockerfile` present in this repository and its respective `twilio/twilio-go` Docker image are currently used by Twilio for testing purposes only.

## Getting help

If you need help installing or using the library, please check the [Twilio Support Help Center](https://support.twilio.com) first, and [file a support ticket](https://twilio.com/help/contact) if you don't find an answer to your question.

If you've instead found a bug in the library or would like new features added, go ahead and open issues or pull requests against this repo!

[apidocs]: https://www.twilio.com/docs/api
[libdocs]: https://pkg.go.dev/github.com/twilio/twilio-go?tab=versions
[content-api-example]: https://github.com/twilio/twilio-go/blob/main/advanced-examples/content.md
