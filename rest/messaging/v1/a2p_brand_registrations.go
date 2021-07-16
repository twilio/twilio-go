/*
 * Twilio - Messaging
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.19.0
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

// Optional parameters for the method 'CreateBrandRegistrations'
type CreateBrandRegistrationsParams struct {
	// A2P Messaging Profile Bundle Sid.
	A2pProfileBundleSid *string `json:"A2pProfileBundleSid,omitempty"`
	// Customer Profile Bundle Sid.
	CustomerProfileBundleSid *string `json:"CustomerProfileBundleSid,omitempty"`
}

func (params *CreateBrandRegistrationsParams) SetA2pProfileBundleSid(A2pProfileBundleSid string) *CreateBrandRegistrationsParams {
	params.A2pProfileBundleSid = &A2pProfileBundleSid
	return params
}
func (params *CreateBrandRegistrationsParams) SetCustomerProfileBundleSid(CustomerProfileBundleSid string) *CreateBrandRegistrationsParams {
	params.CustomerProfileBundleSid = &CustomerProfileBundleSid
	return params
}

func (c *ApiService) CreateBrandRegistrations(params *CreateBrandRegistrationsParams) (*MessagingV1BrandRegistrations, error) {
	path := "/v1/a2p/BrandRegistrations"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.A2pProfileBundleSid != nil {
		data.Set("A2pProfileBundleSid", *params.A2pProfileBundleSid)
	}
	if params != nil && params.CustomerProfileBundleSid != nil {
		data.Set("CustomerProfileBundleSid", *params.CustomerProfileBundleSid)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &MessagingV1BrandRegistrations{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

func (c *ApiService) FetchBrandRegistrations(Sid string) (*MessagingV1BrandRegistrations, error) {
	path := "/v1/a2p/BrandRegistrations/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &MessagingV1BrandRegistrations{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListBrandRegistrations'
type ListBrandRegistrationsParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListBrandRegistrationsParams) SetPageSize(PageSize int) *ListBrandRegistrationsParams {
	params.PageSize = &PageSize
	return params
}

// Retrieve a single page of BrandRegistrations records from the API. Request is executed immediately.
func (c *ApiService) PageBrandRegistrations(params *ListBrandRegistrationsParams, pageToken string, pageNumber string) (*ListBrandRegistrationsResponse, error) {
	path := "/v1/a2p/BrandRegistrations"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	if pageToken != "" {
		data.Set("PageToken", pageToken)
	}
	if pageToken != "" {
		data.Set("Page", pageNumber)
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListBrandRegistrationsResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists BrandRegistrations records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListBrandRegistrations(params *ListBrandRegistrationsParams, limit int) ([]MessagingV1BrandRegistrations, error) {
	params.SetPageSize(client.ReadLimits(params.PageSize, limit))

	response, err := c.PageBrandRegistrations(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []MessagingV1BrandRegistrations

	for response != nil {
		records = append(records, response.Data...)

		var record interface{}
		if record, err = client.GetNext(response, &curRecord, limit, c.getNextListBrandRegistrationsResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListBrandRegistrationsResponse)
	}

	return records, err
}

// Streams BrandRegistrations records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamBrandRegistrations(params *ListBrandRegistrationsParams, limit int) (chan MessagingV1BrandRegistrations, error) {
	params.SetPageSize(client.ReadLimits(params.PageSize, limit))

	response, err := c.PageBrandRegistrations(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan MessagingV1BrandRegistrations, 1)

	go func() {
		for response != nil {
			for item := range response.Data {
				channel <- response.Data[item]
			}

			var record interface{}
			if record, err = client.GetNext(response, &curRecord, limit, c.getNextListBrandRegistrationsResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListBrandRegistrationsResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListBrandRegistrationsResponse(nextPageUri string) (interface{}, error) {
	if nextPageUri == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(c.baseURL+nextPageUri, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListBrandRegistrationsResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
