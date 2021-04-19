# ListSyncMapResponse

## Properties

Name | Type | Description
------------ | ------------- | -------------
**Maps** | Pointer to [**[]SyncV1ServiceSyncMap**](SyncV1ServiceSyncMap.md) |  | [optional] 
**Meta** | Pointer to [**ListServiceResponseMeta**](ListServiceResponse_meta.md) |  | [optional] 

## Methods

### NewListSyncMapResponse

`func NewListSyncMapResponse() *ListSyncMapResponse`

NewListSyncMapResponse instantiates a new ListSyncMapResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSyncMapResponseWithDefaults

`func NewListSyncMapResponseWithDefaults() *ListSyncMapResponse`

NewListSyncMapResponseWithDefaults instantiates a new ListSyncMapResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaps

`func (o *ListSyncMapResponse) GetMaps() []SyncV1ServiceSyncMap`

GetMaps returns the Maps field if non-nil, zero value otherwise.

### GetMapsOk

`func (o *ListSyncMapResponse) GetMapsOk() (*[]SyncV1ServiceSyncMap, bool)`

GetMapsOk returns a tuple with the Maps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaps

`func (o *ListSyncMapResponse) SetMaps(v []SyncV1ServiceSyncMap)`

SetMaps sets Maps field to given value.

### HasMaps

`func (o *ListSyncMapResponse) HasMaps() bool`

HasMaps returns a boolean if a field has been set.

### GetMeta

`func (o *ListSyncMapResponse) GetMeta() ListServiceResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListSyncMapResponse) GetMetaOk() (*ListServiceResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListSyncMapResponse) SetMeta(v ListServiceResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListSyncMapResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


