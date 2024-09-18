# AssistantsV1AssistantSendMessageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identity** | **string** | The unique identity of user for the session. |
**SessionId** | **string** | The unique name for the session. |[optional] 
**Body** | **string** | The query to ask the assistant. |
**Webhook** | **string** | The webhook url to call after the assistant has generated a response or report an error. |[optional] 
**Mode** | **string** | one of the modes 'chat', 'email' or 'voice' |[optional] [default to "chat"]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


