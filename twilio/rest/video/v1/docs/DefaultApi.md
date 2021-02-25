# \DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateComposition**](DefaultApi.md#CreateComposition) | **Post** /v1/Compositions | 
[**CreateCompositionHook**](DefaultApi.md#CreateCompositionHook) | **Post** /v1/CompositionHooks | 
[**CreateCompositionSettings**](DefaultApi.md#CreateCompositionSettings) | **Post** /v1/CompositionSettings/Default | 
[**CreateRecordingSettings**](DefaultApi.md#CreateRecordingSettings) | **Post** /v1/RecordingSettings/Default | 
[**CreateRoom**](DefaultApi.md#CreateRoom) | **Post** /v1/Rooms | 
[**DeleteComposition**](DefaultApi.md#DeleteComposition) | **Delete** /v1/Compositions/{Sid} | 
[**DeleteCompositionHook**](DefaultApi.md#DeleteCompositionHook) | **Delete** /v1/CompositionHooks/{Sid} | 
[**DeleteRecording**](DefaultApi.md#DeleteRecording) | **Delete** /v1/Recordings/{Sid} | 
[**DeleteRoomRecording**](DefaultApi.md#DeleteRoomRecording) | **Delete** /v1/Rooms/{RoomSid}/Recordings/{Sid} | 
[**FetchComposition**](DefaultApi.md#FetchComposition) | **Get** /v1/Compositions/{Sid} | 
[**FetchCompositionHook**](DefaultApi.md#FetchCompositionHook) | **Get** /v1/CompositionHooks/{Sid} | 
[**FetchCompositionSettings**](DefaultApi.md#FetchCompositionSettings) | **Get** /v1/CompositionSettings/Default | 
[**FetchRecording**](DefaultApi.md#FetchRecording) | **Get** /v1/Recordings/{Sid} | 
[**FetchRecordingSettings**](DefaultApi.md#FetchRecordingSettings) | **Get** /v1/RecordingSettings/Default | 
[**FetchRoom**](DefaultApi.md#FetchRoom) | **Get** /v1/Rooms/{Sid} | 
[**FetchRoomParticipant**](DefaultApi.md#FetchRoomParticipant) | **Get** /v1/Rooms/{RoomSid}/Participants/{Sid} | 
[**FetchRoomParticipantPublishedTrack**](DefaultApi.md#FetchRoomParticipantPublishedTrack) | **Get** /v1/Rooms/{RoomSid}/Participants/{ParticipantSid}/PublishedTracks/{Sid} | 
[**FetchRoomParticipantSubscribeRule**](DefaultApi.md#FetchRoomParticipantSubscribeRule) | **Get** /v1/Rooms/{RoomSid}/Participants/{ParticipantSid}/SubscribeRules | 
[**FetchRoomParticipantSubscribedTrack**](DefaultApi.md#FetchRoomParticipantSubscribedTrack) | **Get** /v1/Rooms/{RoomSid}/Participants/{ParticipantSid}/SubscribedTracks/{Sid} | 
[**FetchRoomRecording**](DefaultApi.md#FetchRoomRecording) | **Get** /v1/Rooms/{RoomSid}/Recordings/{Sid} | 
[**FetchRoomRecordingRule**](DefaultApi.md#FetchRoomRecordingRule) | **Get** /v1/Rooms/{RoomSid}/RecordingRules | 
[**ListComposition**](DefaultApi.md#ListComposition) | **Get** /v1/Compositions | 
[**ListCompositionHook**](DefaultApi.md#ListCompositionHook) | **Get** /v1/CompositionHooks | 
[**ListRecording**](DefaultApi.md#ListRecording) | **Get** /v1/Recordings | 
[**ListRoom**](DefaultApi.md#ListRoom) | **Get** /v1/Rooms | 
[**ListRoomParticipant**](DefaultApi.md#ListRoomParticipant) | **Get** /v1/Rooms/{RoomSid}/Participants | 
[**ListRoomParticipantPublishedTrack**](DefaultApi.md#ListRoomParticipantPublishedTrack) | **Get** /v1/Rooms/{RoomSid}/Participants/{ParticipantSid}/PublishedTracks | 
[**ListRoomParticipantSubscribedTrack**](DefaultApi.md#ListRoomParticipantSubscribedTrack) | **Get** /v1/Rooms/{RoomSid}/Participants/{ParticipantSid}/SubscribedTracks | 
[**ListRoomRecording**](DefaultApi.md#ListRoomRecording) | **Get** /v1/Rooms/{RoomSid}/Recordings | 
[**UpdateCompositionHook**](DefaultApi.md#UpdateCompositionHook) | **Post** /v1/CompositionHooks/{Sid} | 
[**UpdateRoom**](DefaultApi.md#UpdateRoom) | **Post** /v1/Rooms/{Sid} | 
[**UpdateRoomParticipant**](DefaultApi.md#UpdateRoomParticipant) | **Post** /v1/Rooms/{RoomSid}/Participants/{Sid} | 
[**UpdateRoomParticipantSubscribeRule**](DefaultApi.md#UpdateRoomParticipantSubscribeRule) | **Post** /v1/Rooms/{RoomSid}/Participants/{ParticipantSid}/SubscribeRules | 
[**UpdateRoomRecordingRule**](DefaultApi.md#UpdateRoomRecordingRule) | **Post** /v1/Rooms/{RoomSid}/RecordingRules | 



## CreateComposition

> VideoV1Composition CreateComposition(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateCompositionRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateCompositionRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**AudioSources** | [**[]string**](string.md)| An array of track names from the same group room to merge into the new composition. Can include zero or more track names. The new composition includes all audio sources specified in &#x60;audio_sources&#x60; except for those specified in &#x60;audio_sources_excluded&#x60;. The track names in this parameter can include an asterisk as a wild card character, which will match zero or more characters in a track name. For example, &#x60;student*&#x60; includes &#x60;student&#x60; as well as &#x60;studentTeam&#x60;. Please, be aware that either video_layout or audio_sources have to be provided to get a valid creation request | 
**AudioSourcesExcluded** | [**[]string**](string.md)| An array of track names to exclude. The new composition includes all audio sources specified in &#x60;audio_sources&#x60; except for those specified in &#x60;audio_sources_excluded&#x60;. The track names in this parameter can include an asterisk as a wild card character, which will match zero or more characters in a track name. For example, &#x60;student*&#x60; excludes &#x60;student&#x60; as well as &#x60;studentTeam&#x60;. This parameter can also be empty. | 
**Format** | **String**| The container format of the composition&#39;s media files. Can be: &#x60;mp4&#x60; or &#x60;webm&#x60; and the default is &#x60;webm&#x60;. If you specify &#x60;mp4&#x60; or &#x60;webm&#x60;, you must also specify one or more &#x60;audio_sources&#x60; and/or a &#x60;video_layout&#x60; element that contains a valid &#x60;video_sources&#x60; list, otherwise an error occurs. | 
**Resolution** | **String**| A string that describes the columns (width) and rows (height) of the generated composed video in pixels. Defaults to &#x60;640x480&#x60;.  The string&#39;s format is &#x60;{width}x{height}&#x60; where:   * 16 &lt;&#x3D; &#x60;{width}&#x60; &lt;&#x3D; 1280 * 16 &lt;&#x3D; &#x60;{height}&#x60; &lt;&#x3D; 1280 * &#x60;{width}&#x60; * &#x60;{height}&#x60; &lt;&#x3D; 921,600  Typical values are:   * HD &#x3D; &#x60;1280x720&#x60; * PAL &#x3D; &#x60;1024x576&#x60; * VGA &#x3D; &#x60;640x480&#x60; * CIF &#x3D; &#x60;320x240&#x60;  Note that the &#x60;resolution&#x60; imposes an aspect ratio to the resulting composition. When the original video tracks are constrained by the aspect ratio, they are scaled to fit. See [Specifying Video Layouts](https://www.twilio.com/docs/video/api/compositions-resource#specifying-video-layouts) for more info. | 
**RoomSid** | **String**| The SID of the Group Room with the media tracks to be used as composition sources. | 
**StatusCallback** | **String**| The URL we should call using the &#x60;status_callback_method&#x60; to send status information to your application on every composition event. If not provided, status callback events will not be dispatched. | 
**StatusCallbackMethod** | **String**| The HTTP method we should use to call &#x60;status_callback&#x60;. Can be: &#x60;POST&#x60; or &#x60;GET&#x60; and the default is &#x60;POST&#x60;. | 
**Trim** | **Bool**| Whether to clip the intervals where there is no active media in the composition. The default is &#x60;true&#x60;. Compositions with &#x60;trim&#x60; enabled are shorter when the Room is created and no Participant joins for a while as well as if all the Participants leave the room and join later, because those gaps will be removed. See [Specifying Video Layouts](https://www.twilio.com/docs/video/api/compositions-resource#specifying-video-layouts) for more info. | 
**VideoLayout** | [**map[string]interface{}**](map[string]interface{}.md)| An object that describes the video layout of the composition in terms of regions. See [Specifying Video Layouts](https://www.twilio.com/docs/video/api/compositions-resource#specifying-video-layouts) for more info. Please, be aware that either video_layout or audio_sources have to be provided to get a valid creation request | 

### Return type

[**VideoV1Composition**](video.v1.composition.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCompositionHook

> VideoV1CompositionHook CreateCompositionHook(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateCompositionHookRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateCompositionHookRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**AudioSources** | [**[]string**](string.md)| An array of track names from the same group room to merge into the compositions created by the composition hook. Can include zero or more track names. A composition triggered by the composition hook includes all audio sources specified in &#x60;audio_sources&#x60; except those specified in &#x60;audio_sources_excluded&#x60;. The track names in this parameter can include an asterisk as a wild card character, which matches zero or more characters in a track name. For example, &#x60;student*&#x60; includes tracks named &#x60;student&#x60; as well as &#x60;studentTeam&#x60;. | 
**AudioSourcesExcluded** | [**[]string**](string.md)| An array of track names to exclude. A composition triggered by the composition hook includes all audio sources specified in &#x60;audio_sources&#x60; except for those specified in &#x60;audio_sources_excluded&#x60;. The track names in this parameter can include an asterisk as a wild card character, which matches zero or more characters in a track name. For example, &#x60;student*&#x60; excludes &#x60;student&#x60; as well as &#x60;studentTeam&#x60;. This parameter can also be empty. | 
**Enabled** | **Bool**| Whether the composition hook is active. When &#x60;true&#x60;, the composition hook will be triggered for every completed Group Room in the account. When &#x60;false&#x60;, the composition hook will never be triggered. | 
**Format** | **String**| The container format of the media files used by the compositions created by the composition hook. Can be: &#x60;mp4&#x60; or &#x60;webm&#x60; and the default is &#x60;webm&#x60;. If &#x60;mp4&#x60; or &#x60;webm&#x60;, &#x60;audio_sources&#x60; must have one or more tracks and/or a &#x60;video_layout&#x60; element must contain a valid &#x60;video_sources&#x60; list, otherwise an error occurs. | 
**FriendlyName** | **String**| A descriptive string that you create to describe the resource. It can be up to  100 characters long and it must be unique within the account. | 
**Resolution** | **String**| A string that describes the columns (width) and rows (height) of the generated composed video in pixels. Defaults to &#x60;640x480&#x60;.  The string&#39;s format is &#x60;{width}x{height}&#x60; where:   * 16 &lt;&#x3D; &#x60;{width}&#x60; &lt;&#x3D; 1280 * 16 &lt;&#x3D; &#x60;{height}&#x60; &lt;&#x3D; 1280 * &#x60;{width}&#x60; * &#x60;{height}&#x60; &lt;&#x3D; 921,600  Typical values are:   * HD &#x3D; &#x60;1280x720&#x60; * PAL &#x3D; &#x60;1024x576&#x60; * VGA &#x3D; &#x60;640x480&#x60; * CIF &#x3D; &#x60;320x240&#x60;  Note that the &#x60;resolution&#x60; imposes an aspect ratio to the resulting composition. When the original video tracks are constrained by the aspect ratio, they are scaled to fit. See [Specifying Video Layouts](https://www.twilio.com/docs/video/api/compositions-resource#specifying-video-layouts) for more info. | 
**StatusCallback** | **String**| The URL we should call using the &#x60;status_callback_method&#x60; to send status information to your application on every composition event. If not provided, status callback events will not be dispatched. | 
**StatusCallbackMethod** | **String**| The HTTP method we should use to call &#x60;status_callback&#x60;. Can be: &#x60;POST&#x60; or &#x60;GET&#x60; and the default is &#x60;POST&#x60;. | 
**Trim** | **Bool**| Whether to clip the intervals where there is no active media in the Compositions triggered by the composition hook. The default is &#x60;true&#x60;. Compositions with &#x60;trim&#x60; enabled are shorter when the Room is created and no Participant joins for a while as well as if all the Participants leave the room and join later, because those gaps will be removed. See [Specifying Video Layouts](https://www.twilio.com/docs/video/api/compositions-resource#specifying-video-layouts) for more info. | 
**VideoLayout** | [**map[string]interface{}**](map[string]interface{}.md)| An object that describes the video layout of the composition hook in terms of regions. See [Specifying Video Layouts](https://www.twilio.com/docs/video/api/compositions-resource#specifying-video-layouts) for more info. | 

### Return type

[**VideoV1CompositionHook**](video.v1.composition_hook.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCompositionSettings

> VideoV1CompositionSettings CreateCompositionSettings(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateCompositionSettingsRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateCompositionSettingsRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**AwsCredentialsSid** | **String**| The SID of the stored Credential resource. | 
**AwsS3Url** | **String**| The URL of the AWS S3 bucket where the compositions should be stored. We only support DNS-compliant URLs like &#x60;https://documentation-example-twilio-bucket/compositions&#x60;, where &#x60;compositions&#x60; is the path in which you want the compositions to be stored. This URL accepts only URI-valid characters, as described in the &lt;a href&#x3D;&#39;https://tools.ietf.org/html/rfc3986#section-2&#39;&gt;RFC 3986&lt;/a&gt;. | 
**AwsStorageEnabled** | **Bool**| Whether all compositions should be written to the &#x60;aws_s3_url&#x60;. When &#x60;false&#x60;, all compositions are stored in our cloud. | 
**EncryptionEnabled** | **Bool**| Whether all compositions should be stored in an encrypted form. The default is &#x60;false&#x60;. | 
**EncryptionKeySid** | **String**| The SID of the Public Key resource to use for encryption. | 
**FriendlyName** | **String**| A descriptive string that you create to describe the resource and show to the user in the console | 

### Return type

[**VideoV1CompositionSettings**](video.v1.composition_settings.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRecordingSettings

> VideoV1RecordingSettings CreateRecordingSettings(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateRecordingSettingsRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateRecordingSettingsRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**AwsCredentialsSid** | **String**| The SID of the stored Credential resource. | 
**AwsS3Url** | **String**| The URL of the AWS S3 bucket where the recordings should be stored. We only support DNS-compliant URLs like &#x60;https://documentation-example-twilio-bucket/recordings&#x60;, where &#x60;recordings&#x60; is the path in which you want the recordings to be stored. This URL accepts only URI-valid characters, as described in the &lt;a href&#x3D;&#39;https://tools.ietf.org/html/rfc3986#section-2&#39;&gt;RFC 3986&lt;/a&gt;. | 
**AwsStorageEnabled** | **Bool**| Whether all recordings should be written to the &#x60;aws_s3_url&#x60;. When &#x60;false&#x60;, all recordings are stored in our cloud. | 
**EncryptionEnabled** | **Bool**| Whether all recordings should be stored in an encrypted form. The default is &#x60;false&#x60;. | 
**EncryptionKeySid** | **String**| The SID of the Public Key resource to use for encryption. | 
**FriendlyName** | **String**| A descriptive string that you create to describe the resource and be shown to users in the console | 

### Return type

[**VideoV1RecordingSettings**](video.v1.recording_settings.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRoom

> VideoV1Room CreateRoom(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateRoomRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateRoomRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**EnableTurn** | **Bool**| Deprecated, now always considered to be true. | 
**MaxParticipants** | **Int32**| The maximum number of concurrent Participants allowed in the room. Peer-to-peer rooms can have up to 10 Participants. Small Group rooms can have up to 4 Participants. Group rooms can have up to 50 Participants. | 
**MediaRegion** | **String**| The region for the media server in Group Rooms.  Can be: one of the [available Media Regions](https://www.twilio.com/docs/video/ip-address-whitelisting#group-rooms-media-servers). ***This feature is not available in &#x60;peer-to-peer&#x60; rooms.*** | 
**RecordParticipantsOnConnect** | **Bool**| Whether to start recording when Participants connect. ***This feature is not available in &#x60;peer-to-peer&#x60; rooms.*** | 
**StatusCallback** | **String**| The URL we should call using the &#x60;status_callback_method&#x60; to send status information to your application on every room event. See [Status Callbacks](https://www.twilio.com/docs/video/api/status-callbacks) for more info. | 
**StatusCallbackMethod** | **String**| The HTTP method we should use to call &#x60;status_callback&#x60;. Can be &#x60;POST&#x60; or &#x60;GET&#x60;. | 
**Type** | **String**| The type of room. Can be: &#x60;go&#x60;, &#x60;peer-to-peer&#x60;, &#x60;group-small&#x60;, or &#x60;group&#x60;. The default value is &#x60;group&#x60;. | 
**UniqueName** | **String**| An application-defined string that uniquely identifies the resource. It can be used as a &#x60;room_sid&#x60; in place of the resource&#39;s &#x60;sid&#x60; in the URL to address the resource. This value is unique for &#x60;in-progress&#x60; rooms. SDK clients can use this name to connect to the room. REST API clients can use this name in place of the Room SID to interact with the room as long as the room is &#x60;in-progress&#x60;. | 
**VideoCodecs** | [**[]string**](string.md)| An array of the video codecs that are supported when publishing a track in the room.  Can be: &#x60;VP8&#x60; and &#x60;H264&#x60;.  ***This feature is not available in &#x60;peer-to-peer&#x60; rooms*** | 

### Return type

[**VideoV1Room**](video.v1.room.md)

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

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| The SID of the Composition resource to delete. | 

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


## DeleteCompositionHook

> DeleteCompositionHook(ctx, Sid)



Delete a Recording CompositionHook resource identified by a `CompositionHook SID`.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| The SID of the CompositionHook resource to delete. | 

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


## DeleteRecording

> DeleteRecording(ctx, Sid)



Delete a Recording resource identified by a Recording SID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| The SID of the Recording resource to delete. | 

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


## DeleteRoomRecording

> DeleteRoomRecording(ctx, RoomSid, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**RoomSid** | **string**| The SID of the room with the RoomRecording resource to delete. | 
**Sid** | **string**| The SID of the RoomRecording resource to delete. | 

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

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| The SID of the Composition resource to fetch. | 

### Return type

[**VideoV1Composition**](video.v1.composition.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchCompositionHook

> VideoV1CompositionHook FetchCompositionHook(ctx, Sid)



Returns a single CompositionHook resource identified by a CompositionHook SID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| The SID of the CompositionHook resource to fetch. | 

### Return type

[**VideoV1CompositionHook**](video.v1.composition_hook.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchCompositionSettings

> VideoV1CompositionSettings FetchCompositionSettings(ctx, )



### Required Parameters

This endpoint does not need any parameter.

### Return type

[**VideoV1CompositionSettings**](video.v1.composition_settings.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchRecording

> VideoV1Recording FetchRecording(ctx, Sid)



Returns a single Recording resource identified by a Recording SID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| The SID of the Recording resource to fetch. | 

### Return type

[**VideoV1Recording**](video.v1.recording.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchRecordingSettings

> VideoV1RecordingSettings FetchRecordingSettings(ctx, )



### Required Parameters

This endpoint does not need any parameter.

### Return type

[**VideoV1RecordingSettings**](video.v1.recording_settings.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchRoom

> VideoV1Room FetchRoom(ctx, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| The SID of the Room resource to fetch. | 

### Return type

[**VideoV1Room**](video.v1.room.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchRoomParticipant

> VideoV1RoomRoomParticipant FetchRoomParticipant(ctx, RoomSid, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**RoomSid** | **string**| The SID of the room with the Participant resource to fetch. | 
**Sid** | **string**| The SID of the RoomParticipant resource to fetch. | 

### Return type

[**VideoV1RoomRoomParticipant**](video.v1.room.room_participant.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchRoomParticipantPublishedTrack

> VideoV1RoomRoomParticipantRoomParticipantPublishedTrack FetchRoomParticipantPublishedTrack(ctx, RoomSid, ParticipantSid, Sid)



Returns a single Track resource represented by TrackName or SID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**RoomSid** | **string**| The SID of the Room resource where the Track resource to fetch is published. | 
**ParticipantSid** | **string**| The SID of the Participant resource with the published track to fetch. | 
**Sid** | **string**| The SID of the RoomParticipantPublishedTrack resource to fetch. | 

### Return type

[**VideoV1RoomRoomParticipantRoomParticipantPublishedTrack**](video.v1.room.room_participant.room_participant_published_track.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchRoomParticipantSubscribeRule

> VideoV1RoomRoomParticipantRoomParticipantSubscribeRule FetchRoomParticipantSubscribeRule(ctx, RoomSid, ParticipantSid)



Returns a list of Subscribe Rules for the Participant.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**RoomSid** | **string**| The SID of the Room resource where the subscribe rules to fetch apply. | 
**ParticipantSid** | **string**| The SID of the Participant resource with the subscribe rules to fetch. | 

### Return type

[**VideoV1RoomRoomParticipantRoomParticipantSubscribeRule**](video.v1.room.room_participant.room_participant_subscribe_rule.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchRoomParticipantSubscribedTrack

> VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack FetchRoomParticipantSubscribedTrack(ctx, RoomSid, ParticipantSid, Sid)



Returns a single Track resource represented by `track_sid`.  Note: This is one resource with the Video API that requires a SID, be Track Name on the subscriber side is not guaranteed to be unique.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**RoomSid** | **string**| The SID of the Room where the Track resource to fetch is subscribed. | 
**ParticipantSid** | **string**| The SID of the participant that subscribes to the Track resource to fetch. | 
**Sid** | **string**| The SID of the RoomParticipantSubscribedTrack resource to fetch. | 

### Return type

[**VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack**](video.v1.room.room_participant.room_participant_subscribed_track.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchRoomRecording

> VideoV1RoomRoomRecording FetchRoomRecording(ctx, RoomSid, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**RoomSid** | **string**| The SID of the Room resource with the recording to fetch. | 
**Sid** | **string**| The SID of the RoomRecording resource to fetch. | 

### Return type

[**VideoV1RoomRoomRecording**](video.v1.room.room_recording.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchRoomRecordingRule

> VideoV1RoomRoomRecordingRule FetchRoomRecordingRule(ctx, RoomSid)



Returns a list of Recording Rules for the Room.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**RoomSid** | **string**| The SID of the Room resource where the recording rules to fetch apply. | 

### Return type

[**VideoV1RoomRoomRecordingRule**](video.v1.room.room_recording_rule.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListComposition

> ListCompositionResponse ListComposition(ctx, optional)



List of all Recording compositions.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListCompositionRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListCompositionRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Status** | **String**| Read only Composition resources with this status. Can be: &#x60;enqueued&#x60;, &#x60;processing&#x60;, &#x60;completed&#x60;, &#x60;deleted&#x60;, or &#x60;failed&#x60;. | 
**DateCreatedAfter** | **Time**| Read only Composition resources created on or after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date-time with time zone. | 
**DateCreatedBefore** | **Time**| Read only Composition resources created before this ISO 8601 date-time with time zone. | 
**RoomSid** | **String**| Read only Composition resources with this Room SID. | 
**PageSize** | **Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListCompositionResponse**](ListCompositionResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCompositionHook

> ListCompositionHookResponse ListCompositionHook(ctx, optional)



List of all Recording CompositionHook resources.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListCompositionHookRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListCompositionHookRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Enabled** | **Bool**| Read only CompositionHook resources with an &#x60;enabled&#x60; value that matches this parameter. | 
**DateCreatedAfter** | **Time**| Read only CompositionHook resources created on or after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) datetime with time zone. | 
**DateCreatedBefore** | **Time**| Read only CompositionHook resources created before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) datetime with time zone. | 
**FriendlyName** | **String**| Read only CompositionHook resources with friendly names that match this string. The match is not case sensitive and can include asterisk &#x60;*&#x60; characters as wildcard match. | 
**PageSize** | **Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListCompositionHookResponse**](ListCompositionHookResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRecording

> ListRecordingResponse ListRecording(ctx, optional)



List of all Track recordings.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListRecordingRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListRecordingRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Status** | **String**| Read only the recordings that have this status. Can be: &#x60;processing&#x60;, &#x60;completed&#x60;, or &#x60;deleted&#x60;. | 
**SourceSid** | **String**| Read only the recordings that have this &#x60;source_sid&#x60;. | 
**GroupingSid** | [**[]string**](string.md)| Read only recordings with this &#x60;grouping_sid&#x60;, which may include a &#x60;participant_sid&#x60; and/or a &#x60;room_sid&#x60;. | 
**DateCreatedAfter** | **Time**| Read only recordings that started on or after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date-time with time zone. | 
**DateCreatedBefore** | **Time**| Read only recordings that started before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date-time with time zone, given as &#x60;YYYY-MM-DDThh:mm:ss+|-hh:mm&#x60; or &#x60;YYYY-MM-DDThh:mm:ssZ&#x60;. | 
**MediaType** | **String**| Read only recordings that have this media type. Can be either &#x60;audio&#x60; or &#x60;video&#x60;. | 
**PageSize** | **Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListRecordingResponse**](ListRecordingResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRoom

> ListRoomResponse ListRoom(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListRoomRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListRoomRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Status** | **String**| Read only the rooms with this status. Can be: &#x60;in-progress&#x60; (default) or &#x60;completed&#x60; | 
**UniqueName** | **String**| Read only rooms with the this &#x60;unique_name&#x60;. | 
**DateCreatedAfter** | **Time**| Read only rooms that started on or after this date, given as &#x60;YYYY-MM-DD&#x60;. | 
**DateCreatedBefore** | **Time**| Read only rooms that started before this date, given as &#x60;YYYY-MM-DD&#x60;. | 
**PageSize** | **Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListRoomResponse**](ListRoomResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRoomParticipant

> ListRoomParticipantResponse ListRoomParticipant(ctx, RoomSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**RoomSid** | **string**| The SID of the room with the Participant resources to read. | 
 **optional** | ***ListRoomParticipantRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListRoomParticipantRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Status** | **String**| Read only the participants with this status. Can be: &#x60;connected&#x60; or &#x60;disconnected&#x60;. For &#x60;in-progress&#x60; Rooms the default Status is &#x60;connected&#x60;, for &#x60;completed&#x60; Rooms only &#x60;disconnected&#x60; Participants are returned. | 
**Identity** | **String**| Read only the Participants with this [User](https://www.twilio.com/docs/chat/rest/user-resource) &#x60;identity&#x60; value. | 
**DateCreatedAfter** | **Time**| Read only Participants that started after this date in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601#UTC) format. | 
**DateCreatedBefore** | **Time**| Read only Participants that started before this date in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601#UTC) format. | 
**PageSize** | **Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListRoomParticipantResponse**](ListRoomParticipantResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRoomParticipantPublishedTrack

> ListRoomParticipantPublishedTrackResponse ListRoomParticipantPublishedTrack(ctx, RoomSid, ParticipantSid, optional)



Returns a list of tracks associated with a given Participant. Only `currently` Published Tracks are in the list resource.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**RoomSid** | **string**| The SID of the Room resource where the Track resources to read are published. | 
**ParticipantSid** | **string**| The SID of the Participant resource with the published tracks to read. | 
 **optional** | ***ListRoomParticipantPublishedTrackRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListRoomParticipantPublishedTrackRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PageSize** | **Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListRoomParticipantPublishedTrackResponse**](ListRoomParticipantPublishedTrackResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRoomParticipantSubscribedTrack

> ListRoomParticipantSubscribedTrackResponse ListRoomParticipantSubscribedTrack(ctx, RoomSid, ParticipantSid, optional)



Returns a list of tracks that are subscribed for the participant.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**RoomSid** | **string**| The SID of the Room resource with the Track resources to read. | 
**ParticipantSid** | **string**| The SID of the participant that subscribes to the Track resources to read. | 
 **optional** | ***ListRoomParticipantSubscribedTrackRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListRoomParticipantSubscribedTrackRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PageSize** | **Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListRoomParticipantSubscribedTrackResponse**](ListRoomParticipantSubscribedTrackResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRoomRecording

> ListRoomRecordingResponse ListRoomRecording(ctx, RoomSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**RoomSid** | **string**| The SID of the room with the RoomRecording resources to read. | 
 **optional** | ***ListRoomRecordingRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListRoomRecordingRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Status** | **String**| Read only the recordings with this status. Can be: &#x60;processing&#x60;, &#x60;completed&#x60;, or &#x60;deleted&#x60;. | 
**SourceSid** | **String**| Read only the recordings that have this &#x60;source_sid&#x60;. | 
**DateCreatedAfter** | **Time**| Read only recordings that started on or after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) datetime with time zone. | 
**DateCreatedBefore** | **Time**| Read only Recordings that started before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) datetime with time zone. | 
**PageSize** | **Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListRoomRecordingResponse**](ListRoomRecordingResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCompositionHook

> VideoV1CompositionHook UpdateCompositionHook(ctx, Sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| The SID of the CompositionHook resource to update. | 
 **optional** | ***UpdateCompositionHookRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateCompositionHookRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**AudioSources** | [**[]string**](string.md)| An array of track names from the same group room to merge into the compositions created by the composition hook. Can include zero or more track names. A composition triggered by the composition hook includes all audio sources specified in &#x60;audio_sources&#x60; except those specified in &#x60;audio_sources_excluded&#x60;. The track names in this parameter can include an asterisk as a wild card character, which matches zero or more characters in a track name. For example, &#x60;student*&#x60; includes tracks named &#x60;student&#x60; as well as &#x60;studentTeam&#x60;. | 
**AudioSourcesExcluded** | [**[]string**](string.md)| An array of track names to exclude. A composition triggered by the composition hook includes all audio sources specified in &#x60;audio_sources&#x60; except for those specified in &#x60;audio_sources_excluded&#x60;. The track names in this parameter can include an asterisk as a wild card character, which matches zero or more characters in a track name. For example, &#x60;student*&#x60; excludes &#x60;student&#x60; as well as &#x60;studentTeam&#x60;. This parameter can also be empty. | 
**Enabled** | **Bool**| Whether the composition hook is active. When &#x60;true&#x60;, the composition hook will be triggered for every completed Group Room in the account. When &#x60;false&#x60;, the composition hook never triggers. | 
**Format** | **String**| The container format of the media files used by the compositions created by the composition hook. Can be: &#x60;mp4&#x60; or &#x60;webm&#x60; and the default is &#x60;webm&#x60;. If &#x60;mp4&#x60; or &#x60;webm&#x60;, &#x60;audio_sources&#x60; must have one or more tracks and/or a &#x60;video_layout&#x60; element must contain a valid &#x60;video_sources&#x60; list, otherwise an error occurs. | 
**FriendlyName** | **String**| A descriptive string that you create to describe the resource. It can be up to  100 characters long and it must be unique within the account. | 
**Resolution** | **String**| A string that describes the columns (width) and rows (height) of the generated composed video in pixels. Defaults to &#x60;640x480&#x60;.  The string&#39;s format is &#x60;{width}x{height}&#x60; where:   * 16 &lt;&#x3D; &#x60;{width}&#x60; &lt;&#x3D; 1280 * 16 &lt;&#x3D; &#x60;{height}&#x60; &lt;&#x3D; 1280 * &#x60;{width}&#x60; * &#x60;{height}&#x60; &lt;&#x3D; 921,600  Typical values are:   * HD &#x3D; &#x60;1280x720&#x60; * PAL &#x3D; &#x60;1024x576&#x60; * VGA &#x3D; &#x60;640x480&#x60; * CIF &#x3D; &#x60;320x240&#x60;  Note that the &#x60;resolution&#x60; imposes an aspect ratio to the resulting composition. When the original video tracks are constrained by the aspect ratio, they are scaled to fit. See [Specifying Video Layouts](https://www.twilio.com/docs/video/api/compositions-resource#specifying-video-layouts) for more info. | 
**StatusCallback** | **String**| The URL we should call using the &#x60;status_callback_method&#x60; to send status information to your application on every composition event. If not provided, status callback events will not be dispatched. | 
**StatusCallbackMethod** | **String**| The HTTP method we should use to call &#x60;status_callback&#x60;. Can be: &#x60;POST&#x60; or &#x60;GET&#x60; and the default is &#x60;POST&#x60;. | 
**Trim** | **Bool**| Whether to clip the intervals where there is no active media in the compositions triggered by the composition hook. The default is &#x60;true&#x60;. Compositions with &#x60;trim&#x60; enabled are shorter when the Room is created and no Participant joins for a while as well as if all the Participants leave the room and join later, because those gaps will be removed. See [Specifying Video Layouts](https://www.twilio.com/docs/video/api/compositions-resource#specifying-video-layouts) for more info. | 
**VideoLayout** | [**map[string]interface{}**](map[string]interface{}.md)| A JSON object that describes the video layout of the composition hook in terms of regions. See [Specifying Video Layouts](https://www.twilio.com/docs/video/api/compositions-resource#specifying-video-layouts) for more info. | 

### Return type

[**VideoV1CompositionHook**](video.v1.composition_hook.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRoom

> VideoV1Room UpdateRoom(ctx, Sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| The SID of the Room resource to update. | 
 **optional** | ***UpdateRoomRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateRoomRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Status** | **String**| The new status of the resource. Set to &#x60;completed&#x60; to end the room. | 

### Return type

[**VideoV1Room**](video.v1.room.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRoomParticipant

> VideoV1RoomRoomParticipant UpdateRoomParticipant(ctx, RoomSid, Sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**RoomSid** | **string**| The SID of the room with the participant to update. | 
**Sid** | **string**| The SID of the RoomParticipant resource to update. | 
 **optional** | ***UpdateRoomParticipantRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateRoomParticipantRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Status** | **String**| The new status of the resource. Can be: &#x60;connected&#x60; or &#x60;disconnected&#x60;. For &#x60;in-progress&#x60; Rooms the default Status is &#x60;connected&#x60;, for &#x60;completed&#x60; Rooms only &#x60;disconnected&#x60; Participants are returned. | 

### Return type

[**VideoV1RoomRoomParticipant**](video.v1.room.room_participant.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRoomParticipantSubscribeRule

> VideoV1RoomRoomParticipantRoomParticipantSubscribeRule UpdateRoomParticipantSubscribeRule(ctx, RoomSid, ParticipantSid, optional)



Update the Subscribe Rules for the Participant

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**RoomSid** | **string**| The SID of the Room resource where the subscribe rules to update apply. | 
**ParticipantSid** | **string**| The SID of the Participant resource to update the Subscribe Rules. | 
 **optional** | ***UpdateRoomParticipantSubscribeRuleRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateRoomParticipantSubscribeRuleRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Rules** | [**map[string]interface{}**](map[string]interface{}.md)| A JSON-encoded array of subscribe rules. See the [Specifying Subscribe Rules](https://www.twilio.com/docs/video/api/track-subscriptions#specifying-sr) section for further information. | 

### Return type

[**VideoV1RoomRoomParticipantRoomParticipantSubscribeRule**](video.v1.room.room_participant.room_participant_subscribe_rule.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRoomRecordingRule

> VideoV1RoomRoomRecordingRule UpdateRoomRecordingRule(ctx, RoomSid, optional)



Update the Recording Rules for the Room

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**RoomSid** | **string**| The SID of the Room resource where the recording rules to update apply. | 
 **optional** | ***UpdateRoomRecordingRuleRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateRoomRecordingRuleRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Rules** | [**map[string]interface{}**](map[string]interface{}.md)| A JSON-encoded array of recording rules. | 

### Return type

[**VideoV1RoomRoomRecordingRule**](video.v1.room.room_recording_rule.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

