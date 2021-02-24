/*
 * Twilio - Taskrouter
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// TaskrouterV1WorkspaceWorkflowWorkflowRealTimeStatistics struct for TaskrouterV1WorkspaceWorkflowWorkflowRealTimeStatistics
type TaskrouterV1WorkspaceWorkflowWorkflowRealTimeStatistics struct {
	AccountSid            *string                 `json:"AccountSid,omitempty"`
	LongestTaskWaitingAge *int32                  `json:"LongestTaskWaitingAge,omitempty"`
	LongestTaskWaitingSid *string                 `json:"LongestTaskWaitingSid,omitempty"`
	TasksByPriority       *map[string]interface{} `json:"TasksByPriority,omitempty"`
	TasksByStatus         *map[string]interface{} `json:"TasksByStatus,omitempty"`
	TotalTasks            *int32                  `json:"TotalTasks,omitempty"`
	Url                   *string                 `json:"Url,omitempty"`
	WorkflowSid           *string                 `json:"WorkflowSid,omitempty"`
	WorkspaceSid          *string                 `json:"WorkspaceSid,omitempty"`
}
