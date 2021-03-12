# BulkexportsV1ExportJob

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Details** | Pointer to **map[string]interface{}** | This is a list of the completed, pending, or errored dates within the export time range, with one entry for each status with more than one day in that status |
**Email** | Pointer to **string** | The optional email to send the completion notification to |
**EndDay** | Pointer to **string** | The end time for the export specified when creating the job |
**FriendlyName** | Pointer to **string** | The friendly name specified when creating the job |
**JobSid** | Pointer to **string** | The job_sid returned when the export was created |
**ResourceType** | Pointer to **string** | The type of communication – Messages, Calls, Conferences, and Participants |
**StartDay** | Pointer to **string** | The start time for the export specified when creating the job |
**Url** | Pointer to **string** |  |
**WebhookMethod** | Pointer to **string** | This is the method used to call the webhook |
**WebhookUrl** | Pointer to **string** | The optional webhook url called on completion |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


