# ChatV1Role

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that we created to identify the Role resource. |
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/api/rest/account) that created the Role resource. |
**ServiceSid** | Pointer to **string** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) the resource is associated with. |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the resource. |
**Type** | Pointer to [**string**](RoleEnumRoleType.md) |  |
**Permissions** | Pointer to **[]string** | An array of the permissions the role has been granted, formatted as a JSON string. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was created specified in [RFC 2822](http://www.ietf.org/rfc/rfc2822.txt) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was last updated specified in [RFC 2822](http://www.ietf.org/rfc/rfc2822.txt) format. |
**Url** | Pointer to **string** | The absolute URL of the Role resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


