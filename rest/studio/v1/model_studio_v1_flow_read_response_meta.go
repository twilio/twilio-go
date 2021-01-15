/*
 * Twilio - Studio
 *
 * This is the public Twilio REST API.
 *
 * API version: 5.15.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// StudioV1FlowReadResponseMeta struct for StudioV1FlowReadResponseMeta
type StudioV1FlowReadResponseMeta struct {
	FirstPageUrl string `json:"first_page_url,omitempty"`
	Key string `json:"key,omitempty"`
	NextPageUrl string `json:"next_page_url,omitempty"`
	Page int32 `json:"page,omitempty"`
	PageSize int32 `json:"page_size,omitempty"`
	PreviousPageUrl string `json:"previous_page_url,omitempty"`
	Url string `json:"url,omitempty"`
}
