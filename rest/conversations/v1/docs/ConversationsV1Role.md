# ConversationsV1Role

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that we created to identify the Role resource. |
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Role resource. |
**ChatServiceSid** | Pointer to **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Role resource is associated with. |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the resource. |
**Type** | Pointer to [**string**](RoleEnumRoleType.md) |  |
**Permissions** | Pointer to **[]string** | An array of the permissions the role has been granted. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**Url** | Pointer to **string** | An absolute API resource URL for this user role. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


