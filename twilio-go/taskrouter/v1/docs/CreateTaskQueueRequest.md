# CreateTaskQueueRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssignmentActivitySid** | **string** | The SID of the Activity to assign Workers when a task is assigned to them. | [optional] 
**FriendlyName** | **string** | A descriptive string that you create to describe the TaskQueue. For example &#x60;Support-Tier 1&#x60;, &#x60;Sales&#x60;, or &#x60;Escalation&#x60;. | 
**MaxReservedWorkers** | **int32** | The maximum number of Workers to reserve for the assignment of a Task in the queue. Can be an integer between 1 and 50, inclusive and defaults to 1. | [optional] 
**ReservationActivitySid** | **string** | The SID of the Activity to assign Workers when a task is reserved for them. | [optional] 
**TargetWorkers** | **string** | A string that describes the Worker selection criteria for any Tasks that enter the TaskQueue. For example, &#x60;&#39;\&quot;language\&quot; &#x3D;&#x3D; \&quot;spanish\&quot;&#39;&#x60;. The default value is &#x60;1&#x3D;&#x3D;1&#x60;. If this value is empty, Tasks will wait in the TaskQueue until they are deleted or moved to another TaskQueue. For more information about Worker selection, see [Describing Worker selection criteria](https://www.twilio.com/docs/taskrouter/api/taskqueues#target-workers). | [optional] 
**TaskOrder** | **string** | How Tasks will be assigned to Workers. Set this parameter to &#x60;LIFO&#x60; to assign most recently created Task first or FIFO to assign the oldest Task first. Default is &#x60;FIFO&#x60;. [Click here](https://www.twilio.com/docs/taskrouter/queue-ordering-last-first-out-lifo) to learn more. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


