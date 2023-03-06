# ObjectsContactRecordsApi

All URIs are relative to *https://context.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchRecord**](ObjectsContactRecordsApi.md#FetchRecord) | **Get** /v1/Objects/Contact/Records/{Sid} | 
[**ListRecord**](ObjectsContactRecordsApi.md#ListRecord) | **Get** /v1/Objects/Contact/Records | 



## FetchRecord

> ContextV1Record FetchRecord(ctx, Sid)



Retrieves the Record instance for an Object from your Account.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | A 34 character string that uniquely identifies the record for an object.

### Other Parameters

Other parameters are passed through a pointer to a FetchRecordParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ContextV1Record**](ContextV1Record.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRecord

> []ContextV1Record ListRecord(ctx, optional)



Retrieve a paginated list of records belonging to the account used to make the request.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListRecordParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]ContextV1Record**](ContextV1Record.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

