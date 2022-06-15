# MediaV1MediaRecording

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**Bitrate** | Pointer to **int** | The bitrate of the media |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |
**Duration** | Pointer to **int** | The duration of the MediaRecording |
**Format** | Pointer to **string** | The format of the MediaRecording |
**Links** | Pointer to **map[string]interface{}** | The URLs of related resources |
**MediaSize** | Pointer to **int** | The size of the recording media |
**ProcessorSid** | Pointer to **string** | The SID of the MediaProcessor |
**Resolution** | Pointer to **string** | The dimensions of the video image in pixels |
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**SourceSid** | Pointer to **string** | The SID of the resource that generated the original media |
**Status** | Pointer to **string** | The status of the MediaRecording |
**StatusCallback** | Pointer to **string** | The URL to which Twilio will send MediaRecording event updates |
**StatusCallbackMethod** | Pointer to **string** | The HTTP method Twilio should use to call the `status_callback` URL |
**Url** | Pointer to **string** | The absolute URL of the resource |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


