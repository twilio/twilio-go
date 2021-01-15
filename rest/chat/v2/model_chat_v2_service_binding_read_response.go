/*
 * Twilio - Chat
 *
 * This is the public Twilio REST API.
 *
 * API version: 5.15.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ChatV2ServiceBindingReadResponse struct for ChatV2ServiceBindingReadResponse
type ChatV2ServiceBindingReadResponse struct {
	Bindings []ChatV2ServiceBinding `json:"bindings,omitempty"`
	Meta ChatV2CredentialReadResponseMeta `json:"meta,omitempty"`
}
