# FlexV1FlexFlow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**ChannelType** | Pointer to **string** | The channel type |
**ChatServiceSid** | Pointer to **string** | The SID of the chat service |
**ContactIdentity** | Pointer to **string** | The channel contact's Identity |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |
**Enabled** | Pointer to **bool** | Whether the Flex Flow is enabled |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the resource |
**Integration** | Pointer to **interface{}** | An object that contains specific parameters for the integration |
**IntegrationType** | Pointer to **string** | The software that will handle inbound messages. |
**JanitorEnabled** | Pointer to **bool** | Remove active Proxy sessions if the corresponding Task is deleted. |
**LongLived** | Pointer to **bool** | Re-use this chat channel for future interactions with a contact |
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**Url** | Pointer to **string** | The absolute URL of the Flex Flow resource |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


