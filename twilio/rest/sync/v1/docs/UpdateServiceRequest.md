# UpdateServiceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AclEnabled** | **bool** | Whether token identities in the Service must be granted access to Sync objects by using the [Permissions](https://www.twilio.com/docs/sync/api/sync-permissions) resource. | [optional] 
**FriendlyName** | **string** | A string that you assign to describe the resource. | [optional] 
**ReachabilityDebouncingEnabled** | **bool** | Whether every &#x60;endpoint_disconnected&#x60; event should occur after a configurable delay. The default is &#x60;false&#x60;, where the &#x60;endpoint_disconnected&#x60; event occurs immediately after disconnection. When &#x60;true&#x60;, intervening reconnections can prevent the &#x60;endpoint_disconnected&#x60; event. | [optional] 
**ReachabilityDebouncingWindow** | **int32** | The reachability event delay in milliseconds if &#x60;reachability_debouncing_enabled&#x60; &#x3D; &#x60;true&#x60;.  Must be between 1,000 and 30,000 and defaults to 5,000. This is the number of milliseconds after the last running client disconnects, and a Sync identity is declared offline, before the webhook is called if all endpoints remain offline. A reconnection from the same identity by any endpoint during this interval prevents the webhook from being called. | [optional] 
**ReachabilityWebhooksEnabled** | **bool** | Whether the service instance should call &#x60;webhook_url&#x60; when client endpoints connect to Sync. The default is &#x60;false&#x60;. | [optional] 
**WebhookUrl** | **string** | The URL we should call when Sync objects are manipulated. | [optional] 
**WebhooksFromRestEnabled** | **bool** | Whether the Service instance should call &#x60;webhook_url&#x60; when the REST API is used to update Sync objects. The default is &#x60;false&#x60;. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


