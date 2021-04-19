# ListDocumentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Documents** | Pointer to [**[]SyncV1ServiceDocument**](SyncV1ServiceDocument.md) |  | [optional] 
**Meta** | Pointer to [**ListServiceResponseMeta**](ListServiceResponse_meta.md) |  | [optional] 

## Methods

### NewListDocumentResponse

`func NewListDocumentResponse() *ListDocumentResponse`

NewListDocumentResponse instantiates a new ListDocumentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDocumentResponseWithDefaults

`func NewListDocumentResponseWithDefaults() *ListDocumentResponse`

NewListDocumentResponseWithDefaults instantiates a new ListDocumentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocuments

`func (o *ListDocumentResponse) GetDocuments() []SyncV1ServiceDocument`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *ListDocumentResponse) GetDocumentsOk() (*[]SyncV1ServiceDocument, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *ListDocumentResponse) SetDocuments(v []SyncV1ServiceDocument)`

SetDocuments sets Documents field to given value.

### HasDocuments

`func (o *ListDocumentResponse) HasDocuments() bool`

HasDocuments returns a boolean if a field has been set.

### GetMeta

`func (o *ListDocumentResponse) GetMeta() ListServiceResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListDocumentResponse) GetMetaOk() (*ListServiceResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListDocumentResponse) SetMeta(v ListServiceResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListDocumentResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


