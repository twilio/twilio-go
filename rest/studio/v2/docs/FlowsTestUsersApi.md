# FlowsTestUsersApi

All URIs are relative to *https://studio.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchTestUser**](FlowsTestUsersApi.md#FetchTestUser) | **Get** /v2/Flows/{Sid}/TestUsers | 
[**UpdateTestUser**](FlowsTestUsersApi.md#UpdateTestUser) | **Post** /v2/Flows/{Sid}/TestUsers | 



## FetchTestUser

> StudioV2TestUser FetchTestUser(ctx, Sid)



Fetch flow test users

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | Unique identifier of the flow.

### Other Parameters

Other parameters are passed through a pointer to a FetchTestUserParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**StudioV2TestUser**](StudioV2TestUser.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTestUser

> StudioV2TestUser UpdateTestUser(ctx, Sidoptional)



Update flow test users

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | Unique identifier of the flow.

### Other Parameters

Other parameters are passed through a pointer to a UpdateTestUserParams struct


Name | Type | Description
------------- | ------------- | -------------
**TestUsers** | **[]string** | List of test user identities that can test draft versions of the flow.

### Return type

[**StudioV2TestUser**](StudioV2TestUser.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

