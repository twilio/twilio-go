/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Supersim
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

// SupersimV1EsimProfile struct for SupersimV1EsimProfile
type SupersimV1EsimProfile struct {
	// The unique string that we created to identify the eSIM Profile resource.
	Sid *string `json:"sid,omitempty"`
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) to which the eSIM Profile resource belongs.
	AccountSid *string `json:"account_sid,omitempty"`
	// The [ICCID](https://en.wikipedia.org/wiki/Subscriber_identity_module#ICCID) associated with the Sim resource.
	Iccid *string `json:"iccid,omitempty"`
	// The SID of the [Sim](https://www.twilio.com/docs/iot/supersim/api/sim-resource) resource that this eSIM Profile controls.
	SimSid *string `json:"sim_sid,omitempty"`
	Status *string `json:"status,omitempty"`
	// Identifier of the eUICC that can claim the eSIM Profile.
	Eid *string `json:"eid,omitempty"`
	// Address of the SM-DP+ server from which the Profile will be downloaded. The URL will appear once the eSIM Profile reaches the status `available`.
	SmdpPlusAddress *string `json:"smdp_plus_address,omitempty"`
	// Unique identifier of the eSIM profile that can be used to identify and download the eSIM profile from the SM-DP+ server. Populated if `generate_matching_id` is set to `true` when creating the eSIM profile reservation.
	MatchingId *string `json:"matching_id,omitempty"`
	// Combined machine-readable activation code for acquiring an eSIM Profile with the Activation Code download method. Can be used in a QR code to download an eSIM profile.
	ActivationCode *string `json:"activation_code,omitempty"`
	// Code indicating the failure if the download of the SIM Profile failed and the eSIM Profile is in `failed` state.
	ErrorCode *string `json:"error_code,omitempty"`
	// Error message describing the failure if the download of the SIM Profile failed and the eSIM Profile is in `failed` state.
	ErrorMessage *string `json:"error_message,omitempty"`
	// The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The date and time in GMT when the resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The absolute URL of the eSIM Profile resource.
	Url *string `json:"url,omitempty"`
}
