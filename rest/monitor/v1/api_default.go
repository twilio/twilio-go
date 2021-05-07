/*
 * Twilio - Monitor
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.15.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"

	twilio "github.com/twilio/twilio-go/client"
)

type DefaultApiService struct {
	baseURL string
	client  twilio.BaseClient
}

func NewDefaultApiService(client twilio.BaseClient) *DefaultApiService {
	return &DefaultApiService{
		client:  client,
		baseURL: "https://monitor.twilio.com",
	}
}

// FetchAlert Method for FetchAlert
//
// param: Sid The SID of the Alert resource to fetch.
//
// return: MonitorV1AlertInstance
func (c *DefaultApiService) FetchAlert(Sid string) (*MonitorV1AlertInstance, error) {
	path := "/v1/Alerts/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &MonitorV1AlertInstance{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// FetchEvent Method for FetchEvent
//
// param: Sid The SID of the Event resource to fetch.
//
// return: MonitorV1Event
func (c *DefaultApiService) FetchEvent(Sid string) (*MonitorV1Event, error) {
	path := "/v1/Events/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &MonitorV1Event{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// ListAlertParams Optional parameters for the method 'ListAlert'
type ListAlertParams struct {
	LogLevel  *string    `json:"LogLevel,omitempty"`
	StartDate *time.Time `json:"StartDate,omitempty"`
	EndDate   *time.Time `json:"EndDate,omitempty"`
	PageSize  *int32     `json:"PageSize,omitempty"`
}

func (params *ListAlertParams) SetLogLevel(LogLevel string) *ListAlertParams {
	params.LogLevel = &LogLevel
	return params
}
func (params *ListAlertParams) SetStartDate(StartDate time.Time) *ListAlertParams {
	params.StartDate = &StartDate
	return params
}
func (params *ListAlertParams) SetEndDate(EndDate time.Time) *ListAlertParams {
	params.EndDate = &EndDate
	return params
}
func (params *ListAlertParams) SetPageSize(PageSize int32) *ListAlertParams {
	params.PageSize = &PageSize
	return params
}

// ListAlert Method for ListAlert
//
// param: optional nil or *ListAlertParams - Optional Parameters:
//
// param: "LogLevel" (string) - Only show alerts for this log-level.  Can be: `error`, `warning`, `notice`, or `debug`.
//
// param: "StartDate" (time.Time) - Only include alerts that occurred on or after this date and time. Specify the date and time in GMT and format as `YYYY-MM-DD` or `YYYY-MM-DDThh:mm:ssZ`. Queries for alerts older than 30 days are not supported.
//
// param: "EndDate" (time.Time) - Only include alerts that occurred on or before this date and time. Specify the date and time in GMT and format as `YYYY-MM-DD` or `YYYY-MM-DDThh:mm:ssZ`. Queries for alerts older than 30 days are not supported.
//
// param: "PageSize" (int32) - How many resources to return in each list page. The default is 50, and the maximum is 1000.
//
// return: ListAlertResponse
func (c *DefaultApiService) ListAlert(params *ListAlertParams) (*ListAlertResponse, error) {
	path := "/v1/Alerts"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.LogLevel != nil {
		data.Set("LogLevel", *params.LogLevel)
	}
	if params != nil && params.StartDate != nil {
		data.Set("StartDate", fmt.Sprint((*params.StartDate).Format(time.RFC3339)))
	}
	if params != nil && params.EndDate != nil {
		data.Set("EndDate", fmt.Sprint((*params.EndDate).Format(time.RFC3339)))
	}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListAlertResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// ListEventParams Optional parameters for the method 'ListEvent'
type ListEventParams struct {
	ActorSid        *string    `json:"ActorSid,omitempty"`
	EventType       *string    `json:"EventType,omitempty"`
	ResourceSid     *string    `json:"ResourceSid,omitempty"`
	SourceIpAddress *string    `json:"SourceIpAddress,omitempty"`
	StartDate       *time.Time `json:"StartDate,omitempty"`
	EndDate         *time.Time `json:"EndDate,omitempty"`
	PageSize        *int32     `json:"PageSize,omitempty"`
}

func (params *ListEventParams) SetActorSid(ActorSid string) *ListEventParams {
	params.ActorSid = &ActorSid
	return params
}
func (params *ListEventParams) SetEventType(EventType string) *ListEventParams {
	params.EventType = &EventType
	return params
}
func (params *ListEventParams) SetResourceSid(ResourceSid string) *ListEventParams {
	params.ResourceSid = &ResourceSid
	return params
}
func (params *ListEventParams) SetSourceIpAddress(SourceIpAddress string) *ListEventParams {
	params.SourceIpAddress = &SourceIpAddress
	return params
}
func (params *ListEventParams) SetStartDate(StartDate time.Time) *ListEventParams {
	params.StartDate = &StartDate
	return params
}
func (params *ListEventParams) SetEndDate(EndDate time.Time) *ListEventParams {
	params.EndDate = &EndDate
	return params
}
func (params *ListEventParams) SetPageSize(PageSize int32) *ListEventParams {
	params.PageSize = &PageSize
	return params
}

// ListEvent Method for ListEvent
//
// Returns a list of events in the account, sorted by event-date.
//
// param: optional nil or *ListEventParams - Optional Parameters:
//
// param: "ActorSid" (string) - Only include events initiated by this Actor. Useful for auditing actions taken by specific users or API credentials.
//
// param: "EventType" (string) - Only include events of this [Event Type](https://www.twilio.com/docs/usage/monitor-events#event-types).
//
// param: "ResourceSid" (string) - Only include events that refer to this resource. Useful for discovering the history of a specific resource.
//
// param: "SourceIpAddress" (string) - Only include events that originated from this IP address. Useful for tracking suspicious activity originating from the API or the Twilio Console.
//
// param: "StartDate" (time.Time) - Only include events that occurred on or after this date. Specify the date in GMT and [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
//
// param: "EndDate" (time.Time) - Only include events that occurred on or before this date. Specify the date in GMT and [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
//
// param: "PageSize" (int32) - How many resources to return in each list page. The default is 50, and the maximum is 1000.
//
// return: ListEventResponse
func (c *DefaultApiService) ListEvent(params *ListEventParams) (*ListEventResponse, error) {
	path := "/v1/Events"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.ActorSid != nil {
		data.Set("ActorSid", *params.ActorSid)
	}
	if params != nil && params.EventType != nil {
		data.Set("EventType", *params.EventType)
	}
	if params != nil && params.ResourceSid != nil {
		data.Set("ResourceSid", *params.ResourceSid)
	}
	if params != nil && params.SourceIpAddress != nil {
		data.Set("SourceIpAddress", *params.SourceIpAddress)
	}
	if params != nil && params.StartDate != nil {
		data.Set("StartDate", fmt.Sprint((*params.StartDate).Format(time.RFC3339)))
	}
	if params != nil && params.EndDate != nil {
		data.Set("EndDate", fmt.Sprint((*params.EndDate).Format(time.RFC3339)))
	}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListEventResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
