/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.15.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListAvailablePhoneNumberTollFreeResponse struct for ListAvailablePhoneNumberTollFreeResponse
type ListAvailablePhoneNumberTollFreeResponse struct {
	AvailablePhoneNumbers []ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberTollFree `json:"available_phone_numbers,omitempty"`
	End                   int32                                                                    `json:"end,omitempty"`
	FirstPageUri          string                                                                   `json:"first_page_uri,omitempty"`
	NextPageUri           string                                                                   `json:"next_page_uri,omitempty"`
	Page                  int32                                                                    `json:"page,omitempty"`
	PageSize              int32                                                                    `json:"page_size,omitempty"`
	PreviousPageUri       string                                                                   `json:"previous_page_uri,omitempty"`
	Start                 int32                                                                    `json:"start,omitempty"`
	Uri                   string                                                                   `json:"uri,omitempty"`
}
