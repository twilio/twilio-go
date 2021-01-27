/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.0.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// UpdateUsageTriggerRequest struct for UpdateUsageTriggerRequest
type UpdateUsageTriggerRequest struct {
	// The HTTP method we should use to call `callback_url`. Can be: `GET` or `POST` and the default is `POST`.
	CallbackMethod string `json:"CallbackMethod,omitempty"`
	// The URL we should call using `callback_method` when the trigger fires.
	CallbackUrl string `json:"CallbackUrl,omitempty"`
	// A descriptive string that you create to describe the resource. It can be up to 64 characters long.
	FriendlyName string `json:"FriendlyName,omitempty"`
}
