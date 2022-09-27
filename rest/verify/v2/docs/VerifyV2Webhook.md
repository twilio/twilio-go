# VerifyV2Webhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**ServiceSid** | Pointer to **string** | Service Sid. |
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the webhook |
**EventTypes** | Pointer to **[]string** | The array of events that this Webhook is subscribed to. |
**Status** | Pointer to [**string**](WebhookEnumStatus.md) |  |
**Version** | Pointer to [**string**](WebhookEnumVersion.md) |  |
**WebhookUrl** | Pointer to **string** | The URL associated with this Webhook. |
**WebhookMethod** | Pointer to [**string**](WebhookEnumMethods.md) |  |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |
**Url** | Pointer to **string** | The absolute URL of the Webhook resource |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


