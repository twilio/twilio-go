// Package twilio provides bindings for Twilio's REST APIs.
package twilio

import (
	"github.com/twilio/twilio-go/client"
	studioV2 "github.com/twilio/twilio-go/rest/studio/v2"
	messagingV1 "github.com/twilio/twilio-go/rest/messaging/v1"
	numbersV2 "github.com/twilio/twilio-go/rest/numbers/v2"
	apiv2010 "github.com/twilio/twilio-go/rest/api/v2010"
	"time"
)

// Twilio provides access to Twilio services.
type Twilio struct {
	*client.Credentials
	*client.Client
	defaultbaseURL *string
	common         service
	StudioV2       *studioV2.DefaultApiService
	MessagingV1    *messagingV1.DefaultApiService
	NumbersV2    *numbersV2.DefaultApiService
	ApiV2010    *apiv2010.DefaultApiService
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
	credentials := &client.Credentials{AccountSID: accountSID, AuthToken: authToken}

	c := &Twilio{
		Credentials: credentials,
		Client: &client.Client{
			Credentials: credentials,
			BaseURL:     "twilio.com",
		},
	}

	c.common.client = c
	c.StudioV2 = studioV2.NewDefaultApiService(c.Client)
	c.MessagingV1 = messagingV1.NewDefaultApiService(c.Client)
	c.NumbersV2 = numbersV2.NewDefaultApiService(c.Client)
	c.ApiV2010 = apiv2010.NewDefaultApiService(c.Client)
	return c
}

// SetTimeout sets the Timeout for Twilio HTTP requests.
func (c *Twilio) SetTimeout(timeout time.Duration) {
	c.Client.SetTimeout(timeout)
}
