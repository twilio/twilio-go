# IntelligenceV2Sentence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MediaChannel** | **int** | The channel number. |[optional] [default to 0]
**SentenceIndex** | **int** | The index of the sentence in the transcript. |[optional] [default to 0]
**StartTime** | Pointer to **string** | Offset from the beginning of the transcript when this sentence starts. |
**EndTime** | Pointer to **string** | Offset from the beginning of the transcript when this sentence ends. |
**Transcript** | Pointer to **string** | Transcript text. |
**Sid** | Pointer to **string** | A 34 character string that uniquely identifies this Sentence. |
**Confidence** | Pointer to **string** |  |
**Words** | Pointer to **[]interface{}** | Detailed information for each of the words of the given Sentence. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


