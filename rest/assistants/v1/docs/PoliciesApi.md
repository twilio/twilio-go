# PoliciesApi

All URIs are relative to *https://assistants.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListPolicies**](PoliciesApi.md#ListPolicies) | **Get** /v1/Policies | List policies



## ListPolicies

> []AssistantsV1Policy ListPolicies(ctx, optional)

List policies

List policies

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListPoliciesParams struct


Name | Type | Description
------------- | ------------- | -------------
**ToolId** | **string** | The tool ID.
**KnowledgeId** | **string** | The knowledge ID.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]AssistantsV1Policy**](AssistantsV1Policy.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

