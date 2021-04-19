# ListActivityResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Activities** | Pointer to [**[]TaskrouterV1WorkspaceActivity**](TaskrouterV1WorkspaceActivity.md) |  | [optional] 
**Meta** | Pointer to [**ListWorkspaceResponseMeta**](ListWorkspaceResponse_meta.md) |  | [optional] 

## Methods

### NewListActivityResponse

`func NewListActivityResponse() *ListActivityResponse`

NewListActivityResponse instantiates a new ListActivityResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListActivityResponseWithDefaults

`func NewListActivityResponseWithDefaults() *ListActivityResponse`

NewListActivityResponseWithDefaults instantiates a new ListActivityResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivities

`func (o *ListActivityResponse) GetActivities() []TaskrouterV1WorkspaceActivity`

GetActivities returns the Activities field if non-nil, zero value otherwise.

### GetActivitiesOk

`func (o *ListActivityResponse) GetActivitiesOk() (*[]TaskrouterV1WorkspaceActivity, bool)`

GetActivitiesOk returns a tuple with the Activities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivities

`func (o *ListActivityResponse) SetActivities(v []TaskrouterV1WorkspaceActivity)`

SetActivities sets Activities field to given value.

### HasActivities

`func (o *ListActivityResponse) HasActivities() bool`

HasActivities returns a boolean if a field has been set.

### GetMeta

`func (o *ListActivityResponse) GetMeta() ListWorkspaceResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListActivityResponse) GetMetaOk() (*ListWorkspaceResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListActivityResponse) SetMeta(v ListWorkspaceResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListActivityResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


