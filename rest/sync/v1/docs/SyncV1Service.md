# SyncV1Service

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**UniqueName** | Pointer to **string** | An application-defined string that uniquely identifies the resource |
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the resource |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |
**Url** | Pointer to **string** | The absolute URL of the Service resource |
**WebhookUrl** | Pointer to **string** | The URL we call when Sync objects are manipulated |
**WebhooksFromRestEnabled** | Pointer to **bool** | Whether the Service instance should call webhook_url when the REST API is used to update Sync objects |
**ReachabilityWebhooksEnabled** | Pointer to **bool** | Whether the service instance calls webhook_url when client endpoints connect to Sync |
**AclEnabled** | Pointer to **bool** | Whether token identities in the Service must be granted access to Sync objects by using the Permissions resource |
**ReachabilityDebouncingEnabled** | Pointer to **bool** | Whether every endpoint_disconnected event occurs after a configurable delay |
**ReachabilityDebouncingWindow** | Pointer to **int** | The reachability event delay in milliseconds |
**Links** | Pointer to **map[string]interface{}** | The URLs of related resources |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


