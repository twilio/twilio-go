# ServicesUsersChannelsApi

All URIs are relative to *https://ip-messaging.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteUserChannel**](ServicesUsersChannelsApi.md#DeleteUserChannel) | **Delete** /v2/Services/{ServiceSid}/Users/{UserSid}/Channels/{ChannelSid} | 
[**FetchUserChannel**](ServicesUsersChannelsApi.md#FetchUserChannel) | **Get** /v2/Services/{ServiceSid}/Users/{UserSid}/Channels/{ChannelSid} | 
[**ListUserChannel**](ServicesUsersChannelsApi.md#ListUserChannel) | **Get** /v2/Services/{ServiceSid}/Users/{UserSid}/Channels | 
[**UpdateUserChannel**](ServicesUsersChannelsApi.md#UpdateUserChannel) | **Post** /v2/Services/{ServiceSid}/Users/{UserSid}/Channels/{ChannelSid} | 



## DeleteUserChannel

> DeleteUserChannel(ctx, ServiceSidUserSidChannelSid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 
**UserSid** | **string** | 
**ChannelSid** | **string** | 

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

> IpMessagingV2UserChannel FetchUserChannel(ctx, ServiceSidUserSidChannelSid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 
**UserSid** | **string** | 
**ChannelSid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a FetchUserChannelParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**IpMessagingV2UserChannel**](IpMessagingV2UserChannel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUserChannel

> []IpMessagingV2UserChannel ListUserChannel(ctx, ServiceSidUserSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 
**UserSid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a ListUserChannelParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]IpMessagingV2UserChannel**](IpMessagingV2UserChannel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserChannel

> IpMessagingV2UserChannel UpdateUserChannel(ctx, ServiceSidUserSidChannelSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 
**UserSid** | **string** | 
**ChannelSid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateUserChannelParams struct


Name | Type | Description
------------- | ------------- | -------------
**NotificationLevel** | **string** | 
**LastConsumedMessageIndex** | **int** | 
**LastConsumptionTimestamp** | **time.Time** | 

### Return type

[**IpMessagingV2UserChannel**](IpMessagingV2UserChannel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

