# IamV1PostLoginActionMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The credential type, e.g. password, email, oauth |[optional] 
**Identifier** | **string** | The contact value the factor resolved, from the user's verified factor directory. Omitted together with identifierType when the factor resolved none, e.g. a passkey login |[optional] 
**IdentifierType** | **string** | The kind of identifier, e.g. email or phone |[optional] 
**CompletedAt** | [**time.Time**](time.Time.md) | RFC 3339 timestamp the factor completed at |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


