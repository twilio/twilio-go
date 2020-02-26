package twilio

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"testing"
	"time"
)

func TestSyncService_marshall(t *testing.T) {
	testJSONMarshal(t, SyncService{}, "{}")

	got := &SyncService{
		SID:          String("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
		AccountSID:   String("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
		FriendlyName: String("friendly_name"),
		UniqueName:   String("my-new-app"),
		DateCreated:  &time.Time{},
		DateUpdated:  &time.Time{},
		URL:          String("https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
		Links: map[string]*string{
			"documents": String("https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Documents"),
		},
		WebhookURL:                    String("http://www.example.com"),
		WebhooksFromRestEnabled:       Bool(false),
		ReachabilityWebhooksEnabled:   Bool(false),
		ACLEnabled:                    Bool(true),
		ReachabilityDebouncingEnabled: Bool(false),
		ReachabilityDebouncingWindow:  Int(5000),
	}

	want := `{
		"sid": "ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		"account_sid": "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		"friendly_name": "friendly_name",
		"unique_name": "my-new-app",
		"date_created": "0001-01-01T00:00:00Z",
		"date_updated": "0001-01-01T00:00:00Z",
		"url": "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		"links": {
			"documents": "https://sync.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Documents"
		},
		"webhook_url": "http://www.example.com",
		"webhooks_from_rest_enabled": false,
		"reachability_webhooks_enabled": false,
		"acl_enabled": true,
		"reachability_debouncing_enabled": false,
		"reachability_debouncing_window": 5000
	}`
	testJSONMarshal(t, got, want)
}

func TestSyncService_Create(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Services", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		f := url.Values{}
		f.Add("FriendlyName", "a")
		f.Add("WebhookUrl", "b")
		f.Add("ReachabilityWebhooksEnabled", "false")
		f.Add("AclEnabled", "false")
		f.Add("ReachabilityDebouncingEnabled", "false")
		f.Add("ReachabilityDebouncingWindow", "5000")
		f.Add("WebhooksFromRestEnabled", "true")

		testFormValues(t, r, f)
		response := `{"friendly_name":"SyncService"}`

		fmt.Fprint(w, response)
	})

	got, err := client.Sync.Service.Create(&SyncServiceParams{
		FriendlyName:                  String("a"),
		WebhookURL:                    String("b"),
		ReachabilityWebhooksEnabled:   Bool(false),
		ACLEnabled:                    Bool(false),
		ReachabilityDebouncingEnabled: Bool(false),
		ReachabilityDebouncingWindow:  Int(5000),
		WebhooksFromRestEnabled:       Bool(true),
	})

	if err != nil {
		t.Errorf("SyncService.Create returned error: %v", err)
	}

	want := &SyncService{FriendlyName: String("SyncService")}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("SyncService.Create returned %+v, want %+v", got, want)
	}
}

func TestSyncService_Fetch(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Services/IS123", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		response := `{"sid":"IS123"}`

		fmt.Fprint(w, response)
	})

	got, err := client.Sync.Service.Fetch("IS123")

	if err != nil {
		t.Errorf("SyncService.Fetch returned error: %v", err)
	}

	want := &SyncService{SID: String("IS123")}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("SyncService.Fetch returned %+v, want %+v", got, want)
	}
}

func TestSyncService_Update(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Services/IS123", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		response := `{"sid":"IS123","friendly_name":"NewName"}`

		fmt.Fprint(w, response)
	})

	got, err := client.Sync.Service.Update("IS123", &SyncServiceParams{
		FriendlyName: String("NewName"),
	})

	if err != nil {
		t.Errorf("SyncService.Update returned error: %v", err)
	}

	want := &SyncService{SID: String("IS123"), FriendlyName: String("NewName")}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("SyncService.Update returned %+v, want %+v", got, want)
	}
}

func TestSyncService_Delete(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Services/IS123", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	err := client.Sync.Service.Delete("IS123")

	if err != nil {
		t.Errorf("SyncService.Delete returned error: %v", err)
	}
}
