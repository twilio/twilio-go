# VideoV1Composition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**Status** | Pointer to [**string**](CompositionEnumStatus.md) |  |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |
**DateCompleted** | Pointer to [**time.Time**](time.Time.md) | Date when the media processing task finished |
**DateDeleted** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the composition generated media was deleted |
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**RoomSid** | Pointer to **string** | The SID of the Group Room that generated the audio and video tracks used in the composition |
**AudioSources** | Pointer to **[]string** | The array of track names to include in the composition |
**AudioSourcesExcluded** | Pointer to **[]string** | The array of track names to exclude from the composition |
**VideoLayout** | Pointer to **interface{}** | An object that describes the video layout of the composition |
**Resolution** | Pointer to **string** | The dimensions of the video image in pixels expressed as columns (width) and rows (height) |
**Trim** | Pointer to **bool** | Whether to remove intervals with no media |
**Format** | Pointer to [**string**](CompositionEnumFormat.md) |  |
**Bitrate** | Pointer to **int** | The average bit rate of the composition's media |
**Size** | Pointer to **int64** | The size of the composed media file in bytes |
**Duration** | Pointer to **int** | The duration of the composition's media file in seconds |
**MediaExternalLocation** | Pointer to **string** | The URL of the media file associated with the composition when stored externally |
**StatusCallback** | Pointer to **string** | The URL called to send status information on every composition event. |
**StatusCallbackMethod** | Pointer to **string** | The HTTP method used to call `status_callback` |
**Url** | Pointer to **string** | The absolute URL of the resource |
**Links** | Pointer to **map[string]interface{}** | The URL of the media file associated with the composition |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


