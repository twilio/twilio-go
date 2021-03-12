/*
 * Twilio - Ip_messaging
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListRoleResponse struct for ListRoleResponse
type ListRoleResponse struct {
	Meta  ListCredentialResponseMeta `json:"Meta,omitempty"`
	Roles []IpMessagingV1ServiceRole `json:"Roles,omitempty"`
}
