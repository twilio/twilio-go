# ListCredentialListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CredentialLists** | Pointer to [**[]TrunkingV1TrunkCredentialList**](TrunkingV1TrunkCredentialList.md) |  | [optional] 
**Meta** | Pointer to [**ListTrunkResponseMeta**](ListTrunkResponse_meta.md) |  | [optional] 

## Methods

### NewListCredentialListResponse

`func NewListCredentialListResponse() *ListCredentialListResponse`

NewListCredentialListResponse instantiates a new ListCredentialListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCredentialListResponseWithDefaults

`func NewListCredentialListResponseWithDefaults() *ListCredentialListResponse`

NewListCredentialListResponseWithDefaults instantiates a new ListCredentialListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentialLists

`func (o *ListCredentialListResponse) GetCredentialLists() []TrunkingV1TrunkCredentialList`

GetCredentialLists returns the CredentialLists field if non-nil, zero value otherwise.

### GetCredentialListsOk

`func (o *ListCredentialListResponse) GetCredentialListsOk() (*[]TrunkingV1TrunkCredentialList, bool)`

GetCredentialListsOk returns a tuple with the CredentialLists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialLists

`func (o *ListCredentialListResponse) SetCredentialLists(v []TrunkingV1TrunkCredentialList)`

SetCredentialLists sets CredentialLists field to given value.

### HasCredentialLists

`func (o *ListCredentialListResponse) HasCredentialLists() bool`

HasCredentialLists returns a boolean if a field has been set.

### GetMeta

`func (o *ListCredentialListResponse) GetMeta() ListTrunkResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListCredentialListResponse) GetMetaOk() (*ListTrunkResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListCredentialListResponse) SetMeta(v ListTrunkResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListCredentialListResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


