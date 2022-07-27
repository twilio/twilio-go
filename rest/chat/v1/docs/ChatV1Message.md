# ChatV1Message

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**Attributes** | Pointer to **string** | The JSON string that stores application-specific data |
**ServiceSid** | Pointer to **string** | The SID of the Service that the resource is associated with |
**To** | Pointer to **string** | The SID of the Channel that the message was sent to |
**ChannelSid** | Pointer to **string** | The unique ID of the Channel the Message resource belongs to |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was last updated |
**WasEdited** | Pointer to **bool** | Whether the message has been edited since  it was created |
**From** | Pointer to **string** | The identity of the message's author |
**Body** | Pointer to **string** | The content of the message |
**Index** | Pointer to **int** | The index of the message within the Channel |
**Url** | Pointer to **string** | The absolute URL of the Message resource |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


