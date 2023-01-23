# ApiV2010Siprec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The SID of the Siprec resource. |
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created this Siprec resource. |
**CallSid** | Pointer to **string** | The SID of the [Call](https://www.twilio.com/docs/voice/api/call-resource) the Siprec resource is associated with. |
**Name** | Pointer to **string** | The user-specified name of this Siprec, if one was given when the Siprec was created. This may be used to stop the Siprec. |
**Status** | Pointer to [**string**](SiprecEnumStatus.md) |  |
**DateUpdated** | Pointer to **string** | The date and time in GMT that this resource was last updated, specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format. |
**Uri** | Pointer to **string** | The URI of the resource, relative to `https://api.twilio.com`. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


