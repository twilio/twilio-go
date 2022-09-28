# FlexV1FlexFlow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the resource |
**ChatServiceSid** | Pointer to **string** | The SID of the chat service |
**ChannelType** | Pointer to [**string**](FlexFlowEnumChannelType.md) |  |
**ContactIdentity** | Pointer to **string** | The channel contact's Identity |
**Enabled** | Pointer to **bool** | Whether the Flex Flow is enabled |
**IntegrationType** | Pointer to [**string**](FlexFlowEnumIntegrationType.md) |  |
**Integration** | Pointer to **interface{}** | An object that contains specific parameters for the integration |
**LongLived** | Pointer to **bool** | Re-use this chat channel for future interactions with a contact |
**JanitorEnabled** | Pointer to **bool** | Remove active Proxy sessions if the corresponding Task is deleted. |
**Url** | Pointer to **string** | The absolute URL of the Flex Flow resource |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


