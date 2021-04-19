# ListAssetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assets** | Pointer to [**[]ServerlessV1ServiceAsset**](ServerlessV1ServiceAsset.md) |  | [optional] 
**Meta** | Pointer to [**ListServiceResponseMeta**](ListServiceResponse_meta.md) |  | [optional] 

## Methods

### NewListAssetResponse

`func NewListAssetResponse() *ListAssetResponse`

NewListAssetResponse instantiates a new ListAssetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAssetResponseWithDefaults

`func NewListAssetResponseWithDefaults() *ListAssetResponse`

NewListAssetResponseWithDefaults instantiates a new ListAssetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssets

`func (o *ListAssetResponse) GetAssets() []ServerlessV1ServiceAsset`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *ListAssetResponse) GetAssetsOk() (*[]ServerlessV1ServiceAsset, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *ListAssetResponse) SetAssets(v []ServerlessV1ServiceAsset)`

SetAssets sets Assets field to given value.

### HasAssets

`func (o *ListAssetResponse) HasAssets() bool`

HasAssets returns a boolean if a field has been set.

### GetMeta

`func (o *ListAssetResponse) GetMeta() ListServiceResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListAssetResponse) GetMetaOk() (*ListServiceResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListAssetResponse) SetMeta(v ListServiceResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListAssetResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


