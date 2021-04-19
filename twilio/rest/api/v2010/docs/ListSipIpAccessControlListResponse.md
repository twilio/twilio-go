# ListSipIpAccessControlListResponse

## Properties

Name | Type | Description
------------ | ------------- | -------------
**End** | Pointer to **int32** |  | [optional] 
**FirstPageUri** | Pointer to **string** |  | [optional] 
**IpAccessControlLists** | Pointer to [**[]ApiV2010AccountSipSipIpAccessControlList**](ApiV2010AccountSipSipIpAccessControlList.md) |  | [optional] 
**NextPageUri** | Pointer to **string** |  | [optional] 
**Page** | Pointer to **int32** |  | [optional] 
**PageSize** | Pointer to **int32** |  | [optional] 
**PreviousPageUri** | Pointer to **string** |  | [optional] 
**Start** | Pointer to **int32** |  | [optional] 
**Uri** | Pointer to **string** |  | [optional] 

## Methods

### NewListSipIpAccessControlListResponse

`func NewListSipIpAccessControlListResponse() *ListSipIpAccessControlListResponse`

NewListSipIpAccessControlListResponse instantiates a new ListSipIpAccessControlListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSipIpAccessControlListResponseWithDefaults

`func NewListSipIpAccessControlListResponseWithDefaults() *ListSipIpAccessControlListResponse`

NewListSipIpAccessControlListResponseWithDefaults instantiates a new ListSipIpAccessControlListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *ListSipIpAccessControlListResponse) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *ListSipIpAccessControlListResponse) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *ListSipIpAccessControlListResponse) SetEnd(v int32)`

SetEnd sets End field to given value.

### HasEnd

`func (o *ListSipIpAccessControlListResponse) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetFirstPageUri

`func (o *ListSipIpAccessControlListResponse) GetFirstPageUri() string`

GetFirstPageUri returns the FirstPageUri field if non-nil, zero value otherwise.

### GetFirstPageUriOk

`func (o *ListSipIpAccessControlListResponse) GetFirstPageUriOk() (*string, bool)`

GetFirstPageUriOk returns a tuple with the FirstPageUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstPageUri

`func (o *ListSipIpAccessControlListResponse) SetFirstPageUri(v string)`

SetFirstPageUri sets FirstPageUri field to given value.

### HasFirstPageUri

`func (o *ListSipIpAccessControlListResponse) HasFirstPageUri() bool`

HasFirstPageUri returns a boolean if a field has been set.

### GetIpAccessControlLists

`func (o *ListSipIpAccessControlListResponse) GetIpAccessControlLists() []ApiV2010AccountSipSipIpAccessControlList`

GetIpAccessControlLists returns the IpAccessControlLists field if non-nil, zero value otherwise.

### GetIpAccessControlListsOk

`func (o *ListSipIpAccessControlListResponse) GetIpAccessControlListsOk() (*[]ApiV2010AccountSipSipIpAccessControlList, bool)`

GetIpAccessControlListsOk returns a tuple with the IpAccessControlLists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAccessControlLists

`func (o *ListSipIpAccessControlListResponse) SetIpAccessControlLists(v []ApiV2010AccountSipSipIpAccessControlList)`

SetIpAccessControlLists sets IpAccessControlLists field to given value.

### HasIpAccessControlLists

`func (o *ListSipIpAccessControlListResponse) HasIpAccessControlLists() bool`

HasIpAccessControlLists returns a boolean if a field has been set.

### GetNextPageUri

`func (o *ListSipIpAccessControlListResponse) GetNextPageUri() string`

GetNextPageUri returns the NextPageUri field if non-nil, zero value otherwise.

### GetNextPageUriOk

`func (o *ListSipIpAccessControlListResponse) GetNextPageUriOk() (*string, bool)`

GetNextPageUriOk returns a tuple with the NextPageUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageUri

`func (o *ListSipIpAccessControlListResponse) SetNextPageUri(v string)`

SetNextPageUri sets NextPageUri field to given value.

### HasNextPageUri

`func (o *ListSipIpAccessControlListResponse) HasNextPageUri() bool`

HasNextPageUri returns a boolean if a field has been set.

### GetPage

`func (o *ListSipIpAccessControlListResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ListSipIpAccessControlListResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ListSipIpAccessControlListResponse) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *ListSipIpAccessControlListResponse) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageSize

`func (o *ListSipIpAccessControlListResponse) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ListSipIpAccessControlListResponse) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ListSipIpAccessControlListResponse) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ListSipIpAccessControlListResponse) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetPreviousPageUri

`func (o *ListSipIpAccessControlListResponse) GetPreviousPageUri() string`

GetPreviousPageUri returns the PreviousPageUri field if non-nil, zero value otherwise.

### GetPreviousPageUriOk

`func (o *ListSipIpAccessControlListResponse) GetPreviousPageUriOk() (*string, bool)`

GetPreviousPageUriOk returns a tuple with the PreviousPageUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousPageUri

`func (o *ListSipIpAccessControlListResponse) SetPreviousPageUri(v string)`

SetPreviousPageUri sets PreviousPageUri field to given value.

### HasPreviousPageUri

`func (o *ListSipIpAccessControlListResponse) HasPreviousPageUri() bool`

HasPreviousPageUri returns a boolean if a field has been set.

### GetStart

`func (o *ListSipIpAccessControlListResponse) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ListSipIpAccessControlListResponse) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ListSipIpAccessControlListResponse) SetStart(v int32)`

SetStart sets Start field to given value.

### HasStart

`func (o *ListSipIpAccessControlListResponse) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetUri

`func (o *ListSipIpAccessControlListResponse) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ListSipIpAccessControlListResponse) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ListSipIpAccessControlListResponse) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *ListSipIpAccessControlListResponse) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


