# ListWorkflowResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**ListWorkspaceResponseMeta**](ListWorkspaceResponse_meta.md) |  | [optional] 
**Workflows** | Pointer to [**[]TaskrouterV1WorkspaceWorkflow**](TaskrouterV1WorkspaceWorkflow.md) |  | [optional] 

## Methods

### NewListWorkflowResponse

`func NewListWorkflowResponse() *ListWorkflowResponse`

NewListWorkflowResponse instantiates a new ListWorkflowResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListWorkflowResponseWithDefaults

`func NewListWorkflowResponseWithDefaults() *ListWorkflowResponse`

NewListWorkflowResponseWithDefaults instantiates a new ListWorkflowResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListWorkflowResponse) GetMeta() ListWorkspaceResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListWorkflowResponse) GetMetaOk() (*ListWorkspaceResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListWorkflowResponse) SetMeta(v ListWorkspaceResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListWorkflowResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetWorkflows

`func (o *ListWorkflowResponse) GetWorkflows() []TaskrouterV1WorkspaceWorkflow`

GetWorkflows returns the Workflows field if non-nil, zero value otherwise.

### GetWorkflowsOk

`func (o *ListWorkflowResponse) GetWorkflowsOk() (*[]TaskrouterV1WorkspaceWorkflow, bool)`

GetWorkflowsOk returns a tuple with the Workflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflows

`func (o *ListWorkflowResponse) SetWorkflows(v []TaskrouterV1WorkspaceWorkflow)`

SetWorkflows sets Workflows field to given value.

### HasWorkflows

`func (o *ListWorkflowResponse) HasWorkflows() bool`

HasWorkflows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


