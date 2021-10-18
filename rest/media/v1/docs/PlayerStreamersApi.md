# PlayerStreamersApi

All URIs are relative to *https://media.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePlayerStreamer**](PlayerStreamersApi.md#CreatePlayerStreamer) | **Post** /v1/PlayerStreamers | 
[**FetchPlayerStreamer**](PlayerStreamersApi.md#FetchPlayerStreamer) | **Get** /v1/PlayerStreamers/{Sid} | 
[**ListPlayerStreamer**](PlayerStreamersApi.md#ListPlayerStreamer) | **Get** /v1/PlayerStreamers | 
[**UpdatePlayerStreamer**](PlayerStreamersApi.md#UpdatePlayerStreamer) | **Post** /v1/PlayerStreamers/{Sid} | 



## CreatePlayerStreamer

> MediaV1PlayerStreamer CreatePlayerStreamer(ctx, optional)



### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreatePlayerStreamerParams struct


Name | Type | Description
------------- | ------------- | -------------
**StatusCallback** | **string** | The URL to which Twilio will send asynchronous webhook requests for every PlayerStreamer event. See [Status Callbacks](/docs/live/status-callbacks) for more details.
**StatusCallbackMethod** | **string** | The HTTP method Twilio should use to call the &#x60;status_callback&#x60; URL. Can be &#x60;POST&#x60; or &#x60;GET&#x60; and the default is &#x60;POST&#x60;.
**Video** | **bool** | Specifies whether the PlayerStreamer is configured to stream video. Defaults to &#x60;true&#x60;.

### Return type

[**MediaV1PlayerStreamer**](MediaV1PlayerStreamer.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchPlayerStreamer

> MediaV1PlayerStreamer FetchPlayerStreamer(ctx, Sid)



Returns a single PlayerStreamer resource identified by a SID.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the PlayerStreamer resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchPlayerStreamerParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**MediaV1PlayerStreamer**](MediaV1PlayerStreamer.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPlayerStreamer

> ListPlayerStreamerResponse ListPlayerStreamer(ctx, optional)



Returns a list of PlayerStreamers.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListPlayerStreamerParams struct


Name | Type | Description
------------- | ------------- | -------------
**Order** | **string** | The sort order of the list by &#x60;date_created&#x60;. Can be: &#x60;asc&#x60; (ascending) or &#x60;desc&#x60; (descending) with &#x60;desc&#x60; as the default.
**Status** | **string** | Status to filter by, with possible values &#x60;created&#x60;, &#x60;started&#x60;, &#x60;ended&#x60;, or &#x60;failed&#x60;.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**ListPlayerStreamerResponse**](ListPlayerStreamerResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePlayerStreamer

> MediaV1PlayerStreamer UpdatePlayerStreamer(ctx, Sidoptional)



Updates a PlayerStreamer resource identified by a SID.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the PlayerStreamer resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdatePlayerStreamerParams struct


Name | Type | Description
------------- | ------------- | -------------
**Status** | **string** | The status the PlayerStreamer should be transitioned to. Can be: &#x60;ended&#x60;.

### Return type

[**MediaV1PlayerStreamer**](MediaV1PlayerStreamer.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

