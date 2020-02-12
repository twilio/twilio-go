package twilio

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestProxyPhoneNumber_marshall(t *testing.T) {
	testJSONMarshal(t, &ProxyPhoneNumber{}, "{}")

	got := &ProxyPhoneNumber{
		SID:             String("PNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
		AccountSID:      String("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
		ProxyServiceSID: String("KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
		FriendlyName:    String("Friendly Name"),
		DateCreated:     &time.Time{},
		DateUpdated:     &time.Time{},
		PhoneNumber:     String("+1987654321"),
		ISOCountry:      String("US"),
		Capabilities: map[string]*bool{
			"sms_outbound":  Bool(true),
			"voice_inbound": Bool(false),
		},
		URL:        String("https://proxy.twilio.com/v1/Services/KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/PhoneNumbers/PNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
		IsReserved: Bool(false),
		InUse:      Int(0),
	}

	want := `{
		"sid": "PNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		"account_sid": "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		"service_sid": "KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		"date_created": "0001-01-01T00:00:00Z",
		"date_updated": "0001-01-01T00:00:00Z",
		"phone_number": "+1987654321",
		"friendly_name": "Friendly Name",
		"iso_country": "US",
		"capabilities": {
			"sms_outbound": true,
			"voice_inbound": false
		},
		"url": "https://proxy.twilio.com/v1/Services/KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/PhoneNumbers/PNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		"is_reserved": false,
		"in_use": 0
	}`

	testJSONMarshal(t, got, want)
}

func TestProxyPhoneNumber_Create(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Services/KS123/PhoneNumbers/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testFormValues(t, r, values{"Sid": "PN123"})
		response := `{"sid":"PN123"}`

		fmt.Fprint(w, response)
	})

	got, err := client.Proxy.PhoneNumber.Create("KS123", &ProxyPhoneNumberCreateParams{PhoneNumberSID: String("PN123")})

	if err != nil {
		t.Errorf("ProxyService.Create returned error: %v", err)
	}

	want := &ProxyPhoneNumber{SID: String("PN123")}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("ProxyService.Create returned %+v, want %+v", got, want)
	}

}

func TestProxyPhoneNumber_Fetch(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Services/KS123/PhoneNumbers/PN123", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		response := `{"sid":"KS123"}`

		fmt.Fprint(w, response)
	})

	got, err := client.Proxy.PhoneNumber.Fetch("KS123", "PN123")

	if err != nil {
		t.Errorf("ProxyService.Fetch returned error: %v", err)
	}

	want := &ProxyPhoneNumber{SID: String("KS123")}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("ProxyService.Fetch returned %+v, want %+v", got, want)
	}

}

func TestProxyPhoneNumber_Update(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Services/KS123/PhoneNumbers/PN123", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		response := `{"sid":"PN123","friendly_name":"NewName"}`

		fmt.Fprint(w, response)
	})

	got, err := client.Proxy.PhoneNumber.Update("KS123", "PN123", &ProxyPhoneNumberUpdateParams{
		PhoneNumberSID: String("PN124"),
	})

	if err != nil {
		t.Errorf("ProxyService.Update returned error: %v", err)
	}

	want := &ProxyPhoneNumber{SID: String("PN123"), FriendlyName: String("NewName")}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("ProxyService.Update returned %+v, want %+v", got, want)
	}

}

func TestProxyPhoneNumber_Delete(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Services/KS123/PhoneNumbers/PN123", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")

	})

	err := client.Proxy.PhoneNumber.Delete("KS123", "PN123")

	if err != nil {
		t.Errorf("ProxyService.Delete returned error: %v", err)
	}
}
