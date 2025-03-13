# ServicesOperatorsApi

All URIs are relative to *https://intelligence.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOperatorAttachment**](ServicesOperatorsApi.md#CreateOperatorAttachment) | **Post** /v2/Services/{ServiceSid}/Operators/{OperatorSid} | Attach an Operator to a Service.
[**DeleteOperatorAttachment**](ServicesOperatorsApi.md#DeleteOperatorAttachment) | **Delete** /v2/Services/{ServiceSid}/Operators/{OperatorSid} | Detach an Operator from a Service.
[**FetchOperatorAttachments**](ServicesOperatorsApi.md#FetchOperatorAttachments) | **Get** /v2/Services/{ServiceSid}/Operators | Retrieve Operators attached to a Service.



## CreateOperatorAttachment

> IntelligenceV2OperatorAttachment CreateOperatorAttachment(ctx, ServiceSidOperatorSid)

Attach an Operator to a Service.

Attach an Operator to a Service.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The unique SID identifier of the Service.
**OperatorSid** | **string** | The unique SID identifier of the Operator. Allows both Custom and Pre-built Operators.

### Other Parameters

Other parameters are passed through a pointer to a CreateOperatorAttachmentParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**IntelligenceV2OperatorAttachment**](IntelligenceV2OperatorAttachment.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOperatorAttachment

> DeleteOperatorAttachment(ctx, ServiceSidOperatorSid)

Detach an Operator from a Service.

Detach an Operator from a Service.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The unique SID identifier of the Service.
**OperatorSid** | **string** | The unique SID identifier of the Operator. Allows both Custom and Pre-built Operators.

### Other Parameters

Other parameters are passed through a pointer to a DeleteOperatorAttachmentParams struct


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


## FetchOperatorAttachments

> IntelligenceV2OperatorAttachments FetchOperatorAttachments(ctx, ServiceSid)

Retrieve Operators attached to a Service.

Retrieve Operators attached to a Service.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The unique SID identifier of the Service.

### Other Parameters

Other parameters are passed through a pointer to a FetchOperatorAttachmentsParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**IntelligenceV2OperatorAttachments**](IntelligenceV2OperatorAttachments.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

