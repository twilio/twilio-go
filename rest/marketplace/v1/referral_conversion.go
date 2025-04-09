/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Marketplace
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"encoding/json"
	"net/url"
)

// Optional parameters for the method 'CreateReferralConversion'
type CreateReferralConversionParams struct {
	//
	CreateReferralConversionRequest *CreateReferralConversionRequest `json:"CreateReferralConversionRequest,omitempty"`
}

func (params *CreateReferralConversionParams) SetCreateReferralConversionRequest(CreateReferralConversionRequest CreateReferralConversionRequest) *CreateReferralConversionParams {
	params.CreateReferralConversionRequest = &CreateReferralConversionRequest
	return params
}

func (c *ApiService) CreateReferralConversion(params *CreateReferralConversionParams) (*MarketplaceV1ReferralConversion, error) {
	path := "/v1/ReferralConversion"

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/json",
	}

	body := []byte{}
	if params != nil && params.CreateReferralConversionRequest != nil {
		b, err := json.Marshal(*params.CreateReferralConversionRequest)
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

	ps := &MarketplaceV1ReferralConversion{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
