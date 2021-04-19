# ListMemberResponse

## Properties

Name | Type | Description
------------ | ------------- | -------------
**Members** | Pointer to [**[]ChatV2ServiceChannelMember**](ChatV2ServiceChannelMember.md) |  | [optional] 
**Meta** | Pointer to [**ListCredentialResponseMeta**](ListCredentialResponse_meta.md) |  | [optional] 

## Methods

### NewListMemberResponse

`func NewListMemberResponse() *ListMemberResponse`

NewListMemberResponse instantiates a new ListMemberResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListMemberResponseWithDefaults

`func NewListMemberResponseWithDefaults() *ListMemberResponse`

NewListMemberResponseWithDefaults instantiates a new ListMemberResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMembers

`func (o *ListMemberResponse) GetMembers() []ChatV2ServiceChannelMember`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *ListMemberResponse) GetMembersOk() (*[]ChatV2ServiceChannelMember, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *ListMemberResponse) SetMembers(v []ChatV2ServiceChannelMember)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *ListMemberResponse) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetMeta

`func (o *ListMemberResponse) GetMeta() ListCredentialResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListMemberResponse) GetMetaOk() (*ListCredentialResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListMemberResponse) SetMeta(v ListCredentialResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListMemberResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


