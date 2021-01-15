# UpdateMessageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | **string** | A valid JSON string that contains application-specific data. | [optional] 
**Body** | **string** | The message to send to the channel. Can be an empty string or &#x60;null&#x60;, which sets the value as an empty string. You can send structured data in the body by serializing it as a string. | [optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The date, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format, to assign to the resource as the date it was created. The default value is the current time set by the Chat service. This parameter should only be used when a Chat&#39;s history is being recreated from a backup/separate source. | [optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The date, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format, to assign to the resource as the date it was last updated. | [optional] 
**From** | **string** | The [Identity](https://www.twilio.com/docs/chat/identity) of the message&#39;s author. | [optional] 
**LastUpdatedBy** | **string** | The [Identity](https://www.twilio.com/docs/chat/identity) of the User who last updated the Message, if applicable. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


