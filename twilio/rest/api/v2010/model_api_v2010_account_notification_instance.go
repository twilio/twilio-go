/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.12.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ApiV2010AccountNotificationInstance struct for ApiV2010AccountNotificationInstance
type ApiV2010AccountNotificationInstance struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"AccountSid,omitempty"`
	// The API version used to generate the notification
	ApiVersion *string `json:"ApiVersion,omitempty"`
	// The SID of the Call the resource is associated with
	CallSid *string `json:"CallSid,omitempty"`
	// The RFC 2822 date and time in GMT that the resource was created
	DateCreated *string `json:"DateCreated,omitempty"`
	// The RFC 2822 date and time in GMT that the resource was last updated
	DateUpdated *string `json:"DateUpdated,omitempty"`
	// A unique error code corresponding to the notification
	ErrorCode *string `json:"ErrorCode,omitempty"`
	// An integer log level
	Log *string `json:"Log,omitempty"`
	// The date the notification was generated
	MessageDate *string `json:"MessageDate,omitempty"`
	// The text of the notification
	MessageText *string `json:"MessageText,omitempty"`
	// A URL for more information about the error code
	MoreInfo *string `json:"MoreInfo,omitempty"`
	// HTTP method used with the request url
	RequestMethod *string `json:"RequestMethod,omitempty"`
	// URL of the resource that generated the notification
	RequestUrl *string `json:"RequestUrl,omitempty"`
	// Twilio-generated HTTP variables sent to the server
	RequestVariables *string `json:"RequestVariables,omitempty"`
	// The HTTP body returned by your server
	ResponseBody *string `json:"ResponseBody,omitempty"`
	// The HTTP headers returned by your server
	ResponseHeaders *string `json:"ResponseHeaders,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"Sid,omitempty"`
	// The URI of the resource, relative to `https://api.twilio.com`
	Uri *string `json:"Uri,omitempty"`
}
