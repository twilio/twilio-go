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

> ListServiceResponse ListService(ctx, optional)



### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListServiceParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListServiceResponse**](ListServiceResponse.md)

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
**ConsumptionReportInterval** | **int** | DEPRECATED. The interval in seconds between consumption reports submission batches from client endpoints.
**DefaultChannelCreatorRoleSid** | **string** | The channel role assigned to a channel creator when they join a new channel. See the [Roles endpoint](https://www.twilio.com/docs/chat/api/roles) for more details.
**DefaultChannelRoleSid** | **string** | The channel role assigned to users when they are added to a channel. See the [Roles endpoint](https://www.twilio.com/docs/chat/api/roles) for more details.
**DefaultServiceRoleSid** | **string** | The service role assigned to users when they are added to the service. See the [Roles endpoint](https://www.twilio.com/docs/chat/api/roles) for more details.
**FriendlyName** | **string** | A descriptive string that you create to describe the resource. It can be up to 64 characters long.
**LimitsChannelMembers** | **int** | The maximum number of Members that can be added to Channels within this Service. Can be up to 1,000.
**LimitsUserChannels** | **int** | The maximum number of Channels Users can be a Member of within this Service. Can be up to 1,000.
**NotificationsAddedToChannelEnabled** | **bool** | Whether to send a notification when a member is added to a channel. Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;.
**NotificationsAddedToChannelTemplate** | **string** | The template to use to create the notification text displayed when a member is added to a channel and &#x60;notifications.added_to_channel.enabled&#x60; is &#x60;true&#x60;.
**NotificationsInvitedToChannelEnabled** | **bool** | Whether to send a notification when a user is invited to a channel. Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;.
**NotificationsInvitedToChannelTemplate** | **string** | The template to use to create the notification text displayed when a user is invited to a channel and &#x60;notifications.invited_to_channel.enabled&#x60; is &#x60;true&#x60;.
**NotificationsNewMessageEnabled** | **bool** | Whether to send a notification when a new message is added to a channel. Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;.
**NotificationsNewMessageTemplate** | **string** | The template to use to create the notification text displayed when a new message is added to a channel and &#x60;notifications.new_message.enabled&#x60; is &#x60;true&#x60;.
**NotificationsRemovedFromChannelEnabled** | **bool** | Whether to send a notification to a user when they are removed from a channel. Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;.
**NotificationsRemovedFromChannelTemplate** | **string** | The template to use to create the notification text displayed to a user when they are removed from a channel and &#x60;notifications.removed_from_channel.enabled&#x60; is &#x60;true&#x60;.
**PostWebhookUrl** | **string** | The URL for post-event webhooks, which are called by using the &#x60;webhook_method&#x60;. See [Webhook Events](https://www.twilio.com/docs/api/chat/webhooks) for more details.
**PreWebhookUrl** | **string** | The URL for pre-event webhooks, which are called by using the &#x60;webhook_method&#x60;. See [Webhook Events](https://www.twilio.com/docs/api/chat/webhooks) for more details.
**ReachabilityEnabled** | **bool** | Whether to enable the [Reachability Indicator](https://www.twilio.com/docs/chat/reachability-indicator) for this Service instance. The default is &#x60;false&#x60;.
**ReadStatusEnabled** | **bool** | Whether to enable the [Message Consumption Horizon](https://www.twilio.com/docs/chat/consumption-horizon) feature. The default is &#x60;true&#x60;.
**TypingIndicatorTimeout** | **int** | How long in seconds after a &#x60;started typing&#x60; event until clients should assume that user is no longer typing, even if no &#x60;ended typing&#x60; message was received.  The default is 5 seconds.
**WebhookFilters** | **[]string** | The list of WebHook events that are enabled for this Service instance. See [Webhook Events](https://www.twilio.com/docs/chat/webhook-events) for more details.
**WebhookMethod** | **string** | The HTTP method to use for calls to the &#x60;pre_webhook_url&#x60; and &#x60;post_webhook_url&#x60; webhooks.  Can be: &#x60;POST&#x60; or &#x60;GET&#x60; and the default is &#x60;POST&#x60;. See [Webhook Events](https://www.twilio.com/docs/chat/webhook-events) for more details.
**WebhooksOnChannelAddMethod** | **string** | The HTTP method to use when calling the &#x60;webhooks.on_channel_add.url&#x60;.
**WebhooksOnChannelAddUrl** | **string** | The URL of the webhook to call in response to the &#x60;on_channel_add&#x60; event using the &#x60;webhooks.on_channel_add.method&#x60; HTTP method.
**WebhooksOnChannelAddedMethod** | **string** | The URL of the webhook to call in response to the &#x60;on_channel_added&#x60; event&#x60;.
**WebhooksOnChannelAddedUrl** | **string** | The URL of the webhook to call in response to the &#x60;on_channel_added&#x60; event using the &#x60;webhooks.on_channel_added.method&#x60; HTTP method.
**WebhooksOnChannelDestroyMethod** | **string** | The HTTP method to use when calling the &#x60;webhooks.on_channel_destroy.url&#x60;.
**WebhooksOnChannelDestroyUrl** | **string** | The URL of the webhook to call in response to the &#x60;on_channel_destroy&#x60; event using the &#x60;webhooks.on_channel_destroy.method&#x60; HTTP method.
**WebhooksOnChannelDestroyedMethod** | **string** | The HTTP method to use when calling the &#x60;webhooks.on_channel_destroyed.url&#x60;.
**WebhooksOnChannelDestroyedUrl** | **string** | The URL of the webhook to call in response to the &#x60;on_channel_added&#x60; event using the &#x60;webhooks.on_channel_destroyed.method&#x60; HTTP method.
**WebhooksOnChannelUpdateMethod** | **string** | The HTTP method to use when calling the &#x60;webhooks.on_channel_update.url&#x60;.
**WebhooksOnChannelUpdateUrl** | **string** | The URL of the webhook to call in response to the &#x60;on_channel_update&#x60; event using the &#x60;webhooks.on_channel_update.method&#x60; HTTP method.
**WebhooksOnChannelUpdatedMethod** | **string** | The HTTP method to use when calling the &#x60;webhooks.on_channel_updated.url&#x60;.
**WebhooksOnChannelUpdatedUrl** | **string** | The URL of the webhook to call in response to the &#x60;on_channel_updated&#x60; event using the &#x60;webhooks.on_channel_updated.method&#x60; HTTP method.
**WebhooksOnMemberAddMethod** | **string** | The HTTP method to use when calling the &#x60;webhooks.on_member_add.url&#x60;.
**WebhooksOnMemberAddUrl** | **string** | The URL of the webhook to call in response to the &#x60;on_member_add&#x60; event using the &#x60;webhooks.on_member_add.method&#x60; HTTP method.
**WebhooksOnMemberAddedMethod** | **string** | The HTTP method to use when calling the &#x60;webhooks.on_channel_updated.url&#x60;.
**WebhooksOnMemberAddedUrl** | **string** | The URL of the webhook to call in response to the &#x60;on_channel_updated&#x60; event using the &#x60;webhooks.on_channel_updated.method&#x60; HTTP method.
**WebhooksOnMemberRemoveMethod** | **string** | The HTTP method to use when calling the &#x60;webhooks.on_member_remove.url&#x60;.
**WebhooksOnMemberRemoveUrl** | **string** | The URL of the webhook to call in response to the &#x60;on_member_remove&#x60; event using the &#x60;webhooks.on_member_remove.method&#x60; HTTP method.
**WebhooksOnMemberRemovedMethod** | **string** | The HTTP method to use when calling the &#x60;webhooks.on_member_removed.url&#x60;.
**WebhooksOnMemberRemovedUrl** | **string** | The URL of the webhook to call in response to the &#x60;on_member_removed&#x60; event using the &#x60;webhooks.on_member_removed.method&#x60; HTTP method.
**WebhooksOnMessageRemoveMethod** | **string** | The HTTP method to use when calling the &#x60;webhooks.on_message_remove.url&#x60;.
**WebhooksOnMessageRemoveUrl** | **string** | The URL of the webhook to call in response to the &#x60;on_message_remove&#x60; event using the &#x60;webhooks.on_message_remove.method&#x60; HTTP method.
**WebhooksOnMessageRemovedMethod** | **string** | The HTTP method to use when calling the &#x60;webhooks.on_message_removed.url&#x60;.
**WebhooksOnMessageRemovedUrl** | **string** | The URL of the webhook to call in response to the &#x60;on_message_removed&#x60; event using the &#x60;webhooks.on_message_removed.method&#x60; HTTP method.
**WebhooksOnMessageSendMethod** | **string** | The HTTP method to use when calling the &#x60;webhooks.on_message_send.url&#x60;.
**WebhooksOnMessageSendUrl** | **string** | The URL of the webhook to call in response to the &#x60;on_message_send&#x60; event using the &#x60;webhooks.on_message_send.method&#x60; HTTP method.
**WebhooksOnMessageSentMethod** | **string** | The URL of the webhook to call in response to the &#x60;on_message_sent&#x60; event&#x60;.
**WebhooksOnMessageSentUrl** | **string** | The URL of the webhook to call in response to the &#x60;on_message_sent&#x60; event using the &#x60;webhooks.on_message_sent.method&#x60; HTTP method.
**WebhooksOnMessageUpdateMethod** | **string** | The HTTP method to use when calling the &#x60;webhooks.on_message_update.url&#x60;.
**WebhooksOnMessageUpdateUrl** | **string** | The URL of the webhook to call in response to the &#x60;on_message_update&#x60; event using the &#x60;webhooks.on_message_update.method&#x60; HTTP method.
**WebhooksOnMessageUpdatedMethod** | **string** | The HTTP method to use when calling the &#x60;webhooks.on_message_updated.url&#x60;.
**WebhooksOnMessageUpdatedUrl** | **string** | The URL of the webhook to call in response to the &#x60;on_message_updated&#x60; event using the &#x60;webhooks.on_message_updated.method&#x60; HTTP method.

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

