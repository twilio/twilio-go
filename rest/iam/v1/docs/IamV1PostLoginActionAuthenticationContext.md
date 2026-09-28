# IamV1PostLoginActionAuthenticationContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Acr** | **string** | The authentication context class reference the login satisfied |[optional] 
**Amr** | **[]string** | The authentication factors the user completed, as RFC 8176 authentication method reference values (e.g. pwd, otp, mfa) |[optional] 
**Methods** | [**[]IamV1PostLoginActionMethod**](IamV1PostLoginActionMethod.md) | The completed factors and the identifier each resolved, in completion order |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


