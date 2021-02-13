/*
 * Twilio - Taskrouter
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.9.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ListTaskResponse struct for ListTaskResponse
type ListTaskResponse struct {
	Meta  ListWorkspaceResponseMeta   `json:"Meta,omitempty"`
	Tasks []TaskrouterV1WorkspaceTask `json:"Tasks,omitempty"`
}
