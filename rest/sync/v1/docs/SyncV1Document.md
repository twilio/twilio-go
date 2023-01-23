# SyncV1Document

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that we created to identify the Document resource. |
**UniqueName** | Pointer to **string** | An application-defined string that uniquely identifies the resource. It can be used in place of the resource's `sid` in the URL to address the resource and can be up to 320 characters long. |
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Document resource. |
**ServiceSid** | Pointer to **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) the resource is associated with. |
**Url** | Pointer to **string** | The absolute URL of the Document resource. |
**Links** | Pointer to **map[string]interface{}** | The URLs of resources related to the Sync Document. |
**Revision** | Pointer to **string** | The current revision of the Sync Document, represented as a string. The `revision` property is used with conditional updates to ensure data consistency. |
**Data** | Pointer to **interface{}** | An arbitrary, schema-less object that the Sync Document stores. Can be up to 16 KiB in length. |
**DateExpires** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the Sync Document expires and will be deleted, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. If the Sync Document does not expire, this value is `null`. The Document resource might not be deleted immediately after it expires. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**CreatedBy** | Pointer to **string** | The identity of the Sync Document's creator. If the Sync Document is created from the client SDK, the value matches the Access Token's `identity` field. If the Sync Document was created from the REST API, the value is `system`. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


