# AccountsUserRolesApi

All URIs are relative to *https://flex-api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchUserRoles**](AccountsUserRolesApi.md#FetchUserRoles) | **Get** /v1/Accounts/UserRoles | 



## FetchUserRoles

> FlexV1UserRoles FetchUserRoles(ctx, optional)



This is used by Flex UI and Quality Management to fetch the Flex Insights roles for the user

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a FetchUserRolesParams struct


Name | Type | Description
------------- | ------------- | -------------
**Token** | **string** | The Token HTTP request header

### Return type

[**FlexV1UserRoles**](FlexV1UserRoles.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

