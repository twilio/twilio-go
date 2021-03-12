# ConversationsV1ServiceServiceUser

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**Attributes** | Pointer to **string** | The JSON Object string that stores application-specific data |
**ChatServiceSid** | Pointer to **string** | The SID of the Conversation Service that the resource is associated with |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the resource |
**Identity** | Pointer to **string** | The string that identifies the resource's User |
**IsNotifiable** | Pointer to **bool** | Whether the User has a potentially valid Push Notification registration for this Conversations Service |
**IsOnline** | Pointer to **bool** | Whether the User is actively connected to this Conversations Service and online |
**RoleSid** | Pointer to **string** | The SID of a service-level Role assigned to the user |
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**Url** | Pointer to **string** | An absolute URL for this user. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


