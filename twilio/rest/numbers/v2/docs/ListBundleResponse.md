# ListBundleResponse

## Properties

Name | Type | Description
------------ | ------------- | -------------
**Meta** | Pointer to [**ListBundleResponseMeta**](ListBundleResponse_meta.md) |  | [optional] 
**Results** | Pointer to [**[]NumbersV2RegulatoryComplianceBundle**](NumbersV2RegulatoryComplianceBundle.md) |  | [optional] 

## Methods

### NewListBundleResponse

`func NewListBundleResponse() *ListBundleResponse`

NewListBundleResponse instantiates a new ListBundleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListBundleResponseWithDefaults

`func NewListBundleResponseWithDefaults() *ListBundleResponse`

NewListBundleResponseWithDefaults instantiates a new ListBundleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListBundleResponse) GetMeta() ListBundleResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListBundleResponse) GetMetaOk() (*ListBundleResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListBundleResponse) SetMeta(v ListBundleResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListBundleResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetResults

`func (o *ListBundleResponse) GetResults() []NumbersV2RegulatoryComplianceBundle`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ListBundleResponse) GetResultsOk() (*[]NumbersV2RegulatoryComplianceBundle, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ListBundleResponse) SetResults(v []NumbersV2RegulatoryComplianceBundle)`

SetResults sets Results field to given value.

### HasResults

`func (o *ListBundleResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


