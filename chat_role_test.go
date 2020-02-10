package twilio

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestChatRole_Create(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Services/IS123/Roles", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testFormValues(t, r, values{"FriendlyName": "ChatRole"})
		response := `{"friendly_name":"ChatRole"}`

		fmt.Fprint(w, response)
	})

	got, err := client.Chat.Role.Create("IS123", &ChatRoleParams{FriendlyName: String("ChatRole")})

	if err != nil {
		t.Errorf("ChatRole.Create returned error: %v", err)
	}

	expected := &ChatRole{FriendlyName: String("ChatRole")}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("ChatRole.Create returned %+v, expected %+v", got, expected)
	}

}

func TestChatRole_Fetch(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Services/IS123/Roles/RL1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		response := `{"sid":"IS123"}`

		fmt.Fprint(w, response)
	})

	got, err := client.Chat.Role.Fetch("IS123", "RL1")

	if err != nil {
		t.Errorf("ChatRole.Fetch returned error: %v", err)
	}

	expected := &ChatRole{Sid: String("IS123")}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("ChatRole.Fetch returned %+v, expected %+v", got, expected)
	}

}

func TestChatRole_Read(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Services/IS123/Roles", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		response := `{"roles": [{"sid":"RL1"}]}`

		fmt.Fprint(w, response)
	})

	got, err := client.Chat.Role.Read("IS123")
	if err != nil {
		t.Errorf("ChatRole.Read returned error: %v", err)
	}

	c := &ChatRole{Sid: String("RL1")}
	expected := &ChatRoles{Roles: []*ChatRole{c}}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("ChatRole.Read returned %+v, expected %+v", got, expected)
	}

}

func TestChatRole_Update(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Services/IS123/Roles/RL1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		response := `{"sid":"RL1","friendly_name":"NewName"}`

		fmt.Fprint(w, response)
	})

	got, err := client.Chat.Role.Update("IS123", "RL1", &ChatRoleParams{
		FriendlyName: String("NewName"),
	})

	if err != nil {
		t.Errorf("ChatRole.Update returned error: %v", err)
	}

	expected := &ChatRole{Sid: String("RL1"), FriendlyName: String("NewName")}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("ChatRole.Update returned %+v, expected %+v", got, expected)
	}

}

func TestChatRole_Delete(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Services/IS123/Roles/RL1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")

	})

	err := client.Chat.Role.Delete("IS123", "RL1")

	if err != nil {
		t.Errorf("ChatRole.Delete returned error: %v", err)
	}
}
