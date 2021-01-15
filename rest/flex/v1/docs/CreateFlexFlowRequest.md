# CreateFlexFlowRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelType** | **string** | The channel type. Can be: &#x60;web&#x60;, &#x60;facebook&#x60;, &#x60;sms&#x60;, &#x60;whatsapp&#x60;, &#x60;line&#x60; or &#x60;custom&#x60;. | 
**ChatServiceSid** | **string** | The SID of the chat service. | 
**ContactIdentity** | **string** | The channel contact&#39;s Identity. | [optional] 
**Enabled** | **bool** | Whether the new Flex Flow is enabled. | [optional] 
**FriendlyName** | **string** | A descriptive string that you create to describe the Flex Flow resource. | 
**IntegrationChannel** | **string** | The Task Channel for the TaskRouter Task that will be created. Applicable and required when integrationType is &#x60;task&#x60;. Set to &#x60;sms&#x60; for SMS, and to &#x60;chat&#x60; otherwise. The default value is &#x60;default&#x60; | [optional] 
**IntegrationCreationOnMessage** | **bool** | In the context of outbound messaging, defines whether to create a Task immediately (and therefore reserve the conversation to current agent), or delay Task creation until the customer sends the first response. Set to false to create immediately, true to delay Task creation. This setting is only applicable for outbound messaging. | [optional] 
**IntegrationFlowSid** | **string** | The SID of the Studio Flow. Required when &#x60;integrationType&#x60; is &#x60;studio&#x60;. | [optional] 
**IntegrationPriority** | **int32** | The Task priority of a new Task. The default priority is 0. Optional when &#x60;integrationType&#x60; is &#x60;task&#x60;, not applicable otherwise. | [optional] 
**IntegrationRetryCount** | **int32** | The number of times to retry the webhook if the first attempt fails. Can be an integer between 0 and 3 (included), default is 0. Optional when integrationType is &#x60;external&#x60;, not applicable otherwise. | [optional] 
**IntegrationTimeout** | **int32** | The Task timeout in seconds for a new Task. Default is 86,400 seconds (24 hours). Optional when &#x60;integrationType&#x60; is &#x60;task&#x60;, not applicable otherwise. | [optional] 
**IntegrationUrl** | **string** | The URL of the external webhook. Required when &#x60;integrationType&#x60; is &#x60;external&#x60;. | [optional] 
**IntegrationWorkflowSid** | **string** | The Workflow SID for a new Task. Required when &#x60;integrationType&#x60; is &#x60;task&#x60;. | [optional] 
**IntegrationWorkspaceSid** | **string** | The Workspace SID for a new Task. Required when &#x60;integrationType&#x60; is &#x60;task&#x60;. | [optional] 
**IntegrationType** | **string** | The integration type. Can be: &#x60;studio&#x60;, &#x60;external&#x60;, or &#x60;task&#x60;. | [optional] 
**JanitorEnabled** | **bool** | When enabled, the Messaging Channel Janitor will remove active Proxy sessions if the associated Task is deleted outside of the Flex UI. Defaults to &#x60;false&#x60;. | [optional] 
**LongLived** | **bool** | When enabled, Flex will keep the chat channel active so that it may be used for subsequent interactions with a contact identity. Defaults to &#x60;false&#x60;. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


