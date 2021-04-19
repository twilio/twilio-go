# ListUserResponse

## Properties

Name | Type | Description
------------ | ------------- | -------------
**Meta** | Pointer to [**ListConversationResponseMeta**](ListConversationResponse_meta.md) |  | [optional] 
**Users** | Pointer to [**[]ConversationsV1User**](ConversationsV1User.md) |  | [optional] 

## Methods

### NewListUserResponse

`func NewListUserResponse() *ListUserResponse`

NewListUserResponse instantiates a new ListUserResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListUserResponseWithDefaults

`func NewListUserResponseWithDefaults() *ListUserResponse`

NewListUserResponseWithDefaults instantiates a new ListUserResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListUserResponse) GetMeta() ListConversationResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListUserResponse) GetMetaOk() (*ListConversationResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListUserResponse) SetMeta(v ListConversationResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListUserResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUsers

`func (o *ListUserResponse) GetUsers() []ConversationsV1User`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *ListUserResponse) GetUsersOk() (*[]ConversationsV1User, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *ListUserResponse) SetUsers(v []ConversationsV1User)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *ListUserResponse) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


