# VerifyV2Challenge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | **string** | A string that uniquely identifies this Challenge. |[optional] 
**AccountSid** | **string** | Account Sid. |[optional] 
**ServiceSid** | **string** | Service Sid. |[optional] 
**EntitySid** | **string** | Entity Sid. |[optional] 
**Identity** | **string** | Unique external identifier of the Entity |[optional] 
**FactorSid** | **string** | Factor Sid. |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The date this Challenge was created |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The date this Challenge was updated |[optional] 
**DateResponded** | [**time.Time**](time.Time.md) | The date this Challenge was responded |[optional] 
**ExpirationDate** | [**time.Time**](time.Time.md) | The date-time when this Challenge expires |[optional] 
**Status** | Pointer to [**string**](ChallengeEnumChallengeStatuses.md) |  |
**RespondedReason** | Pointer to [**string**](ChallengeEnumChallengeReasons.md) |  |
**Details** | Pointer to **interface{}** | Details about the Challenge. |
**HiddenDetails** | Pointer to **interface{}** | Hidden details about the Challenge |
**Metadata** | Pointer to **interface{}** | Metadata of the challenge. |
**FactorType** | Pointer to [**string**](ChallengeEnumFactorTypes.md) |  |
**Url** | **string** | The URL of this resource. |[optional] 
**Links** | **map[string]interface{}** | Nested resource URLs. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


