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
// UpdateSampleRequest struct for UpdateSampleRequest
type UpdateSampleRequest struct {
	// An ISO language-country string of the sample.
	Language string `json:"Language,omitempty"`
	// The communication channel the sample was captured. It can be: *voice*, *sms*, *chat*, *alexa*, *google-assistant*, or *slack*. If not included the value will be null
	SourceChannel string `json:"SourceChannel,omitempty"`
	// The text example of how end-users may express this task. The sample may contain Field tag blocks.
	TaggedText string `json:"TaggedText,omitempty"`
}
