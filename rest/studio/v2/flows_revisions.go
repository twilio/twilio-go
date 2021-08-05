/*
 * Twilio - Studio
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

// Retrieve a specific Flow revision.
func (c *ApiService) FetchFlowRevision(Sid string, Revision string) (*StudioV2FlowRevision, error) {
	path := "/v2/Flows/{Sid}/Revisions/{Revision}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)
	path = strings.Replace(path, "{"+"Revision"+"}", Revision, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &StudioV2FlowRevision{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListFlowRevision'
type ListFlowRevisionParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListFlowRevisionParams) SetPageSize(PageSize int) *ListFlowRevisionParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListFlowRevisionParams) SetLimit(Limit int) *ListFlowRevisionParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of FlowRevision records from the API. Request is executed immediately.
func (c *ApiService) PageFlowRevision(Sid string, params *ListFlowRevisionParams, pageToken string, pageNumber string) (*ListFlowRevisionResponse, error) {
	path := "/v2/Flows/{Sid}/Revisions"

	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

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

	ps := &ListFlowRevisionResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists FlowRevision records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListFlowRevision(Sid string, params *ListFlowRevisionParams) ([]StudioV2FlowRevision, error) {
	if params == nil {
		params = &ListFlowRevisionParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageFlowRevision(Sid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []StudioV2FlowRevision

	for response != nil {
		records = append(records, response.Revisions...)

		var record interface{}
		if record, err = client.GetNext(response, &curRecord, params.Limit, c.getNextListFlowRevisionResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListFlowRevisionResponse)
	}

	return records, err
}

// Streams FlowRevision records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamFlowRevision(Sid string, params *ListFlowRevisionParams) (chan StudioV2FlowRevision, error) {
	if params == nil {
		params = &ListFlowRevisionParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageFlowRevision(Sid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan StudioV2FlowRevision, 1)

	go func() {
		for response != nil {
			for item := range response.Revisions {
				channel <- response.Revisions[item]
			}

			var record interface{}
			if record, err = client.GetNext(response, &curRecord, params.Limit, c.getNextListFlowRevisionResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListFlowRevisionResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListFlowRevisionResponse(nextPageUri string) (interface{}, error) {
	if nextPageUri == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(c.baseURL+nextPageUri, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListFlowRevisionResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
