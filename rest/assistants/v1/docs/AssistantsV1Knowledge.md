# AssistantsV1Knowledge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | The type of knowledge source. |[optional] 
**Id** | **string** | The description of knowledge. |
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Knowledge resource. |[optional] 
**KnowledgeSourceDetails** | **map[string]interface{}** | The details of the knowledge source based on the type. |[optional] 
**Name** | **string** | The name of the knowledge source. |
**Status** | **string** | The status of processing the knowledge source ('QUEUED', 'PROCESSING', 'COMPLETED', 'FAILED') |[optional] 
**Type** | **string** | The type of knowledge source ('Web', 'Database', 'Text', 'File') |
**Url** | **string** | The url of the knowledge resource. |[optional] 
**EmbeddingModel** | **string** | The embedding model to be used for the knowledge source. |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The date and time in GMT when the Knowledge was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**DateUpdated** | [**time.Time**](time.Time.md) | The date and time in GMT when the Knowledge was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


