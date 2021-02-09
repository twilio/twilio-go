# CreateSessionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateExpiry** | [**time.Time**](time.Time.md) | The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date when the Session should expire. If this is value is present, it overrides the &#x60;ttl&#x60; value. | [optional] 
**FailOnParticipantConflict** | **bool** | [Experimental] For accounts with the ProxyAllowParticipantConflict account flag, setting to true enables per-request opt-in to allowing Proxy to reject a Session create (with Participants) request that could cause the same Identifier/ProxyIdentifier pair to be active in multiple Sessions. Depending on the context, this could be a 409 error (Twilio error code 80623) or a 400 error (Twilio error code 80604). If not provided, requests will be allowed to succeed and a Debugger notification (80802) will be emitted. Having multiple, active Participants with the same Identifier/ProxyIdentifier pair causes calls and messages from affected Participants to be routed incorrectly. Please note, the default behavior for accounts without the ProxyAllowParticipantConflict flag is to reject the request as described.  This will eventually be the default for all accounts. | [optional] 
**Mode** | **string** | The Mode of the Session. Can be: &#x60;message-only&#x60;, &#x60;voice-only&#x60;, or &#x60;voice-and-message&#x60; and the default value is &#x60;voice-and-message&#x60;. | [optional] 
**Participants** | **[]map[string]interface{}** | The Participant objects to include in the new session. | [optional] 
**Status** | **string** | The initial status of the Session. Can be: &#x60;open&#x60;, &#x60;in-progress&#x60;, &#x60;closed&#x60;, &#x60;failed&#x60;, or &#x60;unknown&#x60;. The default is &#x60;open&#x60; on create. | [optional] 
**Ttl** | **int32** | The time, in seconds, when the session will expire. The time is measured from the last Session create or the Session&#39;s last Interaction. | [optional] 
**UniqueName** | **string** | An application-defined string that uniquely identifies the resource. This value must be 191 characters or fewer in length and be unique. **This value should not have PII.** | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


