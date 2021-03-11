# ApiV2010AccountSipSipDomain

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**ApiVersion** | Pointer to **string** | The API version used to process the call |
**AuthType** | Pointer to **string** | The types of authentication mapped to the domain |
**ByocTrunkSid** | Pointer to **string** | The SID of the BYOC Trunk resource. |
**DateCreated** | Pointer to **string** | The RFC 2822 date and time in GMT that the resource was created |
**DateUpdated** | Pointer to **string** | The RFC 2822 date and time in GMT that the resource was last updated |
**DomainName** | Pointer to **string** | The unique address on Twilio to route SIP traffic |
**EmergencyCallerSid** | Pointer to **string** | Whether an emergency caller sid is configured for the domain. |
**EmergencyCallingEnabled** | Pointer to **bool** | Whether emergency calling is enabled for the domain. |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the resource |
**Secure** | Pointer to **bool** | Whether secure SIP is enabled for the domain |
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**SipRegistration** | Pointer to **bool** | Whether SIP registration is allowed |
**SubresourceUris** | Pointer to **map[string]interface{}** | A list mapping resources associated with the SIP Domain resource |
**Uri** | Pointer to **string** | The URI of the resource, relative to `https://api.twilio.com` |
**VoiceFallbackMethod** | Pointer to **string** | The HTTP method used with voice_fallback_url |
**VoiceFallbackUrl** | Pointer to **string** | The URL we call when an error occurs while executing TwiML |
**VoiceMethod** | Pointer to **string** | The HTTP method to use with voice_url |
**VoiceStatusCallbackMethod** | Pointer to **string** | The HTTP method we use to call voice_status_callback_url |
**VoiceStatusCallbackUrl** | Pointer to **string** | The URL that we call with status updates |
**VoiceUrl** | Pointer to **string** | The URL we call when receiving a call |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


