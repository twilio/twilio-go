# AssistantsWebhooksApi

All URIs are relative to *https://autopilot.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWebhook**](AssistantsWebhooksApi.md#CreateWebhook) | **Post** /v1/Assistants/{AssistantSid}/Webhooks | 
[**DeleteWebhook**](AssistantsWebhooksApi.md#DeleteWebhook) | **Delete** /v1/Assistants/{AssistantSid}/Webhooks/{Sid} | 
[**FetchWebhook**](AssistantsWebhooksApi.md#FetchWebhook) | **Get** /v1/Assistants/{AssistantSid}/Webhooks/{Sid} | 
[**ListWebhook**](AssistantsWebhooksApi.md#ListWebhook) | **Get** /v1/Assistants/{AssistantSid}/Webhooks | 
[**UpdateWebhook**](AssistantsWebhooksApi.md#UpdateWebhook) | **Post** /v1/Assistants/{AssistantSid}/Webhooks/{Sid} | 



## CreateWebhook

> AutopilotV1Webhook CreateWebhook(ctx, AssistantSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the new resource.

### Other Parameters

Other parameters are passed through a pointer to a CreateWebhookParams struct


Name | Type | Description
------------- | ------------- | -------------
**UniqueName** | **string** | An application-defined string that uniquely identifies the new resource. It can be used as an alternative to the `sid` in the URL path to address the resource. This value must be unique and 64 characters or less in length.
**Events** | **string** | The list of space-separated events that this Webhook will subscribe to.
**WebhookUrl** | **string** | The URL associated with this Webhook.
**WebhookMethod** | **string** | The method to be used when calling the webhook's URL.

### Return type

[**AutopilotV1Webhook**](AutopilotV1Webhook.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWebhook

> DeleteWebhook(ctx, AssistantSidSid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resources to delete.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Webhook resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteWebhookParams struct


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


## FetchWebhook

> AutopilotV1Webhook FetchWebhook(ctx, AssistantSidSid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resource to fetch.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Webhook resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchWebhookParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**AutopilotV1Webhook**](AutopilotV1Webhook.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWebhook

> []AutopilotV1Webhook ListWebhook(ctx, AssistantSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListWebhookParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]AutopilotV1Webhook**](AutopilotV1Webhook.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWebhook

> AutopilotV1Webhook UpdateWebhook(ctx, AssistantSidSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resource to update.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Webhook resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateWebhookParams struct


Name | Type | Description
------------- | ------------- | -------------
**UniqueName** | **string** | An application-defined string that uniquely identifies the new resource. It can be used as an alternative to the `sid` in the URL path to address the resource. This value must be unique and 64 characters or less in length.
**Events** | **string** | The list of space-separated events that this Webhook will subscribe to.
**WebhookUrl** | **string** | The URL associated with this Webhook.
**WebhookMethod** | **string** | The method to be used when calling the webhook's URL.

### Return type

[**AutopilotV1Webhook**](AutopilotV1Webhook.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

