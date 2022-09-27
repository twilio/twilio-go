# SyncV1Document

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**UniqueName** | Pointer to **string** | An application-defined string that uniquely identifies the resource |
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**ServiceSid** | Pointer to **string** | The SID of the Sync Service that the resource is associated with |
**Url** | Pointer to **string** | The absolute URL of the Document resource |
**Links** | Pointer to **map[string]interface{}** | The URLs of resources related to the Sync Document |
**Revision** | Pointer to **string** | The current revision of the Sync Document, represented by a string identifier |
**Data** | Pointer to **interface{}** | An arbitrary, schema-less object that the Sync Document stores |
**DateExpires** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the Sync Document expires |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |
**CreatedBy** | Pointer to **string** | The identity of the Sync Document's creator |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


