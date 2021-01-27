/*
 * Twilio - Verify
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.0.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// UpdateMessagingConfigurationRequest struct for UpdateMessagingConfigurationRequest
type UpdateMessagingConfigurationRequest struct {
	// The SID of the [Messaging Service](https://www.twilio.com/docs/sms/services/api) to be used to send SMS to the country of this configuration.
	MessagingServiceSid string `json:"MessagingServiceSid"`
}
