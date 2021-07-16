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
	"time"

	"github.com/twilio/twilio-go/client"
)

// Optional parameters for the method 'FetchTaskQueueStatistics'
type FetchTaskQueueStatisticsParams struct {
	// Only calculate statistics from this date and time and earlier, specified in GMT as an [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date-time.
	EndDate *time.Time `json:"EndDate,omitempty"`
	// Only calculate statistics since this many minutes in the past. The default is 15 minutes.
	Minutes *int `json:"Minutes,omitempty"`
	// Only calculate statistics from this date and time and later, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	StartDate *time.Time `json:"StartDate,omitempty"`
	// Only calculate real-time and cumulative statistics for the specified TaskChannel. Can be the TaskChannel's SID or its `unique_name`, such as `voice`, `sms`, or `default`.
	TaskChannel *string `json:"TaskChannel,omitempty"`
	// A comma separated list of values that describes the thresholds, in seconds, to calculate statistics on. For each threshold specified, the number of Tasks canceled and reservations accepted above and below the specified thresholds in seconds are computed.
	SplitByWaitTime *string `json:"SplitByWaitTime,omitempty"`
}

func (params *FetchTaskQueueStatisticsParams) SetEndDate(EndDate time.Time) *FetchTaskQueueStatisticsParams {
	params.EndDate = &EndDate
	return params
}
func (params *FetchTaskQueueStatisticsParams) SetMinutes(Minutes int) *FetchTaskQueueStatisticsParams {
	params.Minutes = &Minutes
	return params
}
func (params *FetchTaskQueueStatisticsParams) SetStartDate(StartDate time.Time) *FetchTaskQueueStatisticsParams {
	params.StartDate = &StartDate
	return params
}
func (params *FetchTaskQueueStatisticsParams) SetTaskChannel(TaskChannel string) *FetchTaskQueueStatisticsParams {
	params.TaskChannel = &TaskChannel
	return params
}
func (params *FetchTaskQueueStatisticsParams) SetSplitByWaitTime(SplitByWaitTime string) *FetchTaskQueueStatisticsParams {
	params.SplitByWaitTime = &SplitByWaitTime
	return params
}

