/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Bulk Messaging and Broadcast
 * Bulk Sending is a public Twilio REST API for 1:Many Message creation up to 100 recipients. Broadcast is a public Twilio REST API for 1:Many Message creation up to 10,000 recipients via file upload.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"

    "github.com/twilio/twilio-go/client"
)


// Optional parameters for the method 'CreateBroadcast'
type CreateBroadcastParams struct {
    // Idempotency key provided by the client
    XTwilioRequestKey *string `json:"X-Twilio-Request-Key,omitempty"`
}

func (params *CreateBroadcastParams) SetXTwilioRequestKey(XTwilioRequestKey string) (*CreateBroadcastParams){
    params.XTwilioRequestKey = &XTwilioRequestKey
    return params
}

// Create a new Broadcast
func (c *ApiService) CreateBroadcast(params *CreateBroadcastParams) (*MessagingV1Broadcast, error) {
    path := "/v1/Broadcasts"
    
data := url.Values{}
headers := make(map[string]interface{})



	if params != nil && params.XTwilioRequestKey != nil {
		headers["X-Twilio-Request-Key"] = *params.XTwilioRequestKey
	}

    resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &MessagingV1Broadcast{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}
