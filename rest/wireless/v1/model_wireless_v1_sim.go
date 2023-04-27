/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Wireless
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
// WirelessV1Sim struct for WirelessV1Sim
type WirelessV1Sim struct {
		// The unique string that we created to identify the Sim resource.
	Sid *string `json:"sid,omitempty"`
		// An application-defined string that uniquely identifies the resource. It can be used in place of the resource's `sid` in the URL to address the resource.
	UniqueName *string `json:"unique_name,omitempty"`
		// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) to which the Sim resource belongs.
	AccountSid *string `json:"account_sid,omitempty"`
		// The SID of the [RatePlan resource](https://www.twilio.com/docs/wireless/api/rateplan-resource) to which the Sim resource is assigned.
	RatePlanSid *string `json:"rate_plan_sid,omitempty"`
		// The string that you assigned to describe the Sim resource.
	FriendlyName *string `json:"friendly_name,omitempty"`
		// The [ICCID](https://en.wikipedia.org/wiki/SIM_card#ICCID) associated with the SIM.
	Iccid *string `json:"iccid,omitempty"`
		// Deprecated.
	EId *string `json:"e_id,omitempty"`
	Status *string `json:"status,omitempty"`
	ResetStatus *string `json:"reset_status,omitempty"`
		// The URL we call using the `commands_callback_method` when the SIM originates a machine-to-machine [Command](https://www.twilio.com/docs/wireless/api/command-resource). Your server should respond with an HTTP status code in the 200 range; any response body will be ignored.
	CommandsCallbackUrl *string `json:"commands_callback_url,omitempty"`
		// The HTTP method we use to call `commands_callback_url`.  Can be: `POST` or `GET`. Default is `POST`.
	CommandsCallbackMethod *string `json:"commands_callback_method,omitempty"`
		// Deprecated.
	SmsFallbackMethod *string `json:"sms_fallback_method,omitempty"`
		// Deprecated.
	SmsFallbackUrl *string `json:"sms_fallback_url,omitempty"`
		// Deprecated.
	SmsMethod *string `json:"sms_method,omitempty"`
		// Deprecated.
	SmsUrl *string `json:"sms_url,omitempty"`
		// Deprecated. The HTTP method we use to call `voice_fallback_url`. Can be: `GET` or `POST`. Default is `POST`.
	VoiceFallbackMethod *string `json:"voice_fallback_method,omitempty"`
		// Deprecated. The URL we call using the `voice_fallback_method` when an error occurs while retrieving or executing the TwiML requested from `voice_url`.
	VoiceFallbackUrl *string `json:"voice_fallback_url,omitempty"`
		// Deprecated. The HTTP method we use to call `voice_url`. Can be: `GET` or `POST`. Default is `POST`.
	VoiceMethod *string `json:"voice_method,omitempty"`
		// Deprecated. The URL we call using the `voice_method` when the SIM-connected device makes a voice call.
	VoiceUrl *string `json:"voice_url,omitempty"`
		// The date and time in GMT when the resource was created specified in [ISO 8601](https://www.iso.org/iso-8601-date-and-time-format.html) format.
	DateCreated *time.Time `json:"date_created,omitempty"`
		// The date and time in GMT when the Sim resource was last updated specified in [ISO 8601](https://www.iso.org/iso-8601-date-and-time-format.html) format.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
		// The absolute URL of the resource.
	Url *string `json:"url,omitempty"`
		// The URLs of related subresources.
	Links *map[string]interface{} `json:"links,omitempty"`
		// Deprecated.
	IpAddress *string `json:"ip_address,omitempty"`
}


