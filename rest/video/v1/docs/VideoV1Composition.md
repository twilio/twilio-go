# VideoV1Composition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**Status** | Pointer to [**string**](CompositionEnumStatus.md) |  |
**DateCreated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |[optional] 
**DateCompleted** | Pointer to [**time.Time**](time.Time.md) | Date when the media processing task finished |
**DateDeleted** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the composition generated media was deleted |
**Sid** | **string** | The unique string that identifies the resource |[optional] 
**RoomSid** | **string** | The SID of the Group Room that generated the audio and video tracks used in the composition |[optional] 
**AudioSources** | **[]string** | The array of track names to include in the composition |[optional] 
**AudioSourcesExcluded** | **[]string** | The array of track names to exclude from the composition |[optional] 
**VideoLayout** | Pointer to **interface{}** | An object that describes the video layout of the composition |
**Resolution** | **string** | The dimensions of the video image in pixels expressed as columns (width) and rows (height) |[optional] 
**Trim** | **bool** | Whether to remove intervals with no media |[optional] 
**Format** | Pointer to [**string**](CompositionEnumFormat.md) |  |
**Bitrate** | **int** | The average bit rate of the composition's media |[optional] 
**Size** | **int64** | The size of the composed media file in bytes |[optional] 
**Duration** | **int** | The duration of the composition's media file in seconds |[optional] 
**MediaExternalLocation** | **string** | The URL of the media file associated with the composition when stored externally |[optional] 
**StatusCallback** | **string** | The URL called to send status information on every composition event. |[optional] 
**StatusCallbackMethod** | **string** | The HTTP method used to call `status_callback` |[optional] 
**Url** | **string** | The absolute URL of the resource |[optional] 
**Links** | **map[string]interface{}** | The URL of the media file associated with the composition |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


