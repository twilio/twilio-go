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

// DeliveryStatus the model 'DeliveryStatus'
type DeliveryStatus string

// List of delivery_status
const (
	DELIVERYSTATUS_READ        DeliveryStatus = "read"
	DELIVERYSTATUS_FAILED      DeliveryStatus = "failed"
	DELIVERYSTATUS_DELIVERED   DeliveryStatus = "delivered"
	DELIVERYSTATUS_UNDELIVERED DeliveryStatus = "undelivered"
	DELIVERYSTATUS_SENT        DeliveryStatus = "sent"
)
