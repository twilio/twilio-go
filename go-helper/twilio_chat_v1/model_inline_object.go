/*
 * Twilio - Chat
 *
 * This is the public Twilio REST API.
 *
 * API version: 5.15.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// InlineObject struct for InlineObject
type InlineObject struct {
	// [GCM only] The API key for the project that was obtained from the Google Developer console for your GCM Service application credential.
	ApiKey string `json:"ApiKey,omitempty"`
	// [APN only] The URL encoded representation of the certificate. For example,  `-----BEGIN CERTIFICATE----- MIIFnTCCBIWgAwIBAgIIAjy9H849+E8wDQYJKoZIhvcNAQEFBQAwgZYxCzAJBgNV.....A== -----END CERTIFICATE-----`
	Certificate string `json:"Certificate,omitempty"`
	// A descriptive string that you create to describe the new resource. It can be up to 64 characters long.
	FriendlyName string `json:"FriendlyName,omitempty"`
	// [APN only] The URL encoded representation of the private key. For example, `-----BEGIN RSA PRIVATE KEY----- MIIEpQIBAAKCAQEAuyf/lNrH9ck8DmNyo3fGgvCI1l9s+cmBY3WIz+cUDqmxiieR. -----END RSA PRIVATE KEY-----`
	PrivateKey string `json:"PrivateKey,omitempty"`
	// [APN only] Whether to send the credential to sandbox APNs. Can be `true` to send to sandbox APNs or `false` to send to production.
	Sandbox bool `json:"Sandbox,omitempty"`
	// [FCM only] The **Server key** of your project from the Firebase console, found under Settings / Cloud messaging.
	Secret string `json:"Secret,omitempty"`
	// The type of push-notification service the credential is for. Can be: `gcm`, `fcm`, or `apn`.
	Type string `json:"Type"`
}
