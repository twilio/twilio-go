/*
 * Twilio - Video
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.9.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ListCompositionHookResponse struct for ListCompositionHookResponse
type ListCompositionHookResponse struct {
	CompositionHooks []VideoV1CompositionHook        `json:"CompositionHooks,omitempty"`
	Meta             ListCompositionHookResponseMeta `json:"Meta,omitempty"`
}
