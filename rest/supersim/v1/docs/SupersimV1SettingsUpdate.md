# SupersimV1SettingsUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique identifier of this Settings Update |
**Iccid** | Pointer to **string** | The ICCID associated with the SIM |
**SimSid** | Pointer to **string** | The SID of the Super SIM to which this Settings Update was applied |
**Status** | Pointer to [**string**](SettingsUpdateEnumStatus.md) |  |
**Packages** | Pointer to **[]interface{}** | Array containing the different Settings Packages that will be applied to the SIM after the update completes |
**DateCompleted** | Pointer to [**time.Time**](time.Time.md) | The time when the update successfully completed and the new settings were applied to the SIM |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date this Settings Update was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date this Settings Update was last updated |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


