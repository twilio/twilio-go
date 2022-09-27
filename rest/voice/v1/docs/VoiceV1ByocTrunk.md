# VoiceV1ByocTrunk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the resource |
**VoiceUrl** | Pointer to **string** | The URL we call when receiving a call |
**VoiceMethod** | Pointer to **string** | The HTTP method to use with voice_url |
**VoiceFallbackUrl** | Pointer to **string** | The URL we call when an error occurs while executing TwiML |
**VoiceFallbackMethod** | Pointer to **string** | The HTTP method used with voice_fallback_url |
**StatusCallbackUrl** | Pointer to **string** | The URL that we call with status updates |
**StatusCallbackMethod** | Pointer to **string** | The HTTP method we use to call status_callback_url |
**CnamLookupEnabled** | Pointer to **bool** | Whether Caller ID Name (CNAM) lookup is enabled for the trunk |
**ConnectionPolicySid** | Pointer to **string** | Origination Connection Policy (to your Carrier) |
**FromDomainSid** | Pointer to **string** | The SID of the SIP Domain that should be used in the `From` header of originating calls |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT that the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT that the resource was last updated |
**Url** | Pointer to **string** | The absolute URL of the resource |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


