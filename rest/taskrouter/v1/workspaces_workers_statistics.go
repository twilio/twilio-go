/*
 * Twilio - Taskrouter
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.23.0
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
)

// Optional parameters for the method 'FetchWorkerInstanceStatistics'
type FetchWorkerInstanceStatisticsParams struct {
	// Only calculate statistics since this many minutes in the past. The default 15 minutes. This is helpful for displaying statistics for the last 15 minutes, 240 minutes (4 hours), and 480 minutes (8 hours) to see trends.
	Minutes *int `json:"Minutes,omitempty"`
	// Only calculate statistics from this date and time and later, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	StartDate *time.Time `json:"StartDate,omitempty"`
	// Only include usage that occurred on or before this date, specified in GMT as an [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date-time.
	EndDate *time.Time `json:"EndDate,omitempty"`
	// Only calculate statistics on this TaskChannel. Can be the TaskChannel's SID or its `unique_name`, such as `voice`, `sms`, or `default`.
	TaskChannel *string `json:"TaskChannel,omitempty"`
}

func (params *FetchWorkerInstanceStatisticsParams) SetMinutes(Minutes int) *FetchWorkerInstanceStatisticsParams {
	params.Minutes = &Minutes
	return params
}
func (params *FetchWorkerInstanceStatisticsParams) SetStartDate(StartDate time.Time) *FetchWorkerInstanceStatisticsParams {
	params.StartDate = &StartDate
	return params
}
func (params *FetchWorkerInstanceStatisticsParams) SetEndDate(EndDate time.Time) *FetchWorkerInstanceStatisticsParams {
	params.EndDate = &EndDate
	return params
}
func (params *FetchWorkerInstanceStatisticsParams) SetTaskChannel(TaskChannel string) *FetchWorkerInstanceStatisticsParams {
	params.TaskChannel = &TaskChannel
	return params
}

func (c *ApiService) FetchWorkerInstanceStatistics(WorkspaceSid string, WorkerSid string, params *FetchWorkerInstanceStatisticsParams) (*TaskrouterV1WorkerInstanceStatistics, error) {
	path := "/v1/Workspaces/{WorkspaceSid}/Workers/{WorkerSid}/Statistics"
	path = strings.Replace(path, "{"+"WorkspaceSid"+"}", WorkspaceSid, -1)
	path = strings.Replace(path, "{"+"WorkerSid"+"}", WorkerSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Minutes != nil {
		data.Set("Minutes", fmt.Sprint(*params.Minutes))
	}
	if params != nil && params.StartDate != nil {
		data.Set("StartDate", fmt.Sprint((*params.StartDate).Format(time.RFC3339)))
	}
	if params != nil && params.EndDate != nil {
		data.Set("EndDate", fmt.Sprint((*params.EndDate).Format(time.RFC3339)))
	}
	if params != nil && params.TaskChannel != nil {
		data.Set("TaskChannel", *params.TaskChannel)
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TaskrouterV1WorkerInstanceStatistics{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'FetchWorkerStatistics'
type FetchWorkerStatisticsParams struct {
	// Only calculate statistics since this many minutes in the past. The default 15 minutes. This is helpful for displaying statistics for the last 15 minutes, 240 minutes (4 hours), and 480 minutes (8 hours) to see trends.
	Minutes *int `json:"Minutes,omitempty"`
	// Only calculate statistics from this date and time and later, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	StartDate *time.Time `json:"StartDate,omitempty"`
	// Only calculate statistics from this date and time and earlier, specified in GMT as an [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date-time.
	EndDate *time.Time `json:"EndDate,omitempty"`
	// The SID of the TaskQueue for which to fetch Worker statistics.
	TaskQueueSid *string `json:"TaskQueueSid,omitempty"`
	// The `friendly_name` of the TaskQueue for which to fetch Worker statistics.
	TaskQueueName *string `json:"TaskQueueName,omitempty"`
	// Only include Workers with `friendly_name` values that match this parameter.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// Only calculate statistics on this TaskChannel. Can be the TaskChannel's SID or its `unique_name`, such as `voice`, `sms`, or `default`.
	TaskChannel *string `json:"TaskChannel,omitempty"`
}

func (params *FetchWorkerStatisticsParams) SetMinutes(Minutes int) *FetchWorkerStatisticsParams {
	params.Minutes = &Minutes
	return params
}
func (params *FetchWorkerStatisticsParams) SetStartDate(StartDate time.Time) *FetchWorkerStatisticsParams {
	params.StartDate = &StartDate
	return params
}
func (params *FetchWorkerStatisticsParams) SetEndDate(EndDate time.Time) *FetchWorkerStatisticsParams {
	params.EndDate = &EndDate
	return params
}
func (params *FetchWorkerStatisticsParams) SetTaskQueueSid(TaskQueueSid string) *FetchWorkerStatisticsParams {
	params.TaskQueueSid = &TaskQueueSid
	return params
}
func (params *FetchWorkerStatisticsParams) SetTaskQueueName(TaskQueueName string) *FetchWorkerStatisticsParams {
	params.TaskQueueName = &TaskQueueName
	return params
}
func (params *FetchWorkerStatisticsParams) SetFriendlyName(FriendlyName string) *FetchWorkerStatisticsParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *FetchWorkerStatisticsParams) SetTaskChannel(TaskChannel string) *FetchWorkerStatisticsParams {
	params.TaskChannel = &TaskChannel
	return params
}

func (c *ApiService) FetchWorkerStatistics(WorkspaceSid string, params *FetchWorkerStatisticsParams) (*TaskrouterV1WorkerStatistics, error) {
	path := "/v1/Workspaces/{WorkspaceSid}/Workers/Statistics"
	path = strings.Replace(path, "{"+"WorkspaceSid"+"}", WorkspaceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Minutes != nil {
		data.Set("Minutes", fmt.Sprint(*params.Minutes))
	}
	if params != nil && params.StartDate != nil {
		data.Set("StartDate", fmt.Sprint((*params.StartDate).Format(time.RFC3339)))
	}
	if params != nil && params.EndDate != nil {
		data.Set("EndDate", fmt.Sprint((*params.EndDate).Format(time.RFC3339)))
	}
	if params != nil && params.TaskQueueSid != nil {
		data.Set("TaskQueueSid", *params.TaskQueueSid)
	}
	if params != nil && params.TaskQueueName != nil {
		data.Set("TaskQueueName", *params.TaskQueueName)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.TaskChannel != nil {
		data.Set("TaskChannel", *params.TaskChannel)
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TaskrouterV1WorkerStatistics{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
