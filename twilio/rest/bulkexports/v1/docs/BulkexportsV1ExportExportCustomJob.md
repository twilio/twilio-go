# BulkexportsV1ExportExportCustomJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Details** | Pointer to **map[string]interface{}** | The details of a job state which is an object that contains a status string, a day count integer, and list of days in the job | [optional] 
**Email** | Pointer to **NullableString** | The optional email to send the completion notification to | [optional] 
**EndDay** | Pointer to **NullableString** | The end day for the custom export specified as a string in the format of yyyy-MM-dd. This will be the last day exported. For instance, to export a single day, choose the same day for start and end day. To export the first 4 days of July, you would set the start date to 2020-07-01 and the end date to 2020-07-04. The end date must be the UTC day before yesterday. | [optional] 
**EstimatedCompletionTime** | Pointer to **NullableString** | this is the time estimated until your job is complete. This is calculated each time you request the job list. The time is calculated based on the current rate of job completion (which may vary) and your job queue position | [optional] 
**FriendlyName** | Pointer to **NullableString** | The friendly name specified when creating the job | [optional] 
**JobQueuePosition** | Pointer to **NullableString** | This is the job position from the 1st in line. Your queue position will never increase. As jobs ahead of yours in the queue are processed, the queue position number will decrease | [optional] 
**JobSid** | Pointer to **NullableString** | The unique job_sid returned when the custom export was created. This can be used to look up the status of the job. | [optional] 
**ResourceType** | Pointer to **NullableString** | The type of communication – Messages, Calls, Conferences, and Participants | [optional] 
**StartDay** | Pointer to **NullableString** | The start day for the custom export specified as a string in the format of yyyy-MM-dd | [optional] 
**WebhookMethod** | Pointer to **NullableString** | This is the method used to call the webhook | [optional] 
**WebhookUrl** | Pointer to **NullableString** | The optional webhook url called on completion | [optional] 

## Methods

### NewBulkexportsV1ExportExportCustomJob

`func NewBulkexportsV1ExportExportCustomJob() *BulkexportsV1ExportExportCustomJob`

NewBulkexportsV1ExportExportCustomJob instantiates a new BulkexportsV1ExportExportCustomJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkexportsV1ExportExportCustomJobWithDefaults

`func NewBulkexportsV1ExportExportCustomJobWithDefaults() *BulkexportsV1ExportExportCustomJob`

NewBulkexportsV1ExportExportCustomJobWithDefaults instantiates a new BulkexportsV1ExportExportCustomJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetails

`func (o *BulkexportsV1ExportExportCustomJob) GetDetails() map[string]interface{}`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *BulkexportsV1ExportExportCustomJob) GetDetailsOk() (*map[string]interface{}, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *BulkexportsV1ExportExportCustomJob) SetDetails(v map[string]interface{})`

SetDetails sets Details field to given value.

### HasDetails

`func (o *BulkexportsV1ExportExportCustomJob) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### SetDetailsNil

`func (o *BulkexportsV1ExportExportCustomJob) SetDetailsNil(b bool)`

 SetDetailsNil sets the value for Details to be an explicit nil

### UnsetDetails
`func (o *BulkexportsV1ExportExportCustomJob) UnsetDetails()`

UnsetDetails ensures that no value is present for Details, not even an explicit nil
### GetEmail

