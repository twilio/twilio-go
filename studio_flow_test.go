package twilio

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"testing"
	"time"
)

func TestStudioFlow_marshall(t *testing.T) {
	testJSONMarshal(t, StudioFlow{}, "{}")

	got := &StudioFlow{
		SID:           String("FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
		AccountSID:    String("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
		FriendlyName:  String("Main IVR"),
		Status:        String("draft"),
		Revision:      Int(1),
		CommitMessage: String("First draft"),
		Valid:         Bool(true),
		WebhookURL:    String("/Flows/FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
		DateCreated:   &time.Time{},
		DateUpdated:   nil,
		URL:           String("/FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
		Links: map[string]*string{
			"test_users": String("/FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TestUsers"),
			"revisions":  String("/FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Revisions"),
		},
	}

	want := `{
		"sid": "FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		"account_sid": "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		"definition": null,
		"friendly_name": "Main IVR",
		"status": "draft",
		"revision": 1,
		"commit_message": "First draft",
		"valid": true,
		"errors": null,
		"webhook_url": "/Flows/FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		"date_created": "0001-01-01T00:00:00Z",
		"date_updated": null,
		"url": "/FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		"links": {
		  "test_users": "/FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TestUsers",
		  "revisions": "/FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Revisions"
		}
	  }`
	testJSONMarshal(t, got, want)
}

func TestStudioFlow_Create(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Flows", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		f := url.Values{}
		f.Add("FriendlyName", "a")
		f.Add("Status", "b")
		f.Add("Definition", "c")
		f.Add("CommitMessage", "d")
		testFormValues(t, r, f)
		response := `{"friendly_name":"StudioFlow"}`

		fmt.Fprint(w, response)
	})

	got, err := client.Studio.Flow.Create(&StudioFlowParams{
		FriendlyName:  String("a"),
		Status:        String("b"),
		Definition:    String("c"),
		CommitMessage: String("d"),
	})

	if err != nil {
		t.Errorf("StudioFlow.Create returned error: %v", err)
	}

	want := &StudioFlow{FriendlyName: String("StudioFlow")}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("StudioFlow.Create returned %+v, want %+v", got, want)
	}
}

func TestStudioFlow_Read(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Flows/FW123", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		response := `{"sid":"FW123"}`

		fmt.Fprint(w, response)
	})

	got, err := client.Studio.Flow.Read("FW123")

	if err != nil {
		t.Errorf("StudioFlow.Read returned error: %v", err)
	}

	want := &StudioFlow{SID: String("FW123")}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("StudioFlow.Read returned %+v, want %+v", got, want)
	}
}

func TestStudioFlow_Update(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Flows/FW123", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		response := `{"sid":"FW123","friendly_name":"NewName"}`

		fmt.Fprint(w, response)
	})

	got, err := client.Studio.Flow.Update("FW123", &StudioFlowParams{
		FriendlyName: String("NewName"),
	})

	if err != nil {
		t.Errorf("StudioFlow.Update returned error: %v", err)
	}

	want := &StudioFlow{SID: String("FW123"), FriendlyName: String("NewName")}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("StudioFlow.Update returned %+v, want %+v", got, want)
	}
}

func TestStudioFlow_Delete(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Flows/FW123", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	err := client.Studio.Flow.Delete("FW123")

	if err != nil {
		t.Errorf("StudioFlow.Delete returned error: %v", err)
	}
}
