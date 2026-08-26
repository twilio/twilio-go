# ConversationsV2OrchestratorPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Handler that places the call. Only STUDIO_FLOW is currently supported. |
**SystemFlow** | **string** | Behaviour to run for the call. |[optional] [default to "OUTBOUND_CALL"]
**OnSuccess** | [**ConversationsV2OrchestratorPolicyAction**](ConversationsV2OrchestratorPolicyAction.md) |  |[optional] 
**OnFailure** | [**ConversationsV2OrchestratorPolicyAction**](ConversationsV2OrchestratorPolicyAction.md) |  |[optional] 
**Parameters** | **map[string]interface{}** | Optional key/value pairs passed through to the handler that places the call. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


