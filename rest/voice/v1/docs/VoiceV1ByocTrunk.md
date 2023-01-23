# VoiceV1ByocTrunk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**Sid** | **string** | The unique string that identifies the resource |[optional] 
**FriendlyName** | **string** | The string that you assigned to describe the resource |[optional] 
**VoiceUrl** | **string** | The URL we call when receiving a call |[optional] 
**VoiceMethod** | **string** | The HTTP method to use with voice_url |[optional] 
**VoiceFallbackUrl** | **string** | The URL we call when an error occurs while executing TwiML |[optional] 
**VoiceFallbackMethod** | **string** | The HTTP method used with voice_fallback_url |[optional] 
**StatusCallbackUrl** | **string** | The URL that we call with status updates |[optional] 
**StatusCallbackMethod** | **string** | The HTTP method we use to call status_callback_url |[optional] 
**CnamLookupEnabled** | **bool** | Whether Caller ID Name (CNAM) lookup is enabled for the trunk |[optional] 
**ConnectionPolicySid** | **string** | Origination Connection Policy (to your Carrier) |[optional] 
**FromDomainSid** | **string** | The SID of the SIP Domain that should be used in the `From` header of originating calls |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT that the resource was created |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT that the resource was last updated |[optional] 
**Url** | **string** | The absolute URL of the resource |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


