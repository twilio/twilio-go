# VideoV1CompositionHook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**AudioSources** | Pointer to **[]string** | The array of track names to include in the compositions created by the composition hook |
**AudioSourcesExcluded** | Pointer to **[]string** | The array of track names to exclude from the compositions created by the composition hook |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |
**Enabled** | Pointer to **bool** | Whether the CompositionHook is active |
**Format** | Pointer to **string** | The container format of the media files used by the compositions created by the composition hook |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the resource |
**Resolution** | Pointer to **string** | The dimensions of the video image in pixels expressed as columns (width) and rows (height) |
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**StatusCallback** | Pointer to **string** | The URL to send status information to your application |
**StatusCallbackMethod** | Pointer to **string** | The HTTP method we should use to call status_callback |
**Trim** | Pointer to **bool** | Whether intervals with no media are clipped |
**Url** | Pointer to **string** | The absolute URL of the resource |
**VideoLayout** | Pointer to **map[string]interface{}** | A JSON object that describes the video layout of the Composition |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


