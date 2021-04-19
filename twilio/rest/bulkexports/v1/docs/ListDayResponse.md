# ListDayResponse

## Properties

Name | Type | Description
------------ | ------------- | -------------
**Days** | Pointer to [**[]BulkexportsV1ExportDay**](BulkexportsV1ExportDay.md) |  | [optional] 
**Meta** | Pointer to [**ListDayResponseMeta**](ListDayResponse_meta.md) |  | [optional] 

## Methods

### NewListDayResponse

`func NewListDayResponse() *ListDayResponse`

NewListDayResponse instantiates a new ListDayResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDayResponseWithDefaults

`func NewListDayResponseWithDefaults() *ListDayResponse`

NewListDayResponseWithDefaults instantiates a new ListDayResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDays

`func (o *ListDayResponse) GetDays() []BulkexportsV1ExportDay`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *ListDayResponse) GetDaysOk() (*[]BulkexportsV1ExportDay, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *ListDayResponse) SetDays(v []BulkexportsV1ExportDay)`

SetDays sets Days field to given value.

### HasDays

`func (o *ListDayResponse) HasDays() bool`

HasDays returns a boolean if a field has been set.

### GetMeta

`func (o *ListDayResponse) GetMeta() ListDayResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListDayResponse) GetMetaOk() (*ListDayResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListDayResponse) SetMeta(v ListDayResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListDayResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


