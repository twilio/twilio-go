# NumbersV2ReplaceItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**RegulationSid** | Pointer to **string** | The unique string of a regulation |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the resource |
**Status** | Pointer to [**string**](ReplaceItemsEnumStatus.md) |  |
**ValidUntil** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource will be valid until |
**Email** | Pointer to **string** | The email address |
**StatusCallback** | Pointer to **string** | The URL we call to inform your application of status changes |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


