/*
 * Twilio - Taskrouter
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.0.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// UpdateTaskRequest struct for UpdateTaskRequest
type UpdateTaskRequest struct {
	// The new status of the task. Can be: `canceled`, to cancel a Task that is currently `pending` or `reserved`; `wrapping`, to move the Task to wrapup state; or `completed`, to move a Task to the completed state.
	AssignmentStatus string `json:"AssignmentStatus,omitempty"`
	// The JSON string that describes the custom attributes of the task.
	Attributes string `json:"Attributes,omitempty"`
	// The Task's new priority value. When supplied, the Task takes on the specified priority unless it matches a Workflow Target with a Priority set. Value can be 0 to 2^31^ (2,147,483,647).
	Priority int32 `json:"Priority,omitempty"`
	// The reason that the Task was canceled or completed. This parameter is required only if the Task is canceled or completed. Setting this value queues the task for deletion and logs the reason.
	Reason string `json:"Reason,omitempty"`
	// When MultiTasking is enabled, specify the TaskChannel with the task to update. Can be the TaskChannel's SID or its `unique_name`, such as `voice`, `sms`, or `default`.
	TaskChannel string `json:"TaskChannel,omitempty"`
}
