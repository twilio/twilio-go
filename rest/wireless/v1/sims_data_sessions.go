/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Wireless
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

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
	PageSize *int64 `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int64 `json:"limit,omitempty"`
}

func (params *ListDataSessionParams) SetPageSize(PageSize int64) *ListDataSessionParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListDataSessionParams) SetLimit(Limit int64) *ListDataSessionParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of DataSession records from the API. Request is executed immediately.
func (c *ApiService) PageDataSession(SimSid string, params *ListDataSessionParams, pageToken, pageNumber string) (*ListDataSessionResponse, error) {
	path := "/v1/Sims/{SimSid}/DataSessions"

	path = strings.Replace(path, "{"+"SimSid"+"}", SimSid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

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

	ps := &ListDataSessionResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists DataSession records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListDataSession(SimSid string, params *ListDataSessionParams) ([]WirelessV1DataSession, error) {
	response, errors := c.StreamDataSession(SimSid, params)

	records := make([]WirelessV1DataSession, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams DataSession records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamDataSession(SimSid string, params *ListDataSessionParams) (chan WirelessV1DataSession, chan error) {
	if params == nil {
		params = &ListDataSessionParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan WirelessV1DataSession, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageDataSession(SimSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamDataSession(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamDataSession(response *ListDataSessionResponse, params *ListDataSessionParams, recordChannel chan WirelessV1DataSession, errorChannel chan error) {
	var curRecord int64 = 1

	for response != nil {
		responseRecords := response.DataSessions
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListDataSessionResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListDataSessionResponse)
	}

	close(recordChannel)
	close(errorChannel)
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
