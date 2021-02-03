/*
 * Twilio - Conversations
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.8.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// PushType the model 'PushType'
type PushType string

// List of push_type
const (
	APN PushType = "apn"
	GCM PushType = "gcm"
	FCM PushType = "fcm"
)
