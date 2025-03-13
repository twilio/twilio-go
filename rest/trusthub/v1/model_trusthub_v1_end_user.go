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
	"encoding/json"
	"github.com/twilio/twilio-go/client"
	"time"
)

// TrusthubV1EndUser struct for TrusthubV1EndUser
type TrusthubV1EndUser struct {
	// The unique string created by Twilio to identify the End User resource.
	Sid *string `json:"sid,omitempty"`
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the End User resource.
	AccountSid *string `json:"account_sid,omitempty"`
	// The string that you assigned to describe the resource.
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The type of end user of the Bundle resource - can be `individual` or `business`.
	Type *string `json:"type,omitempty"`
	// The set of parameters that are the attributes of the End Users resource which are listed in the End User Types.
	Attributes *map[string]interface{} `json:"attributes,omitempty"`
	// The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The date and time in GMT when the resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The absolute URL of the End User resource.
	Url *string `json:"url,omitempty"`
}
