/*
 * Twilio - Verify
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.8.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// CreateMessagingConfigurationRequest struct for CreateMessagingConfigurationRequest
type CreateMessagingConfigurationRequest struct {
	// The [ISO-3166-1](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code of the country this configuration will be applied to. If this is a global configuration, Country will take the value `all`.
	Country string `json:"Country"`
	// The SID of the [Messaging Service](https://www.twilio.com/docs/sms/services/api) to be used to send SMS to the country of this configuration.
	MessagingServiceSid string `json:"MessagingServiceSid"`
}
