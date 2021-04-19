# ListIncomingPhoneNumberLocalResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | Pointer to **int32** |  | [optional] 
**FirstPageUri** | Pointer to **string** |  | [optional] 
**IncomingPhoneNumbers** | Pointer to [**[]ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberLocal**](ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberLocal.md) |  | [optional] 
**NextPageUri** | Pointer to **string** |  | [optional] 
**Page** | Pointer to **int32** |  | [optional] 
**PageSize** | Pointer to **int32** |  | [optional] 
**PreviousPageUri** | Pointer to **string** |  | [optional] 
**Start** | Pointer to **int32** |  | [optional] 
**Uri** | Pointer to **string** |  | [optional] 

## Methods

### NewListIncomingPhoneNumberLocalResponse

`func NewListIncomingPhoneNumberLocalResponse() *ListIncomingPhoneNumberLocalResponse`

NewListIncomingPhoneNumberLocalResponse instantiates a new ListIncomingPhoneNumberLocalResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListIncomingPhoneNumberLocalResponseWithDefaults

`func NewListIncomingPhoneNumberLocalResponseWithDefaults() *ListIncomingPhoneNumberLocalResponse`

NewListIncomingPhoneNumberLocalResponseWithDefaults instantiates a new ListIncomingPhoneNumberLocalResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *ListIncomingPhoneNumberLocalResponse) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *ListIncomingPhoneNumberLocalResponse) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *ListIncomingPhoneNumberLocalResponse) SetEnd(v int32)`

SetEnd sets End field to given value.

### HasEnd

`func (o *ListIncomingPhoneNumberLocalResponse) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetFirstPageUri

`func (o *ListIncomingPhoneNumberLocalResponse) GetFirstPageUri() string`

GetFirstPageUri returns the FirstPageUri field if non-nil, zero value otherwise.

### GetFirstPageUriOk

`func (o *ListIncomingPhoneNumberLocalResponse) GetFirstPageUriOk() (*string, bool)`

GetFirstPageUriOk returns a tuple with the FirstPageUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstPageUri

`func (o *ListIncomingPhoneNumberLocalResponse) SetFirstPageUri(v string)`

SetFirstPageUri sets FirstPageUri field to given value.

### HasFirstPageUri

`func (o *ListIncomingPhoneNumberLocalResponse) HasFirstPageUri() bool`

HasFirstPageUri returns a boolean if a field has been set.

### GetIncomingPhoneNumbers

`func (o *ListIncomingPhoneNumberLocalResponse) GetIncomingPhoneNumbers() []ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberLocal`

GetIncomingPhoneNumbers returns the IncomingPhoneNumbers field if non-nil, zero value otherwise.

### GetIncomingPhoneNumbersOk

`func (o *ListIncomingPhoneNumberLocalResponse) GetIncomingPhoneNumbersOk() (*[]ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberLocal, bool)`

GetIncomingPhoneNumbersOk returns a tuple with the IncomingPhoneNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomingPhoneNumbers

`func (o *ListIncomingPhoneNumberLocalResponse) SetIncomingPhoneNumbers(v []ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberLocal)`

SetIncomingPhoneNumbers sets IncomingPhoneNumbers field to given value.

### HasIncomingPhoneNumbers

`func (o *ListIncomingPhoneNumberLocalResponse) HasIncomingPhoneNumbers() bool`

HasIncomingPhoneNumbers returns a boolean if a field has been set.

### GetNextPageUri

`func (o *ListIncomingPhoneNumberLocalResponse) GetNextPageUri() string`

GetNextPageUri returns the NextPageUri field if non-nil, zero value otherwise.

### GetNextPageUriOk

`func (o *ListIncomingPhoneNumberLocalResponse) GetNextPageUriOk() (*string, bool)`

GetNextPageUriOk returns a tuple with the NextPageUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageUri

`func (o *ListIncomingPhoneNumberLocalResponse) SetNextPageUri(v string)`

SetNextPageUri sets NextPageUri field to given value.

### HasNextPageUri

`func (o *ListIncomingPhoneNumberLocalResponse) HasNextPageUri() bool`

HasNextPageUri returns a boolean if a field has been set.

### GetPage

`func (o *ListIncomingPhoneNumberLocalResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ListIncomingPhoneNumberLocalResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ListIncomingPhoneNumberLocalResponse) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *ListIncomingPhoneNumberLocalResponse) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageSize

`func (o *ListIncomingPhoneNumberLocalResponse) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ListIncomingPhoneNumberLocalResponse) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ListIncomingPhoneNumberLocalResponse) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ListIncomingPhoneNumberLocalResponse) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetPreviousPageUri

`func (o *ListIncomingPhoneNumberLocalResponse) GetPreviousPageUri() string`

GetPreviousPageUri returns the PreviousPageUri field if non-nil, zero value otherwise.

### GetPreviousPageUriOk

`func (o *ListIncomingPhoneNumberLocalResponse) GetPreviousPageUriOk() (*string, bool)`

GetPreviousPageUriOk returns a tuple with the PreviousPageUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousPageUri

`func (o *ListIncomingPhoneNumberLocalResponse) SetPreviousPageUri(v string)`

SetPreviousPageUri sets PreviousPageUri field to given value.

### HasPreviousPageUri

`func (o *ListIncomingPhoneNumberLocalResponse) HasPreviousPageUri() bool`

HasPreviousPageUri returns a boolean if a field has been set.

### GetStart

`func (o *ListIncomingPhoneNumberLocalResponse) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ListIncomingPhoneNumberLocalResponse) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ListIncomingPhoneNumberLocalResponse) SetStart(v int32)`

SetStart sets Start field to given value.

### HasStart

`func (o *ListIncomingPhoneNumberLocalResponse) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetUri

`func (o *ListIncomingPhoneNumberLocalResponse) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ListIncomingPhoneNumberLocalResponse) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ListIncomingPhoneNumberLocalResponse) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *ListIncomingPhoneNumberLocalResponse) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


