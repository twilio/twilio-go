# OauthV1OpenidDiscovery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Issuer** | Pointer to **string** | The URL of the party that will create the token and sign it with its private key. |
**AuthorizationEndpoint** | Pointer to **string** | The endpoint that validates all authorization requests. |
**DeviceAuthorizationEndpoint** | Pointer to **string** | The endpoint that validates all device code related authorization requests. |
**TokenEndpoint** | Pointer to **string** | The URL of the token endpoint. After a client has received an authorization code, that code is presented to the token endpoint and exchanged for an identity token, an access token, and a refresh token. |
**UserinfoEndpoint** | Pointer to **string** | The URL of the user info endpoint, which returns user profile information to a client. Keep in mind that the user info endpoint returns only the information that has been requested. |
**RevocationEndpoint** | Pointer to **string** | The endpoint used to revoke access or refresh tokens issued by the authorization server. |
**JwkUri** | Pointer to **string** | The URL of your JSON Web Key Set. This set is a collection of JSON Web Keys, a standard method for representing cryptographic keys in a JSON structure. |
**ResponseTypeSupported** | Pointer to **[]string** | A collection of response type supported by authorization server. |
**SubjectTypeSupported** | Pointer to **[]string** | A collection of subject by authorization server. |
**IdTokenSigningAlgValuesSupported** | Pointer to **[]string** | A collection of JWS signing algorithms supported by authorization server to sign identity token. |
**ScopesSupported** | Pointer to **[]string** | A collection of scopes supported by authorization server for identity token |
**ClaimsSupported** | Pointer to **[]string** | A collection of claims supported by authorization server for identity token |
**Url** | Pointer to **string** |  |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


