/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Messaging
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

// MessagingV1ShortCode struct for MessagingV1ShortCode
type MessagingV1ShortCode struct {
	// The unique string that we created to identify the ShortCode resource.
	Sid *string `json:"sid,omitempty"`
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the ShortCode resource.
	AccountSid *string `json:"account_sid,omitempty"`
	// The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) the resource is associated with.
	ServiceSid *string `json:"service_sid,omitempty"`
	// The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The date and time in GMT when the resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The [E.164](https://www.twilio.com/docs/glossary/what-e164) format of the short code.
	ShortCode *string `json:"short_code,omitempty"`
	// The 2-character [ISO Country Code](https://www.iso.org/iso-3166-country-codes.html) of the number.
	CountryCode *string `json:"country_code,omitempty"`
	// An array of values that describe whether the number can receive calls or messages. Can be: `SMS` and `MMS`.
	Capabilities *[]string `json:"capabilities,omitempty"`
	// The absolute URL of the ShortCode resource.
	Url *string `json:"url,omitempty"`
}
