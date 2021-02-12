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
// PushService the model 'PushService'
type PushService string

// List of push_service
const (
	PUSHSERVICE_GCM PushService = "gcm"
	PUSHSERVICE_APN PushService = "apn"
	PUSHSERVICE_FCM PushService = "fcm"
)
