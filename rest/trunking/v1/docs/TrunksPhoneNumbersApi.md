# TrunksPhoneNumbersApi

All URIs are relative to *https://trunking.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePhoneNumber**](TrunksPhoneNumbersApi.md#CreatePhoneNumber) | **Post** /v1/Trunks/{TrunkSid}/PhoneNumbers | 
[**DeletePhoneNumber**](TrunksPhoneNumbersApi.md#DeletePhoneNumber) | **Delete** /v1/Trunks/{TrunkSid}/PhoneNumbers/{Sid} | 
[**FetchPhoneNumber**](TrunksPhoneNumbersApi.md#FetchPhoneNumber) | **Get** /v1/Trunks/{TrunkSid}/PhoneNumbers/{Sid} | 
[**ListPhoneNumber**](TrunksPhoneNumbersApi.md#ListPhoneNumber) | **Get** /v1/Trunks/{TrunkSid}/PhoneNumbers | 



## CreatePhoneNumber

> TrunkingV1PhoneNumber CreatePhoneNumber(ctx, TrunkSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrunkSid** | **string** | The SID of the Trunk to associate the phone number with.

### Other Parameters

Other parameters are passed through a pointer to a CreatePhoneNumberParams struct


Name | Type | Description
------------- | ------------- | -------------
**PhoneNumberSid** | **string** | The SID of the [Incoming Phone Number](https://www.twilio.com/docs/phone-numbers/api/incomingphonenumber-resource) that you want to associate with the trunk.

### Return type

[**TrunkingV1PhoneNumber**](TrunkingV1PhoneNumber.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePhoneNumber

> DeletePhoneNumber(ctx, TrunkSidSid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrunkSid** | **string** | The SID of the Trunk from which to delete the PhoneNumber resource.
**Sid** | **string** | The unique string that we created to identify the PhoneNumber resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeletePhoneNumberParams struct


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


## FetchPhoneNumber

> TrunkingV1PhoneNumber FetchPhoneNumber(ctx, TrunkSidSid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrunkSid** | **string** | The SID of the Trunk from which to fetch the PhoneNumber resource.
**Sid** | **string** | The unique string that we created to identify the PhoneNumber resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchPhoneNumberParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**TrunkingV1PhoneNumber**](TrunkingV1PhoneNumber.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPhoneNumber

> []TrunkingV1PhoneNumber ListPhoneNumber(ctx, TrunkSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrunkSid** | **string** | The SID of the Trunk from which to read the PhoneNumber resources.

### Other Parameters

Other parameters are passed through a pointer to a ListPhoneNumberParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]TrunkingV1PhoneNumber**](TrunkingV1PhoneNumber.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

