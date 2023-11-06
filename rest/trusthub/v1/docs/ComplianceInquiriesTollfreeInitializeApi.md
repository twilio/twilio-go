# ComplianceInquiriesTollfreeInitializeApi

All URIs are relative to *https://trusthub.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateComplianceTollfreeInquiry**](ComplianceInquiriesTollfreeInitializeApi.md#CreateComplianceTollfreeInquiry) | **Post** /v1/ComplianceInquiries/Tollfree/Initialize | 
[**UpdateComplianceTollfreeInquiry**](ComplianceInquiriesTollfreeInitializeApi.md#UpdateComplianceTollfreeInquiry) | **Post** /v1/ComplianceInquiries/Tollfree/{TollfreeId}/Initialize | 



## CreateComplianceTollfreeInquiry

> TrusthubV1ComplianceTollfreeInquiry CreateComplianceTollfreeInquiry(ctx, optional)



Create a new Compliance Tollfree Verification Inquiry for the authenticated account. This is necessary to start a new embedded session.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateComplianceTollfreeInquiryParams struct


Name | Type | Description
------------- | ------------- | -------------
**Did** | **string** | The Tollfree phone number to be verified

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


## UpdateComplianceTollfreeInquiry

> TrusthubV1ComplianceTollfreeInquiry UpdateComplianceTollfreeInquiry(ctx, TollfreeIdoptional)



Resume a specific Compliance Tollfree Verification Inquiry that has expired, or re-open a rejected Compliance Tollfree Verification Inquiry for editing.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TollfreeId** | **string** | The unique TolfreeId matching the Compliance Tollfree Verification Inquiry that should be resumed or resubmitted. This value will have been returned by the initial Compliance Tollfree Verification Inquiry creation call.

### Other Parameters

Other parameters are passed through a pointer to a UpdateComplianceTollfreeInquiryParams struct


Name | Type | Description
------------- | ------------- | -------------
**Did** | **string** | The Tollfree phone number to be verified

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

