/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Studio
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/twilio/twilio-go/client"
)

// Delete a specific Flow.
func (c *ApiService) DeleteFlow(Sid string) error {
	return c.DeleteFlowWithCtx(context.TODO(), Sid)
}

// Delete a specific Flow.
func (c *ApiService) DeleteFlowWithCtx(ctx context.Context, Sid string) error {
	path := "/v1/Flows/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Delete(ctx, c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

// Retrieve a specific Flow.
func (c *ApiService) FetchFlow(Sid string) (*StudioV1Flow, error) {
	return c.FetchFlowWithCtx(context.TODO(), Sid)
}

// Retrieve a specific Flow.
func (c *ApiService) FetchFlowWithCtx(ctx context.Context, Sid string) (*StudioV1Flow, error) {
	path := "/v1/Flows/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(ctx, c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &StudioV1Flow{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListFlow'
type ListFlowParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListFlowParams) SetPageSize(PageSize int) *ListFlowParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListFlowParams) SetLimit(Limit int) *ListFlowParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Flow records from the API. Request is executed immediately.
func (c *ApiService) PageFlow(params *ListFlowParams, pageToken, pageNumber string) (*ListFlowResponse, error) {
	return c.PageFlowWithCtx(context.TODO(), params, pageToken, pageNumber)
}

// Retrieve a single page of Flow records from the API. Request is executed immediately.
func (c *ApiService) PageFlowWithCtx(ctx context.Context, params *ListFlowParams, pageToken, pageNumber string) (*ListFlowResponse, error) {
	path := "/v1/Flows"

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

	resp, err := c.requestHandler.Get(ctx, c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListFlowResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Flow records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListFlow(params *ListFlowParams) ([]StudioV1Flow, error) {
	return c.ListFlowWithCtx(context.TODO(), params)
}

// Lists Flow records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListFlowWithCtx(ctx context.Context, params *ListFlowParams) ([]StudioV1Flow, error) {
	response, errors := c.StreamFlowWithCtx(ctx, params)

	records := make([]StudioV1Flow, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams Flow records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamFlow(params *ListFlowParams) (chan StudioV1Flow, chan error) {
	return c.StreamFlowWithCtx(context.TODO(), params)
}

// Streams Flow records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamFlowWithCtx(ctx context.Context, params *ListFlowParams) (chan StudioV1Flow, chan error) {
	if params == nil {
		params = &ListFlowParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan StudioV1Flow, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageFlowWithCtx(ctx, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamFlow(ctx, response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamFlow(ctx context.Context, response *ListFlowResponse, params *ListFlowParams, recordChannel chan StudioV1Flow, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Flows
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNextWithCtx(ctx, c.baseURL, response, c.getNextListFlowResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListFlowResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListFlowResponse(ctx context.Context, nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(ctx, nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListFlowResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
