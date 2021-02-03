/*
 * Twilio - Wireless
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.8.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// CreateCommandRequest struct for CreateCommandRequest
type CreateCommandRequest struct {
	// The HTTP method we use to call `callback_url`. Can be: `POST` or `GET`, and the default is `POST`.
	CallbackMethod string `json:"CallbackMethod,omitempty"`
	// The URL we call using the `callback_url` when the Command has finished sending, whether the command was delivered or it failed.
	CallbackUrl string `json:"CallbackUrl,omitempty"`
	// The message body of the Command. Can be plain text in text mode or a Base64 encoded byte string in binary mode.
	Command string `json:"Command"`
	// The mode to use when sending the SMS message. Can be: `text` or `binary`. The default SMS mode is `text`.
	CommandMode string `json:"CommandMode,omitempty"`
	// Whether to request delivery receipt from the recipient. For Commands that request delivery receipt, the Command state transitions to 'delivered' once the server has received a delivery receipt from the device. The default value is `true`.
	DeliveryReceiptRequested bool `json:"DeliveryReceiptRequested,omitempty"`
	// Whether to include the SID of the command in the message body. Can be: `none`, `start`, or `end`, and the default behavior is `none`. When sending a Command to a SIM in text mode, we can automatically include the SID of the Command in the message body, which could be used to ensure that the device does not process the same Command more than once.  A value of `start` will prepend the message with the Command SID, and `end` will append it to the end, separating the Command SID from the message body with a space. The length of the Command SID is included in the 160 character limit so the SMS body must be 128 characters or less before the Command SID is included.
	IncludeSid string `json:"IncludeSid,omitempty"`
	// The `sid` or `unique_name` of the [SIM](https://www.twilio.com/docs/wireless/api/sim-resource) to send the Command to.
	Sim string `json:"Sim,omitempty"`
}
