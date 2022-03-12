# ServicesUsersChannelsApi

All URIs are relative to *https://chat.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListUserChannel**](ServicesUsersChannelsApi.md#ListUserChannel) | **Get** /v1/Services/{ServiceSid}/Users/{UserSid}/Channels | 



## ListUserChannel

> []ChatV1UserChannel ListUserChannel(ctx, ServiceSidUserSidoptional)



List all Channels for a given User.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to read the resources from.
**UserSid** | **string** | The SID of the [User](https://www.twilio.com/docs/api/chat/rest/users) to read the User Channel resources from.

### Other Parameters

Other parameters are passed through a pointer to a ListUserChannelParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]ChatV1UserChannel**](ChatV1UserChannel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

