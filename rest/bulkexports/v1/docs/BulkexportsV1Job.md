# BulkexportsV1Job

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceType** | **string** | The type of communication – Messages, Calls, Conferences, and Participants |[optional] 
**FriendlyName** | **string** | The friendly name specified when creating the job |[optional] 
**Details** | Pointer to **interface{}** | The details of a job state which is an object that contains a `status` string, a day count integer, and list of days in the job |
**StartDay** | **string** | The start time for the export specified when creating the job |[optional] 
**EndDay** | **string** | The end time for the export specified when creating the job |[optional] 
**JobSid** | **string** | The job_sid returned when the export was created |[optional] 
**WebhookUrl** | **string** | The optional webhook url called on completion |[optional] 
**WebhookMethod** | **string** | This is the method used to call the webhook |[optional] 
**Email** | **string** | The optional email to send the completion notification to |[optional] 
**Url** | **string** |  |[optional] 
**JobQueuePosition** | **string** | This is the job position from the 1st in line. Your queue position will never increase. As jobs ahead of yours in the queue are processed, the queue position number will decrease |[optional] 
**EstimatedCompletionTime** | **string** | this is the time estimated until your job is complete. This is calculated each time you request the job list. The time is calculated based on the current rate of job completion (which may vary) and your job queue position |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


