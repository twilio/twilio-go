# AssistantsV1AssistantWithToolsAndKnowledge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Assistant resource. |
**CustomerAi** | **map[string]interface{}** | The Personalization and Perception Engine settings. |
**Id** | **string** | The Assistant ID. |
**Model** | **string** | The default model used by the assistant. |
**Name** | **string** | The name of the assistant. |
**Owner** | **string** | The owner/company of the assistant. |
**Url** | **string** | The url of the assistant resource. |[optional] 
**PersonalityPrompt** | **string** | The personality prompt to be used for assistant. |
**DateCreated** | [**time.Time**](time.Time.md) | The date and time in GMT when the Assistant was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**DateUpdated** | [**time.Time**](time.Time.md) | The date and time in GMT when the Assistant was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**Knowledge** | [**[]AssistantsV1Knowledge**](AssistantsV1Knowledge.md) | The list of knowledge sources associated with the assistant. |
**Tools** | [**[]AssistantsV1Tool**](AssistantsV1Tool.md) | The list of tools associated with the assistant. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


