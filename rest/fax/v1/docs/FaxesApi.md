# FaxesApi

All URIs are relative to *https://fax.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFax**](FaxesApi.md#DeleteFax) | **Delete** /v1/Faxes/{Sid} | 
[**FetchFax**](FaxesApi.md#FetchFax) | **Get** /v1/Faxes/{Sid} | 
[**ListFax**](FaxesApi.md#ListFax) | **Get** /v1/Faxes | 



## DeleteFax

> DeleteFax(ctx, Sid)



Delete a specific fax and its associated media.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Fax resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteFaxParams struct


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


## FetchFax

> FaxV1Fax FetchFax(ctx, Sid)



Fetch a specific fax.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Fax resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchFaxParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**FaxV1Fax**](FaxV1Fax.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFax

> []FaxV1Fax ListFax(ctx, optional)



Retrieve a list of all faxes.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListFaxParams struct


Name | Type | Description
------------- | ------------- | -------------
**From** | **string** | Retrieve only those faxes sent from this phone number, specified in [E.164](https://www.twilio.com/docs/glossary/what-e164) format.
**To** | **string** | Retrieve only those faxes sent to this phone number, specified in [E.164](https://www.twilio.com/docs/glossary/what-e164) format.
**DateCreatedOnOrBefore** | **time.Time** | Retrieve only those faxes with a &#x60;date_created&#x60; that is before or equal to this value, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
**DateCreatedAfter** | **time.Time** | Retrieve only those faxes with a &#x60;date_created&#x60; that is later than this value, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]FaxV1Fax**](FaxV1Fax.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

