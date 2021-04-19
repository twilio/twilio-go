# ListRoleResponse

## Properties

Name | Type | Description
------------ | ------------- | -------------
**Meta** | Pointer to [**ListCredentialResponseMeta**](ListCredentialResponse_meta.md) |  | [optional] 
**Roles** | Pointer to [**[]ChatV2ServiceRole**](ChatV2ServiceRole.md) |  | [optional] 

## Methods

### NewListRoleResponse

`func NewListRoleResponse() *ListRoleResponse`

NewListRoleResponse instantiates a new ListRoleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListRoleResponseWithDefaults

`func NewListRoleResponseWithDefaults() *ListRoleResponse`

NewListRoleResponseWithDefaults instantiates a new ListRoleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListRoleResponse) GetMeta() ListCredentialResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListRoleResponse) GetMetaOk() (*ListCredentialResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListRoleResponse) SetMeta(v ListCredentialResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListRoleResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetRoles

`func (o *ListRoleResponse) GetRoles() []ChatV2ServiceRole`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *ListRoleResponse) GetRolesOk() (*[]ChatV2ServiceRole, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *ListRoleResponse) SetRoles(v []ChatV2ServiceRole)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *ListRoleResponse) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


