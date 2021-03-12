# WirelessV1Sim

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account to which the Sim resource belongs |
**CommandsCallbackMethod** | Pointer to **string** | The HTTP method we use to call commands_callback_url |
**CommandsCallbackUrl** | Pointer to **string** | The URL we call when the SIM originates a machine-to-machine Command |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the Sim resource was last updated |
**EId** | Pointer to **string** | Deprecated |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the Sim resource |
**Iccid** | Pointer to **string** | The ICCID associated with the SIM |
**IpAddress** | Pointer to **string** | Deprecated |
**Links** | Pointer to **map[string]interface{}** | The URLs of related subresources |
**RatePlanSid** | Pointer to **string** | The SID of the RatePlan resource to which the Sim resource is assigned. |
**ResetStatus** | Pointer to **string** | The connectivity reset status of the SIM |
**Sid** | Pointer to **string** | The unique string that identifies the Sim resource |
**SmsFallbackMethod** | Pointer to **string** | The HTTP method we use to call sms_fallback_url |
**SmsFallbackUrl** | Pointer to **string** | The URL we call when an error occurs while retrieving or executing the TwiML requested from the sms_url |
**SmsMethod** | Pointer to **string** | The HTTP method we use to call sms_url |
**SmsUrl** | Pointer to **string** | The URL we call when the SIM-connected device sends an SMS message that is not a Command |
**Status** | Pointer to **string** | The status of the Sim resource |
**UniqueName** | Pointer to **string** | An application-defined string that uniquely identifies the resource |
**Url** | Pointer to **string** | The absolute URL of the resource |
**VoiceFallbackMethod** | Pointer to **string** | The HTTP method we use to call voice_fallback_url |
**VoiceFallbackUrl** | Pointer to **string** | The URL we call when an error occurs while retrieving or executing the TwiML requested from voice_url |
**VoiceMethod** | Pointer to **string** | The HTTP method we use to call voice_url |
**VoiceUrl** | Pointer to **string** | The URL we call when the SIM-connected device makes a voice call |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


