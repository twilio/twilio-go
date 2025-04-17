# SafeListNumbersApi

All URIs are relative to *https://accounts.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSafelist**](SafeListNumbersApi.md#CreateSafelist) | **Post** /v1/SafeList/Numbers | Add a new phone number or phone number 1k prefix to SafeList.
[**DeleteSafelist**](SafeListNumbersApi.md#DeleteSafelist) | **Delete** /v1/SafeList/Numbers | Remove a phone number or phone number 1k prefix from SafeList.
[**FetchSafelist**](SafeListNumbersApi.md#FetchSafelist) | **Get** /v1/SafeList/Numbers | Check if a phone number or phone number 1k prefix exists in SafeList.



## CreateSafelist

> AccountsV1Safelist CreateSafelist(ctx, optional)

Add a new phone number or phone number 1k prefix to SafeList.

Add a new phone number or phone number 1k prefix to SafeList.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateSafelistParams struct


Name | Type | Description
------------- | ------------- | -------------
**PhoneNumber** | **string** | The phone number or phone number 1k prefix to be added in SafeList. Phone numbers must be in [E.164 format](https://www.twilio.com/docs/glossary/what-e164).

### Return type

[**AccountsV1Safelist**](AccountsV1Safelist.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSafelist

> DeleteSafelist(ctx, optional)

Remove a phone number or phone number 1k prefix from SafeList.

Remove a phone number or phone number 1k prefix from SafeList.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a DeleteSafelistParams struct


Name | Type | Description
------------- | ------------- | -------------
**PhoneNumber** | **string** | The phone number or phone number 1k prefix to be removed from SafeList. Phone numbers must be in [E.164 format](https://www.twilio.com/docs/glossary/what-e164).

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

> AccountsV1Safelist FetchSafelist(ctx, optional)

Check if a phone number or phone number 1k prefix exists in SafeList.

Check if a phone number or phone number 1k prefix exists in SafeList.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a FetchSafelistParams struct


Name | Type | Description
------------- | ------------- | -------------
**PhoneNumber** | **string** | The phone number or phone number 1k prefix to be fetched from SafeList. Phone numbers must be in [E.164 format](https://www.twilio.com/docs/glossary/what-e164).

### Return type

[**AccountsV1Safelist**](AccountsV1Safelist.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

