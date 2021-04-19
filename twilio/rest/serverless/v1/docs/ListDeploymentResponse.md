# ListDeploymentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deployments** | Pointer to [**[]ServerlessV1ServiceEnvironmentDeployment**](ServerlessV1ServiceEnvironmentDeployment.md) |  | [optional] 
**Meta** | Pointer to [**ListServiceResponseMeta**](ListServiceResponse_meta.md) |  | [optional] 

## Methods

### NewListDeploymentResponse

`func NewListDeploymentResponse() *ListDeploymentResponse`

NewListDeploymentResponse instantiates a new ListDeploymentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDeploymentResponseWithDefaults

`func NewListDeploymentResponseWithDefaults() *ListDeploymentResponse`

NewListDeploymentResponseWithDefaults instantiates a new ListDeploymentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeployments

`func (o *ListDeploymentResponse) GetDeployments() []ServerlessV1ServiceEnvironmentDeployment`

GetDeployments returns the Deployments field if non-nil, zero value otherwise.

### GetDeploymentsOk

`func (o *ListDeploymentResponse) GetDeploymentsOk() (*[]ServerlessV1ServiceEnvironmentDeployment, bool)`

GetDeploymentsOk returns a tuple with the Deployments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployments

`func (o *ListDeploymentResponse) SetDeployments(v []ServerlessV1ServiceEnvironmentDeployment)`

SetDeployments sets Deployments field to given value.

### HasDeployments

`func (o *ListDeploymentResponse) HasDeployments() bool`

HasDeployments returns a boolean if a field has been set.

### GetMeta

`func (o *ListDeploymentResponse) GetMeta() ListServiceResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListDeploymentResponse) GetMetaOk() (*ListServiceResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListDeploymentResponse) SetMeta(v ListServiceResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListDeploymentResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


