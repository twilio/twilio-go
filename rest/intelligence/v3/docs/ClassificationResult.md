# ClassificationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OutputFormat** | **string** | The output format set on the Operator that generated this result. Determines the structure of the `result` object. |
**Id** | **string** | A unique identifier for the Operator Result. Assigned by Twilio (TTID). |
**AccountId** | **string** | The ID of the Account that created the Language Operator. |
**IntelligenceConfiguration** | [**IntelligenceConfigurationReference**](IntelligenceConfigurationReference.md) |  |
**ConversationId** | **string** | The `id` of the Conversation attached to the Operator Result. |
**Operator** | [**OperatorReference**](OperatorReference.md) |  |
**DateCreated** | [**time.Time**](time.Time.md) | Timestamp for when the Operator Result was created. |
**ReferenceIds** | **[]string** | The `id`s of objects related to this Operator Result. |
**ExecutionDetails** | [**ExecutionDetails**](ExecutionDetails.md) |  |
**Metadata** | [**OperatorResultsResponseBaseMetadata**](OperatorResultsResponseBaseMetadata.md) |  |
**Result** | [**ClassificationResultResult**](ClassificationResultResult.md) |  |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


