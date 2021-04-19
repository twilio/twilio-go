# ListSyncListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Lists** | Pointer to [**[]SyncV1ServiceSyncList**](SyncV1ServiceSyncList.md) |  | [optional] 
**Meta** | Pointer to [**ListServiceResponseMeta**](ListServiceResponse_meta.md) |  | [optional] 

## Methods

### NewListSyncListResponse

`func NewListSyncListResponse() *ListSyncListResponse`

NewListSyncListResponse instantiates a new ListSyncListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSyncListResponseWithDefaults

`func NewListSyncListResponseWithDefaults() *ListSyncListResponse`

NewListSyncListResponseWithDefaults instantiates a new ListSyncListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLists

`func (o *ListSyncListResponse) GetLists() []SyncV1ServiceSyncList`

GetLists returns the Lists field if non-nil, zero value otherwise.

### GetListsOk

`func (o *ListSyncListResponse) GetListsOk() (*[]SyncV1ServiceSyncList, bool)`

GetListsOk returns a tuple with the Lists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLists

`func (o *ListSyncListResponse) SetLists(v []SyncV1ServiceSyncList)`

SetLists sets Lists field to given value.

### HasLists

`func (o *ListSyncListResponse) HasLists() bool`

HasLists returns a boolean if a field has been set.

### GetMeta

`func (o *ListSyncListResponse) GetMeta() ListServiceResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListSyncListResponse) GetMetaOk() (*ListServiceResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListSyncListResponse) SetMeta(v ListServiceResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListSyncListResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


