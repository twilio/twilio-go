# UpdateUserChannelRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastConsumedMessageIndex** | Pointer to **int32** | The index of the last [Message](https://www.twilio.com/docs/chat/rest/message-resource) in the [Channel](https://www.twilio.com/docs/chat/channels) that the Member has read. | [optional] 
**LastConsumptionTimestamp** | [**time.Time**](time.Time.md) | The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp of the last [Message](https://www.twilio.com/docs/chat/rest/message-resource) read event for the Member within the [Channel](https://www.twilio.com/docs/chat/channels). | [optional] 
**NotificationLevel** | **string** | The push notification level to assign to the User Channel. Can be: &#x60;default&#x60; or &#x60;muted&#x60;. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


