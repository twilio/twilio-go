/*
 * Twilio - Conversations
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.9.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// PushType the model 'PushType'
type PushType string

// List of push_type
const (
	PUSHTYPE_APN PushType = "apn"
	PUSHTYPE_GCM PushType = "gcm"
	PUSHTYPE_FCM PushType = "fcm"
)
