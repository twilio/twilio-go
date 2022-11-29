# TrunksOriginationUrlsApi

All URIs are relative to *https://trunking.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOriginationUrl**](TrunksOriginationUrlsApi.md#CreateOriginationUrl) | **Post** /v1/Trunks/{TrunkSid}/OriginationUrls | 
[**DeleteOriginationUrl**](TrunksOriginationUrlsApi.md#DeleteOriginationUrl) | **Delete** /v1/Trunks/{TrunkSid}/OriginationUrls/{Sid} | 
[**FetchOriginationUrl**](TrunksOriginationUrlsApi.md#FetchOriginationUrl) | **Get** /v1/Trunks/{TrunkSid}/OriginationUrls/{Sid} | 
[**ListOriginationUrl**](TrunksOriginationUrlsApi.md#ListOriginationUrl) | **Get** /v1/Trunks/{TrunkSid}/OriginationUrls | 
[**UpdateOriginationUrl**](TrunksOriginationUrlsApi.md#UpdateOriginationUrl) | **Post** /v1/Trunks/{TrunkSid}/OriginationUrls/{Sid} | 



## CreateOriginationUrl

> TrunkingV1OriginationUrl CreateOriginationUrl(ctx, TrunkSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrunkSid** | **string** | The SID of the Trunk to associate the resource with.

### Other Parameters

Other parameters are passed through a pointer to a CreateOriginationUrlParams struct


Name | Type | Description
------------- | ------------- | -------------
**Weight** | **int** | The value that determines the relative share of the load the URI should receive compared to other URIs with the same priority. Can be an integer from 1 to 65535, inclusive, and the default is 10. URLs with higher values receive more load than those with lower ones with the same priority.
**Priority** | **int** | The relative importance of the URI. Can be an integer from 0 to 65535, inclusive, and the default is 10. The lowest number represents the most important URI.
**Enabled** | **bool** | Whether the URL is enabled. The default is `true`.
**FriendlyName** | **string** | A descriptive string that you create to describe the resource. It can be up to 64 characters long.
**SipUrl** | **string** | The SIP address you want Twilio to route your Origination calls to. This must be a `sip:` schema.

### Return type

[**TrunkingV1OriginationUrl**](TrunkingV1OriginationUrl.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOriginationUrl

> DeleteOriginationUrl(ctx, TrunkSidSid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrunkSid** | **string** | The SID of the Trunk from which to delete the OriginationUrl.
**Sid** | **string** | The unique string that we created to identify the OriginationUrl resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteOriginationUrlParams struct


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


## FetchOriginationUrl

> TrunkingV1OriginationUrl FetchOriginationUrl(ctx, TrunkSidSid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrunkSid** | **string** | The SID of the Trunk from which to fetch the OriginationUrl.
**Sid** | **string** | The unique string that we created to identify the OriginationUrl resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchOriginationUrlParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**TrunkingV1OriginationUrl**](TrunkingV1OriginationUrl.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOriginationUrl

> []TrunkingV1OriginationUrl ListOriginationUrl(ctx, TrunkSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrunkSid** | **string** | The SID of the Trunk from which to read the OriginationUrl.

### Other Parameters

Other parameters are passed through a pointer to a ListOriginationUrlParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]TrunkingV1OriginationUrl**](TrunkingV1OriginationUrl.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOriginationUrl

> TrunkingV1OriginationUrl UpdateOriginationUrl(ctx, TrunkSidSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrunkSid** | **string** | The SID of the Trunk from which to update the OriginationUrl.
**Sid** | **string** | The unique string that we created to identify the OriginationUrl resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateOriginationUrlParams struct


Name | Type | Description
------------- | ------------- | -------------
**Weight** | **int** | The value that determines the relative share of the load the URI should receive compared to other URIs with the same priority. Can be an integer from 1 to 65535, inclusive, and the default is 10. URLs with higher values receive more load than those with lower ones with the same priority.
**Priority** | **int** | The relative importance of the URI. Can be an integer from 0 to 65535, inclusive, and the default is 10. The lowest number represents the most important URI.
**Enabled** | **bool** | Whether the URL is enabled. The default is `true`.
**FriendlyName** | **string** | A descriptive string that you create to describe the resource. It can be up to 64 characters long.
**SipUrl** | **string** | The SIP address you want Twilio to route your Origination calls to. This must be a `sip:` schema. `sips` is NOT supported.

### Return type

[**TrunkingV1OriginationUrl**](TrunkingV1OriginationUrl.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

