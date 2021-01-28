# UpdateTaskQueueRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssignmentActivitySid** | **string** | The SID of the Activity to assign Workers when a task is assigned for them. | [optional] 
**FriendlyName** | **string** | A descriptive string that you create to describe the TaskQueue. For example &#x60;Support-Tier 1&#x60;, &#x60;Sales&#x60;, or &#x60;Escalation&#x60;. | [optional] 
**MaxReservedWorkers** | **int32** | The maximum number of Workers to create reservations for the assignment of a task while in the queue. Maximum of 50. | [optional] 
**ReservationActivitySid** | **string** | The SID of the Activity to assign Workers when a task is reserved for them. | [optional] 
**TargetWorkers** | **string** | A string describing the Worker selection criteria for any Tasks that enter the TaskQueue. For example &#39;\&quot;language\&quot; &#x3D;&#x3D; \&quot;spanish\&quot;&#39; If no TargetWorkers parameter is provided, Tasks will wait in the queue until they are either deleted or moved to another queue. Additional examples on how to describing Worker selection criteria below. | [optional] 
**TaskOrder** | **string** | How Tasks will be assigned to Workers. Can be: &#x60;FIFO&#x60; or &#x60;LIFO&#x60; and the default is &#x60;FIFO&#x60;. Use &#x60;FIFO&#x60; to assign the oldest task first and &#x60;LIFO&#x60; to assign the most recent task first. For more information, see [Queue Ordering](https://www.twilio.com/docs/taskrouter/queue-ordering-last-first-out-lifo). | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


