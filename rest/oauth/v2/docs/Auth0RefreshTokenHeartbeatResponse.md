# Auth0RefreshTokenHeartbeatResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsValid** | **bool** | True if session is valid, false if expired or invalid |
**ExpiresAt** | Pointer to [**time.Time**](time.Time.md) | datetime string when the session expires, only present if isValid is true |
**LastActiveAt** | Pointer to [**time.Time**](time.Time.md) | datetime string when the session was refreshed, only present if isValid is true |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


