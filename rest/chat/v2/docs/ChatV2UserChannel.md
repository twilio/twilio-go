# ChatV2UserChannel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**ServiceSid** | **string** | The SID of the Service that the resource is associated with |[optional] 
**ChannelSid** | **string** | The SID of the Channel the resource belongs to |[optional] 
**UserSid** | **string** | The SID of the User the User Channel belongs to |[optional] 
**MemberSid** | **string** | The SID of the User as a Member in the Channel |[optional] 
**Status** | Pointer to [**string**](UserChannelEnumChannelStatus.md) |  |
**LastConsumedMessageIndex** | Pointer to **int** | The index of the last Message in the Channel the Member has read |
**UnreadMessagesCount** | Pointer to **int** | The number of unread Messages in the Channel for the User |
**Links** | **map[string]interface{}** | Absolute URLs to access the Members, Messages , Invites and, if it exists, the last Message for the Channel |[optional] 
**Url** | **string** | The absolute URL of the resource |[optional] 
**NotificationLevel** | Pointer to [**string**](UserChannelEnumNotificationLevel.md) |  |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


