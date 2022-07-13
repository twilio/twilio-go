# ServicesShortCodesApi

All URIs are relative to *https://messaging.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateShortCode**](ServicesShortCodesApi.md#CreateShortCode) | **Post** /v1/Services/{ServiceSid}/ShortCodes | 
[**DeleteShortCode**](ServicesShortCodesApi.md#DeleteShortCode) | **Delete** /v1/Services/{ServiceSid}/ShortCodes/{Sid} | 
[**FetchShortCode**](ServicesShortCodesApi.md#FetchShortCode) | **Get** /v1/Services/{ServiceSid}/ShortCodes/{Sid} | 
[**ListShortCode**](ServicesShortCodesApi.md#ListShortCode) | **Get** /v1/Services/{ServiceSid}/ShortCodes | 



## CreateShortCode

> MessagingV1ShortCode CreateShortCode(ctx, ServiceSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to create the resource under.

### Other Parameters

Other parameters are passed through a pointer to a CreateShortCodeParams struct


Name | Type | Description
------------- | ------------- | -------------
**ShortCodeSid** | **string** | The SID of the ShortCode resource being added to the Service.

### Return type

[**MessagingV1ShortCode**](MessagingV1ShortCode.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteShortCode

> DeleteShortCode(ctx, ServiceSidSid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to delete the resource from.
**Sid** | **string** | The SID of the ShortCode resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteShortCodeParams struct


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


## FetchShortCode

> MessagingV1ShortCode FetchShortCode(ctx, ServiceSidSid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to fetch the resource from.
**Sid** | **string** | The SID of the ShortCode resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchShortCodeParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**MessagingV1ShortCode**](MessagingV1ShortCode.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListShortCode

> []MessagingV1ShortCode ListShortCode(ctx, ServiceSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to read the resources from.

### Other Parameters

Other parameters are passed through a pointer to a ListShortCodeParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]MessagingV1ShortCode**](MessagingV1ShortCode.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

