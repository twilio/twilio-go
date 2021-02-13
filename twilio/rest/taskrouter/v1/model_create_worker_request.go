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

// CreateWorkerRequest struct for CreateWorkerRequest
type CreateWorkerRequest struct {
	// The SID of a valid Activity that will describe the new Worker's initial state. See [Activities](https://www.twilio.com/docs/taskrouter/api/activity) for more information. If not provided, the new Worker's initial state is the `default_activity_sid` configured on the Workspace.
	ActivitySid string `json:"ActivitySid,omitempty"`
	// A valid JSON string that describes the new Worker. For example: `{ \"email\": \"Bob@example.com\", \"phone\": \"+5095551234\" }`. This data is passed to the `assignment_callback_url` when TaskRouter assigns a Task to the Worker. Defaults to {}.
	Attributes string `json:"Attributes,omitempty"`
	// A descriptive string that you create to describe the new Worker. It can be up to 64 characters long.
	FriendlyName string `json:"FriendlyName"`
}
