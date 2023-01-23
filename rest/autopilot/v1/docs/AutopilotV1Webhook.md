# AutopilotV1Webhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | The absolute URL of the Webhook resource |[optional] 
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was created |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was last updated |[optional] 
**AssistantSid** | **string** | The SID of the Assistant that is the parent of the resource |[optional] 
**Sid** | **string** | The unique string that identifies the resource |[optional] 
**UniqueName** | **string** | An application-defined string that uniquely identifies the resource |[optional] 
**Events** | **string** | The list of space-separated events that this Webhook is subscribed to. |[optional] 
**WebhookUrl** | **string** | The URL associated with this Webhook. |[optional] 
**WebhookMethod** | **string** | The method used when calling the webhook's URL. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


