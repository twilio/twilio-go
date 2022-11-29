# ServicesApi

All URIs are relative to *https://chat.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateService**](ServicesApi.md#CreateService) | **Post** /v1/Services | 
[**DeleteService**](ServicesApi.md#DeleteService) | **Delete** /v1/Services/{Sid} | 
[**FetchService**](ServicesApi.md#FetchService) | **Get** /v1/Services/{Sid} | 
[**ListService**](ServicesApi.md#ListService) | **Get** /v1/Services | 
[**UpdateService**](ServicesApi.md#UpdateService) | **Post** /v1/Services/{Sid} | 



## CreateService

> ChatV1Service CreateService(ctx, optional)





### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateServiceParams struct


Name | Type | Description
------------- | ------------- | -------------
**FriendlyName** | **string** | A descriptive string that you create to describe the resource. It can be up to 64 characters long.

### Return type

[**ChatV1Service**](ChatV1Service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteService

> DeleteService(ctx, Sid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Service resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteServiceParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchService

> ChatV1Service FetchService(ctx, Sid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Service resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchServiceParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ChatV1Service**](ChatV1Service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListService

> []ChatV1Service ListService(ctx, optional)





### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListServiceParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]ChatV1Service**](ChatV1Service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateService

> ChatV1Service UpdateService(ctx, Sidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Service resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateServiceParams struct


Name | Type | Description
------------- | ------------- | -------------
**FriendlyName** | **string** | A descriptive string that you create to describe the resource. It can be up to 64 characters long.
**DefaultServiceRoleSid** | **string** | The service role assigned to users when they are added to the service. See the [Roles endpoint](https://www.twilio.com/docs/chat/api/roles) for more details.
**DefaultChannelRoleSid** | **string** | The channel role assigned to users when they are added to a channel. See the [Roles endpoint](https://www.twilio.com/docs/chat/api/roles) for more details.
**DefaultChannelCreatorRoleSid** | **string** | The channel role assigned to a channel creator when they join a new channel. See the [Roles endpoint](https://www.twilio.com/docs/chat/api/roles) for more details.
**ReadStatusEnabled** | **bool** | Whether to enable the [Message Consumption Horizon](https://www.twilio.com/docs/chat/consumption-horizon) feature. The default is `true`.
**ReachabilityEnabled** | **bool** | Whether to enable the [Reachability Indicator](https://www.twilio.com/docs/chat/reachability-indicator) for this Service instance. The default is `false`.
**TypingIndicatorTimeout** | **int** | How long in seconds after a `started typing` event until clients should assume that user is no longer typing, even if no `ended typing` message was received.  The default is 5 seconds.
**ConsumptionReportInterval** | **int** | DEPRECATED. The interval in seconds between consumption reports submission batches from client endpoints.
**NotificationsNewMessageEnabled** | **bool** | Whether to send a notification when a new message is added to a channel. Can be: `true` or `false` and the default is `false`.
**NotificationsNewMessageTemplate** | **string** | The template to use to create the notification text displayed when a new message is added to a channel and `notifications.new_message.enabled` is `true`.
**NotificationsAddedToChannelEnabled** | **bool** | Whether to send a notification when a member is added to a channel. Can be: `true` or `false` and the default is `false`.
**NotificationsAddedToChannelTemplate** | **string** | The template to use to create the notification text displayed when a member is added to a channel and `notifications.added_to_channel.enabled` is `true`.
**NotificationsRemovedFromChannelEnabled** | **bool** | Whether to send a notification to a user when they are removed from a channel. Can be: `true` or `false` and the default is `false`.
**NotificationsRemovedFromChannelTemplate** | **string** | The template to use to create the notification text displayed to a user when they are removed from a channel and `notifications.removed_from_channel.enabled` is `true`.
**NotificationsInvitedToChannelEnabled** | **bool** | Whether to send a notification when a user is invited to a channel. Can be: `true` or `false` and the default is `false`.
**NotificationsInvitedToChannelTemplate** | **string** | The template to use to create the notification text displayed when a user is invited to a channel and `notifications.invited_to_channel.enabled` is `true`.
**PreWebhookUrl** | **string** | The URL for pre-event webhooks, which are called by using the `webhook_method`. See [Webhook Events](https://www.twilio.com/docs/api/chat/webhooks) for more details.
**PostWebhookUrl** | **string** | The URL for post-event webhooks, which are called by using the `webhook_method`. See [Webhook Events](https://www.twilio.com/docs/api/chat/webhooks) for more details.
**WebhookMethod** | **string** | The HTTP method to use for calls to the `pre_webhook_url` and `post_webhook_url` webhooks.  Can be: `POST` or `GET` and the default is `POST`. See [Webhook Events](https://www.twilio.com/docs/chat/webhook-events) for more details.
**WebhookFilters** | **[]string** | The list of WebHook events that are enabled for this Service instance. See [Webhook Events](https://www.twilio.com/docs/chat/webhook-events) for more details.
**WebhooksOnMessageSendUrl** | **string** | The URL of the webhook to call in response to the `on_message_send` event using the `webhooks.on_message_send.method` HTTP method.
**WebhooksOnMessageSendMethod** | **string** | The HTTP method to use when calling the `webhooks.on_message_send.url`.
**WebhooksOnMessageUpdateUrl** | **string** | The URL of the webhook to call in response to the `on_message_update` event using the `webhooks.on_message_update.method` HTTP method.
**WebhooksOnMessageUpdateMethod** | **string** | The HTTP method to use when calling the `webhooks.on_message_update.url`.
**WebhooksOnMessageRemoveUrl** | **string** | The URL of the webhook to call in response to the `on_message_remove` event using the `webhooks.on_message_remove.method` HTTP method.
**WebhooksOnMessageRemoveMethod** | **string** | The HTTP method to use when calling the `webhooks.on_message_remove.url`.
**WebhooksOnChannelAddUrl** | **string** | The URL of the webhook to call in response to the `on_channel_add` event using the `webhooks.on_channel_add.method` HTTP method.
**WebhooksOnChannelAddMethod** | **string** | The HTTP method to use when calling the `webhooks.on_channel_add.url`.
**WebhooksOnChannelDestroyUrl** | **string** | The URL of the webhook to call in response to the `on_channel_destroy` event using the `webhooks.on_channel_destroy.method` HTTP method.
**WebhooksOnChannelDestroyMethod** | **string** | The HTTP method to use when calling the `webhooks.on_channel_destroy.url`.
**WebhooksOnChannelUpdateUrl** | **string** | The URL of the webhook to call in response to the `on_channel_update` event using the `webhooks.on_channel_update.method` HTTP method.
**WebhooksOnChannelUpdateMethod** | **string** | The HTTP method to use when calling the `webhooks.on_channel_update.url`.
**WebhooksOnMemberAddUrl** | **string** | The URL of the webhook to call in response to the `on_member_add` event using the `webhooks.on_member_add.method` HTTP method.
**WebhooksOnMemberAddMethod** | **string** | The HTTP method to use when calling the `webhooks.on_member_add.url`.
**WebhooksOnMemberRemoveUrl** | **string** | The URL of the webhook to call in response to the `on_member_remove` event using the `webhooks.on_member_remove.method` HTTP method.
**WebhooksOnMemberRemoveMethod** | **string** | The HTTP method to use when calling the `webhooks.on_member_remove.url`.
**WebhooksOnMessageSentUrl** | **string** | The URL of the webhook to call in response to the `on_message_sent` event using the `webhooks.on_message_sent.method` HTTP method.
**WebhooksOnMessageSentMethod** | **string** | The URL of the webhook to call in response to the `on_message_sent` event`.
**WebhooksOnMessageUpdatedUrl** | **string** | The URL of the webhook to call in response to the `on_message_updated` event using the `webhooks.on_message_updated.method` HTTP method.
**WebhooksOnMessageUpdatedMethod** | **string** | The HTTP method to use when calling the `webhooks.on_message_updated.url`.
**WebhooksOnMessageRemovedUrl** | **string** | The URL of the webhook to call in response to the `on_message_removed` event using the `webhooks.on_message_removed.method` HTTP method.
**WebhooksOnMessageRemovedMethod** | **string** | The HTTP method to use when calling the `webhooks.on_message_removed.url`.
**WebhooksOnChannelAddedUrl** | **string** | The URL of the webhook to call in response to the `on_channel_added` event using the `webhooks.on_channel_added.method` HTTP method.
**WebhooksOnChannelAddedMethod** | **string** | The URL of the webhook to call in response to the `on_channel_added` event`.
**WebhooksOnChannelDestroyedUrl** | **string** | The URL of the webhook to call in response to the `on_channel_added` event using the `webhooks.on_channel_destroyed.method` HTTP method.
**WebhooksOnChannelDestroyedMethod** | **string** | The HTTP method to use when calling the `webhooks.on_channel_destroyed.url`.
**WebhooksOnChannelUpdatedUrl** | **string** | The URL of the webhook to call in response to the `on_channel_updated` event using the `webhooks.on_channel_updated.method` HTTP method.
**WebhooksOnChannelUpdatedMethod** | **string** | The HTTP method to use when calling the `webhooks.on_channel_updated.url`.
**WebhooksOnMemberAddedUrl** | **string** | The URL of the webhook to call in response to the `on_channel_updated` event using the `webhooks.on_channel_updated.method` HTTP method.
**WebhooksOnMemberAddedMethod** | **string** | The HTTP method to use when calling the `webhooks.on_channel_updated.url`.
**WebhooksOnMemberRemovedUrl** | **string** | The URL of the webhook to call in response to the `on_member_removed` event using the `webhooks.on_member_removed.method` HTTP method.
**WebhooksOnMemberRemovedMethod** | **string** | The HTTP method to use when calling the `webhooks.on_member_removed.url`.
**LimitsChannelMembers** | **int** | The maximum number of Members that can be added to Channels within this Service. Can be up to 1,000.
**LimitsUserChannels** | **int** | The maximum number of Channels Users can be a Member of within this Service. Can be up to 1,000.

### Return type

[**ChatV1Service**](ChatV1Service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

