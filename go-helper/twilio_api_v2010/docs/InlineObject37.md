# InlineObject37

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ByocTrunkSid** | **string** | The SID of the BYOC Trunk(Bring Your Own Carrier) resource that the Sip Domain will be associated with. | [optional] 
**DomainName** | **string** | The unique address you reserve on Twilio to which you route your SIP traffic. Domain names can contain letters, digits, and \&quot;-\&quot;. | [optional] 
**EmergencyCallerSid** | **string** | Whether an emergency caller sid is configured for the domain. If present, this phone number will be used as the callback for the emergency call. | [optional] 
**EmergencyCallingEnabled** | **bool** | Whether emergency calling is enabled for the domain. If enabled, allows emergency calls on the domain from phone numbers with validated addresses. | [optional] 
**FriendlyName** | **string** | A descriptive string that you created to describe the resource. It can be up to 64 characters long. | [optional] 
**Secure** | **bool** | Whether secure SIP is enabled for the domain. If enabled, TLS will be enforced and SRTP will be negotiated on all incoming calls to this sip domain. | [optional] 
**SipRegistration** | **bool** | Whether to allow SIP Endpoints to register with the domain to receive calls. Can be &#x60;true&#x60; or &#x60;false&#x60;. &#x60;true&#x60; allows SIP Endpoints to register with the domain to receive calls, &#x60;false&#x60; does not. | [optional] 
**VoiceFallbackMethod** | **string** | The HTTP method we should use to call &#x60;voice_fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;. | [optional] 
**VoiceFallbackUrl** | **string** | The URL that we should call when an error occurs while retrieving or executing the TwiML requested by &#x60;voice_url&#x60;. | [optional] 
**VoiceMethod** | **string** | The HTTP method we should use to call &#x60;voice_url&#x60; | [optional] 
**VoiceStatusCallbackMethod** | **string** | The HTTP method we should use to call &#x60;voice_status_callback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;. | [optional] 
**VoiceStatusCallbackUrl** | **string** | The URL that we should call to pass status parameters (such as call ended) to your application. | [optional] 
**VoiceUrl** | **string** | The URL we should call when the domain receives a call. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


