# ProxyV1ServiceSession

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**ClosedReason** | Pointer to **string** | The reason the Session ended |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |
**DateEnded** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date when the Session ended |
**DateExpiry** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date when the Session should expire |
**DateLastInteraction** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date when the Session last had an interaction |
**DateStarted** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date when the Session started |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |
**Links** | Pointer to **map[string]interface{}** | The URLs of resources related to the Session |
**Mode** | Pointer to **string** | The Mode of the Session |
**ServiceSid** | Pointer to **string** | The SID of the resource's parent Service |
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**Status** | Pointer to **string** | The status of the Session |
**Ttl** | Pointer to **int32** | When the session will expire |
**UniqueName** | Pointer to **string** | An application-defined string that uniquely identifies the resource |
**Url** | Pointer to **string** | The absolute URL of the Session resource |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


