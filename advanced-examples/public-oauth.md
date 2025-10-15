// Download the helper library from https://www.twilio.com/docs/go/install
```go
package main

import (
	"fmt"
	"os"

	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
)

func main() {
	var GrantType string = "client_credentials"
	var ClientId string = os.Getenv("CLIENT_ID")
	var ClientSecret string = os.Getenv("CLIENT_SECRET")
	var AccountSid string = os.Getenv("ACCOUNT_SID")
	clientCredentialProvider := twilio.ClientCredentialProvider{
		GrantType:    GrantType,
		ClientId:     ClientId,
		ClientSecret: ClientSecret,
	}
	clientParams := twilio.ClientParams{
		AccountSid:               AccountSid,
		ClientCredentialProvider: &clientCredentialProvider,
	}
	client := twilio.NewRestClientWithParams(clientParams)
	params := api.ListMessageParams{}
	params.SetPageSize(50)

	resp, err := client.Api.ListMessage(&params)
	//
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	} else {
		fmt.Println(*resp[0].Body)
	}
}
```
