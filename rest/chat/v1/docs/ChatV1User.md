# ChatV1User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | **string** | The unique string that identifies the resource |[optional] 
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**ServiceSid** | **string** | The SID of the Service that the resource is associated with |[optional] 
**Attributes** | **string** | The JSON string that stores application-specific data |[optional] 
**FriendlyName** | **string** | The string that you assigned to describe the resource |[optional] 
**RoleSid** | **string** | The SID of the assigned to the user |[optional] 
**Identity** | **string** | The string that identifies the resource's User |[optional] 
**IsOnline** | **bool** | Whether the User is actively connected to the Service instance and online |[optional] 
**IsNotifiable** | **bool** | Whether the User has a potentially valid Push Notification registration for the Service instance |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was created |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was last updated |[optional] 
**JoinedChannelsCount** | **int** | The number of Channels this User is a Member of |[optional] 
**Links** | **map[string]interface{}** | The absolute URLs of the Channel and Binding resources related to the user |[optional] 
**Url** | **string** | The absolute URL of the User resource |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