`func (o *BulkexportsV1ExportExportCustomJob) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *BulkexportsV1ExportExportCustomJob) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *BulkexportsV1ExportExportCustomJob) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *BulkexportsV1ExportExportCustomJob) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *BulkexportsV1ExportExportCustomJob) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *BulkexportsV1ExportExportCustomJob) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetEndDay

`func (o *BulkexportsV1ExportExportCustomJob) GetEndDay() string`

GetEndDay returns the EndDay field if non-nil, zero value otherwise.

### GetEndDayOk

`func (o *BulkexportsV1ExportExportCustomJob) GetEndDayOk() (*string, bool)`

GetEndDayOk returns a tuple with the EndDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDay

`func (o *BulkexportsV1ExportExportCustomJob) SetEndDay(v string)`

SetEndDay sets EndDay field to given value.

### HasEndDay

`func (o *BulkexportsV1ExportExportCustomJob) HasEndDay() bool`

HasEndDay returns a boolean if a field has been set.

### SetEndDayNil

`func (o *BulkexportsV1ExportExportCustomJob) SetEndDayNil(b bool)`

 SetEndDayNil sets the value for EndDay to be an explicit nil

### UnsetEndDay
`func (o *BulkexportsV1ExportExportCustomJob) UnsetEndDay()`

UnsetEndDay ensures that no value is present for EndDay, not even an explicit nil
### GetEstimatedCompletionTime

`func (o *BulkexportsV1ExportExportCustomJob) GetEstimatedCompletionTime() string`

GetEstimatedCompletionTime returns the EstimatedCompletionTime field if non-nil, zero value otherwise.

### GetEstimatedCompletionTimeOk

`func (o *BulkexportsV1ExportExportCustomJob) GetEstimatedCompletionTimeOk() (*string, bool)`

GetEstimatedCompletionTimeOk returns a tuple with the EstimatedCompletionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedCompletionTime

`func (o *BulkexportsV1ExportExportCustomJob) SetEstimatedCompletionTime(v string)`

SetEstimatedCompletionTime sets EstimatedCompletionTime field to given value.

### HasEstimatedCompletionTime

`func (o *BulkexportsV1ExportExportCustomJob) HasEstimatedCompletionTime() bool`

HasEstimatedCompletionTime returns a boolean if a field has been set.

### SetEstimatedCompletionTimeNil

`func (o *BulkexportsV1ExportExportCustomJob) SetEstimatedCompletionTimeNil(b bool)`

 SetEstimatedCompletionTimeNil sets the value for EstimatedCompletionTime to be an explicit nil

### UnsetEstimatedCompletionTime
`func (o *BulkexportsV1ExportExportCustomJob) UnsetEstimatedCompletionTime()`

UnsetEstimatedCompletionTime ensures that no value is present for EstimatedCompletionTime, not even an explicit nil
### GetFriendlyName

`func (o *BulkexportsV1ExportExportCustomJob) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *BulkexportsV1ExportExportCustomJob) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *BulkexportsV1ExportExportCustomJob) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *BulkexportsV1ExportExportCustomJob) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *BulkexportsV1ExportExportCustomJob) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *BulkexportsV1ExportExportCustomJob) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetJobQueuePosition

`func (o *BulkexportsV1ExportExportCustomJob) GetJobQueuePosition() string`

GetJobQueuePosition returns the JobQueuePosition field if non-nil, zero value otherwise.

### GetJobQueuePositionOk

`func (o *BulkexportsV1ExportExportCustomJob) GetJobQueuePositionOk() (*string, bool)`

GetJobQueuePositionOk returns a tuple with the JobQueuePosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobQueuePosition

`func (o *BulkexportsV1ExportExportCustomJob) SetJobQueuePosition(v string)`

SetJobQueuePosition sets JobQueuePosition field to given value.

### HasJobQueuePosition

`func (o *BulkexportsV1ExportExportCustomJob) HasJobQueuePosition() bool`

HasJobQueuePosition returns a boolean if a field has been set.

### SetJobQueuePositionNil

`func (o *BulkexportsV1ExportExportCustomJob) SetJobQueuePositionNil(b bool)`

 SetJobQueuePositionNil sets the value for JobQueuePosition to be an explicit nil

### UnsetJobQueuePosition
`func (o *BulkexportsV1ExportExportCustomJob) UnsetJobQueuePosition()`

UnsetJobQueuePosition ensures that no value is present for JobQueuePosition, not even an explicit nil
### GetJobSid

`func (o *BulkexportsV1ExportExportCustomJob) GetJobSid() string`

GetJobSid returns the JobSid field if non-nil, zero value otherwise.

### GetJobSidOk

`func (o *BulkexportsV1ExportExportCustomJob) GetJobSidOk() (*string, bool)`

GetJobSidOk returns a tuple with the JobSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobSid

`func (o *BulkexportsV1ExportExportCustomJob) SetJobSid(v string)`

SetJobSid sets JobSid field to given value.

