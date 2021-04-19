# \DefaultApi

All URIs are relative to *https://video.twilio.com*

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

> VideoV1Composition CreateComposition(ctx).AudioSources(AudioSources).AudioSourcesExcluded(AudioSourcesExcluded).Format(Format).Resolution(Resolution).RoomSid(RoomSid).StatusCallback(StatusCallback).StatusCallbackMethod(StatusCallbackMethod).Trim(Trim).VideoLayout(VideoLayout).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AudioSources := []string{"Inner_example"} // []string | An array of track names from the same group room to merge into the new composition. Can include zero or more track names. The new composition includes all audio sources specified in `audio_sources` except for those specified in `audio_sources_excluded`. The track names in this parameter can include an asterisk as a wild card character, which will match zero or more characters in a track name. For example, `student*` includes `student` as well as `studentTeam`. Please, be aware that either video_layout or audio_sources have to be provided to get a valid creation request (optional)
    AudioSourcesExcluded := []string{"Inner_example"} // []string | An array of track names to exclude. The new composition includes all audio sources specified in `audio_sources` except for those specified in `audio_sources_excluded`. The track names in this parameter can include an asterisk as a wild card character, which will match zero or more characters in a track name. For example, `student*` excludes `student` as well as `studentTeam`. This parameter can also be empty. (optional)
    Format := "Format_example" // string | The container format of the composition's media files. Can be: `mp4` or `webm` and the default is `webm`. If you specify `mp4` or `webm`, you must also specify one or more `audio_sources` and/or a `video_layout` element that contains a valid `video_sources` list, otherwise an error occurs. (optional)
    Resolution := "Resolution_example" // string | A string that describes the columns (width) and rows (height) of the generated composed video in pixels. Defaults to `640x480`.  The string's format is `{width}x{height}` where:   * 16 <= `{width}` <= 1280 * 16 <= `{height}` <= 1280 * `{width}` * `{height}` <= 921,600  Typical values are:   * HD = `1280x720` * PAL = `1024x576` * VGA = `640x480` * CIF = `320x240`  Note that the `resolution` imposes an aspect ratio to the resulting composition. When the original video tracks are constrained by the aspect ratio, they are scaled to fit. See [Specifying Video Layouts](https://www.twilio.com/docs/video/api/compositions-resource#specifying-video-layouts) for more info. (optional)
    RoomSid := "RoomSid_example" // string | The SID of the Group Room with the media tracks to be used as composition sources. (optional)
    StatusCallback := "StatusCallback_example" // string | The URL we should call using the `status_callback_method` to send status information to your application on every composition event. If not provided, status callback events will not be dispatched. (optional)
    StatusCallbackMethod := "StatusCallbackMethod_example" // string | The HTTP method we should use to call `status_callback`. Can be: `POST` or `GET` and the default is `POST`. (optional)
    Trim := true // bool | Whether to clip the intervals where there is no active media in the composition. The default is `true`. Compositions with `trim` enabled are shorter when the Room is created and no Participant joins for a while as well as if all the Participants leave the room and join later, because those gaps will be removed. See [Specifying Video Layouts](https://www.twilio.com/docs/video/api/compositions-resource#specifying-video-layouts) for more info. (optional)
    VideoLayout := TODO // map[string]interface{} | An object that describes the video layout of the composition in terms of regions. See [Specifying Video Layouts](https://www.twilio.com/docs/video/api/compositions-resource#specifying-video-layouts) for more info. Please, be aware that either video_layout or audio_sources have to be provided to get a valid creation request (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateComposition(context.Background()).AudioSources(AudioSources).AudioSourcesExcluded(AudioSourcesExcluded).Format(Format).Resolution(Resolution).RoomSid(RoomSid).StatusCallback(StatusCallback).StatusCallbackMethod(StatusCallbackMethod).Trim(Trim).VideoLayout(VideoLayout).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateComposition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateComposition`: VideoV1Composition
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateComposition`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateCompositionParams struct via the builder pattern


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


## CreateCompositionHook

> VideoV1CompositionHook CreateCompositionHook(ctx).AudioSources(AudioSources).AudioSourcesExcluded(AudioSourcesExcluded).Enabled(Enabled).Format(Format).FriendlyName(FriendlyName).Resolution(Resolution).StatusCallback(StatusCallback).StatusCallbackMethod(StatusCallbackMethod).Trim(Trim).VideoLayout(VideoLayout).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AudioSources := []string{"Inner_example"} // []string | An array of track names from the same group room to merge into the compositions created by the composition hook. Can include zero or more track names. A composition triggered by the composition hook includes all audio sources specified in `audio_sources` except those specified in `audio_sources_excluded`. The track names in this parameter can include an asterisk as a wild card character, which matches zero or more characters in a track name. For example, `student*` includes tracks named `student` as well as `studentTeam`. (optional)
    AudioSourcesExcluded := []string{"Inner_example"} // []string | An array of track names to exclude. A composition triggered by the composition hook includes all audio sources specified in `audio_sources` except for those specified in `audio_sources_excluded`. The track names in this parameter can include an asterisk as a wild card character, which matches zero or more characters in a track name. For example, `student*` excludes `student` as well as `studentTeam`. This parameter can also be empty. (optional)
    Enabled := true // bool | Whether the composition hook is active. When `true`, the composition hook will be triggered for every completed Group Room in the account. When `false`, the composition hook will never be triggered. (optional)
    Format := "Format_example" // string | The container format of the media files used by the compositions created by the composition hook. Can be: `mp4` or `webm` and the default is `webm`. If `mp4` or `webm`, `audio_sources` must have one or more tracks and/or a `video_layout` element must contain a valid `video_sources` list, otherwise an error occurs. (optional)
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you create to describe the resource. It can be up to  100 characters long and it must be unique within the account. (optional)
    Resolution := "Resolution_example" // string | A string that describes the columns (width) and rows (height) of the generated composed video in pixels. Defaults to `640x480`.  The string's format is `{width}x{height}` where:   * 16 <= `{width}` <= 1280 * 16 <= `{height}` <= 1280 * `{width}` * `{height}` <= 921,600  Typical values are:   * HD = `1280x720` * PAL = `1024x576` * VGA = `640x480` * CIF = `320x240`  Note that the `resolution` imposes an aspect ratio to the resulting composition. When the original video tracks are constrained by the aspect ratio, they are scaled to fit. See [Specifying Video Layouts](https://www.twilio.com/docs/video/api/compositions-resource#specifying-video-layouts) for more info. (optional)
    StatusCallback := "StatusCallback_example" // string | The URL we should call using the `status_callback_method` to send status information to your application on every composition event. If not provided, status callback events will not be dispatched. (optional)
    StatusCallbackMethod := "StatusCallbackMethod_example" // string | The HTTP method we should use to call `status_callback`. Can be: `POST` or `GET` and the default is `POST`. (optional)
    Trim := true // bool | Whether to clip the intervals where there is no active media in the Compositions triggered by the composition hook. The default is `true`. Compositions with `trim` enabled are shorter when the Room is created and no Participant joins for a while as well as if all the Participants leave the room and join later, because those gaps will be removed. See [Specifying Video Layouts](https://www.twilio.com/docs/video/api/compositions-resource#specifying-video-layouts) for more info. (optional)
    VideoLayout := TODO // map[string]interface{} | An object that describes the video layout of the composition hook in terms of regions. See [Specifying Video Layouts](https://www.twilio.com/docs/video/api/compositions-resource#specifying-video-layouts) for more info. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateCompositionHook(context.Background()).AudioSources(AudioSources).AudioSourcesExcluded(AudioSourcesExcluded).Enabled(Enabled).Format(Format).FriendlyName(FriendlyName).Resolution(Resolution).StatusCallback(StatusCallback).StatusCallbackMethod(StatusCallbackMethod).Trim(Trim).VideoLayout(VideoLayout).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateCompositionHook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCompositionHook`: VideoV1CompositionHook
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateCompositionHook`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateCompositionHookParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **AudioSources** | **[]string** | An array of track names from the same group room to merge into the compositions created by the composition hook. Can include zero or more track names. A composition triggered by the composition hook includes all audio sources specified in &#x60;audio_sources&#x60; except those specified in &#x60;audio_sources_excluded&#x60;. The track names in this parameter can include an asterisk as a wild card character, which matches zero or more characters in a track name. For example, &#x60;student*&#x60; includes tracks named &#x60;student&#x60; as well as &#x60;studentTeam&#x60;.
 **AudioSourcesExcluded** | **[]string** | An array of track names to exclude. A composition triggered by the composition hook includes all audio sources specified in &#x60;audio_sources&#x60; except for those specified in &#x60;audio_sources_excluded&#x60;. The track names in this parameter can include an asterisk as a wild card character, which matches zero or more characters in a track name. For example, &#x60;student*&#x60; excludes &#x60;student&#x60; as well as &#x60;studentTeam&#x60;. This parameter can also be empty.
 **Enabled** | **bool** | Whether the composition hook is active. When &#x60;true&#x60;, the composition hook will be triggered for every completed Group Room in the account. When &#x60;false&#x60;, the composition hook will never be triggered.
 **Format** | **string** | The container format of the media files used by the compositions created by the composition hook. Can be: &#x60;mp4&#x60; or &#x60;webm&#x60; and the default is &#x60;webm&#x60;. If &#x60;mp4&#x60; or &#x60;webm&#x60;, &#x60;audio_sources&#x60; must have one or more tracks and/or a &#x60;video_layout&#x60; element must contain a valid &#x60;video_sources&#x60; list, otherwise an error occurs.
 **FriendlyName** | **string** | A descriptive string that you create to describe the resource. It can be up to  100 characters long and it must be unique within the account.
 **Resolution** | **string** | A string that describes the columns (width) and rows (height) of the generated composed video in pixels. Defaults to &#x60;640x480&#x60;.  The string&#39;s format is &#x60;{width}x{height}&#x60; where:   * 16 &lt;&#x3D; &#x60;{width}&#x60; &lt;&#x3D; 1280 * 16 &lt;&#x3D; &#x60;{height}&#x60; &lt;&#x3D; 1280 * &#x60;{width}&#x60; * &#x60;{height}&#x60; &lt;&#x3D; 921,600  Typical values are:   * HD &#x3D; &#x60;1280x720&#x60; * PAL &#x3D; &#x60;1024x576&#x60; * VGA &#x3D; &#x60;640x480&#x60; * CIF &#x3D; &#x60;320x240&#x60;  Note that the &#x60;resolution&#x60; imposes an aspect ratio to the resulting composition. When the original video tracks are constrained by the aspect ratio, they are scaled to fit. See [Specifying Video Layouts](https://www.twilio.com/docs/video/api/compositions-resource#specifying-video-layouts) for more info.
 **StatusCallback** | **string** | The URL we should call using the &#x60;status_callback_method&#x60; to send status information to your application on every composition event. If not provided, status callback events will not be dispatched.
 **StatusCallbackMethod** | **string** | The HTTP method we should use to call &#x60;status_callback&#x60;. Can be: &#x60;POST&#x60; or &#x60;GET&#x60; and the default is &#x60;POST&#x60;.
 **Trim** | **bool** | Whether to clip the intervals where there is no active media in the Compositions triggered by the composition hook. The default is &#x60;true&#x60;. Compositions with &#x60;trim&#x60; enabled are shorter when the Room is created and no Participant joins for a while as well as if all the Participants leave the room and join later, because those gaps will be removed. See [Specifying Video Layouts](https://www.twilio.com/docs/video/api/compositions-resource#specifying-video-layouts) for more info.
 **VideoLayout** | [**map[string]interface{}**](map[string]interface{}.md) | An object that describes the video layout of the composition hook in terms of regions. See [Specifying Video Layouts](https://www.twilio.com/docs/video/api/compositions-resource#specifying-video-layouts) for more info.

### Return type

[**VideoV1CompositionHook**](VideoV1CompositionHook.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCompositionSettings

> VideoV1CompositionSettings CreateCompositionSettings(ctx).AwsCredentialsSid(AwsCredentialsSid).AwsS3Url(AwsS3Url).AwsStorageEnabled(AwsStorageEnabled).EncryptionEnabled(EncryptionEnabled).EncryptionKeySid(EncryptionKeySid).FriendlyName(FriendlyName).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AwsCredentialsSid := "AwsCredentialsSid_example" // string | The SID of the stored Credential resource. (optional)
    AwsS3Url := "AwsS3Url_example" // string | The URL of the AWS S3 bucket where the compositions should be stored. We only support DNS-compliant URLs like `https://documentation-example-twilio-bucket/compositions`, where `compositions` is the path in which you want the compositions to be stored. This URL accepts only URI-valid characters, as described in the <a href='https://tools.ietf.org/html/rfc3986#section-2'>RFC 3986</a>. (optional)
    AwsStorageEnabled := true // bool | Whether all compositions should be written to the `aws_s3_url`. When `false`, all compositions are stored in our cloud. (optional)
    EncryptionEnabled := true // bool | Whether all compositions should be stored in an encrypted form. The default is `false`. (optional)
    EncryptionKeySid := "EncryptionKeySid_example" // string | The SID of the Public Key resource to use for encryption. (optional)
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you create to describe the resource and show to the user in the console (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateCompositionSettings(context.Background()).AwsCredentialsSid(AwsCredentialsSid).AwsS3Url(AwsS3Url).AwsStorageEnabled(AwsStorageEnabled).EncryptionEnabled(EncryptionEnabled).EncryptionKeySid(EncryptionKeySid).FriendlyName(FriendlyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateCompositionSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCompositionSettings`: VideoV1CompositionSettings
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateCompositionSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateCompositionSettingsParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **AwsCredentialsSid** | **string** | The SID of the stored Credential resource.
 **AwsS3Url** | **string** | The URL of the AWS S3 bucket where the compositions should be stored. We only support DNS-compliant URLs like &#x60;https://documentation-example-twilio-bucket/compositions&#x60;, where &#x60;compositions&#x60; is the path in which you want the compositions to be stored. This URL accepts only URI-valid characters, as described in the &lt;a href&#x3D;&#39;https://tools.ietf.org/html/rfc3986#section-2&#39;&gt;RFC 3986&lt;/a&gt;.
 **AwsStorageEnabled** | **bool** | Whether all compositions should be written to the &#x60;aws_s3_url&#x60;. When &#x60;false&#x60;, all compositions are stored in our cloud.
 **EncryptionEnabled** | **bool** | Whether all compositions should be stored in an encrypted form. The default is &#x60;false&#x60;.
 **EncryptionKeySid** | **string** | The SID of the Public Key resource to use for encryption.
 **FriendlyName** | **string** | A descriptive string that you create to describe the resource and show to the user in the console

### Return type

[**VideoV1CompositionSettings**](VideoV1CompositionSettings.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRecordingSettings

> VideoV1RecordingSettings CreateRecordingSettings(ctx).AwsCredentialsSid(AwsCredentialsSid).AwsS3Url(AwsS3Url).AwsStorageEnabled(AwsStorageEnabled).EncryptionEnabled(EncryptionEnabled).EncryptionKeySid(EncryptionKeySid).FriendlyName(FriendlyName).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AwsCredentialsSid := "AwsCredentialsSid_example" // string | The SID of the stored Credential resource. (optional)
    AwsS3Url := "AwsS3Url_example" // string | The URL of the AWS S3 bucket where the recordings should be stored. We only support DNS-compliant URLs like `https://documentation-example-twilio-bucket/recordings`, where `recordings` is the path in which you want the recordings to be stored. This URL accepts only URI-valid characters, as described in the <a href='https://tools.ietf.org/html/rfc3986#section-2'>RFC 3986</a>. (optional)
    AwsStorageEnabled := true // bool | Whether all recordings should be written to the `aws_s3_url`. When `false`, all recordings are stored in our cloud. (optional)
    EncryptionEnabled := true // bool | Whether all recordings should be stored in an encrypted form. The default is `false`. (optional)
    EncryptionKeySid := "EncryptionKeySid_example" // string | The SID of the Public Key resource to use for encryption. (optional)
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you create to describe the resource and be shown to users in the console (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateRecordingSettings(context.Background()).AwsCredentialsSid(AwsCredentialsSid).AwsS3Url(AwsS3Url).AwsStorageEnabled(AwsStorageEnabled).EncryptionEnabled(EncryptionEnabled).EncryptionKeySid(EncryptionKeySid).FriendlyName(FriendlyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateRecordingSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRecordingSettings`: VideoV1RecordingSettings
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateRecordingSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateRecordingSettingsParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **AwsCredentialsSid** | **string** | The SID of the stored Credential resource.
 **AwsS3Url** | **string** | The URL of the AWS S3 bucket where the recordings should be stored. We only support DNS-compliant URLs like &#x60;https://documentation-example-twilio-bucket/recordings&#x60;, where &#x60;recordings&#x60; is the path in which you want the recordings to be stored. This URL accepts only URI-valid characters, as described in the &lt;a href&#x3D;&#39;https://tools.ietf.org/html/rfc3986#section-2&#39;&gt;RFC 3986&lt;/a&gt;.
 **AwsStorageEnabled** | **bool** | Whether all recordings should be written to the &#x60;aws_s3_url&#x60;. When &#x60;false&#x60;, all recordings are stored in our cloud.
 **EncryptionEnabled** | **bool** | Whether all recordings should be stored in an encrypted form. The default is &#x60;false&#x60;.
 **EncryptionKeySid** | **string** | The SID of the Public Key resource to use for encryption.
 **FriendlyName** | **string** | A descriptive string that you create to describe the resource and be shown to users in the console

### Return type

[**VideoV1RecordingSettings**](VideoV1RecordingSettings.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRoom

> VideoV1Room CreateRoom(ctx).EnableTurn(EnableTurn).MaxParticipants(MaxParticipants).MediaRegion(MediaRegion).RecordParticipantsOnConnect(RecordParticipantsOnConnect).RecordingRules(RecordingRules).StatusCallback(StatusCallback).StatusCallbackMethod(StatusCallbackMethod).Type(Type).UniqueName(UniqueName).VideoCodecs(VideoCodecs).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    EnableTurn := true // bool | Deprecated, now always considered to be true. (optional)
    MaxParticipants := int32(56) // int32 | The maximum number of concurrent Participants allowed in the room. Peer-to-peer rooms can have up to 10 Participants. Small Group rooms can have up to 4 Participants. Group rooms can have up to 50 Participants. (optional)
    MediaRegion := "MediaRegion_example" // string | The region for the media server in Group Rooms.  Can be: one of the [available Media Regions](https://www.twilio.com/docs/video/ip-address-whitelisting#group-rooms-media-servers). ***This feature is not available in `peer-to-peer` rooms.*** (optional)
    RecordParticipantsOnConnect := true // bool | Whether to start recording when Participants connect. ***This feature is not available in `peer-to-peer` rooms.*** (optional)
    RecordingRules := TODO // map[string]interface{} | A collection of Recording Rules that describe how to include or exclude matching tracks for recording (optional)
    StatusCallback := "StatusCallback_example" // string | The URL we should call using the `status_callback_method` to send status information to your application on every room event. See [Status Callbacks](https://www.twilio.com/docs/video/api/status-callbacks) for more info. (optional)
    StatusCallbackMethod := "StatusCallbackMethod_example" // string | The HTTP method we should use to call `status_callback`. Can be `POST` or `GET`. (optional)
    Type := "Type_example" // string | The type of room. Can be: `go`, `peer-to-peer`, `group-small`, or `group`. The default value is `group`. (optional)
    UniqueName := "UniqueName_example" // string | An application-defined string that uniquely identifies the resource. It can be used as a `room_sid` in place of the resource's `sid` in the URL to address the resource. This value is unique for `in-progress` rooms. SDK clients can use this name to connect to the room. REST API clients can use this name in place of the Room SID to interact with the room as long as the room is `in-progress`. (optional)
    VideoCodecs := []string{"Inner_example"} // []string | An array of the video codecs that are supported when publishing a track in the room.  Can be: `VP8` and `H264`.  ***This feature is not available in `peer-to-peer` rooms*** (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateRoom(context.Background()).EnableTurn(EnableTurn).MaxParticipants(MaxParticipants).MediaRegion(MediaRegion).RecordParticipantsOnConnect(RecordParticipantsOnConnect).RecordingRules(RecordingRules).StatusCallback(StatusCallback).StatusCallbackMethod(StatusCallbackMethod).Type(Type).UniqueName(UniqueName).VideoCodecs(VideoCodecs).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateRoom``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRoom`: VideoV1Room
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateRoom`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateRoomParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **EnableTurn** | **bool** | Deprecated, now always considered to be true.
 **MaxParticipants** | **int32** | The maximum number of concurrent Participants allowed in the room. Peer-to-peer rooms can have up to 10 Participants. Small Group rooms can have up to 4 Participants. Group rooms can have up to 50 Participants.
 **MediaRegion** | **string** | The region for the media server in Group Rooms.  Can be: one of the [available Media Regions](https://www.twilio.com/docs/video/ip-address-whitelisting#group-rooms-media-servers). ***This feature is not available in &#x60;peer-to-peer&#x60; rooms.***
 **RecordParticipantsOnConnect** | **bool** | Whether to start recording when Participants connect. ***This feature is not available in &#x60;peer-to-peer&#x60; rooms.***
 **RecordingRules** | [**map[string]interface{}**](map[string]interface{}.md) | A collection of Recording Rules that describe how to include or exclude matching tracks for recording
 **StatusCallback** | **string** | The URL we should call using the &#x60;status_callback_method&#x60; to send status information to your application on every room event. See [Status Callbacks](https://www.twilio.com/docs/video/api/status-callbacks) for more info.
 **StatusCallbackMethod** | **string** | The HTTP method we should use to call &#x60;status_callback&#x60;. Can be &#x60;POST&#x60; or &#x60;GET&#x60;.
 **Type** | **string** | The type of room. Can be: &#x60;go&#x60;, &#x60;peer-to-peer&#x60;, &#x60;group-small&#x60;, or &#x60;group&#x60;. The default value is &#x60;group&#x60;.
 **UniqueName** | **string** | An application-defined string that uniquely identifies the resource. It can be used as a &#x60;room_sid&#x60; in place of the resource&#39;s &#x60;sid&#x60; in the URL to address the resource. This value is unique for &#x60;in-progress&#x60; rooms. SDK clients can use this name to connect to the room. REST API clients can use this name in place of the Room SID to interact with the room as long as the room is &#x60;in-progress&#x60;.
 **VideoCodecs** | **[]string** | An array of the video codecs that are supported when publishing a track in the room.  Can be: &#x60;VP8&#x60; and &#x60;H264&#x60;.  ***This feature is not available in &#x60;peer-to-peer&#x60; rooms***

### Return type

[**VideoV1Room**](VideoV1Room.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteComposition

> DeleteComposition(ctx, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    Sid := "Sid_example" // string | The SID of the Composition resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteComposition(context.Background(), Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteComposition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Composition resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteCompositionParams struct via the builder pattern


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


## DeleteCompositionHook

> DeleteCompositionHook(ctx, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    Sid := "Sid_example" // string | The SID of the CompositionHook resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteCompositionHook(context.Background(), Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteCompositionHook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the CompositionHook resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteCompositionHookParams struct via the builder pattern


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


## DeleteRecording

> DeleteRecording(ctx, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    Sid := "Sid_example" // string | The SID of the Recording resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteRecording(context.Background(), Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteRecording``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Recording resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteRecordingParams struct via the builder pattern


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


## DeleteRoomRecording

> DeleteRoomRecording(ctx, RoomSid, Sid).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    RoomSid := "RoomSid_example" // string | The SID of the room with the RoomRecording resource to delete.
    Sid := "Sid_example" // string | The SID of the RoomRecording resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteRoomRecording(context.Background(), RoomSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteRoomRecording``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**RoomSid** | **string** | The SID of the room with the RoomRecording resource to delete.
**Sid** | **string** | The SID of the RoomRecording resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteRoomRecordingParams struct via the builder pattern


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

> VideoV1Composition FetchComposition(ctx, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    Sid := "Sid_example" // string | The SID of the Composition resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchComposition(context.Background(), Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchComposition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchComposition`: VideoV1Composition
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchComposition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Composition resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchCompositionParams struct via the builder pattern


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


## FetchCompositionHook

> VideoV1CompositionHook FetchCompositionHook(ctx, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    Sid := "Sid_example" // string | The SID of the CompositionHook resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchCompositionHook(context.Background(), Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchCompositionHook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchCompositionHook`: VideoV1CompositionHook
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchCompositionHook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the CompositionHook resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchCompositionHookParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


### Return type

[**VideoV1CompositionHook**](VideoV1CompositionHook.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchCompositionSettings

> VideoV1CompositionSettings FetchCompositionSettings(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchCompositionSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchCompositionSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchCompositionSettings`: VideoV1CompositionSettings
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchCompositionSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a FetchCompositionSettingsParams struct via the builder pattern


### Return type

[**VideoV1CompositionSettings**](VideoV1CompositionSettings.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchRecording

> VideoV1Recording FetchRecording(ctx, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    Sid := "Sid_example" // string | The SID of the Recording resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchRecording(context.Background(), Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchRecording``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchRecording`: VideoV1Recording
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchRecording`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Recording resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchRecordingParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


### Return type

[**VideoV1Recording**](VideoV1Recording.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchRecordingSettings

> VideoV1RecordingSettings FetchRecordingSettings(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchRecordingSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchRecordingSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchRecordingSettings`: VideoV1RecordingSettings
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchRecordingSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a FetchRecordingSettingsParams struct via the builder pattern


### Return type

[**VideoV1RecordingSettings**](VideoV1RecordingSettings.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchRoom

> VideoV1Room FetchRoom(ctx, Sid).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    Sid := "Sid_example" // string | The SID of the Room resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchRoom(context.Background(), Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchRoom``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchRoom`: VideoV1Room
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchRoom`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Room resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchRoomParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


### Return type

[**VideoV1Room**](VideoV1Room.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchRoomParticipant

> VideoV1RoomRoomParticipant FetchRoomParticipant(ctx, RoomSid, Sid).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    RoomSid := "RoomSid_example" // string | The SID of the room with the Participant resource to fetch.
    Sid := "Sid_example" // string | The SID of the RoomParticipant resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchRoomParticipant(context.Background(), RoomSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchRoomParticipant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchRoomParticipant`: VideoV1RoomRoomParticipant
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchRoomParticipant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**RoomSid** | **string** | The SID of the room with the Participant resource to fetch.
**Sid** | **string** | The SID of the RoomParticipant resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchRoomParticipantParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

[**VideoV1RoomRoomParticipant**](VideoV1RoomRoomParticipant.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchRoomParticipantPublishedTrack

> VideoV1RoomRoomParticipantRoomParticipantPublishedTrack FetchRoomParticipantPublishedTrack(ctx, RoomSid, ParticipantSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    RoomSid := "RoomSid_example" // string | The SID of the Room resource where the Track resource to fetch is published.
    ParticipantSid := "ParticipantSid_example" // string | The SID of the Participant resource with the published track to fetch.
    Sid := "Sid_example" // string | The SID of the RoomParticipantPublishedTrack resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchRoomParticipantPublishedTrack(context.Background(), RoomSid, ParticipantSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchRoomParticipantPublishedTrack``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchRoomParticipantPublishedTrack`: VideoV1RoomRoomParticipantRoomParticipantPublishedTrack
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchRoomParticipantPublishedTrack`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**RoomSid** | **string** | The SID of the Room resource where the Track resource to fetch is published.
**ParticipantSid** | **string** | The SID of the Participant resource with the published track to fetch.
**Sid** | **string** | The SID of the RoomParticipantPublishedTrack resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchRoomParticipantPublishedTrackParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------




### Return type

[**VideoV1RoomRoomParticipantRoomParticipantPublishedTrack**](VideoV1RoomRoomParticipantRoomParticipantPublishedTrack.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchRoomParticipantSubscribeRule

> VideoV1RoomRoomParticipantRoomParticipantSubscribeRule FetchRoomParticipantSubscribeRule(ctx, RoomSid, ParticipantSid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    RoomSid := "RoomSid_example" // string | The SID of the Room resource where the subscribe rules to fetch apply.
    ParticipantSid := "ParticipantSid_example" // string | The SID of the Participant resource with the subscribe rules to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchRoomParticipantSubscribeRule(context.Background(), RoomSid, ParticipantSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchRoomParticipantSubscribeRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchRoomParticipantSubscribeRule`: VideoV1RoomRoomParticipantRoomParticipantSubscribeRule
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchRoomParticipantSubscribeRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**RoomSid** | **string** | The SID of the Room resource where the subscribe rules to fetch apply.
**ParticipantSid** | **string** | The SID of the Participant resource with the subscribe rules to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchRoomParticipantSubscribeRuleParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

[**VideoV1RoomRoomParticipantRoomParticipantSubscribeRule**](VideoV1RoomRoomParticipantRoomParticipantSubscribeRule.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchRoomParticipantSubscribedTrack

> VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack FetchRoomParticipantSubscribedTrack(ctx, RoomSid, ParticipantSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    RoomSid := "RoomSid_example" // string | The SID of the Room where the Track resource to fetch is subscribed.
    ParticipantSid := "ParticipantSid_example" // string | The SID of the participant that subscribes to the Track resource to fetch.
    Sid := "Sid_example" // string | The SID of the RoomParticipantSubscribedTrack resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchRoomParticipantSubscribedTrack(context.Background(), RoomSid, ParticipantSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchRoomParticipantSubscribedTrack``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchRoomParticipantSubscribedTrack`: VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchRoomParticipantSubscribedTrack`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**RoomSid** | **string** | The SID of the Room where the Track resource to fetch is subscribed.
**ParticipantSid** | **string** | The SID of the participant that subscribes to the Track resource to fetch.
**Sid** | **string** | The SID of the RoomParticipantSubscribedTrack resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchRoomParticipantSubscribedTrackParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------




### Return type

[**VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack**](VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchRoomRecording

> VideoV1RoomRoomRecording FetchRoomRecording(ctx, RoomSid, Sid).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    RoomSid := "RoomSid_example" // string | The SID of the Room resource with the recording to fetch.
    Sid := "Sid_example" // string | The SID of the RoomRecording resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchRoomRecording(context.Background(), RoomSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchRoomRecording``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchRoomRecording`: VideoV1RoomRoomRecording
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchRoomRecording`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**RoomSid** | **string** | The SID of the Room resource with the recording to fetch.
**Sid** | **string** | The SID of the RoomRecording resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchRoomRecordingParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

[**VideoV1RoomRoomRecording**](VideoV1RoomRoomRecording.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchRoomRecordingRule

> VideoV1RoomRoomRecordingRule FetchRoomRecordingRule(ctx, RoomSid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    RoomSid := "RoomSid_example" // string | The SID of the Room resource where the recording rules to fetch apply.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchRoomRecordingRule(context.Background(), RoomSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchRoomRecordingRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchRoomRecordingRule`: VideoV1RoomRoomRecordingRule
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchRoomRecordingRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**RoomSid** | **string** | The SID of the Room resource where the recording rules to fetch apply.

### Other Parameters

Other parameters are passed through a pointer to a FetchRoomRecordingRuleParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


### Return type

[**VideoV1RoomRoomRecordingRule**](VideoV1RoomRoomRecordingRule.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListComposition

> ListCompositionResponse ListComposition(ctx).Status(Status).DateCreatedAfter(DateCreatedAfter).DateCreatedBefore(DateCreatedBefore).RoomSid(RoomSid).PageSize(PageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    Status := "Status_example" // string | Read only Composition resources with this status. Can be: `enqueued`, `processing`, `completed`, `deleted`, or `failed`. (optional)
    DateCreatedAfter := time.Now() // time.Time | Read only Composition resources created on or after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date-time with time zone. (optional)
    DateCreatedBefore := time.Now() // time.Time | Read only Composition resources created before this ISO 8601 date-time with time zone. (optional)
    RoomSid := "RoomSid_example" // string | Read only Composition resources with this Room SID. (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListComposition(context.Background()).Status(Status).DateCreatedAfter(DateCreatedAfter).DateCreatedBefore(DateCreatedBefore).RoomSid(RoomSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListComposition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListComposition`: ListCompositionResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListComposition`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListCompositionParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **Status** | **string** | Read only Composition resources with this status. Can be: &#x60;enqueued&#x60;, &#x60;processing&#x60;, &#x60;completed&#x60;, &#x60;deleted&#x60;, or &#x60;failed&#x60;.
 **DateCreatedAfter** | **time.Time** | Read only Composition resources created on or after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date-time with time zone.
 **DateCreatedBefore** | **time.Time** | Read only Composition resources created before this ISO 8601 date-time with time zone.
 **RoomSid** | **string** | Read only Composition resources with this Room SID.
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

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

> ListCompositionHookResponse ListCompositionHook(ctx).Enabled(Enabled).DateCreatedAfter(DateCreatedAfter).DateCreatedBefore(DateCreatedBefore).FriendlyName(FriendlyName).PageSize(PageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    Enabled := true // bool | Read only CompositionHook resources with an `enabled` value that matches this parameter. (optional)
    DateCreatedAfter := time.Now() // time.Time | Read only CompositionHook resources created on or after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) datetime with time zone. (optional)
    DateCreatedBefore := time.Now() // time.Time | Read only CompositionHook resources created before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) datetime with time zone. (optional)
    FriendlyName := "FriendlyName_example" // string | Read only CompositionHook resources with friendly names that match this string. The match is not case sensitive and can include asterisk `*` characters as wildcard match. (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListCompositionHook(context.Background()).Enabled(Enabled).DateCreatedAfter(DateCreatedAfter).DateCreatedBefore(DateCreatedBefore).FriendlyName(FriendlyName).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListCompositionHook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCompositionHook`: ListCompositionHookResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListCompositionHook`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListCompositionHookParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **Enabled** | **bool** | Read only CompositionHook resources with an &#x60;enabled&#x60; value that matches this parameter.
 **DateCreatedAfter** | **time.Time** | Read only CompositionHook resources created on or after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) datetime with time zone.
 **DateCreatedBefore** | **time.Time** | Read only CompositionHook resources created before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) datetime with time zone.
 **FriendlyName** | **string** | Read only CompositionHook resources with friendly names that match this string. The match is not case sensitive and can include asterisk &#x60;*&#x60; characters as wildcard match.
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

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

> ListRecordingResponse ListRecording(ctx).Status(Status).SourceSid(SourceSid).GroupingSid(GroupingSid).DateCreatedAfter(DateCreatedAfter).DateCreatedBefore(DateCreatedBefore).MediaType(MediaType).PageSize(PageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    Status := "Status_example" // string | Read only the recordings that have this status. Can be: `processing`, `completed`, or `deleted`. (optional)
    SourceSid := "SourceSid_example" // string | Read only the recordings that have this `source_sid`. (optional)
    GroupingSid := []string{"Inner_example"} // []string | Read only recordings with this `grouping_sid`, which may include a `participant_sid` and/or a `room_sid`. (optional)
    DateCreatedAfter := time.Now() // time.Time | Read only recordings that started on or after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date-time with time zone. (optional)
    DateCreatedBefore := time.Now() // time.Time | Read only recordings that started before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date-time with time zone, given as `YYYY-MM-DDThh:mm:ss+|-hh:mm` or `YYYY-MM-DDThh:mm:ssZ`. (optional)
    MediaType := "MediaType_example" // string | Read only recordings that have this media type. Can be either `audio` or `video`. (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListRecording(context.Background()).Status(Status).SourceSid(SourceSid).GroupingSid(GroupingSid).DateCreatedAfter(DateCreatedAfter).DateCreatedBefore(DateCreatedBefore).MediaType(MediaType).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListRecording``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRecording`: ListRecordingResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListRecording`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListRecordingParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **Status** | **string** | Read only the recordings that have this status. Can be: &#x60;processing&#x60;, &#x60;completed&#x60;, or &#x60;deleted&#x60;.
 **SourceSid** | **string** | Read only the recordings that have this &#x60;source_sid&#x60;.
 **GroupingSid** | **[]string** | Read only recordings with this &#x60;grouping_sid&#x60;, which may include a &#x60;participant_sid&#x60; and/or a &#x60;room_sid&#x60;.
 **DateCreatedAfter** | **time.Time** | Read only recordings that started on or after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date-time with time zone.
 **DateCreatedBefore** | **time.Time** | Read only recordings that started before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date-time with time zone, given as &#x60;YYYY-MM-DDThh:mm:ss+|-hh:mm&#x60; or &#x60;YYYY-MM-DDThh:mm:ssZ&#x60;.
 **MediaType** | **string** | Read only recordings that have this media type. Can be either &#x60;audio&#x60; or &#x60;video&#x60;.
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

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

> ListRoomResponse ListRoom(ctx).Status(Status).UniqueName(UniqueName).DateCreatedAfter(DateCreatedAfter).DateCreatedBefore(DateCreatedBefore).PageSize(PageSize).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    Status := "Status_example" // string | Read only the rooms with this status. Can be: `in-progress` (default) or `completed` (optional)
    UniqueName := "UniqueName_example" // string | Read only rooms with the this `unique_name`. (optional)
    DateCreatedAfter := time.Now() // time.Time | Read only rooms that started on or after this date, given as `YYYY-MM-DD`. (optional)
    DateCreatedBefore := time.Now() // time.Time | Read only rooms that started before this date, given as `YYYY-MM-DD`. (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListRoom(context.Background()).Status(Status).UniqueName(UniqueName).DateCreatedAfter(DateCreatedAfter).DateCreatedBefore(DateCreatedBefore).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListRoom``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRoom`: ListRoomResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListRoom`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListRoomParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **Status** | **string** | Read only the rooms with this status. Can be: &#x60;in-progress&#x60; (default) or &#x60;completed&#x60;
 **UniqueName** | **string** | Read only rooms with the this &#x60;unique_name&#x60;.
 **DateCreatedAfter** | **time.Time** | Read only rooms that started on or after this date, given as &#x60;YYYY-MM-DD&#x60;.
 **DateCreatedBefore** | **time.Time** | Read only rooms that started before this date, given as &#x60;YYYY-MM-DD&#x60;.
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

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

> ListRoomParticipantResponse ListRoomParticipant(ctx, RoomSid).Status(Status).Identity(Identity).DateCreatedAfter(DateCreatedAfter).DateCreatedBefore(DateCreatedBefore).PageSize(PageSize).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    RoomSid := "RoomSid_example" // string | The SID of the room with the Participant resources to read.
    Status := "Status_example" // string | Read only the participants with this status. Can be: `connected` or `disconnected`. For `in-progress` Rooms the default Status is `connected`, for `completed` Rooms only `disconnected` Participants are returned. (optional)
    Identity := "Identity_example" // string | Read only the Participants with this [User](https://www.twilio.com/docs/chat/rest/user-resource) `identity` value. (optional)
    DateCreatedAfter := time.Now() // time.Time | Read only Participants that started after this date in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601#UTC) format. (optional)
    DateCreatedBefore := time.Now() // time.Time | Read only Participants that started before this date in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601#UTC) format. (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListRoomParticipant(context.Background(), RoomSid).Status(Status).Identity(Identity).DateCreatedAfter(DateCreatedAfter).DateCreatedBefore(DateCreatedBefore).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListRoomParticipant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRoomParticipant`: ListRoomParticipantResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListRoomParticipant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**RoomSid** | **string** | The SID of the room with the Participant resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListRoomParticipantParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **Status** | **string** | Read only the participants with this status. Can be: &#x60;connected&#x60; or &#x60;disconnected&#x60;. For &#x60;in-progress&#x60; Rooms the default Status is &#x60;connected&#x60;, for &#x60;completed&#x60; Rooms only &#x60;disconnected&#x60; Participants are returned.
 **Identity** | **string** | Read only the Participants with this [User](https://www.twilio.com/docs/chat/rest/user-resource) &#x60;identity&#x60; value.
 **DateCreatedAfter** | **time.Time** | Read only Participants that started after this date in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601#UTC) format.
 **DateCreatedBefore** | **time.Time** | Read only Participants that started before this date in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601#UTC) format.
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

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

> ListRoomParticipantPublishedTrackResponse ListRoomParticipantPublishedTrack(ctx, RoomSid, ParticipantSid).PageSize(PageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    RoomSid := "RoomSid_example" // string | The SID of the Room resource where the Track resources to read are published.
    ParticipantSid := "ParticipantSid_example" // string | The SID of the Participant resource with the published tracks to read.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListRoomParticipantPublishedTrack(context.Background(), RoomSid, ParticipantSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListRoomParticipantPublishedTrack``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRoomParticipantPublishedTrack`: ListRoomParticipantPublishedTrackResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListRoomParticipantPublishedTrack`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**RoomSid** | **string** | The SID of the Room resource where the Track resources to read are published.
**ParticipantSid** | **string** | The SID of the Participant resource with the published tracks to read.

### Other Parameters

Other parameters are passed through a pointer to a ListRoomParticipantPublishedTrackParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

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

> ListRoomParticipantSubscribedTrackResponse ListRoomParticipantSubscribedTrack(ctx, RoomSid, ParticipantSid).PageSize(PageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    RoomSid := "RoomSid_example" // string | The SID of the Room resource with the Track resources to read.
    ParticipantSid := "ParticipantSid_example" // string | The SID of the participant that subscribes to the Track resources to read.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListRoomParticipantSubscribedTrack(context.Background(), RoomSid, ParticipantSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListRoomParticipantSubscribedTrack``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRoomParticipantSubscribedTrack`: ListRoomParticipantSubscribedTrackResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListRoomParticipantSubscribedTrack`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**RoomSid** | **string** | The SID of the Room resource with the Track resources to read.
**ParticipantSid** | **string** | The SID of the participant that subscribes to the Track resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListRoomParticipantSubscribedTrackParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

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

> ListRoomRecordingResponse ListRoomRecording(ctx, RoomSid).Status(Status).SourceSid(SourceSid).DateCreatedAfter(DateCreatedAfter).DateCreatedBefore(DateCreatedBefore).PageSize(PageSize).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    RoomSid := "RoomSid_example" // string | The SID of the room with the RoomRecording resources to read.
    Status := "Status_example" // string | Read only the recordings with this status. Can be: `processing`, `completed`, or `deleted`. (optional)
    SourceSid := "SourceSid_example" // string | Read only the recordings that have this `source_sid`. (optional)
    DateCreatedAfter := time.Now() // time.Time | Read only recordings that started on or after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) datetime with time zone. (optional)
    DateCreatedBefore := time.Now() // time.Time | Read only Recordings that started before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) datetime with time zone. (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListRoomRecording(context.Background(), RoomSid).Status(Status).SourceSid(SourceSid).DateCreatedAfter(DateCreatedAfter).DateCreatedBefore(DateCreatedBefore).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListRoomRecording``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRoomRecording`: ListRoomRecordingResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListRoomRecording`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**RoomSid** | **string** | The SID of the room with the RoomRecording resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListRoomRecordingParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **Status** | **string** | Read only the recordings with this status. Can be: &#x60;processing&#x60;, &#x60;completed&#x60;, or &#x60;deleted&#x60;.
 **SourceSid** | **string** | Read only the recordings that have this &#x60;source_sid&#x60;.
 **DateCreatedAfter** | **time.Time** | Read only recordings that started on or after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) datetime with time zone.
 **DateCreatedBefore** | **time.Time** | Read only Recordings that started before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) datetime with time zone.
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

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

> VideoV1CompositionHook UpdateCompositionHook(ctx, Sid).AudioSources(AudioSources).AudioSourcesExcluded(AudioSourcesExcluded).Enabled(Enabled).Format(Format).FriendlyName(FriendlyName).Resolution(Resolution).StatusCallback(StatusCallback).StatusCallbackMethod(StatusCallbackMethod).Trim(Trim).VideoLayout(VideoLayout).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    Sid := "Sid_example" // string | The SID of the CompositionHook resource to update.
    AudioSources := []string{"Inner_example"} // []string | An array of track names from the same group room to merge into the compositions created by the composition hook. Can include zero or more track names. A composition triggered by the composition hook includes all audio sources specified in `audio_sources` except those specified in `audio_sources_excluded`. The track names in this parameter can include an asterisk as a wild card character, which matches zero or more characters in a track name. For example, `student*` includes tracks named `student` as well as `studentTeam`. (optional)
    AudioSourcesExcluded := []string{"Inner_example"} // []string | An array of track names to exclude. A composition triggered by the composition hook includes all audio sources specified in `audio_sources` except for those specified in `audio_sources_excluded`. The track names in this parameter can include an asterisk as a wild card character, which matches zero or more characters in a track name. For example, `student*` excludes `student` as well as `studentTeam`. This parameter can also be empty. (optional)
    Enabled := true // bool | Whether the composition hook is active. When `true`, the composition hook will be triggered for every completed Group Room in the account. When `false`, the composition hook never triggers. (optional)
    Format := "Format_example" // string | The container format of the media files used by the compositions created by the composition hook. Can be: `mp4` or `webm` and the default is `webm`. If `mp4` or `webm`, `audio_sources` must have one or more tracks and/or a `video_layout` element must contain a valid `video_sources` list, otherwise an error occurs. (optional)
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you create to describe the resource. It can be up to  100 characters long and it must be unique within the account. (optional)
    Resolution := "Resolution_example" // string | A string that describes the columns (width) and rows (height) of the generated composed video in pixels. Defaults to `640x480`.  The string's format is `{width}x{height}` where:   * 16 <= `{width}` <= 1280 * 16 <= `{height}` <= 1280 * `{width}` * `{height}` <= 921,600  Typical values are:   * HD = `1280x720` * PAL = `1024x576` * VGA = `640x480` * CIF = `320x240`  Note that the `resolution` imposes an aspect ratio to the resulting composition. When the original video tracks are constrained by the aspect ratio, they are scaled to fit. See [Specifying Video Layouts](https://www.twilio.com/docs/video/api/compositions-resource#specifying-video-layouts) for more info. (optional)
    StatusCallback := "StatusCallback_example" // string | The URL we should call using the `status_callback_method` to send status information to your application on every composition event. If not provided, status callback events will not be dispatched. (optional)
    StatusCallbackMethod := "StatusCallbackMethod_example" // string | The HTTP method we should use to call `status_callback`. Can be: `POST` or `GET` and the default is `POST`. (optional)
    Trim := true // bool | Whether to clip the intervals where there is no active media in the compositions triggered by the composition hook. The default is `true`. Compositions with `trim` enabled are shorter when the Room is created and no Participant joins for a while as well as if all the Participants leave the room and join later, because those gaps will be removed. See [Specifying Video Layouts](https://www.twilio.com/docs/video/api/compositions-resource#specifying-video-layouts) for more info. (optional)
    VideoLayout := TODO // map[string]interface{} | A JSON object that describes the video layout of the composition hook in terms of regions. See [Specifying Video Layouts](https://www.twilio.com/docs/video/api/compositions-resource#specifying-video-layouts) for more info. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateCompositionHook(context.Background(), Sid).AudioSources(AudioSources).AudioSourcesExcluded(AudioSourcesExcluded).Enabled(Enabled).Format(Format).FriendlyName(FriendlyName).Resolution(Resolution).StatusCallback(StatusCallback).StatusCallbackMethod(StatusCallbackMethod).Trim(Trim).VideoLayout(VideoLayout).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateCompositionHook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCompositionHook`: VideoV1CompositionHook
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateCompositionHook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the CompositionHook resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateCompositionHookParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **AudioSources** | **[]string** | An array of track names from the same group room to merge into the compositions created by the composition hook. Can include zero or more track names. A composition triggered by the composition hook includes all audio sources specified in &#x60;audio_sources&#x60; except those specified in &#x60;audio_sources_excluded&#x60;. The track names in this parameter can include an asterisk as a wild card character, which matches zero or more characters in a track name. For example, &#x60;student*&#x60; includes tracks named &#x60;student&#x60; as well as &#x60;studentTeam&#x60;.
 **AudioSourcesExcluded** | **[]string** | An array of track names to exclude. A composition triggered by the composition hook includes all audio sources specified in &#x60;audio_sources&#x60; except for those specified in &#x60;audio_sources_excluded&#x60;. The track names in this parameter can include an asterisk as a wild card character, which matches zero or more characters in a track name. For example, &#x60;student*&#x60; excludes &#x60;student&#x60; as well as &#x60;studentTeam&#x60;. This parameter can also be empty.
 **Enabled** | **bool** | Whether the composition hook is active. When &#x60;true&#x60;, the composition hook will be triggered for every completed Group Room in the account. When &#x60;false&#x60;, the composition hook never triggers.
 **Format** | **string** | The container format of the media files used by the compositions created by the composition hook. Can be: &#x60;mp4&#x60; or &#x60;webm&#x60; and the default is &#x60;webm&#x60;. If &#x60;mp4&#x60; or &#x60;webm&#x60;, &#x60;audio_sources&#x60; must have one or more tracks and/or a &#x60;video_layout&#x60; element must contain a valid &#x60;video_sources&#x60; list, otherwise an error occurs.
 **FriendlyName** | **string** | A descriptive string that you create to describe the resource. It can be up to  100 characters long and it must be unique within the account.
 **Resolution** | **string** | A string that describes the columns (width) and rows (height) of the generated composed video in pixels. Defaults to &#x60;640x480&#x60;.  The string&#39;s format is &#x60;{width}x{height}&#x60; where:   * 16 &lt;&#x3D; &#x60;{width}&#x60; &lt;&#x3D; 1280 * 16 &lt;&#x3D; &#x60;{height}&#x60; &lt;&#x3D; 1280 * &#x60;{width}&#x60; * &#x60;{height}&#x60; &lt;&#x3D; 921,600  Typical values are:   * HD &#x3D; &#x60;1280x720&#x60; * PAL &#x3D; &#x60;1024x576&#x60; * VGA &#x3D; &#x60;640x480&#x60; * CIF &#x3D; &#x60;320x240&#x60;  Note that the &#x60;resolution&#x60; imposes an aspect ratio to the resulting composition. When the original video tracks are constrained by the aspect ratio, they are scaled to fit. See [Specifying Video Layouts](https://www.twilio.com/docs/video/api/compositions-resource#specifying-video-layouts) for more info.
 **StatusCallback** | **string** | The URL we should call using the &#x60;status_callback_method&#x60; to send status information to your application on every composition event. If not provided, status callback events will not be dispatched.
 **StatusCallbackMethod** | **string** | The HTTP method we should use to call &#x60;status_callback&#x60;. Can be: &#x60;POST&#x60; or &#x60;GET&#x60; and the default is &#x60;POST&#x60;.
 **Trim** | **bool** | Whether to clip the intervals where there is no active media in the compositions triggered by the composition hook. The default is &#x60;true&#x60;. Compositions with &#x60;trim&#x60; enabled are shorter when the Room is created and no Participant joins for a while as well as if all the Participants leave the room and join later, because those gaps will be removed. See [Specifying Video Layouts](https://www.twilio.com/docs/video/api/compositions-resource#specifying-video-layouts) for more info.
 **VideoLayout** | [**map[string]interface{}**](map[string]interface{}.md) | A JSON object that describes the video layout of the composition hook in terms of regions. See [Specifying Video Layouts](https://www.twilio.com/docs/video/api/compositions-resource#specifying-video-layouts) for more info.

### Return type

[**VideoV1CompositionHook**](VideoV1CompositionHook.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRoom

> VideoV1Room UpdateRoom(ctx, Sid).Status(Status).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    Sid := "Sid_example" // string | The SID of the Room resource to update.
    Status := "Status_example" // string | The new status of the resource. Set to `completed` to end the room. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateRoom(context.Background(), Sid).Status(Status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateRoom``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRoom`: VideoV1Room
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateRoom`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Room resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateRoomParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **Status** | **string** | The new status of the resource. Set to &#x60;completed&#x60; to end the room.

### Return type

[**VideoV1Room**](VideoV1Room.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRoomParticipant

> VideoV1RoomRoomParticipant UpdateRoomParticipant(ctx, RoomSid, Sid).Status(Status).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    RoomSid := "RoomSid_example" // string | The SID of the room with the participant to update.
    Sid := "Sid_example" // string | The SID of the RoomParticipant resource to update.
    Status := "Status_example" // string | The new status of the resource. Can be: `connected` or `disconnected`. For `in-progress` Rooms the default Status is `connected`, for `completed` Rooms only `disconnected` Participants are returned. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateRoomParticipant(context.Background(), RoomSid, Sid).Status(Status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateRoomParticipant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRoomParticipant`: VideoV1RoomRoomParticipant
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateRoomParticipant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**RoomSid** | **string** | The SID of the room with the participant to update.
**Sid** | **string** | The SID of the RoomParticipant resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateRoomParticipantParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **Status** | **string** | The new status of the resource. Can be: &#x60;connected&#x60; or &#x60;disconnected&#x60;. For &#x60;in-progress&#x60; Rooms the default Status is &#x60;connected&#x60;, for &#x60;completed&#x60; Rooms only &#x60;disconnected&#x60; Participants are returned.

### Return type

[**VideoV1RoomRoomParticipant**](VideoV1RoomRoomParticipant.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRoomParticipantSubscribeRule

> VideoV1RoomRoomParticipantRoomParticipantSubscribeRule UpdateRoomParticipantSubscribeRule(ctx, RoomSid, ParticipantSid).Rules(Rules).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    RoomSid := "RoomSid_example" // string | The SID of the Room resource where the subscribe rules to update apply.
    ParticipantSid := "ParticipantSid_example" // string | The SID of the Participant resource to update the Subscribe Rules.
    Rules := TODO // map[string]interface{} | A JSON-encoded array of subscribe rules. See the [Specifying Subscribe Rules](https://www.twilio.com/docs/video/api/track-subscriptions#specifying-sr) section for further information. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateRoomParticipantSubscribeRule(context.Background(), RoomSid, ParticipantSid).Rules(Rules).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateRoomParticipantSubscribeRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRoomParticipantSubscribeRule`: VideoV1RoomRoomParticipantRoomParticipantSubscribeRule
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateRoomParticipantSubscribeRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**RoomSid** | **string** | The SID of the Room resource where the subscribe rules to update apply.
**ParticipantSid** | **string** | The SID of the Participant resource to update the Subscribe Rules.

### Other Parameters

Other parameters are passed through a pointer to a UpdateRoomParticipantSubscribeRuleParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **Rules** | [**map[string]interface{}**](map[string]interface{}.md) | A JSON-encoded array of subscribe rules. See the [Specifying Subscribe Rules](https://www.twilio.com/docs/video/api/track-subscriptions#specifying-sr) section for further information.

### Return type

[**VideoV1RoomRoomParticipantRoomParticipantSubscribeRule**](VideoV1RoomRoomParticipantRoomParticipantSubscribeRule.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRoomRecordingRule

> VideoV1RoomRoomRecordingRule UpdateRoomRecordingRule(ctx, RoomSid).Rules(Rules).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    RoomSid := "RoomSid_example" // string | The SID of the Room resource where the recording rules to update apply.
    Rules := TODO // map[string]interface{} | A JSON-encoded array of recording rules. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateRoomRecordingRule(context.Background(), RoomSid).Rules(Rules).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateRoomRecordingRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRoomRecordingRule`: VideoV1RoomRoomRecordingRule
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateRoomRecordingRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**RoomSid** | **string** | The SID of the Room resource where the recording rules to update apply.

### Other Parameters

Other parameters are passed through a pointer to a UpdateRoomRecordingRuleParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **Rules** | [**map[string]interface{}**](map[string]interface{}.md) | A JSON-encoded array of recording rules.

### Return type

[**VideoV1RoomRoomRecordingRule**](VideoV1RoomRoomRecordingRule.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

