# ComplianceInquiriesTollfreeInitializeApi

All URIs are relative to *https://trusthub.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateComplianceTollfreeInquiry**](ComplianceInquiriesTollfreeInitializeApi.md#CreateComplianceTollfreeInquiry) | **Post** /v1/ComplianceInquiries/Tollfree/Initialize | 



## CreateComplianceTollfreeInquiry

> TrusthubV1ComplianceTollfreeInquiry CreateComplianceTollfreeInquiry(ctx, optional)



Create a new Compliance Tollfree Verification Inquiry for the authenticated account. This is necessary to start a new embedded session.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateComplianceTollfreeInquiryParams struct


Name | Type | Description
------------- | ------------- | -------------
**TollfreePhoneNumber** | **string** | The Tollfree phone number to be verified
**NotificationEmail** | **string** | The notification email to be triggered when verification status is changed

### Return type

[**TrusthubV1ComplianceTollfreeInquiry**](TrusthubV1ComplianceTollfreeInquiry.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

