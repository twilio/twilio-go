package twilio

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestChatService_marshall(t *testing.T) {
	testJSONMarshal(t, &AvailablePhoneNumberLocal{}, "{}")

	c := &ChatService{
		AccountSID:                   String("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
		ConsumptionReportInterval:    Int(10),
		DateCreated:                  &time.Time{},
		DateUpdated:                  &time.Time{},
		DefaultChannelCreatorRoleSID: String("RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
		DefaultChannelRoleSID:        String("RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
		DefaultServiceRoleSID:        String("RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
		FriendlyName:                 String("friendly_name"),
		Limits: map[string]*int{
			"channel_members": Int(100),
			"user_channels":   Int(250),
		},
		Links: map[string]*string{
			"channels": String("https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels"),
			"users":    String("https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Users"),
			"roles":    String("https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Roles"),
			"bindings": String("https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Bindings"),
		},
		Notifications: &Notifications{
			AddedToChannel: &BaseNotification{
				Sound:    String("String"),
				Enabled:  Bool(false),
				Template: String("Template"),
			},
			LogEnabled: Bool(true),
		},
		PostWebhookURL:         String("post_webhook_url"),
		PreWebhookURL:          String("pre_webhook_url"),
		PreWebhookRetryCount:   Int(2),
		PostWebhookRetryCount:  Int(3),
		ReachabilityEnabled:    Bool(false),
		ReadStatusEnabled:      Bool(false),
		SID:                    String("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
		TypingIndicatorTimeout: Int(100),
		URL:                    String("https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
		WebhookFilters:         []*string{String("webhook_filters")},
		WebhookMethod:          String("webhook_method"),
		Media: &Media{
			SizeLimitMB:          Int(150),
			CompatibilityMessage: String("media compatibility message"),
		},
	}

	got := &ChatServiceList{
		Services: []*ChatService{c},
		Meta: &Meta{
			Page:            Int(0),
			PageSize:        Int(50),
			FirstPageURL:    String("https://chat.twilio.com/v2/Services?PageSize=50&Page=0"),
			PreviousPageURL: String("https://chat.twilio.com/v2/Services?PageSize=50&Page=0"),
			URL:             String("https://chat.twilio.com/v2/Services?PageSize=50&Page=0"),
			NextPageURL:     String("https://chat.twilio.com/v2/Services?PageSize=50&Page=1"),
			Key:             String("services"),
		},
	}

	want := `{
		"services": [
			{
				"account_sid": "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				"consumption_report_interval": 10,
				"date_created": "0001-01-01T00:00:00Z",
				"date_updated": "0001-01-01T00:00:00Z",
				"default_channel_creator_role_sid": "RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				"default_channel_role_sid": "RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				"default_service_role_sid": "RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				"friendly_name": "friendly_name",
				"limits": {
				  "channel_members": 100,
				  "user_channels": 250
				},
				"links": {
				  "channels": "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels",
				  "users": "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Users",
				  "roles": "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Roles",
				  "bindings": "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Bindings"
				},
				"notifications": {
					"removed_from_channel": null,
					"log_enabled": true,
					"added_to_channel": {
						"enabled": false,
						"template": "Template",
						"sound": "String"
					},
					"new_message": null,
					"invited_to_channel": null
				},
				"post_webhook_url": "post_webhook_url",
				"pre_webhook_url": "pre_webhook_url",
				"pre_webhook_retry_count": 2,
				"post_webhook_retry_count": 3,
				"reachability_enabled": false,
				"read_status_enabled": false,
				"sid": "ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				"typing_indicator_timeout": 100,
				"url": "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				"webhook_filters": [
				  "webhook_filters"
				],
				"webhook_method": "webhook_method",
				"media": {
				  "size_limit_mb": 150,
				  "compatibility_message": "media compatibility message"
				}
			}
		],
		"meta": {
			"page": 0,
			"page_size": 50,
			"first_page_url": "https://chat.twilio.com/v2/Services?PageSize=50&Page=0",
			"previous_page_url": "https://chat.twilio.com/v2/Services?PageSize=50&Page=0",
			"url": "https://chat.twilio.com/v2/Services?PageSize=50&Page=0",
			"next_page_url": "https://chat.twilio.com/v2/Services?PageSize=50&Page=1",
			"key": "services"
		}
	}`
	testJSONMarshal(t, got, want)
}

func TestChatService_Create(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Services", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testFormValues(t, r, values{"FriendlyName": "ChatService"})
		response := `{"friendly_name":"ChatService"}`

		fmt.Fprint(w, response)
	})

	got, err := client.Chat.Service.Create(&ChatServiceParams{FriendlyName: String("ChatService")})

	if err != nil {
		t.Errorf("ChatService.Create returned error: %v", err)
	}

	want := &ChatService{FriendlyName: String("ChatService")}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("ChatService.Create returned %+v, want %+v", got, want)
	}

}

func TestChatService_Fetch(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Services/AC123", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		response := `{"sid":"AC123"}`

		fmt.Fprint(w, response)
	})

	got, err := client.Chat.Service.Fetch("AC123")

	if err != nil {
		t.Errorf("ChatService.Fetch returned error: %v", err)
	}

	want := &ChatService{SID: String("AC123")}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("ChatService.Fetch returned %+v, want %+v", got, want)
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
		FriendlyName:           String("NewName"),
		TypingIndicatorTimeout: Int(10),
	})

	if err != nil {
		t.Errorf("ChatService.Update returned error: %v", err)
	}

	want := &ChatService{SID: String("AC123"), FriendlyName: String("NewName"), TypingIndicatorTimeout: Int(10)}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("ChatService.Update returned %+v, want %+v", got, want)
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
