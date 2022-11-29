# MediaProcessorsApi

All URIs are relative to *https://media.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMediaProcessor**](MediaProcessorsApi.md#CreateMediaProcessor) | **Post** /v1/MediaProcessors | 
[**FetchMediaProcessor**](MediaProcessorsApi.md#FetchMediaProcessor) | **Get** /v1/MediaProcessors/{Sid} | 
[**ListMediaProcessor**](MediaProcessorsApi.md#ListMediaProcessor) | **Get** /v1/MediaProcessors | 
[**UpdateMediaProcessor**](MediaProcessorsApi.md#UpdateMediaProcessor) | **Post** /v1/MediaProcessors/{Sid} | 



## CreateMediaProcessor

> MediaV1MediaProcessor CreateMediaProcessor(ctx, optional)





### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateMediaProcessorParams struct


Name | Type | Description
------------- | ------------- | -------------
**Extension** | **string** | The [Media Extension](/docs/live/api/media-extensions-overview) name or URL. Ex: `video-composer-v2`
**ExtensionContext** | **string** | The context of the Media Extension, represented as a JSON dictionary. See the documentation for the specific [Media Extension](/docs/live/api/media-extensions-overview) you are using for more information about the context to send.
**ExtensionEnvironment** | [**interface{}**](interface{}.md) | User-defined environment variables for the Media Extension, represented as a JSON dictionary of key/value strings. See the documentation for the specific [Media Extension](/docs/live/api/media-extensions-overview) you are using for more information about whether you need to provide this.
**StatusCallback** | **string** | The URL to which Twilio will send asynchronous webhook requests for every MediaProcessor event. See [Status Callbacks](/docs/live/status-callbacks) for details.
**StatusCallbackMethod** | **string** | The HTTP method Twilio should use to call the `status_callback` URL. Can be `POST` or `GET` and the default is `POST`.
**MaxDuration** | **int** | The maximum time, in seconds, that the MediaProcessor can run before automatically ends. The default value is 300 seconds, and the maximum value is 90000 seconds. Once this maximum duration is reached, Twilio will end the MediaProcessor, regardless of whether media is still streaming.

### Return type

[**MediaV1MediaProcessor**](MediaV1MediaProcessor.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchMediaProcessor

> MediaV1MediaProcessor FetchMediaProcessor(ctx, Sid)



Returns a single MediaProcessor resource identified by a SID.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the MediaProcessor resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchMediaProcessorParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**MediaV1MediaProcessor**](MediaV1MediaProcessor.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMediaProcessor

> []MediaV1MediaProcessor ListMediaProcessor(ctx, optional)



Returns a list of MediaProcessors.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListMediaProcessorParams struct


Name | Type | Description
------------- | ------------- | -------------
**Order** | **string** | The sort order of the list by `date_created`. Can be: `asc` (ascending) or `desc` (descending) with `desc` as the default.
**Status** | **string** | Status to filter by, with possible values `started`, `ended` or `failed`.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]MediaV1MediaProcessor**](MediaV1MediaProcessor.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMediaProcessor

> MediaV1MediaProcessor UpdateMediaProcessor(ctx, Sidoptional)



Updates a MediaProcessor resource identified by a SID.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the MediaProcessor resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateMediaProcessorParams struct


Name | Type | Description
------------- | ------------- | -------------
**Status** | **string** | 

### Return type

[**MediaV1MediaProcessor**](MediaV1MediaProcessor.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

