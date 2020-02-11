package twilio

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestChatRole_marshall(t *testing.T) {
	testJSONMarshal(t, &AvailablePhoneNumberLocal{}, "{}")

	c := &ChatRole{
		SID:          String("RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
		AccountSID:   String("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
		ServiceSID:   String("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
		FriendlyName: String("channel user"),
		Type:         String("channel"),
		Permissions:  []*string{String("sendMessage"), String("leaveChannel"), String("editOwnMessage"), String("deleteOwnMessage")},
		DateCreated:  &time.Time{},
		DateUpdated:  &time.Time{},
		URL:          String("https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Roles/RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
	}

	got := &ChatRoles{
		Roles: []*ChatRole{c},
		Meta: &Meta{
			Page:            Int(0),
			PageSize:        Int(50),
			FirstPageURL:    String("https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Roles?PageSize=50&Page=0"),
			PreviousPageURL: String("https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Roles?PageSize=50&Page=0"),
			URL:             String("https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Roles?PageSize=50&Page=0"),
			NextPageURL:     String("https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Roles?PageSize=50&Page=1"),
			Key:             String("roles"),
		},
	}

	want := `{
		"roles": [
			{
				"sid": "RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				"account_sid": "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				"service_sid": "ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				"friendly_name": "channel user",
				"type": "channel",
				"permissions": [
					"sendMessage",
					"leaveChannel",
					"editOwnMessage",
					"deleteOwnMessage"
				],
				"date_created": "0001-01-01T00:00:00Z",
				"date_updated": "0001-01-01T00:00:00Z",
				"url": "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Roles/RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
			}
		],
		"meta": {
			"page": 0,
			"page_size": 50,
			"first_page_url": "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Roles?PageSize=50&Page=0",
			"previous_page_url": "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Roles?PageSize=50&Page=0",
			"url": "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Roles?PageSize=50&Page=0",
			"next_page_url": "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Roles?PageSize=50&Page=1",
			"key": "roles"
		}
	}`
	testJSONMarshal(t, got, want)
}

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

	want := &ChatRole{FriendlyName: String("ChatRole")}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("ChatRole.Create returned %+v, want %+v", got, want)
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

	want := &ChatRole{SID: String("IS123")}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("ChatRole.Fetch returned %+v, want %+v", got, want)
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

	c := &ChatRole{SID: String("RL1")}
	want := &ChatRoles{Roles: []*ChatRole{c}}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("ChatRole.Read returned %+v, want %+v", got, want)
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

	want := &ChatRole{SID: String("RL1"), FriendlyName: String("NewName")}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("ChatRole.Update returned %+v, want %+v", got, want)
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
