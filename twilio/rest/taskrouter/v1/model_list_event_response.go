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

// ListEventResponse struct for ListEventResponse
type ListEventResponse struct {
	Events []TaskrouterV1WorkspaceEvent `json:"Events,omitempty"`
	Meta   ListWorkspaceResponseMeta    `json:"Meta,omitempty"`
}
