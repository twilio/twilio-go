/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Voice
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

// Optional parameters for the method 'CreateConnectionPolicy'
type CreateConnectionPolicyParams struct {
	// A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long.
	FriendlyName *string `json:"FriendlyName,omitempty"`
}

func (params *CreateConnectionPolicyParams) SetFriendlyName(FriendlyName string) *CreateConnectionPolicyParams {
	params.FriendlyName = &FriendlyName
	return params
}

func (c *ApiService) CreateConnectionPolicy(params *CreateConnectionPolicyParams) (*VoiceV1ConnectionPolicy, error) {
	return c.CreateConnectionPolicyWithCtx(context.TODO(), params)
}

func (c *ApiService) CreateConnectionPolicyWithCtx(ctx context.Context, params *CreateConnectionPolicyParams) (*VoiceV1ConnectionPolicy, error) {
	path := "/v1/ConnectionPolicies"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}

	resp, err := c.requestHandler.Post(ctx, c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VoiceV1ConnectionPolicy{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

func (c *ApiService) DeleteConnectionPolicy(Sid string) error {
	return c.DeleteConnectionPolicyWithCtx(context.TODO(), Sid)
}

func (c *ApiService) DeleteConnectionPolicyWithCtx(ctx context.Context, Sid string) error {
	path := "/v1/ConnectionPolicies/{Sid}"
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

func (c *ApiService) FetchConnectionPolicy(Sid string) (*VoiceV1ConnectionPolicy, error) {
	return c.FetchConnectionPolicyWithCtx(context.TODO(), Sid)
}

func (c *ApiService) FetchConnectionPolicyWithCtx(ctx context.Context, Sid string) (*VoiceV1ConnectionPolicy, error) {
	path := "/v1/ConnectionPolicies/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(ctx, c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VoiceV1ConnectionPolicy{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListConnectionPolicy'
type ListConnectionPolicyParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListConnectionPolicyParams) SetPageSize(PageSize int) *ListConnectionPolicyParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListConnectionPolicyParams) SetLimit(Limit int) *ListConnectionPolicyParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of ConnectionPolicy records from the API. Request is executed immediately.
func (c *ApiService) PageConnectionPolicy(params *ListConnectionPolicyParams, pageToken, pageNumber string) (*ListConnectionPolicyResponse, error) {
	return c.PageConnectionPolicyWithCtx(context.TODO(), params, pageToken, pageNumber)
}

// Retrieve a single page of ConnectionPolicy records from the API. Request is executed immediately.
func (c *ApiService) PageConnectionPolicyWithCtx(ctx context.Context, params *ListConnectionPolicyParams, pageToken, pageNumber string) (*ListConnectionPolicyResponse, error) {
	path := "/v1/ConnectionPolicies"

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

	ps := &ListConnectionPolicyResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists ConnectionPolicy records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListConnectionPolicy(params *ListConnectionPolicyParams) ([]VoiceV1ConnectionPolicy, error) {
	return c.ListConnectionPolicyWithCtx(context.TODO(), params)
}

// Lists ConnectionPolicy records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListConnectionPolicyWithCtx(ctx context.Context, params *ListConnectionPolicyParams) ([]VoiceV1ConnectionPolicy, error) {
	response, errors := c.StreamConnectionPolicyWithCtx(ctx, params)

	records := make([]VoiceV1ConnectionPolicy, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams ConnectionPolicy records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamConnectionPolicy(params *ListConnectionPolicyParams) (chan VoiceV1ConnectionPolicy, chan error) {
	return c.StreamConnectionPolicyWithCtx(context.TODO(), params)
}

// Streams ConnectionPolicy records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamConnectionPolicyWithCtx(ctx context.Context, params *ListConnectionPolicyParams) (chan VoiceV1ConnectionPolicy, chan error) {
	if params == nil {
		params = &ListConnectionPolicyParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan VoiceV1ConnectionPolicy, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageConnectionPolicyWithCtx(ctx, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamConnectionPolicy(ctx, response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamConnectionPolicy(ctx context.Context, response *ListConnectionPolicyResponse, params *ListConnectionPolicyParams, recordChannel chan VoiceV1ConnectionPolicy, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.ConnectionPolicies
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNextWithCtx(ctx, c.baseURL, response, c.getNextListConnectionPolicyResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListConnectionPolicyResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListConnectionPolicyResponse(ctx context.Context, nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(ctx, nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListConnectionPolicyResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateConnectionPolicy'
type UpdateConnectionPolicyParams struct {
	// A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long.
	FriendlyName *string `json:"FriendlyName,omitempty"`
}

func (params *UpdateConnectionPolicyParams) SetFriendlyName(FriendlyName string) *UpdateConnectionPolicyParams {
	params.FriendlyName = &FriendlyName
	return params
}

func (c *ApiService) UpdateConnectionPolicy(Sid string, params *UpdateConnectionPolicyParams) (*VoiceV1ConnectionPolicy, error) {
	return c.UpdateConnectionPolicyWithCtx(context.TODO(), Sid, params)
}

func (c *ApiService) UpdateConnectionPolicyWithCtx(ctx context.Context, Sid string, params *UpdateConnectionPolicyParams) (*VoiceV1ConnectionPolicy, error) {
	path := "/v1/ConnectionPolicies/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}

	resp, err := c.requestHandler.Post(ctx, c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VoiceV1ConnectionPolicy{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
