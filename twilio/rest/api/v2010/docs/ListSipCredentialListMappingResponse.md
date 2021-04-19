# ListSipCredentialListMappingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CredentialListMappings** | Pointer to [**[]ApiV2010AccountSipSipDomainSipCredentialListMapping**](ApiV2010AccountSipSipDomainSipCredentialListMapping.md) |  | [optional] 
**End** | Pointer to **int32** |  | [optional] 
**FirstPageUri** | Pointer to **string** |  | [optional] 
**NextPageUri** | Pointer to **string** |  | [optional] 
**Page** | Pointer to **int32** |  | [optional] 
**PageSize** | Pointer to **int32** |  | [optional] 
**PreviousPageUri** | Pointer to **string** |  | [optional] 
**Start** | Pointer to **int32** |  | [optional] 
**Uri** | Pointer to **string** |  | [optional] 

## Methods

### NewListSipCredentialListMappingResponse

`func NewListSipCredentialListMappingResponse() *ListSipCredentialListMappingResponse`

NewListSipCredentialListMappingResponse instantiates a new ListSipCredentialListMappingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSipCredentialListMappingResponseWithDefaults

`func NewListSipCredentialListMappingResponseWithDefaults() *ListSipCredentialListMappingResponse`

NewListSipCredentialListMappingResponseWithDefaults instantiates a new ListSipCredentialListMappingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentialListMappings

`func (o *ListSipCredentialListMappingResponse) GetCredentialListMappings() []ApiV2010AccountSipSipDomainSipCredentialListMapping`

GetCredentialListMappings returns the CredentialListMappings field if non-nil, zero value otherwise.

### GetCredentialListMappingsOk

`func (o *ListSipCredentialListMappingResponse) GetCredentialListMappingsOk() (*[]ApiV2010AccountSipSipDomainSipCredentialListMapping, bool)`

GetCredentialListMappingsOk returns a tuple with the CredentialListMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialListMappings

`func (o *ListSipCredentialListMappingResponse) SetCredentialListMappings(v []ApiV2010AccountSipSipDomainSipCredentialListMapping)`

SetCredentialListMappings sets CredentialListMappings field to given value.

### HasCredentialListMappings

`func (o *ListSipCredentialListMappingResponse) HasCredentialListMappings() bool`

HasCredentialListMappings returns a boolean if a field has been set.

### GetEnd

`func (o *ListSipCredentialListMappingResponse) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *ListSipCredentialListMappingResponse) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *ListSipCredentialListMappingResponse) SetEnd(v int32)`

SetEnd sets End field to given value.

### HasEnd

`func (o *ListSipCredentialListMappingResponse) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetFirstPageUri

`func (o *ListSipCredentialListMappingResponse) GetFirstPageUri() string`

GetFirstPageUri returns the FirstPageUri field if non-nil, zero value otherwise.

### GetFirstPageUriOk

`func (o *ListSipCredentialListMappingResponse) GetFirstPageUriOk() (*string, bool)`

GetFirstPageUriOk returns a tuple with the FirstPageUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstPageUri

`func (o *ListSipCredentialListMappingResponse) SetFirstPageUri(v string)`

SetFirstPageUri sets FirstPageUri field to given value.

### HasFirstPageUri

`func (o *ListSipCredentialListMappingResponse) HasFirstPageUri() bool`

HasFirstPageUri returns a boolean if a field has been set.

### GetNextPageUri

`func (o *ListSipCredentialListMappingResponse) GetNextPageUri() string`

GetNextPageUri returns the NextPageUri field if non-nil, zero value otherwise.

### GetNextPageUriOk

`func (o *ListSipCredentialListMappingResponse) GetNextPageUriOk() (*string, bool)`

GetNextPageUriOk returns a tuple with the NextPageUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageUri

`func (o *ListSipCredentialListMappingResponse) SetNextPageUri(v string)`

SetNextPageUri sets NextPageUri field to given value.

### HasNextPageUri

`func (o *ListSipCredentialListMappingResponse) HasNextPageUri() bool`

HasNextPageUri returns a boolean if a field has been set.

### GetPage

`func (o *ListSipCredentialListMappingResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ListSipCredentialListMappingResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ListSipCredentialListMappingResponse) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *ListSipCredentialListMappingResponse) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageSize

`func (o *ListSipCredentialListMappingResponse) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ListSipCredentialListMappingResponse) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ListSipCredentialListMappingResponse) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ListSipCredentialListMappingResponse) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetPreviousPageUri

`func (o *ListSipCredentialListMappingResponse) GetPreviousPageUri() string`

GetPreviousPageUri returns the PreviousPageUri field if non-nil, zero value otherwise.

### GetPreviousPageUriOk

`func (o *ListSipCredentialListMappingResponse) GetPreviousPageUriOk() (*string, bool)`

GetPreviousPageUriOk returns a tuple with the PreviousPageUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousPageUri

`func (o *ListSipCredentialListMappingResponse) SetPreviousPageUri(v string)`

SetPreviousPageUri sets PreviousPageUri field to given value.

### HasPreviousPageUri

`func (o *ListSipCredentialListMappingResponse) HasPreviousPageUri() bool`

HasPreviousPageUri returns a boolean if a field has been set.

### GetStart

`func (o *ListSipCredentialListMappingResponse) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ListSipCredentialListMappingResponse) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ListSipCredentialListMappingResponse) SetStart(v int32)`

SetStart sets Start field to given value.

### HasStart

`func (o *ListSipCredentialListMappingResponse) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetUri

`func (o *ListSipCredentialListMappingResponse) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ListSipCredentialListMappingResponse) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ListSipCredentialListMappingResponse) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *ListSipCredentialListMappingResponse) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


