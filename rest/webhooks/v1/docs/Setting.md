# Setting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique Setting identifier (TypeID/TTID format) |[readonly] 
**Type** | **string** | Setting ownership (read-only):  - **system**: Managed by Twilio   - Twilio provides managed Settings to simplify your onboarding journey   - Cannot be modified or deleted   - Used for default account configuration  - **customer**: Created by you   - Can update friendlyName via PATCH   - Can delete via DELETE (if not in use)   - Core config (auth, signature) is immutable  |[readonly] 
**FriendlyName** | **string** | Human-readable name |
**Description** | **string** | Optional description of the Setting's purpose and usage |[optional] 
**Auth** | Pointer to [**AuthSetting**](AuthSetting.md) |  |
**Signature** | Pointer to [**SignatureSetting**](SignatureSetting.md) |  |
**Connection** | Pointer to [**ConnectionSetting**](ConnectionSetting.md) | HTTP connection settings. If omitted, system defaults are applied. |
**EdgeZones** | Pointer to [**EdgeZones**](EdgeZones.md) | Edge zone configuration for webhook delivery. When omitted, all zones are eligible and Twilio selects freely. When set, `zones` lists preferred zones in priority order.  |
**CreatedAt** | [**time.Time**](time.Time.md) |  |[readonly] 
**UpdatedAt** | [**time.Time**](time.Time.md) |  |[readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


