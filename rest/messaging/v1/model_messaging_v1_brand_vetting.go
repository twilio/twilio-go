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
// MessagingV1BrandVetting struct for MessagingV1BrandVetting
type MessagingV1BrandVetting struct {
		// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the vetting record.
	AccountSid *string `json:"account_sid,omitempty"`
		// The unique string to identify Brand Registration.
	BrandSid *string `json:"brand_sid,omitempty"`
		// The Twilio SID of the third-party vetting record.
	BrandVettingSid *string `json:"brand_vetting_sid,omitempty"`
		// The date and time in GMT when the resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
		// The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateCreated *time.Time `json:"date_created,omitempty"`
		// The unique identifier of the vetting from the third-party provider.
	VettingId *string `json:"vetting_id,omitempty"`
		// The type of vetting that has been conducted. One of “STANDARD” (Aegis) or “POLITICAL” (Campaign Verify).
	VettingClass *string `json:"vetting_class,omitempty"`
		// The status of the import vetting attempt. One of “PENDING,” “SUCCESS,” or “FAILED”.
	VettingStatus *string `json:"vetting_status,omitempty"`
	VettingProvider *string `json:"vetting_provider,omitempty"`
		// The absolute URL of the Brand Vetting resource.
	Url *string `json:"url,omitempty"`
}


