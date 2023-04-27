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
	"encoding/json"
	"fmt"
	"net/url"

    "github.com/twilio/twilio-go/client"
)


// Optional parameters for the method 'FetchTaskQueueRealTimeStatistics'
type FetchTaskQueueRealTimeStatisticsParams struct {
    // The TaskChannel for which to fetch statistics. Can be the TaskChannel's SID or its `unique_name`, such as `voice`, `sms`, or `default`.
    TaskChannel *string `json:"TaskChannel,omitempty"`
}

func (params *FetchTaskQueueRealTimeStatisticsParams) SetTaskChannel(TaskChannel string) (*FetchTaskQueueRealTimeStatisticsParams){
    params.TaskChannel = &TaskChannel
    return params
}

// 
func (c *ApiService) FetchTaskQueueRealTimeStatistics(WorkspaceSid string, TaskQueueSid string, params *FetchTaskQueueRealTimeStatisticsParams) (*TaskrouterV1TaskQueueRealTimeStatistics, error) {
    path := "/v1/Workspaces/{WorkspaceSid}/TaskQueues/{TaskQueueSid}/RealTimeStatistics"
        path = strings.Replace(path, "{"+"WorkspaceSid"+"}", WorkspaceSid, -1)
    path = strings.Replace(path, "{"+"TaskQueueSid"+"}", TaskQueueSid, -1)

data := url.Values{}
headers := make(map[string]interface{})

if params != nil && params.TaskChannel != nil {
    data.Set("TaskChannel", *params.TaskChannel)
}



    resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &TaskrouterV1TaskQueueRealTimeStatistics{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}
