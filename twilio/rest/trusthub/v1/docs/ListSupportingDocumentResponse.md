# ListSupportingDocumentResponse

## Properties

Name | Type | Description
------------ | ------------- | -------------
**Meta** | Pointer to [**ListCustomerProfileResponseMeta**](ListCustomerProfileResponse_meta.md) |  | [optional] 
**Results** | Pointer to [**[]TrusthubV1SupportingDocument**](TrusthubV1SupportingDocument.md) |  | [optional] 

## Methods

### NewListSupportingDocumentResponse

`func NewListSupportingDocumentResponse() *ListSupportingDocumentResponse`

NewListSupportingDocumentResponse instantiates a new ListSupportingDocumentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSupportingDocumentResponseWithDefaults

`func NewListSupportingDocumentResponseWithDefaults() *ListSupportingDocumentResponse`

NewListSupportingDocumentResponseWithDefaults instantiates a new ListSupportingDocumentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListSupportingDocumentResponse) GetMeta() ListCustomerProfileResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListSupportingDocumentResponse) GetMetaOk() (*ListCustomerProfileResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListSupportingDocumentResponse) SetMeta(v ListCustomerProfileResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListSupportingDocumentResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetResults

`func (o *ListSupportingDocumentResponse) GetResults() []TrusthubV1SupportingDocument`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ListSupportingDocumentResponse) GetResultsOk() (*[]TrusthubV1SupportingDocument, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ListSupportingDocumentResponse) SetResults(v []TrusthubV1SupportingDocument)`

SetResults sets Results field to given value.

### HasResults

`func (o *ListSupportingDocumentResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


