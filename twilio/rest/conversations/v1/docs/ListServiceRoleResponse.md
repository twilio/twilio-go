# ListServiceRoleResponse

## Properties

Name | Type | Description
------------ | ------------- | -------------
**Meta** | Pointer to [**ListConversationResponseMeta**](ListConversationResponse_meta.md) |  | [optional] 
**Roles** | Pointer to [**[]ConversationsV1ServiceServiceRole**](ConversationsV1ServiceServiceRole.md) |  | [optional] 

## Methods

### NewListServiceRoleResponse

`func NewListServiceRoleResponse() *ListServiceRoleResponse`

NewListServiceRoleResponse instantiates a new ListServiceRoleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListServiceRoleResponseWithDefaults

`func NewListServiceRoleResponseWithDefaults() *ListServiceRoleResponse`

NewListServiceRoleResponseWithDefaults instantiates a new ListServiceRoleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListServiceRoleResponse) GetMeta() ListConversationResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListServiceRoleResponse) GetMetaOk() (*ListConversationResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListServiceRoleResponse) SetMeta(v ListConversationResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListServiceRoleResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetRoles

`func (o *ListServiceRoleResponse) GetRoles() []ConversationsV1ServiceServiceRole`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *ListServiceRoleResponse) GetRolesOk() (*[]ConversationsV1ServiceServiceRole, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *ListServiceRoleResponse) SetRoles(v []ConversationsV1ServiceServiceRole)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *ListServiceRoleResponse) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


