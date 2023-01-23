# ApiV2010AuthorizedConnectApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the AuthorizedConnectApp resource. |
**ConnectAppCompanyName** | Pointer to **string** | The company name set for the Connect App. |
**ConnectAppDescription** | Pointer to **string** | A detailed description of the Connect App. |
**ConnectAppFriendlyName** | Pointer to **string** | The name of the Connect App. |
**ConnectAppHomepageUrl** | Pointer to **string** | The public URL for the Connect App. |
**ConnectAppSid** | Pointer to **string** | The SID that we assigned to the Connect App. |
**DateCreated** | Pointer to **string** | The date and time in GMT that the resource was created specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format. |
**DateUpdated** | Pointer to **string** | The date and time in GMT that the resource was last updated specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format. |
**Permissions** | Pointer to [**[]string**](AuthorizedConnectAppEnumPermission.md) | The set of permissions that you authorized for the Connect App.  Can be: `get-all` or `post-all`. |
**Uri** | Pointer to **string** | The URI of the resource, relative to `https://api.twilio.com`. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


