/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.8.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// CreateValidationRequestRequest struct for CreateValidationRequestRequest
type CreateValidationRequestRequest struct {
	// The number of seconds to delay before initiating the verification call. Can be an integer between `0` and `60`, inclusive. The default is `0`.
	CallDelay int32 `json:"CallDelay,omitempty"`
	// The digits to dial after connecting the verification call.
	Extension string `json:"Extension,omitempty"`
	// A descriptive string that you create to describe the new caller ID resource. It can be up to 64 characters long. The default value is a formatted version of the phone number.
	FriendlyName string `json:"FriendlyName,omitempty"`
	// The phone number to verify in [E.164](https://www.twilio.com/docs/glossary/what-e164) format, which consists of a + followed by the country code and subscriber number.
	PhoneNumber string `json:"PhoneNumber"`
	// The URL we should call using the `status_callback_method` to send status information about the verification process to your application.
	StatusCallback string `json:"StatusCallback,omitempty"`
	// The HTTP method we should use to call `status_callback`. Can be: `GET` or `POST`, and the default is `POST`.
	StatusCallbackMethod string `json:"StatusCallbackMethod,omitempty"`
}
