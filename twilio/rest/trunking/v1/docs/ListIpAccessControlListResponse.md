# ListIpAccessControlListResponse

## Properties

Name | Type | Description
------------ | ------------- | -------------
**IpAccessControlLists** | Pointer to [**[]TrunkingV1TrunkIpAccessControlList**](TrunkingV1TrunkIpAccessControlList.md) |  | [optional] 
**Meta** | Pointer to [**ListTrunkResponseMeta**](ListTrunkResponse_meta.md) |  | [optional] 

## Methods

### NewListIpAccessControlListResponse

`func NewListIpAccessControlListResponse() *ListIpAccessControlListResponse`

NewListIpAccessControlListResponse instantiates a new ListIpAccessControlListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListIpAccessControlListResponseWithDefaults

`func NewListIpAccessControlListResponseWithDefaults() *ListIpAccessControlListResponse`

NewListIpAccessControlListResponseWithDefaults instantiates a new ListIpAccessControlListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpAccessControlLists

`func (o *ListIpAccessControlListResponse) GetIpAccessControlLists() []TrunkingV1TrunkIpAccessControlList`

GetIpAccessControlLists returns the IpAccessControlLists field if non-nil, zero value otherwise.

### GetIpAccessControlListsOk

`func (o *ListIpAccessControlListResponse) GetIpAccessControlListsOk() (*[]TrunkingV1TrunkIpAccessControlList, bool)`

GetIpAccessControlListsOk returns a tuple with the IpAccessControlLists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAccessControlLists

`func (o *ListIpAccessControlListResponse) SetIpAccessControlLists(v []TrunkingV1TrunkIpAccessControlList)`

SetIpAccessControlLists sets IpAccessControlLists field to given value.

### HasIpAccessControlLists

`func (o *ListIpAccessControlListResponse) HasIpAccessControlLists() bool`

HasIpAccessControlLists returns a boolean if a field has been set.

### GetMeta

`func (o *ListIpAccessControlListResponse) GetMeta() ListTrunkResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListIpAccessControlListResponse) GetMetaOk() (*ListTrunkResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListIpAccessControlListResponse) SetMeta(v ListTrunkResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListIpAccessControlListResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


