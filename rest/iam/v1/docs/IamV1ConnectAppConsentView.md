# IamV1ConnectAppConsentView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The Connect App's friendly name. |[optional] 
**Company** | **string** | The company that owns the Connect App. |[optional] 
**Description** | Pointer to **string** | A description of the Connect App. |
**Permissions** | **[]string** | The permissions the Connect App asks for. |[optional] 
**PresignedImageUrl** | Pointer to **string** | A short-lived URL to download the app logo, or null if the app has no logo. Do not cache it. |
**AlreadyAuthorized** | **bool** | Whether the authenticated account has already authorized this Connect App. |[optional] 
**RedirectUrl** | Pointer to **string** | Set only when alreadyAuthorized is true. The authorizeRedirectUrl with AccountSid (the Connect subaccount) and state added; redirect the user here instead of showing the consent page. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


