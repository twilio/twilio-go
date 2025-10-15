# ApprovePasskeysChallengeRequestResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticatorData** | **string** | The [authenticator data](https://developer.mozilla.org/en-US/docs/Web/API/Web_Authentication_API/Authenticator_data) structure contains information from the authenticator about the processing of a credential creation or authentication request. |
**ClientDataJSON** | **string** | This property contains the JSON-compatible serialization of the data passed from the browser to the authenticator in order to generate this credential. |
**Signature** | **string** | An assertion signature over `authenticatorData` and `clientDataJSON`. The assertion signature is created with the private key of the key pair that was created during the originating `navigator.credentials.create()` call and verified using the public key of that same key pair. |
**UserHandle** | **string** | The user handle stored in the authenticator, specified as `user.id` in the options passed to the originating `navigator.credentials.create()` call. This property should contain a base64url-encoded entity SID. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


