# ApiV2010Message

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to **string** | The message text |
**NumSegments** | Pointer to **string** | The number of messages used to deliver the message body |
**Direction** | Pointer to [**string**](MessageEnumDirection.md) |  |
**From** | Pointer to **string** | The phone number that initiated the message |
**To** | Pointer to **string** | The phone number that received the message |
**DateUpdated** | Pointer to **string** | The RFC 2822 date and time in GMT that the resource was last updated |
**Price** | Pointer to **string** | The amount billed for the message |
**ErrorMessage** | Pointer to **string** | The description of the error_code |
**Uri** | Pointer to **string** | The URI of the resource, relative to `https://api.twilio.com` |
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**NumMedia** | Pointer to **string** | The number of media files associated with the message |
**Status** | Pointer to [**string**](MessageEnumStatus.md) |  |
**MessagingServiceSid** | Pointer to **string** | The SID of the Messaging Service used with the message. |
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**DateSent** | Pointer to **string** | The RFC 2822 date and time in GMT when the message was sent |
**DateCreated** | Pointer to **string** | The RFC 2822 date and time in GMT that the resource was created |
**ErrorCode** | Pointer to **int** | The error code associated with the message |
**PriceUnit** | Pointer to **string** | The currency in which price is measured |
**ApiVersion** | Pointer to **string** | The API version used to process the message |
**SubresourceUris** | Pointer to **map[string]interface{}** | A list of related resources identified by their relative URIs |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


