# CreateNewFactorPasskeyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | A 34 character string that uniquely identifies this Factor. |
**AccountSid** | Pointer to **string** | The unique SID identifier of the Account. |
**ServiceSid** | Pointer to **string** | The unique SID identifier of the Service. |
**EntitySid** | Pointer to **string** | The unique SID identifier of the Entity. |
**Identity** | Pointer to **string** | Customer unique identity for the Entity owner of the Factor. |
**Binding** | Pointer to **interface{}** | Contains the `factor_type` specific secret and metadata. The Binding property is ONLY returned upon Factor creation. |
**Options** | Pointer to **interface{}** |  |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date that this Factor was created, given in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date that this Factor was updated, given in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**FriendlyName** | Pointer to **string** | The friendly name of this Factor. This can be any string up to 64 characters, meant for humans to distinguish between Factors. |
**Status** | **string** | The Status of this Factor. One of `unverified` or `verified`. |[optional] 
**FactorType** | **string** | The Type of this Factor. Currently `push` and `totp` are supported. |[optional] 
**Config** | Pointer to **interface{}** | An object that contains configurations specific to a `factor_type`. |
**Metadata** | Pointer to **interface{}** | Custom metadata associated with the factor. |
**Url** | Pointer to **string** | The URL of this resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


