# ServicesShortCodesApi

All URIs are relative to *https://proxy.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateShortCode**](ServicesShortCodesApi.md#CreateShortCode) | **Post** /v1/Services/{ServiceSid}/ShortCodes | 
[**DeleteShortCode**](ServicesShortCodesApi.md#DeleteShortCode) | **Delete** /v1/Services/{ServiceSid}/ShortCodes/{Sid} | 
[**FetchShortCode**](ServicesShortCodesApi.md#FetchShortCode) | **Get** /v1/Services/{ServiceSid}/ShortCodes/{Sid} | 
[**ListShortCode**](ServicesShortCodesApi.md#ListShortCode) | **Get** /v1/Services/{ServiceSid}/ShortCodes | 
[**UpdateShortCode**](ServicesShortCodesApi.md#UpdateShortCode) | **Post** /v1/Services/{ServiceSid}/ShortCodes/{Sid} | 



## CreateShortCode

> ProxyV1ShortCode CreateShortCode(ctx, ServiceSidoptional)



Add a Short Code to the Proxy Number Pool for the Service.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) resource.

### Other Parameters

Other parameters are passed through a pointer to a CreateShortCodeParams struct


Name | Type | Description
------------- | ------------- | -------------
**Sid** | **string** | The SID of a Twilio [ShortCode](https://www.twilio.com/docs/sms/api/short-code) resource that represents the short code you would like to assign to your Proxy Service.

### Return type

[**ProxyV1ShortCode**](ProxyV1ShortCode.md)

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



Delete a specific Short Code from a Service.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) resource to delete the ShortCode resource from.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the ShortCode resource to delete.

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

> ProxyV1ShortCode FetchShortCode(ctx, ServiceSidSid)



Fetch a specific Short Code.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) to fetch the resource from.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the ShortCode resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchShortCodeParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ProxyV1ShortCode**](ProxyV1ShortCode.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListShortCode

> []ProxyV1ShortCode ListShortCode(ctx, ServiceSidoptional)



Retrieve a list of all Short Codes in the Proxy Number Pool for the Service. A maximum of 100 records will be returned per page.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) to read the resources from.

### Other Parameters

Other parameters are passed through a pointer to a ListShortCodeParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]ProxyV1ShortCode**](ProxyV1ShortCode.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateShortCode

> ProxyV1ShortCode UpdateShortCode(ctx, ServiceSidSidoptional)



Update a specific Short Code.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) of the resource to update.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the ShortCode resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateShortCodeParams struct


Name | Type | Description
------------- | ------------- | -------------
**IsReserved** | **bool** | Whether the short code should be reserved and not be assigned to a participant using proxy pool logic. See [Reserved Phone Numbers](https://www.twilio.com/docs/proxy/reserved-phone-numbers) for more information.

### Return type

[**ProxyV1ShortCode**](ProxyV1ShortCode.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

