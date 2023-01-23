# ChatV2ChannelWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that we created to identify the Channel Webhook resource. |
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Channel Webhook resource. |
**ServiceSid** | Pointer to **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) the Channel Webhook resource is associated with. |
**ChannelSid** | Pointer to **string** | The SID of the [Channel](https://www.twilio.com/docs/chat/channels) the Channel Webhook resource belongs to. |
**Type** | Pointer to **string** | The type of webhook. Can be: `webhook`, `studio`, or `trigger`. |
**Url** | Pointer to **string** | The absolute URL of the Channel Webhook resource. |
**Configuration** | Pointer to **interface{}** | The JSON string that describes how the channel webhook is configured. The configuration object contains the `url`, `method`, `filters`, and `retry_count` values that are configured by the create and update actions. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


