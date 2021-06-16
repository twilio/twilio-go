/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.17.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ApiV2010AccountCall struct for ApiV2010AccountCall
type ApiV2010AccountCall struct {
	// The SID of the Account that created this resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The annotation provided for the call
	Annotation *string `json:"annotation,omitempty"`
	// Either `human` or `machine` if this call was initiated with answering machine detection. Empty otherwise.
	AnsweredBy *string `json:"answered_by,omitempty"`
	// The API Version used to create the call
	ApiVersion *string `json:"api_version,omitempty"`
	// The caller's name if this call was an incoming call to a phone number with caller ID Lookup enabled. Otherwise, empty.
	CallerName *string `json:"caller_name,omitempty"`
	// The RFC 2822 date and time in GMT that this resource was created
	DateCreated *string `json:"date_created,omitempty"`
	// The RFC 2822 date and time in GMT that this resource was last updated
	DateUpdated *string `json:"date_updated,omitempty"`
	// A string describing the direction of the call. `inbound` for inbound calls, `outbound-api` for calls initiated via the REST API or `outbound-dial` for calls initiated by a `Dial` verb.
	Direction *string `json:"direction,omitempty"`
	// The length of the call in seconds.
	Duration *string `json:"duration,omitempty"`
	// The end time of the call. Null if the call did not complete successfully.
	EndTime *string `json:"end_time,omitempty"`
	// The forwarding phone number if this call was an incoming call forwarded from another number (depends on carrier supporting forwarding). Otherwise, empty.
	ForwardedFrom *string `json:"forwarded_from,omitempty"`
	// The phone number, SIP address or Client identifier that made this call. Phone numbers are in E.164 format (e.g., +16175551212). SIP addresses are formatted as `name@company.com`. Client identifiers are formatted `client:name`.
	From *string `json:"from,omitempty"`
	// The calling phone number, SIP address, or Client identifier formatted for display.
	FromFormatted *string `json:"from_formatted,omitempty"`
	// The Group SID associated with this call. If no Group is associated with the call, the field is empty.
	GroupSid *string `json:"group_sid,omitempty"`
	// The SID that identifies the call that created this leg.
	ParentCallSid *string `json:"parent_call_sid,omitempty"`
	// If the call was inbound, this is the SID of the IncomingPhoneNumber resource that received the call. If the call was outbound, it is the SID of the OutgoingCallerId resource from which the call was placed.
	PhoneNumberSid *string `json:"phone_number_sid,omitempty"`
	// The charge for this call, in the currency associated with the account. Populated after the call is completed. May not be immediately available.
	Price *string `json:"price,omitempty"`
	// The currency in which `Price` is measured.
	PriceUnit *string `json:"price_unit,omitempty"`
	// The wait time in milliseconds before the call is placed.
	QueueTime *string `json:"queue_time,omitempty"`
	// The unique string that identifies this resource
	Sid *string `json:"sid,omitempty"`
	// The start time of the call. Null if the call has not yet been dialed.
	StartTime *string `json:"start_time,omitempty"`
	// The status of this call.
	Status *string `json:"status,omitempty"`
	// A list of related subresources identified by their relative URIs
	SubresourceUris *map[string]interface{} `json:"subresource_uris,omitempty"`
	// The phone number, SIP address or Client identifier that received this call. Phone numbers are in E.164 format (e.g., +16175551212). SIP addresses are formatted as `name@company.com`. Client identifiers are formatted `client:name`.
	To *string `json:"to,omitempty"`
	// The phone number, SIP address or Client identifier that received this call. Formatted for display.
	ToFormatted *string `json:"to_formatted,omitempty"`
	// The (optional) unique identifier of the trunk resource that was used for this call.
	TrunkSid *string `json:"trunk_sid,omitempty"`
	// The URI of this resource, relative to `https://api.twilio.com`
	Uri *string `json:"uri,omitempty"`
}
