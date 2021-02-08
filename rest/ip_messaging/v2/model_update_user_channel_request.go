/*
 * Twilio - Ip_messaging
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.8.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
import (
	"time"
)
// UpdateUserChannelRequest struct for UpdateUserChannelRequest
type UpdateUserChannelRequest struct {
	LastConsumedMessageIndex *int32 `json:"LastConsumedMessageIndex,omitempty"`
	LastConsumptionTimestamp time.Time `json:"LastConsumptionTimestamp,omitempty"`
	NotificationLevel string `json:"NotificationLevel,omitempty"`
}
