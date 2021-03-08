# twilio-go

[![Build Status](https://travis-ci.com/twilio/twilio-go.png?branch=main)](https://travis-ci.com/twilio/twilio-go)
[![Learn OSS Contribution in TwilioQuest](https://img.shields.io/static/v1?label=TwilioQuest&message=Learn%20to%20contribute%21&color=F22F46&labelColor=1f243c&style=flat-square&logo=data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAIAAAACACAMAAAD04JH5AAAASFBMVEUAAAAZGRkcHBwjIyMoKCgAAABgYGBoaGiAgICMjIyzs7PJycnMzMzNzc3UoBfd3d3m5ubqrhfrMEDu7u739/f4vSb/3AD///9tbdyEAAAABXRSTlMAAAAAAMJrBrEAAAKoSURBVHgB7ZrRcuI6EESdyxXGYoNFvMD//+l2bSszRgyUYpFAsXOeiJGmj4NkuWx1Qeh+Ekl9DgEXOBwOx+Px5xyQhDykfgq4wG63MxxaR4ddIkg6Ul3g84vCIcjPBA5gmUMeXESrlukuoK33+33uID8TWeLAdOWsKpJYzwVMB7bOzYSGOciyUlXSn0/ABXTosJ1M1SbypZ4O4MbZuIDMU02PMbauhhHMHXbmebmALIiEbbbbbUrpF1gwE9kFfRNAJaP+FQEXCCTGyJ4ngDrjOFo3jEL5JdqjF/pueR4cCeCGgAtwmuRS6gDwaRiGvu+DMFwSBLTE3+jF8JyuV1okPZ+AC4hDFhCHyHQjdjPHUKFDlHSJkHQXMB3KpSwXNGJPcwwTdZiXlRN0gSp0zpWxNtM0beYE0nRH6QIbO7rawwXaBYz0j78gxjokDuv12gVeUuBD0MDi0OQCLvDaAho4juP1Q/jkAncXqIcCfd+7gAu4QLMACCLxpRsSuQh0igu0C9Svhi7weAGZg50L3IE3cai4IfkNZAC8dfdhsUD3CgKBVC9JE5ABAFzg4QL/taYPAAWrHdYcgfLaIgAXWJ7OV38n1LEF8tt2TH29E+QAoDoO5Ve/LtCQDmKM9kPbvCEBApK+IXzbcSJ0cIGF6e8gpcRhUDogWZ8JnaWjPXc/fNnBBUKRngiHgTUSivSzDRDgHZQOLvBQgf8rRt+VdBUUhwkU6VpJ+xcOwQUqZr+mR0kvBUgv6cB4+37hQAkXqE8PwGisGhJtN4xAHMzrsgvI7rccXqSvKh6jltGlrOHA3Xk1At3LC4QiPdX9/0ndHpGVvTjR4bZA1ypAKgVcwE5vx74ulwIugDt8e/X7JgfkucBMIAr26ndnB4UCLnDOqvteQsHlgX9N4A+c4cW3DXSPbwAAAABJRU5ErkJggg==)](https://twil.io/learn-open-source)

## Documentation

The documentation for the Twilio API can be found [here][apidocs].

### Supported Go Versions

This library supports the following Go implementations:

* 1.14, 1.15, 1.16

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

accountSID := "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
authToken := "YYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYY"

client := twilio.NewClient(accountSID, authToken)
```

We suggest storing your credentials as environment variables and then use it in your code. Why? You'll never have to worry about committing your credentials and accidentally posting them somewhere public.

```go
import (
	"github.com/twilio/twilio-go/twilio"
	"os"
)

accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
authToken := os.Getenv("TWILIO_AUTH_TOKEN")
client := twilio.NewClient(accountSid, authToken)
```

### Buy a phone number

```go
package main
import (
	"fmt"
	twilio "github.com/twilio/twilio-go/twilio"
	openapi "github.com/twilio/twilio-go/twilio/rest/api/v2010"
	"os"
)

func main() {
	accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")
	phoneNumber := os.Getenv("TWILIO_PHONE_NUMBER")

	client := twilio.NewClient(accountSid, authToken)

	params := &openapi.CreateIncomingPhoneNumberParams{}
	params.PhoneNumber = &phoneNumber

	resp, err := client.ApiV2010.CreateIncomingPhoneNumber(accountSid, params)
	if err != nil {
		fmt.Println(err.Error())
		err = nil
	} else {
		fmt.Println(resp)
	}
}
```

### Send a text message

```go
package main
import (
	"fmt"
	twilio "github.com/twilio/twilio-go/twilio"
	openapi "github.com/twilio/twilio-go/twilio/rest/api/v2010"
	"os"
)
func main() {
	accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")
	from := os.Getenv("TWILIO_FROM_PHONE_NUMBER")
	to := os.Getenv("TWILIO_TO_PHONE_NUMBER")

	client := twilio.NewClient(accountSid, authToken)

	text := "Hello there"

	params := &openapi.CreateMessageParams{}
	params.To = &to
	params.From = &from
	params.Body = &text


	resp, err := client.ApiV2010.CreateMessage(accountSid, params)
	if err != nil {
		fmt.Println(err.Error())
		err = nil
	} else {
		fmt.Println(resp)
	}
}
```

### Make a call
``` go
package main

import (
	"fmt"
	twilio "github.com/twilio/twilio-go/twilio"
	openapi "github.com/twilio/twilio-go/twilio/rest/api/v2010"
	"os"
)

func main() {
	accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")
	from := os.Getenv("TWILIO_FROM_PHONE_NUMBER")
	to := os.Getenv("TWILIO_TO_PHONE_NUMBER")

	client := twilio.NewClient(accountSid, authToken)

	callurl := "http://twimlets.com/holdmusic?Bucket=com.twilio.music.ambient"

	params := &openapi.CreateCallParams{}
	params.To = &to
	params.From = &from
	params.Url = &callurl

	resp, err := client.ApiV2010.CreateCall(accountSid, params)
	if err != nil {
		fmt.Println(err.Error())
		err = nil
	} else {
		fmt.Println(resp)
	}
}
```

### Handling Exceptions
```go
package main
import (
	"fmt"
	twilio "github.com/twilio/twilio-go/twilio"
	"github.com/twilio/twilio-go/framework/error"
	openapi "github.com/twilio/twilio-go/twilio/rest/api/v2010"
	"os"
)

func main() {
	accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")
	phoneNumber := os.Getenv("TWILIO_PHONE_NUMBER")

	client := twilio.NewClient(accountSid, authToken)

	params := &openapi.CreateIncomingPhoneNumberParams{}
	params.PhoneNumber = &phoneNumber

	resp, err := client.ApiV2010.CreateIncomingPhoneNumber(accountSid, params)
	if err != nil {
		twilioError := err.(*error.TwilioRestError)
		fmt.Println(twilioError.Error())
	}
	fmt.Println(resp)
```

For more descriptive exception types, please see the [Twilio documentation](https://www.twilio.com/docs/libraries/go/usage-guide#exceptions).

### Building
To build *twilio-go* run:

```bash
go build ./...
```

### Testing

To execute the test suite run:

```bash
go test ./...
```

### Getting help

If you need help installing or using the library, please check the [Twilio Support Help Center](https://support.twilio.com) first, and [file a support ticket](https://twilio.com/help/contact) if you don't find an answer to your question.

All the code [here](twilio/rest) was generated by [twilio-oai-generator](https://github.com/twilio/twilio-oai-generator) by leveraging [openapi-generator](https://github.com/OpenAPITools/openapi-generator) and [twilio-oai](https://github.com/twilio/twilio-oai). If you find an issue with the generation or the openapi specs, please go ahead and open an issue or a PR against the relevant repositories. 

[apidocs]: https://www.twilio.com/docs/api
[twiml]: https://www.twilio.com/docs/api/twiml
[libdocs]: https://twilio.github.io/twilio-go

