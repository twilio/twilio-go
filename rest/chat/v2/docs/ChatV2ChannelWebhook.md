# ChatV2ChannelWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | **string** | The unique string that identifies the resource |[optional] 
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**ServiceSid** | **string** | The SID of the Service that the Channel Webhook resource is associated with |[optional] 
**ChannelSid** | **string** | The SID of the Channel the Channel Webhook resource belongs to |[optional] 
**Type** | **string** | The type of webhook |[optional] 
**Url** | **string** | The absolute URL of the Channel Webhook resource |[optional] 
**Configuration** | Pointer to **interface{}** | The JSON string that describes the configuration object for the channel webhook |
**DateCreated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


