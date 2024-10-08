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

// NumbersV1PortingPortInPhoneNumber struct for NumbersV1PortingPortInPhoneNumber
type NumbersV1PortingPortInPhoneNumber struct {
	// The unique identifier for the port in request that this phone number is associated with.
	PortInRequestSid *string `json:"port_in_request_sid,omitempty"`
	// The unique identifier for this phone number associated with this port in request.
	PhoneNumberSid *string `json:"phone_number_sid,omitempty"`
	// URL reference for this resource.
	Url *string `json:"url,omitempty"`
	// Account Sid or subaccount where the phone number(s) will be Ported.
	AccountSid *string `json:"account_sid,omitempty"`
	// The number type of the phone number. This can be: toll-free, local, mobile or unknown. This field may be null if the number is not portable or if the portability for a number has not yet been evaluated.
	PhoneNumberType *string `json:"phone_number_type,omitempty"`
	// The timestamp for when this port in phone number was created.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The ISO country code that this number is associated with. This field may be null if the number is not portable or if the portability for a number has not yet been evaluated.
	Country *string `json:"country,omitempty"`
	// Indicates if the phone number is missing required fields such as a PIN or account number. This field may be null if the number is not portable or if the portability for a number has not yet been evaluated.
	MissingRequiredFields *bool `json:"missing_required_fields,omitempty"`
	// Timestamp indicating when the Port In Phone Number resource was last modified.
	LastUpdated *time.Time `json:"last_updated,omitempty"`
	// Phone number to be ported. This will be in the E164 Format.
	PhoneNumber *string `json:"phone_number,omitempty"`
	// If the number is portable by Twilio or not. This field may be null if the number portability has not yet been evaluated. If a number is not portable reference the `not_portability_reason_code` and `not_portability_reason` fields for more details
	Portable *bool `json:"portable,omitempty"`
	// The not portability reason code description. This field may be null if the number is portable or if the portability for a number has not yet been evaluated.
	NotPortabilityReason *string `json:"not_portability_reason,omitempty"`
	// The not portability reason code. This field may be null if the number is portable or if the portability for a number has not yet been evaluated.
	NotPortabilityReasonCode *int `json:"not_portability_reason_code,omitempty"`
	// The status of the port in phone number.
	PortInPhoneNumberStatus *string `json:"port_in_phone_number_status,omitempty"`
	// The pin required by the losing carrier to do the port out.
	PortOutPin *int `json:"port_out_pin,omitempty"`
	// The description of the rejection reason provided by the losing carrier. This field may be null if the number has not been rejected by the losing carrier.
	RejectionReason *string `json:"rejection_reason,omitempty"`
	// The code for the rejection reason provided by the losing carrier. This field may be null if the number has not been rejected by the losing carrier.
	RejectionReasonCode *int `json:"rejection_reason_code,omitempty"`
	// The timestamp the phone number will be ported. This will only be set once a port date has been confirmed. Not all carriers can guarantee a specific time on the port date. Twilio will try its best to get the port completed by this time on the port date. Please subscribe to webhooks for confirmation on when a port has actually been completed.
	PortDate *time.Time `json:"port_date,omitempty"`
}
