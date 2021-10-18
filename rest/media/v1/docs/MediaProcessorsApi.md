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
**Extension** | **string** | The [Media Extension](/docs/live/api/media-extensions-overview) name or URL. Ex: &#x60;video-composer-v1-preview&#x60;
**ExtensionContext** | **string** | The context of the Media Extension, represented as a JSON dictionary. See the documentation for the specific [Media Extension](/docs/live/api/media-extensions-overview) you are using for more information about the context to send.
**ExtensionEnvironment** | [**map[string]interface{}**](map[string]interface{}.md) | User-defined environment variables for the Media Extension, represented as a JSON dictionary of key/value strings. See the documentation for the specific [Media Extension](/docs/live/api/media-extensions-overview) you are using for more information about whether you need to provide this.
**StatusCallback** | **string** | The URL to which Twilio will send asynchronous webhook requests for every MediaProcessor event. See [Status Callbacks](/docs/live/status-callbacks) for details.
**StatusCallbackMethod** | **string** | The HTTP method Twilio should use to call the &#x60;status_callback&#x60; URL. Can be &#x60;POST&#x60; or &#x60;GET&#x60; and the default is &#x60;POST&#x60;.

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

> ListMediaProcessorResponse ListMediaProcessor(ctx, optional)



Returns a list of MediaProcessors.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListMediaProcessorParams struct


Name | Type | Description
------------- | ------------- | -------------
**Order** | **string** | The sort order of the list by &#x60;date_created&#x60;. Can be: &#x60;asc&#x60; (ascending) or &#x60;desc&#x60; (descending) with &#x60;desc&#x60; as the default.
**Status** | **string** | Status to filter by, with possible values &#x60;started&#x60;, &#x60;ended&#x60; or &#x60;failed&#x60;.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**ListMediaProcessorResponse**](ListMediaProcessorResponse.md)

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
**Status** | **string** | The status of the MediaProcessor. Can be &#x60;ended&#x60;.

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

