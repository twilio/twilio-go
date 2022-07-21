# VerifyV2Challenge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | A string that uniquely identifies this Challenge. |
**AccountSid** | Pointer to **string** | Account Sid. |
**ServiceSid** | Pointer to **string** | Service Sid. |
**EntitySid** | Pointer to **string** | Entity Sid. |
**Identity** | Pointer to **string** | Unique external identifier of the Entity |
**FactorSid** | Pointer to **string** | Factor Sid. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date this Challenge was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date this Challenge was updated |
**DateResponded** | Pointer to [**time.Time**](time.Time.md) | The date this Challenge was responded |
**ExpirationDate** | Pointer to [**time.Time**](time.Time.md) | The date-time when this Challenge expires |
**Status** | Pointer to [**string**](ChallengeEnumChallengeStatuses.md) |  |
**RespondedReason** | Pointer to [**string**](ChallengeEnumChallengeReasons.md) |  |
**Details** | Pointer to **interface{}** | Details about the Challenge. |
**HiddenDetails** | Pointer to **interface{}** | Hidden details about the Challenge |
**Metadata** | Pointer to **interface{}** | Metadata of the challenge. |
**FactorType** | Pointer to [**string**](ChallengeEnumFactorTypes.md) |  |
**Url** | Pointer to **string** | The URL of this resource. |
**Links** | Pointer to **map[string]interface{}** | Nested resource URLs. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


