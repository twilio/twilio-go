# FrontlineV1User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that we created to identify the User resource. |
**Identity** | Pointer to **string** | The application-defined string that uniquely identifies the resource's User. This value is often a username or an email address, and is case-sensitive. |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the User. |
**Avatar** | Pointer to **string** | The avatar URL which will be shown in Frontline application. |
**State** | Pointer to [**string**](UserEnumStateType.md) |  |
**IsAvailable** | Pointer to **bool** | Whether the User is available for new conversations. Defaults to `false` for new users. |
**Url** | Pointer to **string** | An absolute API resource URL for this user. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


