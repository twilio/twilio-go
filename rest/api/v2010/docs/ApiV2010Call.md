# ApiV2010Call

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created this resource |
**AnsweredBy** | Pointer to **string** | Either `human` or `machine` if this call was initiated with answering machine detection. Empty otherwise. |
**ApiVersion** | Pointer to **string** | The API Version used to create the call |
**CallerName** | Pointer to **string** | The caller's name if this call was an incoming call to a phone number with caller ID Lookup enabled. Otherwise, empty. |
**DateCreated** | Pointer to **string** | The RFC 2822 date and time in GMT that this resource was created |
**DateUpdated** | Pointer to **string** | The RFC 2822 date and time in GMT that this resource was last updated |
**Direction** | Pointer to **string** | A string describing the direction of the call. `inbound` for inbound calls, `outbound-api` for calls initiated via the REST API or `outbound-dial` for calls initiated by a `Dial` verb. |
**Duration** | Pointer to **string** | The length of the call in seconds. |
**EndTime** | Pointer to **string** | The end time of the call. Null if the call did not complete successfully. |
**ForwardedFrom** | Pointer to **string** | The forwarding phone number if this call was an incoming call forwarded from another number (depends on carrier supporting forwarding). Otherwise, empty. |
**From** | Pointer to **string** | The phone number, SIP address or Client identifier that made this call. Phone numbers are in E.164 format (e.g., +16175551212). SIP addresses are formatted as `name@company.com`. Client identifiers are formatted `client:name`. |
**FromFormatted** | Pointer to **string** | The calling phone number, SIP address, or Client identifier formatted for display. |
**GroupSid** | Pointer to **string** | The Group SID associated with this call. If no Group is associated with the call, the field is empty. |
**ParentCallSid** | Pointer to **string** | The SID that identifies the call that created this leg. |
**PhoneNumberSid** | Pointer to **string** | If the call was inbound, this is the SID of the IncomingPhoneNumber resource that received the call. If the call was outbound, it is the SID of the OutgoingCallerId resource from which the call was placed. |
**Price** | Pointer to **string** | The charge for this call, in the currency associated with the account. Populated after the call is completed. May not be immediately available. |
**PriceUnit** | Pointer to **string** | The currency in which `Price` is measured. |
**QueueTime** | Pointer to **string** | The wait time in milliseconds before the call is placed. |
**Sid** | Pointer to **string** | The unique string that identifies this resource |
**StartTime** | Pointer to **string** | The start time of the call. Null if the call has not yet been dialed. |
**Status** | Pointer to **string** | The status of this call. |
**SubresourceUris** | Pointer to **map[string]interface{}** | A list of related subresources identified by their relative URIs |
**To** | Pointer to **string** | The phone number, SIP address or Client identifier that received this call. Phone numbers are in E.164 format (e.g., +16175551212). SIP addresses are formatted as `name@company.com`. Client identifiers are formatted `client:name`. |
**ToFormatted** | Pointer to **string** | The phone number, SIP address or Client identifier that received this call. Formatted for display. |
**TrunkSid** | Pointer to **string** | The (optional) unique identifier of the trunk resource that was used for this call. |
**Uri** | Pointer to **string** | The URI of this resource, relative to `https://api.twilio.com` |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


