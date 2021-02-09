# UpdateSimRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) to which the Sim resource should belong. The Account SID can only be that of the requesting Account or that of a [Subaccount](https://www.twilio.com/docs/iam/api/subaccounts) of the requesting Account. Only valid when the Sim resource&#39;s status is &#x60;new&#x60;. For more information, see the [Move SIMs between Subaccounts documentation](https://www.twilio.com/docs/wireless/api/sim-resource#move-sims-between-subaccounts). | [optional] 
**CallbackMethod** | **string** | The HTTP method we should use to call &#x60;callback_url&#x60;. Can be: &#x60;POST&#x60; or &#x60;GET&#x60;. The default is &#x60;POST&#x60;. | [optional] 
**CallbackUrl** | **string** | The URL we should call using the &#x60;callback_url&#x60; when the SIM has finished updating. When the SIM transitions from &#x60;new&#x60; to &#x60;ready&#x60; or from any status to &#x60;deactivated&#x60;, we call this URL when the status changes to an intermediate status (&#x60;ready&#x60; or &#x60;deactivated&#x60;) and again when the status changes to its final status (&#x60;active&#x60; or &#x60;canceled&#x60;). | [optional] 
**CommandsCallbackMethod** | **string** | The HTTP method we should use to call &#x60;commands_callback_url&#x60;. Can be: &#x60;POST&#x60; or &#x60;GET&#x60;. The default is &#x60;POST&#x60;. | [optional] 
**CommandsCallbackUrl** | **string** | The URL we should call using the &#x60;commands_callback_method&#x60; when the SIM sends a [Command](https://www.twilio.com/docs/wireless/api/command-resource). Your server should respond with an HTTP status code in the 200 range; any response body is ignored. | [optional] 
**FriendlyName** | **string** | A descriptive string that you create to describe the Sim resource. It does not need to be unique. | [optional] 
**RatePlan** | **string** | The SID or unique name of the [RatePlan resource](https://www.twilio.com/docs/wireless/api/rateplan-resource) to which the Sim resource should be assigned. | [optional] 
**ResetStatus** | **string** | Initiate a connectivity reset on the SIM. Set to &#x60;resetting&#x60; to initiate a connectivity reset on the SIM. No other value is valid. | [optional] 
**SmsFallbackMethod** | **string** | The HTTP method we should use to call &#x60;sms_fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;. Default is &#x60;POST&#x60;. | [optional] 
**SmsFallbackUrl** | **string** | The URL we should call using the &#x60;sms_fallback_method&#x60; when an error occurs while retrieving or executing the TwiML requested from &#x60;sms_url&#x60;. | [optional] 
**SmsMethod** | **string** | The HTTP method we should use to call &#x60;sms_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;. Default is &#x60;POST&#x60;. | [optional] 
**SmsUrl** | **string** | The URL we should call using the &#x60;sms_method&#x60; when the SIM-connected device sends an SMS message that is not a [Command](https://www.twilio.com/docs/wireless/api/command-resource). | [optional] 
**Status** | **string** | The new status of the Sim resource. Can be: &#x60;ready&#x60;, &#x60;active&#x60;, &#x60;suspended&#x60;, or &#x60;deactivated&#x60;. | [optional] 
**UniqueName** | **string** | An application-defined string that uniquely identifies the resource. It can be used in place of the &#x60;sid&#x60; in the URL path to address the resource. | [optional] 
**VoiceFallbackMethod** | **string** | The HTTP method we should use to call &#x60;voice_fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;. | [optional] 
**VoiceFallbackUrl** | **string** | The URL we should call using the &#x60;voice_fallback_method&#x60; when an error occurs while retrieving or executing the TwiML requested from &#x60;voice_url&#x60;. | [optional] 
**VoiceMethod** | **string** | The HTTP method we should use when we call &#x60;voice_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;. | [optional] 
**VoiceUrl** | **string** | The URL we should call using the &#x60;voice_method&#x60; when the SIM-connected device makes a voice call. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


