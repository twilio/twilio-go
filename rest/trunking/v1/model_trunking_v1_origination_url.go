/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Trunking
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

// TrunkingV1OriginationUrl struct for TrunkingV1OriginationUrl
type TrunkingV1OriginationUrl struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the OriginationUrl resource.
	AccountSid *string `json:"account_sid,omitempty"`
	// The unique string that we created to identify the OriginationUrl resource.
	Sid *string `json:"sid,omitempty"`
	// The SID of the Trunk that owns the Origination URL.
	TrunkSid *string `json:"trunk_sid,omitempty"`
	// The value that determines the relative share of the load the URI should receive compared to other URIs with the same priority. Can be an integer from 1 to 65535, inclusive, and the default is 10. URLs with higher values receive more load than those with lower ones with the same priority.
	Weight int `json:"weight,omitempty"`
	// Whether the URL is enabled. The default is `true`.
	Enabled *bool `json:"enabled,omitempty"`
	// The SIP address you want Twilio to route your Origination calls to. This must be a `sip:` schema.
	SipUrl *string `json:"sip_url,omitempty"`
	// The string that you assigned to describe the resource.
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The relative importance of the URI. Can be an integer from 0 to 65535, inclusive, and the default is 10. The lowest number represents the most important URI.
	Priority int `json:"priority,omitempty"`
	// The date and time in GMT when the resource was created specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The date and time in GMT when the resource was last updated specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The absolute URL of the resource.
	Url *string `json:"url,omitempty"`
}
