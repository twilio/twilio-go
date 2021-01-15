# UpdateSessionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateExpiry** | [**time.Time**](time.Time.md) | The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date when the Session should expire. If this is value is present, it overrides the &#x60;ttl&#x60; value. | [optional] 
**FailOnParticipantConflict** | **bool** | [Experimental] For accounts with the ProxyAllowParticipantConflict account flag, setting to true enables per-request opt-in to allowing Proxy to return a 400 error (Twilio error code 80604) when a request to set a Session to in-progress would cause Participants with the same Identifier/ProxyIdentifier pair to be active in multiple Sessions. If not provided, requests will be allowed to succeed, and a Debugger notification (80801) will be emitted. Having multiple, active Participants with the same Identifier/ProxyIdentifier pair causes calls and messages from affected Participants to be routed incorrectly. Please note, the default behavior for accounts without the ProxyAllowParticipantConflict flag is to reject the request as described.  This will eventually be the default for all accounts. | [optional] 
**Status** | **string** | The new status of the resource. Can be: &#x60;in-progress&#x60; to re-open a session or &#x60;closed&#x60; to close a session. | [optional] 
**Ttl** | **int32** | The time, in seconds, when the session will expire. The time is measured from the last Session create or the Session&#39;s last Interaction. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


