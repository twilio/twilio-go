# AssistantsV1AssistantSendMessageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | success or failure based on whether the request successfully generated a response. |
**Flagged** | **bool** | If successful, this property will denote whether the response was flagged or not. |[optional] 
**Aborted** | **bool** | This property will denote whether the request was aborted or not. |[optional] 
**SessionId** | **string** | The unique name for the session. |
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that sent the Message. |
**Body** | **string** | If successful, the body of the generated response |[optional] 
**Error** | **string** | The error message if generation was not successful |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


