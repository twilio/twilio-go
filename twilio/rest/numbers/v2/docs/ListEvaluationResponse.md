# ListEvaluationResponse

## Properties

Name | Type | Description
------------ | ------------- | -------------
**Meta** | Pointer to [**ListBundleResponseMeta**](ListBundleResponse_meta.md) |  | [optional] 
**Results** | Pointer to [**[]NumbersV2RegulatoryComplianceBundleEvaluation**](NumbersV2RegulatoryComplianceBundleEvaluation.md) |  | [optional] 

## Methods

### NewListEvaluationResponse

`func NewListEvaluationResponse() *ListEvaluationResponse`

NewListEvaluationResponse instantiates a new ListEvaluationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListEvaluationResponseWithDefaults

`func NewListEvaluationResponseWithDefaults() *ListEvaluationResponse`

NewListEvaluationResponseWithDefaults instantiates a new ListEvaluationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListEvaluationResponse) GetMeta() ListBundleResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListEvaluationResponse) GetMetaOk() (*ListBundleResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListEvaluationResponse) SetMeta(v ListBundleResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListEvaluationResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetResults

`func (o *ListEvaluationResponse) GetResults() []NumbersV2RegulatoryComplianceBundleEvaluation`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ListEvaluationResponse) GetResultsOk() (*[]NumbersV2RegulatoryComplianceBundleEvaluation, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ListEvaluationResponse) SetResults(v []NumbersV2RegulatoryComplianceBundleEvaluation)`

SetResults sets Results field to given value.

### HasResults

`func (o *ListEvaluationResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


