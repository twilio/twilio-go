# VideoV1Composition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**AudioSources** | Pointer to **[]string** | The array of track names to include in the composition |
**AudioSourcesExcluded** | Pointer to **[]string** | The array of track names to exclude from the composition |
**Bitrate** | Pointer to **int** | The average bit rate of the composition's media |
**DateCompleted** | Pointer to [**time.Time**](time.Time.md) | Date when the media processing task finished |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |
**DateDeleted** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the composition generated media was deleted |
**Duration** | Pointer to **int** | The duration of the composition's media file in seconds |
**Format** | Pointer to **string** | The container format of the composition's media files as specified in the POST request that created the Composition resource |
**Links** | Pointer to **map[string]interface{}** | The URL of the media file associated with the composition |
**MediaExternalLocation** | Pointer to **string** | The URL of the media file associated with the composition when stored externally |
**Resolution** | Pointer to **string** | The dimensions of the video image in pixels expressed as columns (width) and rows (height) |
**RoomSid** | Pointer to **string** | The SID of the Group Room that generated the audio and video tracks used in the composition |
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**Size** | Pointer to **int** | The size of the composed media file in bytes |
**Status** | Pointer to **string** | The status of the composition |
**Trim** | Pointer to **bool** | Whether to remove intervals with no media |
**Url** | Pointer to **string** | The absolute URL of the resource |
**VideoLayout** | Pointer to **map[string]interface{}** | An object that describes the video layout of the composition |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


