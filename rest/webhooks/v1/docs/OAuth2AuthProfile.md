# OAuth2AuthProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique AuthProfile identifier (TypeID/TTID format) |[readonly] 
**Type** | **string** | Authentication type |
**FriendlyName** | **string** | Human-readable name for identification |
**CreatedAt** | [**time.Time**](time.Time.md) |  |[readonly] 
**UpdatedAt** | [**time.Time**](time.Time.md) |  |[readonly] 
**TokenUrl** | **string** | OAuth2 token endpoint URL |
**ClientId** | **string** | OAuth2 client identifier |
**ClientSecret** | **string** | OAuth2 client secret (never returned, stored securely as secret) |[optional] 
**Scope** | **string** | OAuth2 scope to request |[optional] 
**Audience** | Pointer to **string** | OAuth2 audience parameter |
**ClientCredentialsAuthMethod** | **string** | Client credentials authentication method (defaults to CLIENT_SECRET_POST) |[optional] [default to "CLIENT_SECRET_POST"]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


