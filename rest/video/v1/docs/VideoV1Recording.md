# VideoV1Recording

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**Status** | Pointer to [**string**](RecordingEnumStatus.md) |  |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**SourceSid** | Pointer to **string** | The SID of the recording source |
**Size** | Pointer to **int64** | The size of the recorded track, in bytes |
**Url** | Pointer to **string** | The absolute URL of the resource |
**Type** | Pointer to [**string**](RecordingEnumType.md) |  |
**Duration** | Pointer to **int** | The duration of the recording in seconds |
**ContainerFormat** | Pointer to [**string**](RecordingEnumFormat.md) |  |
**Codec** | Pointer to [**string**](RecordingEnumCodec.md) |  |
**GroupingSids** | Pointer to **interface{}** | A list of SIDs related to the recording |
**TrackName** | Pointer to **string** | The name that was given to the source track of the recording |
**Offset** | Pointer to **int64** | The number of milliseconds between a point in time that is common to all rooms in a group and when the source room of the recording started |
**MediaExternalLocation** | Pointer to **string** | The URL of the media file associated with the recording when stored externally |
**StatusCallback** | Pointer to **string** | The URL called to send status information on every recording event. |
**StatusCallbackMethod** | Pointer to **string** | The HTTP method used to call `status_callback` |
**Links** | Pointer to **map[string]interface{}** | The URLs of related resources |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


