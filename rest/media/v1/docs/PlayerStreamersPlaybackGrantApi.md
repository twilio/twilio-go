# PlayerStreamersPlaybackGrantApi

All URIs are relative to *https://media.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePlayerStreamerPlaybackGrant**](PlayerStreamersPlaybackGrantApi.md#CreatePlayerStreamerPlaybackGrant) | **Post** /v1/PlayerStreamers/{Sid}/PlaybackGrant | 
[**FetchPlayerStreamerPlaybackGrant**](PlayerStreamersPlaybackGrantApi.md#FetchPlayerStreamerPlaybackGrant) | **Get** /v1/PlayerStreamers/{Sid}/PlaybackGrant | 



## CreatePlayerStreamerPlaybackGrant

> MediaV1PlayerStreamerPlaybackGrant CreatePlayerStreamerPlaybackGrant(ctx, Sidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The unique string generated to identify the PlayerStreamer resource associated with this PlaybackGrant

### Other Parameters

Other parameters are passed through a pointer to a CreatePlayerStreamerPlaybackGrantParams struct


Name | Type | Description
------------- | ------------- | -------------
**Ttl** | **int** | The time to live of the PlaybackGrant. Default value is 15 seconds. Maximum value is 60 seconds.
**AccessControlAllowOrigin** | **string** | The full origin URL where the livestream can be streamed. If this is not provided, it can be streamed from any domain.

### Return type

[**MediaV1PlayerStreamerPlaybackGrant**](MediaV1PlayerStreamerPlaybackGrant.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchPlayerStreamerPlaybackGrant

> MediaV1PlayerStreamerPlaybackGrant FetchPlayerStreamerPlaybackGrant(ctx, Sid)



**This method is not enabled.** Returns a single PlaybackGrant resource identified by a SID.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the PlayerStreamer resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchPlayerStreamerPlaybackGrantParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**MediaV1PlayerStreamerPlaybackGrant**](MediaV1PlayerStreamerPlaybackGrant.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

