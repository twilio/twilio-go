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
// ApiV2010AccountMessageMediaReadResponse struct for ApiV2010AccountMessageMediaReadResponse
type ApiV2010AccountMessageMediaReadResponse struct {
	End int32 `json:"End,omitempty"`
	FirstPageUri string `json:"FirstPageUri,omitempty"`
	MediaList []ApiV2010AccountMessageMedia `json:"MediaList,omitempty"`
	NextPageUri string `json:"NextPageUri,omitempty"`
	Page int32 `json:"Page,omitempty"`
	PageSize int32 `json:"PageSize,omitempty"`
	PreviousPageUri string `json:"PreviousPageUri,omitempty"`
	Start int32 `json:"Start,omitempty"`
	Uri string `json:"Uri,omitempty"`
}
