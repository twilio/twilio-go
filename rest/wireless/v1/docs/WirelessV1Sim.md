# WirelessV1Sim

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that identifies the Sim resource |
**UniqueName** | Pointer to **string** | An application-defined string that uniquely identifies the resource |
**AccountSid** | Pointer to **string** | The SID of the Account to which the Sim resource belongs |
**RatePlanSid** | Pointer to **string** | The SID of the RatePlan resource to which the Sim resource is assigned. |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the Sim resource |
**Iccid** | Pointer to **string** | The ICCID associated with the SIM |
**EId** | Pointer to **string** | Deprecated |
**Status** | Pointer to [**string**](SimEnumStatus.md) |  |
**ResetStatus** | Pointer to [**string**](SimEnumResetStatus.md) |  |
**CommandsCallbackUrl** | Pointer to **string** | The URL we call when the SIM originates a machine-to-machine Command |
**CommandsCallbackMethod** | Pointer to **string** | The HTTP method we use to call commands_callback_url |
**SmsFallbackMethod** | Pointer to **string** | Deprecated |
**SmsFallbackUrl** | Pointer to **string** | Deprecated |
**SmsMethod** | Pointer to **string** | Deprecated |
**SmsUrl** | Pointer to **string** | Deprecated |
**VoiceFallbackMethod** | Pointer to **string** | Deprecated. The HTTP method we use to call voice_fallback_url |
**VoiceFallbackUrl** | Pointer to **string** | Deprecated. The URL we call when an error occurs while retrieving or executing the TwiML requested from voice_url |
**VoiceMethod** | Pointer to **string** | Deprecated. The HTTP method we use to call voice_url |
**VoiceUrl** | Pointer to **string** | Deprecated. The URL we call when the SIM-connected device makes a voice call |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the Sim resource was last updated |
**Url** | Pointer to **string** | The absolute URL of the resource |
**Links** | Pointer to **map[string]interface{}** | The URLs of related subresources |
**IpAddress** | Pointer to **string** | Deprecated |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


