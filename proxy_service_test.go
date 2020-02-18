package twilio

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"testing"
	"time"
)

func TestProxyService_marshall(t *testing.T) {
	testJSONMarshal(t, ProxyService{}, "{}")

	got := &ProxyService{
		SID:                     String("KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
		AccountSID:              String("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
		ChatInstanceSID:         String("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
		UniqueName:              String("My Service"),
		DefaultTTL:              Int(3600),
		CallbackURL:             String("http://www.example.com"),
		GeoMatchLevel:           String("country"),
		NumberSelectionBehavior: String("prefer_sticky"),
		InterceptCallbackURL:    String("http://www.example.com"),
		OutOfSessionCallbackURL: String("http://www.example.com"),
		DateCreated:             &time.Time{},
		DateUpdated:             &time.Time{},
		URL:                     String("https://proxy.twilio.com/v1/Services/KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
		Links: map[string]*string{
			"sessions":      String("https://proxy.twilio.com/v1/Services/KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Sessions"),
			"phone_numbers": String("https://proxy.twilio.com/v1/Services/KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/PhoneNumbers"),
			"short_codes":   String("https://proxy.twilio.com/v1/Services/KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/ShortCodes"),
		},
	}

	want := `{
		"sid": "KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		"account_sid": "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		"chat_instance_sid": "ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		"unique_name": "My Service",
		"default_ttl": 3600,
		"callback_url": "http://www.example.com",
		"geo_match_level": "country",
		"number_selection_behavior": "prefer_sticky",
		"intercept_callback_url": "http://www.example.com",
		"out_of_session_callback_url": "http://www.example.com",
		"date_created": "0001-01-01T00:00:00Z",
		"date_updated": "0001-01-01T00:00:00Z",
		"url": "https://proxy.twilio.com/v1/Services/KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		"links": {
			"sessions": "https://proxy.twilio.com/v1/Services/KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Sessions",
			"phone_numbers": "https://proxy.twilio.com/v1/Services/KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/PhoneNumbers",
			"short_codes": "https://proxy.twilio.com/v1/Services/KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/ShortCodes"
		}
	}`
	testJSONMarshal(t, got, want)
}

func TestProxyService_Create(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Services", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		f := url.Values{}
		f.Add("ChatInstanceSid", "a")
		f.Add("UniqueName", "b")
		f.Add("DefaultTtl", "10")
		f.Add("CallbackUrl", "c")
		f.Add("GeoMatchLevel", "d")
		f.Add("NumberSelectionBehavior", "e")
		f.Add("InterceptCallbackUrl", "f")
		f.Add("OutOfSessionCallbackUrl", "g")

		testFormValues(t, r, f)
		response := `{"unique_name":"ProxyService"}`

		fmt.Fprint(w, response)
	})

	got, err := client.Proxy.Service.Create(&ProxyServiceParams{
		ChatInstanceSID:         String("a"),
		UniqueName:              String("b"),
		DefaultTTL:              Int(10),
		CallbackURL:             String("c"),
		GeoMatchLevel:           String("d"),
		NumberSelectionBehavior: String("e"),
		InterceptCallbackURL:    String("f"),
		OutOfSessionCallbackURL: String("g"),
	})

	if err != nil {
		t.Errorf("ProxyService.Create returned error: %v", err)
	}

	want := &ProxyService{UniqueName: String("ProxyService")}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("ProxyService.Create returned %+v, want %+v", got, want)
	}
}

func TestProxyService_Fetch(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Services/KS123", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		response := `{"sid":"KS123"}`

		fmt.Fprint(w, response)
	})

	got, err := client.Proxy.Service.Fetch("KS123")

	if err != nil {
		t.Errorf("ProxyService.Fetch returned error: %v", err)
	}

	want := &ProxyService{SID: String("KS123")}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("ProxyService.Fetch returned %+v, want %+v", got, want)
	}
}

func TestProxyService_Update(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Services/KS123", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		response := `{"sid":"KS123","unique_name":"NewName"}`

		fmt.Fprint(w, response)
	})

	got, err := client.Proxy.Service.Update("KS123", &ProxyServiceParams{
		UniqueName: String("NewName"),
	})

	if err != nil {
		t.Errorf("ProxyService.Update returned error: %v", err)
	}

	want := &ProxyService{SID: String("KS123"), UniqueName: String("NewName")}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("ProxyService.Update returned %+v, want %+v", got, want)
	}
}

func TestProxyService_Delete(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Services/KS123", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	err := client.Proxy.Service.Delete("KS123")

	if err != nil {
		t.Errorf("ProxyService.Delete returned error: %v", err)
	}
}
