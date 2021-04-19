# ListEnvironmentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environments** | Pointer to [**[]ServerlessV1ServiceEnvironment**](ServerlessV1ServiceEnvironment.md) |  | [optional] 
**Meta** | Pointer to [**ListServiceResponseMeta**](ListServiceResponse_meta.md) |  | [optional] 

## Methods

### NewListEnvironmentResponse

`func NewListEnvironmentResponse() *ListEnvironmentResponse`

NewListEnvironmentResponse instantiates a new ListEnvironmentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListEnvironmentResponseWithDefaults

`func NewListEnvironmentResponseWithDefaults() *ListEnvironmentResponse`

NewListEnvironmentResponseWithDefaults instantiates a new ListEnvironmentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironments

`func (o *ListEnvironmentResponse) GetEnvironments() []ServerlessV1ServiceEnvironment`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *ListEnvironmentResponse) GetEnvironmentsOk() (*[]ServerlessV1ServiceEnvironment, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *ListEnvironmentResponse) SetEnvironments(v []ServerlessV1ServiceEnvironment)`

SetEnvironments sets Environments field to given value.

### HasEnvironments

`func (o *ListEnvironmentResponse) HasEnvironments() bool`

HasEnvironments returns a boolean if a field has been set.

### GetMeta

`func (o *ListEnvironmentResponse) GetMeta() ListServiceResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListEnvironmentResponse) GetMetaOk() (*ListServiceResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListEnvironmentResponse) SetMeta(v ListServiceResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListEnvironmentResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


