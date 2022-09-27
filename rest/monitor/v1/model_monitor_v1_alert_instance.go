/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Monitor
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"time"
)

// MonitorV1AlertInstance struct for MonitorV1AlertInstance
type MonitorV1AlertInstance struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The text of the alert
	AlertText *string `json:"alert_text,omitempty"`
	// The API version used when the alert was generated
	ApiVersion *string `json:"api_version,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The date and time when the alert was generated specified in ISO 8601 format
	DateGenerated *time.Time `json:"date_generated,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The error code for the condition that generated the alert
	ErrorCode *string `json:"error_code,omitempty"`
	// The log level
	LogLevel *string `json:"log_level,omitempty"`
	// The URL of the page in our Error Dictionary with more information about the error condition
	MoreInfo *string `json:"more_info,omitempty"`
	// The method used by the request that generated the alert
	RequestMethod *string `json:"request_method,omitempty"`
	// The URL of the request that generated the alert
	RequestUrl *string `json:"request_url,omitempty"`
	// The variables passed in the request that generated the alert
	RequestVariables *string `json:"request_variables,omitempty"`
	// The SID of the resource for which the alert was generated
	ResourceSid *string `json:"resource_sid,omitempty"`
	// The response body of the request that generated the alert
	ResponseBody *string `json:"response_body,omitempty"`
	// The response headers of the request that generated the alert
	ResponseHeaders *string `json:"response_headers,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// The absolute URL of the Alert resource
	Url *string `json:"url,omitempty"`
	// The request headers of the request that generated the alert
	RequestHeaders *string `json:"request_headers,omitempty"`
	// The SID of the service or resource that generated the alert
	ServiceSid *string `json:"service_sid,omitempty"`
}
