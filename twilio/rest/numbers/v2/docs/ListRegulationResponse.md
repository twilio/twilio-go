# ListRegulationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**ListBundleResponseMeta**](ListBundleResponse_meta.md) |  | [optional] 
**Results** | Pointer to [**[]NumbersV2RegulatoryComplianceRegulation**](NumbersV2RegulatoryComplianceRegulation.md) |  | [optional] 

## Methods

### NewListRegulationResponse

`func NewListRegulationResponse() *ListRegulationResponse`

NewListRegulationResponse instantiates a new ListRegulationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListRegulationResponseWithDefaults

`func NewListRegulationResponseWithDefaults() *ListRegulationResponse`

NewListRegulationResponseWithDefaults instantiates a new ListRegulationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListRegulationResponse) GetMeta() ListBundleResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListRegulationResponse) GetMetaOk() (*ListBundleResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListRegulationResponse) SetMeta(v ListBundleResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListRegulationResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetResults

`func (o *ListRegulationResponse) GetResults() []NumbersV2RegulatoryComplianceRegulation`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ListRegulationResponse) GetResultsOk() (*[]NumbersV2RegulatoryComplianceRegulation, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ListRegulationResponse) SetResults(v []NumbersV2RegulatoryComplianceRegulation)`

SetResults sets Results field to given value.

### HasResults

`func (o *ListRegulationResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


