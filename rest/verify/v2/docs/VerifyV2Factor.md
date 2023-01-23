# VerifyV2Factor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | **string** | A string that uniquely identifies this Factor. |[optional] 
**AccountSid** | **string** | Account Sid. |[optional] 
**ServiceSid** | **string** | Service Sid. |[optional] 
**EntitySid** | **string** | Entity Sid. |[optional] 
**Identity** | **string** | Unique external identifier of the Entity |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The date this Factor was created |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The date this Factor was updated |[optional] 
**FriendlyName** | **string** | A human readable description of this resource. |[optional] 
**Status** | Pointer to [**string**](FactorEnumFactorStatuses.md) |  |
**FactorType** | Pointer to [**string**](FactorEnumFactorTypes.md) |  |
**Config** | Pointer to **interface{}** | Configurations for a `factor_type`. |
**Metadata** | Pointer to **interface{}** | Metadata of the factor. |
**Url** | **string** | The URL of this resource. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


