/*
 * Twilio - Serverless
 *
 * This is the public Twilio REST API.
 *
 * API version: 5.15.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
import (
	"time"
)
// ServerlessV1ServiceEnvironmentLog struct for ServerlessV1ServiceEnvironmentLog
type ServerlessV1ServiceEnvironmentLog struct {
	AccountSid string `json:"AccountSid,omitempty"`
	BuildSid string `json:"BuildSid,omitempty"`
	DateCreated time.Time `json:"DateCreated,omitempty"`
	DeploymentSid string `json:"DeploymentSid,omitempty"`
	EnvironmentSid string `json:"EnvironmentSid,omitempty"`
	FunctionSid string `json:"FunctionSid,omitempty"`
	Level string `json:"Level,omitempty"`
	Message string `json:"Message,omitempty"`
	RequestSid string `json:"RequestSid,omitempty"`
	ServiceSid string `json:"ServiceSid,omitempty"`
	Sid string `json:"Sid,omitempty"`
	Url string `json:"Url,omitempty"`
}
