# ChatV2User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**ServiceSid** | Pointer to **string** | The SID of the Service that the resource is associated with |
**Attributes** | Pointer to **string** | The JSON string that stores application-specific data |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the resource |
**RoleSid** | Pointer to **string** | The SID of the Role assigned to the user |
**Identity** | Pointer to **string** | The string that identifies the resource's User |
**IsOnline** | Pointer to **bool** | Whether the User is actively connected to the Service instance and online |
**IsNotifiable** | Pointer to **bool** | Whether the User has a potentially valid Push Notification registration for the Service instance |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |
**JoinedChannelsCount** | Pointer to **int** | The number of Channels the User is a Member of |
**Links** | Pointer to **map[string]interface{}** | The absolute URLs of the Channel and Binding resources related to the user |
**Url** | Pointer to **string** | The absolute URL of the User resource |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


