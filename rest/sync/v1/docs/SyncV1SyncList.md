# SyncV1SyncList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that we created to identify the Sync List resource. |
**UniqueName** | Pointer to **string** | An application-defined string that uniquely identifies the resource. It can be used in place of the resource's `sid` in the URL to address the resource. |
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Sync List resource. |
**ServiceSid** | Pointer to **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) the resource is associated with. |
**Url** | Pointer to **string** | The absolute URL of the Sync List resource. |
**Links** | Pointer to **map[string]interface{}** | The URLs of the Sync List's nested resources. |
**Revision** | Pointer to **string** | The current revision of the Sync List, represented as a string. |
**DateExpires** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the Sync List expires and will be deleted, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. If the Sync List does not expire, this value is `null`. The Sync List might not be deleted immediately after it expires. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**CreatedBy** | Pointer to **string** | The identity of the Sync List's creator. If the Sync List is created from the client SDK, the value matches the Access Token's `identity` field. If the Sync List was created from the REST API, the value is `system`. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


