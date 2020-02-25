package twilio

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"testing"
	"time"
)

func TestRuntimeService_marshall(t *testing.T) {
	testJSONMarshal(t, RuntimeService{}, "{}")

	got := &RuntimeService{
		SID:                String("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
		AccountSID:         String("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
		FriendlyName:       String("My New App"),
		UniqueName:         String("my-new-app"),
		IncludeCredentials: Bool(true),
		DateCreated:        &time.Time{},
		DateUpdated:        &time.Time{},
		URL:                String("https://serverless.twilio.com/v1/Services/ZS00000000000000000000000000000000"),
		Links: map[string]*string{
			"environments": String("https://serverless.twilio.com/v1/Services/ZS00000000000000000000000000000000/Environments"),
			"functions":    String("https://serverless.twilio.com/v1/Services/ZS00000000000000000000000000000000/Functions"),
			"assets":       String("https://serverless.twilio.com/v1/Services/ZS00000000000000000000000000000000/Assets"),
			"builds":       String("https://serverless.twilio.com/v1/Services/ZS00000000000000000000000000000000/Builds"),
		},
	}

	want := `{
		"sid": "ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		"account_sid": "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		"friendly_name": "My New App",
		"unique_name": "my-new-app",
		"include_credentials": true,
		"date_created": "0001-01-01T00:00:00Z",
		"date_updated": "0001-01-01T00:00:00Z",
		"url": "https://serverless.twilio.com/v1/Services/ZS00000000000000000000000000000000",
		"links": {
			"environments": "https://serverless.twilio.com/v1/Services/ZS00000000000000000000000000000000/Environments",
			"functions": "https://serverless.twilio.com/v1/Services/ZS00000000000000000000000000000000/Functions",
			"assets": "https://serverless.twilio.com/v1/Services/ZS00000000000000000000000000000000/Assets",
			"builds": "https://serverless.twilio.com/v1/Services/ZS00000000000000000000000000000000/Builds"
		}
	}`
	testJSONMarshal(t, got, want)
}

func TestRuntimeService_Create(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Services", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		f := url.Values{}
		f.Add("FriendlyName", "a")
		f.Add("UniqueName", "b")
		f.Add("IncludeCredentials", "true")

		testFormValues(t, r, f)
		response := `{"unique_name":"RuntimeService"}`

		fmt.Fprint(w, response)
	})

	got, err := client.Serverless.Service.Create(&RuntimeServiceParams{
		FriendlyName:       String("a"),
		UniqueName:         String("b"),
		IncludeCredentials: Bool(true),
	})

	if err != nil {
		t.Errorf("RuntimeService.Create returned error: %v", err)
	}

	want := &RuntimeService{UniqueName: String("RuntimeService")}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("RuntimeService.Create returned %+v, want %+v", got, want)
	}
}

func TestRuntimeService_Fetch(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Services/ZS123", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		response := `{"sid":"ZS123"}`

		fmt.Fprint(w, response)
	})

	got, err := client.Serverless.Service.Fetch("ZS123")

	if err != nil {
		t.Errorf("RuntimeService.Fetch returned error: %v", err)
	}

	want := &RuntimeService{SID: String("ZS123")}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("RuntimeService.Fetch returned %+v, want %+v", got, want)
	}
}

func TestRuntimeService_Update(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Services/ZS123", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		response := `{"sid":"ZS123","unique_name":"NewName"}`

		fmt.Fprint(w, response)
	})

	got, err := client.Serverless.Service.Update("ZS123", &RuntimeServiceParams{
		UniqueName: String("NewName"),
	})

	if err != nil {
		t.Errorf("RuntimeService.Update returned error: %v", err)
	}

	want := &RuntimeService{SID: String("ZS123"), UniqueName: String("NewName")}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("RuntimeService.Update returned %+v, want %+v", got, want)
	}
}

func TestRuntimeService_Delete(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Services/ZS123", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	err := client.Serverless.Service.Delete("ZS123")

	if err != nil {
		t.Errorf("RuntimeService.Delete returned error: %v", err)
	}
}
