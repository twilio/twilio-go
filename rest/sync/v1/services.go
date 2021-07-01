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

// Optional parameters for the method 'CreateService'
type CreateServiceParams struct {
	// Whether token identities in the Service must be granted access to Sync objects by using the [Permissions](https://www.twilio.com/docs/sync/api/sync-permissions) resource.
	AclEnabled *bool `json:"AclEnabled,omitempty"`
	// A string that you assign to describe the resource.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// Whether every `endpoint_disconnected` event should occur after a configurable delay. The default is `false`, where the `endpoint_disconnected` event occurs immediately after disconnection. When `true`, intervening reconnections can prevent the `endpoint_disconnected` event.
	ReachabilityDebouncingEnabled *bool `json:"ReachabilityDebouncingEnabled,omitempty"`
	// The reachability event delay in milliseconds if `reachability_debouncing_enabled` = `true`.  Must be between 1,000 and 30,000 and defaults to 5,000. This is the number of milliseconds after the last running client disconnects, and a Sync identity is declared offline, before the `webhook_url` is called if all endpoints remain offline. A reconnection from the same identity by any endpoint during this interval prevents the call to `webhook_url`.
	ReachabilityDebouncingWindow *int `json:"ReachabilityDebouncingWindow,omitempty"`
	// Whether the service instance should call `webhook_url` when client endpoints connect to Sync. The default is `false`.
	ReachabilityWebhooksEnabled *bool `json:"ReachabilityWebhooksEnabled,omitempty"`
	// The URL we should call when Sync objects are manipulated.
	WebhookUrl *string `json:"WebhookUrl,omitempty"`
	// Whether the Service instance should call `webhook_url` when the REST API is used to update Sync objects. The default is `false`.
	WebhooksFromRestEnabled *bool `json:"WebhooksFromRestEnabled,omitempty"`
}

func (params *CreateServiceParams) SetAclEnabled(AclEnabled bool) *CreateServiceParams {
	params.AclEnabled = &AclEnabled
	return params
}
func (params *CreateServiceParams) SetFriendlyName(FriendlyName string) *CreateServiceParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *CreateServiceParams) SetReachabilityDebouncingEnabled(ReachabilityDebouncingEnabled bool) *CreateServiceParams {
	params.ReachabilityDebouncingEnabled = &ReachabilityDebouncingEnabled
	return params
}
func (params *CreateServiceParams) SetReachabilityDebouncingWindow(ReachabilityDebouncingWindow int) *CreateServiceParams {
	params.ReachabilityDebouncingWindow = &ReachabilityDebouncingWindow
	return params
}
func (params *CreateServiceParams) SetReachabilityWebhooksEnabled(ReachabilityWebhooksEnabled bool) *CreateServiceParams {
	params.ReachabilityWebhooksEnabled = &ReachabilityWebhooksEnabled
	return params
}
func (params *CreateServiceParams) SetWebhookUrl(WebhookUrl string) *CreateServiceParams {
	params.WebhookUrl = &WebhookUrl
	return params
}
func (params *CreateServiceParams) SetWebhooksFromRestEnabled(WebhooksFromRestEnabled bool) *CreateServiceParams {
	params.WebhooksFromRestEnabled = &WebhooksFromRestEnabled
	return params
}

