/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Serverless
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"encoding/json"
	"github.com/twilio/twilio-go/client"
	"time"
)

// ServerlessV1Deployment struct for ServerlessV1Deployment
type ServerlessV1Deployment struct {
	// The unique string that we created to identify the Deployment resource.
	Sid *string `json:"sid,omitempty"`
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Deployment resource.
	AccountSid *string `json:"account_sid,omitempty"`
	// The SID of the Service that the Deployment resource is associated with.
	ServiceSid *string `json:"service_sid,omitempty"`
	// The SID of the Environment for the Deployment.
	EnvironmentSid *string `json:"environment_sid,omitempty"`
	// The SID of the Build for the deployment.
	BuildSid *string `json:"build_sid,omitempty"`
	// The date and time in GMT when the Deployment resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The date and time in GMT when the Deployment resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The absolute URL of the Deployment resource.
	Url *string `json:"url,omitempty"`
}
