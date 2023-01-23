# ChatV1Credential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that we created to identify the Credential resource. |
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/api/rest/account) that created the Credential resource. |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the resource. |
**Type** | Pointer to [**string**](CredentialEnumPushService.md) |  |
**Sandbox** | Pointer to **string** | [APN only] Whether to send the credential to sandbox APNs. Can be `true` to send to sandbox APNs or `false` to send to production. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was created specified in [RFC 2822](http://www.ietf.org/rfc/rfc2822.txt) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was last updated specified in [RFC 2822](http://www.ietf.org/rfc/rfc2822.txt) format. |
**Url** | Pointer to **string** | The absolute URL of the Credential resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


