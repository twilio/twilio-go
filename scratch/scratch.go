package main

import (
	"encoding/json"
	"fmt"

	twilio "github.com/twilio/twilio-go/client"
)

func main() {
	client := twilio.NewClient("ACa44bde7854a694c6bc5245f9da9b41cd", "f22025da58fdd63e7cd7c6e47fd065ca")
	foo, _ := client.ProxyService.Delete("KSae27df9e1e6fe1a8268afeb7f2172bc5", nil)
	s, err := json.MarshalIndent(foo, "", "\t")

	fmt.Printf("err, %+v", err)

	fmt.Printf("Proxyasdfasdf, %s", string(s))
}
