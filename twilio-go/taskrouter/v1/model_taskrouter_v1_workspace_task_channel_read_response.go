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
// TaskrouterV1WorkspaceTaskChannelReadResponse struct for TaskrouterV1WorkspaceTaskChannelReadResponse
type TaskrouterV1WorkspaceTaskChannelReadResponse struct {
	Channels []TaskrouterV1WorkspaceTaskChannel `json:"channels,omitempty"`
	Meta TaskrouterV1WorkspaceReadResponseMeta `json:"meta,omitempty"`
}
