/*
 * Twilio - Events
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

// Optional parameters for the method 'CreateSubscribedEvent'
type CreateSubscribedEventParams struct {
	// The schema version that the subscription should use.
	SchemaVersion *int `json:"SchemaVersion,omitempty"`
	// Type of event being subscribed to.
	Type *string `json:"Type,omitempty"`
}

func (params *CreateSubscribedEventParams) SetSchemaVersion(SchemaVersion int) *CreateSubscribedEventParams {
	params.SchemaVersion = &SchemaVersion
	return params
}
func (params *CreateSubscribedEventParams) SetType(Type string) *CreateSubscribedEventParams {
	params.Type = &Type
	return params
}

// Create a new Subscribed Event type for the subscription
func (c *ApiService) CreateSubscribedEvent(SubscriptionSid string, params *CreateSubscribedEventParams) (*EventsV1SubscriptionSubscribedEvent, error) {
	path := "/v1/Subscriptions/{SubscriptionSid}/SubscribedEvents"
	path = strings.Replace(path, "{"+"SubscriptionSid"+"}", SubscriptionSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.SchemaVersion != nil {
		data.Set("SchemaVersion", fmt.Sprint(*params.SchemaVersion))
	}
	if params != nil && params.Type != nil {
		data.Set("Type", *params.Type)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &EventsV1SubscriptionSubscribedEvent{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Remove an event type from a subscription.
func (c *ApiService) DeleteSubscribedEvent(SubscriptionSid string, Type string) error {
	path := "/v1/Subscriptions/{SubscriptionSid}/SubscribedEvents/{Type}"
	path = strings.Replace(path, "{"+"SubscriptionSid"+"}", SubscriptionSid, -1)
	path = strings.Replace(path, "{"+"Type"+"}", Type, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

// Read an Event for a Subscription.
func (c *ApiService) FetchSubscribedEvent(SubscriptionSid string, Type string) (*EventsV1SubscriptionSubscribedEvent, error) {
	path := "/v1/Subscriptions/{SubscriptionSid}/SubscribedEvents/{Type}"
	path = strings.Replace(path, "{"+"SubscriptionSid"+"}", SubscriptionSid, -1)
	path = strings.Replace(path, "{"+"Type"+"}", Type, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &EventsV1SubscriptionSubscribedEvent{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListSubscribedEvent'
type ListSubscribedEventParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListSubscribedEventParams) SetPageSize(PageSize int) *ListSubscribedEventParams {
	params.PageSize = &PageSize
	return params
}

// Retrieve a list of all Subscribed Event types for a Subscription.
func (c *ApiService) ListSubscribedEvent(SubscriptionSid string, params *ListSubscribedEventParams) (*ListSubscribedEventResponse, error) {
	path := "/v1/Subscriptions/{SubscriptionSid}/SubscribedEvents"
	path = strings.Replace(path, "{"+"SubscriptionSid"+"}", SubscriptionSid, -1)

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

	ps := &ListSubscribedEventResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

//Retrieve a single page of SubscribedEvent records from the API. Request is executed immediately.
func (c *ApiService) SubscribedEventPage(SubscriptionSid string, params *ListSubscribedEventParams, pageToken string, pageNumber string, pageSize string) *client.Page {
	path := "/v1/Subscriptions/{SubscriptionSid}/SubscribedEvents"
	path = strings.Replace(path, "{"+"SubscriptionSid"+"}", SubscriptionSid, -1)

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

//Streams SubscribedEvent records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) SubscribedEventStream(SubscriptionSid string, params *ListSubscribedEventParams, meta client.PaginationData) chan map[string]interface{} {
	limits := c.requestHandler.ReadLimits(meta)
	page := c.SubscribedEventPage(SubscriptionSid, params, "", "", fmt.Sprint(limits.PageSize))
	return c.requestHandler.Stream(page, limits.Limit, limits.PageLimit)
}

//Lists SubscribedEvent records from the API as a list. Unlike stream, this operation is eager and will loads 'limit' records into memory before returning.
func (c *ApiService) SubscribedEventList(SubscriptionSid string, params *ListSubscribedEventParams, meta client.PaginationData) []interface{} {
	limits := c.requestHandler.ReadLimits(meta)
	page := c.SubscribedEventPage(SubscriptionSid, params, "", "", fmt.Sprint(limits.PageSize))
	return c.requestHandler.List(page, limits.Limit, limits.PageLimit)
}

// Optional parameters for the method 'UpdateSubscribedEvent'
type UpdateSubscribedEventParams struct {
	// The schema version that the subscription should use.
	SchemaVersion *int `json:"SchemaVersion,omitempty"`
}

func (params *UpdateSubscribedEventParams) SetSchemaVersion(SchemaVersion int) *UpdateSubscribedEventParams {
	params.SchemaVersion = &SchemaVersion
	return params
}

// Update an Event for a Subscription.
func (c *ApiService) UpdateSubscribedEvent(SubscriptionSid string, Type string, params *UpdateSubscribedEventParams) (*EventsV1SubscriptionSubscribedEvent, error) {
	path := "/v1/Subscriptions/{SubscriptionSid}/SubscribedEvents/{Type}"
	path = strings.Replace(path, "{"+"SubscriptionSid"+"}", SubscriptionSid, -1)
	path = strings.Replace(path, "{"+"Type"+"}", Type, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.SchemaVersion != nil {
		data.Set("SchemaVersion", fmt.Sprint(*params.SchemaVersion))
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &EventsV1SubscriptionSubscribedEvent{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
