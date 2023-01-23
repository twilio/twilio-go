# SyncV1Service

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | **string** | The unique string that identifies the resource |[optional] 
**UniqueName** | **string** | An application-defined string that uniquely identifies the resource |[optional] 
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**FriendlyName** | **string** | The string that you assigned to describe the resource |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |[optional] 
**Url** | **string** | The absolute URL of the Service resource |[optional] 
**WebhookUrl** | **string** | The URL we call when Sync objects are manipulated |[optional] 
**WebhooksFromRestEnabled** | **bool** | Whether the Service instance should call webhook_url when the REST API is used to update Sync objects |[optional] 
**ReachabilityWebhooksEnabled** | **bool** | Whether the service instance calls webhook_url when client endpoints connect to Sync |[optional] 
**AclEnabled** | **bool** | Whether token identities in the Service must be granted access to Sync objects by using the Permissions resource |[optional] 
**ReachabilityDebouncingEnabled** | **bool** | Whether every endpoint_disconnected event occurs after a configurable delay |[optional] 
**ReachabilityDebouncingWindow** | **int** | The reachability event delay in milliseconds |[optional] 
**Links** | **map[string]interface{}** | The URLs of related resources |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


