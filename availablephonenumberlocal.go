package twilio

import (
	"encoding/json"
	"fmt"

	twilio "github.com/twilio/twilio-go/internal"
)

// AvailablePhoneNumberLocal represents an available local phone number.
type AvailablePhoneNumberLocal struct {
	FriendlyName        string          `json:"friendly_name"`
	PhoneNumber         string          `json:"phone_number"`
	Lata                int             `json:"lata,string"`
	Locality            string          `json:"locality"`
	RateCenter          string          `json:"rate_center"`
	Latitude            float64         `json:"latitude,string"`
	Longitude           float64         `json:"longitude,string"`
	Region              string          `json:"region"`
	PostalCode          string          `json:"postal_code"`
	ISOCountry          string          `json:"iso_country"`
	AddressRequirements string          `json:"address_requirements"`
	Beta                bool            `json:"beta"`
	Capabilities        map[string]bool `json:"capabilities"`
}

// AvailablePhoneNumbersLocal represents a paginated set of AvailablePhoneNumberLocal structs.
type AvailablePhoneNumbersLocal struct {
	NumPages              int                          `json:"num_pages"`
	Page                  int                          `json:"page"`
	PageSize              int                          `json:"page_size"`
	Start                 int                          `json:"start"`
	Total                 int                          `json:"total"`
	End                   int                          `json:"end"`
	FirstPageURI          string                       `json:"first_page_uri"`
	LastPageURI           string                       `json:"last_page_uri"`
	NextPageURI           string                       `json:"next_page_uri"`
	URI                   string                       `json:"uri"`
	PreviousPageURI       string                       `json:"previous_page_uri"`
	AvailablePhoneNumbers []*AvailablePhoneNumberLocal `json:"available_phone_numbers"`
}

// AvailablePhoneNumberLocalReadParams is the set of parameters that can be used during a local phonenumber search.
type AvailablePhoneNumberLocalReadParams struct {
	// Read on AvailablePhoneNumberLocal requires account_sid in URI
	// Read on AvailablePhoneNumberLocal requires country_code in URI
	FaxEnabled                    bool   `url:"FaxEnabled,omitempty"`
	SmsEnabled                    bool   `url:"SmsEnabled,omitempty"`
	MmsEnabled                    bool   `url:"MmsEnabled,omitempty"`
	VoiceEnabled                  bool   `url:"VoiceEnabled,omitempty"`
	ExcludeAllAddressRequired     bool   `url:"ExcludeAllAddressRequired,omitempty"`
	ExcludeLocalAddressRequired   bool   `url:"ExcludeLocalAddressRequired,omitempty"`
	ExcludeForeignAddressRequired bool   `url:"ExcludeForeignAddressRequired,omitempty"`
	Beta                          bool   `url:"Beta,omitempty"`
	Distance                      int    `url:"Distance,omitempty"`
	AreaCode                      int    `url:"AreaCode,omitempty"`
	InPostalCode                  string `url:"InPostalCode,omitempty"`
	NearNumber                    string `url:"NearNumber,omitempty"`
	NearLatLong                   string `url:"NearLatLong,omitempty"`
	Contains                      string `url:"Contains,omitempty"`
	InRegion                      string `url:"InRegion,omitempty"`
	InRateCenter                  string `url:"InRateCenter,omitempty"`
	InLata                        string `url:"InLata,omitempty"`
	InLocality                    string `url:"InLocality,omitempty"`
}

//
// PhoneNumberClient is the entrypoint for the AvailablePhoneNumber Local resource.
// See: https://www.twilio.com/docs/phone-numbers/api/availablephonenumberlocal-resource
type PhoneNumberClient struct {
	client     *twilio.Client
	serviceURL string
}

// Initialize constructs a new TaskRouter client.
func NewPhoneNumberClient(client *twilio.Client) *PhoneNumberClient {
	pn := new(PhoneNumberClient)
	pn.client = client
	pn.serviceURL = fmt.Sprintf("https://api.%s/2010-04-01", pn.client.BaseURL)
	return pn
}

// ReadAvailableLocalPhoneNumbers returns available local phone numbers.
func (pn PhoneNumberClient) ReadAvailableLocalPhoneNumbers(
	params *AvailablePhoneNumberLocalReadParams) (*AvailablePhoneNumbersLocal, error) {
	url := fmt.Sprintf("%s/Accounts/%s/AvailablePhoneNumbers/US/Local.json", pn.serviceURL, pn.client.AccountSid)

	resp, err := pn.client.Get(url, params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	avPNL := new(AvailablePhoneNumbersLocal)
	if decodeErr := json.NewDecoder(resp.Body).Decode(avPNL); decodeErr != nil {
		return nil, decodeErr
	}

	return avPNL, err
}
