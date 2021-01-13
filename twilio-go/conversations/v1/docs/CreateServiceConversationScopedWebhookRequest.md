# CreateServiceConversationScopedWebhookRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigurationFilters** | **[]string** | The list of events, firing webhook event for this Conversation. | [optional] 
**ConfigurationFlowSid** | **string** | The studio flow SID, where the webhook should be sent to. | [optional] 
**ConfigurationMethod** | **string** | The HTTP method to be used when sending a webhook request. | [optional] 
**ConfigurationReplayAfter** | **int32** | The message index for which and it&#39;s successors the webhook will be replayed. Not set by default | [optional] 
**ConfigurationTriggers** | **[]string** | The list of keywords, firing webhook event for this Conversation. | [optional] 
**ConfigurationUrl** | **string** | The absolute url the webhook request should be sent to. | [optional] 
**Target** | **string** | The target of this webhook: &#x60;webhook&#x60;, &#x60;studio&#x60;, &#x60;trigger&#x60; | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


