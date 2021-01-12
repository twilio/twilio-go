/*
 * Twilio - Sync
 *
 * This is the public Twilio REST API.
 *
 * API version: 5.15.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// InlineResponse2001 struct for InlineResponse2001
type InlineResponse2001 struct {
	Documents []SyncV1ServiceDocument `json:"documents,omitempty"`
	Meta InlineResponse200Meta `json:"meta,omitempty"`
}
