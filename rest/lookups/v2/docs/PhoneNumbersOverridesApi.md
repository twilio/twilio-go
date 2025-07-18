# PhoneNumbersOverridesApi

All URIs are relative to *https://lookups.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLookupPhoneNumberOverrides**](PhoneNumbersOverridesApi.md#CreateLookupPhoneNumberOverrides) | **Post** /v2/PhoneNumbers/{PhoneNumber}/Overrides/{Field} | Create Override for a Phone Number for a specific field
[**DeleteLookupPhoneNumberOverrides**](PhoneNumbersOverridesApi.md#DeleteLookupPhoneNumberOverrides) | **Delete** /v2/PhoneNumbers/{PhoneNumber}/Overrides/{Field} | Delete an Override for a Phone Number for a specific field
[**FetchLookupPhoneNumberOverrides**](PhoneNumbersOverridesApi.md#FetchLookupPhoneNumberOverrides) | **Get** /v2/PhoneNumbers/{PhoneNumber}/Overrides/{Field} | Get Overrides for a Phone Number for a specific field.
[**UpdateLookupPhoneNumberOverrides**](PhoneNumbersOverridesApi.md#UpdateLookupPhoneNumberOverrides) | **Put** /v2/PhoneNumbers/{PhoneNumber}/Overrides/{Field} | Update Override for a Phone Number for a specific field



## CreateLookupPhoneNumberOverrides

> OverridesResponse CreateLookupPhoneNumberOverrides(ctx, FieldPhoneNumberoptional)

Create Override for a Phone Number for a specific field

Create an Override for a specific package and phone number.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Field** | **string** | 
**PhoneNumber** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a CreateLookupPhoneNumberOverridesParams struct


Name | Type | Description
------------- | ------------- | -------------
**OverridesRequest** | [**OverridesRequest**](OverridesRequest.md) | 

### Return type

[**OverridesResponse**](OverridesResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLookupPhoneNumberOverrides

> DeleteLookupPhoneNumberOverrides(ctx, FieldPhoneNumber)

Delete an Override for a Phone Number for a specific field

Delete an Override for a specific package and phone number.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Field** | **string** | 
**PhoneNumber** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a DeleteLookupPhoneNumberOverridesParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchLookupPhoneNumberOverrides

> OverridesResponse FetchLookupPhoneNumberOverrides(ctx, FieldPhoneNumber)

Get Overrides for a Phone Number for a specific field.

Retrieve an Override for a specific package and phone number.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Field** | **string** | 
**PhoneNumber** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a FetchLookupPhoneNumberOverridesParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**OverridesResponse**](OverridesResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLookupPhoneNumberOverrides

> OverridesResponse UpdateLookupPhoneNumberOverrides(ctx, FieldPhoneNumberoptional)

Update Override for a Phone Number for a specific field

Update an Override for a specific package and phone number.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Field** | **string** | 
**PhoneNumber** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateLookupPhoneNumberOverridesParams struct


Name | Type | Description
------------- | ------------- | -------------
**OverridesRequest** | [**OverridesRequest**](OverridesRequest.md) | 

### Return type

[**OverridesResponse**](OverridesResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

