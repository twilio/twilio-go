# SyncV1SyncMap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**UniqueName** | Pointer to **string** | An application-defined string that uniquely identifies the resource |
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**ServiceSid** | Pointer to **string** | The SID of the Sync Service that the resource is associated with |
**Url** | Pointer to **string** | The absolute URL of the Sync Map resource |
**Links** | Pointer to **map[string]interface{}** | The URLs of the Sync Map's nested resources |
**Revision** | Pointer to **string** | The current revision of the Sync Map, represented as a string |
**DateExpires** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the Sync Map expires |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |
**CreatedBy** | Pointer to **string** | The identity of the Sync Map's creator |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


