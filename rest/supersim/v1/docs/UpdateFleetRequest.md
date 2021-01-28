# UpdateFleetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommandsMethod** | **string** | A string representing the HTTP method to use when making a request to &#x60;commands_url&#x60;. Can be one of &#x60;POST&#x60; or &#x60;GET&#x60;. Defaults to &#x60;POST&#x60;. | [optional] 
**CommandsUrl** | **string** | The URL that will receive a webhook when a SIM in the Fleet is used to send an SMS from your device (mobile originated) to the Commands number. Your server should respond with an HTTP status code in the 200 range; any response body will be ignored. | [optional] 
**NetworkAccessProfile** | **string** | The SID or unique name of the Network Access Profile that will control which cellular networks the Fleet&#39;s SIMs can connect to. | [optional] 
**UniqueName** | **string** | An application-defined string that uniquely identifies the resource. It can be used in place of the resource&#39;s &#x60;sid&#x60; in the URL to address the resource. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


