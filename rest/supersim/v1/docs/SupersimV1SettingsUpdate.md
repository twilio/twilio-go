# SupersimV1SettingsUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique identifier of this Settings Update. |
**Iccid** | Pointer to **string** | The [ICCID](https://en.wikipedia.org/wiki/SIM_card#ICCID) associated with the SIM. |
**SimSid** | Pointer to **string** | The SID of the Super SIM to which this Settings Update was applied. |
**Status** | Pointer to [**string**](SettingsUpdateEnumStatus.md) |  |
**Packages** | Pointer to **[]interface{}** | Array containing the different Settings Packages that will be applied to the SIM after the update completes. Each object within the array indicates the name and the version of the Settings Package that will be on the SIM if the update is successful. |
**DateCompleted** | Pointer to [**time.Time**](time.Time.md) | The time, given in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format, when the update successfully completed and the new settings were applied to the SIM. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date that this Settings Update was created, given in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date that this Settings Update was updated, given in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


