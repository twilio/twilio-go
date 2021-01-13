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
// TaskrouterV1WorkspaceTaskQueueReadResponse struct for TaskrouterV1WorkspaceTaskQueueReadResponse
type TaskrouterV1WorkspaceTaskQueueReadResponse struct {
	Meta TaskrouterV1WorkspaceReadResponseMeta `json:"meta,omitempty"`
	TaskQueues []TaskrouterV1WorkspaceTaskQueue `json:"task_queues,omitempty"`
}
