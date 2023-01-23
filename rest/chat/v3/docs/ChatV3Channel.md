# ChatV3Channel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | **string** | The unique string that identifies the resource |[optional] 
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**ServiceSid** | **string** | The SID of the Service that the resource is associated with |[optional] 
**FriendlyName** | **string** | The string that you assigned to describe the resource |[optional] 
**UniqueName** | **string** | An application-defined string that uniquely identifies the resource |[optional] 
**Attributes** | **string** | The JSON string that stores application-specific data |[optional] 
**Type** | Pointer to [**string**](ChannelEnumChannelType.md) |  |
**DateCreated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |[optional] 
**CreatedBy** | **string** | The identity of the User that created the channel |[optional] 
**MembersCount** | **int** | The number of Members in the Channel |[optional] 
**MessagesCount** | **int** | The number of Messages that have been passed in the Channel |[optional] 
**MessagingServiceSid** | **string** | The unique ID of the Messaging Service this channel belongs to. |[optional] 
**Url** | **string** | The absolute URL of the Channel resource |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


