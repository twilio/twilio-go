# FlexV1FlexFlow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Flex Flow resource and owns this Workflow. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**Sid** | Pointer to **string** | The unique string that we created to identify the Flex Flow resource. |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the resource. |
**ChatServiceSid** | Pointer to **string** | The SID of the chat service. |
**ChannelType** | Pointer to [**string**](FlexFlowEnumChannelType.md) |  |
**ContactIdentity** | Pointer to **string** | The channel contact's Identity. |
**Enabled** | Pointer to **bool** | Whether the Flex Flow is enabled. |
**IntegrationType** | Pointer to [**string**](FlexFlowEnumIntegrationType.md) |  |
**Integration** | Pointer to **interface{}** | An object that contains specific parameters for the integration. |
**LongLived** | Pointer to **bool** | When enabled, Flex will keep the chat channel active so that it may be used for subsequent interactions with a contact identity. Defaults to `false`. |
**JanitorEnabled** | Pointer to **bool** | When enabled, the Messaging Channel Janitor will remove active Proxy sessions if the associated Task is deleted outside of the Flex UI. Defaults to `false`. |
**Url** | Pointer to **string** | The absolute URL of the Flex Flow resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


