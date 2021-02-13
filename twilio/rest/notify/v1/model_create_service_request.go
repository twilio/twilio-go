/*
 * Twilio - Notify
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.9.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// CreateServiceRequest struct for CreateServiceRequest
type CreateServiceRequest struct {
	// Deprecated.
	AlexaSkillId string `json:"AlexaSkillId,omitempty"`
	// The SID of the [Credential](https://www.twilio.com/docs/notify/api/credential-resource) to use for APN Bindings.
	ApnCredentialSid string `json:"ApnCredentialSid,omitempty"`
	// Deprecated.
	DefaultAlexaNotificationProtocolVersion string `json:"DefaultAlexaNotificationProtocolVersion,omitempty"`
	// The protocol version to use for sending APNS notifications. Can be overridden on a Binding by Binding basis when creating a [Binding](https://www.twilio.com/docs/notify/api/binding-resource) resource.
	DefaultApnNotificationProtocolVersion string `json:"DefaultApnNotificationProtocolVersion,omitempty"`
	// The protocol version to use for sending FCM notifications. Can be overridden on a Binding by Binding basis when creating a [Binding](https://www.twilio.com/docs/notify/api/binding-resource) resource.
	DefaultFcmNotificationProtocolVersion string `json:"DefaultFcmNotificationProtocolVersion,omitempty"`
	// The protocol version to use for sending GCM notifications. Can be overridden on a Binding by Binding basis when creating a [Binding](https://www.twilio.com/docs/notify/api/binding-resource) resource.
	DefaultGcmNotificationProtocolVersion string `json:"DefaultGcmNotificationProtocolVersion,omitempty"`
	// Callback configuration that enables delivery callbacks, default false
	DeliveryCallbackEnabled bool `json:"DeliveryCallbackEnabled,omitempty"`
	// URL to send delivery status callback.
	DeliveryCallbackUrl string `json:"DeliveryCallbackUrl,omitempty"`
	// Deprecated.
	FacebookMessengerPageId string `json:"FacebookMessengerPageId,omitempty"`
	// The SID of the [Credential](https://www.twilio.com/docs/notify/api/credential-resource) to use for FCM Bindings.
	FcmCredentialSid string `json:"FcmCredentialSid,omitempty"`
	// A descriptive string that you create to describe the resource. It can be up to 64 characters long.
	FriendlyName string `json:"FriendlyName,omitempty"`
	// The SID of the [Credential](https://www.twilio.com/docs/notify/api/credential-resource) to use for GCM Bindings.
	GcmCredentialSid string `json:"GcmCredentialSid,omitempty"`
	// Whether to log notifications. Can be: `true` or `false` and the default is `true`.
	LogEnabled bool `json:"LogEnabled,omitempty"`
	// The SID of the [Messaging Service](https://www.twilio.com/docs/sms/send-messages#messaging-services) to use for SMS Bindings. This parameter must be set in order to send SMS notifications.
	MessagingServiceSid string `json:"MessagingServiceSid,omitempty"`
}
