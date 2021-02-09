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
	client := twilio.NewClient(accountSid, authToken)

	from := "9089137656"
	to := "4803520524"
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