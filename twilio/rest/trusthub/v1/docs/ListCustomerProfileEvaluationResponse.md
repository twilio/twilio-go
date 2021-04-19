# ListCustomerProfileEvaluationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**ListCustomerProfileResponseMeta**](ListCustomerProfileResponse_meta.md) |  | [optional] 
**Results** | Pointer to [**[]TrusthubV1CustomerProfileCustomerProfileEvaluation**](TrusthubV1CustomerProfileCustomerProfileEvaluation.md) |  | [optional] 

## Methods

### NewListCustomerProfileEvaluationResponse

`func NewListCustomerProfileEvaluationResponse() *ListCustomerProfileEvaluationResponse`

NewListCustomerProfileEvaluationResponse instantiates a new ListCustomerProfileEvaluationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCustomerProfileEvaluationResponseWithDefaults

`func NewListCustomerProfileEvaluationResponseWithDefaults() *ListCustomerProfileEvaluationResponse`

NewListCustomerProfileEvaluationResponseWithDefaults instantiates a new ListCustomerProfileEvaluationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListCustomerProfileEvaluationResponse) GetMeta() ListCustomerProfileResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListCustomerProfileEvaluationResponse) GetMetaOk() (*ListCustomerProfileResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListCustomerProfileEvaluationResponse) SetMeta(v ListCustomerProfileResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListCustomerProfileEvaluationResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetResults

`func (o *ListCustomerProfileEvaluationResponse) GetResults() []TrusthubV1CustomerProfileCustomerProfileEvaluation`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ListCustomerProfileEvaluationResponse) GetResultsOk() (*[]TrusthubV1CustomerProfileCustomerProfileEvaluation, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ListCustomerProfileEvaluationResponse) SetResults(v []TrusthubV1CustomerProfileCustomerProfileEvaluation)`

SetResults sets Results field to given value.

### HasResults

`func (o *ListCustomerProfileEvaluationResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


