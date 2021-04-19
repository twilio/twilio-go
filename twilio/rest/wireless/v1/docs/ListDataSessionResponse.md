# ListDataSessionResponse

## Properties

Name | Type | Description
------------ | ------------- | -------------
**DataSessions** | Pointer to [**[]WirelessV1SimDataSession**](WirelessV1SimDataSession.md) |  | [optional] 
**Meta** | Pointer to [**ListCommandResponseMeta**](ListCommandResponse_meta.md) |  | [optional] 

## Methods

### NewListDataSessionResponse

`func NewListDataSessionResponse() *ListDataSessionResponse`

NewListDataSessionResponse instantiates a new ListDataSessionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDataSessionResponseWithDefaults

`func NewListDataSessionResponseWithDefaults() *ListDataSessionResponse`

NewListDataSessionResponseWithDefaults instantiates a new ListDataSessionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataSessions

`func (o *ListDataSessionResponse) GetDataSessions() []WirelessV1SimDataSession`

GetDataSessions returns the DataSessions field if non-nil, zero value otherwise.

### GetDataSessionsOk

`func (o *ListDataSessionResponse) GetDataSessionsOk() (*[]WirelessV1SimDataSession, bool)`

GetDataSessionsOk returns a tuple with the DataSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSessions

`func (o *ListDataSessionResponse) SetDataSessions(v []WirelessV1SimDataSession)`

SetDataSessions sets DataSessions field to given value.

### HasDataSessions

`func (o *ListDataSessionResponse) HasDataSessions() bool`

HasDataSessions returns a boolean if a field has been set.

### GetMeta

`func (o *ListDataSessionResponse) GetMeta() ListCommandResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListDataSessionResponse) GetMetaOk() (*ListCommandResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListDataSessionResponse) SetMeta(v ListCommandResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListDataSessionResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


