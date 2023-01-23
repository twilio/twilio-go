# ChatV1Message

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | **string** | The unique string that identifies the resource |[optional] 
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**Attributes** | **string** | The JSON string that stores application-specific data |[optional] 
**ServiceSid** | **string** | The SID of the Service that the resource is associated with |[optional] 
**To** | **string** | The SID of the Channel that the message was sent to |[optional] 
**ChannelSid** | **string** | The unique ID of the Channel the Message resource belongs to |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was created |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was last updated |[optional] 
**WasEdited** | **bool** | Whether the message has been edited since  it was created |[optional] 
**From** | **string** | The identity of the message's author |[optional] 
**Body** | **string** | The content of the message |[optional] 
**Index** | **int** | The index of the message within the Channel |[optional] 
**Url** | **string** | The absolute URL of the Message resource |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


