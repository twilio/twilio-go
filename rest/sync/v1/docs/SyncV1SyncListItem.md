# SyncV1SyncListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | **int** | The automatically generated index of the List Item |[optional] 
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**ServiceSid** | **string** | The SID of the Sync Service that the resource is associated with |[optional] 
**ListSid** | **string** | The SID of the Sync List that contains the List Item |[optional] 
**Url** | **string** | The absolute URL of the List Item resource |[optional] 
**Revision** | **string** | The current revision of the item, represented as a string |[optional] 
**Data** | Pointer to **interface{}** | An arbitrary, schema-less object that the List Item stores |
**DateExpires** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the List Item expires |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |[optional] 
**CreatedBy** | **string** | The identity of the List Item's creator |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


