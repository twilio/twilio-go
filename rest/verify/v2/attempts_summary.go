/*
 * Twilio - Verify
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.29.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

// Optional parameters for the method 'FetchVerificationAttemptsSummary'
type FetchVerificationAttemptsSummaryParams struct {
	// Filter used to consider only Verification Attempts of the given verify service on the summary aggregation.
	VerifyServiceSid *string `json:"VerifyServiceSid,omitempty"`
	// Datetime filter used to consider only Verification Attempts created after this datetime on the summary aggregation. Given as GMT in RFC 2822 format.
	DateCreatedAfter *time.Time `json:"DateCreatedAfter,omitempty"`
	// Datetime filter used to consider only Verification Attempts created before this datetime on the summary aggregation. Given as GMT in RFC 2822 format.
	DateCreatedBefore *time.Time `json:"DateCreatedBefore,omitempty"`
	// Filter used to consider only Verification Attempts sent to the specified destination country on the summary aggregation.
	Country *string `json:"Country,omitempty"`
	// Filter Verification Attempts considered on the summary aggregation by communication channel. Valid values are `SMS` and `CALL`
	Channel *string `json:"Channel,omitempty"`
	// Filter the Verification Attempts considered on the summary aggregation by Destination prefix. It is the prefix of a phone number in E.164 format.
	DestinationPrefix *string `json:"DestinationPrefix,omitempty"`
}

func (params *FetchVerificationAttemptsSummaryParams) SetVerifyServiceSid(VerifyServiceSid string) *FetchVerificationAttemptsSummaryParams {
	params.VerifyServiceSid = &VerifyServiceSid
	return params
}
func (params *FetchVerificationAttemptsSummaryParams) SetDateCreatedAfter(DateCreatedAfter time.Time) *FetchVerificationAttemptsSummaryParams {
	params.DateCreatedAfter = &DateCreatedAfter
	return params
}
func (params *FetchVerificationAttemptsSummaryParams) SetDateCreatedBefore(DateCreatedBefore time.Time) *FetchVerificationAttemptsSummaryParams {
	params.DateCreatedBefore = &DateCreatedBefore
	return params
}
func (params *FetchVerificationAttemptsSummaryParams) SetCountry(Country string) *FetchVerificationAttemptsSummaryParams {
	params.Country = &Country
	return params
}
func (params *FetchVerificationAttemptsSummaryParams) SetChannel(Channel string) *FetchVerificationAttemptsSummaryParams {
	params.Channel = &Channel
	return params
}
func (params *FetchVerificationAttemptsSummaryParams) SetDestinationPrefix(DestinationPrefix string) *FetchVerificationAttemptsSummaryParams {
	params.DestinationPrefix = &DestinationPrefix
	return params
}

// Get a summary of how many attempts were made and how many were converted.
func (c *ApiService) FetchVerificationAttemptsSummary(params *FetchVerificationAttemptsSummaryParams) (*VerifyV2VerificationAttemptsSummary, error) {
	path := "/v2/Attempts/Summary"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.VerifyServiceSid != nil {
		data.Set("VerifyServiceSid", *params.VerifyServiceSid)
	}
	if params != nil && params.DateCreatedAfter != nil {
		data.Set("DateCreatedAfter", fmt.Sprint((*params.DateCreatedAfter).Format(time.RFC3339)))
	}
	if params != nil && params.DateCreatedBefore != nil {
		data.Set("DateCreatedBefore", fmt.Sprint((*params.DateCreatedBefore).Format(time.RFC3339)))
	}
	if params != nil && params.Country != nil {
		data.Set("Country", *params.Country)
	}
	if params != nil && params.Channel != nil {
		data.Set("Channel", *params.Channel)
	}
	if params != nil && params.DestinationPrefix != nil {
		data.Set("DestinationPrefix", *params.DestinationPrefix)
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VerifyV2VerificationAttemptsSummary{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
