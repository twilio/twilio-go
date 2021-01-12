# InlineObject22

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallDelay** | **int32** | The number of seconds to delay before initiating the verification call. Can be an integer between &#x60;0&#x60; and &#x60;60&#x60;, inclusive. The default is &#x60;0&#x60;. | [optional] 
**Extension** | **string** | The digits to dial after connecting the verification call. | [optional] 
**FriendlyName** | **string** | A descriptive string that you create to describe the new caller ID resource. It can be up to 64 characters long. The default value is a formatted version of the phone number. | [optional] 
**PhoneNumber** | **string** | The phone number to verify in [E.164](https://www.twilio.com/docs/glossary/what-e164) format, which consists of a + followed by the country code and subscriber number. | 
**StatusCallback** | **string** | The URL we should call using the &#x60;status_callback_method&#x60; to send status information about the verification process to your application. | [optional] 
**StatusCallbackMethod** | **string** | The HTTP method we should use to call &#x60;status_callback&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;, and the default is &#x60;POST&#x60;. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


