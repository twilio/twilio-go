// Package twilio provides bindings for Twilio's REST APIs.
package twilio

import (
	"net/http"
	"time"

	client "github.com/twilio/twilio-go/client"
	studio "github.com/twilio/twilio-go/studio"
)

// Twilio provides access to Twilio services.
type Twilio struct {
	*client.Credentials
	*client.Client
	defaultbaseURL *string
	common         service
	Studio         *studio.ApiService
}

type service struct {
	client *Twilio
}

// Meta holds relevant pagination resources.
type Meta struct {
	FirstPageURL    *string `json:"first_page_url"`
	Key             *string `json:"key"`
	LastPageURL     *string `json:"last_page_url,omitempty"`
	NextPageURL     *string `json:"next_page_url"`
	Page            *int    `json:"page"`
	PageSize        *int    `json:"page_size"`
	PreviousPageURL *string `json:"previous_page_url"`
	URL             *string `json:"url"`
}

// NewClient provides an initialized Twilio client.
func NewClient(accountSID string, authToken string) *Twilio {
	var httpClient = &http.Client{
		Timeout: time.Second * 10,
	}

	credentials := &client.Credentials{AccountSID: accountSID, AuthToken: authToken}

	c := &Twilio{
		Credentials: credentials,
		Client: &client.Client{
			Credentials: credentials,
			HTTPClient:  httpClient,
			BaseURL:     "twilio.com",
		},
	}

	c.common.client = c
	c.Studio = studio.NewApiService(c.Client)

	return c
}
