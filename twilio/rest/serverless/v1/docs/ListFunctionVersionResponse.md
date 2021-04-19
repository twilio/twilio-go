# ListFunctionVersionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FunctionVersions** | Pointer to [**[]ServerlessV1ServiceFunctionFunctionVersion**](ServerlessV1ServiceFunctionFunctionVersion.md) |  | [optional] 
**Meta** | Pointer to [**ListServiceResponseMeta**](ListServiceResponse_meta.md) |  | [optional] 

## Methods

### NewListFunctionVersionResponse

`func NewListFunctionVersionResponse() *ListFunctionVersionResponse`

NewListFunctionVersionResponse instantiates a new ListFunctionVersionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListFunctionVersionResponseWithDefaults

`func NewListFunctionVersionResponseWithDefaults() *ListFunctionVersionResponse`

NewListFunctionVersionResponseWithDefaults instantiates a new ListFunctionVersionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFunctionVersions

`func (o *ListFunctionVersionResponse) GetFunctionVersions() []ServerlessV1ServiceFunctionFunctionVersion`

GetFunctionVersions returns the FunctionVersions field if non-nil, zero value otherwise.

### GetFunctionVersionsOk

`func (o *ListFunctionVersionResponse) GetFunctionVersionsOk() (*[]ServerlessV1ServiceFunctionFunctionVersion, bool)`

GetFunctionVersionsOk returns a tuple with the FunctionVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionVersions

`func (o *ListFunctionVersionResponse) SetFunctionVersions(v []ServerlessV1ServiceFunctionFunctionVersion)`

SetFunctionVersions sets FunctionVersions field to given value.

### HasFunctionVersions

`func (o *ListFunctionVersionResponse) HasFunctionVersions() bool`

HasFunctionVersions returns a boolean if a field has been set.

### GetMeta

`func (o *ListFunctionVersionResponse) GetMeta() ListServiceResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListFunctionVersionResponse) GetMetaOk() (*ListServiceResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListFunctionVersionResponse) SetMeta(v ListServiceResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListFunctionVersionResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


