# VoiceV1ByocTrunk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**CnamLookupEnabled** | Pointer to **bool** | Whether Caller ID Name (CNAM) lookup is enabled for the trunk |
**ConnectionPolicySid** | Pointer to **string** | Origination Connection Policy (to your Carrier) |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT that the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT that the resource was last updated |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the resource |
**FromDomainSid** | Pointer to **string** | The SID of the SIP Domain that should be used in the `From` header of originating calls |
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**StatusCallbackMethod** | Pointer to **string** | The HTTP method we use to call status_callback_url |
**StatusCallbackUrl** | Pointer to **string** | The URL that we call with status updates |
**Url** | Pointer to **string** | The absolute URL of the resource |
**VoiceFallbackMethod** | Pointer to **string** | The HTTP method used with voice_fallback_url |
**VoiceFallbackUrl** | Pointer to **string** | The URL we call when an error occurs while executing TwiML |
**VoiceMethod** | Pointer to **string** | The HTTP method to use with voice_url |
**VoiceUrl** | Pointer to **string** | The URL we call when receiving a call |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


