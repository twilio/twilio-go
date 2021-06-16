/*
 * Twilio - Taskrouter
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.17.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// TaskrouterV1WorkspaceWorkspaceStatistics struct for TaskrouterV1WorkspaceWorkspaceStatistics
type TaskrouterV1WorkspaceWorkspaceStatistics struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// An object that contains the cumulative statistics for the Workspace
	Cumulative *map[string]interface{} `json:"cumulative,omitempty"`
	// n object that contains the real-time statistics for the Workspace
	Realtime *map[string]interface{} `json:"realtime,omitempty"`
	// The absolute URL of the Workspace statistics resource
	Url *string `json:"url,omitempty"`
	// The SID of the Workspace
	WorkspaceSid *string `json:"workspace_sid,omitempty"`
}
