# ListEndUserTypeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndUserTypes** | Pointer to [**[]NumbersV2RegulatoryComplianceEndUserType**](NumbersV2RegulatoryComplianceEndUserType.md) |  | [optional] 
**Meta** | Pointer to [**ListBundleResponseMeta**](ListBundleResponse_meta.md) |  | [optional] 

## Methods

### NewListEndUserTypeResponse

`func NewListEndUserTypeResponse() *ListEndUserTypeResponse`

NewListEndUserTypeResponse instantiates a new ListEndUserTypeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListEndUserTypeResponseWithDefaults

`func NewListEndUserTypeResponseWithDefaults() *ListEndUserTypeResponse`

NewListEndUserTypeResponseWithDefaults instantiates a new ListEndUserTypeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndUserTypes

`func (o *ListEndUserTypeResponse) GetEndUserTypes() []NumbersV2RegulatoryComplianceEndUserType`

GetEndUserTypes returns the EndUserTypes field if non-nil, zero value otherwise.

### GetEndUserTypesOk

`func (o *ListEndUserTypeResponse) GetEndUserTypesOk() (*[]NumbersV2RegulatoryComplianceEndUserType, bool)`

GetEndUserTypesOk returns a tuple with the EndUserTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndUserTypes

`func (o *ListEndUserTypeResponse) SetEndUserTypes(v []NumbersV2RegulatoryComplianceEndUserType)`

SetEndUserTypes sets EndUserTypes field to given value.

### HasEndUserTypes

`func (o *ListEndUserTypeResponse) HasEndUserTypes() bool`

HasEndUserTypes returns a boolean if a field has been set.

### GetMeta

`func (o *ListEndUserTypeResponse) GetMeta() ListBundleResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListEndUserTypeResponse) GetMetaOk() (*ListBundleResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListEndUserTypeResponse) SetMeta(v ListBundleResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListEndUserTypeResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


