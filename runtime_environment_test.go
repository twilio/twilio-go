package twilio

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"testing"
	"time"
)

func TestRuntimeEnvironment_marshall(t *testing.T) {
	testJSONMarshal(t, &RuntimeEnvironment{}, "{}")

	got := &RuntimeEnvironment{
		SID:          String("ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
		AccountSID:   String("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
		ServiceSID:   String("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
		BuildSID:     nil,
		UniqueName:   String("staging"),
		DomainSuffix: String("stage"),
		DomainName:   String("foobar-1234-stage.twil.io"),
		DateCreated:  &time.Time{},
		DateUpdated:  &time.Time{},
		URL:          String("/Environments/ZE00000000000000000000000000000000"),
		Links: map[string]*string{
			"variables":   String("/Environments/ZE00000000000000000000000000000000/Variables"),
			"deployments": String("/Environments/ZE00000000000000000000000000000000/Deployments"),
			"logs":        String("/Environments/ZE00000000000000000000000000000000/Logs"),
		},
	}

	want := `{
		"sid": "ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		"account_sid": "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		"service_sid": "ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		"build_sid": null,
		"unique_name": "staging",
		"domain_suffix": "stage",
		"domain_name": "foobar-1234-stage.twil.io",
		"date_created": "2018-11-10T20:00:00Z",
		"date_updated": "2018-11-10T20:00:00Z",
		"url": "/Environments/ZE00000000000000000000000000000000",
		"links": {
			"variables": "/Environments/ZE00000000000000000000000000000000/Variables",
			"deployments": "/Environments/ZE00000000000000000000000000000000/Deployments",
			"logs": "/Environments/ZE00000000000000000000000000000000/Logs"
		}
	}`

	testJSONMarshal(t, got, want)
}

func TestRuntimeEnvironment_Create(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Services/ZS123/Environments", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		f := url.Values{}
		f.Add("UniqueName", "foo")
		f.Add("DomainSuffix", "bar")
		testFormValues(t, r, f)

		response := `{"sid":"ZE123"}`

		fmt.Fprint(w, response)
	})

	got, err := client.Serverless.Environment.Create("ZS123", &RuntimeEnvironmentParams{
		UniqueName:   String("foo"),
		DomainSuffix: String("bar"),
	})

	if err != nil {
		t.Errorf("RuntimeEnvironment.Create returned error: %v", err)
	}

	want := &RuntimeEnvironment{SID: String("ZE123")}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("RuntimeEnvironment.Create returned %+v, want %+v", got, want)
	}
}

func TestRuntimeEnvironment_Fetch(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Services/ZS123/Environments/ZE123", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		response := `{"sid":"ZS123"}`

		fmt.Fprint(w, response)
	})

	got, err := client.Serverless.Environment.Fetch("ZS123", "ZE123")

	if err != nil {
		t.Errorf("RuntimeEnvironment.Fetch returned error: %v", err)
	}

	want := &RuntimeEnvironment{SID: String("ZS123")}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("RuntimeEnvironment.Fetch returned %+v, want %+v", got, want)
	}
}

func TestRuntimeEnvironment_Delete(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Services/ZS123/Environments/ZE123", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	err := client.Serverless.Environment.Delete("ZS123", "ZE123")

	if err != nil {
		t.Errorf("RuntimeEnvironment.Delete returned error: %v", err)
	}
}
