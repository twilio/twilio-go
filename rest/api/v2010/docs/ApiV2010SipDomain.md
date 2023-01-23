# ApiV2010SipDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**ApiVersion** | **string** | The API version used to process the call |[optional] 
**AuthType** | **string** | The types of authentication mapped to the domain |[optional] 
**DateCreated** | **string** | The RFC 2822 date and time in GMT that the resource was created |[optional] 
**DateUpdated** | **string** | The RFC 2822 date and time in GMT that the resource was last updated |[optional] 
**DomainName** | **string** | The unique address on Twilio to route SIP traffic |[optional] 
**FriendlyName** | **string** | The string that you assigned to describe the resource |[optional] 
**Sid** | **string** | The unique string that identifies the resource |[optional] 
**Uri** | **string** | The URI of the resource, relative to `https://api.twilio.com` |[optional] 
**VoiceFallbackMethod** | **string** | The HTTP method used with voice_fallback_url |[optional] 
**VoiceFallbackUrl** | **string** | The URL we call when an error occurs while executing TwiML |[optional] 
**VoiceMethod** | **string** | The HTTP method to use with voice_url |[optional] 
**VoiceStatusCallbackMethod** | **string** | The HTTP method we use to call voice_status_callback_url |[optional] 
**VoiceStatusCallbackUrl** | **string** | The URL that we call with status updates |[optional] 
**VoiceUrl** | **string** | The URL we call when receiving a call |[optional] 
**SubresourceUris** | **map[string]interface{}** | A list mapping resources associated with the SIP Domain resource |[optional] 
**SipRegistration** | **bool** | Whether SIP registration is allowed |[optional] 
**EmergencyCallingEnabled** | **bool** | Whether emergency calling is enabled for the domain. |[optional] 
**Secure** | **bool** | Whether secure SIP is enabled for the domain |[optional] 
**ByocTrunkSid** | **string** | The SID of the BYOC Trunk resource. |[optional] 
**EmergencyCallerSid** | **string** | Whether an emergency caller sid is configured for the domain. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


