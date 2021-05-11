/*
 * Twilio - Taskrouter
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.15.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// TaskrouterV1WorkspaceWorkerWorkerInstanceStatistics struct for TaskrouterV1WorkspaceWorkerWorkerInstanceStatistics
type TaskrouterV1WorkspaceWorkerWorkerInstanceStatistics struct {
	AccountSid   *string                 `json:"account_sid,omitempty"`
	Cumulative   *map[string]interface{} `json:"cumulative,omitempty"`
	Url          *string                 `json:"url,omitempty"`
	WorkerSid    *string                 `json:"worker_sid,omitempty"`
	WorkspaceSid *string                 `json:"workspace_sid,omitempty"`
}
