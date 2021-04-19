# ListIncomingPhoneNumberMobileResponse

## Properties

Name | Type | Description
------------ | ------------- | -------------
**End** | Pointer to **int32** |  | [optional] 
**FirstPageUri** | Pointer to **string** |  | [optional] 
**IncomingPhoneNumbers** | Pointer to [**[]ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile**](ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile.md) |  | [optional] 
**NextPageUri** | Pointer to **string** |  | [optional] 
**Page** | Pointer to **int32** |  | [optional] 
**PageSize** | Pointer to **int32** |  | [optional] 
**PreviousPageUri** | Pointer to **string** |  | [optional] 
**Start** | Pointer to **int32** |  | [optional] 
**Uri** | Pointer to **string** |  | [optional] 

## Methods

### NewListIncomingPhoneNumberMobileResponse

`func NewListIncomingPhoneNumberMobileResponse() *ListIncomingPhoneNumberMobileResponse`

NewListIncomingPhoneNumberMobileResponse instantiates a new ListIncomingPhoneNumberMobileResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListIncomingPhoneNumberMobileResponseWithDefaults

`func NewListIncomingPhoneNumberMobileResponseWithDefaults() *ListIncomingPhoneNumberMobileResponse`

NewListIncomingPhoneNumberMobileResponseWithDefaults instantiates a new ListIncomingPhoneNumberMobileResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *ListIncomingPhoneNumberMobileResponse) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *ListIncomingPhoneNumberMobileResponse) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *ListIncomingPhoneNumberMobileResponse) SetEnd(v int32)`

SetEnd sets End field to given value.

### HasEnd

`func (o *ListIncomingPhoneNumberMobileResponse) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetFirstPageUri

`func (o *ListIncomingPhoneNumberMobileResponse) GetFirstPageUri() string`

GetFirstPageUri returns the FirstPageUri field if non-nil, zero value otherwise.

### GetFirstPageUriOk

`func (o *ListIncomingPhoneNumberMobileResponse) GetFirstPageUriOk() (*string, bool)`

GetFirstPageUriOk returns a tuple with the FirstPageUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstPageUri

`func (o *ListIncomingPhoneNumberMobileResponse) SetFirstPageUri(v string)`

SetFirstPageUri sets FirstPageUri field to given value.

### HasFirstPageUri

`func (o *ListIncomingPhoneNumberMobileResponse) HasFirstPageUri() bool`

HasFirstPageUri returns a boolean if a field has been set.

### GetIncomingPhoneNumbers

`func (o *ListIncomingPhoneNumberMobileResponse) GetIncomingPhoneNumbers() []ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile`

GetIncomingPhoneNumbers returns the IncomingPhoneNumbers field if non-nil, zero value otherwise.

### GetIncomingPhoneNumbersOk

`func (o *ListIncomingPhoneNumberMobileResponse) GetIncomingPhoneNumbersOk() (*[]ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile, bool)`

GetIncomingPhoneNumbersOk returns a tuple with the IncomingPhoneNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomingPhoneNumbers

`func (o *ListIncomingPhoneNumberMobileResponse) SetIncomingPhoneNumbers(v []ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile)`

SetIncomingPhoneNumbers sets IncomingPhoneNumbers field to given value.

### HasIncomingPhoneNumbers

`func (o *ListIncomingPhoneNumberMobileResponse) HasIncomingPhoneNumbers() bool`

HasIncomingPhoneNumbers returns a boolean if a field has been set.

### GetNextPageUri

`func (o *ListIncomingPhoneNumberMobileResponse) GetNextPageUri() string`

GetNextPageUri returns the NextPageUri field if non-nil, zero value otherwise.

### GetNextPageUriOk

`func (o *ListIncomingPhoneNumberMobileResponse) GetNextPageUriOk() (*string, bool)`

GetNextPageUriOk returns a tuple with the NextPageUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageUri

`func (o *ListIncomingPhoneNumberMobileResponse) SetNextPageUri(v string)`

SetNextPageUri sets NextPageUri field to given value.

### HasNextPageUri

`func (o *ListIncomingPhoneNumberMobileResponse) HasNextPageUri() bool`

HasNextPageUri returns a boolean if a field has been set.

### GetPage

`func (o *ListIncomingPhoneNumberMobileResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ListIncomingPhoneNumberMobileResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ListIncomingPhoneNumberMobileResponse) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *ListIncomingPhoneNumberMobileResponse) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageSize

`func (o *ListIncomingPhoneNumberMobileResponse) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ListIncomingPhoneNumberMobileResponse) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ListIncomingPhoneNumberMobileResponse) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ListIncomingPhoneNumberMobileResponse) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetPreviousPageUri

`func (o *ListIncomingPhoneNumberMobileResponse) GetPreviousPageUri() string`

GetPreviousPageUri returns the PreviousPageUri field if non-nil, zero value otherwise.

### GetPreviousPageUriOk

`func (o *ListIncomingPhoneNumberMobileResponse) GetPreviousPageUriOk() (*string, bool)`

GetPreviousPageUriOk returns a tuple with the PreviousPageUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousPageUri

`func (o *ListIncomingPhoneNumberMobileResponse) SetPreviousPageUri(v string)`

SetPreviousPageUri sets PreviousPageUri field to given value.

### HasPreviousPageUri

`func (o *ListIncomingPhoneNumberMobileResponse) HasPreviousPageUri() bool`

HasPreviousPageUri returns a boolean if a field has been set.

### GetStart

`func (o *ListIncomingPhoneNumberMobileResponse) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ListIncomingPhoneNumberMobileResponse) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ListIncomingPhoneNumberMobileResponse) SetStart(v int32)`

SetStart sets Start field to given value.

### HasStart

`func (o *ListIncomingPhoneNumberMobileResponse) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetUri

`func (o *ListIncomingPhoneNumberMobileResponse) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ListIncomingPhoneNumberMobileResponse) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ListIncomingPhoneNumberMobileResponse) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *ListIncomingPhoneNumberMobileResponse) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


