# VerifyV2NewFactor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | A string that uniquely identifies this Factor. |
**AccountSid** | Pointer to **string** | Account Sid. |
**ServiceSid** | Pointer to **string** | Service Sid. |
**EntitySid** | Pointer to **string** | Entity Sid. |
**Identity** | Pointer to **string** | Unique external identifier of the Entity |
**Binding** | Pointer to **interface{}** | Binding of the factor |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date this Factor was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date this Factor was updated |
**FriendlyName** | Pointer to **string** | A human readable description of this resource. |
**Status** | Pointer to [**string**](NewFactorEnumFactorStatuses.md) |  |
**FactorType** | Pointer to [**string**](NewFactorEnumFactorTypes.md) |  |
**Config** | Pointer to **interface{}** | Configurations for a `factor_type`. |
**Metadata** | Pointer to **interface{}** | Metadata of the factor. |
**Url** | Pointer to **string** | The URL of this resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


