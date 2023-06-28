# PortingPortabilityApi

All URIs are relative to *https://numbers.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePortingBulkPortability**](PortingPortabilityApi.md#CreatePortingBulkPortability) | **Post** /v1/Porting/Portability | 
[**FetchPortingBulkPortability**](PortingPortabilityApi.md#FetchPortingBulkPortability) | **Get** /v1/Porting/Portability/{Sid} | 



## CreatePortingBulkPortability

> NumbersV1PortingBulkPortability CreatePortingBulkPortability(ctx, optional)



Allows to check if a list of phone numbers can be ported to Twilio or not. This is done asynchronous for each phone number.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreatePortingBulkPortabilityParams struct


Name | Type | Description
------------- | ------------- | -------------
**PhoneNumbers** | **[]string** | The phone numbers which portability is to be checked. This should be a list of strings. Phone numbers are in E.164 format (e.g. +16175551212). .

### Return type

[**NumbersV1PortingBulkPortability**](NumbersV1PortingBulkPortability.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchPortingBulkPortability

> NumbersV1PortingBulkPortability FetchPortingBulkPortability(ctx, Sid)



Fetch a previous portability check. This should return the current status of the validation and the result for all the numbers provided, given that they have been validated (as this process is performed asynchronously).

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | A 34 character string that uniquely identifies the Portability check.

### Other Parameters

Other parameters are passed through a pointer to a FetchPortingBulkPortabilityParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**NumbersV1PortingBulkPortability**](NumbersV1PortingBulkPortability.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

