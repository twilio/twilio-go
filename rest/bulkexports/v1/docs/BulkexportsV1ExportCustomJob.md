# BulkexportsV1ExportCustomJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FriendlyName** | Pointer to **string** | The friendly name specified when creating the job |
**ResourceType** | Pointer to **string** | The type of communication – Messages, Calls, Conferences, and Participants |
**StartDay** | Pointer to **string** | The start day for the custom export specified as a string in the format of yyyy-MM-dd |
**EndDay** | Pointer to **string** | The end day for the custom export specified as a string in the format of yyyy-MM-dd. This will be the last day exported. For instance, to export a single day, choose the same day for start and end day. To export the first 4 days of July, you would set the start date to 2020-07-01 and the end date to 2020-07-04. The end date must be the UTC day before yesterday. |
**WebhookUrl** | Pointer to **string** | The optional webhook url called on completion |
**WebhookMethod** | Pointer to **string** | This is the method used to call the webhook |
**Email** | Pointer to **string** | The optional email to send the completion notification to |
**JobSid** | Pointer to **string** | The unique job_sid returned when the custom export was created. This can be used to look up the status of the job. |
**Details** | Pointer to **interface{}** | The details of a job state which is an object that contains a `status` string, a day count integer, and list of days in the job |
**JobQueuePosition** | Pointer to **string** | This is the job position from the 1st in line. Your queue position will never increase. As jobs ahead of yours in the queue are processed, the queue position number will decrease |
**EstimatedCompletionTime** | Pointer to **string** | this is the time estimated until your job is complete. This is calculated each time you request the job list. The time is calculated based on the current rate of job completion (which may vary) and your job queue position |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


