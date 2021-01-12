/*
 * Twilio - Taskrouter
 *
 * This is the public Twilio REST API.
 *
 * API version: 5.15.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// InlineObject12 struct for InlineObject12
type InlineObject12 struct {
	// The SID of a valid Activity that will describe the Worker's initial state. See [Activities](https://www.twilio.com/docs/taskrouter/api/activity) for more information.
	ActivitySid string `json:"ActivitySid,omitempty"`
	// The JSON string that describes the Worker. For example: `{ \"email\": \"Bob@example.com\", \"phone\": \"+5095551234\" }`. This data is passed to the `assignment_callback_url` when TaskRouter assigns a Task to the Worker. Defaults to {}.
	Attributes string `json:"Attributes,omitempty"`
	// A descriptive string that you create to describe the Worker. It can be up to 64 characters long.
	FriendlyName string `json:"FriendlyName,omitempty"`
	// Whether to reject pending reservations.
	RejectPendingReservations bool `json:"RejectPendingReservations,omitempty"`
}
