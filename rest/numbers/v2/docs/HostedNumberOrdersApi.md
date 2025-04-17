# HostedNumberOrdersApi

All URIs are relative to *https://numbers.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateHostedNumberOrder**](HostedNumberOrdersApi.md#CreateHostedNumberOrder) | **Post** /v2/HostedNumber/Orders | Host a phone number&#39;s capability on Twilio&#39;s platform.
[**DeleteHostedNumberOrder**](HostedNumberOrdersApi.md#DeleteHostedNumberOrder) | **Delete** /v2/HostedNumber/Orders/{Sid} | Cancel the HostedNumberOrder (only available when the status is in &#x60;received&#x60;).
[**FetchHostedNumberOrder**](HostedNumberOrdersApi.md#FetchHostedNumberOrder) | **Get** /v2/HostedNumber/Orders/{Sid} | Fetch a specific HostedNumberOrder.
[**ListHostedNumberOrder**](HostedNumberOrdersApi.md#ListHostedNumberOrder) | **Get** /v2/HostedNumber/Orders | Retrieve a list of HostedNumberOrders belonging to the account initiating the request.
[**UpdateHostedNumberOrder**](HostedNumberOrdersApi.md#UpdateHostedNumberOrder) | **Post** /v2/HostedNumber/Orders/{Sid} | Updates a specific HostedNumberOrder.



## CreateHostedNumberOrder

> NumbersV2HostedNumberOrder CreateHostedNumberOrder(ctx, optional)

Host a phone number's capability on Twilio's platform.

Host a phone number's capability on Twilio's platform.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateHostedNumberOrderParams struct


Name | Type | Description
------------- | ------------- | -------------
**PhoneNumber** | **string** | The number to host in [+E.164](https://en.wikipedia.org/wiki/E.164) format
**ContactPhoneNumber** | **string** | The contact phone number of the person authorized to sign the Authorization Document.
**AddressSid** | **string** | Optional. A 34 character string that uniquely identifies the Address resource that represents the address of the owner of this phone number.
**Email** | **string** | Optional. Email of the owner of this phone number that is being hosted.
**AccountSid** | **string** | This defaults to the AccountSid of the authorization the user is using. This can be provided to specify a subaccount to add the HostedNumberOrder to.
**FriendlyName** | **string** | A 128 character string that is a human readable text that describes this resource.
**CcEmails** | **[]string** | Optional. A list of emails that the LOA document for this HostedNumberOrder will be carbon copied to.
**SmsUrl** | **string** | The URL that Twilio should request when somebody sends an SMS to the phone number. This will be copied onto the IncomingPhoneNumber resource.
**SmsMethod** | **string** | The HTTP method that should be used to request the SmsUrl. Must be either `GET` or `POST`.  This will be copied onto the IncomingPhoneNumber resource.
**SmsFallbackUrl** | **string** | A URL that Twilio will request if an error occurs requesting or executing the TwiML defined by SmsUrl. This will be copied onto the IncomingPhoneNumber resource.
**SmsCapability** | **bool** | Used to specify that the SMS capability will be hosted on Twilio's platform.
**SmsFallbackMethod** | **string** | The HTTP method that should be used to request the SmsFallbackUrl. Must be either `GET` or `POST`. This will be copied onto the IncomingPhoneNumber resource.
**StatusCallbackUrl** | **string** | Optional. The Status Callback URL attached to the IncomingPhoneNumber resource.
**StatusCallbackMethod** | **string** | Optional. The Status Callback Method attached to the IncomingPhoneNumber resource.
**SmsApplicationSid** | **string** | Optional. The 34 character sid of the application Twilio should use to handle SMS messages sent to this number. If a `SmsApplicationSid` is present, Twilio will ignore all of the SMS urls above and use those set on the application.
**ContactTitle** | **string** | The title of the person authorized to sign the Authorization Document for this phone number.

### Return type

[**NumbersV2HostedNumberOrder**](NumbersV2HostedNumberOrder.md)

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

Cancel the HostedNumberOrder (only available when the status is in `received`).

Cancel the HostedNumberOrder (only available when the status is in `received`).

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | A 34 character string that uniquely identifies this HostedNumberOrder.

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

> NumbersV2HostedNumberOrder FetchHostedNumberOrder(ctx, Sid)

Fetch a specific HostedNumberOrder.

Fetch a specific HostedNumberOrder.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | A 34 character string that uniquely identifies this HostedNumberOrder.

### Other Parameters

Other parameters are passed through a pointer to a FetchHostedNumberOrderParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**NumbersV2HostedNumberOrder**](NumbersV2HostedNumberOrder.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHostedNumberOrder

> []NumbersV2HostedNumberOrder ListHostedNumberOrder(ctx, optional)

Retrieve a list of HostedNumberOrders belonging to the account initiating the request.

Retrieve a list of HostedNumberOrders belonging to the account initiating the request.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListHostedNumberOrderParams struct


Name | Type | Description
------------- | ------------- | -------------
**Status** | [**string**](stringstring.md) | The Status of this HostedNumberOrder. One of `received`, `pending-verification`, `verified`, `pending-loa`, `carrier-processing`, `testing`, `completed`, `failed`, or `action-required`.
**SmsCapability** | **bool** | Whether the SMS capability will be hosted on our platform. Can be `true` of `false`.
**PhoneNumber** | **string** | An E164 formatted phone number hosted by this HostedNumberOrder.
**IncomingPhoneNumberSid** | **string** | A 34 character string that uniquely identifies the IncomingPhoneNumber resource created by this HostedNumberOrder.
**FriendlyName** | **string** | A human readable description of this resource, up to 128 characters.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]NumbersV2HostedNumberOrder**](NumbersV2HostedNumberOrder.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateHostedNumberOrder

> NumbersV2HostedNumberOrder UpdateHostedNumberOrder(ctx, Sidoptional)

Updates a specific HostedNumberOrder.

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
**Status** | [**string**](string.md) | 
**VerificationCallDelay** | **int** | The number of seconds to wait before initiating the ownership verification call. Can be a value between 0 and 60, inclusive.
**VerificationCallExtension** | **string** | The numerical extension to dial when making the ownership verification call.

### Return type

[**NumbersV2HostedNumberOrder**](NumbersV2HostedNumberOrder.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

