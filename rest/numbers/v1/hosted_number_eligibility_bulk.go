/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Numbers
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

// Optional parameters for the method 'CreateBulkEligibility'
type CreateBulkEligibilityParams struct {
	//
	Body *map[string]interface{} `json:"body,omitempty"`
}

func (params *CreateBulkEligibilityParams) SetBody(Body map[string]interface{}) *CreateBulkEligibilityParams {
	params.Body = &Body
	return params
}

// Create a bulk eligibility check for a set of numbers that you want to host in Twilio.
func (c *ApiService) CreateBulkEligibility(params *CreateBulkEligibilityParams) (*NumbersV1BulkEligibility, error) {
	path := "/v1/HostedNumber/Eligibility/Bulk"

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/json",
	}

	body := []byte{}
	if params != nil && params.Body != nil {
		b, err := json.Marshal(*params.Body)
		if err != nil {
			return nil, err
		}
		body = b
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers, body...)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &NumbersV1BulkEligibility{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Fetch an eligibility bulk check that you requested to host in Twilio.
func (c *ApiService) FetchBulkEligibility(RequestId string) (*NumbersV1BulkEligibility, error) {
	path := "/v1/HostedNumber/Eligibility/Bulk/{RequestId}"
	path = strings.Replace(path, "{"+"RequestId"+"}", RequestId, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &NumbersV1BulkEligibility{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
