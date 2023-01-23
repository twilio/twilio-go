# MediaV1MediaRecording

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |[optional] 
**Duration** | **int** | The duration of the MediaRecording |[optional] 
**Format** | Pointer to [**string**](MediaRecordingEnumFormat.md) |  |
**Links** | **map[string]interface{}** | The URLs of related resources |[optional] 
**ProcessorSid** | **string** | The SID of the MediaProcessor |[optional] 
**Resolution** | **string** | The dimensions of the video image in pixels |[optional] 
**SourceSid** | **string** | The SID of the resource that generated the original media |[optional] 
**Sid** | **string** | The unique string that identifies the resource |[optional] 
**MediaSize** | **int64** | The size of the recording media |[optional] 
**Status** | Pointer to [**string**](MediaRecordingEnumStatus.md) |  |
**StatusCallback** | **string** | The URL to which Twilio will send MediaRecording event updates |[optional] 
**StatusCallbackMethod** | **string** | The HTTP method Twilio should use to call the `status_callback` URL |[optional] 
**Url** | **string** | The absolute URL of the resource |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


