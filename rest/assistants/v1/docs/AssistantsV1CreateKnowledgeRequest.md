# AssistantsV1CreateKnowledgeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssistantId** | **string** | The Assistant ID. |[optional] 
**Description** | **string** | The description of the knowledge source. |[optional] 
**KnowledgeSourceDetails** | **map[string]interface{}** | The details of the knowledge source based on the type. |[optional] 
**Name** | **string** | The name of the tool. |
**Policy** | [**AssistantsV1CreatePolicyRequest**](AssistantsV1CreatePolicyRequest.md) |  |[optional] 
**Type** | **string** | The type of the knowledge source. |
**EmbeddingModel** | **string** | The embedding model to be used for the knowledge source. It's required for 'Database' type but disallowed for other types. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


