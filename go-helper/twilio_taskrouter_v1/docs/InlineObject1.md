# InlineObject1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultActivitySid** | **string** | The SID of the Activity that will be used when new Workers are created in the Workspace. | [optional] 
**EventCallbackUrl** | **string** | The URL we should call when an event occurs. See [Workspace Events](https://www.twilio.com/docs/taskrouter/api/event) for more information. | [optional] 
**EventsFilter** | **string** | The list of Workspace events for which to call event_callback_url. For example if &#x60;EventsFilter&#x3D;task.created,task.canceled,worker.activity.update&#x60;, then TaskRouter will call event_callback_url only when a task is created, canceled, or a Worker activity is updated. | [optional] 
**FriendlyName** | **string** | A descriptive string that you create to describe the Workspace resource. For example: &#x60;Sales Call Center&#x60; or &#x60;Customer Support Team&#x60;. | [optional] 
**MultiTaskEnabled** | **bool** | Whether to enable multi-tasking. Can be: &#x60;true&#x60; to enable multi-tasking, or &#x60;false&#x60; to disable it. The default is &#x60;false&#x60;. Multi-tasking allows Workers to handle multiple Tasks simultaneously. When enabled (&#x60;true&#x60;), each Worker can receive parallel reservations up to the per-channel maximums defined in the Workers section. Otherwise, each Worker will only receive a new reservation when the previous task is completed. Learn more at [Multitasking][https://www.twilio.com/docs/taskrouter/multitasking]. | [optional] 
**PrioritizeQueueOrder** | **string** | The type of TaskQueue to prioritize when Workers are receiving Tasks from both types of TaskQueues. Can be: &#x60;LIFO&#x60; or &#x60;FIFO&#x60; and the default is &#x60;FIFO&#x60;. For more information, see [Queue Ordering][https://www.twilio.com/docs/taskrouter/queue-ordering-last-first-out-lifo]. | [optional] 
**TimeoutActivitySid** | **string** | The SID of the Activity that will be assigned to a Worker when a Task reservation times out without a response. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


