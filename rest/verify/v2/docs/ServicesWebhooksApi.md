# ServicesWebhooksApi

All URIs are relative to *https://verify.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWebhook**](ServicesWebhooksApi.md#CreateWebhook) | **Post** /v2/Services/{ServiceSid}/Webhooks | 
[**DeleteWebhook**](ServicesWebhooksApi.md#DeleteWebhook) | **Delete** /v2/Services/{ServiceSid}/Webhooks/{Sid} | 
[**FetchWebhook**](ServicesWebhooksApi.md#FetchWebhook) | **Get** /v2/Services/{ServiceSid}/Webhooks/{Sid} | 
[**ListWebhook**](ServicesWebhooksApi.md#ListWebhook) | **Get** /v2/Services/{ServiceSid}/Webhooks | 
[**UpdateWebhook**](ServicesWebhooksApi.md#UpdateWebhook) | **Post** /v2/Services/{ServiceSid}/Webhooks/{Sid} | 



## CreateWebhook

> VerifyV2Webhook CreateWebhook(ctx, ServiceSidoptional)



Create a new Webhook for the Service

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The unique SID identifier of the Service.

### Other Parameters

Other parameters are passed through a pointer to a CreateWebhookParams struct


Name | Type | Description
------------- | ------------- | -------------
**FriendlyName** | **string** | The string that you assigned to describe the webhook. **This value should not contain PII.**
**EventTypes** | **[]string** | The array of events that this Webhook is subscribed to. Possible event types: `*, factor.deleted, factor.created, factor.verified, challenge.approved, challenge.denied`
**WebhookUrl** | **string** | The URL associated with this Webhook.
**Status** | **string** | 
**Version** | **string** | 

### Return type

[**VerifyV2Webhook**](VerifyV2Webhook.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWebhook

> DeleteWebhook(ctx, ServiceSidSid)



Delete a specific Webhook.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The unique SID identifier of the Service.
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

> VerifyV2Webhook FetchWebhook(ctx, ServiceSidSid)



Fetch a specific Webhook.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The unique SID identifier of the Service.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Webhook resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchWebhookParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**VerifyV2Webhook**](VerifyV2Webhook.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWebhook

> []VerifyV2Webhook ListWebhook(ctx, ServiceSidoptional)



Retrieve a list of all Webhooks for a Service.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The unique SID identifier of the Service.

### Other Parameters

Other parameters are passed through a pointer to a ListWebhookParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]VerifyV2Webhook**](VerifyV2Webhook.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWebhook

> VerifyV2Webhook UpdateWebhook(ctx, ServiceSidSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The unique SID identifier of the Service.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Webhook resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateWebhookParams struct


Name | Type | Description
------------- | ------------- | -------------
**FriendlyName** | **string** | The string that you assigned to describe the webhook. **This value should not contain PII.**
**EventTypes** | **[]string** | The array of events that this Webhook is subscribed to. Possible event types: `*, factor.deleted, factor.created, factor.verified, challenge.approved, challenge.denied`
**WebhookUrl** | **string** | The URL associated with this Webhook.
**Status** | **string** | 
**Version** | **string** | 

### Return type

[**VerifyV2Webhook**](VerifyV2Webhook.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

