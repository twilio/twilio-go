# InlineObject10

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigurationFilters** | **[]string** | The events that cause us to call the Channel Webhook. Used when &#x60;type&#x60; is &#x60;webhook&#x60;. This parameter takes only one event. To specify more than one event, repeat this parameter for each event. For the list of possible events, see [Webhook Event Triggers](https://www.twilio.com/docs/chat/webhook-events#webhook-event-trigger). | [optional] 
**ConfigurationFlowSid** | **string** | The SID of the Studio [Flow](https://www.twilio.com/docs/studio/rest-api/flow) to call when an event in &#x60;configuration.filters&#x60; occurs. Used only when &#x60;type&#x60; &#x3D; &#x60;studio&#x60;. | [optional] 
**ConfigurationMethod** | **string** | The HTTP method used to call &#x60;configuration.url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and the default is &#x60;POST&#x60;. | [optional] 
**ConfigurationRetryCount** | **int32** | The number of times to retry the webhook if the first attempt fails. Can be an integer between 0 and 3, inclusive, and the default is 0. | [optional] 
**ConfigurationTriggers** | **[]string** | A string that will cause us to call the webhook when it is present in a message body. This parameter takes only one trigger string. To specify more than one, repeat this parameter for each trigger string up to a total of 5 trigger strings. Used only when &#x60;type&#x60; &#x3D; &#x60;trigger&#x60;. | [optional] 
**ConfigurationUrl** | **string** | The URL of the webhook to call using the &#x60;configuration.method&#x60;. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


