package twilio

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestIncomingPhoneNumber_marshall(t *testing.T) {
	testJSONMarshal(t, &AvailablePhoneNumberLocal{}, "{}")

	got := &IncomingPhoneNumber{
		AccountSID:          String("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
		AddressRequirements: String("none"),
		AddressSID:          String("ADXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
		APIVersion:          String("2010-04-01"),
		Beta:                Bool(false),
		Capabilities: map[string]*bool{
			"mms":   Bool(true),
			"sms":   Bool(false),
			"voice": Bool(true),
		},
		DateCreated:          &time.Time{},
		DateUpdated:          &time.Time{},
		EmergencyStatus:      String("Active"),
		EmergencyAddressSID:  String("ADXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
		FriendlyName:         String("(808) 925-5327"),
		IdentitySID:          String("RIXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
		Origin:               String("origin"),
		PhoneNumber:          String("+18089255327"),
		SID:                  String("PNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
		SMSApplicationSID:    nil,
		SMSFallbackMethod:    String("POST"),
		SMSFallbackURL:       String(""),
		SMSMethod:            String("POST"),
		SMSURL:               String(""),
		StatusCallback:       String(""),
		StatusCallbackMethod: String("POST"),
		TrunkSID:             nil,
		URI:                  String("/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/IncomingPhoneNumbers/PNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json"),
		VoiceApplicationSID:  nil,
		VoiceCallerIDLookup:  Bool(false),
		VoiceFallbackMethod:  String("POST"),
		VoiceFallbackURL:     nil,
		VoiceMethod:          String("POST"),
		VoiceURL:             nil,
		BundleSID:            String("BUXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
	}

	want := `{
		"account_sid": "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		"address_requirements": "none",
		"address_sid": "ADXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		"api_version": "2010-04-01",
		"beta": false,
		"capabilities": {
		  "mms": true,
		  "sms": false,
		  "voice": true
		},
		"date_created": "0001-01-01T00:00:00Z",
		"date_updated": "0001-01-01T00:00:00Z",
		"emergency_status": "Active",
		"emergency_address_sid": "ADXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		"friendly_name": "(808) 925-5327",
		"identity_sid": "RIXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		"origin": "origin",
		"phone_number": "+18089255327",
		"sid": "PNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		"sms_application_sid": null,
		"sms_fallback_method": "POST",
		"sms_fallback_url": "",
		"sms_method": "POST",
		"sms_url": "",
		"status_callback": "",
		"status_callback_method": "POST",
		"trunk_sid": null,
		"uri": "/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/IncomingPhoneNumbers/PNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json",
		"voice_application_sid": null,
		"voice_caller_id_lookup": false,
		"voice_fallback_method": "POST",
		"voice_fallback_url": null,
		"voice_method": "POST",
		"voice_url": null,
		"bundle_sid": "BUXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
	}`

	testJSONMarshal(t, got, want)
}

func TestIncomingPhoneNumber_Create(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Accounts/AC123/IncomingPhoneNumbers.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testFormValues(t, r, values{"FriendlyName": "IncomingPhoneNumber"})
		response := `{"friendly_name":"IncomingPhoneNumber"}`

		fmt.Fprint(w, response)
	})

	got, err := client.IncomingPhoneNumbers.Create(&IncomingPhoneNumberParams{FriendlyName: String("IncomingPhoneNumber")})

	if err != nil {
		t.Errorf("IncomingPhoneNumber.Create returned error: %v", err)
	}

	want := &IncomingPhoneNumber{FriendlyName: String("IncomingPhoneNumber")}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("IncomingPhoneNumber.Create returned %+v, want %+v", got, want)
	}

}

func TestIncomingPhoneNumber_Read(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Accounts/AC123/IncomingPhoneNumbers/PN1.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		response := `{"sid":"PN1"}`

		fmt.Fprint(w, response)
	})

	got, err := client.IncomingPhoneNumbers.Read("PN1")

	if err != nil {
		t.Errorf("IncomingPhoneNumber.Read returned error: %v", err)
	}

	want := &IncomingPhoneNumber{SID: String("PN1")}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("IncomingPhoneNumber.Read returned %+v, want %+v", got, want)
	}

}

func TestIncomingPhoneNumber_Update(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Accounts/AC123/IncomingPhoneNumbers/PN1.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		response := `{"sid":"PN1","friendly_name":"NewName"}`

		fmt.Fprint(w, response)
	})

	got, err := client.IncomingPhoneNumbers.Update("PN1", &IncomingPhoneNumberParams{
		FriendlyName: String("NewName"),
	})

	if err != nil {
		t.Errorf("IncomingPhoneNumber.Update returned error: %v", err)
	}

	want := &IncomingPhoneNumber{SID: String("PN1"), FriendlyName: String("NewName")}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("IncomingPhoneNumber.Update returned %+v, want %+v", got, want)
	}

}

func TestIncomingPhoneNumber_Delete(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Accounts/AC123/IncomingPhoneNumbers/PN1.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	err := client.IncomingPhoneNumbers.Delete("PN1")

	if err != nil {
		t.Errorf("IncomingPhoneNumber.Delete returned error: %v", err)
	}
}
