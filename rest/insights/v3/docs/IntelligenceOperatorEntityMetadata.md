# IntelligenceOperatorEntityMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceType** | [**SourceType**](SourceType.md) |  |
**OperatorId** | **string** | The Intelligence operator ID. Part of the owner check for idempotent upserts — a mismatch against the existing registration's `(sourceType, operatorId)` returns `409`.  |
**OperatorName** | **string** | Optional, display-only name of the operator as supplied at registration time. Not part of any uniqueness or validation logic.  |[optional] 
**Field** | **string** | Annotation fieldPath key; generic identifier that determines the output-schema field. Resolved against the operator identified by `operatorId`.  |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


