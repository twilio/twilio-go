# InlineObject1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filters** | **[]string** | The list of webhook event triggers that are enabled for this Service: &#x60;onMessageAdded&#x60;, &#x60;onMessageUpdated&#x60;, &#x60;onMessageRemoved&#x60;, &#x60;onConversationUpdated&#x60;, &#x60;onConversationRemoved&#x60;, &#x60;onParticipantAdded&#x60;, &#x60;onParticipantUpdated&#x60;, &#x60;onParticipantRemoved&#x60; | [optional] 
**Method** | **string** | The HTTP method to be used when sending a webhook request. | [optional] 
**PostWebhookUrl** | **string** | The absolute url the post-event webhook request should be sent to. | [optional] 
**PreWebhookUrl** | **string** | The absolute url the pre-event webhook request should be sent to. | [optional] 
**Target** | **string** | The routing target of the webhook. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


