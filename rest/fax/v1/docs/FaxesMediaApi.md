# FaxesMediaApi

All URIs are relative to *https://fax.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFaxMedia**](FaxesMediaApi.md#DeleteFaxMedia) | **Delete** /v1/Faxes/{FaxSid}/Media/{Sid} | 
[**FetchFaxMedia**](FaxesMediaApi.md#FetchFaxMedia) | **Get** /v1/Faxes/{FaxSid}/Media/{Sid} | 
[**ListFaxMedia**](FaxesMediaApi.md#ListFaxMedia) | **Get** /v1/Faxes/{FaxSid}/Media | 



## DeleteFaxMedia

> DeleteFaxMedia(ctx, FaxSidSid)



Delete a specific fax media instance.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FaxSid** | **string** | The SID of the fax with the FaxMedia resource to delete.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the FaxMedia resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteFaxMediaParams struct


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


## FetchFaxMedia

> FaxV1FaxFaxMedia FetchFaxMedia(ctx, FaxSidSid)



Fetch a specific fax media instance.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FaxSid** | **string** | The SID of the fax with the FaxMedia resource to fetch.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the FaxMedia resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchFaxMediaParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**FaxV1FaxFaxMedia**](FaxV1FaxFaxMedia.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFaxMedia

> ListFaxMediaResponse ListFaxMedia(ctx, FaxSidoptional)



Retrieve a list of all fax media instances for the specified fax.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FaxSid** | **string** | The SID of the fax with the FaxMedia resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListFaxMediaParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**ListFaxMediaResponse**](ListFaxMediaResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

