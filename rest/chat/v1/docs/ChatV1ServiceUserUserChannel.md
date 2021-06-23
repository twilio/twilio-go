# ChatV1ServiceUserUserChannel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**ChannelSid** | Pointer to **string** | The SID of the Channel the resource belongs to |
**LastConsumedMessageIndex** | Pointer to **int** | The index of the last Message in the Channel the Member has read |
**Links** | Pointer to **map[string]interface{}** | Absolute URLs to access the Members, Messages , Invites and, if it exists, the last Message for the Channel |
**MemberSid** | Pointer to **string** | The SID of the User as a Member in the Channel |
**ServiceSid** | Pointer to **string** | The SID of the Service that the resource is associated with |
**Status** | Pointer to **string** | The status of the User on the Channel |
**UnreadMessagesCount** | Pointer to **int** | The number of unread Messages in the Channel for the User |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


