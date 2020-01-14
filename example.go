package main

import (
	"./src/rest"
	"./src/rest/chat/v2"
	"fmt"
)

func main() {
	twilioClient := twilio.NewClient("ACxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx", "your_auth_token")

	initialName := "Chat Service from Go Client"
	fmt.Println("Creating Chat Service with name:", initialName)
	csParams := &chat.ServiceParams{FriendlyName: initialName}
	resp, err := twilioClient.Chat.Create(csParams)
	if err != nil {
		panic("Failed to create the chat service")
	}
	fmt.Println("Created Chat Service with name:", resp.FriendlyName, "and sid:", resp.Sid)

	fmt.Println("Reading Chat Service with sid:", resp.Sid)
	resp, err = twilioClient.Chat.Read(resp.Sid, &chat.ServiceParams{})
	if err != nil {
		panic("Failed to read the chat service")
	}
	fmt.Println("Read Chat Service and got name:", resp.FriendlyName)

	newName := "Chat Service from Go Client!"
	fmt.Println("Updating name of Chat Service with sid:", resp.Sid, "to", newName)
	updateParams := &chat.ServiceParams{FriendlyName: newName}
	resp, err = twilioClient.Chat.Update(resp.Sid, updateParams)
	if err != nil {
		panic("Failed to update the chat service")
	}
	fmt.Println("Updated name for Chat Service with sid:", resp.Sid, "is", resp.FriendlyName)

	fmt.Println(resp.Sid, resp.FriendlyName)
	resp, err = twilioClient.Chat.Delete(resp.Sid, &chat.ServiceParams{})
	if err != nil {
		panic("Failed to delete the chat service")
	}
	fmt.Println("Deleted chat service")
}
