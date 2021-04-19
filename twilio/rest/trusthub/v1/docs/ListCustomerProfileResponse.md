# ListCustomerProfileResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**ListCustomerProfileResponseMeta**](ListCustomerProfileResponse_meta.md) |  | [optional] 
**Results** | Pointer to [**[]TrusthubV1CustomerProfile**](TrusthubV1CustomerProfile.md) |  | [optional] 

## Methods

### NewListCustomerProfileResponse

`func NewListCustomerProfileResponse() *ListCustomerProfileResponse`

NewListCustomerProfileResponse instantiates a new ListCustomerProfileResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCustomerProfileResponseWithDefaults

`func NewListCustomerProfileResponseWithDefaults() *ListCustomerProfileResponse`

NewListCustomerProfileResponseWithDefaults instantiates a new ListCustomerProfileResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListCustomerProfileResponse) GetMeta() ListCustomerProfileResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListCustomerProfileResponse) GetMetaOk() (*ListCustomerProfileResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListCustomerProfileResponse) SetMeta(v ListCustomerProfileResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListCustomerProfileResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetResults

`func (o *ListCustomerProfileResponse) GetResults() []TrusthubV1CustomerProfile`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ListCustomerProfileResponse) GetResultsOk() (*[]TrusthubV1CustomerProfile, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ListCustomerProfileResponse) SetResults(v []TrusthubV1CustomerProfile)`

SetResults sets Results field to given value.

### HasResults

`func (o *ListCustomerProfileResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


