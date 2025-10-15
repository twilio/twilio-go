# VerifyPasskeysFactorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A [base64url](https://base64.guru/standards/base64url) encoded representation of `rawId`. |[optional] 
**RawId** | **string** | The globally unique identifier for this `PublicKeyCredential`. |[optional] 
**AuthenticatorAttachment** | **string** | A string that indicates the mechanism by which the WebAuthn implementation is attached to the authenticator at the time the associated  `navigator.credentials.create()` or `navigator.credentials.get()` call completes. |[optional] 
**Type** | **string** | The valid credential types supported by the API. The values of this enumeration are used for versioning the `AuthenticatorAssertion` and `AuthenticatorAttestation` structures according to the type of the authenticator. |[optional] [default to "public-key"]
**Response** | [**VerifyPasskeysFactorRequestResponse**](VerifyPasskeysFactorRequestResponse.md) |  |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


