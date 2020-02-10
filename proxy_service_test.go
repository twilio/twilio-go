package twilio

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestProxyService_Create(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Services", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testFormValues(t, r, values{"UniqueName": "ProxyChatService"})
		response := `{"unique_name":"ProxyChatService"}`

		fmt.Fprint(w, response)
	})

	got, err := client.Proxy.Service.Create(&ProxyServiceParams{UniqueName: String("ProxyChatService")})

	if err != nil {
		t.Errorf("ProxyService.Create returned error: %v", err)
	}

	expected := &ProxyService{UniqueName: String("ProxyChatService")}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("ProxyService.Create returned %+v, expected %+v", got, expected)
	}

}

func TestProxyService_Read(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Services/KS123", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		response := `{"sid":"KS123"}`

		fmt.Fprint(w, response)
	})

	got, err := client.Proxy.Service.Read("KS123", nil)

	if err != nil {
		t.Errorf("ProxyService.Read returned error: %v", err)
	}

	expected := &ProxyService{Sid: String("KS123")}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("ProxyService.Read returned %+v, expected %+v", got, expected)
	}

}

func TestProxyService_Update(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Services/KS123", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		response := `{"sid":"KS123","unique_name":"NewName","default_ttl":10}`

		fmt.Fprint(w, response)
	})

	got, err := client.Proxy.Service.Update("KS123", &ProxyServiceParams{
		UniqueName: String("NewName"),
		DefaultTTL: Int(10),
	})

	if err != nil {
		t.Errorf("ProxyService.Update returned error: %v", err)
	}

	expected := &ProxyService{Sid: String("KS123"), UniqueName: String("NewName"), DefaultTTL: Int(10)}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("ProxyService.Update returned %+v, expected %+v", got, expected)
	}

}

func TestProxyService_Delete(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Services/KS123", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")

	})

	err := client.Proxy.Service.Delete("KS123", nil)

	if err != nil {
		t.Errorf("ProxyService.Delete returned error: %v", err)
	}
}
