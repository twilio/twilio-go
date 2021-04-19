# ListModelBuildResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**ListAssistantResponseMeta**](ListAssistantResponse_meta.md) |  | [optional] 
**ModelBuilds** | Pointer to [**[]AutopilotV1AssistantModelBuild**](AutopilotV1AssistantModelBuild.md) |  | [optional] 

## Methods

### NewListModelBuildResponse

`func NewListModelBuildResponse() *ListModelBuildResponse`

NewListModelBuildResponse instantiates a new ListModelBuildResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListModelBuildResponseWithDefaults

`func NewListModelBuildResponseWithDefaults() *ListModelBuildResponse`

NewListModelBuildResponseWithDefaults instantiates a new ListModelBuildResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListModelBuildResponse) GetMeta() ListAssistantResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListModelBuildResponse) GetMetaOk() (*ListAssistantResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListModelBuildResponse) SetMeta(v ListAssistantResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListModelBuildResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetModelBuilds

`func (o *ListModelBuildResponse) GetModelBuilds() []AutopilotV1AssistantModelBuild`

GetModelBuilds returns the ModelBuilds field if non-nil, zero value otherwise.

### GetModelBuildsOk

`func (o *ListModelBuildResponse) GetModelBuildsOk() (*[]AutopilotV1AssistantModelBuild, bool)`

GetModelBuildsOk returns a tuple with the ModelBuilds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelBuilds

`func (o *ListModelBuildResponse) SetModelBuilds(v []AutopilotV1AssistantModelBuild)`

SetModelBuilds sets ModelBuilds field to given value.

### HasModelBuilds

`func (o *ListModelBuildResponse) HasModelBuilds() bool`

HasModelBuilds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


