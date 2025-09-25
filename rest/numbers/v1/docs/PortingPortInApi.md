# PortingPortInApi

All URIs are relative to *https://numbers.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePortingPortIn**](PortingPortInApi.md#CreatePortingPortIn) | **Post** /v1/Porting/PortIn | Allows to create a new port in request
[**DeletePortingPortIn**](PortingPortInApi.md#DeletePortingPortIn) | **Delete** /v1/Porting/PortIn/{PortInRequestSid} | Allows to cancel a port in request by SID
[**FetchPortingPortIn**](PortingPortInApi.md#FetchPortingPortIn) | **Get** /v1/Porting/PortIn/{PortInRequestSid} | Fetch a port in request by SID



## CreatePortingPortIn

> NumbersV1PortingPortIn CreatePortingPortIn(ctx, optional)

Allows to create a new port in request

Allows to create a new port in request

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreatePortingPortInParams struct


Name | Type | Description
------------- | ------------- | -------------
**NumbersV1PortingPortInCreate** | [**NumbersV1PortingPortInCreate**](NumbersV1PortingPortInCreate.md) | 

### Return type

[**NumbersV1PortingPortIn**](NumbersV1PortingPortIn.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePortingPortIn

> DeletePortingPortIn(ctx, PortInRequestSid)

Allows to cancel a port in request by SID

Allows to cancel a port in request by SID

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**PortInRequestSid** | **string** | The SID of the Port In request. This is a unique identifier of the port in request.

### Other Parameters

Other parameters are passed through a pointer to a DeletePortingPortInParams struct


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


## FetchPortingPortIn

> NumbersV1PortingPortIn FetchPortingPortIn(ctx, PortInRequestSid)

Fetch a port in request by SID

Fetch a port in request by SID

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**PortInRequestSid** | **string** | The SID of the Port In request. This is a unique identifier of the port in request.

### Other Parameters

Other parameters are passed through a pointer to a FetchPortingPortInParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**NumbersV1PortingPortIn**](NumbersV1PortingPortIn.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

