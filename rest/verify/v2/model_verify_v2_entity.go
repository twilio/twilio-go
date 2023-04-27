/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Verify
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
// VerifyV2Entity struct for VerifyV2Entity
type VerifyV2Entity struct {
		// A 34 character string that uniquely identifies this Entity.
	Sid *string `json:"sid,omitempty"`
		// The unique external identifier for the Entity of the Service. This identifier should be immutable, not PII, length between 8 and 64 characters, and generated by your external system, such as your user's UUID, GUID, or SID. It can only contain dash (-) separated alphanumeric characters.
	Identity *string `json:"identity,omitempty"`
		// The unique SID identifier of the Account.
	AccountSid *string `json:"account_sid,omitempty"`
		// The unique SID identifier of the Service.
	ServiceSid *string `json:"service_sid,omitempty"`
		// The date that this Entity was created, given in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateCreated *time.Time `json:"date_created,omitempty"`
		// The date that this Entity was updated, given in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
		// The URL of this resource.
	Url *string `json:"url,omitempty"`
		// Contains a dictionary of URL links to nested resources of this Entity.
	Links *map[string]interface{} `json:"links,omitempty"`
}


