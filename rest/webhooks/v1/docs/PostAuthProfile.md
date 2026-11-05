# PostAuthProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Authentication type |
**FriendlyName** | Pointer to **string** | Human-readable name for identification (optional, defaults to blank) |
**TokenUrl** | **string** | OAuth2 token endpoint URL |
**ClientId** | **string** | OAuth2 client identifier |
**ClientSecret** | **string** | OAuth2 client secret (never returned, stored securely as secret) |
**Scope** | **string** | OAuth2 scope to request |[optional] 
**Audience** | **string** | OAuth2 audience parameter |[optional] 
**ClientCredentialsAuthMethod** | **string** | Client credentials authentication method (defaults to CLIENT_SECRET_POST) |[optional] [default to "CLIENT_SECRET_POST"]
**Username** | **string** | Digest auth username |
**Password** | **string** | Digest auth password  (never returned, stored securely as secret) |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


