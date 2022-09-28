# ChatV2Channel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**ServiceSid** | Pointer to **string** | The SID of the Service that the resource is associated with |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the resource |
**UniqueName** | Pointer to **string** | An application-defined string that uniquely identifies the resource |
**Attributes** | Pointer to **string** | The JSON string that stores application-specific data |
**Type** | Pointer to [**string**](ChannelEnumChannelType.md) |  |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |
**CreatedBy** | Pointer to **string** | The identity of the User that created the channel |
**MembersCount** | Pointer to **int** | The number of Members in the Channel |
**MessagesCount** | Pointer to **int** | The number of Messages that have been passed in the Channel |
**Url** | Pointer to **string** | The absolute URL of the Channel resource |
**Links** | Pointer to **map[string]interface{}** | Absolute URLs to access the Members, Messages , Invites and, if it exists, the last Message for the Channel |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


