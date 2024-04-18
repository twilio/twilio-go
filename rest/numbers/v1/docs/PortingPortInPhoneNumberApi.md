# PortingPortInPhoneNumberApi

All URIs are relative to *https://numbers.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePortingPortInPhoneNumber**](PortingPortInPhoneNumberApi.md#DeletePortingPortInPhoneNumber) | **Delete** /v1/Porting/PortIn/{PortInRequestSid}/PhoneNumber/{PhoneNumberSid} | 



## DeletePortingPortInPhoneNumber

> DeletePortingPortInPhoneNumber(ctx, PortInRequestSidPhoneNumberSid)



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

