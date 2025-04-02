// Download the helper library from https://www.twilio.com/docs/go/install
package main

import (
	"fmt"
	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
	"os"
)

func main() {
	var GrantType string = "client_credentials"
	var ClientId string = "client-id"
	var ClientSecret string = "client-secret"
	clientCredentialProvider := twilio.ClientCredentialProvider{
		GrantType:    GrantType,
		ClientId:     ClientId,
		ClientSecret: ClientSecret,
	}
	clientParams := twilio.ClientParams{
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
