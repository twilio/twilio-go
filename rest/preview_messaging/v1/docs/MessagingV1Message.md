# MessagingV1Message

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**To** | **string** | The destination phone number in [E.164](https://www.twilio.com/docs/glossary/what-e164) format for SMS/MMS or [Channel user address](https://www.twilio.com/docs/sms/channels#channel-addresses) for other 3rd-party channels. |[optional] 
**Body** | **string** | The text of the message you want to send. Can be up to 1,600 characters in length. Overrides the request-level body and content template if provided. |[optional] 
**ContentVariables** | **map[string]string** | Key-value pairs of variable names to substitution values. Refer to the [Twilio Content API Resources](https://www.twilio.com/docs/content-api/content-api-resources#send-a-message-with-preconfigured-content) for more details. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


