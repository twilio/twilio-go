/*
 * Twilio - Sync
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

// Optional parameters for the method 'CreateSyncStream'
type CreateSyncStreamParams struct {
	// How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the Stream expires and is deleted (time-to-live).
	Ttl *int `json:"Ttl,omitempty"`
	// An application-defined string that uniquely identifies the resource. This value must be unique within its Service and it can be up to 320 characters long. The `unique_name` value can be used as an alternative to the `sid` in the URL path to address the resource.
	UniqueName *string `json:"UniqueName,omitempty"`
}

func (params *CreateSyncStreamParams) SetTtl(Ttl int) *CreateSyncStreamParams {
	params.Ttl = &Ttl
	return params
}
func (params *CreateSyncStreamParams) SetUniqueName(UniqueName string) *CreateSyncStreamParams {
	params.UniqueName = &UniqueName
	return params
}

// Create a new Stream.
func (c *ApiService) CreateSyncStream(ServiceSid string, params *CreateSyncStreamParams) (*SyncV1ServiceSyncStream, error) {
	path := "/v1/Services/{ServiceSid}/Streams"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Ttl != nil {
		data.Set("Ttl", fmt.Sprint(*params.Ttl))
	}
	if params != nil && params.UniqueName != nil {
		data.Set("UniqueName", *params.UniqueName)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &SyncV1ServiceSyncStream{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Delete a specific Stream.
func (c *ApiService) DeleteSyncStream(ServiceSid string, Sid string) error {
	path := "/v1/Services/{ServiceSid}/Streams/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
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

// Fetch a specific Stream.
func (c *ApiService) FetchSyncStream(ServiceSid string, Sid string) (*SyncV1ServiceSyncStream, error) {
	path := "/v1/Services/{ServiceSid}/Streams/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &SyncV1ServiceSyncStream{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListSyncStream'
type ListSyncStreamParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListSyncStreamParams) SetPageSize(PageSize int) *ListSyncStreamParams {
	params.PageSize = &PageSize
	return params
}

// Retrieve a list of all Streams in a Service Instance.
func (c *ApiService) ListSyncStream(ServiceSid string, params *ListSyncStreamParams) (*ListSyncStreamResponse, error) {
	path := "/v1/Services/{ServiceSid}/Streams"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListSyncStreamResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

//Retrieve a single page of  records from the API. Request is executed immediately.
func (c *ApiService) ServicesStreamsPage(ServiceSid string, params *ListSyncStreamParams, pageToken string, pageNumber string) (*ListSyncStreamResponse, error) {
	path := "/v1/Services/{ServiceSid}/Streams"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	data.Set("PageToken", pageToken)
	data.Set("PageNumber", pageNumber)

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListSyncStreamResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

//Lists ServicesStreams records from the API as a list. Unlike stream, this operation is eager and will loads 'limit' records into memory before returning.
func (c *ApiService) ServicesStreamsList(ServiceSid string, params *ListSyncStreamParams, limit int) ([]ListSyncStreamResponse, error) {
	params.SetPageSize(c.requestHandler.ReadLimits(params.PageSize, limit))
	response, err := c.ListSyncStream(ServiceSid, params)
	if err != nil {
		return nil, err
	}

	page := client.NewPage(c.baseURL, response)

	resp := c.requestHandler.List(page, limit, 0)
	ret := make([]ListSyncStreamResponse, len(resp))

	for i := range resp {
		jsonStr, _ := json.Marshal(resp[i])
		ps := ListSyncStreamResponse{}
		if err := json.Unmarshal(jsonStr, &ps); err != nil {
			return ret, err
		}

		ret[i] = ps
	}

	return ret, nil
}

//Streams ServicesStreams records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) ServicesStreamsStream(ServiceSid string, params *ListSyncStreamParams, limit int) (chan interface{}, error) {
	params.SetPageSize(c.requestHandler.ReadLimits(params.PageSize, limit))
	response, err := c.ListSyncStream(ServiceSid, params)
	if err != nil {
		return nil, err
	}

	page := client.NewPage(c.baseURL, response)

	ps := ListSyncStreamResponse{}
	return c.requestHandler.Stream(page, limit, 0, ps), nil
}

// Optional parameters for the method 'UpdateSyncStream'
type UpdateSyncStreamParams struct {
	// How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the Stream expires and is deleted (time-to-live).
	Ttl *int `json:"Ttl,omitempty"`
}

func (params *UpdateSyncStreamParams) SetTtl(Ttl int) *UpdateSyncStreamParams {
	params.Ttl = &Ttl
	return params
}

// Update a specific Stream.
func (c *ApiService) UpdateSyncStream(ServiceSid string, Sid string, params *UpdateSyncStreamParams) (*SyncV1ServiceSyncStream, error) {
	path := "/v1/Services/{ServiceSid}/Streams/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Ttl != nil {
		data.Set("Ttl", fmt.Sprint(*params.Ttl))
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &SyncV1ServiceSyncStream{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
