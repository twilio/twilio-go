# RuleUpdateRequestPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Optional field used when updating an existing Rule within an Intelligence Configuration. When provided, the Rule with this `id` is updated; when omitted, a new Rule is created.  |[optional] 
**Operators** | [**[]Operator**](Operator.md) | List of Operators to be executed by the Rule. Minimum of one (1) and maximum of five (5) Operators allowed per Rule. |
**Triggers** | [**[]Trigger**](Trigger.md) | List of Triggers that determine when to activate the Rule. Maximum of one (1) Trigger allowed per Rule. |[optional] 
**Actions** | [**[]Action**](Action.md) | List of Actions to be performed after the Rule is triggered. Maximum of two (2) Actions allowed per Rule. |
**Context** | [**Context**](Context.md) |  |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


