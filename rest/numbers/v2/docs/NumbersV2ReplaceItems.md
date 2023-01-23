# NumbersV2ReplaceItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that we created to identify the Bundle resource. |
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Bundle resource. |
**RegulationSid** | Pointer to **string** | The unique string of a regulation that is associated to the Bundle resource. |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the resource. |
**Status** | Pointer to [**string**](ReplaceItemsEnumStatus.md) |  |
**ValidUntil** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format when the resource will be valid until. |
**Email** | Pointer to **string** | The email address that will receive updates when the Bundle resource changes status. |
**StatusCallback** | Pointer to **string** | The URL we call to inform your application of status changes. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


