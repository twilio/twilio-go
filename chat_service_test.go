package twilio

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestChatService_Create(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Services", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testFormValues(t, r, values{"FriendlyName": "ChatService"})
		response := `{"friendly_name":"ChatService"}`

		fmt.Fprint(w, response)
	})

	got, err := client.Chat.Service.Create(&ChatServiceParams{FriendlyName: "ChatService"})

	if err != nil {
		t.Errorf("ChatService.Create returned error: %v", err)
	}

	expected := &ChatService{FriendlyName: "ChatService"}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("ChatService.Create returned %+v, expected %+v", got, expected)
	}

}

func TestChatService_Read(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Services/AC123", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		response := `{"sid":"AC123"}`

		fmt.Fprint(w, response)
	})

	got, err := client.Chat.Service.Read("AC123")

	if err != nil {
		t.Errorf("ChatService.Read returned error: %v", err)
	}

	expected := &ChatService{Sid: "AC123"}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("ChatService.Read returned %+v, expected %+v", got, expected)
	}

}

func TestChatService_Update(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Services/AC123", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		response := `{"sid":"AC123","friendly_name":"NewName","typing_indicator_timeout":10}`

		fmt.Fprint(w, response)
	})

	got, err := client.Chat.Service.Update("AC123", &ChatServiceParams{
		FriendlyName:           "NewName",
		TypingIndicatorTimeout: 10,
	})

	if err != nil {
		t.Errorf("ChatService.Update returned error: %v", err)
	}

	expected := &ChatService{Sid: "AC123", FriendlyName: "NewName", TypingIndicatorTimeout: 10}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("ChatService.Update returned %+v, expected %+v", got, expected)
	}

}

func TestChatService_Delete(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Services/AC123", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")

	})

	err := client.Chat.Service.Delete("AC123")

	if err != nil {
		t.Errorf("ChatService.Delete returned error: %v", err)
	}
}
