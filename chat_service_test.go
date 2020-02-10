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
		response := `{"unique_name":"ChatService"}`

		fmt.Fprint(w, response)
	})

	got, err := client.Proxy.Service.Create(&ChatServiceParams{FriendlyName: String("ChatService")})

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

	mux.HandleFunc("/Services/KS123", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		response := `{"sid":"KS123"}`

		fmt.Fprint(w, response)
	})

	got, err := client.Proxy.Service.Read("KS123", nil)

	if err != nil {
		t.Errorf("ChatService.Read returned error: %v", err)
	}

	expected := &ChatService{Sid: "KS123"}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("ChatService.Read returned %+v, expected %+v", got, expected)
	}

}

func TestChatService_Update(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Services/KS123", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		response := `{"sid":"KS123","unique_name":"NewName","default_ttl":10}`

		fmt.Fprint(w, response)
	})

	got, err := client.Proxy.Service.Update("KS123", &ChatServiceParams{
		FriendlyName: String("NewName"),
		DefaultTTL:   Int(10),
	})

	if err != nil {
		t.Errorf("ChatService.Update returned error: %v", err)
	}

	expected := &ChatService{Sid: "KS123", FriendlyName: "NewName", DefaultTTL: 10}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("ChatService.Update returned %+v, expected %+v", got, expected)
	}

}

func TestChatService_Delete(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Services/KS123", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")

	})

	err := client.Proxy.Service.Delete("KS123", nil)

	if err != nil {
		t.Errorf("ChatService.Delete returned error: %v", err)
	}
}
