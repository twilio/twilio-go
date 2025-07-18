# IntelligenceV2OperatorResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OperatorType** | Pointer to [**string**](OperatorResultEnumOperatorType.md) |  |
**Name** | Pointer to **string** | The name of the applied Language Understanding. |
**OperatorSid** | Pointer to **string** | A 34 character string that identifies this Language Understanding operator sid. |
**ExtractMatch** | Pointer to **bool** | Boolean to tell if extract Language Understanding Processing model matches results. |
**MatchProbability** | Pointer to **float32** | Percentage of 'matching' class needed to consider a sentence matches |
**NormalizedResult** | Pointer to **string** | Normalized output of extraction stage which matches Label. |
**UtteranceResults** | Pointer to **[]interface{}** | List of mapped utterance object which matches sentences. |
**UtteranceMatch** | Pointer to **bool** | Boolean to tell if Utterance matches results. |
**PredictedLabel** | Pointer to **string** | The 'matching' class. This might be available on conversation classify model outputs. |
**PredictedProbability** | Pointer to **float32** | Percentage of 'matching' class needed to consider a sentence matches. |
**LabelProbabilities** | Pointer to **interface{}** | The labels probabilities. This might be available on conversation classify model outputs. |
**ExtractResults** | Pointer to **interface{}** | List of text extraction results. This might be available on classify-extract model outputs. |
**TextGenerationResults** | Pointer to **interface{}** | Output of a text generation operator for example Conversation Sumamary. |
**JsonResults** | Pointer to **interface{}** |  |
**TranscriptSid** | Pointer to **string** | A 34 character string that uniquely identifies this Transcript. |
**Url** | Pointer to **string** | The URL of this resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


