# FlexV1FlexFlow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |[optional] 
**Sid** | **string** | The unique string that identifies the resource |[optional] 
**FriendlyName** | **string** | The string that you assigned to describe the resource |[optional] 
**ChatServiceSid** | **string** | The SID of the chat service |[optional] 
**ChannelType** | Pointer to [**string**](FlexFlowEnumChannelType.md) |  |
**ContactIdentity** | **string** | The channel contact's Identity |[optional] 
**Enabled** | **bool** | Whether the Flex Flow is enabled |[optional] 
**IntegrationType** | Pointer to [**string**](FlexFlowEnumIntegrationType.md) |  |
**Integration** | Pointer to **interface{}** | An object that contains specific parameters for the integration |
**LongLived** | **bool** | Re-use this chat channel for future interactions with a contact |[optional] 
**JanitorEnabled** | **bool** | Remove active Proxy sessions if the corresponding Task is deleted. |[optional] 
**Url** | **string** | The absolute URL of the Flex Flow resource |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


