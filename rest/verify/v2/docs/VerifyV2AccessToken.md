# VerifyV2AccessToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | **string** | A string that uniquely identifies this Access Token. |[optional] 
**AccountSid** | **string** | Account Sid. |[optional] 
**ServiceSid** | **string** | Verify Service Sid. |[optional] 
**EntityIdentity** | **string** | Unique external identifier of the Entity |[optional] 
**FactorType** | Pointer to [**string**](AccessTokenEnumFactorTypes.md) |  |
**FactorFriendlyName** | **string** | A human readable description of this factor. |[optional] 
**Token** | **string** | Generated access token. |[optional] 
**Url** | **string** | The URL of this resource. |[optional] 
**Ttl** | **int** | How long, in seconds, the access token is valid. |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The date this access token was created |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


