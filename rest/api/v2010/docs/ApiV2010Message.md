# ApiV2010Message

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | **string** | The message text |[optional] 
**NumSegments** | **string** | The number of messages used to deliver the message body |[optional] 
**Direction** | Pointer to [**string**](MessageEnumDirection.md) |  |
**From** | **string** | The phone number that initiated the message |[optional] 
**To** | **string** | The phone number that received the message |[optional] 
**DateUpdated** | **string** | The RFC 2822 date and time in GMT that the resource was last updated |[optional] 
**Price** | **string** | The amount billed for the message |[optional] 
**ErrorMessage** | **string** | The description of the error_code |[optional] 
**Uri** | **string** | The URI of the resource, relative to `https://api.twilio.com` |[optional] 
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**NumMedia** | **string** | The number of media files associated with the message |[optional] 
**Status** | Pointer to [**string**](MessageEnumStatus.md) |  |
**MessagingServiceSid** | **string** | The SID of the Messaging Service used with the message. |[optional] 
**Sid** | **string** | The unique string that identifies the resource |[optional] 
**DateSent** | **string** | The RFC 2822 date and time in GMT when the message was sent |[optional] 
**DateCreated** | **string** | The RFC 2822 date and time in GMT that the resource was created |[optional] 
**ErrorCode** | Pointer to **int** | The error code associated with the message |
**PriceUnit** | **string** | The currency in which price is measured |[optional] 
**ApiVersion** | **string** | The API version used to process the message |[optional] 
**SubresourceUris** | **map[string]interface{}** | A list of related resources identified by their relative URIs |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


