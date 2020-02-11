package twilio

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestStudioFlow_Create(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Flows", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testFormValues(t, r, values{"FriendlyName": "StudioFlow"})
		response := `{"friendly_name":"StudioFlow"}`

		fmt.Fprint(w, response)
	})

	got, err := client.Studio.Flow.Create(&StudioFlowParams{FriendlyName: String("StudioFlow")})

	if err != nil {
		t.Errorf("StudioFlow.Create returned error: %v", err)
	}

	want := &StudioFlow{FriendlyName: String("StudioFlow")}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("StudioFlow.Create returned %+v, want %+v", got, want)
	}

}

func TestStudioFlow_Read(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Flows/FW123", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		response := `{"sid":"FW123"}`

		fmt.Fprint(w, response)
	})

	got, err := client.Studio.Flow.Read("FW123")

	if err != nil {
		t.Errorf("StudioFlow.Read returned error: %v", err)
	}

	want := &StudioFlow{SID: String("FW123")}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("StudioFlow.Read returned %+v, want %+v", got, want)
	}

}

func TestStudioFlow_Update(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Flows/FW123", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		response := `{"sid":"FW123","friendly_name":"NewName"}`

		fmt.Fprint(w, response)
	})

	got, err := client.Studio.Flow.Update("FW123", &StudioFlowParams{
		FriendlyName: String("NewName"),
	})

	if err != nil {
		t.Errorf("StudioFlow.Update returned error: %v", err)
	}

	want := &StudioFlow{SID: String("FW123"), FriendlyName: String("NewName")}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("StudioFlow.Update returned %+v, want %+v", got, want)
	}

}

func TestStudioFlow_Delete(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Flows/FW123", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")

	})

	err := client.Studio.Flow.Delete("FW123")

	if err != nil {
		t.Errorf("StudioFlow.Delete returned error: %v", err)
	}
}
