# ExportsConfigurationApi

All URIs are relative to *https://bulkexports.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchExportConfiguration**](ExportsConfigurationApi.md#FetchExportConfiguration) | **Get** /v1/Exports/{ResourceType}/Configuration | 
[**UpdateExportConfiguration**](ExportsConfigurationApi.md#UpdateExportConfiguration) | **Post** /v1/Exports/{ResourceType}/Configuration | 



## FetchExportConfiguration

> BulkexportsV1ExportConfiguration FetchExportConfiguration(ctx, ResourceType)



Fetch a specific Export Configuration.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ResourceType** | **string** | The type of communication – Messages, Calls, Conferences, and Participants

### Other Parameters

Other parameters are passed through a pointer to a FetchExportConfigurationParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**BulkexportsV1ExportConfiguration**](BulkexportsV1ExportConfiguration.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateExportConfiguration

> BulkexportsV1ExportConfiguration UpdateExportConfiguration(ctx, ResourceTypeoptional)



Update a specific Export Configuration.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ResourceType** | **string** | The type of communication – Messages, Calls, Conferences, and Participants

### Other Parameters

Other parameters are passed through a pointer to a UpdateExportConfigurationParams struct


Name | Type | Description
------------- | ------------- | -------------
**Enabled** | **bool** | If true, Twilio will automatically generate every day's file when the day is over.
**WebhookUrl** | **string** | Stores the URL destination for the method specified in webhook_method.
**WebhookMethod** | **string** | Sets whether Twilio should call a webhook URL when the automatic generation is complete, using GET or POST. The actual destination is set in the webhook_url

### Return type

[**BulkexportsV1ExportConfiguration**](BulkexportsV1ExportConfiguration.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

