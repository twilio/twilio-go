# ConversationsV1ServiceWebhookConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The unique ID of the Account responsible for this service. |
**ChatServiceSid** | Pointer to **string** | The unique ID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) this conversation belongs to. |
**Filters** | Pointer to **[]string** | The list of events that your configured webhook targets will receive. Events not configured here will not fire. |
**Method** | Pointer to **string** | The HTTP method to be used when sending a webhook request |
**PostWebhookUrl** | Pointer to **string** | The absolute url the post-event webhook request should be sent to. |
**PreWebhookUrl** | Pointer to **string** | The absolute url the pre-event webhook request should be sent to. |
**Url** | Pointer to **string** | An absolute URL for this webhook. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


