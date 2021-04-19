# ListUsageRecordResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**ListCommandResponseMeta**](ListCommandResponse_meta.md) |  | [optional] 
**UsageRecords** | Pointer to [**[]WirelessV1SimUsageRecord**](WirelessV1SimUsageRecord.md) |  | [optional] 

## Methods

### NewListUsageRecordResponse

`func NewListUsageRecordResponse() *ListUsageRecordResponse`

NewListUsageRecordResponse instantiates a new ListUsageRecordResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListUsageRecordResponseWithDefaults

`func NewListUsageRecordResponseWithDefaults() *ListUsageRecordResponse`

NewListUsageRecordResponseWithDefaults instantiates a new ListUsageRecordResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListUsageRecordResponse) GetMeta() ListCommandResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListUsageRecordResponse) GetMetaOk() (*ListCommandResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListUsageRecordResponse) SetMeta(v ListCommandResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListUsageRecordResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUsageRecords

`func (o *ListUsageRecordResponse) GetUsageRecords() []WirelessV1SimUsageRecord`

GetUsageRecords returns the UsageRecords field if non-nil, zero value otherwise.

### GetUsageRecordsOk

`func (o *ListUsageRecordResponse) GetUsageRecordsOk() (*[]WirelessV1SimUsageRecord, bool)`

GetUsageRecordsOk returns a tuple with the UsageRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageRecords

`func (o *ListUsageRecordResponse) SetUsageRecords(v []WirelessV1SimUsageRecord)`

SetUsageRecords sets UsageRecords field to given value.

### HasUsageRecords

`func (o *ListUsageRecordResponse) HasUsageRecords() bool`

HasUsageRecords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


