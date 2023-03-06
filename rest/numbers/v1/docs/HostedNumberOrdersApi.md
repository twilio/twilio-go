# HostedNumberOrdersApi

All URIs are relative to *https://numbers.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateHostedNumberOrder**](HostedNumberOrdersApi.md#CreateHostedNumberOrder) | **Post** /v1/HostedNumber/Orders | 
[**DeleteHostedNumberOrder**](HostedNumberOrdersApi.md#DeleteHostedNumberOrder) | **Delete** /v1/HostedNumber/Orders/{Sid} | 
[**FetchHostedNumberOrder**](HostedNumberOrdersApi.md#FetchHostedNumberOrder) | **Get** /v1/HostedNumber/Orders/{Sid} | 
[**ListHostedNumberOrder**](HostedNumberOrdersApi.md#ListHostedNumberOrder) | **Get** /v1/HostedNumber/Orders | 
[**UpdateHostedNumberOrder**](HostedNumberOrdersApi.md#UpdateHostedNumberOrder) | **Post** /v1/HostedNumber/Orders/{Sid} | 



## CreateHostedNumberOrder

> NumbersV1HostedNumberOrder CreateHostedNumberOrder(ctx, optional)



Host a phone number's capability on our platform.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateHostedNumberOrderParams struct


Name | Type | Description
------------- | ------------- | -------------
**PhoneNumber** | **string** | The [E.164](https://www.twilio.com/docs/glossary/what-e164) formatted phone number hosted by the new hosted number order.  E.164 phone numbers consist of a + followed by the country code and subscriber number without punctuation characters. For example, +14155551234.
**SmsCapability** | **bool** | Whether the SMS capability will be hosted on our platform. Can be `true` of `false`.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
**FriendlyName** | **string** | A descriptive string that you create to describe the resource. It can be up to 128 characters long.
**SmsApplicationSid** | **string** | The SID of the Application we should use to handle SMS messages sent to this number. If a `sms_application_sid` is present, we will ignore all of the SMS URLs and use those set on the application.
**SmsFallbackMethod** | **string** | The HTTP method we should use to call the `sms_fallback_url`. Can be: `POST` or `GET` and the default is `POST`.
**SmsFallbackUrl** | **string** | The URL we should call if there is a problem calling `sms_url` when the IncomingPhoneNumber receives an SMS message.
**SmsMethod** | **string** | The HTTP method we should use to call the `sms_url`. Can be: `POST` or `GET` and the default is `POST`.
**SmsUrl** | **string** | The URL we should call using the `sms_method` when the IncomingPhoneNumber receives an SMS message.
**StatusCallbackMethod** | **string** | The HTTP method we should use when calling `status_callback_url`.
**StatusCallbackUrl** | **string** | The URL we should call using the `status_callback_method` to send status information to your application.
**VerificationCallDelay** | **int** | The number of seconds to wait before initiating the ownership verification call. Can be a value between 0 and 60, inclusive.
**VerificationCallExtension** | **string** | The numerical extension to dial when making the ownership verification call.
**VerificationDocumentSid** | **string** | The SID of the identity document resource that represents the document used to verify ownership of the number to be hosted.
**VerificationType** | **string** | 

### Return type

[**NumbersV1HostedNumberOrder**](NumbersV1HostedNumberOrder.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteHostedNumberOrder

> DeleteHostedNumberOrder(ctx, Sid)



Cancel the HostedNumberOrder (only available when the status is in `twilio-processing`, `received`, `pending-verification`, `verified`, and `pending-loa`).

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the HostedNumberOrder resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteHostedNumberOrderParams struct


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


## FetchHostedNumberOrder

> NumbersV1HostedNumberOrder FetchHostedNumberOrder(ctx, Sid)



Fetch a specific HostedNumberOrder.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the HostedNumberOrder resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchHostedNumberOrderParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**NumbersV1HostedNumberOrder**](NumbersV1HostedNumberOrder.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHostedNumberOrder

> []NumbersV1HostedNumberOrder ListHostedNumberOrder(ctx, optional)



Retrieve a list of HostedNumberOrders belonging to the account initiating the request.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListHostedNumberOrderParams struct


Name | Type | Description
------------- | ------------- | -------------
**FriendlyName** | **string** | The string that identifies the HostedNumberOrder resources to read.
**IncomingPhoneNumberSid** | **string** | The SID of the IncomingPhoneNumber resource created by this HostedNumberOrder.
**PhoneNumber** | **string** | The [E.164](https://www.twilio.com/docs/glossary/what-e164) formatted phone numbers of the HostedNumberOrder resources to read. 
**Status** | **string** | The status of the resources to read. Can be: `twilio-processing`, `received`, `pending-verification`, `verified`, `pending-loa`, `carrier-processing`, `testing`, `completed`, `failed`, or `action-required`.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]NumbersV1HostedNumberOrder**](NumbersV1HostedNumberOrder.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateHostedNumberOrder

> NumbersV1HostedNumberOrder UpdateHostedNumberOrder(ctx, Sidoptional)



Updates a specific HostedNumberOrder.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the HostedNumberOrder resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateHostedNumberOrderParams struct


Name | Type | Description
------------- | ------------- | -------------
**FriendlyName** | **string** | A descriptive string that you create to describe the resource. It can be up to 128 characters long.
**Status** | **string** | 
**VerificationCallDelay** | **int** | The number of seconds to wait before initiating the ownership verification call. Can be a value between 0 and 60, inclusive.
**VerificationCallExtension** | **string** | The numerical extension to dial when making the ownership verification call.
**VerificationDocumentSid** | **string** | The SID of the Identity Document that represents the document for verifying ownership of the number to be hosted. Required when `verification_type` is `phone-bill`.
**VerificationType** | **string** | 

### Return type

[**NumbersV1HostedNumberOrder**](NumbersV1HostedNumberOrder.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

