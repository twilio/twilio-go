# UpdateServiceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConsumptionReportInterval** | **int32** | DEPRECATED. The interval in seconds between consumption reports submission batches from client endpoints. | [optional] 
**DefaultChannelCreatorRoleSid** | **string** | The channel role assigned to a channel creator when they join a new channel. See the [Role resource](https://www.twilio.com/docs/chat/rest/role-resource) for more info about roles. | [optional] 
**DefaultChannelRoleSid** | **string** | The channel role assigned to users when they are added to a channel. See the [Role resource](https://www.twilio.com/docs/chat/rest/role-resource) for more info about roles. | [optional] 
**DefaultServiceRoleSid** | **string** | The service role assigned to users when they are added to the service. See the [Role resource](https://www.twilio.com/docs/chat/rest/role-resource) for more info about roles. | [optional] 
**FriendlyName** | **string** | A descriptive string that you create to describe the resource. | [optional] 
**LimitsChannelMembers** | **int32** | The maximum number of Members that can be added to Channels within this Service. Can be up to 1,000. | [optional] 
**LimitsUserChannels** | **int32** | The maximum number of Channels Users can be a Member of within this Service. Can be up to 1,000. | [optional] 
**MediaCompatibilityMessage** | **string** | The message to send when a media message has no text. Can be used as placeholder message. | [optional] 
**NotificationsAddedToChannelEnabled** | **bool** | Whether to send a notification when a member is added to a channel. The default is &#x60;false&#x60;. | [optional] 
**NotificationsAddedToChannelSound** | **string** | The name of the sound to play when a member is added to a channel and &#x60;notifications.added_to_channel.enabled&#x60; is &#x60;true&#x60;. | [optional] 
**NotificationsAddedToChannelTemplate** | **string** | The template to use to create the notification text displayed when a member is added to a channel and &#x60;notifications.added_to_channel.enabled&#x60; is &#x60;true&#x60;. | [optional] 
**NotificationsInvitedToChannelEnabled** | **bool** | Whether to send a notification when a user is invited to a channel. The default is &#x60;false&#x60;. | [optional] 
**NotificationsInvitedToChannelSound** | **string** | The name of the sound to play when a user is invited to a channel and &#x60;notifications.invited_to_channel.enabled&#x60; is &#x60;true&#x60;. | [optional] 
**NotificationsInvitedToChannelTemplate** | **string** | The template to use to create the notification text displayed when a user is invited to a channel and &#x60;notifications.invited_to_channel.enabled&#x60; is &#x60;true&#x60;. | [optional] 
**NotificationsLogEnabled** | **bool** | Whether to log notifications. The default is &#x60;false&#x60;. | [optional] 
**NotificationsNewMessageBadgeCountEnabled** | **bool** | Whether the new message badge is enabled. The default is &#x60;false&#x60;. | [optional] 
**NotificationsNewMessageEnabled** | **bool** | Whether to send a notification when a new message is added to a channel. The default is &#x60;false&#x60;. | [optional] 
**NotificationsNewMessageSound** | **string** | The name of the sound to play when a new message is added to a channel and &#x60;notifications.new_message.enabled&#x60; is &#x60;true&#x60;. | [optional] 
**NotificationsNewMessageTemplate** | **string** | The template to use to create the notification text displayed when a new message is added to a channel and &#x60;notifications.new_message.enabled&#x60; is &#x60;true&#x60;. | [optional] 
**NotificationsRemovedFromChannelEnabled** | **bool** | Whether to send a notification to a user when they are removed from a channel. The default is &#x60;false&#x60;. | [optional] 
**NotificationsRemovedFromChannelSound** | **string** | The name of the sound to play to a user when they are removed from a channel and &#x60;notifications.removed_from_channel.enabled&#x60; is &#x60;true&#x60;. | [optional] 
**NotificationsRemovedFromChannelTemplate** | **string** | The template to use to create the notification text displayed to a user when they are removed from a channel and &#x60;notifications.removed_from_channel.enabled&#x60; is &#x60;true&#x60;. | [optional] 
**PostWebhookRetryCount** | **int32** | The number of times to retry a call to the &#x60;post_webhook_url&#x60; if the request times out (after 5 seconds) or it receives a 429, 503, or 504 HTTP response. The default is 0, which means the call won&#39;t be retried. | [optional] 
**PostWebhookUrl** | **string** | The URL for post-event webhooks, which are called by using the &#x60;webhook_method&#x60;. See [Webhook Events](https://www.twilio.com/docs/chat/webhook-events) for more details. | [optional] 
**PreWebhookRetryCount** | **int32** | The number of times to retry a call to the &#x60;pre_webhook_url&#x60; if the request times out (after 5 seconds) or it receives a 429, 503, or 504 HTTP response. Default retry count is 0 times, which means the call won&#39;t be retried. | [optional] 
**PreWebhookUrl** | **string** | The URL for pre-event webhooks, which are called by using the &#x60;webhook_method&#x60;. See [Webhook Events](https://www.twilio.com/docs/chat/webhook-events) for more details. | [optional] 
**ReachabilityEnabled** | **bool** | Whether to enable the [Reachability Indicator](https://www.twilio.com/docs/chat/reachability-indicator) for this Service instance. The default is &#x60;false&#x60;. | [optional] 
**ReadStatusEnabled** | **bool** | Whether to enable the [Message Consumption Horizon](https://www.twilio.com/docs/chat/consumption-horizon) feature. The default is &#x60;true&#x60;. | [optional] 
**TypingIndicatorTimeout** | **int32** | How long in seconds after a &#x60;started typing&#x60; event until clients should assume that user is no longer typing, even if no &#x60;ended typing&#x60; message was received.  The default is 5 seconds. | [optional] 
**WebhookFilters** | **[]string** | The list of webhook events that are enabled for this Service instance. See [Webhook Events](https://www.twilio.com/docs/chat/webhook-events) for more details. | [optional] 
**WebhookMethod** | **string** | The HTTP method to use for calls to the &#x60;pre_webhook_url&#x60; and &#x60;post_webhook_url&#x60; webhooks.  Can be: &#x60;POST&#x60; or &#x60;GET&#x60; and the default is &#x60;POST&#x60;. See [Webhook Events](https://www.twilio.com/docs/chat/webhook-events) for more details. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


