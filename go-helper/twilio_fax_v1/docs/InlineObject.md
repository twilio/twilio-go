# InlineObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | **string** | The number the fax was sent from. Can be the phone number in [E.164](https://www.twilio.com/docs/glossary/what-e164) format or the SIP &#x60;from&#x60; value. The caller ID displayed to the recipient uses this value. If this is a phone number, it must be a Twilio number or a verified outgoing caller id from your account. If &#x60;to&#x60; is a SIP address, this can be any alphanumeric string (and also the characters &#x60;+&#x60;, &#x60;_&#x60;, &#x60;.&#x60;, and &#x60;-&#x60;), which will be used in the &#x60;from&#x60; header of the SIP request. | [optional] 
**MediaUrl** | **string** | The URL of the PDF that contains the fax. See our [security](https://www.twilio.com/docs/usage/security) page for information on how to ensure the request for your media comes from Twilio. | 
**Quality** | **string** | The [Fax Quality value](https://www.twilio.com/docs/fax/api/fax-resource#fax-quality-values) that describes the fax quality. Can be: &#x60;standard&#x60;, &#x60;fine&#x60;, or &#x60;superfine&#x60; and defaults to &#x60;fine&#x60;. | [optional] 
**SipAuthPassword** | **string** | The password to use with &#x60;sip_auth_username&#x60; to authenticate faxes sent to a SIP address. | [optional] 
**SipAuthUsername** | **string** | The username to use with the &#x60;sip_auth_password&#x60; to authenticate faxes sent to a SIP address. | [optional] 
**StatusCallback** | **string** | The URL we should call using the &#x60;POST&#x60; method to send [status information](https://www.twilio.com/docs/fax/api/fax-resource#fax-status-callback) to your application when the status of the fax changes. | [optional] 
**StoreMedia** | **bool** | Whether to store a copy of the sent media on our servers for later retrieval. Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;true&#x60;. | [optional] 
**To** | **string** | The phone number to receive the fax in [E.164](https://www.twilio.com/docs/glossary/what-e164) format or the recipient&#39;s SIP URI. | 
**Ttl** | **int32** | How long in minutes from when the fax is initiated that we should try to send the fax. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


