package twilio

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestAvailablePhoneNumberLocal_marshall(t *testing.T) { //nolint
	testJSONMarshal(t, &AvailablePhoneNumberLocal{}, "{}")

	a := &AvailablePhoneNumberLocal{
		AddressRequirements: String("none"),
		Beta:                Bool(false),
		Capabilities: map[string]*bool{
			"mms":   Bool(true),
			"sms":   Bool(false),
			"voice": Bool(true),
		},
		FriendlyName: String("(808) 925-1571"),
		ISOCountry:   String("US"),
		LATA:         Int(834),
		Latitude:     Float64(19.720000),
		Locality:     String("Hilo"),
		Longitude:    Float64(-155.090000),
		PhoneNumber:  String("+18089251571"),
		PostalCode:   String("96720"),
		RateCenter:   String("HILO"),
		Region:       String("HI"),
	}

	got := &AvailablePhoneNumbersLocal{
		AvailablePhoneNumbers: []*AvailablePhoneNumberLocal{a},
		End:                   Int(1),
		FirstPageURI:          String("/AvailablePhoneNumbers/US/Local.json?PageSize=50&Page=0"),
		LastPageURI:           String("/AvailablePhoneNumbers/US/Local.json?PageSize=50&Page=0"),
		NextPageURI:           String("/AvailablePhoneNumbers/US/Local.json?PageSize=50&Page=50"),
		NumPages:              Int(1),
		Page:                  Int(0),
		PageSize:              Int(50),
		PreviousPageURI:       String("/AvailablePhoneNumbers/US/Local.json?PageSize=50&Page=0"),
		Start:                 Int(0),
		Total:                 Int(1),
		URI:                   String("/AvailablePhoneNumbers/US/Local.json?PageSize=1"),
	}
	want := `{
		"available_phone_numbers": [
		  {
			"address_requirements": "none",
			"beta": false,
			"capabilities": {
			  "mms": true,
			  "sms": false,
			  "voice": true
			},
			"friendly_name": "(808) 925-1571",
			"iso_country": "US",
			"lata": "834",
			"latitude": "19.720000",
			"locality": "Hilo",
			"longitude": "-155.090000",
			"phone_number": "+18089251571",
			"postal_code": "96720",
			"rate_center": "HILO",
			"region": "HI"
		  }
		],
		"end": 1,
		"first_page_uri": "/AvailablePhoneNumbers/US/Local.json?PageSize=50&Page=0",
		"last_page_uri": "/AvailablePhoneNumbers/US/Local.json?PageSize=50&Page=0",
		"next_page_uri": "/AvailablePhoneNumbers/US/Local.json?PageSize=50&Page=50",
		"num_pages": 1,
		"page": 0,
		"page_size": 50,
		"previous_page_uri": "/AvailablePhoneNumbers/US/Local.json?PageSize=50&Page=0",
		"start": 0,
		"total": 1,
		"uri": "/AvailablePhoneNumbers/US/Local.json?PageSize=1"
	}`
	testJSONMarshal(t, got, want)
}

func TestAvailablePhoneNumberLocal_Read(t *testing.T) { //nolint
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Accounts/AC123/AvailablePhoneNumbers/US/Local.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testQueryValues(t, r, values{
			"FaxEnabled":                    "true",
			"SmsEnabled":                    "true",
			"MmsEnabled":                    "true",
			"VoiceEnabled":                  "true",
			"ExcludeAllAddressRequired":     "true",
			"ExcludeLocalAddressRequired":   "true",
			"ExcludeForeignAddressRequired": "true",
			"Beta":                          "false",
			"Distance":                      "10",
			"AreaCode":                      "0",
			"InPostalCode":                  "a",
			"NearNumber":                    "b",
			"NearLatLong":                   "c",
			"Contains":                      "d",
			"InRegion":                      "e",
			"InRateCenter":                  "f",
			"InLata":                        "g",
			"InLocality":                    "h",
		})
		response := `{"available_phone_numbers": [{"phone_number":"+18089251571"}]}`
		fmt.Fprint(w, response)
	})

	got, err := client.AvailablePhoneNumbers.Read(&AvailablePhoneNumberLocalReadParams{
		FaxEnabled:                    Bool(true),
		SMSEnabled:                    Bool(true),
		MMSEnabled:                    Bool(true),
		VoiceEnabled:                  Bool(true),
		ExcludeAllAddressRequired:     Bool(true),
		ExcludeLocalAddressRequired:   Bool(true),
		ExcludeForeignAddressRequired: Bool(true),
		Beta:                          Bool(false),
		Distance:                      Int(10),
		AreaCode:                      Int(0),
		InPostalCode:                  String("a"),
		NearNumber:                    String("b"),
		NearLatLong:                   String("c"),
		Contains:                      String("d"),
		InRegion:                      String("e"),
		InRateCenter:                  String("f"),
		InLATA:                        String("g"),
		InLocality:                    String("h"),
	})
	if err != nil {
		t.Errorf("AvailablePhoneNumberLocal.Read returned error: %v", err)
	}

	c := &AvailablePhoneNumberLocal{PhoneNumber: String("+18089251571")}
	want := &AvailablePhoneNumbersLocal{AvailablePhoneNumbers: []*AvailablePhoneNumberLocal{c}}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("AvailablePhoneNumberLocal.Read returned %+v, want %+v", got, want)
	}
}
