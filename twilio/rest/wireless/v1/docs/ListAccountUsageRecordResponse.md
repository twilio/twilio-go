# ListAccountUsageRecordResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**ListCommandResponseMeta**](ListCommandResponse_meta.md) |  | [optional] 
**UsageRecords** | Pointer to [**[]WirelessV1AccountUsageRecord**](WirelessV1AccountUsageRecord.md) |  | [optional] 

## Methods

### NewListAccountUsageRecordResponse

`func NewListAccountUsageRecordResponse() *ListAccountUsageRecordResponse`

NewListAccountUsageRecordResponse instantiates a new ListAccountUsageRecordResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAccountUsageRecordResponseWithDefaults

`func NewListAccountUsageRecordResponseWithDefaults() *ListAccountUsageRecordResponse`

NewListAccountUsageRecordResponseWithDefaults instantiates a new ListAccountUsageRecordResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListAccountUsageRecordResponse) GetMeta() ListCommandResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListAccountUsageRecordResponse) GetMetaOk() (*ListCommandResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListAccountUsageRecordResponse) SetMeta(v ListCommandResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListAccountUsageRecordResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUsageRecords

`func (o *ListAccountUsageRecordResponse) GetUsageRecords() []WirelessV1AccountUsageRecord`

GetUsageRecords returns the UsageRecords field if non-nil, zero value otherwise.

### GetUsageRecordsOk

`func (o *ListAccountUsageRecordResponse) GetUsageRecordsOk() (*[]WirelessV1AccountUsageRecord, bool)`

GetUsageRecordsOk returns a tuple with the UsageRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageRecords

`func (o *ListAccountUsageRecordResponse) SetUsageRecords(v []WirelessV1AccountUsageRecord)`

SetUsageRecords sets UsageRecords field to given value.

### HasUsageRecords

`func (o *ListAccountUsageRecordResponse) HasUsageRecords() bool`

HasUsageRecords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


