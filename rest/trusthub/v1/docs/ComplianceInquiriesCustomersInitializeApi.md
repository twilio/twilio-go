# ComplianceInquiriesCustomersInitializeApi

All URIs are relative to *https://trusthub.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateComplianceInquiry**](ComplianceInquiriesCustomersInitializeApi.md#CreateComplianceInquiry) | **Post** /v1/ComplianceInquiries/Customers/Initialize | Create a new Compliance Inquiry for the authenticated account. This is necessary to start a new embedded session.
[**UpdateComplianceInquiry**](ComplianceInquiriesCustomersInitializeApi.md#UpdateComplianceInquiry) | **Post** /v1/ComplianceInquiries/Customers/{CustomerId}/Initialize | Resume a specific Compliance Inquiry that has expired, or re-open a rejected Compliance Inquiry for editing.



## CreateComplianceInquiry

> TrusthubV1ComplianceInquiry CreateComplianceInquiry(ctx, optional)

Create a new Compliance Inquiry for the authenticated account. This is necessary to start a new embedded session.

Create a new Compliance Inquiry for the authenticated account. This is necessary to start a new embedded session.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateComplianceInquiryParams struct


Name | Type | Description
------------- | ------------- | -------------
**NotificationEmail** | **string** | The email address that approval status updates will be sent to. If not specified, the email address associated with your primary customer profile will be used.
**ThemeSetId** | **string** | Theme id for styling the inquiry form.
**PrimaryProfileSid** | **string** | The unique SID identifier of the Primary Customer Profile that should be used as a parent. Only necessary when creating a secondary Customer Profile.

### Return type

[**TrusthubV1ComplianceInquiry**](TrusthubV1ComplianceInquiry.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateComplianceInquiry

> TrusthubV1ComplianceInquiry UpdateComplianceInquiry(ctx, CustomerIdoptional)

Resume a specific Compliance Inquiry that has expired, or re-open a rejected Compliance Inquiry for editing.

Resume a specific Compliance Inquiry that has expired, or re-open a rejected Compliance Inquiry for editing.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**CustomerId** | **string** | The unique CustomerId matching the Customer Profile/Compliance Inquiry that should be resumed or resubmitted. This value will have been returned by the initial Compliance Inquiry creation call.

### Other Parameters

Other parameters are passed through a pointer to a UpdateComplianceInquiryParams struct


Name | Type | Description
------------- | ------------- | -------------
**PrimaryProfileSid** | **string** | The unique SID identifier of the Primary Customer Profile that should be used as a parent. Only necessary when creating a secondary Customer Profile.
**ThemeSetId** | **string** | Theme id for styling the inquiry form.

### Return type

[**TrusthubV1ComplianceInquiry**](TrusthubV1ComplianceInquiry.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

