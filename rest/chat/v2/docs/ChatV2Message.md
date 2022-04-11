# ChatV2Message

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**Attributes** | Pointer to **string** | The JSON string that stores application-specific data |
**Body** | Pointer to **string** | The content of the message |
**ChannelSid** | Pointer to **string** | The SID of the Channel the Message resource belongs to |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was last updated |
**From** | Pointer to **string** | The Identity of the message's author |
**Index** | Pointer to **int** | The index of the message within the Channel |
**LastUpdatedBy** | Pointer to **string** | The Identity of the User who last updated the Message |
**Media** | Pointer to **interface{}** | A Media object that describes the Message's media if attached; otherwise, null |
**ServiceSid** | Pointer to **string** | The SID of the Service that the resource is associated with |
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**To** | Pointer to **string** | The SID of the Channel that the message was sent to |
**Type** | Pointer to **string** | The Message type |
**Url** | Pointer to **string** | The absolute URL of the Message resource |
**WasEdited** | Pointer to **bool** | Whether the message has been edited since  it was created |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


