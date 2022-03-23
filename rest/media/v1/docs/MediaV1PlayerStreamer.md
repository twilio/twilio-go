# MediaV1PlayerStreamer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |
**EndedReason** | Pointer to **string** | The reason why a PlayerStreamer ended |
**Links** | Pointer to **map[string]interface{}** | The URLs of related resources |
**MaxDuration** | Pointer to **int** | Maximum PlayerStreamer duration in seconds |
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**Status** | Pointer to **string** | The status of the PlayerStreamer |
**StatusCallback** | Pointer to **string** | The URL to which Twilio will send PlayerStreamer event updates |
**StatusCallbackMethod** | Pointer to **string** | The HTTP method Twilio should use to call the `status_callback` URL |
**Url** | Pointer to **string** | The absolute URL of the resource |
**Video** | Pointer to **bool** | Whether the PlayerStreamer is configured to stream video |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


