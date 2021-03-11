# SyncV1ServiceDocument

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**CreatedBy** | Pointer to **string** | The identity of the Sync Document's creator |
**Data** | Pointer to **map[string]interface{}** | An arbitrary, schema-less object that the Sync Document stores |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |
**DateExpires** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the Sync Document expires |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |
**Links** | Pointer to **map[string]interface{}** | The URLs of resources related to the Sync Document |
**Revision** | Pointer to **string** | The current revision of the Sync Document, represented by a string identifier |
**ServiceSid** | Pointer to **string** | The SID of the Sync Service that the resource is associated with |
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**UniqueName** | Pointer to **string** | An application-defined string that uniquely identifies the resource |
**Url** | Pointer to **string** | The absolute URL of the Document resource |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


