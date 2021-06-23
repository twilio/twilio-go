# SyncV1ServiceSyncListSyncListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**CreatedBy** | Pointer to **string** | The identity of the List Item's creator |
**Data** | Pointer to **map[string]interface{}** | An arbitrary, schema-less object that the List Item stores |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |
**DateExpires** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the List Item expires |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |
**Index** | Pointer to **int** | The automatically generated index of the List Item |
**ListSid** | Pointer to **string** | The SID of the Sync List that contains the List Item |
**Revision** | Pointer to **string** | The current revision of the item, represented as a string |
**ServiceSid** | Pointer to **string** | The SID of the Sync Service that the resource is associated with |
**Url** | Pointer to **string** | The absolute URL of the List Item resource |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


