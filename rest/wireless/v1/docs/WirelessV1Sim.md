# WirelessV1Sim

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | **string** | The unique string that identifies the Sim resource |[optional] 
**UniqueName** | **string** | An application-defined string that uniquely identifies the resource |[optional] 
**AccountSid** | **string** | The SID of the Account to which the Sim resource belongs |[optional] 
**RatePlanSid** | **string** | The SID of the RatePlan resource to which the Sim resource is assigned. |[optional] 
**FriendlyName** | **string** | The string that you assigned to describe the Sim resource |[optional] 
**Iccid** | **string** | The ICCID associated with the SIM |[optional] 
**EId** | **string** | Deprecated |[optional] 
**Status** | Pointer to [**string**](SimEnumStatus.md) |  |
**ResetStatus** | Pointer to [**string**](SimEnumResetStatus.md) |  |
**CommandsCallbackUrl** | **string** | The URL we call when the SIM originates a machine-to-machine Command |[optional] 
**CommandsCallbackMethod** | **string** | The HTTP method we use to call commands_callback_url |[optional] 
**SmsFallbackMethod** | **string** | Deprecated |[optional] 
**SmsFallbackUrl** | **string** | Deprecated |[optional] 
**SmsMethod** | **string** | Deprecated |[optional] 
**SmsUrl** | **string** | Deprecated |[optional] 
**VoiceFallbackMethod** | **string** | Deprecated. The HTTP method we use to call voice_fallback_url |[optional] 
**VoiceFallbackUrl** | **string** | Deprecated. The URL we call when an error occurs while retrieving or executing the TwiML requested from voice_url |[optional] 
**VoiceMethod** | **string** | Deprecated. The HTTP method we use to call voice_url |[optional] 
**VoiceUrl** | **string** | Deprecated. The URL we call when the SIM-connected device makes a voice call |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the Sim resource was last updated |[optional] 
**Url** | **string** | The absolute URL of the resource |[optional] 
**Links** | **map[string]interface{}** | The URLs of related subresources |[optional] 
**IpAddress** | **string** | Deprecated |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


