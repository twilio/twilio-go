# OperatorParameter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The data type of the parameter (e.g., STRING, INTEGER). Available values: STRING, INTEGER, NUMBER, BOOLEAN, KNOWLEDGE_BASE_AND_SOURCE_IDS  KNOWLEDGE_BASE_AND_SOURCE_IDS is a special type of parameter that refers to a Memora Knowledge Source, prefixed with its Knowledge Base container. Support for KB type = plaintext only. During Intelligence Configuration creation, this parameter is linked to a specific Knowledge Source. The operator receives the resolved plaintext at runtime and injects it into the prompt. The value of this parameter is expected to be passed in the following format: knowledge_base_id:knowledge_source_id. |
**Default** | Pointer to **interface{}** | Default value to use in the prompt if no value is provided by the Intelligence Configuration. Note: knowledge_base_and_source_ids does not support the default attribute, but all other param types do allow for a default value. |
**Required** | **bool** | Whether this parameter must be set at Operator execution time. Defaults to false if not provided. |[optional] 
**Description** | **string** | A human-readable description of the parameter. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


