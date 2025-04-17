# AssistantsV1Feedback

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssistantId** | **string** | The Assistant ID. |
**Id** | **string** | The Feedback ID. |
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Feedback. |[optional] 
**UserSid** | **string** | The SID of the User created the Feedback. |[optional] 
**MessageId** | **string** | The Message ID. |
**Score** | **float32** | The Score to provide as Feedback (0-1) |
**SessionId** | **string** | The Session ID. |
**Text** | **string** | The text to be given as feedback. |
**DateCreated** | [**time.Time**](time.Time.md) | The date and time in GMT when the Feedback was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**DateUpdated** | [**time.Time**](time.Time.md) | The date and time in GMT when the Feedback was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


