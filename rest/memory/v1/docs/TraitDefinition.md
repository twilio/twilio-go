# TraitDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataType** | **string** | Data type of the Trait, such as STRING, NUMBER, BOOLEAN, ARRAY. For DELETE operations in PATCH requests, set this to an empty string (\"\") to mark the trait for deletion. |
**Description** | **string** | Description of the Trait, providing additional context or information. |[optional] 
**IdTypePromotion** | **string** | The name of the identifier type to promote the trait value to, such as ''email'', ''phone'', ''user_id'', etc. This allows the trait to be mapped to an identifier in Identity Resolution. The identifier type should be configured in the Identity Resolution Settings. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


