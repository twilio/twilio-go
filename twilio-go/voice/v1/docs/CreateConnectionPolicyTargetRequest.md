# CreateConnectionPolicyTargetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Whether the Target is enabled. The default is &#x60;true&#x60;. | [optional] 
**FriendlyName** | **string** | A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long. | [optional] 
**Priority** | **int32** | The relative importance of the target. Can be an integer from 0 to 65535, inclusive, and the default is 10. The lowest number represents the most important target. | [optional] 
**Target** | **string** | The SIP address you want Twilio to route your calls to. This must be a &#x60;sip:&#x60; schema. &#x60;sips&#x60; is NOT supported. | 
**Weight** | **int32** | The value that determines the relative share of the load the Target should receive compared to other Targets with the same priority. Can be an integer from 1 to 65535, inclusive, and the default is 10. Targets with higher values receive more load than those with lower ones with the same priority. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


