/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 5.15.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// InlineResponse20036 struct for InlineResponse20036
type InlineResponse20036 struct {
	AddOnResults []ApiV2010AccountRecordingRecordingAddOnResult `json:"add_on_results,omitempty"`
	End int32 `json:"end,omitempty"`
	FirstPageUri string `json:"first_page_uri,omitempty"`
	NextPageUri string `json:"next_page_uri,omitempty"`
	Page int32 `json:"page,omitempty"`
	PageSize int32 `json:"page_size,omitempty"`
	PreviousPageUri string `json:"previous_page_uri,omitempty"`
	Start int32 `json:"start,omitempty"`
	Uri string `json:"uri,omitempty"`
}
