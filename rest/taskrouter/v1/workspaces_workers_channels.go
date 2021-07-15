/*
 * Twilio - Taskrouter
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.19.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"

	"strings"
)

func (c *ApiService) FetchWorkerChannel(WorkspaceSid string, WorkerSid string, Sid string) (*TaskrouterV1WorkspaceWorkerWorkerChannel, error) {
	path := "/v1/Workspaces/{WorkspaceSid}/Workers/{WorkerSid}/Channels/{Sid}"
	path = strings.Replace(path, "{"+"WorkspaceSid"+"}", WorkspaceSid, -1)
	path = strings.Replace(path, "{"+"WorkerSid"+"}", WorkerSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TaskrouterV1WorkspaceWorkerWorkerChannel{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListWorkerChannel'
type ListWorkerChannelParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListWorkerChannelParams) SetPageSize(PageSize int) *ListWorkerChannelParams {
	params.PageSize = &PageSize
	return params
}

func (c *ApiService) ListWorkerChannel(WorkspaceSid string, WorkerSid string, params *ListWorkerChannelParams) (*ListWorkerChannelResponse, error) {
	path := "/v1/Workspaces/{WorkspaceSid}/Workers/{WorkerSid}/Channels"
	path = strings.Replace(path, "{"+"WorkspaceSid"+"}", WorkspaceSid, -1)
	path = strings.Replace(path, "{"+"WorkerSid"+"}", WorkerSid, -1)

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

	ps := &ListWorkerChannelResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'UpdateWorkerChannel'
type UpdateWorkerChannelParams struct {
	// Whether the WorkerChannel is available. Set to `false` to prevent the Worker from receiving any new Tasks of this TaskChannel type.
	Available *bool `json:"Available,omitempty"`
	// The total number of Tasks that the Worker should handle for the TaskChannel type. TaskRouter creates reservations for Tasks of this TaskChannel type up to the specified capacity. If the capacity is 0, no new reservations will be created.
	Capacity *int `json:"Capacity,omitempty"`
}

func (params *UpdateWorkerChannelParams) SetAvailable(Available bool) *UpdateWorkerChannelParams {
	params.Available = &Available
	return params
}
func (params *UpdateWorkerChannelParams) SetCapacity(Capacity int) *UpdateWorkerChannelParams {
	params.Capacity = &Capacity
	return params
}

func (c *ApiService) UpdateWorkerChannel(WorkspaceSid string, WorkerSid string, Sid string, params *UpdateWorkerChannelParams) (*TaskrouterV1WorkspaceWorkerWorkerChannel, error) {
	path := "/v1/Workspaces/{WorkspaceSid}/Workers/{WorkerSid}/Channels/{Sid}"
	path = strings.Replace(path, "{"+"WorkspaceSid"+"}", WorkspaceSid, -1)
	path = strings.Replace(path, "{"+"WorkerSid"+"}", WorkerSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Available != nil {
		data.Set("Available", fmt.Sprint(*params.Available))
	}
	if params != nil && params.Capacity != nil {
		data.Set("Capacity", fmt.Sprint(*params.Capacity))
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TaskrouterV1WorkspaceWorkerWorkerChannel{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
