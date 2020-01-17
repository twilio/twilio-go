package phonenumber

import (
	"encoding/json"
	"fmt"

	twilio "github.com/twilio/twilio-go"
)

type Client struct {
	Request twilio.Request
}

// Read returns available local phone numbers.
func (c Client) Read(sid string, params *twilio.AvailablePhoneNumberLocalReadParams) (*twilio.AvailablePhoneNumberLocal, error) {
	resp, err := c.Request.Get(fmt.Sprintf("/Services/%s", sid))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	pn := &twilio.AvailablePhoneNumberLocal{}
	if decodeErr := json.NewDecoder(resp.Body).Decode(pn); decodeErr != nil {
		return nil, decodeErr
	}

	return pn, err
}
