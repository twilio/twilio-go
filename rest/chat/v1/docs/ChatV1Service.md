# ChatV1Service

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that we created to identify the Service resource. |
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/api/rest/account) that created the Service resource. |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the resource. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was created specified in [RFC 2822](http://www.ietf.org/rfc/rfc2822.txt) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was last updated specified in [RFC 2822](http://www.ietf.org/rfc/rfc2822.txt) format. |
**DefaultServiceRoleSid** | Pointer to **string** | The service role assigned to users when they are added to the service. See the [Roles endpoint](https://www.twilio.com/docs/chat/api/roles) for more details. |
**DefaultChannelRoleSid** | Pointer to **string** | The channel role assigned to users when they are added to a channel. See the [Roles endpoint](https://www.twilio.com/docs/chat/api/roles) for more details. |
**DefaultChannelCreatorRoleSid** | Pointer to **string** | The channel role assigned to a channel creator when they join a new channel. See the [Roles endpoint](https://www.twilio.com/docs/chat/api/roles) for more details. |
**ReadStatusEnabled** | Pointer to **bool** | Whether the [Message Consumption Horizon](https://www.twilio.com/docs/chat/consumption-horizon) feature is enabled. The default is `true`. |
**ReachabilityEnabled** | Pointer to **bool** | Whether the [Reachability Indicator](https://www.twilio.com/docs/chat/reachability-indicator) is enabled for this Service instance. The default is `false`. |
**TypingIndicatorTimeout** | Pointer to **int** | How long in seconds after a `started typing` event until clients should assume that user is no longer typing, even if no `ended typing` message was received.  The default is 5 seconds. |
**ConsumptionReportInterval** | Pointer to **int** | DEPRECATED. The interval in seconds between consumption reports submission batches from client endpoints. |
**Limits** | Pointer to **interface{}** | An object that describes the limits of the service instance. The `limits` object contains  `channel_members` to describe the members/channel limit and `user_channels` to describe the channels/user limit. `channel_members` can be 1,000 or less, with a default of 250. `user_channels` can be 1,000 or less, with a default value of 100. |
**Webhooks** | Pointer to **interface{}** | An object that contains information about the webhooks configured for this service. |
**PreWebhookUrl** | Pointer to **string** | The URL for pre-event webhooks, which are called by using the `webhook_method`. See [Webhook Events](https://www.twilio.com/docs/api/chat/webhooks) for more details. |
**PostWebhookUrl** | Pointer to **string** | The URL for post-event webhooks, which are called by using the `webhook_method`. See [Webhook Events](https://www.twilio.com/docs/api/chat/webhooks) for more details. |
**WebhookMethod** | Pointer to **string** | The HTTP method to use for calls to the `pre_webhook_url` and `post_webhook_url` webhooks.  Can be: `POST` or `GET` and the default is `POST`. See [Webhook Events](https://www.twilio.com/docs/chat/webhook-events) for more details. |
**WebhookFilters** | Pointer to **[]string** | The list of WebHook events that are enabled for this Service instance. See [Webhook Events](https://www.twilio.com/docs/chat/webhook-events) for more details. |
**Notifications** | Pointer to **interface{}** | The notification configuration for the Service instance. See [Push Notification Configuration](https://www.twilio.com/docs/chat/push-notification-configuration) for more information. |
**Url** | Pointer to **string** | The absolute URL of the Service resource. |
**Links** | Pointer to **map[string]interface{}** | The absolute URLs of the Service's [Channels](https://www.twilio.com/docs/chat/api/channels), [Roles](https://www.twilio.com/docs/chat/api/roles), and [Users](https://www.twilio.com/docs/chat/api/users). |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


