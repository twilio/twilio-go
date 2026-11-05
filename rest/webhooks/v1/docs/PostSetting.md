# PostSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FriendlyName** | Pointer to **string** | Human-readable name (optional, defaults to blank) |
**Description** | **string** | Optional description of the Setting's purpose and usage |[optional] 
**Auth** | Pointer to [**PostAuthSetting**](PostAuthSetting.md) |  |
**Signature** | Pointer to [**SignatureSetting**](SignatureSetting.md) |  |
**Connection** | Pointer to [**ConnectionSetting**](ConnectionSetting.md) | HTTP connection settings. If omitted, system defaults are applied. |
**EdgeZones** | Pointer to [**EdgeZones**](EdgeZones.md) | Edge zone configuration for webhook delivery. When omitted, all zones are eligible and Twilio selects freely. When set, `zones` lists preferred zones in priority order.  |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


