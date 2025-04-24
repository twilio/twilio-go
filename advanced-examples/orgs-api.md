```go
package main

import (
	"fmt"
	"os"

	"github.com/twilio/twilio-go"
	preview_iam "github.com/twilio/twilio-go/rest/preview_iam/organization"
)

func main() {
	var GrantType string = "client_credentials"
	var ClientId string = os.Getenv("CLIENT_ID")
	var ClientSecret string = os.Getenv("CLIENT_SECRET")
	var AccountSid string = os.Getenv("ACCOUNT_SID")
	var orgSid string = os.Getenv("ORG_SID")
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
	listAccounts, _ := client.PreviewIamOrganization.ListOrganizationAccounts(orgSid, &preview_iam.ListOrganizationAccountsParams{})
	for _, account := range listAccounts {
		fmt.Println(account.AccountSid)
		fmt.Println(account.OwnerSid)
	}
}
```
