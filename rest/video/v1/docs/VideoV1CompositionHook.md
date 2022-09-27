# VideoV1CompositionHook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the resource |
**Enabled** | Pointer to **bool** | Whether the CompositionHook is active |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**AudioSources** | Pointer to **[]string** | The array of track names to include in the compositions created by the composition hook |
**AudioSourcesExcluded** | Pointer to **[]string** | The array of track names to exclude from the compositions created by the composition hook |
**VideoLayout** | Pointer to **interface{}** | A JSON object that describes the video layout of the Composition |
**Resolution** | Pointer to **string** | The dimensions of the video image in pixels expressed as columns (width) and rows (height) |
**Trim** | Pointer to **bool** | Whether intervals with no media are clipped |
**Format** | Pointer to [**string**](CompositionHookEnumFormat.md) |  |
**StatusCallback** | Pointer to **string** | The URL to send status information to your application |
**StatusCallbackMethod** | Pointer to **string** | The HTTP method we should use to call status_callback |
**Url** | Pointer to **string** | The absolute URL of the resource |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


