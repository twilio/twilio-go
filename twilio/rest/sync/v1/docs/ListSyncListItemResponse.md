# ListSyncListItemResponse

## Properties

Name | Type | Description
------------ | ------------- | -------------
**Items** | Pointer to [**[]SyncV1ServiceSyncListSyncListItem**](SyncV1ServiceSyncListSyncListItem.md) |  | [optional] 
**Meta** | Pointer to [**ListServiceResponseMeta**](ListServiceResponse_meta.md) |  | [optional] 

## Methods

### NewListSyncListItemResponse

`func NewListSyncListItemResponse() *ListSyncListItemResponse`

NewListSyncListItemResponse instantiates a new ListSyncListItemResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSyncListItemResponseWithDefaults

`func NewListSyncListItemResponseWithDefaults() *ListSyncListItemResponse`

NewListSyncListItemResponseWithDefaults instantiates a new ListSyncListItemResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ListSyncListItemResponse) GetItems() []SyncV1ServiceSyncListSyncListItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListSyncListItemResponse) GetItemsOk() (*[]SyncV1ServiceSyncListSyncListItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListSyncListItemResponse) SetItems(v []SyncV1ServiceSyncListSyncListItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *ListSyncListItemResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetMeta

`func (o *ListSyncListItemResponse) GetMeta() ListServiceResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListSyncListItemResponse) GetMetaOk() (*ListServiceResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListSyncListItemResponse) SetMeta(v ListServiceResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListSyncListItemResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


