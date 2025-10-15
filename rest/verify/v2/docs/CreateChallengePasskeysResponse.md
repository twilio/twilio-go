# CreateChallengePasskeysResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Options** | **map[string]interface{}** | An object that contains challenge options. Currently only used for `passkeys`. |[optional] 
**Sid** | Pointer to **string** | A 34 character string that uniquely identifies this Challenge. |
**AccountSid** | Pointer to **string** | The unique SID identifier of the Account. |
**ServiceSid** | Pointer to **string** | The unique SID identifier of the Service. |
**EntitySid** | Pointer to **string** | The unique SID identifier of the Entity. |
**Identity** | Pointer to **string** | Customer unique identity for the Entity owner of the Challenge. |
**FactorSid** | Pointer to **string** | The unique SID identifier of the Factor. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date that this Challenge was created, given in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date that this Challenge was updated, given in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**DateResponded** | Pointer to [**time.Time**](time.Time.md) | The date that this Challenge was responded, given in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**ExpirationDate** | Pointer to [**time.Time**](time.Time.md) | The date-time when this Challenge expires, given in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**Status** | **string** | The Status of this Challenge. One of `pending`, `expired`, `approved` or `denied`. |[optional] 
**RespondedReason** | **string** | Reason for the Challenge to be in certain `status`. One of `none`, `not_needed` or `not_requested`. |[optional] 
**Details** | Pointer to **interface{}** | Details provided to give context about the Challenge. |
**HiddenDetails** | Pointer to **interface{}** | Details provided to give context about the Challenge. |
**Metadata** | Pointer to **interface{}** | Custom metadata associated with the challenge. |
**FactorType** | **string** | The Factor Type of this Challenge. Currently `push` and `totp` are supported. |[optional] 
**Url** | Pointer to **string** | The URL of this resource. |
**Links** | Pointer to **map[string]interface{}** | Contains a dictionary of URL links to nested resources of this Challenge. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


