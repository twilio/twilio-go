# PortingPortabilityPhoneNumberApi

All URIs are relative to *https://numbers.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchPortingPortability**](PortingPortabilityPhoneNumberApi.md#FetchPortingPortability) | **Get** /v1/Porting/Portability/PhoneNumber/{PhoneNumber} | Check if a single phone number can be ported to Twilio



## FetchPortingPortability

> NumbersV1PortingPortability FetchPortingPortability(ctx, PhoneNumberoptional)

Check if a single phone number can be ported to Twilio

Check if a single phone number can be ported to Twilio

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**PhoneNumber** | **string** | Phone number to check portability in e164 format.

### Other Parameters

Other parameters are passed through a pointer to a FetchPortingPortabilityParams struct


Name | Type | Description
------------- | ------------- | -------------
**TargetAccountSid** | **string** | Account Sid to which the number will be ported. This can be used to determine if a sub account already has the number in its inventory or a different sub account. If this is not provided, the authenticated account will be assumed to be the target account.
**AddressSid** | **string** | Address Sid of customer to which the number will be ported.

### Return type

[**NumbersV1PortingPortability**](NumbersV1PortingPortability.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

