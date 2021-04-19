# ListWorkspaceResponse

## Properties

Name | Type | Description
------------ | ------------- | -------------
**Meta** | Pointer to [**ListWorkspaceResponseMeta**](ListWorkspaceResponse_meta.md) |  | [optional] 
**Workspaces** | Pointer to [**[]TaskrouterV1Workspace**](TaskrouterV1Workspace.md) |  | [optional] 

## Methods

### NewListWorkspaceResponse

`func NewListWorkspaceResponse() *ListWorkspaceResponse`

NewListWorkspaceResponse instantiates a new ListWorkspaceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListWorkspaceResponseWithDefaults

`func NewListWorkspaceResponseWithDefaults() *ListWorkspaceResponse`

NewListWorkspaceResponseWithDefaults instantiates a new ListWorkspaceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListWorkspaceResponse) GetMeta() ListWorkspaceResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListWorkspaceResponse) GetMetaOk() (*ListWorkspaceResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListWorkspaceResponse) SetMeta(v ListWorkspaceResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListWorkspaceResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetWorkspaces

`func (o *ListWorkspaceResponse) GetWorkspaces() []TaskrouterV1Workspace`

GetWorkspaces returns the Workspaces field if non-nil, zero value otherwise.

### GetWorkspacesOk

`func (o *ListWorkspaceResponse) GetWorkspacesOk() (*[]TaskrouterV1Workspace, bool)`

GetWorkspacesOk returns a tuple with the Workspaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaces

`func (o *ListWorkspaceResponse) SetWorkspaces(v []TaskrouterV1Workspace)`

SetWorkspaces sets Workspaces field to given value.

### HasWorkspaces

`func (o *ListWorkspaceResponse) HasWorkspaces() bool`

HasWorkspaces returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


