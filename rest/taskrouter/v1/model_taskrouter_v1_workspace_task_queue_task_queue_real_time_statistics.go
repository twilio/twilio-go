/*
 * Twilio - Taskrouter
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.8.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics struct for TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics
type TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics struct {
	AccountSid string `json:"AccountSid,omitempty"`
	ActivityStatistics []map[string]interface{} `json:"ActivityStatistics,omitempty"`
	LongestRelativeTaskAgeInQueue int32 `json:"LongestRelativeTaskAgeInQueue,omitempty"`
	LongestRelativeTaskSidInQueue string `json:"LongestRelativeTaskSidInQueue,omitempty"`
	LongestTaskWaitingAge int32 `json:"LongestTaskWaitingAge,omitempty"`
	LongestTaskWaitingSid string `json:"LongestTaskWaitingSid,omitempty"`
	TaskQueueSid string `json:"TaskQueueSid,omitempty"`
	TasksByPriority map[string]interface{} `json:"TasksByPriority,omitempty"`
	TasksByStatus map[string]interface{} `json:"TasksByStatus,omitempty"`
	TotalAvailableWorkers int32 `json:"TotalAvailableWorkers,omitempty"`
	TotalEligibleWorkers int32 `json:"TotalEligibleWorkers,omitempty"`
	TotalTasks int32 `json:"TotalTasks,omitempty"`
	Url string `json:"Url,omitempty"`
	WorkspaceSid string `json:"WorkspaceSid,omitempty"`
}
