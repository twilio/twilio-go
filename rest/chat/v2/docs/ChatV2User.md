# ChatV2User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**Attributes** | Pointer to **string** | The JSON string that stores application-specific data |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the resource |
**Identity** | Pointer to **string** | The string that identifies the resource's User |
**IsNotifiable** | Pointer to **bool** | Whether the User has a potentially valid Push Notification registration for the Service instance |
**IsOnline** | Pointer to **bool** | Whether the User is actively connected to the Service instance and online |
**JoinedChannelsCount** | Pointer to **int** | The number of Channels the User is a Member of |
**Links** | Pointer to **map[string]interface{}** | The absolute URLs of the Channel and Binding resources related to the user |
**RoleSid** | Pointer to **string** | The SID of the Role assigned to the user |
**ServiceSid** | Pointer to **string** | The SID of the Service that the resource is associated with |
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**Url** | Pointer to **string** | The absolute URL of the User resource |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


