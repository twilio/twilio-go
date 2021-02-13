/*
 * Twilio - Preview
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.9.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ListDeploymentResponse struct for ListDeploymentResponse
type ListDeploymentResponse struct {
	Deployments []PreviewDeployedDevicesFleetDeployment `json:"Deployments,omitempty"`
	Meta        ListDayResponseMeta                     `json:"Meta,omitempty"`
}
