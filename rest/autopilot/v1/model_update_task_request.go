/*
 * Twilio - Autopilot
 *
 * This is the public Twilio REST API.
 *
 * API version: 5.15.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// UpdateTaskRequest struct for UpdateTaskRequest
type UpdateTaskRequest struct {
	// The JSON string that specifies the [actions](https://www.twilio.com/docs/autopilot/actions) that instruct the Assistant on how to perform the task.
	Actions map[string]interface{} `json:"Actions,omitempty"`
	// The URL from which the Assistant can fetch actions.
	ActionsUrl string `json:"ActionsUrl,omitempty"`
	// A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long.
	FriendlyName string `json:"FriendlyName,omitempty"`
	// An application-defined string that uniquely identifies the resource. This value must be 64 characters or less in length and be unique. It can be used as an alternative to the `sid` in the URL path to address the resource.
	UniqueName string `json:"UniqueName,omitempty"`
}
