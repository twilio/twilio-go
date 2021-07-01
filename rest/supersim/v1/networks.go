/*
 * Twilio - Supersim
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

	"github.com/twilio/twilio-go/client"
)

// Fetch a Network resource.
func (c *ApiService) FetchNetwork(Sid string) (*SupersimV1Network, error) {
	path := "/v1/Networks/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &SupersimV1Network{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListNetwork'
type ListNetworkParams struct {
	// The [ISO country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) of the Network resources to read.
	IsoCountry *string `json:"IsoCountry,omitempty"`
	// The 'mobile country code' of a country. Network resources with this `mcc` in their `identifiers` will be read.
	Mcc *string `json:"Mcc,omitempty"`
	// The 'mobile network code' of a mobile operator network. Network resources with this `mnc` in their `identifiers` will be read.
	Mnc *string `json:"Mnc,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListNetworkParams) SetIsoCountry(IsoCountry string) *ListNetworkParams {
	params.IsoCountry = &IsoCountry
	return params
}
func (params *ListNetworkParams) SetMcc(Mcc string) *ListNetworkParams {
	params.Mcc = &Mcc
	return params
}
func (params *ListNetworkParams) SetMnc(Mnc string) *ListNetworkParams {
	params.Mnc = &Mnc
	return params
}
func (params *ListNetworkParams) SetPageSize(PageSize int) *ListNetworkParams {
	params.PageSize = &PageSize
	return params
}

// Retrieve a list of Network resources.
func (c *ApiService) ListNetwork(params *ListNetworkParams) (*ListNetworkResponse, error) {
	path := "/v1/Networks"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.IsoCountry != nil {
		data.Set("IsoCountry", *params.IsoCountry)
	}
	if params != nil && params.Mcc != nil {
		data.Set("Mcc", *params.Mcc)
	}
	if params != nil && params.Mnc != nil {
		data.Set("Mnc", *params.Mnc)
	}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListNetworkResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

//Retrieve a single page of Network records from the API. Request is executed immediately.
func (c *ApiService) NetworkPage(params *ListNetworkParams, pageToken string, pageNumber string, pageSize string) *client.Page {
	path := "/v1/Networks"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.IsoCountry != nil {
		data.Set("IsoCountry", *params.IsoCountry)
	}
	if params != nil && params.Mcc != nil {
		data.Set("Mcc", *params.Mcc)
	}
	if params != nil && params.Mnc != nil {
		data.Set("Mnc", *params.Mnc)
	}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	data.Set("PageToken", pageToken)
	data.Set("PageNumber", pageNumber)
	data.Set("PageSize", pageSize)

	response, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil
	}

	return client.NewPage(c.baseURL, response)
}

//Streams Network records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) NetworkStream(params *ListNetworkParams, meta client.PaginationData) chan map[string]interface{} {
	limits := c.requestHandler.ReadLimits(meta)
	page := c.NetworkPage(params, "", "", fmt.Sprint(limits.PageSize))
	return c.requestHandler.Stream(page, limits.Limit, limits.PageLimit)
}

//Lists Network records from the API as a list. Unlike stream, this operation is eager and will loads 'limit' records into memory before returning.
func (c *ApiService) NetworkList(params *ListNetworkParams, meta client.PaginationData) []interface{} {
	limits := c.requestHandler.ReadLimits(meta)
	page := c.NetworkPage(params, "", "", fmt.Sprint(limits.PageSize))
	return c.requestHandler.List(page, limits.Limit, limits.PageLimit)
}
