# ListTaskQueueResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**ListWorkspaceResponseMeta**](ListWorkspaceResponse_meta.md) |  | [optional] 
**TaskQueues** | Pointer to [**[]TaskrouterV1WorkspaceTaskQueue**](TaskrouterV1WorkspaceTaskQueue.md) |  | [optional] 

## Methods

### NewListTaskQueueResponse

`func NewListTaskQueueResponse() *ListTaskQueueResponse`

NewListTaskQueueResponse instantiates a new ListTaskQueueResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTaskQueueResponseWithDefaults

`func NewListTaskQueueResponseWithDefaults() *ListTaskQueueResponse`

NewListTaskQueueResponseWithDefaults instantiates a new ListTaskQueueResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListTaskQueueResponse) GetMeta() ListWorkspaceResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListTaskQueueResponse) GetMetaOk() (*ListWorkspaceResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListTaskQueueResponse) SetMeta(v ListWorkspaceResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListTaskQueueResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetTaskQueues

`func (o *ListTaskQueueResponse) GetTaskQueues() []TaskrouterV1WorkspaceTaskQueue`

GetTaskQueues returns the TaskQueues field if non-nil, zero value otherwise.

### GetTaskQueuesOk

`func (o *ListTaskQueueResponse) GetTaskQueuesOk() (*[]TaskrouterV1WorkspaceTaskQueue, bool)`

GetTaskQueuesOk returns a tuple with the TaskQueues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskQueues

`func (o *ListTaskQueueResponse) SetTaskQueues(v []TaskrouterV1WorkspaceTaskQueue)`

SetTaskQueues sets TaskQueues field to given value.

### HasTaskQueues

`func (o *ListTaskQueueResponse) HasTaskQueues() bool`

HasTaskQueues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


