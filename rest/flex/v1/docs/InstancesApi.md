# InstancesApi

All URIs are relative to *https://flex-api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInstance**](InstancesApi.md#CreateInstance) | **Post** /v1/instances | This endpoint allows you to create a Flex Instance



## CreateInstance

> FlexV1Instance CreateInstance(ctx, optional)

This endpoint allows you to create a Flex Instance

This endpoint allows you to create a Flex Instance

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateInstanceParams struct


Name | Type | Description
------------- | ------------- | -------------
**CreateInstanceRequestBody** | [**CreateInstanceRequestBody**](CreateInstanceRequestBody.md) | 

### Return type

[**FlexV1Instance**](FlexV1Instance.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

