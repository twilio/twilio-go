# ApiV2010ConnectApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the ConnectApp resource. |
**AuthorizeRedirectUrl** | Pointer to **string** | The URL we redirect the user to after we authenticate the user and obtain authorization to access the Connect App. |
**CompanyName** | Pointer to **string** | The company name set for the Connect App. |
**DeauthorizeCallbackMethod** | Pointer to **string** | The HTTP method we use to call `deauthorize_callback_url`. |
**DeauthorizeCallbackUrl** | Pointer to **string** | The URL we call using the `deauthorize_callback_method` to de-authorize the Connect App. |
**Description** | Pointer to **string** | The description of the Connect App. |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the resource. |
**HomepageUrl** | Pointer to **string** | The public URL where users can obtain more information about this Connect App. |
**Permissions** | Pointer to [**[]string**](ConnectAppEnumPermission.md) | The set of permissions that your ConnectApp requests. |
**Sid** | Pointer to **string** | The unique string that that we created to identify the ConnectApp resource. |
**Uri** | Pointer to **string** | The URI of the resource, relative to `https://api.twilio.com`. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


