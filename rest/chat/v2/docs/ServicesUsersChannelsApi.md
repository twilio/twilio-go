# ServicesUsersChannelsApi

All URIs are relative to *https://chat.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteUserChannel**](ServicesUsersChannelsApi.md#DeleteUserChannel) | **Delete** /v2/Services/{ServiceSid}/Users/{UserSid}/Channels/{ChannelSid} | 
[**FetchUserChannel**](ServicesUsersChannelsApi.md#FetchUserChannel) | **Get** /v2/Services/{ServiceSid}/Users/{UserSid}/Channels/{ChannelSid} | 
[**ListUserChannel**](ServicesUsersChannelsApi.md#ListUserChannel) | **Get** /v2/Services/{ServiceSid}/Users/{UserSid}/Channels | 
[**UpdateUserChannel**](ServicesUsersChannelsApi.md#UpdateUserChannel) | **Post** /v2/Services/{ServiceSid}/Users/{UserSid}/Channels/{ChannelSid} | 



## DeleteUserChannel

> DeleteUserChannel(ctx, ServiceSidUserSidChannelSid)



Removes User from selected Channel.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to read the resources from.
**UserSid** | **string** | The SID of the [User](https://www.twilio.com/docs/api/chat/rest/users) to read the User Channel resources from.
**ChannelSid** | **string** | The SID of the [Channel](https://www.twilio.com/docs/api/chat/rest/channels) the resource belongs to.

### Other Parameters

Other parameters are passed through a pointer to a DeleteUserChannelParams struct


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


## FetchUserChannel

> ChatV2ServiceUserUserChannel FetchUserChannel(ctx, ServiceSidUserSidChannelSid)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to fetch the User Channel resource from.
**UserSid** | **string** | The SID of the [User](https://www.twilio.com/docs/chat/rest/user-resource) to fetch the User Channel resource from. This value can be either the &#x60;sid&#x60; or the &#x60;identity&#x60; of the User resource.
**ChannelSid** | **string** | The SID of the [Channel](https://www.twilio.com/docs/chat/channels) that has the User Channel to fetch. This value can be either the &#x60;sid&#x60; or the &#x60;unique_name&#x60; of the Channel to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchUserChannelParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ChatV2ServiceUserUserChannel**](ChatV2ServiceUserUserChannel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUserChannel

> ListUserChannelResponse ListUserChannel(ctx, ServiceSidUserSidoptional)



List all Channels for a given User.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to read the User Channel resources from.
**UserSid** | **string** | The SID of the [User](https://www.twilio.com/docs/chat/rest/user-resource) to read the User Channel resources from. This value can be either the &#x60;sid&#x60; or the &#x60;identity&#x60; of the User resource.

### Other Parameters

Other parameters are passed through a pointer to a ListUserChannelParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListUserChannelResponse**](ListUserChannelResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserChannel

> ChatV2ServiceUserUserChannel UpdateUserChannel(ctx, ServiceSidUserSidChannelSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to update the User Channel resource in.
**UserSid** | **string** | The SID of the [User](https://www.twilio.com/docs/chat/rest/user-resource) to update the User Channel resource from. This value can be either the &#x60;sid&#x60; or the &#x60;identity&#x60; of the User resource.
**ChannelSid** | **string** | The SID of the [Channel](https://www.twilio.com/docs/chat/channels) with the User Channel resource to update. This value can be the Channel resource&#39;s &#x60;sid&#x60; or &#x60;unique_name&#x60;.

### Other Parameters

Other parameters are passed through a pointer to a UpdateUserChannelParams struct


Name | Type | Description
------------- | ------------- | -------------
**LastConsumedMessageIndex** | **int** | The index of the last [Message](https://www.twilio.com/docs/chat/rest/message-resource) in the [Channel](https://www.twilio.com/docs/chat/channels) that the Member has read.
**LastConsumptionTimestamp** | **time.Time** | The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp of the last [Message](https://www.twilio.com/docs/chat/rest/message-resource) read event for the Member within the [Channel](https://www.twilio.com/docs/chat/channels).
**NotificationLevel** | **string** | The push notification level to assign to the User Channel. Can be: &#x60;default&#x60; or &#x60;muted&#x60;.

### Return type

[**ChatV2ServiceUserUserChannel**](ChatV2ServiceUserUserChannel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

