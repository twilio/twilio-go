# UpdateSyncMapItemRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionTtl** | **int32** | How long, in seconds, before the Map Item&#39;s parent Sync Map expires (time-to-live) and is deleted.  Can be an integer from 0 to 31,536,000 (1 year). The default value is &#x60;0&#x60;, which means the parent Sync Map does not expire. This parameter can only be used when the Map Item&#39;s &#x60;data&#x60; or &#x60;ttl&#x60; is updated in the same request. The Sync Map will be deleted automatically after it expires, but there can be a delay between the expiration time and the resources&#39;s deletion. | [optional] 
**Data** | [**map[string]interface{}**](.md) | A JSON string that represents an arbitrary, schema-less object that the Map Item stores. Can be up to 16KB in length. | [optional] 
**ItemTtl** | **int32** | How long, in seconds, before the Map Item expires (time-to-live) and is deleted.  Can be an integer from 0 to 31,536,000 (1 year). The default value is &#x60;0&#x60;, which means the Map Item does not expire. The Map Item will be deleted automatically after it expires, but there can be a delay between the expiration time and the resources&#39;s deletion. | [optional] 
**Ttl** | **int32** | An alias for &#x60;item_ttl&#x60;. If both parameters are provided, this value is ignored. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


