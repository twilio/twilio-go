# CreateCommandRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallbackMethod** | **string** | The HTTP method we use to call &#x60;callback_url&#x60;. Can be: &#x60;POST&#x60; or &#x60;GET&#x60;, and the default is &#x60;POST&#x60;. | [optional] 
**CallbackUrl** | **string** | The URL we call using the &#x60;callback_url&#x60; when the Command has finished sending, whether the command was delivered or it failed. | [optional] 
**Command** | **string** | The message body of the Command. Can be plain text in text mode or a Base64 encoded byte string in binary mode. | 
**CommandMode** | **string** | The mode to use when sending the SMS message. Can be: &#x60;text&#x60; or &#x60;binary&#x60;. The default SMS mode is &#x60;text&#x60;. | [optional] 
**DeliveryReceiptRequested** | **bool** | Whether to request delivery receipt from the recipient. For Commands that request delivery receipt, the Command state transitions to &#39;delivered&#39; once the server has received a delivery receipt from the device. The default value is &#x60;true&#x60;. | [optional] 
**IncludeSid** | **string** | Whether to include the SID of the command in the message body. Can be: &#x60;none&#x60;, &#x60;start&#x60;, or &#x60;end&#x60;, and the default behavior is &#x60;none&#x60;. When sending a Command to a SIM in text mode, we can automatically include the SID of the Command in the message body, which could be used to ensure that the device does not process the same Command more than once.  A value of &#x60;start&#x60; will prepend the message with the Command SID, and &#x60;end&#x60; will append it to the end, separating the Command SID from the message body with a space. The length of the Command SID is included in the 160 character limit so the SMS body must be 128 characters or less before the Command SID is included. | [optional] 
**Sim** | **string** | The &#x60;sid&#x60; or &#x60;unique_name&#x60; of the [SIM](https://www.twilio.com/docs/wireless/api/sim-resource) to send the Command to. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


