/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.9.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ListIncomingPhoneNumberResponse struct for ListIncomingPhoneNumberResponse
type ListIncomingPhoneNumberResponse struct {
	End                  int32                                `json:"End,omitempty"`
	FirstPageUri         string                               `json:"FirstPageUri,omitempty"`
	IncomingPhoneNumbers []ApiV2010AccountIncomingPhoneNumber `json:"IncomingPhoneNumbers,omitempty"`
	NextPageUri          string                               `json:"NextPageUri,omitempty"`
	Page                 int32                                `json:"Page,omitempty"`
	PageSize             int32                                `json:"PageSize,omitempty"`
	PreviousPageUri      string                               `json:"PreviousPageUri,omitempty"`
	Start                int32                                `json:"Start,omitempty"`
	Uri                  string                               `json:"Uri,omitempty"`
}
