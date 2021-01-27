package main

import (
	"fmt"
	"github.com/twilio/twilio-go/twilio"
	"io/ioutil"
	"net/url"
	"os"
)

func main() {
	client := twilio.NewClient(os.Getenv("TWILIO_ACCOUNT_SID"), os.Getenv("TWILIO_AUTH_TOKEN"))

	resp, err := client.ApiV2010.FetchAccount(os.Getenv("TWILIO_ACCOUNT_SID"));

	if err != nil {
		fmt.Println("Error="+err.Error())
		err = nil
	} else {
		fmt.Print("Success")
		fmt.Println(resp)
	}
	data := url.Values{}
	data.Set("Date", "2020-09-05")
	redirectResp, err := client.Get("https://messaging.twilio.com/v1/Deactivations", data, nil)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		bodyBytes, err := ioutil.ReadAll(redirectResp.Body)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(string(bodyBytes))
	}
}
