# UpdateByocTrunkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CnamLookupEnabled** | **bool** | Whether Caller ID Name (CNAM) lookup is enabled for the trunk. If enabled, all inbound calls to the BYOC Trunk from the United States and Canada automatically perform a CNAM Lookup and display Caller ID data on your phone. See [CNAM Lookups](https://www.twilio.com/docs/sip-trunking#CNAM) for more information. | [optional] 
**ConnectionPolicySid** | **string** | The SID of the Connection Policy that Twilio will use when routing traffic to your communications infrastructure. | [optional] 
**FriendlyName** | **string** | A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long. | [optional] 
**FromDomainSid** | **string** | The SID of the SIP Domain that should be used in the &#x60;From&#x60; header of originating calls sent to your SIP infrastructure. If your SIP infrastructure allows users to \&quot;call back\&quot; an incoming call, configure this with a [SIP Domain](https://www.twilio.com/docs/voice/api/sending-sip) to ensure proper routing. If not configured, the from domain will default to \&quot;sip.twilio.com\&quot;. | [optional] 
**StatusCallbackMethod** | **string** | The HTTP method we should use to call &#x60;status_callback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;. | [optional] 
**StatusCallbackUrl** | **string** | The URL that we should call to pass status parameters (such as call ended) to your application. | [optional] 
**VoiceFallbackMethod** | **string** | The HTTP method we should use to call &#x60;voice_fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;. | [optional] 
**VoiceFallbackUrl** | **string** | The URL that we should call when an error occurs while retrieving or executing the TwiML requested by &#x60;voice_url&#x60;. | [optional] 
**VoiceMethod** | **string** | The HTTP method we should use to call &#x60;voice_url&#x60; | [optional] 
**VoiceUrl** | **string** | The URL we should call when the BYOC Trunk receives a call. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


