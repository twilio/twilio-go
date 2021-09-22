/*
 * Twilio - Taskrouter
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.20.3
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// TaskrouterV1WorkflowRealTimeStatistics struct for TaskrouterV1WorkflowRealTimeStatistics
type TaskrouterV1WorkflowRealTimeStatistics struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The age of the longest waiting Task
	LongestTaskWaitingAge *int `json:"longest_task_waiting_age,omitempty"`
	// The SID of the longest waiting Task
	LongestTaskWaitingSid *string `json:"longest_task_waiting_sid,omitempty"`
	// The number of Tasks by priority
	TasksByPriority *map[string]interface{} `json:"tasks_by_priority,omitempty"`
	// The number of Tasks by their current status
	TasksByStatus *map[string]interface{} `json:"tasks_by_status,omitempty"`
	// The total number of Tasks
	TotalTasks *int `json:"total_tasks,omitempty"`
	// The absolute URL of the Workflow statistics resource
	Url *string `json:"url,omitempty"`
	// Returns the list of Tasks that are being controlled by the Workflow with the specified SID value
	WorkflowSid *string `json:"workflow_sid,omitempty"`
	// The SID of the Workspace that contains the Workflow.
	WorkspaceSid *string `json:"workspace_sid,omitempty"`
}