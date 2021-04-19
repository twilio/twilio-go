# ListEventTypeResponse

## Properties

Name | Type | Description
------------ | ------------- | -------------
**Meta** | Pointer to [**ListVersionResponseMeta**](ListVersionResponse_meta.md) |  | [optional] 
**Types** | Pointer to [**[]EventsV1EventType**](EventsV1EventType.md) |  | [optional] 

## Methods

### NewListEventTypeResponse

`func NewListEventTypeResponse() *ListEventTypeResponse`

NewListEventTypeResponse instantiates a new ListEventTypeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListEventTypeResponseWithDefaults

`func NewListEventTypeResponseWithDefaults() *ListEventTypeResponse`

NewListEventTypeResponseWithDefaults instantiates a new ListEventTypeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListEventTypeResponse) GetMeta() ListVersionResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListEventTypeResponse) GetMetaOk() (*ListVersionResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListEventTypeResponse) SetMeta(v ListVersionResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListEventTypeResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetTypes

`func (o *ListEventTypeResponse) GetTypes() []EventsV1EventType`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *ListEventTypeResponse) GetTypesOk() (*[]EventsV1EventType, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *ListEventTypeResponse) SetTypes(v []EventsV1EventType)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *ListEventTypeResponse) HasTypes() bool`

HasTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


