# SyncV1Service

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that we created to identify the Service resource. |
**UniqueName** | Pointer to **string** | An application-defined string that uniquely identifies the resource. It can be used in place of the resource's `sid` in the URL to address the resource. It is a read-only property, it cannot be assigned using REST API. |
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Service resource. |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the resource. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**Url** | Pointer to **string** | The absolute URL of the Service resource. |
**WebhookUrl** | Pointer to **string** | The URL we call when Sync objects are manipulated. |
**WebhooksFromRestEnabled** | Pointer to **bool** | Whether the Service instance should call `webhook_url` when the REST API is used to update Sync objects. The default is `false`. |
**ReachabilityWebhooksEnabled** | Pointer to **bool** | Whether the service instance calls `webhook_url` when client endpoints connect to Sync. The default is `false`. |
**AclEnabled** | Pointer to **bool** | Whether token identities in the Service must be granted access to Sync objects by using the [Permissions](https://www.twilio.com/docs/sync/api/sync-permissions) resource. It is disabled (false) by default. |
**ReachabilityDebouncingEnabled** | Pointer to **bool** | Whether every `endpoint_disconnected` event should occur after a configurable delay. The default is `false`, where the `endpoint_disconnected` event occurs immediately after disconnection. When `true`, intervening reconnections can prevent the `endpoint_disconnected` event. |
**ReachabilityDebouncingWindow** | Pointer to **int** | The reachability event delay in milliseconds if `reachability_debouncing_enabled` = `true`.  Must be between 1,000 and 30,000 and defaults to 5,000. This is the number of milliseconds after the last running client disconnects, and a Sync identity is declared offline, before `webhook_url` is called, if all endpoints remain offline. A reconnection from the same identity by any endpoint during this interval prevents the reachability event from occurring. |
**Links** | Pointer to **map[string]interface{}** | The URLs of related resources. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


