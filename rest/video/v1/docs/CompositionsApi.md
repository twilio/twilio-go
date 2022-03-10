# CompositionsApi

All URIs are relative to *https://video.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateComposition**](CompositionsApi.md#CreateComposition) | **Post** /v1/Compositions | 
[**DeleteComposition**](CompositionsApi.md#DeleteComposition) | **Delete** /v1/Compositions/{Sid} | 
[**FetchComposition**](CompositionsApi.md#FetchComposition) | **Get** /v1/Compositions/{Sid} | 
[**ListComposition**](CompositionsApi.md#ListComposition) | **Get** /v1/Compositions | 



## CreateComposition

> VideoV1Composition CreateComposition(ctx, optional)



### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateCompositionParams struct


Name | Type | Description
------------- | ------------- | -------------
**AudioSources** | **[]string** | An array of track names from the same group room to merge into the new composition. Can include zero or more track names. The new composition includes all audio sources specified in &#x60;audio_sources&#x60; except for those specified in &#x60;audio_sources_excluded&#x60;. The track names in this parameter can include an asterisk as a wild card character, which will match zero or more characters in a track name. For example, &#x60;student*&#x60; includes &#x60;student&#x60; as well as &#x60;studentTeam&#x60;. Please, be aware that either video_layout or audio_sources have to be provided to get a valid creation request
**AudioSourcesExcluded** | **[]string** | An array of track names to exclude. The new composition includes all audio sources specified in &#x60;audio_sources&#x60; except for those specified in &#x60;audio_sources_excluded&#x60;. The track names in this parameter can include an asterisk as a wild card character, which will match zero or more characters in a track name. For example, &#x60;student*&#x60; excludes &#x60;student&#x60; as well as &#x60;studentTeam&#x60;. This parameter can also be empty.
**Format** | **string** | The container format of the composition&#39;s media files. Can be: &#x60;mp4&#x60; or &#x60;webm&#x60; and the default is &#x60;webm&#x60;. If you specify &#x60;mp4&#x60; or &#x60;webm&#x60;, you must also specify one or more &#x60;audio_sources&#x60; and/or a &#x60;video_layout&#x60; element that contains a valid &#x60;video_sources&#x60; list, otherwise an error occurs.
**Resolution** | **string** | A string that describes the columns (width) and rows (height) of the generated composed video in pixels. Defaults to &#x60;640x480&#x60;.  The string&#39;s format is &#x60;{width}x{height}&#x60; where:   * 16 &lt;&#x3D; &#x60;{width}&#x60; &lt;&#x3D; 1280 * 16 &lt;&#x3D; &#x60;{height}&#x60; &lt;&#x3D; 1280 * &#x60;{width}&#x60; * &#x60;{height}&#x60; &lt;&#x3D; 921,600  Typical values are:   * HD &#x3D; &#x60;1280x720&#x60; * PAL &#x3D; &#x60;1024x576&#x60; * VGA &#x3D; &#x60;640x480&#x60; * CIF &#x3D; &#x60;320x240&#x60;  Note that the &#x60;resolution&#x60; imposes an aspect ratio to the resulting composition. When the original video tracks are constrained by the aspect ratio, they are scaled to fit. See [Specifying Video Layouts](https://www.twilio.com/docs/video/api/compositions-resource#specifying-video-layouts) for more info.
**RoomSid** | **string** | The SID of the Group Room with the media tracks to be used as composition sources.
**StatusCallback** | **string** | The URL we should call using the &#x60;status_callback_method&#x60; to send status information to your application on every composition event. If not provided, status callback events will not be dispatched.
**StatusCallbackMethod** | **string** | The HTTP method we should use to call &#x60;status_callback&#x60;. Can be: &#x60;POST&#x60; or &#x60;GET&#x60; and the default is &#x60;POST&#x60;.
**Trim** | **bool** | Whether to clip the intervals where there is no active media in the composition. The default is &#x60;true&#x60;. Compositions with &#x60;trim&#x60; enabled are shorter when the Room is created and no Participant joins for a while as well as if all the Participants leave the room and join later, because those gaps will be removed. See [Specifying Video Layouts](https://www.twilio.com/docs/video/api/compositions-resource#specifying-video-layouts) for more info.
**VideoLayout** | [**map[string]interface{}**](map[string]interface{}.md) | An object that describes the video layout of the composition in terms of regions. See [Specifying Video Layouts](https://www.twilio.com/docs/video/api/compositions-resource#specifying-video-layouts) for more info. Please, be aware that either video_layout or audio_sources have to be provided to get a valid creation request

### Return type

[**VideoV1Composition**](VideoV1Composition.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteComposition

> DeleteComposition(ctx, Sid)



Delete a Recording Composition resource identified by a Composition SID.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Composition resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteCompositionParams struct


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


## FetchComposition

> VideoV1Composition FetchComposition(ctx, Sid)



Returns a single Composition resource identified by a Composition SID.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Composition resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchCompositionParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**VideoV1Composition**](VideoV1Composition.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListComposition

> []VideoV1Composition ListComposition(ctx, optional)



List of all Recording compositions.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListCompositionParams struct


Name | Type | Description
------------- | ------------- | -------------
**Status** | **string** | Read only Composition resources with this status. Can be: &#x60;enqueued&#x60;, &#x60;processing&#x60;, &#x60;completed&#x60;, &#x60;deleted&#x60;, or &#x60;failed&#x60;.
**DateCreatedAfter** | **time.Time** | Read only Composition resources created on or after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date-time with time zone.
**DateCreatedBefore** | **time.Time** | Read only Composition resources created before this ISO 8601 date-time with time zone.
**RoomSid** | **string** | Read only Composition resources with this Room SID.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]VideoV1Composition**](VideoV1Composition.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