func (c *ApiService) CreateService(params *CreateServiceParams) (*SyncV1Service, error) {
	path := "/v1/Services"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.AclEnabled != nil {
		data.Set("AclEnabled", fmt.Sprint(*params.AclEnabled))
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.ReachabilityDebouncingEnabled != nil {
		data.Set("ReachabilityDebouncingEnabled", fmt.Sprint(*params.ReachabilityDebouncingEnabled))
	}
	if params != nil && params.ReachabilityDebouncingWindow != nil {
		data.Set("ReachabilityDebouncingWindow", fmt.Sprint(*params.ReachabilityDebouncingWindow))
	}
	if params != nil && params.ReachabilityWebhooksEnabled != nil {
		data.Set("ReachabilityWebhooksEnabled", fmt.Sprint(*params.ReachabilityWebhooksEnabled))
	}
	if params != nil && params.WebhookUrl != nil {
		data.Set("WebhookUrl", *params.WebhookUrl)
	}
	if params != nil && params.WebhooksFromRestEnabled != nil {
		data.Set("WebhooksFromRestEnabled", fmt.Sprint(*params.WebhooksFromRestEnabled))
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &SyncV1Service{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

func (c *ApiService) DeleteService(Sid string) error {
	path := "/v1/Services/{Sid}"
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

func (c *ApiService) FetchService(Sid string) (*SyncV1Service, error) {
	path := "/v1/Services/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &SyncV1Service{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListService'
type ListServiceParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListServiceParams) SetPageSize(PageSize int) *ListServiceParams {
	params.PageSize = &PageSize
	return params
}

func (c *ApiService) ListService(params *ListServiceParams) (*ListServiceResponse, error) {
	path := "/v1/Services"

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

	ps := &ListServiceResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

//Retrieve a single page of Service records from the API. Request is executed immediately.
func (c *ApiService) ServicePage(params *ListServiceParams, pageToken string, pageNumber string, pageSize string) *client.Page {
	path := "/v1/Services"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	data.Set("PageToken", pageToken)
	data.Set("PageNumber", pageNumber)
	data.Set("PageSize", pageSize)

	response, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil
	}

	return client.NewPage(c.baseURL, response)
}

//Streams Service records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) ServiceStream(params *ListServiceParams, meta client.PaginationData) chan map[string]interface{} {
	limits := c.requestHandler.ReadLimits(meta)
	page := c.ServicePage(params, "", "", fmt.Sprint(limits.PageSize))
	return c.requestHandler.Stream(page, limits.Limit, limits.PageLimit)
}

//Lists Service records from the API as a list. Unlike stream, this operation is eager and will loads 'limit' records into memory before returning.
func (c *ApiService) ServiceList(params *ListServiceParams, meta client.PaginationData) []interface{} {
	limits := c.requestHandler.ReadLimits(meta)
	page := c.ServicePage(params, "", "", fmt.Sprint(limits.PageSize))
	return c.requestHandler.List(page, limits.Limit, limits.PageLimit)
}

// Optional parameters for the method 'UpdateService'
type UpdateServiceParams struct {
	// Whether token identities in the Service must be granted access to Sync objects by using the [Permissions](https://www.twilio.com/docs/sync/api/sync-permissions) resource.
	AclEnabled *bool `json:"AclEnabled,omitempty"`
	// A string that you assign to describe the resource.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// Whether every `endpoint_disconnected` event should occur after a configurable delay. The default is `false`, where the `endpoint_disconnected` event occurs immediately after disconnection. When `true`, intervening reconnections can prevent the `endpoint_disconnected` event.
	ReachabilityDebouncingEnabled *bool `json:"ReachabilityDebouncingEnabled,omitempty"`
	// The reachability event delay in milliseconds if `reachability_debouncing_enabled` = `true`.  Must be between 1,000 and 30,000 and defaults to 5,000. This is the number of milliseconds after the last running client disconnects, and a Sync identity is declared offline, before the webhook is called if all endpoints remain offline. A reconnection from the same identity by any endpoint during this interval prevents the webhook from being called.
	ReachabilityDebouncingWindow *int `json:"ReachabilityDebouncingWindow,omitempty"`
	// Whether the service instance should call `webhook_url` when client endpoints connect to Sync. The default is `false`.
	ReachabilityWebhooksEnabled *bool `json:"ReachabilityWebhooksEnabled,omitempty"`
	// The URL we should call when Sync objects are manipulated.
	WebhookUrl *string `json:"WebhookUrl,omitempty"`
	// Whether the Service instance should call `webhook_url` when the REST API is used to update Sync objects. The default is `false`.
	WebhooksFromRestEnabled *bool `json:"WebhooksFromRestEnabled,omitempty"`
}

func (params *UpdateServiceParams) SetAclEnabled(AclEnabled bool) *UpdateServiceParams {
	params.AclEnabled = &AclEnabled
	return params
}
func (params *UpdateServiceParams) SetFriendlyName(FriendlyName string) *UpdateServiceParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *UpdateServiceParams) SetReachabilityDebouncingEnabled(ReachabilityDebouncingEnabled bool) *UpdateServiceParams {
	params.ReachabilityDebouncingEnabled = &ReachabilityDebouncingEnabled
	return params
}
func (params *UpdateServiceParams) SetReachabilityDebouncingWindow(ReachabilityDebouncingWindow int) *UpdateServiceParams {
	params.ReachabilityDebouncingWindow = &ReachabilityDebouncingWindow
	return params
}
func (params *UpdateServiceParams) SetReachabilityWebhooksEnabled(ReachabilityWebhooksEnabled bool) *UpdateServiceParams {
	params.ReachabilityWebhooksEnabled = &ReachabilityWebhooksEnabled
	return params
}
func (params *UpdateServiceParams) SetWebhookUrl(WebhookUrl string) *UpdateServiceParams {
	params.WebhookUrl = &WebhookUrl
	return params
}
func (params *UpdateServiceParams) SetWebhooksFromRestEnabled(WebhooksFromRestEnabled bool) *UpdateServiceParams {
	params.WebhooksFromRestEnabled = &WebhooksFromRestEnabled
	return params
}

func (c *ApiService) UpdateService(Sid string, params *UpdateServiceParams) (*SyncV1Service, error) {
	path := "/v1/Services/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.AclEnabled != nil {
		data.Set("AclEnabled", fmt.Sprint(*params.AclEnabled))
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.ReachabilityDebouncingEnabled != nil {
		data.Set("ReachabilityDebouncingEnabled", fmt.Sprint(*params.ReachabilityDebouncingEnabled))
	}
	if params != nil && params.ReachabilityDebouncingWindow != nil {
		data.Set("ReachabilityDebouncingWindow", fmt.Sprint(*params.ReachabilityDebouncingWindow))
	}
	if params != nil && params.ReachabilityWebhooksEnabled != nil {
		data.Set("ReachabilityWebhooksEnabled", fmt.Sprint(*params.ReachabilityWebhooksEnabled))
	}
	if params != nil && params.WebhookUrl != nil {
		data.Set("WebhookUrl", *params.WebhookUrl)
	}
	if params != nil && params.WebhooksFromRestEnabled != nil {
		data.Set("WebhooksFromRestEnabled", fmt.Sprint(*params.WebhooksFromRestEnabled))
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &SyncV1Service{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
