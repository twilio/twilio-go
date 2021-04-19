# ListTaskResponse

## Properties

Name | Type | Description
------------ | ------------- | -------------
**Meta** | Pointer to [**ListAssistantResponseMeta**](ListAssistantResponse_meta.md) |  | [optional] 
**Tasks** | Pointer to [**[]AutopilotV1AssistantTask**](AutopilotV1AssistantTask.md) |  | [optional] 

## Methods

### NewListTaskResponse

`func NewListTaskResponse() *ListTaskResponse`

NewListTaskResponse instantiates a new ListTaskResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTaskResponseWithDefaults

`func NewListTaskResponseWithDefaults() *ListTaskResponse`

NewListTaskResponseWithDefaults instantiates a new ListTaskResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListTaskResponse) GetMeta() ListAssistantResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListTaskResponse) GetMetaOk() (*ListAssistantResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListTaskResponse) SetMeta(v ListAssistantResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListTaskResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetTasks

`func (o *ListTaskResponse) GetTasks() []AutopilotV1AssistantTask`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *ListTaskResponse) GetTasksOk() (*[]AutopilotV1AssistantTask, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *ListTaskResponse) SetTasks(v []AutopilotV1AssistantTask)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *ListTaskResponse) HasTasks() bool`

HasTasks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


