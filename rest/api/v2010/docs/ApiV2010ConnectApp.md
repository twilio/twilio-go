# ApiV2010ConnectApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**AuthorizeRedirectUrl** | **string** | The URL to redirect the user to after authorization |[optional] 
**CompanyName** | **string** | The company name set for the Connect App |[optional] 
**DeauthorizeCallbackMethod** | **string** | The HTTP method we use to call deauthorize_callback_url |[optional] 
**DeauthorizeCallbackUrl** | **string** | The URL we call to de-authorize the Connect App |[optional] 
**Description** | **string** | The description of the Connect App |[optional] 
**FriendlyName** | **string** | The string that you assigned to describe the resource |[optional] 
**HomepageUrl** | **string** | The URL users can obtain more information |[optional] 
**Permissions** | Pointer to [**[]string**](ConnectAppEnumPermission.md) | The set of permissions that your ConnectApp requests |
**Sid** | **string** | The unique string that identifies the resource |[optional] 
**Uri** | **string** | The URI of the resource, relative to `https://api.twilio.com` |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


