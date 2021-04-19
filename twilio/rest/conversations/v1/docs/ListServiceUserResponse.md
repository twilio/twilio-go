# ListServiceUserResponse

## Properties

Name | Type | Description
------------ | ------------- | -------------
**Meta** | Pointer to [**ListConversationResponseMeta**](ListConversationResponse_meta.md) |  | [optional] 
**Users** | Pointer to [**[]ConversationsV1ServiceServiceUser**](ConversationsV1ServiceServiceUser.md) |  | [optional] 

## Methods

### NewListServiceUserResponse

`func NewListServiceUserResponse() *ListServiceUserResponse`

NewListServiceUserResponse instantiates a new ListServiceUserResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListServiceUserResponseWithDefaults

`func NewListServiceUserResponseWithDefaults() *ListServiceUserResponse`

NewListServiceUserResponseWithDefaults instantiates a new ListServiceUserResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListServiceUserResponse) GetMeta() ListConversationResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListServiceUserResponse) GetMetaOk() (*ListConversationResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListServiceUserResponse) SetMeta(v ListConversationResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListServiceUserResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUsers

`func (o *ListServiceUserResponse) GetUsers() []ConversationsV1ServiceServiceUser`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *ListServiceUserResponse) GetUsersOk() (*[]ConversationsV1ServiceServiceUser, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *ListServiceUserResponse) SetUsers(v []ConversationsV1ServiceServiceUser)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *ListServiceUserResponse) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


