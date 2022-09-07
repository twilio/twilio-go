# SafeListNumbersApi

All URIs are relative to *https://verify.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSafelist**](SafeListNumbersApi.md#CreateSafelist) | **Post** /v2/SafeList/Numbers | 
[**DeleteSafelist**](SafeListNumbersApi.md#DeleteSafelist) | **Delete** /v2/SafeList/Numbers/{PhoneNumber} | 
[**FetchSafelist**](SafeListNumbersApi.md#FetchSafelist) | **Get** /v2/SafeList/Numbers/{PhoneNumber} | 



## CreateSafelist

> VerifyV2Safelist CreateSafelist(ctx, optional)



Add a new phone number to SafeList.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateSafelistParams struct


Name | Type | Description
------------- | ------------- | -------------
**PhoneNumber** | **string** | The phone number to be added in SafeList. Phone numbers must be in [E.164 format](https://www.twilio.com/docs/glossary/what-e164).

### Return type

[**VerifyV2Safelist**](VerifyV2Safelist.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSafelist

> DeleteSafelist(ctx, PhoneNumber)



Remove a phone number from SafeList.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**PhoneNumber** | **string** | The phone number to be removed from SafeList. Phone numbers must be in [E.164 format](https://www.twilio.com/docs/glossary/what-e164).

### Other Parameters

Other parameters are passed through a pointer to a DeleteSafelistParams struct


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


## FetchSafelist

> VerifyV2Safelist FetchSafelist(ctx, PhoneNumber)



Check if a phone number exists in SafeList.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**PhoneNumber** | **string** | The phone number to be fetched from SafeList. Phone numbers must be in [E.164 format](https://www.twilio.com/docs/glossary/what-e164).

### Other Parameters

Other parameters are passed through a pointer to a FetchSafelistParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**VerifyV2Safelist**](VerifyV2Safelist.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

