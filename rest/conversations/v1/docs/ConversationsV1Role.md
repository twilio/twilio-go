# ConversationsV1Role

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | **string** | The unique string that identifies the resource |[optional] 
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**ChatServiceSid** | **string** | The SID of the Conversation Service that the resource is associated with |[optional] 
**FriendlyName** | **string** | The string that you assigned to describe the resource |[optional] 
**Type** | Pointer to [**string**](RoleEnumRoleType.md) |  |
**Permissions** | **[]string** | An array of the permissions the role has been granted |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |[optional] 
**Url** | **string** | An absolute URL for this user role. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


