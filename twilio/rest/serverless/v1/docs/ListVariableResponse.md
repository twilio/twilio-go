# ListVariableResponse

## Properties

Name | Type | Description
------------ | ------------- | -------------
**Meta** | Pointer to [**ListServiceResponseMeta**](ListServiceResponse_meta.md) |  | [optional] 
**Variables** | Pointer to [**[]ServerlessV1ServiceEnvironmentVariable**](ServerlessV1ServiceEnvironmentVariable.md) |  | [optional] 

## Methods

### NewListVariableResponse

`func NewListVariableResponse() *ListVariableResponse`

NewListVariableResponse instantiates a new ListVariableResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListVariableResponseWithDefaults

`func NewListVariableResponseWithDefaults() *ListVariableResponse`

NewListVariableResponseWithDefaults instantiates a new ListVariableResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListVariableResponse) GetMeta() ListServiceResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListVariableResponse) GetMetaOk() (*ListServiceResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListVariableResponse) SetMeta(v ListServiceResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListVariableResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetVariables

`func (o *ListVariableResponse) GetVariables() []ServerlessV1ServiceEnvironmentVariable`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *ListVariableResponse) GetVariablesOk() (*[]ServerlessV1ServiceEnvironmentVariable, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *ListVariableResponse) SetVariables(v []ServerlessV1ServiceEnvironmentVariable)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *ListVariableResponse) HasVariables() bool`

HasVariables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


