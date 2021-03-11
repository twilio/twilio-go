# DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateExportCustomJob**](DefaultApi.md#CreateExportCustomJob) | **Post** /v1/Exports/{ResourceType}/Jobs | 
[**DeleteJob**](DefaultApi.md#DeleteJob) | **Delete** /v1/Exports/Jobs/{JobSid} | 
[**FetchDay**](DefaultApi.md#FetchDay) | **Get** /v1/Exports/{ResourceType}/Days/{Day} | 
[**FetchExport**](DefaultApi.md#FetchExport) | **Get** /v1/Exports/{ResourceType} | 
[**FetchExportConfiguration**](DefaultApi.md#FetchExportConfiguration) | **Get** /v1/Exports/{ResourceType}/Configuration | 
[**FetchJob**](DefaultApi.md#FetchJob) | **Get** /v1/Exports/Jobs/{JobSid} | 
[**ListDay**](DefaultApi.md#ListDay) | **Get** /v1/Exports/{ResourceType}/Days | 
[**ListExportCustomJob**](DefaultApi.md#ListExportCustomJob) | **Get** /v1/Exports/{ResourceType}/Jobs | 
[**UpdateExportConfiguration**](DefaultApi.md#UpdateExportConfiguration) | **Post** /v1/Exports/{ResourceType}/Configuration | 



## CreateExportCustomJob

> BulkexportsV1ExportExportCustomJob CreateExportCustomJob(ctx, ResourceTypeoptional)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ResourceType** | **string** | The type of communication – Messages or Calls, Conferences, and Participants | 

### Other Parameters

Other parameters are passed through a pointer to a CreateExportCustomJobParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Email** | **string** | The optional email to send the completion notification to. You can set both webhook, and email, or one or the other. If you set neither, the job will run but you will have to query to determine your job&#39;s status. | 
**EndDay** | **string** | The end day for the custom export specified as a string in the format of yyyy-mm-dd. End day is inclusive and must be 2 days earlier than the current UTC day. | 
**FriendlyName** | **string** | The friendly name specified when creating the job | 
**StartDay** | **string** | The start day for the custom export specified as a string in the format of yyyy-mm-dd | 
**WebhookMethod** | **string** | This is the method used to call the webhook on completion of the job. If this is supplied, &#x60;WebhookUrl&#x60; must also be supplied. | 
**WebhookUrl** | **string** | The optional webhook url called on completion of the job. If this is supplied, &#x60;WebhookMethod&#x60; must also be supplied. If you set neither webhook nor email, you will have to check your job&#39;s status manually. | 

### Return type

[**BulkexportsV1ExportExportCustomJob**](BulkexportsV1ExportExportCustomJob.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteJob

> DeleteJob(ctx, JobSid)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**JobSid** | **string** | The unique string that that we created to identify the Bulk Export job | 

### Other Parameters

Other parameters are passed through a pointer to a DeleteJobParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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


## FetchDay

> FetchDay(ctx, ResourceTypeDay)



Fetch a specific Day.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ResourceType** | **string** | The type of communication – Messages, Calls, Conferences, and Participants | 
**Day** | **string** | The ISO 8601 format date of the resources in the file, for a UTC day | 

### Other Parameters

Other parameters are passed through a pointer to a FetchDayParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchExport

> BulkexportsV1Export FetchExport(ctx, ResourceType)



Fetch a specific Export.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ResourceType** | **string** | The type of communication – Messages, Calls, Conferences, and Participants | 

### Other Parameters

Other parameters are passed through a pointer to a FetchExportParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**BulkexportsV1Export**](BulkexportsV1Export.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchExportConfiguration

> BulkexportsV1ExportConfiguration FetchExportConfiguration(ctx, ResourceType)



Fetch a specific Export Configuration.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ResourceType** | **string** | The type of communication – Messages, Calls, Conferences, and Participants | 

### Other Parameters

Other parameters are passed through a pointer to a FetchExportConfigurationParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**BulkexportsV1ExportConfiguration**](BulkexportsV1ExportConfiguration.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchJob

> BulkexportsV1ExportJob FetchJob(ctx, JobSid)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**JobSid** | **string** | The unique string that that we created to identify the Bulk Export job | 

### Other Parameters

Other parameters are passed through a pointer to a FetchJobParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**BulkexportsV1ExportJob**](BulkexportsV1ExportJob.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDay

> ListDayResponse ListDay(ctx, ResourceTypeoptional)



Retrieve a list of all Days for a resource.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ResourceType** | **string** | The type of communication – Messages, Calls, Conferences, and Participants | 

### Other Parameters

Other parameters are passed through a pointer to a ListDayParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListDayResponse**](ListDayResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListExportCustomJob

> ListExportCustomJobResponse ListExportCustomJob(ctx, ResourceTypeoptional)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ResourceType** | **string** | The type of communication – Messages, Calls, Conferences, and Participants | 

### Other Parameters

Other parameters are passed through a pointer to a ListExportCustomJobParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListExportCustomJobResponse**](ListExportCustomJobResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateExportConfiguration

> BulkexportsV1ExportConfiguration UpdateExportConfiguration(ctx, ResourceTypeoptional)



Update a specific Export Configuration.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ResourceType** | **string** | The type of communication – Messages, Calls, Conferences, and Participants | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateExportConfigurationParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Enabled** | **bool** | If true, Twilio will automatically generate every day&#39;s file when the day is over. | 
**WebhookMethod** | **string** | Sets whether Twilio should call a webhook URL when the automatic generation is complete, using GET or POST. The actual destination is set in the webhook_url | 
**WebhookUrl** | **string** | Stores the URL destination for the method specified in webhook_method. | 

### Return type

[**BulkexportsV1ExportConfiguration**](BulkexportsV1ExportConfiguration.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

