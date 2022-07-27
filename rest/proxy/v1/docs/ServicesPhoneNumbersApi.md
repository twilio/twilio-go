# ServicesPhoneNumbersApi

All URIs are relative to *https://proxy.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePhoneNumber**](ServicesPhoneNumbersApi.md#CreatePhoneNumber) | **Post** /v1/Services/{ServiceSid}/PhoneNumbers | 
[**DeletePhoneNumber**](ServicesPhoneNumbersApi.md#DeletePhoneNumber) | **Delete** /v1/Services/{ServiceSid}/PhoneNumbers/{Sid} | 
[**FetchPhoneNumber**](ServicesPhoneNumbersApi.md#FetchPhoneNumber) | **Get** /v1/Services/{ServiceSid}/PhoneNumbers/{Sid} | 
[**ListPhoneNumber**](ServicesPhoneNumbersApi.md#ListPhoneNumber) | **Get** /v1/Services/{ServiceSid}/PhoneNumbers | 
[**UpdatePhoneNumber**](ServicesPhoneNumbersApi.md#UpdatePhoneNumber) | **Post** /v1/Services/{ServiceSid}/PhoneNumbers/{Sid} | 



## CreatePhoneNumber

> ProxyV1PhoneNumber CreatePhoneNumber(ctx, ServiceSidoptional)



Add a Phone Number to a Service's Proxy Number Pool.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID parent [Service](https://www.twilio.com/docs/proxy/api/service) resource of the new PhoneNumber resource.

### Other Parameters

Other parameters are passed through a pointer to a CreatePhoneNumberParams struct


Name | Type | Description
------------- | ------------- | -------------
**Sid** | **string** | The SID of a Twilio [IncomingPhoneNumber](https://www.twilio.com/docs/phone-numbers/api/incomingphonenumber-resource) resource that represents the Twilio Number you would like to assign to your Proxy Service.
**PhoneNumber** | **string** | The phone number in [E.164](https://www.twilio.com/docs/glossary/what-e164) format.  E.164 phone numbers consist of a + followed by the country code and subscriber number without punctuation characters. For example, +14155551234.
**IsReserved** | **bool** | Whether the new phone number should be reserved and not be assigned to a participant using proxy pool logic. See [Reserved Phone Numbers](https://www.twilio.com/docs/proxy/reserved-phone-numbers) for more information.

### Return type

[**ProxyV1PhoneNumber**](ProxyV1PhoneNumber.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePhoneNumber

> DeletePhoneNumber(ctx, ServiceSidSid)



Delete a specific Phone Number from a Service.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) of the PhoneNumber resource to delete.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the PhoneNumber resource to delete.

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

> ProxyV1PhoneNumber FetchPhoneNumber(ctx, ServiceSidSid)



Fetch a specific Phone Number.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) of the PhoneNumber resource to fetch.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the PhoneNumber resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchPhoneNumberParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ProxyV1PhoneNumber**](ProxyV1PhoneNumber.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPhoneNumber

> []ProxyV1PhoneNumber ListPhoneNumber(ctx, ServiceSidoptional)



Retrieve a list of all Phone Numbers in the Proxy Number Pool for a Service. A maximum of 100 records will be returned per page.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) of the PhoneNumber resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListPhoneNumberParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]ProxyV1PhoneNumber**](ProxyV1PhoneNumber.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePhoneNumber

> ProxyV1PhoneNumber UpdatePhoneNumber(ctx, ServiceSidSidoptional)



Update a specific Proxy Number.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) of the PhoneNumber resource to update.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the PhoneNumber resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdatePhoneNumberParams struct


Name | Type | Description
------------- | ------------- | -------------
**IsReserved** | **bool** | Whether the phone number should be reserved and not be assigned to a participant using proxy pool logic. See [Reserved Phone Numbers](https://www.twilio.com/docs/proxy/reserved-phone-numbers) for more information.

### Return type

[**ProxyV1PhoneNumber**](ProxyV1PhoneNumber.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

