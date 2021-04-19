# BulkexportsV1ExportJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Details** | Pointer to **map[string]interface{}** | This is a list of the completed, pending, or errored dates within the export time range, with one entry for each status with more than one day in that status | [optional] 
**Email** | Pointer to **NullableString** | The optional email to send the completion notification to | [optional] 
**EndDay** | Pointer to **NullableString** | The end time for the export specified when creating the job | [optional] 
**EstimatedCompletionTime** | Pointer to **NullableString** | this is the time estimated until your job is complete. This is calculated each time you request the job list. The time is calculated based on the current rate of job completion (which may vary) and your job queue position | [optional] 
**FriendlyName** | Pointer to **NullableString** | The friendly name specified when creating the job | [optional] 
**JobQueuePosition** | Pointer to **NullableString** | This is the job position from the 1st in line. Your queue position will never increase. As jobs ahead of yours in the queue are processed, the queue position number will decrease | [optional] 
**JobSid** | Pointer to **NullableString** | The job_sid returned when the export was created | [optional] 
**ResourceType** | Pointer to **NullableString** | The type of communication – Messages, Calls, Conferences, and Participants | [optional] 
**StartDay** | Pointer to **NullableString** | The start time for the export specified when creating the job | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 
**WebhookMethod** | Pointer to **NullableString** | This is the method used to call the webhook | [optional] 
**WebhookUrl** | Pointer to **NullableString** | The optional webhook url called on completion | [optional] 

## Methods

### NewBulkexportsV1ExportJob

`func NewBulkexportsV1ExportJob() *BulkexportsV1ExportJob`

NewBulkexportsV1ExportJob instantiates a new BulkexportsV1ExportJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkexportsV1ExportJobWithDefaults

`func NewBulkexportsV1ExportJobWithDefaults() *BulkexportsV1ExportJob`

NewBulkexportsV1ExportJobWithDefaults instantiates a new BulkexportsV1ExportJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetails

`func (o *BulkexportsV1ExportJob) GetDetails() map[string]interface{}`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *BulkexportsV1ExportJob) GetDetailsOk() (*map[string]interface{}, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *BulkexportsV1ExportJob) SetDetails(v map[string]interface{})`

SetDetails sets Details field to given value.

### HasDetails

`func (o *BulkexportsV1ExportJob) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### SetDetailsNil

`func (o *BulkexportsV1ExportJob) SetDetailsNil(b bool)`

 SetDetailsNil sets the value for Details to be an explicit nil

### UnsetDetails
`func (o *BulkexportsV1ExportJob) UnsetDetails()`

UnsetDetails ensures that no value is present for Details, not even an explicit nil
### GetEmail

`func (o *BulkexportsV1ExportJob) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *BulkexportsV1ExportJob) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *BulkexportsV1ExportJob) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *BulkexportsV1ExportJob) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *BulkexportsV1ExportJob) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *BulkexportsV1ExportJob) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetEndDay

`func (o *BulkexportsV1ExportJob) GetEndDay() string`

GetEndDay returns the EndDay field if non-nil, zero value otherwise.

### GetEndDayOk

`func (o *BulkexportsV1ExportJob) GetEndDayOk() (*string, bool)`

GetEndDayOk returns a tuple with the EndDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDay

`func (o *BulkexportsV1ExportJob) SetEndDay(v string)`

SetEndDay sets EndDay field to given value.

### HasEndDay

`func (o *BulkexportsV1ExportJob) HasEndDay() bool`

HasEndDay returns a boolean if a field has been set.

### SetEndDayNil

`func (o *BulkexportsV1ExportJob) SetEndDayNil(b bool)`

 SetEndDayNil sets the value for EndDay to be an explicit nil

### UnsetEndDay
`func (o *BulkexportsV1ExportJob) UnsetEndDay()`

UnsetEndDay ensures that no value is present for EndDay, not even an explicit nil
### GetEstimatedCompletionTime

`func (o *BulkexportsV1ExportJob) GetEstimatedCompletionTime() string`

GetEstimatedCompletionTime returns the EstimatedCompletionTime field if non-nil, zero value otherwise.

### GetEstimatedCompletionTimeOk

`func (o *BulkexportsV1ExportJob) GetEstimatedCompletionTimeOk() (*string, bool)`

GetEstimatedCompletionTimeOk returns a tuple with the EstimatedCompletionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedCompletionTime

`func (o *BulkexportsV1ExportJob) SetEstimatedCompletionTime(v string)`

SetEstimatedCompletionTime sets EstimatedCompletionTime field to given value.

### HasEstimatedCompletionTime

`func (o *BulkexportsV1ExportJob) HasEstimatedCompletionTime() bool`

HasEstimatedCompletionTime returns a boolean if a field has been set.

### SetEstimatedCompletionTimeNil

`func (o *BulkexportsV1ExportJob) SetEstimatedCompletionTimeNil(b bool)`

 SetEstimatedCompletionTimeNil sets the value for EstimatedCompletionTime to be an explicit nil

### UnsetEstimatedCompletionTime
`func (o *BulkexportsV1ExportJob) UnsetEstimatedCompletionTime()`

UnsetEstimatedCompletionTime ensures that no value is present for EstimatedCompletionTime, not even an explicit nil
### GetFriendlyName

`func (o *BulkexportsV1ExportJob) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *BulkexportsV1ExportJob) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *BulkexportsV1ExportJob) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *BulkexportsV1ExportJob) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *BulkexportsV1ExportJob) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *BulkexportsV1ExportJob) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetJobQueuePosition

`func (o *BulkexportsV1ExportJob) GetJobQueuePosition() string`

