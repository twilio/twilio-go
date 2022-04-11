# ChatV2Service

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**ConsumptionReportInterval** | Pointer to **int** | DEPRECATED |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was last updated |
**DefaultChannelCreatorRoleSid** | Pointer to **string** | The channel role assigned to a channel creator when they join a new channel |
**DefaultChannelRoleSid** | Pointer to **string** | The channel role assigned to users when they are added to a channel |
**DefaultServiceRoleSid** | Pointer to **string** | The service role assigned to users when they are added to the service |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the resource |
**Limits** | Pointer to **interface{}** | An object that describes the limits of the service instance |
**Links** | Pointer to **map[string]interface{}** | The absolute URLs of the Service's Channels, Roles, and Users |
**Media** | Pointer to **interface{}** | The properties of the media that the service supports |
**Notifications** | Pointer to **interface{}** | The notification configuration for the Service instance |
**PostWebhookRetryCount** | Pointer to **int** | The number of times calls to the `post_webhook_url` will be retried |
**PostWebhookUrl** | Pointer to **string** | The URL for post-event webhooks |
**PreWebhookRetryCount** | Pointer to **int** | Count of times webhook will be retried in case of timeout or 429/503/504 HTTP responses |
**PreWebhookUrl** | Pointer to **string** | The webhook URL for pre-event webhooks |
**ReachabilityEnabled** | Pointer to **bool** | Whether the Reachability Indicator feature is enabled for this Service instance |
**ReadStatusEnabled** | Pointer to **bool** | Whether the Message Consumption Horizon feature is enabled |
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**TypingIndicatorTimeout** | Pointer to **int** | How long in seconds to wait before assuming the user is no longer typing |
**Url** | Pointer to **string** | The absolute URL of the Service resource |
**WebhookFilters** | Pointer to **[]string** | The list of webhook events that are enabled for this Service instance |
**WebhookMethod** | Pointer to **string** | The HTTP method  to use for both PRE and POST webhooks |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


