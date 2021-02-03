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
// ChannelStatus the model 'ChannelStatus'
type ChannelStatus string

// List of channel_status
const (
	JOINED ChannelStatus = "joined"
	INVITED ChannelStatus = "invited"
	NOT_PARTICIPATING ChannelStatus = "not_participating"
)
