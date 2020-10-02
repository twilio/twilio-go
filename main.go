// Package twilio provides bindings for Twilio's REST APIs.
package main

import (
	"fmt"
	"github.com/twilio/twilio-go/twilio"
	"os"
)

func main() {
	client := twilio.NewClient(os.Getenv("TWILIO_ACCOUNT_SID"), os.Getenv("TWILIO_AUTH_TOKEN"))

	response, err := client.Studio.V2Api.V2FlowsList(nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(response.Flows)
}
