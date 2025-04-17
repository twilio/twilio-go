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

// TaskrouterV1WorkflowRealTimeStatistics struct for TaskrouterV1WorkflowRealTimeStatistics
type TaskrouterV1WorkflowRealTimeStatistics struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Workflow resource.
	AccountSid *string `json:"account_sid,omitempty"`
	// The age of the longest waiting Task.
	LongestTaskWaitingAge int `json:"longest_task_waiting_age,omitempty"`
	// The SID of the longest waiting Task.
	LongestTaskWaitingSid *string `json:"longest_task_waiting_sid,omitempty"`
	// The number of Tasks by priority. For example: `{\"0\": \"10\", \"99\": \"5\"}` shows 10 Tasks at priority 0 and 5 at priority 99.
	TasksByPriority *map[string]interface{} `json:"tasks_by_priority,omitempty"`
	// The number of Tasks by their current status. For example: `{\"pending\": \"1\", \"reserved\": \"3\", \"assigned\": \"2\", \"completed\": \"5\"}`.
	TasksByStatus *map[string]interface{} `json:"tasks_by_status,omitempty"`
	// The total number of Tasks.
	TotalTasks int `json:"total_tasks,omitempty"`
	// Returns the list of Tasks that are being controlled by the Workflow with the specified SID value.
	WorkflowSid *string `json:"workflow_sid,omitempty"`
	// The SID of the Workspace that contains the Workflow.
	WorkspaceSid *string `json:"workspace_sid,omitempty"`
	// The absolute URL of the Workflow statistics resource.
	Url *string `json:"url,omitempty"`
}
