# MediaV1MediaRecording

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the MediaRecording resource. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**Duration** | Pointer to **int** | The duration of the MediaRecording in seconds. |
**Format** | Pointer to [**string**](MediaRecordingEnumFormat.md) |  |
**Links** | Pointer to **map[string]interface{}** | The URLs of related resources. |
**ProcessorSid** | Pointer to **string** | The SID of the MediaProcessor resource which produced the MediaRecording. |
**Resolution** | Pointer to **string** | The dimensions of the video image in pixels expressed as columns (width) and rows (height). |
**SourceSid** | Pointer to **string** | The SID of the resource that generated the original media track(s) of the MediaRecording. |
**Sid** | Pointer to **string** | The unique string generated to identify the MediaRecording resource. |
**MediaSize** | Pointer to **int64** | The size of the recording media in bytes. |
**Status** | Pointer to [**string**](MediaRecordingEnumStatus.md) |  |
**StatusCallback** | Pointer to **string** | The URL to which Twilio will send asynchronous webhook requests for every MediaRecording event. See [Status Callbacks](/docs/live/status-callbacks) for more details. |
**StatusCallbackMethod** | Pointer to **string** | The HTTP method Twilio should use to call the `status_callback` URL. Can be `POST` or `GET` and the default is `POST`. |
**Url** | Pointer to **string** | The absolute URL of the resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


