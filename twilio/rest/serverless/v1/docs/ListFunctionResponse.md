# ListFunctionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Functions** | Pointer to [**[]ServerlessV1ServiceFunction**](ServerlessV1ServiceFunction.md) |  | [optional] 
**Meta** | Pointer to [**ListServiceResponseMeta**](ListServiceResponse_meta.md) |  | [optional] 

## Methods

### NewListFunctionResponse

`func NewListFunctionResponse() *ListFunctionResponse`

NewListFunctionResponse instantiates a new ListFunctionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListFunctionResponseWithDefaults

`func NewListFunctionResponseWithDefaults() *ListFunctionResponse`

NewListFunctionResponseWithDefaults instantiates a new ListFunctionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFunctions

`func (o *ListFunctionResponse) GetFunctions() []ServerlessV1ServiceFunction`

GetFunctions returns the Functions field if non-nil, zero value otherwise.

### GetFunctionsOk

`func (o *ListFunctionResponse) GetFunctionsOk() (*[]ServerlessV1ServiceFunction, bool)`

GetFunctionsOk returns a tuple with the Functions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctions

`func (o *ListFunctionResponse) SetFunctions(v []ServerlessV1ServiceFunction)`

SetFunctions sets Functions field to given value.

### HasFunctions

`func (o *ListFunctionResponse) HasFunctions() bool`

HasFunctions returns a boolean if a field has been set.

### GetMeta

`func (o *ListFunctionResponse) GetMeta() ListServiceResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListFunctionResponse) GetMetaOk() (*ListServiceResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListFunctionResponse) SetMeta(v ListServiceResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListFunctionResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


