# ApiV2010Account

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthToken** | Pointer to **string** | The authorization token for this account. This token should be kept a secret, so no sharing. |
**DateCreated** | Pointer to **string** | The date that this account was created, in GMT in RFC 2822 format |
**DateUpdated** | Pointer to **string** | The date that this account was last updated, in GMT in RFC 2822 format. |
**FriendlyName** | Pointer to **string** | A human readable description of this account, up to 64 characters long. By default the FriendlyName is your email address. |
**OwnerAccountSid** | Pointer to **string** | The unique 34 character id that represents the parent of this account. The OwnerAccountSid of a parent account is it's own sid. |
**Sid** | Pointer to **string** | A 34 character string that uniquely identifies this resource. |
**Status** | Pointer to [**string**](AccountEnumStatus.md) |  |
**SubresourceUris** | Pointer to **map[string]interface{}** | A Map of various subresources available for the given Account Instance |
**Type** | Pointer to [**string**](AccountEnumType.md) |  |
**Uri** | Pointer to **string** | The URI for this resource, relative to `https://api.twilio.com` |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


