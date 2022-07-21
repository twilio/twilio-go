# ProxyV1Session

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**ServiceSid** | Pointer to **string** | The SID of the resource's parent Service |
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**DateStarted** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date when the Session started |
**DateEnded** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date when the Session ended |
**DateLastInteraction** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date when the Session last had an interaction |
**DateExpiry** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date when the Session should expire |
**UniqueName** | Pointer to **string** | An application-defined string that uniquely identifies the resource |
**Status** | Pointer to [**string**](SessionEnumStatus.md) |  |
**ClosedReason** | Pointer to **string** | The reason the Session ended |
**Ttl** | Pointer to **int** | When the session will expire |
**Mode** | Pointer to [**string**](SessionEnumMode.md) |  |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |
**Url** | Pointer to **string** | The absolute URL of the Session resource |
**Links** | Pointer to **map[string]interface{}** | The URLs of resources related to the Session |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


