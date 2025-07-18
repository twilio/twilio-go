```go
package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/twilio/twilio-go"
	content "github.com/twilio/twilio-go/rest/content/v1"
)

func createContent(client *twilio.RestClient) {
	twilioText := &content.TwilioText{Body: "Hello World"}
	contentType := &content.Types{TwilioText: twilioText}
	createContentRequest := &content.ContentCreateRequest{Language: "en", Types: *contentType}
	params := &content.CreateContentParams{}
	params.SetContentCreateRequest(*createContentRequest)

	resp, err := client.ContentV1.CreateContent(params)
	if err != nil {
		fmt.Println("Error creating content: " + err.Error())
	} else {
		response, _ := json.Marshal(*resp)
		fmt.Println("Response: " + string(response))
	}
}

func main() {
	accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})
	createContent(client)
}
```
