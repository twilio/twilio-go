# VerifyV2AccessToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | A 34 character string that uniquely identifies this Access Token. |
**AccountSid** | Pointer to **string** | The unique SID identifier of the Account. |
**ServiceSid** | Pointer to **string** | The unique SID identifier of the Verify Service. |
**EntityIdentity** | Pointer to **string** | The unique external identifier for the Entity of the Service. |
**FactorType** | Pointer to [**string**](AccessTokenEnumFactorTypes.md) |  |
**FactorFriendlyName** | Pointer to **string** | A human readable description of this factor, up to 64 characters. For a push factor, this can be the device's name. |
**Token** | Pointer to **string** | The access token generated for enrollment, this is an encrypted json web token. |
**Url** | Pointer to **string** | The URL of this resource. |
**Ttl** | Pointer to **int** | How long, in seconds, the access token is valid. Max: 5 minutes |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date that this access token was created, given in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


