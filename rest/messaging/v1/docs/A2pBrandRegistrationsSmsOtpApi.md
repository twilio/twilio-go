# A2pBrandRegistrationsSmsOtpApi

All URIs are relative to *https://messaging.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBrandRegistrationOtp**](A2pBrandRegistrationsSmsOtpApi.md#CreateBrandRegistrationOtp) | **Post** /v1/a2p/BrandRegistrations/{BrandRegistrationSid}/SmsOtp | 



## CreateBrandRegistrationOtp

> MessagingV1BrandRegistrationOtp CreateBrandRegistrationOtp(ctx, BrandRegistrationSid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**BrandRegistrationSid** | **string** | Brand Registration Sid of Sole Proprietor Brand.

### Other Parameters

Other parameters are passed through a pointer to a CreateBrandRegistrationOtpParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**MessagingV1BrandRegistrationOtp**](MessagingV1BrandRegistrationOtp.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

