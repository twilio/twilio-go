/*
 * Twilio - Taskrouter
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.9.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// QueueOrder the model 'QueueOrder'
type QueueOrder string

// List of queue_order
const (
	QUEUEORDER_FIFO QueueOrder = "FIFO"
	QUEUEORDER_LIFO QueueOrder = "LIFO"
)
