# MediaV1PlayerStreamer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |[optional] 
**Video** | **bool** | Whether the PlayerStreamer is configured to stream video |[optional] 
**Links** | **map[string]interface{}** | The URLs of related resources |[optional] 
**Sid** | **string** | The unique string that identifies the resource |[optional] 
**Status** | Pointer to [**string**](PlayerStreamerEnumStatus.md) |  |
**Url** | **string** | The absolute URL of the resource |[optional] 
**StatusCallback** | **string** | The URL to which Twilio will send PlayerStreamer event updates |[optional] 
**StatusCallbackMethod** | **string** | The HTTP method Twilio should use to call the `status_callback` URL |[optional] 
**EndedReason** | Pointer to [**string**](PlayerStreamerEnumEndedReason.md) |  |
**MaxDuration** | **int** | Maximum PlayerStreamer duration in seconds |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


