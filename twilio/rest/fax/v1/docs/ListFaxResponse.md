# ListFaxResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Faxes** | Pointer to [**[]FaxV1Fax**](FaxV1Fax.md) |  | [optional] 
**Meta** | Pointer to [**ListFaxResponseMeta**](ListFaxResponse_meta.md) |  | [optional] 

## Methods

### NewListFaxResponse

`func NewListFaxResponse() *ListFaxResponse`

NewListFaxResponse instantiates a new ListFaxResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListFaxResponseWithDefaults

`func NewListFaxResponseWithDefaults() *ListFaxResponse`

NewListFaxResponseWithDefaults instantiates a new ListFaxResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFaxes

`func (o *ListFaxResponse) GetFaxes() []FaxV1Fax`

GetFaxes returns the Faxes field if non-nil, zero value otherwise.

### GetFaxesOk

`func (o *ListFaxResponse) GetFaxesOk() (*[]FaxV1Fax, bool)`

GetFaxesOk returns a tuple with the Faxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaxes

`func (o *ListFaxResponse) SetFaxes(v []FaxV1Fax)`

SetFaxes sets Faxes field to given value.

### HasFaxes

`func (o *ListFaxResponse) HasFaxes() bool`

HasFaxes returns a boolean if a field has been set.

### GetMeta

`func (o *ListFaxResponse) GetMeta() ListFaxResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListFaxResponse) GetMetaOk() (*ListFaxResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListFaxResponse) SetMeta(v ListFaxResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListFaxResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


