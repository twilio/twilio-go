# ListPoliciesResponse

## Properties

Name | Type | Description
------------ | ------------- | -------------
**Meta** | Pointer to [**ListCustomerProfileResponseMeta**](ListCustomerProfileResponse_meta.md) |  | [optional] 
**Results** | Pointer to [**[]TrusthubV1Policies**](TrusthubV1Policies.md) |  | [optional] 

## Methods

### NewListPoliciesResponse

`func NewListPoliciesResponse() *ListPoliciesResponse`

NewListPoliciesResponse instantiates a new ListPoliciesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPoliciesResponseWithDefaults

`func NewListPoliciesResponseWithDefaults() *ListPoliciesResponse`

NewListPoliciesResponseWithDefaults instantiates a new ListPoliciesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListPoliciesResponse) GetMeta() ListCustomerProfileResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListPoliciesResponse) GetMetaOk() (*ListCustomerProfileResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListPoliciesResponse) SetMeta(v ListCustomerProfileResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListPoliciesResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetResults

`func (o *ListPoliciesResponse) GetResults() []TrusthubV1Policies`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ListPoliciesResponse) GetResultsOk() (*[]TrusthubV1Policies, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ListPoliciesResponse) SetResults(v []TrusthubV1Policies)`

SetResults sets Results field to given value.

### HasResults

`func (o *ListPoliciesResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


