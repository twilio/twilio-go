/*
 * Twilio - Voice
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.18.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"

	"strings"
)

// Optional parameters for the method 'ListDialingPermissionsHrsPrefixes'
type ListDialingPermissionsHrsPrefixesParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListDialingPermissionsHrsPrefixesParams) SetPageSize(PageSize int) *ListDialingPermissionsHrsPrefixesParams {
	params.PageSize = &PageSize
	return params
}

// Fetch the high-risk special services prefixes from the country resource corresponding to the [ISO country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)
func (c *ApiService) ListDialingPermissionsHrsPrefixes(IsoCode string, params *ListDialingPermissionsHrsPrefixesParams) (*ListDialingPermissionsHrsPrefixesResponse, error) {
	path := "/v1/DialingPermissions/Countries/{IsoCode}/HighRiskSpecialPrefixes"
	path = strings.Replace(path, "{"+"IsoCode"+"}", IsoCode, -1)

	data := url.Values{}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListDialingPermissionsHrsPrefixesResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
