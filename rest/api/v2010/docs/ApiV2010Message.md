# ApiV2010Message

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**ApiVersion** | Pointer to **string** | The API version used to process the message |
**Body** | Pointer to **string** | The message text |
**DateCreated** | Pointer to **string** | The RFC 2822 date and time in GMT that the resource was created |
**DateSent** | Pointer to **string** | The RFC 2822 date and time in GMT when the message was sent |
**DateUpdated** | Pointer to **string** | The RFC 2822 date and time in GMT that the resource was last updated |
**Direction** | Pointer to **string** | The direction of the message |
**ErrorCode** | Pointer to **int** | The error code associated with the message |
**ErrorMessage** | Pointer to **string** | The description of the error_code |
**From** | Pointer to **string** | The phone number that initiated the message |
**MessagingServiceSid** | Pointer to **string** | The SID of the Messaging Service used with the message. |
**NumMedia** | Pointer to **string** | The number of media files associated with the message |
**NumSegments** | Pointer to **string** | The number of messages used to deliver the message body |
**Price** | Pointer to **string** | The amount billed for the message |
**PriceUnit** | Pointer to **string** | The currency in which price is measured |
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**Status** | Pointer to **string** | The status of the message |
**SubresourceUris** | Pointer to **map[string]interface{}** | A list of related resources identified by their relative URIs |
**To** | Pointer to **string** | The phone number that received the message |
**Uri** | Pointer to **string** | The URI of the resource, relative to `https://api.twilio.com` |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


