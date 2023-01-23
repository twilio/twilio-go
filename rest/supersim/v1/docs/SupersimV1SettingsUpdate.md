# SupersimV1SettingsUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | **string** | The unique identifier of this Settings Update |[optional] 
**Iccid** | **string** | The ICCID associated with the SIM |[optional] 
**SimSid** | **string** | The SID of the Super SIM to which this Settings Update was applied |[optional] 
**Status** | Pointer to [**string**](SettingsUpdateEnumStatus.md) |  |
**Packages** | **[]interface{}** | Array containing the different Settings Packages that will be applied to the SIM after the update completes |[optional] 
**DateCompleted** | [**time.Time**](time.Time.md) | The time when the update successfully completed and the new settings were applied to the SIM |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The date this Settings Update was created |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The date this Settings Update was last updated |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


