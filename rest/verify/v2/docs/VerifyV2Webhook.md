# VerifyV2Webhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that we created to identify the Webhook resource. |
**ServiceSid** | Pointer to **string** | The unique SID identifier of the Service. |
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Service resource. |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the webhook. **This value should not contain PII.** |
**EventTypes** | Pointer to **[]string** | The array of events that this Webhook is subscribed to. Possible event types: `*, factor.deleted, factor.created, factor.verified, challenge.approved, challenge.denied` |
**Status** | Pointer to [**string**](WebhookEnumStatus.md) |  |
**Version** | Pointer to [**string**](WebhookEnumVersion.md) |  |
**WebhookUrl** | Pointer to **string** | The URL associated with this Webhook. |
**WebhookMethod** | Pointer to [**string**](WebhookEnumMethods.md) |  |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**Url** | Pointer to **string** | The absolute URL of the Webhook resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


