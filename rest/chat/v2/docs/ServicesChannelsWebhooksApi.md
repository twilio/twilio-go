# ServicesChannelsWebhooksApi

All URIs are relative to *https://chat.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateChannelWebhook**](ServicesChannelsWebhooksApi.md#CreateChannelWebhook) | **Post** /v2/Services/{ServiceSid}/Channels/{ChannelSid}/Webhooks | 
[**DeleteChannelWebhook**](ServicesChannelsWebhooksApi.md#DeleteChannelWebhook) | **Delete** /v2/Services/{ServiceSid}/Channels/{ChannelSid}/Webhooks/{Sid} | 
[**FetchChannelWebhook**](ServicesChannelsWebhooksApi.md#FetchChannelWebhook) | **Get** /v2/Services/{ServiceSid}/Channels/{ChannelSid}/Webhooks/{Sid} | 
[**ListChannelWebhook**](ServicesChannelsWebhooksApi.md#ListChannelWebhook) | **Get** /v2/Services/{ServiceSid}/Channels/{ChannelSid}/Webhooks | 
[**UpdateChannelWebhook**](ServicesChannelsWebhooksApi.md#UpdateChannelWebhook) | **Post** /v2/Services/{ServiceSid}/Channels/{ChannelSid}/Webhooks/{Sid} | 



## CreateChannelWebhook

> ChatV2ChannelWebhook CreateChannelWebhook(ctx, ServiceSidChannelSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) with the Channel to create the Webhook resource under.
**ChannelSid** | **string** | The SID of the [Channel](https://www.twilio.com/docs/chat/channels) the new Channel Webhook resource belongs to. This value can be the Channel resource's `sid` or `unique_name`.

### Other Parameters

Other parameters are passed through a pointer to a CreateChannelWebhookParams struct


Name | Type | Description
------------- | ------------- | -------------
**Type** | **string** | 
**ConfigurationUrl** | **string** | The URL of the webhook to call using the `configuration.method`.
**ConfigurationMethod** | **string** | 
**ConfigurationFilters** | **[]string** | The events that cause us to call the Channel Webhook. Used when `type` is `webhook`. This parameter takes only one event. To specify more than one event, repeat this parameter for each event. For the list of possible events, see [Webhook Event Triggers](https://www.twilio.com/docs/chat/webhook-events#webhook-event-trigger).
**ConfigurationTriggers** | **[]string** | A string that will cause us to call the webhook when it is present in a message body. This parameter takes only one trigger string. To specify more than one, repeat this parameter for each trigger string up to a total of 5 trigger strings. Used only when `type` = `trigger`.
**ConfigurationFlowSid** | **string** | The SID of the Studio [Flow](https://www.twilio.com/docs/studio/rest-api/flow) to call when an event in `configuration.filters` occurs. Used only when `type` is `studio`.
**ConfigurationRetryCount** | **int** | The number of times to retry the webhook if the first attempt fails. Can be an integer between 0 and 3, inclusive, and the default is 0.

### Return type

[**ChatV2ChannelWebhook**](ChatV2ChannelWebhook.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteChannelWebhook

> DeleteChannelWebhook(ctx, ServiceSidChannelSidSid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) with the Channel to delete the Webhook resource from.
**ChannelSid** | **string** | The SID of the [Channel](https://www.twilio.com/docs/chat/channels) the Channel Webhook resource to delete belongs to. This value can be the Channel resource's `sid` or `unique_name`.
**Sid** | **string** | The SID of the Channel Webhook resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteChannelWebhookParams struct


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


## FetchChannelWebhook

> ChatV2ChannelWebhook FetchChannelWebhook(ctx, ServiceSidChannelSidSid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) with the Channel to fetch the Webhook resource from.
**ChannelSid** | **string** | The SID of the [Channel](https://www.twilio.com/docs/chat/channels) the Channel Webhook resource to fetch belongs to. This value can be the Channel resource's `sid` or `unique_name`.
**Sid** | **string** | The SID of the Channel Webhook resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchChannelWebhookParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ChatV2ChannelWebhook**](ChatV2ChannelWebhook.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListChannelWebhook

> []ChatV2ChannelWebhook ListChannelWebhook(ctx, ServiceSidChannelSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) with the Channel to read the resources from.
**ChannelSid** | **string** | The SID of the [Channel](https://www.twilio.com/docs/chat/channels) the Channel Webhook resources to read belong to. This value can be the Channel resource's `sid` or `unique_name`.

### Other Parameters

Other parameters are passed through a pointer to a ListChannelWebhookParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]ChatV2ChannelWebhook**](ChatV2ChannelWebhook.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateChannelWebhook

> ChatV2ChannelWebhook UpdateChannelWebhook(ctx, ServiceSidChannelSidSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) with the Channel that has the Webhook resource to update.
**ChannelSid** | **string** | The SID of the [Channel](https://www.twilio.com/docs/chat/channels) the Channel Webhook resource to update belongs to. This value can be the Channel resource's `sid` or `unique_name`.
**Sid** | **string** | The SID of the Channel Webhook resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateChannelWebhookParams struct


Name | Type | Description
------------- | ------------- | -------------
**ConfigurationUrl** | **string** | The URL of the webhook to call using the `configuration.method`.
**ConfigurationMethod** | **string** | 
**ConfigurationFilters** | **[]string** | The events that cause us to call the Channel Webhook. Used when `type` is `webhook`. This parameter takes only one event. To specify more than one event, repeat this parameter for each event. For the list of possible events, see [Webhook Event Triggers](https://www.twilio.com/docs/chat/webhook-events#webhook-event-trigger).
**ConfigurationTriggers** | **[]string** | A string that will cause us to call the webhook when it is present in a message body. This parameter takes only one trigger string. To specify more than one, repeat this parameter for each trigger string up to a total of 5 trigger strings. Used only when `type` = `trigger`.
**ConfigurationFlowSid** | **string** | The SID of the Studio [Flow](https://www.twilio.com/docs/studio/rest-api/flow) to call when an event in `configuration.filters` occurs. Used only when `type` = `studio`.
**ConfigurationRetryCount** | **int** | The number of times to retry the webhook if the first attempt fails. Can be an integer between 0 and 3, inclusive, and the default is 0.

### Return type

[**ChatV2ChannelWebhook**](ChatV2ChannelWebhook.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

