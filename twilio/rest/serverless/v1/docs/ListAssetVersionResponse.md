# ListAssetVersionResponse

## Properties

Name | Type | Description
------------ | ------------- | -------------
**AssetVersions** | Pointer to [**[]ServerlessV1ServiceAssetAssetVersion**](ServerlessV1ServiceAssetAssetVersion.md) |  | [optional] 
**Meta** | Pointer to [**ListServiceResponseMeta**](ListServiceResponse_meta.md) |  | [optional] 

## Methods

### NewListAssetVersionResponse

`func NewListAssetVersionResponse() *ListAssetVersionResponse`

NewListAssetVersionResponse instantiates a new ListAssetVersionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAssetVersionResponseWithDefaults

`func NewListAssetVersionResponseWithDefaults() *ListAssetVersionResponse`

NewListAssetVersionResponseWithDefaults instantiates a new ListAssetVersionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetVersions

`func (o *ListAssetVersionResponse) GetAssetVersions() []ServerlessV1ServiceAssetAssetVersion`

GetAssetVersions returns the AssetVersions field if non-nil, zero value otherwise.

### GetAssetVersionsOk

`func (o *ListAssetVersionResponse) GetAssetVersionsOk() (*[]ServerlessV1ServiceAssetAssetVersion, bool)`

GetAssetVersionsOk returns a tuple with the AssetVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetVersions

`func (o *ListAssetVersionResponse) SetAssetVersions(v []ServerlessV1ServiceAssetAssetVersion)`

SetAssetVersions sets AssetVersions field to given value.

### HasAssetVersions

`func (o *ListAssetVersionResponse) HasAssetVersions() bool`

HasAssetVersions returns a boolean if a field has been set.

### GetMeta

`func (o *ListAssetVersionResponse) GetMeta() ListServiceResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListAssetVersionResponse) GetMetaOk() (*ListServiceResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListAssetVersionResponse) SetMeta(v ListServiceResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListAssetVersionResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


