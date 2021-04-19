# ListShortCodeResponse

## Properties

Name | Type | Description
------------ | ------------- | -------------
**Meta** | Pointer to [**ListServiceResponseMeta**](ListServiceResponse_meta.md) |  | [optional] 
**ShortCodes** | Pointer to [**[]MessagingV1ServiceShortCode**](MessagingV1ServiceShortCode.md) |  | [optional] 

## Methods

### NewListShortCodeResponse

`func NewListShortCodeResponse() *ListShortCodeResponse`

NewListShortCodeResponse instantiates a new ListShortCodeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListShortCodeResponseWithDefaults

`func NewListShortCodeResponseWithDefaults() *ListShortCodeResponse`

NewListShortCodeResponseWithDefaults instantiates a new ListShortCodeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListShortCodeResponse) GetMeta() ListServiceResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListShortCodeResponse) GetMetaOk() (*ListServiceResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListShortCodeResponse) SetMeta(v ListServiceResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListShortCodeResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetShortCodes

`func (o *ListShortCodeResponse) GetShortCodes() []MessagingV1ServiceShortCode`

GetShortCodes returns the ShortCodes field if non-nil, zero value otherwise.

### GetShortCodesOk

`func (o *ListShortCodeResponse) GetShortCodesOk() (*[]MessagingV1ServiceShortCode, bool)`

GetShortCodesOk returns a tuple with the ShortCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortCodes

`func (o *ListShortCodeResponse) SetShortCodes(v []MessagingV1ServiceShortCode)`

SetShortCodes sets ShortCodes field to given value.

### HasShortCodes

`func (o *ListShortCodeResponse) HasShortCodes() bool`

HasShortCodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


