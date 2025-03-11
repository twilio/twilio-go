# PortingPortInPhoneNumberApi

All URIs are relative to *https://numbers.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePortingPortInPhoneNumber**](PortingPortInPhoneNumberApi.md#DeletePortingPortInPhoneNumber) | **Delete** /v1/Porting/PortIn/{PortInRequestSid}/PhoneNumber/{PhoneNumberSid} | Allows to cancel a port in request phone number by SID
[**FetchPortingPortInPhoneNumber**](PortingPortInPhoneNumberApi.md#FetchPortingPortInPhoneNumber) | **Get** /v1/Porting/PortIn/{PortInRequestSid}/PhoneNumber/{PhoneNumberSid} | Fetch a phone number by port in request SID and phone number SID



## DeletePortingPortInPhoneNumber

> DeletePortingPortInPhoneNumber(ctx, PortInRequestSidPhoneNumberSid)

Allows to cancel a port in request phone number by SID

Allows to cancel a port in request phone number by SID

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**PortInRequestSid** | **string** | The SID of the Port In request. This is a unique identifier of the port in request.
**PhoneNumberSid** | **string** | The SID of the Port In request phone number. This is a unique identifier of the phone number.

### Other Parameters

Other parameters are passed through a pointer to a DeletePortingPortInPhoneNumberParams struct


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


## FetchPortingPortInPhoneNumber

> NumbersV1PortingPortInPhoneNumber FetchPortingPortInPhoneNumber(ctx, PortInRequestSidPhoneNumberSid)

Fetch a phone number by port in request SID and phone number SID

Fetch a phone number by port in request SID and phone number SID

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**PortInRequestSid** | **string** | The SID of the Port In request. This is a unique identifier of the port in request.
**PhoneNumberSid** | **string** | The SID of the Phone number. This is a unique identifier of the phone number.

### Other Parameters

Other parameters are passed through a pointer to a FetchPortingPortInPhoneNumberParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**NumbersV1PortingPortInPhoneNumber**](NumbersV1PortingPortInPhoneNumber.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

