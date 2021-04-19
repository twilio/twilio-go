# ListEndUserResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**ListBundleResponseMeta**](ListBundleResponse_meta.md) |  | [optional] 
**Results** | Pointer to [**[]NumbersV2RegulatoryComplianceEndUser**](NumbersV2RegulatoryComplianceEndUser.md) |  | [optional] 

## Methods

### NewListEndUserResponse

`func NewListEndUserResponse() *ListEndUserResponse`

NewListEndUserResponse instantiates a new ListEndUserResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListEndUserResponseWithDefaults

`func NewListEndUserResponseWithDefaults() *ListEndUserResponse`

NewListEndUserResponseWithDefaults instantiates a new ListEndUserResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListEndUserResponse) GetMeta() ListBundleResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListEndUserResponse) GetMetaOk() (*ListBundleResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListEndUserResponse) SetMeta(v ListBundleResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListEndUserResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetResults

`func (o *ListEndUserResponse) GetResults() []NumbersV2RegulatoryComplianceEndUser`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ListEndUserResponse) GetResultsOk() (*[]NumbersV2RegulatoryComplianceEndUser, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ListEndUserResponse) SetResults(v []NumbersV2RegulatoryComplianceEndUser)`

SetResults sets Results field to given value.

### HasResults

`func (o *ListEndUserResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


