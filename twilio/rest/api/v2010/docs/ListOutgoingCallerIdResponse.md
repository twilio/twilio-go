# ListOutgoingCallerIdResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | Pointer to **int32** |  | [optional] 
**FirstPageUri** | Pointer to **string** |  | [optional] 
**NextPageUri** | Pointer to **string** |  | [optional] 
**OutgoingCallerIds** | Pointer to [**[]ApiV2010AccountOutgoingCallerId**](ApiV2010AccountOutgoingCallerId.md) |  | [optional] 
**Page** | Pointer to **int32** |  | [optional] 
**PageSize** | Pointer to **int32** |  | [optional] 
**PreviousPageUri** | Pointer to **string** |  | [optional] 
**Start** | Pointer to **int32** |  | [optional] 
**Uri** | Pointer to **string** |  | [optional] 

## Methods

### NewListOutgoingCallerIdResponse

`func NewListOutgoingCallerIdResponse() *ListOutgoingCallerIdResponse`

NewListOutgoingCallerIdResponse instantiates a new ListOutgoingCallerIdResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListOutgoingCallerIdResponseWithDefaults

`func NewListOutgoingCallerIdResponseWithDefaults() *ListOutgoingCallerIdResponse`

NewListOutgoingCallerIdResponseWithDefaults instantiates a new ListOutgoingCallerIdResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *ListOutgoingCallerIdResponse) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *ListOutgoingCallerIdResponse) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *ListOutgoingCallerIdResponse) SetEnd(v int32)`

SetEnd sets End field to given value.

### HasEnd

`func (o *ListOutgoingCallerIdResponse) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetFirstPageUri

`func (o *ListOutgoingCallerIdResponse) GetFirstPageUri() string`

GetFirstPageUri returns the FirstPageUri field if non-nil, zero value otherwise.

### GetFirstPageUriOk

`func (o *ListOutgoingCallerIdResponse) GetFirstPageUriOk() (*string, bool)`

GetFirstPageUriOk returns a tuple with the FirstPageUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstPageUri

`func (o *ListOutgoingCallerIdResponse) SetFirstPageUri(v string)`

SetFirstPageUri sets FirstPageUri field to given value.

### HasFirstPageUri

`func (o *ListOutgoingCallerIdResponse) HasFirstPageUri() bool`

HasFirstPageUri returns a boolean if a field has been set.

### GetNextPageUri

`func (o *ListOutgoingCallerIdResponse) GetNextPageUri() string`

GetNextPageUri returns the NextPageUri field if non-nil, zero value otherwise.

### GetNextPageUriOk

`func (o *ListOutgoingCallerIdResponse) GetNextPageUriOk() (*string, bool)`

GetNextPageUriOk returns a tuple with the NextPageUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageUri

`func (o *ListOutgoingCallerIdResponse) SetNextPageUri(v string)`

SetNextPageUri sets NextPageUri field to given value.

### HasNextPageUri

`func (o *ListOutgoingCallerIdResponse) HasNextPageUri() bool`

HasNextPageUri returns a boolean if a field has been set.

### GetOutgoingCallerIds

`func (o *ListOutgoingCallerIdResponse) GetOutgoingCallerIds() []ApiV2010AccountOutgoingCallerId`

GetOutgoingCallerIds returns the OutgoingCallerIds field if non-nil, zero value otherwise.

### GetOutgoingCallerIdsOk

`func (o *ListOutgoingCallerIdResponse) GetOutgoingCallerIdsOk() (*[]ApiV2010AccountOutgoingCallerId, bool)`

GetOutgoingCallerIdsOk returns a tuple with the OutgoingCallerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutgoingCallerIds

`func (o *ListOutgoingCallerIdResponse) SetOutgoingCallerIds(v []ApiV2010AccountOutgoingCallerId)`

SetOutgoingCallerIds sets OutgoingCallerIds field to given value.

### HasOutgoingCallerIds

`func (o *ListOutgoingCallerIdResponse) HasOutgoingCallerIds() bool`

HasOutgoingCallerIds returns a boolean if a field has been set.

### GetPage

`func (o *ListOutgoingCallerIdResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ListOutgoingCallerIdResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ListOutgoingCallerIdResponse) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *ListOutgoingCallerIdResponse) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageSize

`func (o *ListOutgoingCallerIdResponse) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ListOutgoingCallerIdResponse) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ListOutgoingCallerIdResponse) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ListOutgoingCallerIdResponse) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetPreviousPageUri

`func (o *ListOutgoingCallerIdResponse) GetPreviousPageUri() string`

GetPreviousPageUri returns the PreviousPageUri field if non-nil, zero value otherwise.

### GetPreviousPageUriOk

`func (o *ListOutgoingCallerIdResponse) GetPreviousPageUriOk() (*string, bool)`

GetPreviousPageUriOk returns a tuple with the PreviousPageUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousPageUri

`func (o *ListOutgoingCallerIdResponse) SetPreviousPageUri(v string)`

SetPreviousPageUri sets PreviousPageUri field to given value.

### HasPreviousPageUri

`func (o *ListOutgoingCallerIdResponse) HasPreviousPageUri() bool`

HasPreviousPageUri returns a boolean if a field has been set.

### GetStart

`func (o *ListOutgoingCallerIdResponse) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ListOutgoingCallerIdResponse) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ListOutgoingCallerIdResponse) SetStart(v int32)`

SetStart sets Start field to given value.

### HasStart

`func (o *ListOutgoingCallerIdResponse) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetUri

`func (o *ListOutgoingCallerIdResponse) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ListOutgoingCallerIdResponse) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ListOutgoingCallerIdResponse) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *ListOutgoingCallerIdResponse) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