func (c *ApiService) FetchTaskQueueStatistics(WorkspaceSid string, TaskQueueSid string, params *FetchTaskQueueStatisticsParams) (*TaskrouterV1WorkspaceTaskQueueTaskQueueStatistics, error) {
	path := "/v1/Workspaces/{WorkspaceSid}/TaskQueues/{TaskQueueSid}/Statistics"
	path = strings.Replace(path, "{"+"WorkspaceSid"+"}", WorkspaceSid, -1)
	path = strings.Replace(path, "{"+"TaskQueueSid"+"}", TaskQueueSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.EndDate != nil {
		data.Set("EndDate", fmt.Sprint((*params.EndDate).Format(time.RFC3339)))
	}
	if params != nil && params.Minutes != nil {
		data.Set("Minutes", fmt.Sprint(*params.Minutes))
	}
	if params != nil && params.StartDate != nil {
		data.Set("StartDate", fmt.Sprint((*params.StartDate).Format(time.RFC3339)))
	}
	if params != nil && params.TaskChannel != nil {
		data.Set("TaskChannel", *params.TaskChannel)
	}
	if params != nil && params.SplitByWaitTime != nil {
		data.Set("SplitByWaitTime", *params.SplitByWaitTime)
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TaskrouterV1WorkspaceTaskQueueTaskQueueStatistics{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListTaskQueuesStatistics'
type ListTaskQueuesStatisticsParams struct {
	// Only calculate statistics from this date and time and earlier, specified in GMT as an [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date-time.
	EndDate *time.Time `json:"EndDate,omitempty"`
	// The `friendly_name` of the TaskQueue statistics to read.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// Only calculate statistics since this many minutes in the past. The default is 15 minutes.
	Minutes *int `json:"Minutes,omitempty"`
	// Only calculate statistics from this date and time and later, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	StartDate *time.Time `json:"StartDate,omitempty"`
	// Only calculate statistics on this TaskChannel. Can be the TaskChannel's SID or its `unique_name`, such as `voice`, `sms`, or `default`.
	TaskChannel *string `json:"TaskChannel,omitempty"`
	// A comma separated list of values that describes the thresholds, in seconds, to calculate statistics on. For each threshold specified, the number of Tasks canceled and reservations accepted above and below the specified thresholds in seconds are computed.
	SplitByWaitTime *string `json:"SplitByWaitTime,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListTaskQueuesStatisticsParams) SetEndDate(EndDate time.Time) *ListTaskQueuesStatisticsParams {
	params.EndDate = &EndDate
	return params
}
func (params *ListTaskQueuesStatisticsParams) SetFriendlyName(FriendlyName string) *ListTaskQueuesStatisticsParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *ListTaskQueuesStatisticsParams) SetMinutes(Minutes int) *ListTaskQueuesStatisticsParams {
	params.Minutes = &Minutes
	return params
}
func (params *ListTaskQueuesStatisticsParams) SetStartDate(StartDate time.Time) *ListTaskQueuesStatisticsParams {
	params.StartDate = &StartDate
	return params
}
func (params *ListTaskQueuesStatisticsParams) SetTaskChannel(TaskChannel string) *ListTaskQueuesStatisticsParams {
	params.TaskChannel = &TaskChannel
	return params
}
func (params *ListTaskQueuesStatisticsParams) SetSplitByWaitTime(SplitByWaitTime string) *ListTaskQueuesStatisticsParams {
	params.SplitByWaitTime = &SplitByWaitTime
	return params
}
func (params *ListTaskQueuesStatisticsParams) SetPageSize(PageSize int) *ListTaskQueuesStatisticsParams {
	params.PageSize = &PageSize
	return params
}

// Retrieve a single page of TaskQueuesStatistics records from the API. Request is executed immediately.
func (c *ApiService) PageTaskQueuesStatistics(WorkspaceSid string, params *ListTaskQueuesStatisticsParams, pageToken string, pageNumber string) (*ListTaskQueuesStatisticsResponse, error) {
	path := "/v1/Workspaces/{WorkspaceSid}/TaskQueues/Statistics"

	path = strings.Replace(path, "{"+"WorkspaceSid"+"}", WorkspaceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.EndDate != nil {
		data.Set("EndDate", fmt.Sprint((*params.EndDate).Format(time.RFC3339)))
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.Minutes != nil {
		data.Set("Minutes", fmt.Sprint(*params.Minutes))
	}
	if params != nil && params.StartDate != nil {
		data.Set("StartDate", fmt.Sprint((*params.StartDate).Format(time.RFC3339)))
	}
	if params != nil && params.TaskChannel != nil {
		data.Set("TaskChannel", *params.TaskChannel)
	}
	if params != nil && params.SplitByWaitTime != nil {
		data.Set("SplitByWaitTime", *params.SplitByWaitTime)
	}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	if pageToken != "" {
		data.Set("PageToken", pageToken)
	}
	if pageToken != "" {
		data.Set("Page", pageNumber)
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListTaskQueuesStatisticsResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists TaskQueuesStatistics records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListTaskQueuesStatistics(WorkspaceSid string, params *ListTaskQueuesStatisticsParams, limit int) ([]TaskrouterV1WorkspaceTaskQueueTaskQueuesStatistics, error) {
	params.SetPageSize(client.ReadLimits(params.PageSize, limit))

	response, err := c.PageTaskQueuesStatistics(WorkspaceSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []TaskrouterV1WorkspaceTaskQueueTaskQueuesStatistics

	for response != nil {
		records = append(records, response.TaskQueuesStatistics...)

		var record interface{}
		if record, err = client.GetNext(response, &curRecord, limit, c.getNextListTaskQueuesStatisticsResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListTaskQueuesStatisticsResponse)
	}

	return records, err
}

// Streams TaskQueuesStatistics records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamTaskQueuesStatistics(WorkspaceSid string, params *ListTaskQueuesStatisticsParams, limit int) (chan TaskrouterV1WorkspaceTaskQueueTaskQueuesStatistics, error) {
	params.SetPageSize(client.ReadLimits(params.PageSize, limit))

	response, err := c.PageTaskQueuesStatistics(WorkspaceSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan TaskrouterV1WorkspaceTaskQueueTaskQueuesStatistics, 1)

	go func() {
		for response != nil {
			for item := range response.TaskQueuesStatistics {
				channel <- response.TaskQueuesStatistics[item]
			}

			var record interface{}
			if record, err = client.GetNext(response, &curRecord, limit, c.getNextListTaskQueuesStatisticsResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListTaskQueuesStatisticsResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListTaskQueuesStatisticsResponse(nextPageUri string) (interface{}, error) {
	if nextPageUri == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(c.baseURL+nextPageUri, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListTaskQueuesStatisticsResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
