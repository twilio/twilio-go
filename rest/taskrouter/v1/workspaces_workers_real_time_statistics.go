/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Taskrouter
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
	"net/url"
	"strings"
)

// Optional parameters for the method 'FetchWorkersRealTimeStatistics'
type FetchWorkersRealTimeStatisticsParams struct {
	// Only calculate real-time statistics on this TaskChannel. Can be the TaskChannel's SID or its `unique_name`, such as `voice`, `sms`, or `default`.
	TaskChannel *string `json:"TaskChannel,omitempty"`
}

func (params *FetchWorkersRealTimeStatisticsParams) SetTaskChannel(TaskChannel string) *FetchWorkersRealTimeStatisticsParams {
	params.TaskChannel = &TaskChannel
	return params
}

//
func (c *ApiService) FetchWorkersRealTimeStatistics(WorkspaceSid string, params *FetchWorkersRealTimeStatisticsParams) (*TaskrouterV1WorkersRealTimeStatistics, error) {
	return c.FetchWorkersRealTimeStatisticsWithCtx(context.TODO(), WorkspaceSid, params)
}

//
func (c *ApiService) FetchWorkersRealTimeStatisticsWithCtx(ctx context.Context, WorkspaceSid string, params *FetchWorkersRealTimeStatisticsParams) (*TaskrouterV1WorkersRealTimeStatistics, error) {
	path := "/v1/Workspaces/{WorkspaceSid}/Workers/RealTimeStatistics"
	path = strings.Replace(path, "{"+"WorkspaceSid"+"}", WorkspaceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.TaskChannel != nil {
		data.Set("TaskChannel", *params.TaskChannel)
	}

	resp, err := c.requestHandler.Get(ctx, c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TaskrouterV1WorkersRealTimeStatistics{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
