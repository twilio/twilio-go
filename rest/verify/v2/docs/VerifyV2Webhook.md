# VerifyV2Webhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | **string** | The unique string that identifies the resource |[optional] 
**ServiceSid** | **string** | Service Sid. |[optional] 
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**FriendlyName** | **string** | The string that you assigned to describe the webhook |[optional] 
**EventTypes** | **[]string** | The array of events that this Webhook is subscribed to. |[optional] 
**Status** | Pointer to [**string**](WebhookEnumStatus.md) |  |
**Version** | Pointer to [**string**](WebhookEnumVersion.md) |  |
**WebhookUrl** | **string** | The URL associated with this Webhook. |[optional] 
**WebhookMethod** | Pointer to [**string**](WebhookEnumMethods.md) |  |
**DateCreated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |[optional] 
**Url** | **string** | The absolute URL of the Webhook resource |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


