package main
import (
	"fmt"
	twilio "github.com/twilio/twilio-go/client"
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