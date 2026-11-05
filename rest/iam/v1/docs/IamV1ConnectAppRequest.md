# IamV1ConnectAppRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FriendlyName** | **string** | The name of the Connect App. |
**CompanyName** | **string** | The company that owns the Connect App. |
**Description** | Pointer to **string** | A description of the Connect App. |
**HomepageUrl** | **string** | The public homepage of the Connect App. |
**AuthorizeRedirectUrl** | **string** | Where the user is sent after authorizing the Connect App. |
**DeauthorizeCallbackUrl** | Pointer to **string** | Called when an account deauthorizes the Connect App. |
**DeauthorizeCallbackMethod** | Pointer to **string** | The HTTP method for deauthorizeCallbackUrl. Defaults to POST. |
**TosUrl** | Pointer to **string** | The Connect App's terms of service. |
**Permissions** | **[]string** | The permissions the Connect App asks for. |
**ImageKey** | Pointer to **string** | The imageKey from POST /v1/Account/ConnectApps/image/presigned-url, after the image was uploaded. Omit on update to keep the current image. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


