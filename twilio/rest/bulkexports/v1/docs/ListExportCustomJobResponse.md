# ListExportCustomJobResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jobs** | Pointer to [**[]BulkexportsV1ExportExportCustomJob**](BulkexportsV1ExportExportCustomJob.md) |  | [optional] 
**Meta** | Pointer to [**ListDayResponseMeta**](ListDayResponse_meta.md) |  | [optional] 

## Methods

### NewListExportCustomJobResponse

`func NewListExportCustomJobResponse() *ListExportCustomJobResponse`

NewListExportCustomJobResponse instantiates a new ListExportCustomJobResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListExportCustomJobResponseWithDefaults

`func NewListExportCustomJobResponseWithDefaults() *ListExportCustomJobResponse`

NewListExportCustomJobResponseWithDefaults instantiates a new ListExportCustomJobResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobs

`func (o *ListExportCustomJobResponse) GetJobs() []BulkexportsV1ExportExportCustomJob`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *ListExportCustomJobResponse) GetJobsOk() (*[]BulkexportsV1ExportExportCustomJob, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *ListExportCustomJobResponse) SetJobs(v []BulkexportsV1ExportExportCustomJob)`

SetJobs sets Jobs field to given value.

### HasJobs

`func (o *ListExportCustomJobResponse) HasJobs() bool`

HasJobs returns a boolean if a field has been set.

### GetMeta

`func (o *ListExportCustomJobResponse) GetMeta() ListDayResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListExportCustomJobResponse) GetMetaOk() (*ListDayResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListExportCustomJobResponse) SetMeta(v ListDayResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListExportCustomJobResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


