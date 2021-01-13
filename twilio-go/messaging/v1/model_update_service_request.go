/*
 * Twilio - Messaging
 *
 * This is the public Twilio REST API.
 *
 * API version: 5.15.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// UpdateServiceRequest struct for UpdateServiceRequest
type UpdateServiceRequest struct {
	// Whether to enable [Area Code Geomatch](https://www.twilio.com/docs/sms/services#area-code-geomatch) on the Service Instance.
	AreaCodeGeomatch bool `json:"AreaCodeGeomatch,omitempty"`
	// The HTTP method we should use to call `fallback_url`. Can be: `GET` or `POST`.
	FallbackMethod string `json:"FallbackMethod,omitempty"`
	// Whether to enable [Fallback to Long Code](https://www.twilio.com/docs/sms/services#fallback-to-long-code) for messages sent through the Service instance.
	FallbackToLongCode bool `json:"FallbackToLongCode,omitempty"`
	// The URL that we should call using `fallback_method` if an error occurs while retrieving or executing the TwiML from the Inbound Request URL.
	FallbackUrl string `json:"FallbackUrl,omitempty"`
	// A descriptive string that you create to describe the resource. It can be up to 64 characters long.
	FriendlyName string `json:"FriendlyName,omitempty"`
	// The HTTP method we should use to call `inbound_request_url`. Can be `GET` or `POST` and the default is `POST`.
	InboundMethod string `json:"InboundMethod,omitempty"`
	// The URL we should call using `inbound_method` when a message is received by any phone number or short code in the Service. When this property is `null`, receiving inbound messages is disabled.
	InboundRequestUrl string `json:"InboundRequestUrl,omitempty"`
	// Whether to enable the [MMS Converter](https://www.twilio.com/docs/sms/services#mms-converter) for messages sent through the Service instance.
	MmsConverter bool `json:"MmsConverter,omitempty"`
	// Reserved.
	ScanMessageContent string `json:"ScanMessageContent,omitempty"`
	// Whether to enable [Smart Encoding](https://www.twilio.com/docs/sms/services#smart-encoding) for messages sent through the Service instance.
	SmartEncoding bool `json:"SmartEncoding,omitempty"`
	// The URL we should call to [pass status updates](https://www.twilio.com/docs/sms/api/message-resource#message-status-values) about message delivery.
	StatusCallback string `json:"StatusCallback,omitempty"`
	// Whether to enable [Sticky Sender](https://www.twilio.com/docs/sms/services#sticky-sender) on the Service instance.
	StickySender bool `json:"StickySender,omitempty"`
	// Reserved.
	SynchronousValidation bool `json:"SynchronousValidation,omitempty"`
	// How long, in seconds, messages sent from the Service are valid. Can be an integer from `1` to `14,400`.
	ValidityPeriod int32 `json:"ValidityPeriod,omitempty"`
}
