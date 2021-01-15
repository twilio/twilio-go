/*
 * Twilio - Taskrouter
 *
 * This is the public Twilio REST API.
 *
 * API version: 5.15.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// TaskrouterV1WorkspaceWorkflowWorkflowStatistics struct for TaskrouterV1WorkspaceWorkflowWorkflowStatistics
type TaskrouterV1WorkspaceWorkflowWorkflowStatistics struct {
	AccountSid string `json:"account_sid,omitempty"`
	Cumulative map[string]interface{} `json:"cumulative,omitempty"`
	Realtime map[string]interface{} `json:"realtime,omitempty"`
	Url string `json:"url,omitempty"`
	WorkflowSid string `json:"workflow_sid,omitempty"`
	WorkspaceSid string `json:"workspace_sid,omitempty"`
}
