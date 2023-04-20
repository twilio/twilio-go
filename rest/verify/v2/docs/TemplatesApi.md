# TemplatesApi

All URIs are relative to *https://verify.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListVerificationTemplate**](TemplatesApi.md#ListVerificationTemplate) | **Get** /v2/Templates | 



## ListVerificationTemplate

> []VerifyV2VerificationTemplate ListVerificationTemplate(ctx, optional)



List all the available templates for a given Account.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListVerificationTemplateParams struct


Name | Type | Description
------------- | ------------- | -------------
**FriendlyName** | **string** | String filter used to query templates with a given friendly name.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]VerifyV2VerificationTemplate**](VerifyV2VerificationTemplate.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

