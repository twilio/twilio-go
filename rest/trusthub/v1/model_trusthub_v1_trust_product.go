/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Trusthub
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

// TrusthubV1TrustProduct struct for TrusthubV1TrustProduct
type TrusthubV1TrustProduct struct {
	// The unique string that we created to identify the Trust Product resource.
	Sid *string `json:"sid,omitempty"`
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Trust Product resource.
	AccountSid *string `json:"account_sid,omitempty"`
	// The unique string of the policy that is associated with the Trust Product resource.
	PolicySid *string `json:"policy_sid,omitempty"`
	// The string that you assigned to describe the resource.
	FriendlyName *string `json:"friendly_name,omitempty"`
	Status       *string `json:"status,omitempty"`
	// The date and time in GMT in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format until which the resource will be valid.
	ValidUntil *time.Time `json:"valid_until,omitempty"`
	// The email address that will receive updates when the Trust Product resource changes status.
	Email *string `json:"email,omitempty"`
	// The URL we call to inform your application of status changes.
	StatusCallback *string `json:"status_callback,omitempty"`
	// The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The date and time in GMT when the resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The absolute URL of the Trust Product resource.
	Url *string `json:"url,omitempty"`
	// The URLs of the Assigned Items of the Trust Product resource.
	Links *map[string]interface{} `json:"links,omitempty"`
	// The error codes associated with the rejection of the Trust Product.
	Errors *[]map[string]interface{} `json:"errors,omitempty"`
}