### HasJobSid

`func (o *BulkexportsV1ExportExportCustomJob) HasJobSid() bool`

HasJobSid returns a boolean if a field has been set.

### SetJobSidNil

`func (o *BulkexportsV1ExportExportCustomJob) SetJobSidNil(b bool)`

 SetJobSidNil sets the value for JobSid to be an explicit nil

### UnsetJobSid
`func (o *BulkexportsV1ExportExportCustomJob) UnsetJobSid()`

UnsetJobSid ensures that no value is present for JobSid, not even an explicit nil
### GetResourceType

`func (o *BulkexportsV1ExportExportCustomJob) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *BulkexportsV1ExportExportCustomJob) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *BulkexportsV1ExportExportCustomJob) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *BulkexportsV1ExportExportCustomJob) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### SetResourceTypeNil

`func (o *BulkexportsV1ExportExportCustomJob) SetResourceTypeNil(b bool)`

 SetResourceTypeNil sets the value for ResourceType to be an explicit nil

### UnsetResourceType
`func (o *BulkexportsV1ExportExportCustomJob) UnsetResourceType()`

UnsetResourceType ensures that no value is present for ResourceType, not even an explicit nil
### GetStartDay

`func (o *BulkexportsV1ExportExportCustomJob) GetStartDay() string`

GetStartDay returns the StartDay field if non-nil, zero value otherwise.

### GetStartDayOk

`func (o *BulkexportsV1ExportExportCustomJob) GetStartDayOk() (*string, bool)`

GetStartDayOk returns a tuple with the StartDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDay

`func (o *BulkexportsV1ExportExportCustomJob) SetStartDay(v string)`

SetStartDay sets StartDay field to given value.

### HasStartDay

`func (o *BulkexportsV1ExportExportCustomJob) HasStartDay() bool`

HasStartDay returns a boolean if a field has been set.

### SetStartDayNil

`func (o *BulkexportsV1ExportExportCustomJob) SetStartDayNil(b bool)`

 SetStartDayNil sets the value for StartDay to be an explicit nil

### UnsetStartDay
`func (o *BulkexportsV1ExportExportCustomJob) UnsetStartDay()`

UnsetStartDay ensures that no value is present for StartDay, not even an explicit nil
### GetWebhookMethod

`func (o *BulkexportsV1ExportExportCustomJob) GetWebhookMethod() string`

GetWebhookMethod returns the WebhookMethod field if non-nil, zero value otherwise.

### GetWebhookMethodOk

`func (o *BulkexportsV1ExportExportCustomJob) GetWebhookMethodOk() (*string, bool)`

GetWebhookMethodOk returns a tuple with the WebhookMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookMethod

`func (o *BulkexportsV1ExportExportCustomJob) SetWebhookMethod(v string)`

SetWebhookMethod sets WebhookMethod field to given value.

### HasWebhookMethod

`func (o *BulkexportsV1ExportExportCustomJob) HasWebhookMethod() bool`

HasWebhookMethod returns a boolean if a field has been set.

### SetWebhookMethodNil

`func (o *BulkexportsV1ExportExportCustomJob) SetWebhookMethodNil(b bool)`

 SetWebhookMethodNil sets the value for WebhookMethod to be an explicit nil

### UnsetWebhookMethod
`func (o *BulkexportsV1ExportExportCustomJob) UnsetWebhookMethod()`

UnsetWebhookMethod ensures that no value is present for WebhookMethod, not even an explicit nil
### GetWebhookUrl

`func (o *BulkexportsV1ExportExportCustomJob) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *BulkexportsV1ExportExportCustomJob) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *BulkexportsV1ExportExportCustomJob) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *BulkexportsV1ExportExportCustomJob) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### SetWebhookUrlNil

`func (o *BulkexportsV1ExportExportCustomJob) SetWebhookUrlNil(b bool)`

 SetWebhookUrlNil sets the value for WebhookUrl to be an explicit nil

### UnsetWebhookUrl
`func (o *BulkexportsV1ExportExportCustomJob) UnsetWebhookUrl()`

UnsetWebhookUrl ensures that no value is present for WebhookUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


