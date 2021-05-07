# NotifyV1Credential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was last updated |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the resource |
**Sandbox** | Pointer to **string** | [APN only] Whether to send the credential to sandbox APNs |
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**Type** | Pointer to **string** | The Credential type, one of `gcm`, `fcm`, or `apn` |
**Url** | Pointer to **string** | The absolute URL of the Credential resource |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


