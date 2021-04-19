# ListFactorResponse

## Properties

Name | Type | Description
------------ | ------------- | -------------
**Factors** | Pointer to [**[]VerifyV2ServiceEntityFactor**](VerifyV2ServiceEntityFactor.md) |  | [optional] 
**Meta** | Pointer to [**ListVerificationAttemptResponseMeta**](ListVerificationAttemptResponse_meta.md) |  | [optional] 

## Methods

### NewListFactorResponse

`func NewListFactorResponse() *ListFactorResponse`

NewListFactorResponse instantiates a new ListFactorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListFactorResponseWithDefaults

`func NewListFactorResponseWithDefaults() *ListFactorResponse`

NewListFactorResponseWithDefaults instantiates a new ListFactorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFactors

`func (o *ListFactorResponse) GetFactors() []VerifyV2ServiceEntityFactor`

GetFactors returns the Factors field if non-nil, zero value otherwise.

### GetFactorsOk

`func (o *ListFactorResponse) GetFactorsOk() (*[]VerifyV2ServiceEntityFactor, bool)`

GetFactorsOk returns a tuple with the Factors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactors

`func (o *ListFactorResponse) SetFactors(v []VerifyV2ServiceEntityFactor)`

SetFactors sets Factors field to given value.

### HasFactors

`func (o *ListFactorResponse) HasFactors() bool`

HasFactors returns a boolean if a field has been set.

### GetMeta

`func (o *ListFactorResponse) GetMeta() ListVerificationAttemptResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListFactorResponse) GetMetaOk() (*ListVerificationAttemptResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListFactorResponse) SetMeta(v ListVerificationAttemptResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListFactorResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


