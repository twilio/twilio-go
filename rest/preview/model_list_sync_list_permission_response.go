/*
 * Twilio - Preview
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.0.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ListSyncListPermissionResponse struct for ListSyncListPermissionResponse
type ListSyncListPermissionResponse struct {
	Meta ListDayResponseMeta `json:"Meta,omitempty"`
	Permissions []PreviewSyncServiceSyncListSyncListPermission `json:"Permissions,omitempty"`
}
