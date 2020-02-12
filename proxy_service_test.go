package twilio

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestProxyService_marshall(t *testing.T) {
	testJSONMarshal(t, ProxyService{}, "{}")

	// testJSONMarshal(t, got, want)
}

func TestProxyService_Create(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Services", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testFormValues(t, r, values{"UniqueName": "ProxyService"})
		response := `{"unique_name":"ProxyService"}`

		fmt.Fprint(w, response)
	})

	got, err := client.Proxy.Service.Create(&ProxyServiceParams{UniqueName: String("ProxyService")})

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
