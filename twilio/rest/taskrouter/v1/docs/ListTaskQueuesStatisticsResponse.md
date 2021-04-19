# ListTaskQueuesStatisticsResponse

## Properties

Name | Type | Description
------------ | ------------- | -------------
**Meta** | Pointer to [**ListWorkspaceResponseMeta**](ListWorkspaceResponse_meta.md) |  | [optional] 
**TaskQueuesStatistics** | Pointer to [**[]TaskrouterV1WorkspaceTaskQueueTaskQueuesStatistics**](TaskrouterV1WorkspaceTaskQueueTaskQueuesStatistics.md) |  | [optional] 

## Methods

### NewListTaskQueuesStatisticsResponse

`func NewListTaskQueuesStatisticsResponse() *ListTaskQueuesStatisticsResponse`

NewListTaskQueuesStatisticsResponse instantiates a new ListTaskQueuesStatisticsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTaskQueuesStatisticsResponseWithDefaults

`func NewListTaskQueuesStatisticsResponseWithDefaults() *ListTaskQueuesStatisticsResponse`

NewListTaskQueuesStatisticsResponseWithDefaults instantiates a new ListTaskQueuesStatisticsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListTaskQueuesStatisticsResponse) GetMeta() ListWorkspaceResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListTaskQueuesStatisticsResponse) GetMetaOk() (*ListWorkspaceResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListTaskQueuesStatisticsResponse) SetMeta(v ListWorkspaceResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListTaskQueuesStatisticsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetTaskQueuesStatistics

`func (o *ListTaskQueuesStatisticsResponse) GetTaskQueuesStatistics() []TaskrouterV1WorkspaceTaskQueueTaskQueuesStatistics`

GetTaskQueuesStatistics returns the TaskQueuesStatistics field if non-nil, zero value otherwise.

### GetTaskQueuesStatisticsOk

`func (o *ListTaskQueuesStatisticsResponse) GetTaskQueuesStatisticsOk() (*[]TaskrouterV1WorkspaceTaskQueueTaskQueuesStatistics, bool)`

GetTaskQueuesStatisticsOk returns a tuple with the TaskQueuesStatistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskQueuesStatistics

`func (o *ListTaskQueuesStatisticsResponse) SetTaskQueuesStatistics(v []TaskrouterV1WorkspaceTaskQueueTaskQueuesStatistics)`

SetTaskQueuesStatistics sets TaskQueuesStatistics field to given value.

### HasTaskQueuesStatistics

`func (o *ListTaskQueuesStatisticsResponse) HasTaskQueuesStatistics() bool`

HasTaskQueuesStatistics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


