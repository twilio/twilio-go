# ConversationsV1ConfigurationWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The unique ID of the Account responsible for this conversation. |
**Method** | Pointer to [**string**](ConfigurationWebhookEnumMethod.md) |  |
**Filters** | Pointer to **[]string** | The list of webhook event triggers that are enabled for this Service. |
**PreWebhookUrl** | Pointer to **string** | The absolute url the pre-event webhook request should be sent to. |
**PostWebhookUrl** | Pointer to **string** | The absolute url the post-event webhook request should be sent to. |
**Target** | Pointer to [**string**](ConfigurationWebhookEnumTarget.md) |  |
**Url** | Pointer to **string** | An absolute URL for this webhook. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


