# CreateConfigurationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | **string** | The display name of the Intelligence Configuration describing its purpose. |
**Description** | **string** | The description of the Intelligence Configuration further explaining its purpose. |[optional] 
**Rules** | [**[]RuleCreationRequestPayload**](RuleCreationRequestPayload.md) | List of Intelligence Configuration Rules that govern when and how Language Operators run. Each Rule represents a bundle of Operators, Triggers, Context, and Actions to be executed by the Intelligence Configuration on a Conversation. A maximum of five (5) Rules are allowed per Intelligence Configuration.  To create an Intelligence Configuration without any Rules configured yet, pass an empty array (`\"rules\": []`). The Configuration will not execute any Language Operators until at least one Rule has been added.  |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


