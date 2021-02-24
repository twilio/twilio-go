/*
 * Twilio - Trunking
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ListCredentialListResponse struct for ListCredentialListResponse
type ListCredentialListResponse struct {
	CredentialLists []TrunkingV1TrunkCredentialList `json:"CredentialLists,omitempty"`
	Meta            ListTrunkResponseMeta           `json:"Meta,omitempty"`
}
