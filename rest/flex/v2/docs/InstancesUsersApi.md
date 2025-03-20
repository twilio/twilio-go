# InstancesUsersApi

All URIs are relative to *https://flex-api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchFlexUser**](InstancesUsersApi.md#FetchFlexUser) | **Get** /v2/Instances/{InstanceSid}/Users/{FlexUserSid} | Fetch flex user for the given flex user sid
[**UpdateFlexUser**](InstancesUsersApi.md#UpdateFlexUser) | **Post** /v2/Instances/{InstanceSid}/Users/{FlexUserSid} | Update flex user for the given flex user sid



## FetchFlexUser

> FlexV2FlexUser FetchFlexUser(ctx, InstanceSidFlexUserSid)

Fetch flex user for the given flex user sid

Fetch flex user for the given flex user sid

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**InstanceSid** | **string** | The unique ID created by Twilio to identify a Flex instance.
**FlexUserSid** | **string** | The unique id for the flex user to be retrieved.

### Other Parameters

Other parameters are passed through a pointer to a FetchFlexUserParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**FlexV2FlexUser**](FlexV2FlexUser.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFlexUser

> FlexV2FlexUser UpdateFlexUser(ctx, InstanceSidFlexUserSidoptional)

Update flex user for the given flex user sid

Update flex user for the given flex user sid

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**InstanceSid** | **string** | The unique ID created by Twilio to identify a Flex instance.
**FlexUserSid** | **string** | The unique id for the flex user.

### Other Parameters

Other parameters are passed through a pointer to a UpdateFlexUserParams struct


Name | Type | Description
------------- | ------------- | -------------
**Email** | **string** | Email of the User.
**UserSid** | **string** | The unique SID identifier of the Twilio Unified User.
**Locale** | **string** | The locale preference of the user.

### Return type

[**FlexV2FlexUser**](FlexV2FlexUser.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

