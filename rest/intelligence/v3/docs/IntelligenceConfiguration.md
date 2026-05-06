# IntelligenceConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | The ID of the Account that created the Intelligence Configuration. |
**Id** | **string** | The unique identifier for the Intelligence Configuration. Assigned by Twilio (TTID). |
**DisplayName** | **string** | The display name of the Intelligence Configuration describing its purpose. |
**Description** | **string** | The description of the Intelligence Configuration further explaining its purpose. |[optional] 
**Version** | **int** | The numeric version of the Intelligence Configuration. Automatically incremented with each update on the resource, used to ensure integrity when updating the Configuration. |
**Rules** | [**[]Rule**](Rule.md) | List of Intelligence Configuration Rules that govern when and how Language Operators run. Each Rule represents a bundle of Operators, Triggers, Context, and Actions to be executed by the Intelligence Configuration on a Conversation. A maximum of five (5) Rules are allowed per Intelligence Configuration.  |
**DateCreated** | [**time.Time**](time.Time.md) | Timestamp of when the Intelligence Configuration was created. |
**DateUpdated** | [**time.Time**](time.Time.md) | Timestamp of when the Intelligence Configuration was last updated. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


