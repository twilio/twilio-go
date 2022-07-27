# SyncV1SyncListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | Pointer to **int** | The automatically generated index of the List Item |
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**ServiceSid** | Pointer to **string** | The SID of the Sync Service that the resource is associated with |
**ListSid** | Pointer to **string** | The SID of the Sync List that contains the List Item |
**Url** | Pointer to **string** | The absolute URL of the List Item resource |
**Revision** | Pointer to **string** | The current revision of the item, represented as a string |
**Data** | Pointer to **interface{}** | An arbitrary, schema-less object that the List Item stores |
**DateExpires** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the List Item expires |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |
**CreatedBy** | Pointer to **string** | The identity of the List Item's creator |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


