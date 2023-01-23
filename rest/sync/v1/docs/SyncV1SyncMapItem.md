# SyncV1SyncMapItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | The unique, user-defined key for the Map Item |[optional] 
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**ServiceSid** | **string** | The SID of the Sync Service that the resource is associated with |[optional] 
**MapSid** | **string** | The SID of the Sync Map that contains the Map Item |[optional] 
**Url** | **string** | The absolute URL of the Map Item resource |[optional] 
**Revision** | **string** | The current revision of the Map Item, represented as a string |[optional] 
**Data** | Pointer to **interface{}** | An arbitrary, schema-less object that the Map Item stores |
**DateExpires** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the Map Item expires |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |[optional] 
**CreatedBy** | **string** | The identity of the Map Item's creator |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


