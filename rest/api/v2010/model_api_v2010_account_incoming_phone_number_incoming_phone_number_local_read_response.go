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
// ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberLocalReadResponse struct for ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberLocalReadResponse
type ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberLocalReadResponse struct {
	End int32 `json:"end,omitempty"`
	FirstPageUri string `json:"first_page_uri,omitempty"`
	IncomingPhoneNumbers []ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberLocal `json:"incoming_phone_numbers,omitempty"`
	NextPageUri string `json:"next_page_uri,omitempty"`
	Page int32 `json:"page,omitempty"`
	PageSize int32 `json:"page_size,omitempty"`
	PreviousPageUri string `json:"previous_page_uri,omitempty"`
	Start int32 `json:"start,omitempty"`
	Uri string `json:"uri,omitempty"`
}
