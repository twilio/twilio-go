# UpdateConfigurationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | **string** | The display name of the Intelligence Configuration describing its purpose. |
**Description** | **string** | The description of the Intelligence Configuration further explaining its purpose. |[optional] 
**Rules** | [**[]RuleUpdateRequestPayload**](RuleUpdateRequestPayload.md) | List of Intelligence Configuration Rules that govern when and how Language Operators run. Each Rule represents a bundle of Operators, Triggers, Context, and Actions to be executed by the Intelligence Configuration on a Conversation. A maximum of five (5) Rules are allowed per Intelligence Configuration.  |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


