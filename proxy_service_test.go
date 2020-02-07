package twilio

import (
	"fmt"
	"net/http"
	"testing"
)

func TestProxyService_Create(t *testing.T) {
	client, mux, _, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Services", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		// testFormValues(t, r, values{"page": "2"})
		// fmt.Fprint(w, `[{"id":1}, {"id":2}]`)
	})

	c, _ := client.Proxy.Service.Create(&ProxyServiceParams{
		UniqueName: String("TESTING"),
	})

	fmt.Printf("FOO: %v", c)

}
