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

// NumbersV1PortingPortability struct for NumbersV1PortingPortability
type NumbersV1PortingPortability struct {
	// The phone number which portability is to be checked. Phone numbers are in E.164 format (e.g. +16175551212).
	PhoneNumber *string `json:"phone_number,omitempty"`
	// Account Sid that the phone number belongs to in Twilio. This is only returned for phone numbers that already exist in Twilio’s inventory and belong to your account or sub account.
	AccountSid *string `json:"account_sid,omitempty"`
	// Boolean flag indicates if the phone number can be ported into Twilio through the Porting API or not.
	Portable *bool `json:"portable,omitempty"`
	// Indicates if the port in process will require a personal identification number (PIN) and an account number for this phone number. If this is true you will be required to submit both a PIN and account number from the losing carrier for this number when opening a port in request. These fields will be required in order to complete the port in process to Twilio.
	PinAndAccountNumberRequired *bool `json:"pin_and_account_number_required,omitempty"`
	// Reason why the phone number cannot be ported into Twilio, `null` otherwise.
	NotPortableReason *string `json:"not_portable_reason,omitempty"`
	// The Portability Reason Code for the phone number if it cannot be ported into Twilio, `null` otherwise.
	NotPortableReasonCode *int    `json:"not_portable_reason_code,omitempty"`
	NumberType            *string `json:"number_type,omitempty"`
	// Country the phone number belongs to.
	Country *string `json:"country,omitempty"`
	// This is the url of the request that you're trying to reach out to locate the resource.
	Url *string `json:"url,omitempty"`
}
