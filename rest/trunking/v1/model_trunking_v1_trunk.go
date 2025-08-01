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
	"time"
)

// TrunkingV1Trunk struct for TrunkingV1Trunk
type TrunkingV1Trunk struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Trunk resource.
	AccountSid *string `json:"account_sid,omitempty"`
	// The unique address you reserve on Twilio to which you route your SIP traffic. Domain names can contain letters, digits, and `-` and must end with `pstn.twilio.com`. See [Termination Settings](https://www.twilio.com/docs/sip-trunking#termination) for more information.
	DomainName *string `json:"domain_name,omitempty"`
	// The HTTP method we use to call the `disaster_recovery_url`. Can be: `GET` or `POST`.
	DisasterRecoveryMethod *string `json:"disaster_recovery_method,omitempty"`
	// The URL we call using the `disaster_recovery_method` if an error occurs while sending SIP traffic towards the configured Origination URL. We retrieve TwiML from this URL and execute the instructions like any other normal TwiML call. See [Disaster Recovery](https://www.twilio.com/docs/sip-trunking#disaster-recovery) for more information.
	DisasterRecoveryUrl *string `json:"disaster_recovery_url,omitempty"`
	// The string that you assigned to describe the resource.
	FriendlyName *string `json:"friendly_name,omitempty"`
	// Whether Secure Trunking is enabled for the trunk. If enabled, all calls going through the trunk will be secure using SRTP for media and TLS for signaling. If disabled, then RTP will be used for media. See [Secure Trunking](https://www.twilio.com/docs/sip-trunking#securetrunking) for more information.
	Secure *bool `json:"secure,omitempty"`
	// The recording settings for the trunk. Can be: `do-not-record`, `record-from-ringing`, `record-from-answer`. If set to `record-from-ringing` or `record-from-answer`, all calls going through the trunk will be recorded. The only way to change recording parameters is on a sub-resource of a Trunk after it has been created. e.g.`/Trunks/[Trunk_SID]/Recording -XPOST -d'Mode=record-from-answer'`. See [Recording](https://www.twilio.com/docs/sip-trunking#recording) for more information.
	Recording        *interface{} `json:"recording,omitempty"`
	TransferMode     *string      `json:"transfer_mode,omitempty"`
	TransferCallerId *string      `json:"transfer_caller_id,omitempty"`
	// Whether Caller ID Name (CNAM) lookup is enabled for the trunk. If enabled, all inbound calls to the SIP Trunk from the United States and Canada automatically perform a CNAM Lookup and display Caller ID data on your phone. See [CNAM Lookups](https://www.twilio.com/docs/sip-trunking#CNAM) for more information.
	CnamLookupEnabled *bool `json:"cnam_lookup_enabled,omitempty"`
	// The types of authentication mapped to the domain. Can be: `IP_ACL` and `CREDENTIAL_LIST`. If both are mapped, the values are returned in a comma delimited list. If empty, the domain will not receive any traffic.
	AuthType *string `json:"auth_type,omitempty"`
	// Whether Symmetric RTP is enabled for the trunk. When Symmetric RTP is disabled, Twilio will send RTP to the destination negotiated in the SDP. Disabling Symmetric RTP is considered to be more secure and therefore recommended. See [Symmetric RTP](https://www.twilio.com/docs/sip-trunking#symmetric-rtp) for more information.
	SymmetricRtpEnabled *bool `json:"symmetric_rtp_enabled,omitempty"`
	// Reserved.
	AuthTypeSet *[]string `json:"auth_type_set,omitempty"`
	// The date and time in GMT when the resource was created specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The date and time in GMT when the resource was last updated specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The unique string that we created to identify the Trunk resource.
	Sid *string `json:"sid,omitempty"`
	// The absolute URL of the resource.
	Url *string `json:"url,omitempty"`
	// The URLs of related resources.
	Links *map[string]interface{} `json:"links,omitempty"`
}
