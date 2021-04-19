# ListSipIpAccessControlListMappingResponse

## Properties

Name | Type | Description
------------ | ------------- | -------------
**End** | Pointer to **int32** |  | [optional] 
**FirstPageUri** | Pointer to **string** |  | [optional] 
**IpAccessControlListMappings** | Pointer to [**[]ApiV2010AccountSipSipDomainSipIpAccessControlListMapping**](ApiV2010AccountSipSipDomainSipIpAccessControlListMapping.md) |  | [optional] 
**NextPageUri** | Pointer to **string** |  | [optional] 
**Page** | Pointer to **int32** |  | [optional] 
**PageSize** | Pointer to **int32** |  | [optional] 
**PreviousPageUri** | Pointer to **string** |  | [optional] 
**Start** | Pointer to **int32** |  | [optional] 
**Uri** | Pointer to **string** |  | [optional] 

## Methods

### NewListSipIpAccessControlListMappingResponse

`func NewListSipIpAccessControlListMappingResponse() *ListSipIpAccessControlListMappingResponse`

NewListSipIpAccessControlListMappingResponse instantiates a new ListSipIpAccessControlListMappingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSipIpAccessControlListMappingResponseWithDefaults

`func NewListSipIpAccessControlListMappingResponseWithDefaults() *ListSipIpAccessControlListMappingResponse`

NewListSipIpAccessControlListMappingResponseWithDefaults instantiates a new ListSipIpAccessControlListMappingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *ListSipIpAccessControlListMappingResponse) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *ListSipIpAccessControlListMappingResponse) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *ListSipIpAccessControlListMappingResponse) SetEnd(v int32)`

SetEnd sets End field to given value.

### HasEnd

`func (o *ListSipIpAccessControlListMappingResponse) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetFirstPageUri

`func (o *ListSipIpAccessControlListMappingResponse) GetFirstPageUri() string`

GetFirstPageUri returns the FirstPageUri field if non-nil, zero value otherwise.

### GetFirstPageUriOk

`func (o *ListSipIpAccessControlListMappingResponse) GetFirstPageUriOk() (*string, bool)`

GetFirstPageUriOk returns a tuple with the FirstPageUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstPageUri

`func (o *ListSipIpAccessControlListMappingResponse) SetFirstPageUri(v string)`

SetFirstPageUri sets FirstPageUri field to given value.

### HasFirstPageUri

`func (o *ListSipIpAccessControlListMappingResponse) HasFirstPageUri() bool`

HasFirstPageUri returns a boolean if a field has been set.

### GetIpAccessControlListMappings

`func (o *ListSipIpAccessControlListMappingResponse) GetIpAccessControlListMappings() []ApiV2010AccountSipSipDomainSipIpAccessControlListMapping`

GetIpAccessControlListMappings returns the IpAccessControlListMappings field if non-nil, zero value otherwise.

### GetIpAccessControlListMappingsOk

`func (o *ListSipIpAccessControlListMappingResponse) GetIpAccessControlListMappingsOk() (*[]ApiV2010AccountSipSipDomainSipIpAccessControlListMapping, bool)`

GetIpAccessControlListMappingsOk returns a tuple with the IpAccessControlListMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAccessControlListMappings

`func (o *ListSipIpAccessControlListMappingResponse) SetIpAccessControlListMappings(v []ApiV2010AccountSipSipDomainSipIpAccessControlListMapping)`

SetIpAccessControlListMappings sets IpAccessControlListMappings field to given value.

### HasIpAccessControlListMappings

`func (o *ListSipIpAccessControlListMappingResponse) HasIpAccessControlListMappings() bool`

HasIpAccessControlListMappings returns a boolean if a field has been set.

### GetNextPageUri

`func (o *ListSipIpAccessControlListMappingResponse) GetNextPageUri() string`

GetNextPageUri returns the NextPageUri field if non-nil, zero value otherwise.

### GetNextPageUriOk

`func (o *ListSipIpAccessControlListMappingResponse) GetNextPageUriOk() (*string, bool)`

GetNextPageUriOk returns a tuple with the NextPageUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageUri

`func (o *ListSipIpAccessControlListMappingResponse) SetNextPageUri(v string)`

SetNextPageUri sets NextPageUri field to given value.

### HasNextPageUri

`func (o *ListSipIpAccessControlListMappingResponse) HasNextPageUri() bool`

HasNextPageUri returns a boolean if a field has been set.

### GetPage

`func (o *ListSipIpAccessControlListMappingResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ListSipIpAccessControlListMappingResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ListSipIpAccessControlListMappingResponse) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *ListSipIpAccessControlListMappingResponse) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageSize

`func (o *ListSipIpAccessControlListMappingResponse) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ListSipIpAccessControlListMappingResponse) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ListSipIpAccessControlListMappingResponse) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ListSipIpAccessControlListMappingResponse) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetPreviousPageUri

`func (o *ListSipIpAccessControlListMappingResponse) GetPreviousPageUri() string`

GetPreviousPageUri returns the PreviousPageUri field if non-nil, zero value otherwise.

### GetPreviousPageUriOk

`func (o *ListSipIpAccessControlListMappingResponse) GetPreviousPageUriOk() (*string, bool)`

GetPreviousPageUriOk returns a tuple with the PreviousPageUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousPageUri

`func (o *ListSipIpAccessControlListMappingResponse) SetPreviousPageUri(v string)`

SetPreviousPageUri sets PreviousPageUri field to given value.

### HasPreviousPageUri

`func (o *ListSipIpAccessControlListMappingResponse) HasPreviousPageUri() bool`

HasPreviousPageUri returns a boolean if a field has been set.

### GetStart

`func (o *ListSipIpAccessControlListMappingResponse) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ListSipIpAccessControlListMappingResponse) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ListSipIpAccessControlListMappingResponse) SetStart(v int32)`

SetStart sets Start field to given value.

### HasStart

`func (o *ListSipIpAccessControlListMappingResponse) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetUri

`func (o *ListSipIpAccessControlListMappingResponse) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ListSipIpAccessControlListMappingResponse) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ListSipIpAccessControlListMappingResponse) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *ListSipIpAccessControlListMappingResponse) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


