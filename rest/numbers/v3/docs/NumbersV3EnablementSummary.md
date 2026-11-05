# NumbersV3EnablementSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegistrationId** | **string** | The Compliance Registration ID. |
**AccountId** | **string** | The Account SID that created the association. |
**Status** | **string** | The association's status, set as the Compliance Registration changes state. |
**Errors** | [**[]NumbersV3NumbersSendersEnablementSummaryErrors**](NumbersV3NumbersSendersEnablementSummaryErrors.md) | Reserved for errors that block this association. Always empty today. Branch on `status`, not on the presence of errors.  |
**RegulationId** | **string** | The Regulation SID this association satisfies. |[optional] 
**RegulationName** | **string** | Human-readable name of the regulation. |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | When the association was created. |
**DateUpdated** | [**time.Time**](time.Time.md) | When the association was last updated. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


