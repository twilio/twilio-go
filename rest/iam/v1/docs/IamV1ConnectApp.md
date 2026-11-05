# IamV1ConnectApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | **string** | The unique identifier of the Connect App. |[optional] 
**AccountSid** | **string** | The account that owns the Connect App. |[optional] 
**Active** | **bool** | Whether the Connect App is active. Inactive apps are never returned by the list endpoint. |[optional] 
**FriendlyName** | Pointer to **string** | The name shown to a user being asked to authorize the Connect App. |
**CompanyName** | **string** | The name of the company that publishes the Connect App. |[optional] 
**Description** | Pointer to **string** | A description of what the Connect App does. |
**HomepageUrl** | **string** | The public homepage of the Connect App. |[optional] 
**AuthorizeRedirectUrl** | **string** | Where the user is sent after allowing or denying authorization. |[optional] 
**DeauthorizeCallbackUrl** | Pointer to **string** | The URL Twilio calls when an account deauthorizes the Connect App. |
**DeauthorizeCallbackMethod** | Pointer to **string** | The HTTP method used for the deauthorize callback. |
**ImageUrl** | Pointer to **string** | An internal storage reference for the app logo, as a relative path. It is not a browser-renderable URL; use presignedImageUrl to display the logo. |
**PresignedImageUrl** | Pointer to **string** | A short-lived URL to download the app logo. Returned only when fetching a single Connect App, and only if the app has a logo. Do not cache it. |
**TosUrl** | Pointer to **string** | The terms of service for the Connect App. |
**Permissions** | Pointer to **[]string** | The permissions the Connect App requests. |
**DateCreated** | [**time.Time**](time.Time.md) | When the Connect App was created. |[optional] 
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | When the Connect App was last updated. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


