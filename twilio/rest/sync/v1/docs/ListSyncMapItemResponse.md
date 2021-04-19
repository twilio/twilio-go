# ListSyncMapItemResponse

## Properties

Name | Type | Description
------------ | ------------- | -------------
**Items** | Pointer to [**[]SyncV1ServiceSyncMapSyncMapItem**](SyncV1ServiceSyncMapSyncMapItem.md) |  | [optional] 
**Meta** | Pointer to [**ListServiceResponseMeta**](ListServiceResponse_meta.md) |  | [optional] 

## Methods

### NewListSyncMapItemResponse

`func NewListSyncMapItemResponse() *ListSyncMapItemResponse`

NewListSyncMapItemResponse instantiates a new ListSyncMapItemResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSyncMapItemResponseWithDefaults

`func NewListSyncMapItemResponseWithDefaults() *ListSyncMapItemResponse`

NewListSyncMapItemResponseWithDefaults instantiates a new ListSyncMapItemResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ListSyncMapItemResponse) GetItems() []SyncV1ServiceSyncMapSyncMapItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListSyncMapItemResponse) GetItemsOk() (*[]SyncV1ServiceSyncMapSyncMapItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListSyncMapItemResponse) SetItems(v []SyncV1ServiceSyncMapSyncMapItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *ListSyncMapItemResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetMeta

`func (o *ListSyncMapItemResponse) GetMeta() ListServiceResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListSyncMapItemResponse) GetMetaOk() (*ListServiceResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListSyncMapItemResponse) SetMeta(v ListServiceResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListSyncMapItemResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


