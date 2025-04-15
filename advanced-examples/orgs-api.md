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
	listAccounts, _ := client.PreviewIam.ListOrganizationAccounts(orgSid, &preview_iam.ListOrganizationAccountsParams{})
	for _, account := range listAccounts {
		fmt.Println(account.AccountSid)
		fmt.Println(account.OwnerSid)
	}
}
```
