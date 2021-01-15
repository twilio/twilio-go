# CreateDocumentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**map[string]interface{}**](.md) | A JSON string that represents an arbitrary, schema-less object that the Sync Document stores. Can be up to 16KB in length. | [optional] 
**Ttl** | **int32** | How long, in seconds, before the Sync Document expires and is deleted (the Sync Document&#39;s time-to-live). Can be an integer from 0 to 31,536,000 (1 year). The default value is &#x60;0&#x60;, which means the Sync Document does not expire. The Sync Document will be deleted automatically after it expires, but there can be a delay between the expiration time and the resources&#39;s deletion. | [optional] 
**UniqueName** | **string** | An application-defined string that uniquely identifies the Sync Document | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


