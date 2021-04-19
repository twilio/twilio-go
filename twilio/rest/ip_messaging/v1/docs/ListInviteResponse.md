# ListInviteResponse

## Properties

Name | Type | Description
------------ | ------------- | -------------
**Invites** | Pointer to [**[]IpMessagingV1ServiceChannelInvite**](IpMessagingV1ServiceChannelInvite.md) |  | [optional] 
**Meta** | Pointer to [**ListCredentialResponseMeta**](ListCredentialResponse_meta.md) |  | [optional] 

## Methods

### NewListInviteResponse

`func NewListInviteResponse() *ListInviteResponse`

NewListInviteResponse instantiates a new ListInviteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListInviteResponseWithDefaults

`func NewListInviteResponseWithDefaults() *ListInviteResponse`

NewListInviteResponseWithDefaults instantiates a new ListInviteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvites

`func (o *ListInviteResponse) GetInvites() []IpMessagingV1ServiceChannelInvite`

GetInvites returns the Invites field if non-nil, zero value otherwise.

### GetInvitesOk

`func (o *ListInviteResponse) GetInvitesOk() (*[]IpMessagingV1ServiceChannelInvite, bool)`

GetInvitesOk returns a tuple with the Invites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvites

`func (o *ListInviteResponse) SetInvites(v []IpMessagingV1ServiceChannelInvite)`

SetInvites sets Invites field to given value.

### HasInvites

`func (o *ListInviteResponse) HasInvites() bool`

HasInvites returns a boolean if a field has been set.

### GetMeta

`func (o *ListInviteResponse) GetMeta() ListCredentialResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListInviteResponse) GetMetaOk() (*ListCredentialResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListInviteResponse) SetMeta(v ListCredentialResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListInviteResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


