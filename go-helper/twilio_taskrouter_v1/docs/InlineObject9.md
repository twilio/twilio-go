# InlineObject9

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssignmentStatus** | **string** | The new status of the task. Can be: &#x60;canceled&#x60;, to cancel a Task that is currently &#x60;pending&#x60; or &#x60;reserved&#x60;; &#x60;wrapping&#x60;, to move the Task to wrapup state; or &#x60;completed&#x60;, to move a Task to the completed state. | [optional] 
**Attributes** | **string** | The JSON string that describes the custom attributes of the task. | [optional] 
**Priority** | **int32** | The Task&#39;s new priority value. When supplied, the Task takes on the specified priority unless it matches a Workflow Target with a Priority set. Value can be 0 to 2^31^ (2,147,483,647). | [optional] 
**Reason** | **string** | The reason that the Task was canceled or completed. This parameter is required only if the Task is canceled or completed. Setting this value queues the task for deletion and logs the reason. | [optional] 
**TaskChannel** | **string** | When MultiTasking is enabled, specify the TaskChannel with the task to update. Can be the TaskChannel&#39;s SID or its &#x60;unique_name&#x60;, such as &#x60;voice&#x60;, &#x60;sms&#x60;, or &#x60;default&#x60;. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


