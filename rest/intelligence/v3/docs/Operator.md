# Operator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier for the Language Operator to be executed by the Rule. Assigned by Twilio (TTID). |
**Version** | **int** | The specific version of the Language Operator to execute. When provided, the Rule will use this exact version of the Operator. When omitted, the latest active version of the Operator is used at execution time.  |[optional] 
**Parameters** | **map[string]interface{}** | Key-value mapping for parameters defined as part of the Operator schema. The key and value passed by the Rule must match the name and data type of the parameter defined in the Operator, respectively. These parameters will customize the behavior of the Operator when executed by the Rule via runtime substitution into the prompt. Note: For parameters of type `knowledge_base_and_source_ids`, the value must be passed in the following format: `knowledge_base_id:knowledge_source_id`.  |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


