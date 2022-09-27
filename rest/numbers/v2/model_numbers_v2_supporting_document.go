/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Numbers
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

// NumbersV2SupportingDocument struct for NumbersV2SupportingDocument
type NumbersV2SupportingDocument struct {
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The string that you assigned to describe the resource
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The image type of the file
	MimeType *string `json:"mime_type,omitempty"`
	Status   *string `json:"status,omitempty"`
	// The failure reason of the Supporting Document Resource.
	FailureReason *string `json:"failure_reason,omitempty"`
	// The type of the Supporting Document
	Type *string `json:"type,omitempty"`
	// The set of parameters that compose the Supporting Documents resource
	Attributes *interface{} `json:"attributes,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The absolute URL of the Supporting Document resource
	Url *string `json:"url,omitempty"`
}
