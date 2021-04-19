# ListWorkerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**ListWorkspaceResponseMeta**](ListWorkspaceResponse_meta.md) |  | [optional] 
**Workers** | Pointer to [**[]TaskrouterV1WorkspaceWorker**](TaskrouterV1WorkspaceWorker.md) |  | [optional] 

## Methods

### NewListWorkerResponse

`func NewListWorkerResponse() *ListWorkerResponse`

NewListWorkerResponse instantiates a new ListWorkerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListWorkerResponseWithDefaults

`func NewListWorkerResponseWithDefaults() *ListWorkerResponse`

NewListWorkerResponseWithDefaults instantiates a new ListWorkerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListWorkerResponse) GetMeta() ListWorkspaceResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListWorkerResponse) GetMetaOk() (*ListWorkspaceResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListWorkerResponse) SetMeta(v ListWorkspaceResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListWorkerResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetWorkers

`func (o *ListWorkerResponse) GetWorkers() []TaskrouterV1WorkspaceWorker`

GetWorkers returns the Workers field if non-nil, zero value otherwise.

### GetWorkersOk

`func (o *ListWorkerResponse) GetWorkersOk() (*[]TaskrouterV1WorkspaceWorker, bool)`

GetWorkersOk returns a tuple with the Workers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkers

`func (o *ListWorkerResponse) SetWorkers(v []TaskrouterV1WorkspaceWorker)`

SetWorkers sets Workers field to given value.

### HasWorkers

`func (o *ListWorkerResponse) HasWorkers() bool`

HasWorkers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


