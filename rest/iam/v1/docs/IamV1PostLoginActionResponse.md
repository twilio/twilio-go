# IamV1PostLoginActionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | Whether the login should be allowed or denied |[optional] 
**IdToken** | Pointer to [**IamV1PostLoginActionTokenClaims**](IamV1PostLoginActionTokenClaims.md) |  |
**AccessToken** | Pointer to [**IamV1PostLoginActionTokenClaims**](IamV1PostLoginActionTokenClaims.md) |  |
**Reason** | **string** | The customer-facing message explaining why the login was denied. Present only when action is deny |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


