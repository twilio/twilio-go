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
// AutopilotV1AssistantModelBuildReadResponse struct for AutopilotV1AssistantModelBuildReadResponse
type AutopilotV1AssistantModelBuildReadResponse struct {
	Meta AutopilotV1AssistantReadResponseMeta `json:"meta,omitempty"`
	ModelBuilds []AutopilotV1AssistantModelBuild `json:"model_builds,omitempty"`
}
