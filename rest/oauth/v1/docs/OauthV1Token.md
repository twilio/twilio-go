# OauthV1Token

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | Pointer to **string** | Token which carries the necessary information to access a Twilio resource directly. |
**RefreshToken** | Pointer to **string** | Token which carries the information necessary to get a new access token. |
**IdToken** | Pointer to **string** |  |
**RefreshTokenExpiresAt** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the refresh token expires in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format. |
**AccessTokenExpiresAt** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the refresh token expires in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


