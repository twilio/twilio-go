package twilio

import (
	"encoding/json"
	"fmt"
)

// AvailablePhoneNumberLocal represents an available local phone number.
type AvailablePhoneNumberLocal struct {
	FriendlyName        *string          `json:"friendly_name"`
	PhoneNumber         *string          `json:"phone_number"`
	LATA                *int             `json:"lata,string"`
	Locality            *string          `json:"locality"`
	RateCenter          *string          `json:"rate_center"`
	Latitude            *float64         `json:"latitude,string"`
	Longitude           *float64         `json:"longitude,string"`
	Region              *string          `json:"region"`
	PostalCode          *string          `json:"postal_code"`
	ISOCountry          *string          `json:"iso_country"`
	AddressRequirements *string          `json:"address_requirements"`
	Beta                *bool            `json:"beta"`
	Capabilities        map[string]*bool `json:"capabilities"`
}

// AvailablePhoneNumbersLocal represents a paginated set of AvailablePhoneNumberLocal structs.
type AvailablePhoneNumbersLocal struct {
	NumPages              *int                         `json:"num_pages"`
	Page                  *int                         `json:"page"`
	PageSize              *int                         `json:"page_size"`
	Start                 *int                         `json:"start"`
	Total                 *int                         `json:"total"`
	End                   *int                         `json:"end"`
	FirstPageURI          *string                      `json:"first_page_uri"`
	LastPageURI           *string                      `json:"last_page_uri"`
	NextPageURI           *string                      `json:"next_page_uri"`
	URI                   *string                      `json:"uri"`
	PreviousPageURI       *string                      `json:"previous_page_uri"`
	AvailablePhoneNumbers []*AvailablePhoneNumberLocal `json:"available_phone_numbers"`
}

// AvailablePhoneNumberLocalReadParams is the set of parameters that can be used during a local phonenumber search.
type AvailablePhoneNumberLocalReadParams struct {
	// Read on AvailablePhoneNumberLocal requires account_sid in URI
	// Read on AvailablePhoneNumberLocal requires country_code in URI
	FaxEnabled                    *bool   `form:",omitempty"`
	SMSEnabled                    *bool   `form:"SmsEnabled,omitempty"`
	MMSEnabled                    *bool   `form:"MmsEnabled,omitempty"`
	VoiceEnabled                  *bool   `form:",omitempty"`
	ExcludeAllAddressRequired     *bool   `form:",omitempty"`
	ExcludeLocalAddressRequired   *bool   `form:",omitempty"`
	ExcludeForeignAddressRequired *bool   `form:",omitempty"`
	Beta                          *bool   `form:",omitempty"`
	Distance                      *int    `form:",omitempty"`
	AreaCode                      *int    `form:",omitempty"`
	InPostalCode                  *string `form:",omitempty"`
	NearNumber                    *string `form:",omitempty"`
	NearLatLong                   *string `form:",omitempty"`
	Contains                      *string `form:",omitempty"`
	InRegion                      *string `form:",omitempty"`
	InRateCenter                  *string `form:",omitempty"`
	InLATA                        *string `form:"InLata,omitempty"`
	InLocality                    *string `form:",omitempty"`
}

// availablePhoneNumbersClient is the entrypoint for the AvailablePhoneNumber Local resource.
// See: https://www.twilio.com/docs/phone-numbers/api/availablephonenumberlocal-resource
type availablePhoneNumbersClient struct {
	client  *Twilio
	baseURL string
}

// newAvailablePhoneNumbersClient constructs a new PhoneNumber client.
func newAvailablePhoneNumbersClient(client *Twilio) *availablePhoneNumbersClient {
	c := new(availablePhoneNumbersClient)
	c.client = client
	c.baseURL = fmt.Sprintf("https://api.%s/2010-04-01", c.client.BaseURL)

	return c
}

// Read returns available local phone numbers.
func (c *availablePhoneNumbersClient) Read(
	params *AvailablePhoneNumberLocalReadParams,
) (*AvailablePhoneNumbersLocal, error) {
	path := fmt.Sprintf("/Accounts/%s/AvailablePhoneNumbers/US/Local.json", c.client.AccountSID)

	resp, err := c.client.Get(c.url(path), params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	p := new(AvailablePhoneNumbersLocal)
	if err := json.NewDecoder(resp.Body).Decode(p); err != nil {
		return nil, err
	}

	return p, err
}

func (c *availablePhoneNumbersClient) url(path string) string {
	if c.client.defaultbaseURL != nil {
		return *c.client.defaultbaseURL + path
	}

	return c.baseURL + path
}
