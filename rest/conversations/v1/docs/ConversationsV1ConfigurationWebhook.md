# ConversationsV1ConfigurationWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | **string** | The unique ID of the Account responsible for this conversation. |[optional] 
**Method** | Pointer to [**string**](ConfigurationWebhookEnumMethod.md) |  |
**Filters** | **[]string** | The list of webhook event triggers that are enabled for this Service. |[optional] 
**PreWebhookUrl** | **string** | The absolute url the pre-event webhook request should be sent to. |[optional] 
**PostWebhookUrl** | **string** | The absolute url the post-event webhook request should be sent to. |[optional] 
**Target** | Pointer to [**string**](ConfigurationWebhookEnumTarget.md) |  |
**Url** | **string** | An absolute URL for this webhook. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


