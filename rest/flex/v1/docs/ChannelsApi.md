# ChannelsApi

All URIs are relative to *https://flex-api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateChannel**](ChannelsApi.md#CreateChannel) | **Post** /v1/Channels | 
[**DeleteChannel**](ChannelsApi.md#DeleteChannel) | **Delete** /v1/Channels/{Sid} | 
[**FetchChannel**](ChannelsApi.md#FetchChannel) | **Get** /v1/Channels/{Sid} | 
[**ListChannel**](ChannelsApi.md#ListChannel) | **Get** /v1/Channels | 



## CreateChannel

> FlexV1Channel CreateChannel(ctx, optional)





### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateChannelParams struct


Name | Type | Description
------------- | ------------- | -------------
**FlexFlowSid** | **string** | The SID of the Flex Flow.
**Identity** | **string** | The `identity` value that uniquely identifies the new resource's chat User.
**ChatUserFriendlyName** | **string** | The chat participant's friendly name.
**ChatFriendlyName** | **string** | The chat channel's friendly name.
**Target** | **string** | The Target Contact Identity, for example the phone number of an SMS.
**ChatUniqueName** | **string** | The chat channel's unique name.
**PreEngagementData** | **string** | The pre-engagement data.
**TaskSid** | **string** | The SID of the TaskRouter Task. Only valid when integration type is `task`. `null` for integration types `studio` & `external`
**TaskAttributes** | **string** | The Task attributes to be added for the TaskRouter Task.
**LongLived** | **bool** | Whether to create the channel as long-lived.

### Return type

[**FlexV1Channel**](FlexV1Channel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteChannel

> DeleteChannel(ctx, Sid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Flex chat channel resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteChannelParams struct


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


## FetchChannel

> FlexV1Channel FetchChannel(ctx, Sid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Flex chat channel resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchChannelParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**FlexV1Channel**](FlexV1Channel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListChannel

> []FlexV1Channel ListChannel(ctx, optional)





### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListChannelParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]FlexV1Channel**](FlexV1Channel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

