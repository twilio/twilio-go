/*
 * Twilio - Autopilot
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.8.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ListAssistantResponseMeta struct for ListAssistantResponseMeta
type ListAssistantResponseMeta struct {
	FirstPageUrl string `json:"FirstPageUrl,omitempty"`
	Key string `json:"Key,omitempty"`
	NextPageUrl string `json:"NextPageUrl,omitempty"`
	Page int32 `json:"Page,omitempty"`
	PageSize int32 `json:"PageSize,omitempty"`
	PreviousPageUrl string `json:"PreviousPageUrl,omitempty"`
	Url string `json:"Url,omitempty"`
}
