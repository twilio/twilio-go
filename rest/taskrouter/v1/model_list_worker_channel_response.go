/*
 * Twilio - Taskrouter
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.29.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListWorkerChannelResponse struct for ListWorkerChannelResponse
type ListWorkerChannelResponse struct {
	Channels []TaskrouterV1WorkerChannel `json:"channels,omitempty"`
	Meta     ListWorkspaceResponseMeta   `json:"meta,omitempty"`
}
