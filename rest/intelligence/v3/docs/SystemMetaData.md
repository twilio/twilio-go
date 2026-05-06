# SystemMetaData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResolvedModel** | **string** | The underlying LLM model used for this execution. |[optional] 
**LatencyMs** | **int** | Operator execution time in milliseconds. |
**InputCharacters** | **int** | Number of characters in the input sent to the model. Aligns with the billing unit. |
**OutputCharacters** | **int** | Number of characters in the model's response. Aligns with the billing unit. |
**InputTruncated** | **bool** | Whether the conversation input was truncated to fit the model's context window. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


