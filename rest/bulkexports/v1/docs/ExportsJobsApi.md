# ExportsJobsApi

All URIs are relative to *https://bulkexports.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateExportCustomJob**](ExportsJobsApi.md#CreateExportCustomJob) | **Post** /v1/Exports/{ResourceType}/Jobs | 
[**DeleteJob**](ExportsJobsApi.md#DeleteJob) | **Delete** /v1/Exports/Jobs/{JobSid} | 
[**FetchJob**](ExportsJobsApi.md#FetchJob) | **Get** /v1/Exports/Jobs/{JobSid} | 
[**ListExportCustomJob**](ExportsJobsApi.md#ListExportCustomJob) | **Get** /v1/Exports/{ResourceType}/Jobs | 



## CreateExportCustomJob

> BulkexportsV1ExportCustomJob CreateExportCustomJob(ctx, ResourceTypeoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ResourceType** | **string** | The type of communication – Messages or Calls, Conferences, and Participants

### Other Parameters

Other parameters are passed through a pointer to a CreateExportCustomJobParams struct


Name | Type | Description
------------- | ------------- | -------------
**StartDay** | **string** | The start day for the custom export specified as a string in the format of yyyy-mm-dd
**EndDay** | **string** | The end day for the custom export specified as a string in the format of yyyy-mm-dd. End day is inclusive and must be 2 days earlier than the current UTC day.
**FriendlyName** | **string** | The friendly name specified when creating the job
**WebhookUrl** | **string** | The optional webhook url called on completion of the job. If this is supplied, `WebhookMethod` must also be supplied. If you set neither webhook nor email, you will have to check your job's status manually.
**WebhookMethod** | **string** | This is the method used to call the webhook on completion of the job. If this is supplied, `WebhookUrl` must also be supplied.
**Email** | **string** | The optional email to send the completion notification to. You can set both webhook, and email, or one or the other. If you set neither, the job will run but you will have to query to determine your job's status.

### Return type

[**BulkexportsV1ExportCustomJob**](BulkexportsV1ExportCustomJob.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteJob

> DeleteJob(ctx, JobSid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**JobSid** | **string** | The unique string that that we created to identify the Bulk Export job

### Other Parameters

Other parameters are passed through a pointer to a DeleteJobParams struct


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


## FetchJob

> BulkexportsV1Job FetchJob(ctx, JobSid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**JobSid** | **string** | The unique string that that we created to identify the Bulk Export job

### Other Parameters

Other parameters are passed through a pointer to a FetchJobParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**BulkexportsV1Job**](BulkexportsV1Job.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListExportCustomJob

> []BulkexportsV1ExportCustomJob ListExportCustomJob(ctx, ResourceTypeoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ResourceType** | **string** | The type of communication – Messages, Calls, Conferences, and Participants

### Other Parameters

Other parameters are passed through a pointer to a ListExportCustomJobParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]BulkexportsV1ExportCustomJob**](BulkexportsV1ExportCustomJob.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

