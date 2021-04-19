# ListSinkResponse

## Properties

Name | Type | Description
------------ | ------------- | -------------
**Meta** | Pointer to [**ListVersionResponseMeta**](ListVersionResponse_meta.md) |  | [optional] 
**Sinks** | Pointer to [**[]EventsV1Sink**](EventsV1Sink.md) |  | [optional] 

## Methods

### NewListSinkResponse

`func NewListSinkResponse() *ListSinkResponse`

NewListSinkResponse instantiates a new ListSinkResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSinkResponseWithDefaults

`func NewListSinkResponseWithDefaults() *ListSinkResponse`

NewListSinkResponseWithDefaults instantiates a new ListSinkResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListSinkResponse) GetMeta() ListVersionResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListSinkResponse) GetMetaOk() (*ListVersionResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListSinkResponse) SetMeta(v ListVersionResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListSinkResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetSinks

`func (o *ListSinkResponse) GetSinks() []EventsV1Sink`

GetSinks returns the Sinks field if non-nil, zero value otherwise.

### GetSinksOk

`func (o *ListSinkResponse) GetSinksOk() (*[]EventsV1Sink, bool)`

GetSinksOk returns a tuple with the Sinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSinks

`func (o *ListSinkResponse) SetSinks(v []EventsV1Sink)`

SetSinks sets Sinks field to given value.

### HasSinks

`func (o *ListSinkResponse) HasSinks() bool`

HasSinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


