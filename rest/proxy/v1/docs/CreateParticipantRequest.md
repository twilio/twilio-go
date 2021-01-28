# CreateParticipantRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailOnParticipantConflict** | **bool** | [Experimental] For accounts with the ProxyAllowParticipantConflict account flag, setting to true enables per-request opt-in to allowing Proxy to reject a Participant create request that could cause the same Identifier/ProxyIdentifier pair to be active in multiple Sessions. Depending on the context, this could be a 409 error (Twilio error code 80623) or a 400 error (Twilio error code 80604). If not provided, requests will be allowed to succeed and a Debugger notification (80802) will be emitted. Having multiple, active Participants with the same Identifier/ProxyIdentifier pair causes calls and messages from affected Participants to be routed incorrectly. Please note, the default behavior for accounts without the ProxyAllowParticipantConflict flag is to reject the request as described.  This will eventually be the default for all accounts. | [optional] 
**FriendlyName** | **string** | The string that you assigned to describe the participant. This value must be 255 characters or fewer. **This value should not have PII.** | [optional] 
**Identifier** | **string** | The phone number of the Participant. | 
**ProxyIdentifier** | **string** | The proxy phone number to use for the Participant. If not specified, Proxy will select a number from the pool. | [optional] 
**ProxyIdentifierSid** | **string** | The SID of the Proxy Identifier to assign to the Participant. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


