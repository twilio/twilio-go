# InlineObject6

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | **string** | The SID of the Account to which the Sim resource should belong. The Account SID can only be that of the requesting Account or that of a Subaccount of the requesting Account. Only valid when the Sim resource&#39;s status is new. | [optional] 
**CallbackMethod** | **string** | The HTTP method we should use to call &#x60;callback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and the default is POST. | [optional] 
**CallbackUrl** | **string** | The URL we should call using the &#x60;callback_method&#x60; after an asynchronous update has finished. | [optional] 
**Fleet** | **string** | The SID or unique name of the Fleet to which the SIM resource should be assigned. | [optional] 
**Status** | **string** | The new status of the resource. Can be: &#x60;ready&#x60;, &#x60;active&#x60;, or &#x60;inactive&#x60;. See the [Super SIM Status Values](https://www.twilio.com/docs/iot/supersim/api/sim-resource#status-values) for more info. | [optional] 
**UniqueName** | **string** | An application-defined string that uniquely identifies the resource. It can be used in place of the resource&#39;s &#x60;sid&#x60; in the URL to address the resource. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


