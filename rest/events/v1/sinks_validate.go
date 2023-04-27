/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Events
 * This is the public Twilio REST API.
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


// Optional parameters for the method 'CreateSinkValidate'
type CreateSinkValidateParams struct {
    // A 34 character string that uniquely identifies the test event for a Sink being validated.
    TestId *string `json:"TestId,omitempty"`
}

func (params *CreateSinkValidateParams) SetTestId(TestId string) (*CreateSinkValidateParams){
    params.TestId = &TestId
    return params
}

// Validate that a test event for a Sink was received.
func (c *ApiService) CreateSinkValidate(Sid string, params *CreateSinkValidateParams) (*EventsV1SinkValidate, error) {
    path := "/v1/Sinks/{Sid}/Validate"
        path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

data := url.Values{}
headers := make(map[string]interface{})

if params != nil && params.TestId != nil {
    data.Set("TestId", *params.TestId)
}



    resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &EventsV1SinkValidate{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}
