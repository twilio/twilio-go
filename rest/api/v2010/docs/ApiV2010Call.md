# ApiV2010Call

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | **string** | The unique string that identifies this resource |[optional] 
**DateCreated** | **string** | The RFC 2822 date and time in GMT that this resource was created |[optional] 
**DateUpdated** | **string** | The RFC 2822 date and time in GMT that this resource was last updated |[optional] 
**ParentCallSid** | **string** | The SID that identifies the call that created this leg. |[optional] 
**AccountSid** | **string** | The SID of the Account that created this resource |[optional] 
**To** | **string** | The phone number, SIP address or Client identifier that received this call. Phone numbers are in E.164 format (e.g., +16175551212). SIP addresses are formatted as `name@company.com`. Client identifiers are formatted `client:name`. |[optional] 
**ToFormatted** | **string** | The phone number, SIP address or Client identifier that received this call. Formatted for display. |[optional] 
**From** | **string** | The phone number, SIP address or Client identifier that made this call. Phone numbers are in E.164 format (e.g., +16175551212). SIP addresses are formatted as `name@company.com`. Client identifiers are formatted `client:name`. |[optional] 
**FromFormatted** | **string** | The calling phone number, SIP address, or Client identifier formatted for display. |[optional] 
**PhoneNumberSid** | **string** | If the call was inbound, this is the SID of the IncomingPhoneNumber resource that received the call. If the call was outbound, it is the SID of the OutgoingCallerId resource from which the call was placed. |[optional] 
**Status** | Pointer to [**string**](CallEnumStatus.md) |  |
**StartTime** | **string** | The start time of the call. Null if the call has not yet been dialed. |[optional] 
**EndTime** | **string** | The end time of the call. Null if the call did not complete successfully. |[optional] 
**Duration** | **string** | The length of the call in seconds. |[optional] 
**Price** | Pointer to **string** | The charge for this call, in the currency associated with the account. Populated after the call is completed. May not be immediately available. |
**PriceUnit** | **string** | The currency in which `Price` is measured. |[optional] 
**Direction** | Pointer to **string** | A string describing the direction of the call. `inbound` for inbound calls, `outbound-api` for calls initiated via the REST API or `outbound-dial` for calls initiated by a `Dial` verb. |
**AnsweredBy** | **string** | Either `human` or `machine` if this call was initiated with answering machine detection. Empty otherwise. |[optional] 
**ApiVersion** | **string** | The API Version used to create the call |[optional] 
**ForwardedFrom** | **string** | The forwarding phone number if this call was an incoming call forwarded from another number (depends on carrier supporting forwarding). Otherwise, empty. |[optional] 
**GroupSid** | **string** | The Group SID associated with this call. If no Group is associated with the call, the field is empty. |[optional] 
**CallerName** | **string** | The caller's name if this call was an incoming call to a phone number with caller ID Lookup enabled. Otherwise, empty. |[optional] 
**QueueTime** | **string** | The wait time in milliseconds before the call is placed. |[optional] 
**TrunkSid** | **string** | The (optional) unique identifier of the trunk resource that was used for this call. |[optional] 
**Uri** | **string** | The URI of this resource, relative to `https://api.twilio.com` |[optional] 
**SubresourceUris** | **map[string]interface{}** | A list of related subresources identified by their relative URIs |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


