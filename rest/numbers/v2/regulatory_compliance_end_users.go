/*
 * Twilio - Numbers
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

// Optional parameters for the method 'CreateEndUser'
type CreateEndUserParams struct {
	// The set of parameters that are the attributes of the End User resource which are derived End User Types.
	Attributes *map[string]interface{} `json:"Attributes,omitempty"`
	// The string that you assigned to describe the resource.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The type of end user of the Bundle resource - can be `individual` or `business`.
	Type *string `json:"Type,omitempty"`
}

func (params *CreateEndUserParams) SetAttributes(Attributes map[string]interface{}) *CreateEndUserParams {
	params.Attributes = &Attributes
	return params
}
func (params *CreateEndUserParams) SetFriendlyName(FriendlyName string) *CreateEndUserParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *CreateEndUserParams) SetType(Type string) *CreateEndUserParams {
	params.Type = &Type
	return params
}

// Create a new End User.
func (c *ApiService) CreateEndUser(params *CreateEndUserParams) (*NumbersV2RegulatoryComplianceEndUser, error) {
	path := "/v2/RegulatoryCompliance/EndUsers"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Attributes != nil {
		v, err := json.Marshal(params.Attributes)

		if err != nil {
			return nil, err
		}

		data.Set("Attributes", string(v))
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.Type != nil {
		data.Set("Type", *params.Type)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &NumbersV2RegulatoryComplianceEndUser{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Delete a specific End User.
func (c *ApiService) DeleteEndUser(Sid string) error {
	path := "/v2/RegulatoryCompliance/EndUsers/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

// Fetch specific End User Instance.
func (c *ApiService) FetchEndUser(Sid string) (*NumbersV2RegulatoryComplianceEndUser, error) {
	path := "/v2/RegulatoryCompliance/EndUsers/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &NumbersV2RegulatoryComplianceEndUser{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListEndUser'
type ListEndUserParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListEndUserParams) SetPageSize(PageSize int) *ListEndUserParams {
	params.PageSize = &PageSize
	return params
}

//Retrieve a single page of EndUser records from the API. Request is executed immediately.
func (c *ApiService) PageEndUser(params *ListEndUserParams, pageToken string, pageNumber string) (*ListEndUserResponse, error) {
	path := "/v2/RegulatoryCompliance/EndUsers"

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

	ps := &ListEndUserResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

//Lists EndUser records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListEndUser(params *ListEndUserParams, limit *int) ([]*ListEndUserResponse, error) {
	params.SetPageSize(client.ReadLimits(params.PageSize, limit))

	response, err := c.PageEndUser(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []*ListEndUserResponse

	for response != nil {
		records = append(records, response)

		var record interface{}
		if record, err = client.GetNext(response, &curRecord, limit, c.getNextListEndUserResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListEndUserResponse)
	}

	return records, err
}

//Streams EndUser records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamEndUser(params *ListEndUserParams, limit *int) (chan *ListEndUserResponse, error) {
	params.SetPageSize(client.ReadLimits(params.PageSize, limit))

	response, err := c.PageEndUser(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan *ListEndUserResponse, 1)

	go func() {
		for response != nil {
			channel <- response

			var record interface{}
			if record, err = client.GetNext(response, &curRecord, limit, c.getNextListEndUserResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListEndUserResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListEndUserResponse(nextPageUri string) (interface{}, error) {
	if nextPageUri == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(c.baseURL+nextPageUri, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListEndUserResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateEndUser'
type UpdateEndUserParams struct {
	// The set of parameters that are the attributes of the End User resource which are derived End User Types.
	Attributes *map[string]interface{} `json:"Attributes,omitempty"`
	// The string that you assigned to describe the resource.
	FriendlyName *string `json:"FriendlyName,omitempty"`
}

func (params *UpdateEndUserParams) SetAttributes(Attributes map[string]interface{}) *UpdateEndUserParams {
	params.Attributes = &Attributes
	return params
}
func (params *UpdateEndUserParams) SetFriendlyName(FriendlyName string) *UpdateEndUserParams {
	params.FriendlyName = &FriendlyName
	return params
}

// Update an existing End User.
func (c *ApiService) UpdateEndUser(Sid string, params *UpdateEndUserParams) (*NumbersV2RegulatoryComplianceEndUser, error) {
	path := "/v2/RegulatoryCompliance/EndUsers/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Attributes != nil {
		v, err := json.Marshal(params.Attributes)

		if err != nil {
			return nil, err
		}

		data.Set("Attributes", string(v))
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &NumbersV2RegulatoryComplianceEndUser{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
