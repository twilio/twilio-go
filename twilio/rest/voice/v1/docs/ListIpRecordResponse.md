# ListIpRecordResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpRecords** | Pointer to [**[]VoiceV1IpRecord**](VoiceV1IpRecord.md) |  | [optional] 
**Meta** | Pointer to [**ListByocTrunkResponseMeta**](ListByocTrunkResponse_meta.md) |  | [optional] 

## Methods

### NewListIpRecordResponse

`func NewListIpRecordResponse() *ListIpRecordResponse`

NewListIpRecordResponse instantiates a new ListIpRecordResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListIpRecordResponseWithDefaults

`func NewListIpRecordResponseWithDefaults() *ListIpRecordResponse`

NewListIpRecordResponseWithDefaults instantiates a new ListIpRecordResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpRecords

`func (o *ListIpRecordResponse) GetIpRecords() []VoiceV1IpRecord`

GetIpRecords returns the IpRecords field if non-nil, zero value otherwise.

### GetIpRecordsOk

`func (o *ListIpRecordResponse) GetIpRecordsOk() (*[]VoiceV1IpRecord, bool)`

GetIpRecordsOk returns a tuple with the IpRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRecords

`func (o *ListIpRecordResponse) SetIpRecords(v []VoiceV1IpRecord)`

SetIpRecords sets IpRecords field to given value.

### HasIpRecords

`func (o *ListIpRecordResponse) HasIpRecords() bool`

HasIpRecords returns a boolean if a field has been set.

### GetMeta

`func (o *ListIpRecordResponse) GetMeta() ListByocTrunkResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListIpRecordResponse) GetMetaOk() (*ListByocTrunkResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListIpRecordResponse) SetMeta(v ListByocTrunkResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListIpRecordResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


