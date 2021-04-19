# ListAlphaSenderResponse

## Properties

Name | Type | Description
------------ | ------------- | -------------
**AlphaSenders** | Pointer to [**[]MessagingV1ServiceAlphaSender**](MessagingV1ServiceAlphaSender.md) |  | [optional] 
**Meta** | Pointer to [**ListServiceResponseMeta**](ListServiceResponse_meta.md) |  | [optional] 

## Methods

### NewListAlphaSenderResponse

`func NewListAlphaSenderResponse() *ListAlphaSenderResponse`

NewListAlphaSenderResponse instantiates a new ListAlphaSenderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAlphaSenderResponseWithDefaults

`func NewListAlphaSenderResponseWithDefaults() *ListAlphaSenderResponse`

NewListAlphaSenderResponseWithDefaults instantiates a new ListAlphaSenderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlphaSenders

`func (o *ListAlphaSenderResponse) GetAlphaSenders() []MessagingV1ServiceAlphaSender`

GetAlphaSenders returns the AlphaSenders field if non-nil, zero value otherwise.

### GetAlphaSendersOk

`func (o *ListAlphaSenderResponse) GetAlphaSendersOk() (*[]MessagingV1ServiceAlphaSender, bool)`

GetAlphaSendersOk returns a tuple with the AlphaSenders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlphaSenders

`func (o *ListAlphaSenderResponse) SetAlphaSenders(v []MessagingV1ServiceAlphaSender)`

SetAlphaSenders sets AlphaSenders field to given value.

### HasAlphaSenders

`func (o *ListAlphaSenderResponse) HasAlphaSenders() bool`

HasAlphaSenders returns a boolean if a field has been set.

### GetMeta

`func (o *ListAlphaSenderResponse) GetMeta() ListServiceResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListAlphaSenderResponse) GetMetaOk() (*ListServiceResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListAlphaSenderResponse) SetMeta(v ListServiceResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListAlphaSenderResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


