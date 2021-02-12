/*
 * Twilio - Events
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.9.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"encoding/json"
	"fmt"
	twilio "github.com/twilio/twilio-go/client"
	"net/url"
)

type DefaultApiService struct {
	baseURL string
	client  *twilio.Client
}

func NewDefaultApiService(client *twilio.Client) *DefaultApiService {
	return &DefaultApiService {
		client: client,
		baseURL: fmt.Sprintf("https://events.twilio.com"),
	}
}
// CreateSinkParams Optional parameters for the method 'CreateSink'
type CreateSinkParams struct {
	Description *string `json:"Description,omitempty"`
	SinkConfiguration *map[string]interface{} `json:"SinkConfiguration,omitempty"`
	SinkType *string `json:"SinkType,omitempty"`
}

/*
CreateSink Method for CreateSink
Create a new Sink
 * @param optional nil or *CreateSinkOpts - Optional Parameters:
 * @param "Description" (string) - A human readable description for the Sink **This value should not contain PII.**
 * @param "SinkConfiguration" (map[string]interface{}) - The information required for Twilio to connect to the provided Sink encoded as JSON.
 * @param "SinkType" (string) - The Sink type. Can only be \\\"kinesis\\\" or \\\"webhook\\\" currently.
@return EventsV1Sink
*/
func (c *DefaultApiService) CreateSink(params *CreateSinkParams) (*EventsV1Sink, error) {
	path := "/v1/Sinks"


	data := url.Values{}
	headers := 0

	if params != nil && params.Description != nil {
		data.Set("Description", *params.Description) 
	}
	if params != nil && params.SinkConfiguration != nil {
		v, err := json.Marshal(params.SinkConfiguration)

		if err != nil {
			return nil, err
		}

		data.Set("SinkConfiguration", fmt.Sprint(v))
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

/*
CreateSinkTest Method for CreateSinkTest
Create a new Sink Test Event for the given Sink.
 * @param Sid A 34 character string that uniquely identifies the Sink to be Tested.
@return EventsV1SinkSinkTest
*/
func (c *DefaultApiService) CreateSinkTest(Sid string) (*EventsV1SinkSinkTest, error) {
	path := "/v1/Sinks/{Sid}/Test"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)


	data := url.Values{}
	headers := 0



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
// CreateSinkValidateParams Optional parameters for the method 'CreateSinkValidate'
type CreateSinkValidateParams struct {
	TestId *string `json:"TestId,omitempty"`
}

/*
CreateSinkValidate Method for CreateSinkValidate
Validate that a test event for a Sink was received.
 * @param Sid
 * @param optional nil or *CreateSinkValidateOpts - Optional Parameters:
 * @param "TestId" (string) - A 34 character string that uniquely identifies the test event for a Sink being validated.
@return EventsV1SinkSinkValidate
*/
func (c *DefaultApiService) CreateSinkValidate(Sid string, params *CreateSinkValidateParams) (*EventsV1SinkSinkValidate, error) {
	path := "/v1/Sinks/{Sid}/Validate"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)


	data := url.Values{}
	headers := 0

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
// CreateSubscriptionParams Optional parameters for the method 'CreateSubscription'
type CreateSubscriptionParams struct {
	Description *string `json:"Description,omitempty"`
	SinkSid *string `json:"SinkSid,omitempty"`
	Types *[]map[string]interface{} `json:"Types,omitempty"`
}

/*
CreateSubscription Method for CreateSubscription
Create a new Subscription.
 * @param optional nil or *CreateSubscriptionOpts - Optional Parameters:
 * @param "Description" (string) - A human readable description for the Subscription **This value should not contain PII.**
 * @param "SinkSid" (string) - The SID of the sink that events selected by this subscription should be sent to. Sink must be active for the subscription to be created.
 * @param "Types" ([]map[string]interface{}) - Contains a dictionary of URL links to nested resources of this Subscription.
@return EventsV1Subscription
*/
func (c *DefaultApiService) CreateSubscription(params *CreateSubscriptionParams) (*EventsV1Subscription, error) {
	path := "/v1/Subscriptions"


	data := url.Values{}
	headers := 0

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

		data.Set("Types", fmt.Sprint(v))
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

/*
DeleteSink Method for DeleteSink
Delete a specific Sink.
 * @param Sid A 34 character string that uniquely identifies this Sink.
*/
func (c *DefaultApiService) DeleteSink(Sid string) (error) {
	path := "/v1/Sinks/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)


	data := url.Values{}
	headers := 0



	resp, err := c.client.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

/*
DeleteSubscription Method for DeleteSubscription
Delete a specific Subscription.
 * @param Sid A 34 character string that uniquely identifies this Subscription.
*/
func (c *DefaultApiService) DeleteSubscription(Sid string) (error) {
	path := "/v1/Subscriptions/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)


	data := url.Values{}
	headers := 0



	resp, err := c.client.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

/*
FetchEventType Method for FetchEventType
Fetch a specific Event Type.
 * @param Type A string that uniquely identifies this Event Type.
@return EventsV1EventType
*/
func (c *DefaultApiService) FetchEventType(Type string) (*EventsV1EventType, error) {
	path := "/v1/Types/{Type}"
	path = strings.Replace(path, "{"+"Type"+"}", Type, -1)


	data := url.Values{}
	headers := 0



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

/*
FetchSchema Method for FetchSchema
Fetch a specific schema with its nested versions.
 * @param Id The unique identifier of the schema. Each schema can have multiple versions, that share the same id.
@return EventsV1Schema
*/
func (c *DefaultApiService) FetchSchema(Id string) (*EventsV1Schema, error) {
	path := "/v1/Schemas/{Id}"
	path = strings.Replace(path, "{"+"Id"+"}", Id, -1)


	data := url.Values{}
	headers := 0



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

/*
FetchSink Method for FetchSink
Fetch a specific Sink.
 * @param Sid A 34 character string that uniquely identifies this Sink.
@return EventsV1Sink
*/
func (c *DefaultApiService) FetchSink(Sid string) (*EventsV1Sink, error) {
	path := "/v1/Sinks/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)


	data := url.Values{}
	headers := 0



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

/*
FetchSubscription Method for FetchSubscription
Fetch a specific Subscription.
 * @param Sid A 34 character string that uniquely identifies this Subscription.
@return EventsV1Subscription
*/
func (c *DefaultApiService) FetchSubscription(Sid string) (*EventsV1Subscription, error) {
	path := "/v1/Subscriptions/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)


	data := url.Values{}
	headers := 0



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

/*
FetchVersion Method for FetchVersion
Fetch a specific schema and version.
 * @param Id The unique identifier of the schema. Each schema can have multiple versions, that share the same id.
 * @param SchemaVersion The version of the schema
@return EventsV1SchemaVersion
*/
func (c *DefaultApiService) FetchVersion(Id string, SchemaVersion int32) (*EventsV1SchemaVersion, error) {
	path := "/v1/Schemas/{Id}/Versions/{SchemaVersion}"
	path = strings.Replace(path, "{"+"Id"+"}", Id, -1)
	path = strings.Replace(path, "{"+"SchemaVersion"+"}", fmt.Sprint(SchemaVersion), -1)


	data := url.Values{}
	headers := 0



	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &EventsV1SchemaVersion{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
// ListEventTypeParams Optional parameters for the method 'ListEventType'
type ListEventTypeParams struct {
	PageSize *int32 `json:"PageSize,omitempty"`
}

/*
ListEventType Method for ListEventType
Retrieve a paginated list of all the available Event Types.
 * @param optional nil or *ListEventTypeOpts - Optional Parameters:
 * @param "PageSize" (int32) - How many resources to return in each list page. The default is 50, and the maximum is 1000.
@return ListEventTypeResponse
*/
func (c *DefaultApiService) ListEventType(params *ListEventTypeParams) (*ListEventTypeResponse, error) {
	path := "/v1/Types"


	data := url.Values{}
	headers := 0

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
// ListSinkParams Optional parameters for the method 'ListSink'
type ListSinkParams struct {
	PageSize *int32 `json:"PageSize,omitempty"`
}

/*
ListSink Method for ListSink
Retrieve a paginated list of Sinks belonging to the account used to make the request.
 * @param optional nil or *ListSinkOpts - Optional Parameters:
 * @param "PageSize" (int32) - How many resources to return in each list page. The default is 50, and the maximum is 1000.
@return ListSinkResponse
*/
func (c *DefaultApiService) ListSink(params *ListSinkParams) (*ListSinkResponse, error) {
	path := "/v1/Sinks"


	data := url.Values{}
	headers := 0

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
// ListSubscribedEventParams Optional parameters for the method 'ListSubscribedEvent'
type ListSubscribedEventParams struct {
	PageSize *int32 `json:"PageSize,omitempty"`
}

/*
ListSubscribedEvent Method for ListSubscribedEvent
Retrieve a list of all Subscribed Event types for a Subscription.
 * @param SubscriptionSid The unique SID identifier of the Subscription.
 * @param optional nil or *ListSubscribedEventOpts - Optional Parameters:
 * @param "PageSize" (int32) - How many resources to return in each list page. The default is 50, and the maximum is 1000.
@return ListSubscribedEventResponse
*/
func (c *DefaultApiService) ListSubscribedEvent(SubscriptionSid string, params *ListSubscribedEventParams) (*ListSubscribedEventResponse, error) {
	path := "/v1/Subscriptions/{SubscriptionSid}/SubscribedEvents"
	path = strings.Replace(path, "{"+"SubscriptionSid"+"}", SubscriptionSid, -1)


	data := url.Values{}
	headers := 0

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
// ListSubscriptionParams Optional parameters for the method 'ListSubscription'
type ListSubscriptionParams struct {
	SinkSid *string `json:"SinkSid,omitempty"`
	PageSize *int32 `json:"PageSize,omitempty"`
}

/*
ListSubscription Method for ListSubscription
Retrieve a paginated list of Subscriptions belonging to the account used to make the request.
 * @param optional nil or *ListSubscriptionOpts - Optional Parameters:
 * @param "SinkSid" (string) - The SID of the sink that the list of Subscriptions should be filtered by.
 * @param "PageSize" (int32) - How many resources to return in each list page. The default is 50, and the maximum is 1000.
@return ListSubscriptionResponse
*/
func (c *DefaultApiService) ListSubscription(params *ListSubscriptionParams) (*ListSubscriptionResponse, error) {
	path := "/v1/Subscriptions"


	data := url.Values{}
	headers := 0

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
// ListVersionParams Optional parameters for the method 'ListVersion'
type ListVersionParams struct {
	PageSize *int32 `json:"PageSize,omitempty"`
}

/*
ListVersion Method for ListVersion
Retrieve a paginated list of versions of the schema.
 * @param Id The unique identifier of the schema. Each schema can have multiple versions, that share the same id.
 * @param optional nil or *ListVersionOpts - Optional Parameters:
 * @param "PageSize" (int32) - How many resources to return in each list page. The default is 50, and the maximum is 1000.
@return ListVersionResponse
*/
func (c *DefaultApiService) ListVersion(Id string, params *ListVersionParams) (*ListVersionResponse, error) {
	path := "/v1/Schemas/{Id}/Versions"
	path = strings.Replace(path, "{"+"Id"+"}", Id, -1)


	data := url.Values{}
	headers := 0

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize)) 
	}


	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListVersionResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
// UpdateSubscriptionParams Optional parameters for the method 'UpdateSubscription'
type UpdateSubscriptionParams struct {
	Description *string `json:"Description,omitempty"`
	SinkSid *string `json:"SinkSid,omitempty"`
}

/*
UpdateSubscription Method for UpdateSubscription
Update a Subscription.
 * @param Sid A 34 character string that uniquely identifies this Subscription.
 * @param optional nil or *UpdateSubscriptionOpts - Optional Parameters:
 * @param "Description" (string) - A human readable description for the Subscription.
 * @param "SinkSid" (string) - The SID of the sink that events selected by this subscription should be sent to. Sink must be active for the subscription to be created.
@return EventsV1Subscription
*/
func (c *DefaultApiService) UpdateSubscription(Sid string, params *UpdateSubscriptionParams) (*EventsV1Subscription, error) {
	path := "/v1/Subscriptions/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)


	data := url.Values{}
	headers := 0

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
