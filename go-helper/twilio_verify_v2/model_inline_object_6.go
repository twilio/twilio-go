/*
 * Twilio - Verify
 *
 * This is the public Twilio REST API.
 *
 * API version: 5.15.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// InlineObject6 struct for InlineObject6
type InlineObject6 struct {
	// The optional payload needed to verify the Factor for the first time. E.g. for a TOTP, the numeric code.
	AuthPayload string `json:"AuthPayload,omitempty"`
	// For APN, the device token. For FCM the registration token. It used to send the push notifications. Required when `factor_type` is `push`
	ConfigNotificationToken string `json:"Config.NotificationToken,omitempty"`
	// The Verify Push SDK version used to configure the factor
	ConfigSdkVersion string `json:"Config.SdkVersion,omitempty"`
	// The new friendly name of this Factor
	FriendlyName string `json:"FriendlyName,omitempty"`
}
