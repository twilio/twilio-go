# ListSupportingDocumentTypeResponse

## Properties

Name | Type | Description
------------ | ------------- | -------------
**Meta** | Pointer to [**ListBundleResponseMeta**](ListBundleResponse_meta.md) |  | [optional] 
**SupportingDocumentTypes** | Pointer to [**[]NumbersV2RegulatoryComplianceSupportingDocumentType**](NumbersV2RegulatoryComplianceSupportingDocumentType.md) |  | [optional] 

## Methods

### NewListSupportingDocumentTypeResponse

`func NewListSupportingDocumentTypeResponse() *ListSupportingDocumentTypeResponse`

NewListSupportingDocumentTypeResponse instantiates a new ListSupportingDocumentTypeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSupportingDocumentTypeResponseWithDefaults

`func NewListSupportingDocumentTypeResponseWithDefaults() *ListSupportingDocumentTypeResponse`

NewListSupportingDocumentTypeResponseWithDefaults instantiates a new ListSupportingDocumentTypeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListSupportingDocumentTypeResponse) GetMeta() ListBundleResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListSupportingDocumentTypeResponse) GetMetaOk() (*ListBundleResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListSupportingDocumentTypeResponse) SetMeta(v ListBundleResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListSupportingDocumentTypeResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetSupportingDocumentTypes

`func (o *ListSupportingDocumentTypeResponse) GetSupportingDocumentTypes() []NumbersV2RegulatoryComplianceSupportingDocumentType`

GetSupportingDocumentTypes returns the SupportingDocumentTypes field if non-nil, zero value otherwise.

### GetSupportingDocumentTypesOk

`func (o *ListSupportingDocumentTypeResponse) GetSupportingDocumentTypesOk() (*[]NumbersV2RegulatoryComplianceSupportingDocumentType, bool)`

GetSupportingDocumentTypesOk returns a tuple with the SupportingDocumentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportingDocumentTypes

`func (o *ListSupportingDocumentTypeResponse) SetSupportingDocumentTypes(v []NumbersV2RegulatoryComplianceSupportingDocumentType)`

SetSupportingDocumentTypes sets SupportingDocumentTypes field to given value.

### HasSupportingDocumentTypes

`func (o *ListSupportingDocumentTypeResponse) HasSupportingDocumentTypes() bool`

HasSupportingDocumentTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


