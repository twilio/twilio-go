package twilio

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

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

	expected := &IncomingPhoneNumber{FriendlyName: String("IncomingPhoneNumber")}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("IncomingPhoneNumber.Create returned %+v, expected %+v", got, expected)
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

	expected := &IncomingPhoneNumber{Sid: String("PN1")}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("IncomingPhoneNumber.Read returned %+v, expected %+v", got, expected)
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

	expected := &IncomingPhoneNumber{Sid: String("PN1"), FriendlyName: String("NewName")}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("IncomingPhoneNumber.Update returned %+v, expected %+v", got, expected)
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
