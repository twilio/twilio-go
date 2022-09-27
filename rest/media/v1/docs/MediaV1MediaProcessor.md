# MediaV1MediaProcessor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |
**Extension** | Pointer to **string** | The Media Extension name or URL |
**ExtensionContext** | Pointer to **string** | The Media Extension context |
**Status** | Pointer to [**string**](MediaProcessorEnumStatus.md) |  |
**Url** | Pointer to **string** | The absolute URL of the resource |
**EndedReason** | Pointer to **string** | The reason why a MediaProcessor ended |
**StatusCallback** | Pointer to **string** | The URL to which Twilio will send MediaProcessor event updates |
**StatusCallbackMethod** | Pointer to **string** | The HTTP method Twilio should use to call the `status_callback` URL |
**MaxDuration** | Pointer to **int** | Maximum MediaProcessor duration in seconds |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


