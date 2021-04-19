# ListFaxMediaResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Media** | Pointer to [**[]FaxV1FaxFaxMedia**](FaxV1FaxFaxMedia.md) |  | [optional] 
**Meta** | Pointer to [**ListFaxResponseMeta**](ListFaxResponse_meta.md) |  | [optional] 

## Methods

### NewListFaxMediaResponse

`func NewListFaxMediaResponse() *ListFaxMediaResponse`

NewListFaxMediaResponse instantiates a new ListFaxMediaResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListFaxMediaResponseWithDefaults

`func NewListFaxMediaResponseWithDefaults() *ListFaxMediaResponse`

NewListFaxMediaResponseWithDefaults instantiates a new ListFaxMediaResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMedia

`func (o *ListFaxMediaResponse) GetMedia() []FaxV1FaxFaxMedia`

GetMedia returns the Media field if non-nil, zero value otherwise.

### GetMediaOk

`func (o *ListFaxMediaResponse) GetMediaOk() (*[]FaxV1FaxFaxMedia, bool)`

GetMediaOk returns a tuple with the Media field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedia

`func (o *ListFaxMediaResponse) SetMedia(v []FaxV1FaxFaxMedia)`

SetMedia sets Media field to given value.

### HasMedia

`func (o *ListFaxMediaResponse) HasMedia() bool`

HasMedia returns a boolean if a field has been set.

### GetMeta

`func (o *ListFaxMediaResponse) GetMeta() ListFaxResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListFaxMediaResponse) GetMetaOk() (*ListFaxResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListFaxMediaResponse) SetMeta(v ListFaxResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListFaxMediaResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


