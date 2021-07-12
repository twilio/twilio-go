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

// Optional parameters for the method 'ListBillingPeriod'
type ListBillingPeriodParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListBillingPeriodParams) SetPageSize(PageSize int) *ListBillingPeriodParams {
	params.PageSize = &PageSize
	return params
}

//Retrieve a single page of BillingPeriod records from the API. Request is executed immediately.
func (c *ApiService) PageBillingPeriod(SimSid string, params *ListBillingPeriodParams, pageToken string, pageNumber string) (*ListBillingPeriodResponse, error) {
	path := "/v1/Sims/{SimSid}/BillingPeriods"

	path = strings.Replace(path, "{"+"SimSid"+"}", SimSid, -1)

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

	ps := &ListBillingPeriodResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

//Lists BillingPeriod records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListBillingPeriod(SimSid string, params *ListBillingPeriodParams, limit *int) ([]*ListBillingPeriodResponse, error) {
	params.SetPageSize(client.ReadLimits(params.PageSize, limit))

	response, err := c.PageBillingPeriod(SimSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []*ListBillingPeriodResponse

	for response != nil {
		records = append(records, response)

		var record interface{}
		if record, err = client.GetNext(response, &curRecord, limit, c.getNextListBillingPeriodResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListBillingPeriodResponse)
	}

	return records, err
}

//Streams BillingPeriod records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamBillingPeriod(SimSid string, params *ListBillingPeriodParams, limit *int) (chan *ListBillingPeriodResponse, error) {
	params.SetPageSize(client.ReadLimits(params.PageSize, limit))

	response, err := c.PageBillingPeriod(SimSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan *ListBillingPeriodResponse, 1)

	go func() {
		for response != nil {
			channel <- response

			var record interface{}
			if record, err = client.GetNext(response, &curRecord, limit, c.getNextListBillingPeriodResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListBillingPeriodResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListBillingPeriodResponse(nextPageUri string) (interface{}, error) {
	if nextPageUri == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(c.baseURL+nextPageUri, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListBillingPeriodResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
