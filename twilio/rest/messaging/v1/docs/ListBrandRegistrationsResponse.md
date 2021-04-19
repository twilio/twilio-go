# ListBrandRegistrationsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]MessagingV1BrandRegistrations**](MessagingV1BrandRegistrations.md) |  | [optional] 
**Meta** | Pointer to [**ListServiceResponseMeta**](ListServiceResponse_meta.md) |  | [optional] 

## Methods

### NewListBrandRegistrationsResponse

`func NewListBrandRegistrationsResponse() *ListBrandRegistrationsResponse`

NewListBrandRegistrationsResponse instantiates a new ListBrandRegistrationsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListBrandRegistrationsResponseWithDefaults

`func NewListBrandRegistrationsResponseWithDefaults() *ListBrandRegistrationsResponse`

NewListBrandRegistrationsResponseWithDefaults instantiates a new ListBrandRegistrationsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListBrandRegistrationsResponse) GetData() []MessagingV1BrandRegistrations`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListBrandRegistrationsResponse) GetDataOk() (*[]MessagingV1BrandRegistrations, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListBrandRegistrationsResponse) SetData(v []MessagingV1BrandRegistrations)`

SetData sets Data field to given value.

### HasData

`func (o *ListBrandRegistrationsResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *ListBrandRegistrationsResponse) GetMeta() ListServiceResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListBrandRegistrationsResponse) GetMetaOk() (*ListServiceResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListBrandRegistrationsResponse) SetMeta(v ListServiceResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListBrandRegistrationsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


