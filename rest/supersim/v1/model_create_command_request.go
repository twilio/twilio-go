/*
 * Twilio - Supersim
 *
 * This is the public Twilio REST API.
 *
 * API version: 5.15.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// CreateCommandRequest struct for CreateCommandRequest
type CreateCommandRequest struct {
	// The HTTP method we should use to call `callback_url`. Can be: `GET` or `POST` and the default is POST.
	CallbackMethod string `json:"CallbackMethod,omitempty"`
	// The URL we should call using the `callback_method` after we have sent the command.
	CallbackUrl string `json:"CallbackUrl,omitempty"`
	// The message body of the command.
	Command string `json:"Command"`
	// The `sid` or `unique_name` of the [SIM](https://www.twilio.com/docs/wireless/api/sim-resource) to send the Command to.
	Sim string `json:"Sim"`
}
