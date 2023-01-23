# VideoV1CompositionHook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**FriendlyName** | **string** | The string that you assigned to describe the resource |[optional] 
**Enabled** | **bool** | Whether the CompositionHook is active |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |[optional] 
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |
**Sid** | **string** | The unique string that identifies the resource |[optional] 
**AudioSources** | **[]string** | The array of track names to include in the compositions created by the composition hook |[optional] 
**AudioSourcesExcluded** | **[]string** | The array of track names to exclude from the compositions created by the composition hook |[optional] 
**VideoLayout** | Pointer to **interface{}** | A JSON object that describes the video layout of the Composition |
**Resolution** | **string** | The dimensions of the video image in pixels expressed as columns (width) and rows (height) |[optional] 
**Trim** | **bool** | Whether intervals with no media are clipped |[optional] 
**Format** | Pointer to [**string**](CompositionHookEnumFormat.md) |  |
**StatusCallback** | Pointer to **string** | The URL to send status information to your application |
**StatusCallbackMethod** | **string** | The HTTP method we should use to call status_callback |[optional] 
**Url** | **string** | The absolute URL of the resource |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


