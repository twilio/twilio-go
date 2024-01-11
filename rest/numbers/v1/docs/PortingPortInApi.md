# PortingPortInApi

All URIs are relative to *https://numbers.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchPortingPortInFetch**](PortingPortInApi.md#FetchPortingPortInFetch) | **Get** /v1/Porting/PortIn/{PortInRequestSid} | 



## FetchPortingPortInFetch

> NumbersV1PortingPortInFetch FetchPortingPortInFetch(ctx, PortInRequestSid)



Fetch a port in request by SID

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**PortInRequestSid** | **string** | The SID of the Port In request. This is a unique identifier of the port in request.

### Other Parameters

Other parameters are passed through a pointer to a FetchPortingPortInFetchParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**NumbersV1PortingPortInFetch**](NumbersV1PortingPortInFetch.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

