# FormsApi

All URIs are relative to *https://verify.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchForm**](FormsApi.md#FetchForm) | **Get** /v2/Forms/{FormType} | 



## FetchForm

> VerifyV2Form FetchForm(ctx, FormType)



Fetch the forms for a specific Form Type.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FormType** | **string** | The Type of this Form. Currently only `form-push` is supported.

### Other Parameters

Other parameters are passed through a pointer to a FetchFormParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**VerifyV2Form**](VerifyV2Form.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

