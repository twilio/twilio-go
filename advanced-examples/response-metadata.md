```go
package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
)

func createMessage(client *twilio.RestClient, to string, from string) {
	params := &api.CreateMessageParams{}
	params.SetTo(to)
	params.SetFrom(from)
	params.SetBody("Hello World")
	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println("Error sending SMS message: " + err.Error())
	} else {
		response, _ := json.Marshal(*resp)
		fmt.Println("Response: " + string(response))
	}
}

func createMessageWithMetadata(client *twilio.RestClient, to string, from string) {
	params := &api.CreateMessageParams{}
	params.SetTo(to)
	params.SetFrom(from)
	params.SetBody("Hello World")
	resp, err := client.Api.CreateMessageWithMetadata(params)
	if err != nil {
		fmt.Println("Error sending SMS message: " + err.Error())
	} else {
		statusCode := resp.GetStatusCode()
		headers := resp.GetHeaders()
		headersJson, _ := json.Marshal(headers)
		fmt.Println("Status code: ", statusCode)
		fmt.Println("headers: " + string(headersJson))
		channel := resp.GetResource()
		response, _ := json.Marshal(channel)
		fmt.Println("Response: " + string(response))
	}
}

func fetchMessage(client *twilio.RestClient, accountSid string, sid string) {
	params := &api.FetchMessageParams{}
	params.SetPathAccountSid(accountSid)

	resp, err := client.Api.FetchMessage(sid, params)
	if err != nil {
		fmt.Println("Error sending SMS message: " + err.Error())
	} else {
		response, _ := json.Marshal(*resp)
		fmt.Println("Response: " + string(response))
	}
}

func fetchMessageWithMetadata(client *twilio.RestClient, accountSid string, sid string) {
	params := &api.FetchMessageParams{}
	params.SetPathAccountSid(accountSid)

	resp, err := client.Api.FetchMessageWithMetadata(sid, params)
	if err != nil {
		fmt.Println("Error sending SMS message: " + err.Error())
	} else {
		statusCode := resp.GetStatusCode()
		headers := resp.GetHeaders()
		headersJson, _ := json.Marshal(headers)
		fmt.Println("Status code: ", statusCode)
		fmt.Println("headers: " + string(headersJson))
		channel := resp.GetResource()
		response, _ := json.Marshal(channel)
		fmt.Println("Response: " + string(response))
	}
}

func updateMessage(client *twilio.RestClient, accountSid string, sid string) {
	params := &api.UpdateMessageParams{}
	params.SetBody("")
	params.SetPathAccountSid(accountSid)

	resp, err := client.Api.UpdateMessage(sid, params)
	if err != nil {
		fmt.Println("Error sending SMS message: " + err.Error())
	} else {
		response, _ := json.Marshal(*resp)
		fmt.Println("Response: " + string(response))
	}
}

func updateMessageWithMetadata(client *twilio.RestClient, accountSid string, sid string) {
	params := &api.UpdateMessageParams{}
	params.SetBody("")
	params.SetPathAccountSid(accountSid)

	resp, err := client.Api.UpdateMessageWithMetadata(sid, params)
	if err != nil {
		fmt.Println("Error sending SMS message: " + err.Error())
	} else {
		statusCode := resp.GetStatusCode()
		headers := resp.GetHeaders()
		headersJson, _ := json.Marshal(headers)
		fmt.Println("Status code: ", statusCode)
		fmt.Println("headers: " + string(headersJson))
		channel := resp.GetResource()
		response, _ := json.Marshal(channel)
		fmt.Println("Response: " + string(response))
	}
}

func deleteMessage(client *twilio.RestClient, accountSid string, sid string) {
	params := &api.DeleteMessageParams{}
	params.SetPathAccountSid(accountSid)

	err := client.Api.DeleteMessage(sid, params)
	if err != nil {
		fmt.Println("Error sending SMS message: " + err.Error())
	} else {
		fmt.Println("Deletion successful")
	}
}

func deleteMessageWithMetadata(client *twilio.RestClient, accountSid string, sid string) {
	params := &api.DeleteMessageParams{}
	params.SetPathAccountSid(accountSid)

	resp, err := client.Api.DeleteMessageWithMetadata(sid, params)
	if err != nil {
		fmt.Println("Error sending SMS message: " + err.Error())
	} else {
		statusCode := resp.GetStatusCode()
		headers := resp.GetHeaders()
		headersJson, _ := json.Marshal(headers)
		fmt.Println("Status code: ", statusCode)
		fmt.Println("headers: " + string(headersJson))
		channel := resp.GetResource()
		response, _ := json.Marshal(channel)
		fmt.Println("Response: " + string(response))
	}
}

func pageMessage(client *twilio.RestClient, pageSize int, limit int) {
	params := &api.ListMessageParams{}
	params.SetPageSize(pageSize)
	params.SetLimit(limit)
	resp, err := client.Api.PageMessage(params, "", "")
	if err != nil {
		fmt.Println(err)
	} else {
		resource, _ := json.Marshal(resp)
		fmt.Println("Resource: " + string(resource))
	}
}

func pageMessageWithMetadata(client *twilio.RestClient, pageSize int, limit int) {
	params := &api.ListMessageParams{}
	params.SetPageSize(pageSize)
	params.SetLimit(limit)
	resp, err := client.Api.PageMessageWithMetadata(params, "", "")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		statusCode := resp.GetStatusCode()
		headers := resp.GetHeaders()
		resource := resp.GetResource()
		response, _ := json.Marshal(resource)
		headersJson, _ := json.Marshal(headers)
		fmt.Println("Status code: ", statusCode)
		fmt.Println("headers: " + string(headersJson))
		fmt.Println("Resource: " + string(response))
	}
}

func streamMessage(client *twilio.RestClient, pageSize int, limit int) {
	params := &api.ListMessageParams{}
	params.SetPageSize(pageSize)
	params.SetLimit(limit)

	channel, _ := client.Api.StreamMessage(params)
	for record := range channel {
		fmt.Println("Body: ", *record.Body)
	}
}

func streamMessageWithMetadata(client *twilio.RestClient, pageSize int, limit int) {
	params := &api.ListMessageParams{}
	params.SetPageSize(pageSize)
	params.SetLimit(limit)

	resp, _ := client.Api.StreamMessageWithMetadata(params)

	statusCode := resp.GetStatusCode()
	headers := resp.GetHeaders()
	headersJson, _ := json.Marshal(headers)
	fmt.Println("Status code: ", statusCode)
	fmt.Println("headers: " + string(headersJson))
	channel := resp.GetResource()
	for record := range channel {
		fmt.Println("Body: ", *record.Body)
	}
}

func listMessage(client *twilio.RestClient, pageSize int, limit int) {
	params := &api.ListMessageParams{}
	params.SetPageSize(pageSize)
	params.SetLimit(limit)

	resp, _ := client.Api.ListMessage(params)
	for record := range resp {
		fmt.Println("Body: ", *resp[record].Body)
	}
}

func listMessageWithMetadata(client *twilio.RestClient, pageSize int, limit int) {
	params := &api.ListMessageParams{}
	params.SetPageSize(pageSize)
	params.SetLimit(limit)

	resp, err := client.Api.ListMessageWithMetadata(params)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		statusCode := resp.GetStatusCode()
		headers := resp.GetHeaders()
		headersJson, _ := json.Marshal(headers)
		fmt.Println("Status code: ", statusCode)
		fmt.Println("headers: " + string(headersJson))

		resource := resp.GetResource()
		for record := range resource {
			fmt.Println("Body: ", *resource[record].Body)
		}
	}
}

func main() {
	accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")
	receiverPhoneNumber := os.Getenv("RECEIVER_PHONE_NUMBER")
	twilioPhoneNumber := os.Getenv("TWILIO_PHONE_NUMBER")

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	// create
	createMessage(client, receiverPhoneNumber, twilioPhoneNumber)
	createMessageWithMetadata(client, receiverPhoneNumber, twilioPhoneNumber)

	// fetch
	sid := os.Getenv("MESSAGE_SID")
	fetchMessage(client, accountSid, sid)
	fetchMessageWithMetadata(client, accountSid, sid)

	// update
	updateMessage(client, accountSid, sid)
	updateMessageWithMetadata(client, accountSid, sid)

	// delete
	deleteMessage(client, accountSid, sid)
	deleteMessageWithMetadata(client, accountSid, sid)

	pageSize := 2
	limit := 5

	// page
	pageMessage(client, pageSize, limit)
	pageMessageWithMetadata(client, pageSize, limit)

	// stream
	streamMessage(client, pageSize, limit)
	streamMessageWithMetadata(client, pageSize, limit)

	// list
	listMessage(client, pageSize, limit)
	listMessageWithMetadata(client, pageSize, limit)
}

```
