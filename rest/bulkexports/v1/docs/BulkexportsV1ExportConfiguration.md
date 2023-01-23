# BulkexportsV1ExportConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | If true, Twilio will automatically generate every day's file when the day is over. |
**WebhookUrl** | Pointer to **string** | Stores the URL destination for the method specified in webhook_method. |
**WebhookMethod** | Pointer to **string** | Sets whether Twilio should call a webhook URL when the automatic generation is complete, using GET or POST. The actual destination is set in the webhook_url |
**ResourceType** | Pointer to **string** | The type of communication – Messages, Calls, Conferences, and Participants |
**Url** | Pointer to **string** | The URL of this resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


