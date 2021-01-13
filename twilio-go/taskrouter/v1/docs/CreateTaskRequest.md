# CreateTaskRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | **string** | A URL-encoded JSON string with the attributes of the new task. This value is passed to the Workflow&#39;s &#x60;assignment_callback_url&#x60; when the Task is assigned to a Worker. For example: &#x60;{ \&quot;task_type\&quot;: \&quot;call\&quot;, \&quot;twilio_call_sid\&quot;: \&quot;CAxxx\&quot;, \&quot;customer_ticket_number\&quot;: \&quot;12345\&quot; }&#x60;. | [optional] 
**Priority** | **int32** | The priority to assign the new task and override the default. When supplied, the new Task will have this priority unless it matches a Workflow Target with a Priority set. When not supplied, the new Task will have the priority of the matching Workflow Target. Value can be 0 to 2^31^ (2,147,483,647). | [optional] 
**TaskChannel** | **string** | When MultiTasking is enabled, specify the TaskChannel by passing either its &#x60;unique_name&#x60; or &#x60;sid&#x60;. Default value is &#x60;default&#x60;. | [optional] 
**Timeout** | **int32** | The amount of time in seconds the new task is allowed to live. Can be up to a maximum of 2 weeks (1,209,600 seconds). The default value is 24 hours (86,400 seconds). On timeout, the &#x60;task.canceled&#x60; event will fire with description &#x60;Task TTL Exceeded&#x60;. | [optional] 
**WorkflowSid** | **string** | The SID of the Workflow that you would like to handle routing for the new Task. If there is only one Workflow defined for the Workspace that you are posting the new task to, this parameter is optional. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


