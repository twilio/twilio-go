# NumbersV1CreateEmbeddedRegistrationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegulationId** | **string** | The regulation for this registration. |
**RegulationVersion** | **int** | The regulation version. |
**FriendlyName** | **string** | Human-readable name for the registration. |
**StatusNotificationEmail** | **string** | Email address for registration status notifications. |[optional] 
**StatusCallbackUrl** | Pointer to **string** | The URL of this resource. |
**Comments** | **string** | Additional comments about the registration. |[optional] 
**ThemeSetId** | **string** | Theme ID for the Compliance Embeddable UI. |[optional] 
**Data** | **map[string]interface{}** | Registration data organized by section (alphanumericSender, business, useCase, authorizedRepresentative, officer, businessAddress). |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


