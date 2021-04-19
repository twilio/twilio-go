# ListBuildResponse

## Properties

Name | Type | Description
------------ | ------------- | -------------
**Builds** | Pointer to [**[]ServerlessV1ServiceBuild**](ServerlessV1ServiceBuild.md) |  | [optional] 
**Meta** | Pointer to [**ListServiceResponseMeta**](ListServiceResponse_meta.md) |  | [optional] 

## Methods

### NewListBuildResponse

`func NewListBuildResponse() *ListBuildResponse`

NewListBuildResponse instantiates a new ListBuildResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListBuildResponseWithDefaults

`func NewListBuildResponseWithDefaults() *ListBuildResponse`

NewListBuildResponseWithDefaults instantiates a new ListBuildResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuilds

`func (o *ListBuildResponse) GetBuilds() []ServerlessV1ServiceBuild`

GetBuilds returns the Builds field if non-nil, zero value otherwise.

### GetBuildsOk

`func (o *ListBuildResponse) GetBuildsOk() (*[]ServerlessV1ServiceBuild, bool)`

GetBuildsOk returns a tuple with the Builds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuilds

`func (o *ListBuildResponse) SetBuilds(v []ServerlessV1ServiceBuild)`

SetBuilds sets Builds field to given value.

### HasBuilds

`func (o *ListBuildResponse) HasBuilds() bool`

HasBuilds returns a boolean if a field has been set.

### GetMeta

`func (o *ListBuildResponse) GetMeta() ListServiceResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListBuildResponse) GetMetaOk() (*ListServiceResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListBuildResponse) SetMeta(v ListServiceResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListBuildResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


