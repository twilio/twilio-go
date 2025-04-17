# AssistantsV1Message

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The message ID. |[optional] 
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Message resource. |[optional] 
**AssistantId** | **string** | The Assistant ID. |[optional] 
**SessionId** | **string** | The Session ID. |[optional] 
**Identity** | **string** | The identity of the user. |[optional] 
**Role** | **string** | The role of the user associated with the message. |[optional] 
**Content** | **map[string]interface{}** | The content of the message. |[optional] 
**Meta** | **map[string]interface{}** | The metadata of the message. |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The date and time in GMT when the Message was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The date and time in GMT when the Message was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


