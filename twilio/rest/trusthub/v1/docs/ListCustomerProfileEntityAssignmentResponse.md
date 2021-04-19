# ListCustomerProfileEntityAssignmentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**ListCustomerProfileResponseMeta**](ListCustomerProfileResponse_meta.md) |  | [optional] 
**Results** | Pointer to [**[]TrusthubV1CustomerProfileCustomerProfileEntityAssignment**](TrusthubV1CustomerProfileCustomerProfileEntityAssignment.md) |  | [optional] 

## Methods

### NewListCustomerProfileEntityAssignmentResponse

`func NewListCustomerProfileEntityAssignmentResponse() *ListCustomerProfileEntityAssignmentResponse`

NewListCustomerProfileEntityAssignmentResponse instantiates a new ListCustomerProfileEntityAssignmentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCustomerProfileEntityAssignmentResponseWithDefaults

`func NewListCustomerProfileEntityAssignmentResponseWithDefaults() *ListCustomerProfileEntityAssignmentResponse`

NewListCustomerProfileEntityAssignmentResponseWithDefaults instantiates a new ListCustomerProfileEntityAssignmentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListCustomerProfileEntityAssignmentResponse) GetMeta() ListCustomerProfileResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListCustomerProfileEntityAssignmentResponse) GetMetaOk() (*ListCustomerProfileResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListCustomerProfileEntityAssignmentResponse) SetMeta(v ListCustomerProfileResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListCustomerProfileEntityAssignmentResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetResults

`func (o *ListCustomerProfileEntityAssignmentResponse) GetResults() []TrusthubV1CustomerProfileCustomerProfileEntityAssignment`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ListCustomerProfileEntityAssignmentResponse) GetResultsOk() (*[]TrusthubV1CustomerProfileCustomerProfileEntityAssignment, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ListCustomerProfileEntityAssignmentResponse) SetResults(v []TrusthubV1CustomerProfileCustomerProfileEntityAssignment)`

SetResults sets Results field to given value.

### HasResults

`func (o *ListCustomerProfileEntityAssignmentResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


