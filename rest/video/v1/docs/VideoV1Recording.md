# VideoV1Recording

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**Status** | Pointer to [**string**](RecordingEnumStatus.md) |  |
**DateCreated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |[optional] 
**Sid** | **string** | The unique string that identifies the resource |[optional] 
**SourceSid** | **string** | The SID of the recording source |[optional] 
**Size** | **int64** | The size of the recorded track, in bytes |[optional] 
**Url** | **string** | The absolute URL of the resource |[optional] 
**Type** | Pointer to [**string**](RecordingEnumType.md) |  |
**Duration** | Pointer to **int** | The duration of the recording in seconds |
**ContainerFormat** | Pointer to [**string**](RecordingEnumFormat.md) |  |
**Codec** | Pointer to [**string**](RecordingEnumCodec.md) |  |
**GroupingSids** | Pointer to **interface{}** | A list of SIDs related to the recording |
**TrackName** | **string** | The name that was given to the source track of the recording |[optional] 
**Offset** | **int64** | The number of milliseconds between a point in time that is common to all rooms in a group and when the source room of the recording started |[optional] 
**MediaExternalLocation** | **string** | The URL of the media file associated with the recording when stored externally |[optional] 
**StatusCallback** | **string** | The URL called to send status information on every recording event. |[optional] 
**StatusCallbackMethod** | **string** | The HTTP method used to call `status_callback` |[optional] 
**Links** | **map[string]interface{}** | The URLs of related resources |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


