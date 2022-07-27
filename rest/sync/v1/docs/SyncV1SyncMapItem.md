# SyncV1SyncMapItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | The unique, user-defined key for the Map Item |
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**ServiceSid** | Pointer to **string** | The SID of the Sync Service that the resource is associated with |
**MapSid** | Pointer to **string** | The SID of the Sync Map that contains the Map Item |
**Url** | Pointer to **string** | The absolute URL of the Map Item resource |
**Revision** | Pointer to **string** | The current revision of the Map Item, represented as a string |
**Data** | Pointer to **interface{}** | An arbitrary, schema-less object that the Map Item stores |
**DateExpires** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the Map Item expires |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |
**CreatedBy** | Pointer to **string** | The identity of the Map Item's creator |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


