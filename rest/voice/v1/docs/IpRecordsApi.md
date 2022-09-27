# IpRecordsApi

All URIs are relative to *https://voice.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIpRecord**](IpRecordsApi.md#CreateIpRecord) | **Post** /v1/IpRecords | 
[**DeleteIpRecord**](IpRecordsApi.md#DeleteIpRecord) | **Delete** /v1/IpRecords/{Sid} | 
[**FetchIpRecord**](IpRecordsApi.md#FetchIpRecord) | **Get** /v1/IpRecords/{Sid} | 
[**ListIpRecord**](IpRecordsApi.md#ListIpRecord) | **Get** /v1/IpRecords | 
[**UpdateIpRecord**](IpRecordsApi.md#UpdateIpRecord) | **Post** /v1/IpRecords/{Sid} | 



## CreateIpRecord

> VoiceV1IpRecord CreateIpRecord(ctx, optional)





### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateIpRecordParams struct


Name | Type | Description
------------- | ------------- | -------------
**IpAddress** | **string** | An IP address in dotted decimal notation, IPv4 only.
**FriendlyName** | **string** | A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long.
**CidrPrefixLength** | **int** | An integer representing the length of the [CIDR](https://tools.ietf.org/html/rfc4632) prefix to use with this IP address. By default the entire IP address is used, which for IPv4 is value 32.

### Return type

[**VoiceV1IpRecord**](VoiceV1IpRecord.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIpRecord

> DeleteIpRecord(ctx, Sid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the IP Record resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteIpRecordParams struct


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


## FetchIpRecord

> VoiceV1IpRecord FetchIpRecord(ctx, Sid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the IP Record resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchIpRecordParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**VoiceV1IpRecord**](VoiceV1IpRecord.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIpRecord

> []VoiceV1IpRecord ListIpRecord(ctx, optional)





### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListIpRecordParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]VoiceV1IpRecord**](VoiceV1IpRecord.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIpRecord

> VoiceV1IpRecord UpdateIpRecord(ctx, Sidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the IP Record resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateIpRecordParams struct


Name | Type | Description
------------- | ------------- | -------------
**FriendlyName** | **string** | A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long.

### Return type

[**VoiceV1IpRecord**](VoiceV1IpRecord.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

