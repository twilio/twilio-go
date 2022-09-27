# SourceIpMappingsApi

All URIs are relative to *https://voice.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSourceIpMapping**](SourceIpMappingsApi.md#CreateSourceIpMapping) | **Post** /v1/SourceIpMappings | 
[**DeleteSourceIpMapping**](SourceIpMappingsApi.md#DeleteSourceIpMapping) | **Delete** /v1/SourceIpMappings/{Sid} | 
[**FetchSourceIpMapping**](SourceIpMappingsApi.md#FetchSourceIpMapping) | **Get** /v1/SourceIpMappings/{Sid} | 
[**ListSourceIpMapping**](SourceIpMappingsApi.md#ListSourceIpMapping) | **Get** /v1/SourceIpMappings | 
[**UpdateSourceIpMapping**](SourceIpMappingsApi.md#UpdateSourceIpMapping) | **Post** /v1/SourceIpMappings/{Sid} | 



## CreateSourceIpMapping

> VoiceV1SourceIpMapping CreateSourceIpMapping(ctx, optional)





### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateSourceIpMappingParams struct


Name | Type | Description
------------- | ------------- | -------------
**IpRecordSid** | **string** | The Twilio-provided string that uniquely identifies the IP Record resource to map from.
**SipDomainSid** | **string** | The SID of the SIP Domain that the IP Record should be mapped to.

### Return type

[**VoiceV1SourceIpMapping**](VoiceV1SourceIpMapping.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSourceIpMapping

> DeleteSourceIpMapping(ctx, Sid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the IP Record resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteSourceIpMappingParams struct


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


## FetchSourceIpMapping

> VoiceV1SourceIpMapping FetchSourceIpMapping(ctx, Sid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the IP Record resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchSourceIpMappingParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**VoiceV1SourceIpMapping**](VoiceV1SourceIpMapping.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSourceIpMapping

> []VoiceV1SourceIpMapping ListSourceIpMapping(ctx, optional)





### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListSourceIpMappingParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]VoiceV1SourceIpMapping**](VoiceV1SourceIpMapping.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSourceIpMapping

> VoiceV1SourceIpMapping UpdateSourceIpMapping(ctx, Sidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the IP Record resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateSourceIpMappingParams struct


Name | Type | Description
------------- | ------------- | -------------
**SipDomainSid** | **string** | The SID of the SIP Domain that the IP Record should be mapped to.

### Return type

[**VoiceV1SourceIpMapping**](VoiceV1SourceIpMapping.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

