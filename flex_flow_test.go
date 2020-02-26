package twilio

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"testing"
	"time"
)

func TestFlexFlow_marshall(t *testing.T) { //nolint
	testJSONMarshal(t, &FlexFlow{}, "{}")

	c := &FlexFlow{
		AccountSID:      String("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
		DateCreated:     &time.Time{},
		DateUpdated:     &time.Time{},
		FriendlyName:    String("friendly_name"),
		SID:             String("FOXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
		URL:             String("https://flex-api.twilio.com/v1/FlexFlows/FOXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
		ChatServiceSID:  String("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
		ChannelType:     String("sms"),
		ContactIdentity: String("+15555555555"),
		Enabled:         Bool(false),
		IntegrationType: String("studio"),
		Integration: &Integration{FlowSID: String("FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
			RetryCount: Int(3),
		},
		LongLived:      Bool(false),
		JanitorEnabled: Bool(false),
	}

	got := &FlexFlowList{
		FlexFlows: []*FlexFlow{c},
		Meta: &Meta{
			Page:            Int(0),
			PageSize:        Int(20),
			FirstPageURL:    String("https://flex-api.twilio.com/v1/FlexFlows?PageSize=20&Page=0"),
			PreviousPageURL: nil,
			URL:             String("https://flex-api.twilio.com/v1/FlexFlows?PageSize=20&Page=0"),
			NextPageURL:     nil,
			Key:             String("flex_flows"),
		},
	}

	want := `{
	"flex_flows": [{
		"account_sid": "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		"date_created": "0001-01-01T00:00:00Z",
		"date_updated": "0001-01-01T00:00:00Z",
		"friendly_name": "friendly_name",
		"sid": "FOXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		"url": "https://flex-api.twilio.com/v1/FlexFlows/FOXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		"chat_service_sid": "ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		"channel_type": "sms",
		"contact_identity": "+15555555555",
		"enabled": false,
		"integration_type": "studio",
		"integration": {
			"retry_count": 3,
			"flow_sid": "FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
		},
		"long_lived": false,
		"janitor_enabled": false	
	}],
	"meta": {
		"page": 0,
		"page_size": 20,
		"first_page_url": "https://flex-api.twilio.com/v1/FlexFlows?PageSize=20&Page=0",
		"previous_page_url": null,
		"url": "https://flex-api.twilio.com/v1/FlexFlows?PageSize=20&Page=0",
		"next_page_url": null,
		"key": "flex_flows"
	}
	}`
	testJSONMarshal(t, got, want)
}

func TestFlexFlow_Create(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		f := url.Values{}
		f.Add("FriendlyName", "FlexFlow")
		f.Add("Url", "https://flex-api.twilio.com/v1/FlexFlows/FOXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
		f.Add("ChatServiceSid", "ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
		f.Add("ChannelType", "sms")
		f.Add("ContactIdentity", "+15555555555")
		f.Add("Enabled", "false")
		f.Add("IntegrationType", "studio")
		f.Add("Integration.RetryCount", "3")
		f.Add("Integration.FlowSid", "FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
		f.Add("LongLived", "false")
		f.Add("JanitorEnabled", "false")

		testFormValues(t, r, f)
		response := `{"friendly_name":"FlexFlow"}`

		fmt.Fprint(w, response)
	})

	got, err := client.FlexFlow.Create(&FlexFlowParams{
		FriendlyName:    String("FlexFlow"),
		URL:             String("https://flex-api.twilio.com/v1/FlexFlows/FOXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
		ChatServiceSID:  String("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
		ChannelType:     String("sms"),
		ContactIdentity: String("+15555555555"),
		Enabled:         Bool(false),
		IntegrationType: String("studio"),
		Integration: &Integration{
			RetryCount: Int(3),
			FlowSID:    String("FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
		},
		LongLived:      Bool(false),
		JanitorEnabled: Bool(false),
	})

	if err != nil {
		t.Errorf("FlexFlow.Create returned error: %v", err)
	}

	want := &FlexFlow{FriendlyName: String("FlexFlow")}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("FlexFlow.Create returned %+v, want %+v", got, want)
	}
}

func TestFlexFlow_Fetch(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/FOXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		response := `{"sid":"FOXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"}`

		fmt.Fprint(w, response)
	})

	got, err := client.FlexFlow.Fetch("FOXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

	if err != nil {
		t.Errorf("FlexFlow.Fetch returned error: %v", err)
	}

	want := &FlexFlow{SID: String("FOXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("FlexFlow.Fetch returned %+v, want %+v", got, want)
	}
}

func TestFlexFlow_Read(t *testing.T) { //nolint
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testQueryValues(t, r, values{"FriendlyName": "friendly_name"})
		response := `{"flex_flows": [{"sid":"FOXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"}]}`
		fmt.Fprint(w, response)
	})

	got, err := client.FlexFlow.Read("friendly_name")
	if err != nil {
		t.Errorf("FlexFlow.Read returned error: %v", err)
	}

	c := &FlexFlow{SID: String("FOXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")}
	want := &FlexFlowList{FlexFlows: []*FlexFlow{c}}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("FlexFlow.Read returned %+v, want %+v", got, want)
	}
}

func TestFlexFlow_Update(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/FOXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		response := `{"sid":"AC123","friendly_name":"NewName"}`

		fmt.Fprint(w, response)
	})

	got, err := client.FlexFlow.Update("FOXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", &FlexFlowParams{
		FriendlyName: String("NewName"),
	})

	if err != nil {
		t.Errorf("FlexFlow.Update returned error: %v", err)
	}

	want := &FlexFlow{SID: String("AC123"), FriendlyName: String("NewName")}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("FlexFlow.Update returned %+v, want %+v", got, want)
	}
}

func TestFlexFlow_Delete(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/FOXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	err := client.FlexFlow.Delete("FOXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

	if err != nil {
		t.Errorf("FlexFlow.Delete returned error: %v", err)
	}
}
