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

// TaskrouterV1WorkersRealTimeStatistics struct for TaskrouterV1WorkersRealTimeStatistics
type TaskrouterV1WorkersRealTimeStatistics struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The number of current Workers by Activity
	ActivityStatistics *[]map[string]interface{} `json:"activity_statistics,omitempty"`
	// The total number of Workers
	TotalWorkers *int `json:"total_workers,omitempty"`
	// The absolute URL of the Workers statistics resource
	Url *string `json:"url,omitempty"`
	// The SID of the Workspace that contains the Workers
	WorkspaceSid *string `json:"workspace_sid,omitempty"`
}
