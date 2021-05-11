/*
 * Twilio - Events
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.0.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	twilio "github.com/twilio/twilio-go/client"
)

type DefaultApiService struct {
	baseURL string
	client  twilio.BaseClient
}

func NewDefaultApiService(client twilio.BaseClient) *DefaultApiService {
	return &DefaultApiService{
		client:  client,
		baseURL: "https://events.twilio.com",
	}
}

// Optional parameters for the method 'CreateSink'
type CreateSinkParams struct {
	// A human readable description for the Sink **This value should not contain PII.**
	Description *string `json:"Description,omitempty"`
	// The information required for Twilio to connect to the provided Sink encoded as JSON.
	SinkConfiguration *map[string]interface{} `json:"SinkConfiguration,omitempty"`
	// The Sink type. Can only be \\\"kinesis\\\" or \\\"webhook\\\" currently.
	SinkType *string `json:"SinkType,omitempty"`
}

func (params *CreateSinkParams) SetDescription(Description string) *CreateSinkParams {
	params.Description = &Description
	return params
}
func (params *CreateSinkParams) SetSinkConfiguration(SinkConfiguration map[string]interface{}) *CreateSinkParams {
	params.SinkConfiguration = &SinkConfiguration
	return params
}
func (params *CreateSinkParams) SetSinkType(SinkType string) *CreateSinkParams {
	params.SinkType = &SinkType
	return params
}

// Create a new Sink
func (c *DefaultApiService) CreateSink(params *CreateSinkParams) (*EventsV1Sink, error) {
	path := "/v1/Sinks"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Description != nil {
		data.Set("Description", *params.Description)
	}
	if params != nil && params.SinkConfiguration != nil {
		v, err := json.Marshal(params.SinkConfiguration)

		if err != nil {
			return nil, err
		}

		data.Set("SinkConfiguration", string(v))
	}
	if params != nil && params.SinkType != nil {
		data.Set("SinkType", *params.SinkType)
	}

	resp, err := c.client.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &EventsV1Sink{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Create a new Sink Test Event for the given Sink.
func (c *DefaultApiService) CreateSinkTest(Sid string) (*EventsV1SinkSinkTest, error) {
	path := "/v1/Sinks/{Sid}/Test"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.client.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &EventsV1SinkSinkTest{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'CreateSinkValidate'
type CreateSinkValidateParams struct {
	// A 34 character string that uniquely identifies the test event for a Sink being validated.
	TestId *string `json:"TestId,omitempty"`
}

func (params *CreateSinkValidateParams) SetTestId(TestId string) *CreateSinkValidateParams {
	params.TestId = &TestId
	return params
}

// Validate that a test event for a Sink was received.
func (c *DefaultApiService) CreateSinkValidate(Sid string, params *CreateSinkValidateParams) (*EventsV1SinkSinkValidate, error) {
	path := "/v1/Sinks/{Sid}/Validate"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.TestId != nil {
		data.Set("TestId", *params.TestId)
	}

	resp, err := c.client.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &EventsV1SinkSinkValidate{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'CreateSubscribedEvent'
type CreateSubscribedEventParams struct {
	// The schema version that the subscription should use.
	SchemaVersion *int32 `json:"SchemaVersion,omitempty"`
	// Type of event being subscribed to.
	Type *string `json:"Type,omitempty"`
}

func (params *CreateSubscribedEventParams) SetSchemaVersion(SchemaVersion int32) *CreateSubscribedEventParams {
	params.SchemaVersion = &SchemaVersion
	return params
}
func (params *CreateSubscribedEventParams) SetType(Type string) *CreateSubscribedEventParams {
	params.Type = &Type
	return params
}

// Create a new Subscribed Event type for the subscription
func (c *DefaultApiService) CreateSubscribedEvent(SubscriptionSid string, params *CreateSubscribedEventParams) (*EventsV1SubscriptionSubscribedEvent, error) {
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

	resp, err := c.client.Post(c.baseURL+path, data, headers)
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

// Optional parameters for the method 'CreateSubscription'
type CreateSubscriptionParams struct {
	// A human readable description for the Subscription **This value should not contain PII.**
	Description *string `json:"Description,omitempty"`
	// The SID of the sink that events selected by this subscription should be sent to. Sink must be active for the subscription to be created.
	SinkSid *string `json:"SinkSid,omitempty"`
	// An array of objects containing the subscribed Event Types
	Types *[]map[string]interface{} `json:"Types,omitempty"`
}

func (params *CreateSubscriptionParams) SetDescription(Description string) *CreateSubscriptionParams {
	params.Description = &Description
	return params
}
func (params *CreateSubscriptionParams) SetSinkSid(SinkSid string) *CreateSubscriptionParams {
	params.SinkSid = &SinkSid
	return params
}
func (params *CreateSubscriptionParams) SetTypes(Types []map[string]interface{}) *CreateSubscriptionParams {
	params.Types = &Types
	return params
}

// Create a new Subscription.
func (c *DefaultApiService) CreateSubscription(params *CreateSubscriptionParams) (*EventsV1Subscription, error) {
	path := "/v1/Subscriptions"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Description != nil {
		data.Set("Description", *params.Description)
	}
	if params != nil && params.SinkSid != nil {
		data.Set("SinkSid", *params.SinkSid)
	}
	if params != nil && params.Types != nil {
		v, err := json.Marshal(params.Types)

		if err != nil {
			return nil, err
		}

		data.Set("Types", string(v))
	}

	resp, err := c.client.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &EventsV1Subscription{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Delete a specific Sink.
func (c *DefaultApiService) DeleteSink(Sid string) error {
	path := "/v1/Sinks/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.client.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

// Remove an event type from a subscription.
func (c *DefaultApiService) DeleteSubscribedEvent(SubscriptionSid string, Type string) error {
	path := "/v1/Subscriptions/{SubscriptionSid}/SubscribedEvents/{Type}"
	path = strings.Replace(path, "{"+"SubscriptionSid"+"}", SubscriptionSid, -1)
	path = strings.Replace(path, "{"+"Type"+"}", Type, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.client.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

// Delete a specific Subscription.
func (c *DefaultApiService) DeleteSubscription(Sid string) error {
	path := "/v1/Subscriptions/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.client.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

// Fetch a specific Event Type.
func (c *DefaultApiService) FetchEventType(Type string) (*EventsV1EventType, error) {
	path := "/v1/Types/{Type}"
	path = strings.Replace(path, "{"+"Type"+"}", Type, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.client.Get(c.baseURL+path, data, headers)
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

// Fetch a specific schema with its nested versions.
func (c *DefaultApiService) FetchSchema(Id string) (*EventsV1Schema, error) {
	path := "/v1/Schemas/{Id}"
	path = strings.Replace(path, "{"+"Id"+"}", Id, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &EventsV1Schema{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Fetch a specific schema and version.
func (c *DefaultApiService) FetchSchemaVersion(Id string, SchemaVersion int32) (*EventsV1SchemaSchemaVersion, error) {
	path := "/v1/Schemas/{Id}/Versions/{SchemaVersion}"
	path = strings.Replace(path, "{"+"Id"+"}", Id, -1)
	path = strings.Replace(path, "{"+"SchemaVersion"+"}", fmt.Sprint(SchemaVersion), -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &EventsV1SchemaSchemaVersion{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Fetch a specific Sink.
func (c *DefaultApiService) FetchSink(Sid string) (*EventsV1Sink, error) {
	path := "/v1/Sinks/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &EventsV1Sink{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Read an Event for a Subscription.
func (c *DefaultApiService) FetchSubscribedEvent(SubscriptionSid string, Type string) (*EventsV1SubscriptionSubscribedEvent, error) {
	path := "/v1/Subscriptions/{SubscriptionSid}/SubscribedEvents/{Type}"
	path = strings.Replace(path, "{"+"SubscriptionSid"+"}", SubscriptionSid, -1)
	path = strings.Replace(path, "{"+"Type"+"}", Type, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.client.Get(c.baseURL+path, data, headers)
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

// Fetch a specific Subscription.
func (c *DefaultApiService) FetchSubscription(Sid string) (*EventsV1Subscription, error) {
	path := "/v1/Subscriptions/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &EventsV1Subscription{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListEventType'
type ListEventTypeParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int32 `json:"PageSize,omitempty"`
}

func (params *ListEventTypeParams) SetPageSize(PageSize int32) *ListEventTypeParams {
	params.PageSize = &PageSize
	return params
}

// Retrieve a paginated list of all the available Event Types.
func (c *DefaultApiService) ListEventType(params *ListEventTypeParams) (*ListEventTypeResponse, error) {
	path := "/v1/Types"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.client.Get(c.baseURL+path, data, headers)
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

// Optional parameters for the method 'ListSchemaVersion'
type ListSchemaVersionParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int32 `json:"PageSize,omitempty"`
}

func (params *ListSchemaVersionParams) SetPageSize(PageSize int32) *ListSchemaVersionParams {
	params.PageSize = &PageSize
	return params
}

// Retrieve a paginated list of versions of the schema.
func (c *DefaultApiService) ListSchemaVersion(Id string, params *ListSchemaVersionParams) (*ListSchemaVersionResponse, error) {
	path := "/v1/Schemas/{Id}/Versions"
	path = strings.Replace(path, "{"+"Id"+"}", Id, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListSchemaVersionResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListSink'
type ListSinkParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int32 `json:"PageSize,omitempty"`
}

func (params *ListSinkParams) SetPageSize(PageSize int32) *ListSinkParams {
	params.PageSize = &PageSize
	return params
}

// Retrieve a paginated list of Sinks belonging to the account used to make the request.
func (c *DefaultApiService) ListSink(params *ListSinkParams) (*ListSinkResponse, error) {
	path := "/v1/Sinks"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListSinkResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListSubscribedEvent'
type ListSubscribedEventParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int32 `json:"PageSize,omitempty"`
}

func (params *ListSubscribedEventParams) SetPageSize(PageSize int32) *ListSubscribedEventParams {
	params.PageSize = &PageSize
	return params
}

// Retrieve a list of all Subscribed Event types for a Subscription.
func (c *DefaultApiService) ListSubscribedEvent(SubscriptionSid string, params *ListSubscribedEventParams) (*ListSubscribedEventResponse, error) {
	path := "/v1/Subscriptions/{SubscriptionSid}/SubscribedEvents"
	path = strings.Replace(path, "{"+"SubscriptionSid"+"}", SubscriptionSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.client.Get(c.baseURL+path, data, headers)
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

// Optional parameters for the method 'ListSubscription'
type ListSubscriptionParams struct {
	// The SID of the sink that the list of Subscriptions should be filtered by.
	SinkSid *string `json:"SinkSid,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int32 `json:"PageSize,omitempty"`
}

func (params *ListSubscriptionParams) SetSinkSid(SinkSid string) *ListSubscriptionParams {
	params.SinkSid = &SinkSid
	return params
}
func (params *ListSubscriptionParams) SetPageSize(PageSize int32) *ListSubscriptionParams {
	params.PageSize = &PageSize
	return params
}

// Retrieve a paginated list of Subscriptions belonging to the account used to make the request.
func (c *DefaultApiService) ListSubscription(params *ListSubscriptionParams) (*ListSubscriptionResponse, error) {
	path := "/v1/Subscriptions"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.SinkSid != nil {
		data.Set("SinkSid", *params.SinkSid)
	}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListSubscriptionResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'UpdateSink'
type UpdateSinkParams struct {
	// A human readable description for the Sink **This value should not contain PII.**
	Description *string `json:"Description,omitempty"`
}

func (params *UpdateSinkParams) SetDescription(Description string) *UpdateSinkParams {
	params.Description = &Description
	return params
}

// Update a specific Sink
func (c *DefaultApiService) UpdateSink(Sid string, params *UpdateSinkParams) (*EventsV1Sink, error) {
	path := "/v1/Sinks/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Description != nil {
		data.Set("Description", *params.Description)
	}

	resp, err := c.client.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &EventsV1Sink{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'UpdateSubscribedEvent'
type UpdateSubscribedEventParams struct {
	// The schema version that the subscription should use.
	SchemaVersion *int32 `json:"SchemaVersion,omitempty"`
}

func (params *UpdateSubscribedEventParams) SetSchemaVersion(SchemaVersion int32) *UpdateSubscribedEventParams {
	params.SchemaVersion = &SchemaVersion
	return params
}

// Update an Event for a Subscription.
func (c *DefaultApiService) UpdateSubscribedEvent(SubscriptionSid string, Type string, params *UpdateSubscribedEventParams) (*EventsV1SubscriptionSubscribedEvent, error) {
	path := "/v1/Subscriptions/{SubscriptionSid}/SubscribedEvents/{Type}"
	path = strings.Replace(path, "{"+"SubscriptionSid"+"}", SubscriptionSid, -1)
	path = strings.Replace(path, "{"+"Type"+"}", Type, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.SchemaVersion != nil {
		data.Set("SchemaVersion", fmt.Sprint(*params.SchemaVersion))
	}

	resp, err := c.client.Post(c.baseURL+path, data, headers)
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

// Optional parameters for the method 'UpdateSubscription'
type UpdateSubscriptionParams struct {
	// A human readable description for the Subscription.
	Description *string `json:"Description,omitempty"`
	// The SID of the sink that events selected by this subscription should be sent to. Sink must be active for the subscription to be created.
	SinkSid *string `json:"SinkSid,omitempty"`
}

func (params *UpdateSubscriptionParams) SetDescription(Description string) *UpdateSubscriptionParams {
	params.Description = &Description
	return params
}
func (params *UpdateSubscriptionParams) SetSinkSid(SinkSid string) *UpdateSubscriptionParams {
	params.SinkSid = &SinkSid
	return params
}

// Update a Subscription.
func (c *DefaultApiService) UpdateSubscription(Sid string, params *UpdateSubscriptionParams) (*EventsV1Subscription, error) {
	path := "/v1/Subscriptions/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Description != nil {
		data.Set("Description", *params.Description)
	}
	if params != nil && params.SinkSid != nil {
		data.Set("SinkSid", *params.SinkSid)
	}

	resp, err := c.client.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &EventsV1Subscription{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
