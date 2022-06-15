/*
 * Twilio - Api
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
	"strings"

	"github.com/twilio/twilio-go/client"
)

// Optional parameters for the method 'ListCallEvent'
type ListCallEventParams struct {
	// The unique SID identifier of the Account.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListCallEventParams) SetPathAccountSid(PathAccountSid string) *ListCallEventParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *ListCallEventParams) SetPageSize(PageSize int) *ListCallEventParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListCallEventParams) SetLimit(Limit int) *ListCallEventParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of CallEvent records from the API. Request is executed immediately.
func (c *ApiService) PageCallEvent(CallSid string, params *ListCallEventParams, pageToken, pageNumber string) (*ListCallEventResponse, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Calls/{CallSid}/Events.json"

	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"CallSid"+"}", CallSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	if pageToken != "" {
		data.Set("PageToken", pageToken)
	}
	if pageNumber != "" {
		data.Set("Page", pageNumber)
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListCallEventResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists CallEvent records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListCallEvent(CallSid string, params *ListCallEventParams) ([]ApiV2010CallEvent, error) {
	response, errors := c.StreamCallEvent(CallSid, params)

	records := make([]ApiV2010CallEvent, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams CallEvent records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamCallEvent(CallSid string, params *ListCallEventParams) (chan ApiV2010CallEvent, chan error) {
	if params == nil {
		params = &ListCallEventParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan ApiV2010CallEvent, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageCallEvent(CallSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamCallEvent(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamCallEvent(response *ListCallEventResponse, params *ListCallEventParams, recordChannel chan ApiV2010CallEvent, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Events
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListCallEventResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListCallEventResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListCallEventResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListCallEventResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
