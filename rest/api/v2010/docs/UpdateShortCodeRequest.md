# UpdateShortCodeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | The API version to use to start a new TwiML session. Can be: &#x60;2010-04-01&#x60; or &#x60;2008-08-01&#x60;. | [optional] 
**FriendlyName** | **string** | A descriptive string that you created to describe this resource. It can be up to 64 characters long. By default, the &#x60;FriendlyName&#x60; is the short code. | [optional] 
**SmsFallbackMethod** | **string** | The HTTP method that we should use to call the &#x60;sms_fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;. | [optional] 
**SmsFallbackUrl** | **string** | The URL that we should call if an error occurs while retrieving or executing the TwiML from &#x60;sms_url&#x60;. | [optional] 
**SmsMethod** | **string** | The HTTP method we should use when calling the &#x60;sms_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;. | [optional] 
**SmsUrl** | **string** | The URL we should call when receiving an incoming SMS message to this short code. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


