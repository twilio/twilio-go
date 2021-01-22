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
// TaskrouterV1WorkspaceWorkflowReadResponse struct for TaskrouterV1WorkspaceWorkflowReadResponse
type TaskrouterV1WorkspaceWorkflowReadResponse struct {
	Meta TaskrouterV1WorkspaceReadResponseMeta `json:"Meta,omitempty"`
	Workflows []TaskrouterV1WorkspaceWorkflow `json:"Workflows,omitempty"`
}
