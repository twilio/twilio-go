# NumbersV1ProofOfAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that we created to identify the ProofOfAddress resource. |
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the ProofOfAddress resource. |
**AddressSid** | Pointer to **string** | The SID of the [Address](https://www.twilio.com/docs/usage/api/address) resource that is associated with the ProofOfAddress resource. |
**IdentitySid** | Pointer to **string** | The SID of the Identity resource that is associated with the ProofOfAddress resource. |
**Status** | Pointer to **string** | The verification status of the address. |
**FailureReason** | Pointer to **string** | The reason the address verification failed. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**Url** | Pointer to **string** | The absolute URL of the ProofOfAddress resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


