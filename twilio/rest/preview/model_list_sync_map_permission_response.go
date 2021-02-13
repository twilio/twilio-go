/*
 * Twilio - Preview
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.9.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ListSyncMapPermissionResponse struct for ListSyncMapPermissionResponse
type ListSyncMapPermissionResponse struct {
	Meta        ListDayResponseMeta                          `json:"Meta,omitempty"`
	Permissions []PreviewSyncServiceSyncMapSyncMapPermission `json:"Permissions,omitempty"`
}
