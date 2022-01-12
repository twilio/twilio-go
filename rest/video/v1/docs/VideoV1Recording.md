# VideoV1Recording

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**Codec** | Pointer to **string** | The codec used to encode the track |
**ContainerFormat** | Pointer to **string** | The file format for the recording |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |
**Duration** | Pointer to **int** | The duration of the recording in seconds |
**GroupingSids** | Pointer to **map[string]interface{}** | A list of SIDs related to the recording |
**Links** | Pointer to **map[string]interface{}** | The URLs of related resources |
**MediaExternalLocation** | Pointer to **string** | The URL of the media file associated with the recording when stored externally |
**Offset** | Pointer to **int** | The number of milliseconds between a point in time that is common to all rooms in a group and when the source room of the recording started |
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**Size** | Pointer to **int** | The size of the recorded track, in bytes |
**SourceSid** | Pointer to **string** | The SID of the recording source |
**Status** | Pointer to **string** | The status of the recording |
**TrackName** | Pointer to **string** | The name that was given to the source track of the recording |
**Type** | Pointer to **string** | The recording's media type |
**Url** | Pointer to **string** | The absolute URL of the resource |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


