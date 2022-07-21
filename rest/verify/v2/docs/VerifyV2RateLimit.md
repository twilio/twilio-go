# VerifyV2RateLimit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | A string that uniquely identifies this Rate Limit. |
**ServiceSid** | Pointer to **string** | The SID of the Service that the resource is associated with |
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**UniqueName** | Pointer to **string** | A unique, developer assigned name of this Rate Limit. |
**Description** | Pointer to **string** | Description of this Rate Limit |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was last updated |
**Url** | Pointer to **string** | The URL of this resource. |
**Links** | Pointer to **map[string]interface{}** | The URLs of related resources |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


