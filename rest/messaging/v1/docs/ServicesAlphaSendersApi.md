# ServicesAlphaSendersApi

All URIs are relative to *https://messaging.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAlphaSender**](ServicesAlphaSendersApi.md#CreateAlphaSender) | **Post** /v1/Services/{ServiceSid}/AlphaSenders | 
[**DeleteAlphaSender**](ServicesAlphaSendersApi.md#DeleteAlphaSender) | **Delete** /v1/Services/{ServiceSid}/AlphaSenders/{Sid} | 
[**FetchAlphaSender**](ServicesAlphaSendersApi.md#FetchAlphaSender) | **Get** /v1/Services/{ServiceSid}/AlphaSenders/{Sid} | 
[**ListAlphaSender**](ServicesAlphaSendersApi.md#ListAlphaSender) | **Get** /v1/Services/{ServiceSid}/AlphaSenders | 



## CreateAlphaSender

> MessagingV1AlphaSender CreateAlphaSender(ctx, ServiceSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to create the resource under.

### Other Parameters

Other parameters are passed through a pointer to a CreateAlphaSenderParams struct


Name | Type | Description
------------- | ------------- | -------------
**AlphaSender** | **string** | The Alphanumeric Sender ID string. Can be up to 11 characters long. Valid characters are A-Z, a-z, 0-9, space, hyphen `-`, plus `+`, underscore `_` and ampersand `&`. This value cannot contain only numbers.

### Return type

[**MessagingV1AlphaSender**](MessagingV1AlphaSender.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAlphaSender

> DeleteAlphaSender(ctx, ServiceSidSid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to delete the resource from.
**Sid** | **string** | The SID of the AlphaSender resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteAlphaSenderParams struct


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


## FetchAlphaSender

> MessagingV1AlphaSender FetchAlphaSender(ctx, ServiceSidSid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to fetch the resource from.
**Sid** | **string** | The SID of the AlphaSender resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchAlphaSenderParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**MessagingV1AlphaSender**](MessagingV1AlphaSender.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAlphaSender

> []MessagingV1AlphaSender ListAlphaSender(ctx, ServiceSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to read the resources from.

### Other Parameters

Other parameters are passed through a pointer to a ListAlphaSenderParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]MessagingV1AlphaSender**](MessagingV1AlphaSender.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

