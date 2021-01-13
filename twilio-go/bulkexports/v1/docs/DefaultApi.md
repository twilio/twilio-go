# \DefaultApi

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

> BulkexportsV1ExportExportCustomJob CreateExportCustomJob(ctx, resourceType, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceType** | **string**| The type of communication – Messages or Calls | 
 **optional** | ***CreateExportCustomJobOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateExportCustomJobOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **email** | **optional.String**| The optional email to send the completion notification to. You can set both webhook, and email, or one or the other. If you set neither, the job will run but you will have to query to determine your job&#39;s status. | 
 **endDay** | **optional.String**| The end day for the custom export specified as a string in the format of yyyy-mm-dd. End day is inclusive and must be 2 days earlier than the current UTC day. | 
 **friendlyName** | **optional.String**| The friendly name specified when creating the job | 
 **startDay** | **optional.String**| The start day for the custom export specified as a string in the format of yyyy-mm-dd | 
 **webhookMethod** | **optional.String**| This is the method used to call the webhook on completion of the job. If this is supplied, &#x60;WebhookUrl&#x60; must also be supplied. | 
 **webhookUrl** | **optional.String**| The optional webhook url called on completion of the job. If this is supplied, &#x60;WebhookMethod&#x60; must also be supplied. If you set neither webhook nor email, you will have to check your job&#39;s status manually. | 

### Return type

[**BulkexportsV1ExportExportCustomJob**](bulkexports.v1.export.export_custom_job.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteJob

> DeleteJob(ctx, jobSid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobSid** | **string**| The unique string that that we created to identify the Bulk Export job | 

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

> FetchDay(ctx, resourceType, day)



Fetch a specific Day.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceType** | **string**| The type of communication – Messages, Calls | 
**day** | **string**| The ISO 8601 format date of the resources in the file, for a UTC day | 

### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchExport

> BulkexportsV1Export FetchExport(ctx, resourceType)



Fetch a specific Export.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceType** | **string**| The type of communication – Messages, Calls | 

### Return type

[**BulkexportsV1Export**](bulkexports.v1.export.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchExportConfiguration

> BulkexportsV1ExportConfiguration FetchExportConfiguration(ctx, resourceType)



Fetch a specific Export Configuration.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceType** | **string**| The type of communication – Messages, Calls | 

### Return type

[**BulkexportsV1ExportConfiguration**](bulkexports.v1.export_configuration.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchJob

> BulkexportsV1ExportJob FetchJob(ctx, jobSid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobSid** | **string**|  | 

### Return type

[**BulkexportsV1ExportJob**](bulkexports.v1.export.job.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDay

> BulkexportsV1ExportDayReadResponse ListDay(ctx, resourceType, optional)



Retrieve a list of all Days for a resource.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceType** | **string**| The type of communication – Messages, Calls | 
 **optional** | ***ListDayOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListDayOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**BulkexportsV1ExportDayReadResponse**](bulkexports_v1_export_dayReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListExportCustomJob

> BulkexportsV1ExportExportCustomJobReadResponse ListExportCustomJob(ctx, resourceType, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceType** | **string**| The type of communication – Messages, Calls | 
 **optional** | ***ListExportCustomJobOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListExportCustomJobOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**BulkexportsV1ExportExportCustomJobReadResponse**](bulkexports_v1_export_export_custom_jobReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateExportConfiguration

> BulkexportsV1ExportConfiguration UpdateExportConfiguration(ctx, resourceType, optional)



Update a specific Export Configuration.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceType** | **string**| The type of communication – Messages, Calls | 
 **optional** | ***UpdateExportConfigurationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateExportConfigurationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **enabled** | **optional.Bool**| If true, Twilio will automatically generate every day&#39;s file when the day is over. | 
 **webhookMethod** | **optional.String**| Sets whether Twilio should call a webhook URL when the automatic generation is complete, using GET or POST. The actual destination is set in the webhook_url | 
 **webhookUrl** | **optional.String**| Stores the URL destination for the method specified in webhook_method. | 

### Return type

[**BulkexportsV1ExportConfiguration**](bulkexports.v1.export_configuration.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

