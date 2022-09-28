# VerifyV2AccessToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | A string that uniquely identifies this Access Token. |
**AccountSid** | Pointer to **string** | Account Sid. |
**ServiceSid** | Pointer to **string** | Verify Service Sid. |
**EntityIdentity** | Pointer to **string** | Unique external identifier of the Entity |
**FactorType** | Pointer to [**string**](AccessTokenEnumFactorTypes.md) |  |
**FactorFriendlyName** | Pointer to **string** | A human readable description of this factor. |
**Token** | Pointer to **string** | Generated access token. |
**Url** | Pointer to **string** | The URL of this resource. |
**Ttl** | Pointer to **int** | How long, in seconds, the access token is valid. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date this access token was created |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


