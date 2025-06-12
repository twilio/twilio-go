# ServicesDestinationAlphaSendersApi

All URIs are relative to *https://messaging.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDestinationAlphaSender**](ServicesDestinationAlphaSendersApi.md#CreateDestinationAlphaSender) | **Post** /v1/Services/{ServiceSid}/DestinationAlphaSenders | 
[**DeleteDestinationAlphaSender**](ServicesDestinationAlphaSendersApi.md#DeleteDestinationAlphaSender) | **Delete** /v1/Services/{ServiceSid}/DestinationAlphaSenders/{Sid} | 
[**FetchDestinationAlphaSender**](ServicesDestinationAlphaSendersApi.md#FetchDestinationAlphaSender) | **Get** /v1/Services/{ServiceSid}/DestinationAlphaSenders/{Sid} | 
[**ListDestinationAlphaSender**](ServicesDestinationAlphaSendersApi.md#ListDestinationAlphaSender) | **Get** /v1/Services/{ServiceSid}/DestinationAlphaSenders | 



## CreateDestinationAlphaSender

> MessagingV1DestinationAlphaSender CreateDestinationAlphaSender(ctx, ServiceSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to create the resource under.

### Other Parameters

Other parameters are passed through a pointer to a CreateDestinationAlphaSenderParams struct


Name | Type | Description
------------- | ------------- | -------------
**AlphaSender** | **string** | The Alphanumeric Sender ID string. Can be up to 11 characters long. Valid characters are A-Z, a-z, 0-9, space, hyphen `-`, plus `+`, underscore `_` and ampersand `&`. This value cannot contain only numbers.
**IsoCountryCode** | **string** | The Optional Two Character ISO Country Code the Alphanumeric Sender ID will be used for. If the IsoCountryCode is not provided, a default Alpha Sender will be created that can be used across all countries.

### Return type

[**MessagingV1DestinationAlphaSender**](MessagingV1DestinationAlphaSender.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDestinationAlphaSender

> DeleteDestinationAlphaSender(ctx, ServiceSidSid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to delete the resource from.
**Sid** | **string** | The SID of the AlphaSender resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteDestinationAlphaSenderParams struct


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


## FetchDestinationAlphaSender

> MessagingV1DestinationAlphaSender FetchDestinationAlphaSender(ctx, ServiceSidSid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to fetch the resource from.
**Sid** | **string** | The SID of the AlphaSender resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchDestinationAlphaSenderParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**MessagingV1DestinationAlphaSender**](MessagingV1DestinationAlphaSender.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDestinationAlphaSender

> []MessagingV1DestinationAlphaSender ListDestinationAlphaSender(ctx, ServiceSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to read the resources from.

### Other Parameters

Other parameters are passed through a pointer to a ListDestinationAlphaSenderParams struct


Name | Type | Description
------------- | ------------- | -------------
**IsoCountryCode** | **string** | Optional filter to return only alphanumeric sender IDs associated with the specified two-character ISO country code.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]MessagingV1DestinationAlphaSender**](MessagingV1DestinationAlphaSender.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

