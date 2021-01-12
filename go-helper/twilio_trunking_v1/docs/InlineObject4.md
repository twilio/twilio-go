# InlineObject4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Whether the URL is enabled. The default is &#x60;true&#x60;. | 
**FriendlyName** | **string** | A descriptive string that you create to describe the resource. It can be up to 64 characters long. | 
**Priority** | **int32** | The relative importance of the URI. Can be an integer from 0 to 65535, inclusive, and the default is 10. The lowest number represents the most important URI. | 
**SipUrl** | **string** | The SIP address you want Twilio to route your Origination calls to. This must be a &#x60;sip:&#x60; schema. | 
**Weight** | **int32** | The value that determines the relative share of the load the URI should receive compared to other URIs with the same priority. Can be an integer from 1 to 65535, inclusive, and the default is 10. URLs with higher values receive more load than those with lower ones with the same priority. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


