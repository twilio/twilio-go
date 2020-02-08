package twilio

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

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

	expected := &ProxyPhoneNumber{SID: "PN123"}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("ProxyService.Create returned %+v, expected %+v", got, expected)
	}

}

func TestProxyPhoneNumber_Read(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Services/KS123/PhoneNumbers/PN123", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		response := `{"sid":"KS123"}`

		fmt.Fprint(w, response)
	})

	got, err := client.Proxy.PhoneNumber.Read("KS123", "PN123")

	if err != nil {
		t.Errorf("ProxyService.Read returned error: %v", err)
	}

	expected := &ProxyPhoneNumber{SID: "KS123"}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("ProxyService.Read returned %+v, expected %+v", got, expected)
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

	expected := &ProxyPhoneNumber{SID: "PN123", FriendlyName: "NewName"}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("ProxyService.Update returned %+v, expected %+v", got, expected)
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
