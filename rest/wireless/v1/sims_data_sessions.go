/*
 * Twilio - Wireless
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

// Optional parameters for the method 'ListDataSession'
type ListDataSessionParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListDataSessionParams) SetPageSize(PageSize int) *ListDataSessionParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListDataSessionParams) SetLimit(Limit int) *ListDataSessionParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of DataSession records from the API. Request is executed immediately.
func (c *ApiService) PageDataSession(SimSid string, params *ListDataSessionParams, pageToken string, pageNumber string) (*ListDataSessionResponse, error) {
	path := "/v1/Sims/{SimSid}/DataSessions"

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

	ps := &ListDataSessionResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists DataSession records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListDataSession(SimSid string, params *ListDataSessionParams) ([]WirelessV1SimDataSession, error) {
	if params == nil {
		params = &ListDataSessionParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageDataSession(SimSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []WirelessV1SimDataSession

	for response != nil {
		records = append(records, response.DataSessions...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListDataSessionResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListDataSessionResponse)
	}

	return records, err
}

// Streams DataSession records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamDataSession(SimSid string, params *ListDataSessionParams) (chan WirelessV1SimDataSession, error) {
	if params == nil {
		params = &ListDataSessionParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageDataSession(SimSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan WirelessV1SimDataSession, 1)

	go func() {
		for response != nil {
			for item := range response.DataSessions {
				channel <- response.DataSessions[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListDataSessionResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListDataSessionResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListDataSessionResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListDataSessionResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
