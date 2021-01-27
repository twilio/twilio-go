# CreateFleetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommandsEnabled** | **bool** | Defines whether SIMs in the Fleet are capable of sending and receiving machine-to-machine SMS via Commands. Defaults to &#x60;true&#x60;. | [optional] 
**CommandsMethod** | **string** | A string representing the HTTP method to use when making a request to &#x60;commands_url&#x60;. Can be one of &#x60;POST&#x60; or &#x60;GET&#x60;. Defaults to &#x60;POST&#x60;. | [optional] 
**CommandsUrl** | **string** | The URL that will receive a webhook when a SIM in the Fleet is used to send an SMS from your device (mobile originated) to the Commands number. Your server should respond with an HTTP status code in the 200 range; any response body will be ignored. | [optional] 
**DataEnabled** | **bool** | Defines whether SIMs in the Fleet are capable of using 2G/3G/4G/LTE/CAT-M data connectivity. Defaults to &#x60;true&#x60;. | [optional] 
**DataLimit** | **int32** | The total data usage (download and upload combined) in Megabytes that each Sim resource assigned to the Fleet resource can consume during a billing period (normally one month). Value must be between 1MB (1) and 2TB (2,000,000). Defaults to 1GB (1,000). | [optional] 
**NetworkAccessProfile** | **string** | The SID or unique name of the Network Access Profile that will control which cellular networks the Fleet&#39;s SIMs can connect to. | 
**UniqueName** | **string** | An application-defined string that uniquely identifies the resource. It can be used in place of the resource&#39;s &#x60;sid&#x60; in the URL to address the resource. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


