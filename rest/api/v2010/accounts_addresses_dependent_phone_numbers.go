/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.19.1
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

// Optional parameters for the method 'ListDependentPhoneNumber'
type ListDependentPhoneNumberParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the DependentPhoneNumber resources to read.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListDependentPhoneNumberParams) SetPathAccountSid(PathAccountSid string) *ListDependentPhoneNumberParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *ListDependentPhoneNumberParams) SetPageSize(PageSize int) *ListDependentPhoneNumberParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListDependentPhoneNumberParams) SetLimit(Limit int) *ListDependentPhoneNumberParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of DependentPhoneNumber records from the API. Request is executed immediately.
func (c *ApiService) PageDependentPhoneNumber(AddressSid string, params *ListDependentPhoneNumberParams, pageToken string, pageNumber string) (*ListDependentPhoneNumberResponse, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Addresses/{AddressSid}/DependentPhoneNumbers.json"

	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"AddressSid"+"}", AddressSid, -1)

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

	ps := &ListDependentPhoneNumberResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists DependentPhoneNumber records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListDependentPhoneNumber(AddressSid string, params *ListDependentPhoneNumberParams) ([]ApiV2010DependentPhoneNumber, error) {
	if params == nil {
		params = &ListDependentPhoneNumberParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageDependentPhoneNumber(AddressSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []ApiV2010DependentPhoneNumber

	for response != nil {
		records = append(records, response.DependentPhoneNumbers...)

		var record interface{}
		if record, err = client.GetNext(response, &curRecord, params.Limit, c.getNextListDependentPhoneNumberResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListDependentPhoneNumberResponse)
	}

	return records, err
}

// Streams DependentPhoneNumber records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamDependentPhoneNumber(AddressSid string, params *ListDependentPhoneNumberParams) (chan ApiV2010DependentPhoneNumber, error) {
	if params == nil {
		params = &ListDependentPhoneNumberParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageDependentPhoneNumber(AddressSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan ApiV2010DependentPhoneNumber, 1)

	go func() {
		for response != nil {
			for item := range response.DependentPhoneNumbers {
				channel <- response.DependentPhoneNumbers[item]
			}

			var record interface{}
			if record, err = client.GetNext(response, &curRecord, params.Limit, c.getNextListDependentPhoneNumberResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListDependentPhoneNumberResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListDependentPhoneNumberResponse(nextPageUri string) (interface{}, error) {
	if nextPageUri == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(c.baseURL+nextPageUri, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListDependentPhoneNumberResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
