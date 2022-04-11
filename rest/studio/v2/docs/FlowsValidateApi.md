# FlowsValidateApi

All URIs are relative to *https://studio.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateFlowValidate**](FlowsValidateApi.md#UpdateFlowValidate) | **Post** /v2/Flows/Validate | 



## UpdateFlowValidate

> StudioV2FlowValidate UpdateFlowValidate(ctx, optional)



Validate flow JSON definition

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a UpdateFlowValidateParams struct


Name | Type | Description
------------- | ------------- | -------------
**CommitMessage** | **string** | Description of change made in the revision.
**Definition** | [**interface{}**](interface{}.md) | JSON representation of flow definition.
**FriendlyName** | **string** | The string that you assigned to describe the Flow.
**Status** | **string** | The status of the Flow. Can be: &#x60;draft&#x60; or &#x60;published&#x60;.

### Return type

[**StudioV2FlowValidate**](StudioV2FlowValidate.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

