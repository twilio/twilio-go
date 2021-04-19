# ListTrustProductResponse

## Properties

Name | Type | Description
------------ | ------------- | -------------
**Meta** | Pointer to [**ListCustomerProfileResponseMeta**](ListCustomerProfileResponse_meta.md) |  | [optional] 
**Results** | Pointer to [**[]TrusthubV1TrustProduct**](TrusthubV1TrustProduct.md) |  | [optional] 

## Methods

### NewListTrustProductResponse

`func NewListTrustProductResponse() *ListTrustProductResponse`

NewListTrustProductResponse instantiates a new ListTrustProductResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTrustProductResponseWithDefaults

`func NewListTrustProductResponseWithDefaults() *ListTrustProductResponse`

NewListTrustProductResponseWithDefaults instantiates a new ListTrustProductResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListTrustProductResponse) GetMeta() ListCustomerProfileResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListTrustProductResponse) GetMetaOk() (*ListCustomerProfileResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListTrustProductResponse) SetMeta(v ListCustomerProfileResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListTrustProductResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetResults

`func (o *ListTrustProductResponse) GetResults() []TrusthubV1TrustProduct`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ListTrustProductResponse) GetResultsOk() (*[]TrusthubV1TrustProduct, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ListTrustProductResponse) SetResults(v []TrusthubV1TrustProduct)`

SetResults sets Results field to given value.

### HasResults

`func (o *ListTrustProductResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


