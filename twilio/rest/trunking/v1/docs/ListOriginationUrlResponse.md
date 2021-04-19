# ListOriginationUrlResponse

## Properties

Name | Type | Description
------------ | ------------- | -------------
**Meta** | Pointer to [**ListTrunkResponseMeta**](ListTrunkResponse_meta.md) |  | [optional] 
**OriginationUrls** | Pointer to [**[]TrunkingV1TrunkOriginationUrl**](TrunkingV1TrunkOriginationUrl.md) |  | [optional] 

## Methods

### NewListOriginationUrlResponse

`func NewListOriginationUrlResponse() *ListOriginationUrlResponse`

NewListOriginationUrlResponse instantiates a new ListOriginationUrlResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListOriginationUrlResponseWithDefaults

`func NewListOriginationUrlResponseWithDefaults() *ListOriginationUrlResponse`

NewListOriginationUrlResponseWithDefaults instantiates a new ListOriginationUrlResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListOriginationUrlResponse) GetMeta() ListTrunkResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListOriginationUrlResponse) GetMetaOk() (*ListTrunkResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListOriginationUrlResponse) SetMeta(v ListTrunkResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListOriginationUrlResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetOriginationUrls

`func (o *ListOriginationUrlResponse) GetOriginationUrls() []TrunkingV1TrunkOriginationUrl`

GetOriginationUrls returns the OriginationUrls field if non-nil, zero value otherwise.

### GetOriginationUrlsOk

`func (o *ListOriginationUrlResponse) GetOriginationUrlsOk() (*[]TrunkingV1TrunkOriginationUrl, bool)`

GetOriginationUrlsOk returns a tuple with the OriginationUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginationUrls

`func (o *ListOriginationUrlResponse) SetOriginationUrls(v []TrunkingV1TrunkOriginationUrl)`

SetOriginationUrls sets OriginationUrls field to given value.

### HasOriginationUrls

`func (o *ListOriginationUrlResponse) HasOriginationUrls() bool`

HasOriginationUrls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


