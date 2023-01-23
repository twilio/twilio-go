# ConversationsV1User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | **string** | The unique string that identifies the resource |[optional] 
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**ChatServiceSid** | **string** | The SID of the Conversation Service that the resource is associated with |[optional] 
**RoleSid** | **string** | The SID of a service-level Role assigned to the user |[optional] 
**Identity** | **string** | The string that identifies the resource's User |[optional] 
**FriendlyName** | **string** | The string that you assigned to describe the resource |[optional] 
**Attributes** | **string** | The JSON Object string that stores application-specific data |[optional] 
**IsOnline** | **bool** | Whether the User is actively connected to this Conversations Service and online |[optional] 
**IsNotifiable** | **bool** | Whether the User has a potentially valid Push Notification registration for this Conversations Service |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |[optional] 
**Url** | **string** | An absolute URL for this user. |[optional] 
**Links** | **map[string]interface{}** |  |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