GetJobQueuePosition returns the JobQueuePosition field if non-nil, zero value otherwise.

### GetJobQueuePositionOk

`func (o *BulkexportsV1ExportJob) GetJobQueuePositionOk() (*string, bool)`

GetJobQueuePositionOk returns a tuple with the JobQueuePosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobQueuePosition

`func (o *BulkexportsV1ExportJob) SetJobQueuePosition(v string)`

SetJobQueuePosition sets JobQueuePosition field to given value.

### HasJobQueuePosition

`func (o *BulkexportsV1ExportJob) HasJobQueuePosition() bool`

HasJobQueuePosition returns a boolean if a field has been set.

### SetJobQueuePositionNil

`func (o *BulkexportsV1ExportJob) SetJobQueuePositionNil(b bool)`

 SetJobQueuePositionNil sets the value for JobQueuePosition to be an explicit nil

### UnsetJobQueuePosition
`func (o *BulkexportsV1ExportJob) UnsetJobQueuePosition()`

UnsetJobQueuePosition ensures that no value is present for JobQueuePosition, not even an explicit nil
### GetJobSid

`func (o *BulkexportsV1ExportJob) GetJobSid() string`

GetJobSid returns the JobSid field if non-nil, zero value otherwise.

### GetJobSidOk

`func (o *BulkexportsV1ExportJob) GetJobSidOk() (*string, bool)`

GetJobSidOk returns a tuple with the JobSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobSid

`func (o *BulkexportsV1ExportJob) SetJobSid(v string)`

SetJobSid sets JobSid field to given value.

### HasJobSid

`func (o *BulkexportsV1ExportJob) HasJobSid() bool`

HasJobSid returns a boolean if a field has been set.

### SetJobSidNil

`func (o *BulkexportsV1ExportJob) SetJobSidNil(b bool)`

 SetJobSidNil sets the value for JobSid to be an explicit nil

### UnsetJobSid
`func (o *BulkexportsV1ExportJob) UnsetJobSid()`

UnsetJobSid ensures that no value is present for JobSid, not even an explicit nil
### GetResourceType

`func (o *BulkexportsV1ExportJob) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *BulkexportsV1ExportJob) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *BulkexportsV1ExportJob) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *BulkexportsV1ExportJob) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### SetResourceTypeNil

`func (o *BulkexportsV1ExportJob) SetResourceTypeNil(b bool)`

 SetResourceTypeNil sets the value for ResourceType to be an explicit nil

### UnsetResourceType
`func (o *BulkexportsV1ExportJob) UnsetResourceType()`

UnsetResourceType ensures that no value is present for ResourceType, not even an explicit nil
### GetStartDay

`func (o *BulkexportsV1ExportJob) GetStartDay() string`

GetStartDay returns the StartDay field if non-nil, zero value otherwise.

### GetStartDayOk

`func (o *BulkexportsV1ExportJob) GetStartDayOk() (*string, bool)`

GetStartDayOk returns a tuple with the StartDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDay

`func (o *BulkexportsV1ExportJob) SetStartDay(v string)`

SetStartDay sets StartDay field to given value.

### HasStartDay

`func (o *BulkexportsV1ExportJob) HasStartDay() bool`

HasStartDay returns a boolean if a field has been set.

### SetStartDayNil

`func (o *BulkexportsV1ExportJob) SetStartDayNil(b bool)`

 SetStartDayNil sets the value for StartDay to be an explicit nil

### UnsetStartDay
`func (o *BulkexportsV1ExportJob) UnsetStartDay()`

UnsetStartDay ensures that no value is present for StartDay, not even an explicit nil
### GetUrl

`func (o *BulkexportsV1ExportJob) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *BulkexportsV1ExportJob) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *BulkexportsV1ExportJob) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *BulkexportsV1ExportJob) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *BulkexportsV1ExportJob) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *BulkexportsV1ExportJob) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetWebhookMethod

`func (o *BulkexportsV1ExportJob) GetWebhookMethod() string`

GetWebhookMethod returns the WebhookMethod field if non-nil, zero value otherwise.

### GetWebhookMethodOk

`func (o *BulkexportsV1ExportJob) GetWebhookMethodOk() (*string, bool)`

GetWebhookMethodOk returns a tuple with the WebhookMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookMethod

`func (o *BulkexportsV1ExportJob) SetWebhookMethod(v string)`

SetWebhookMethod sets WebhookMethod field to given value.

### HasWebhookMethod

`func (o *BulkexportsV1ExportJob) HasWebhookMethod() bool`

HasWebhookMethod returns a boolean if a field has been set.

### SetWebhookMethodNil

`func (o *BulkexportsV1ExportJob) SetWebhookMethodNil(b bool)`

 SetWebhookMethodNil sets the value for WebhookMethod to be an explicit nil

### UnsetWebhookMethod
`func (o *BulkexportsV1ExportJob) UnsetWebhookMethod()`

UnsetWebhookMethod ensures that no value is present for WebhookMethod, not even an explicit nil
### GetWebhookUrl

`func (o *BulkexportsV1ExportJob) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *BulkexportsV1ExportJob) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *BulkexportsV1ExportJob) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *BulkexportsV1ExportJob) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### SetWebhookUrlNil

`func (o *BulkexportsV1ExportJob) SetWebhookUrlNil(b bool)`

 SetWebhookUrlNil sets the value for WebhookUrl to be an explicit nil

### UnsetWebhookUrl
`func (o *BulkexportsV1ExportJob) UnsetWebhookUrl()`

UnsetWebhookUrl ensures that no value is present for WebhookUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


