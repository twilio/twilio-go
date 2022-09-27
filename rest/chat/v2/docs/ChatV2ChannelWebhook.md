# ChatV2ChannelWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**ServiceSid** | Pointer to **string** | The SID of the Service that the Channel Webhook resource is associated with |
**ChannelSid** | Pointer to **string** | The SID of the Channel the Channel Webhook resource belongs to |
**Type** | Pointer to **string** | The type of webhook |
**Url** | Pointer to **string** | The absolute URL of the Channel Webhook resource |
**Configuration** | Pointer to **interface{}** | The JSON string that describes the configuration object for the channel webhook |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


