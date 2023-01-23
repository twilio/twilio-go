# ConversationsV1ServiceWebhookConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | **string** | The unique ID of the Account responsible for this service. |[optional] 
**ChatServiceSid** | **string** | The unique ID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) this conversation belongs to. |[optional] 
**PreWebhookUrl** | **string** | The absolute url the pre-event webhook request should be sent to. |[optional] 
**PostWebhookUrl** | **string** | The absolute url the post-event webhook request should be sent to. |[optional] 
**Filters** | **[]string** | The list of events that your configured webhook targets will receive. Events not configured here will not fire. |[optional] 
**Method** | Pointer to [**string**](ServiceWebhookConfigurationEnumMethod.md) |  |
**Url** | **string** | An absolute URL for this webhook. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


