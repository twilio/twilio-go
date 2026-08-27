# IamV1VendorOauthAppCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  |[optional] 
**FriendlyName** | **string** |  |[optional] 
**OwnerSid** | Pointer to **string** |  |
**Description** | **string** |  |[optional] 
**ClientSid** | Pointer to **string** |  |
**TokenEndpointAuthMethod** | **string** | Determines how the client authenticates. Account OAuth apps on v1 only support 'client_secret_basic'. For PKCE (none), use the v2 API. |[optional] 
**Policy** | [**IamV1OrganizationVendoroauthappPolicy**](IamV1OrganizationVendoroauthappPolicy.md) |  |[optional] 
**AccessTokenTtl** | **int** |  |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


