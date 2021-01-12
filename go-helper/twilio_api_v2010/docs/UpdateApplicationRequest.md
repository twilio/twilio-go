# UpdateApplicationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | The API version to use to start a new TwiML session. Can be: &#x60;2010-04-01&#x60; or &#x60;2008-08-01&#x60;. The default value is your account&#39;s default API version. | [optional] 
**FriendlyName** | **string** | A descriptive string that you create to describe the resource. It can be up to 64 characters long. | [optional] 
**MessageStatusCallback** | **string** | The URL we should call using a POST method to send message status information to your application. | [optional] 
**SmsFallbackMethod** | **string** | The HTTP method we should use to call &#x60;sms_fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;. | [optional] 
**SmsFallbackUrl** | **string** | The URL that we should call when an error occurs while retrieving or executing the TwiML from &#x60;sms_url&#x60;. | [optional] 
**SmsMethod** | **string** | The HTTP method we should use to call &#x60;sms_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;. | [optional] 
**SmsStatusCallback** | **string** | Same as message_status_callback: The URL we should call using a POST method to send status information about SMS messages sent by the application. Deprecated, included for backwards compatibility. | [optional] 
**SmsUrl** | **string** | The URL we should call when the phone number receives an incoming SMS message. | [optional] 
**StatusCallback** | **string** | The URL we should call using the &#x60;status_callback_method&#x60; to send status information to your application. | [optional] 
**StatusCallbackMethod** | **string** | The HTTP method we should use to call &#x60;status_callback&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;. | [optional] 
**VoiceCallerIdLookup** | **bool** | Whether we should look up the caller&#39;s caller-ID name from the CNAM database (additional charges apply). Can be: &#x60;true&#x60; or &#x60;false&#x60;. | [optional] 
**VoiceFallbackMethod** | **string** | The HTTP method we should use to call &#x60;voice_fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;. | [optional] 
**VoiceFallbackUrl** | **string** | The URL that we should call when an error occurs retrieving or executing the TwiML requested by &#x60;url&#x60;. | [optional] 
**VoiceMethod** | **string** | The HTTP method we should use to call &#x60;voice_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;. | [optional] 
**VoiceUrl** | **string** | The URL we should call when the phone number assigned to this application receives a call. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


