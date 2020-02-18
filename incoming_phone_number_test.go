package twilio

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"testing"
	"time"
)

func TestIncomingPhoneNumber_marshall(t *testing.T) { //nolint
	testJSONMarshal(t, &IncomingPhoneNumber{}, "{}")

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
		URI:                  String("/IncomingPhoneNumbers/PNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json"),
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
		"uri": "/IncomingPhoneNumbers/PNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json",
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

func TestIncomingPhoneNumber_Create(t *testing.T) { //nolint
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Accounts/AC123/IncomingPhoneNumbers.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		f := url.Values{}
		f.Add("ApiVersion", "v1")
		f.Add("FriendlyName", "a")
		f.Add("SmsApplicationSid", "b")
		f.Add("SmsFallbackMethod", "c")
		f.Add("PhoneNumber", "d")
		f.Add("AreaCode", "e")
		f.Add("SmsFallbackUrl", "f")
		f.Add("SmsMethod", "g")
		f.Add("SmsUrl", "h")
		f.Add("StatusCallback", "i")
		f.Add("AddressSid", "j")
		f.Add("StatusCallbackMethod", "k")
		f.Add("VoiceApplicationSid", "l")
		f.Add("VoiceCallerIdLookup", "true")
		f.Add("VoiceFallbackMethod", "m")
		f.Add("VoiceFallbackUrl", "n")
		f.Add("VoiceMethod", "o")
		f.Add("VoiceUrl", "p")
		f.Add("VoiceReceiveMode", "q")
		f.Add("EmergencyStatus", "r")
		f.Add("EmergencyAddressSid", "s")
		f.Add("BundleSid", "t")
		f.Add("IdentitySid", "u")
		f.Add("TrunkSid", "v")

		testFormValues(t, r, f)
		response := `{"friendly_name":"IncomingPhoneNumber"}`

		fmt.Fprint(w, response)
	})

	got, err := client.IncomingPhoneNumbers.Create(&IncomingPhoneNumberParams{
		APIVersion:           String("v1"),
		FriendlyName:         String("a"),
		SMSApplicationSID:    String("b"),
		SMSFallbackMethod:    String("c"),
		PhoneNumber:          String("d"),
		AreaCode:             String("e"),
		SMSFallbackURL:       String("f"),
		SMSMethod:            String("g"),
		SMSURL:               String("h"),
		StatusCallback:       String("i"),
		AddressSID:           String("j"),
		StatusCallbackMethod: String("k"),
		VoiceApplicationSID:  String("l"),
		VoiceCallerIDLookup:  Bool(true),
		VoiceFallbackMethod:  String("m"),
		VoiceFallbackURL:     String("n"),
		VoiceMethod:          String("o"),
		VoiceURL:             String("p"),
		VoiceReceiveMode:     String("q"),
		EmergencyStatus:      String("r"),
		EmergencyAddressSID:  String("s"),
		BundleSID:            String("t"),
		IdentitySID:          String("u"),
		TrunkSID:             String("v"),
	})

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
