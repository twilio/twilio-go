# UpdateWorkflowRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssignmentCallbackUrl** | **string** | The URL from your application that will process task assignment events. See [Handling Task Assignment Callback](https://www.twilio.com/docs/taskrouter/handle-assignment-callbacks) for more details. | [optional] 
**Configuration** | **string** | A JSON string that contains the rules to apply to the Workflow. See [Configuring Workflows](https://www.twilio.com/docs/taskrouter/workflow-configuration) for more information. | [optional] 
**FallbackAssignmentCallbackUrl** | **string** | The URL that we should call when a call to the &#x60;assignment_callback_url&#x60; fails. | [optional] 
**FriendlyName** | **string** | A descriptive string that you create to describe the Workflow resource. For example, &#x60;Inbound Call Workflow&#x60; or &#x60;2014 Outbound Campaign&#x60;. | [optional] 
**ReEvaluateTasks** | **string** | Whether or not to re-evaluate Tasks. The default is &#x60;false&#x60;, which means Tasks in the Workflow will not be processed through the assignment loop again. | [optional] 
**TaskReservationTimeout** | **int32** | How long TaskRouter will wait for a confirmation response from your application after it assigns a Task to a Worker. Can be up to &#x60;86,400&#x60; (24 hours) and the default is &#x60;120&#x60;. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


