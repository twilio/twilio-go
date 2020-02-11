package twilio

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestAvailablePhoneNumberLocal_Read(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Accounts/AC123/AvailablePhoneNumbers/US/Local.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		response := `{"available_phone_numbers": [{"phone_number":"+18089251571"}]}`

		fmt.Fprint(w, response)
	})

	got, err := client.AvailablePhoneNumbers.Read(&AvailablePhoneNumberLocalReadParams{
		SMSEnabled: Bool(true),
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
