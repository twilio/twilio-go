/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Events
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

	"github.com/twilio/twilio-go/client"
)

// Fetch a specific Event Type.
func (c *ApiService) FetchEventType(Type string) (*EventsV1EventType, error) {
	path := "/v1/Types/{Type}"
	path = strings.Replace(path, "{"+"Type"+"}", Type, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &EventsV1EventType{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListEventType'
type ListEventTypeParams struct {
	// A string parameter filtering the results to return only the Event Types using a given schema.
	SchemaId *string `json:"SchemaId,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListEventTypeParams) SetSchemaId(SchemaId string) *ListEventTypeParams {
	params.SchemaId = &SchemaId
	return params
}
func (params *ListEventTypeParams) SetPageSize(PageSize int) *ListEventTypeParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListEventTypeParams) SetLimit(Limit int) *ListEventTypeParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of EventType records from the API. Request is executed immediately.
func (c *ApiService) PageEventType(params *ListEventTypeParams, pageToken, pageNumber string) (*ListEventTypeResponse, error) {
	path := "/v1/Types"

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.SchemaId != nil {
		data.Set("SchemaId", *params.SchemaId)
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

	ps := &ListEventTypeResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists EventType records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListEventType(params *ListEventTypeParams) ([]EventsV1EventType, error) {
	response, errors := c.StreamEventType(params)

	records := make([]EventsV1EventType, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams EventType records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamEventType(params *ListEventTypeParams) (chan EventsV1EventType, chan error) {
	if params == nil {
		params = &ListEventTypeParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan EventsV1EventType, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageEventType(params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamEventType(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamEventType(response *ListEventTypeResponse, params *ListEventTypeParams, recordChannel chan EventsV1EventType, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Types
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListEventTypeResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListEventTypeResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListEventTypeResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListEventTypeResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
