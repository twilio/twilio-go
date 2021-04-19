# ListRecordingAddOnResultResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddOnResults** | Pointer to [**[]ApiV2010AccountRecordingRecordingAddOnResult**](ApiV2010AccountRecordingRecordingAddOnResult.md) |  | [optional] 
**End** | Pointer to **int32** |  | [optional] 
**FirstPageUri** | Pointer to **string** |  | [optional] 
**NextPageUri** | Pointer to **string** |  | [optional] 
**Page** | Pointer to **int32** |  | [optional] 
**PageSize** | Pointer to **int32** |  | [optional] 
**PreviousPageUri** | Pointer to **string** |  | [optional] 
**Start** | Pointer to **int32** |  | [optional] 
**Uri** | Pointer to **string** |  | [optional] 

## Methods

### NewListRecordingAddOnResultResponse

`func NewListRecordingAddOnResultResponse() *ListRecordingAddOnResultResponse`

NewListRecordingAddOnResultResponse instantiates a new ListRecordingAddOnResultResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListRecordingAddOnResultResponseWithDefaults

`func NewListRecordingAddOnResultResponseWithDefaults() *ListRecordingAddOnResultResponse`

NewListRecordingAddOnResultResponseWithDefaults instantiates a new ListRecordingAddOnResultResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddOnResults

`func (o *ListRecordingAddOnResultResponse) GetAddOnResults() []ApiV2010AccountRecordingRecordingAddOnResult`

GetAddOnResults returns the AddOnResults field if non-nil, zero value otherwise.

### GetAddOnResultsOk

`func (o *ListRecordingAddOnResultResponse) GetAddOnResultsOk() (*[]ApiV2010AccountRecordingRecordingAddOnResult, bool)`

GetAddOnResultsOk returns a tuple with the AddOnResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddOnResults

`func (o *ListRecordingAddOnResultResponse) SetAddOnResults(v []ApiV2010AccountRecordingRecordingAddOnResult)`

SetAddOnResults sets AddOnResults field to given value.

### HasAddOnResults

`func (o *ListRecordingAddOnResultResponse) HasAddOnResults() bool`

HasAddOnResults returns a boolean if a field has been set.

### GetEnd

`func (o *ListRecordingAddOnResultResponse) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *ListRecordingAddOnResultResponse) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *ListRecordingAddOnResultResponse) SetEnd(v int32)`

SetEnd sets End field to given value.

### HasEnd

`func (o *ListRecordingAddOnResultResponse) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetFirstPageUri

`func (o *ListRecordingAddOnResultResponse) GetFirstPageUri() string`

GetFirstPageUri returns the FirstPageUri field if non-nil, zero value otherwise.

### GetFirstPageUriOk

`func (o *ListRecordingAddOnResultResponse) GetFirstPageUriOk() (*string, bool)`

GetFirstPageUriOk returns a tuple with the FirstPageUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstPageUri

`func (o *ListRecordingAddOnResultResponse) SetFirstPageUri(v string)`

SetFirstPageUri sets FirstPageUri field to given value.

### HasFirstPageUri

`func (o *ListRecordingAddOnResultResponse) HasFirstPageUri() bool`

HasFirstPageUri returns a boolean if a field has been set.

### GetNextPageUri

`func (o *ListRecordingAddOnResultResponse) GetNextPageUri() string`

GetNextPageUri returns the NextPageUri field if non-nil, zero value otherwise.

### GetNextPageUriOk

`func (o *ListRecordingAddOnResultResponse) GetNextPageUriOk() (*string, bool)`

GetNextPageUriOk returns a tuple with the NextPageUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageUri

`func (o *ListRecordingAddOnResultResponse) SetNextPageUri(v string)`

SetNextPageUri sets NextPageUri field to given value.

### HasNextPageUri

`func (o *ListRecordingAddOnResultResponse) HasNextPageUri() bool`

HasNextPageUri returns a boolean if a field has been set.

### GetPage

`func (o *ListRecordingAddOnResultResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ListRecordingAddOnResultResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ListRecordingAddOnResultResponse) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *ListRecordingAddOnResultResponse) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageSize

`func (o *ListRecordingAddOnResultResponse) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ListRecordingAddOnResultResponse) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ListRecordingAddOnResultResponse) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ListRecordingAddOnResultResponse) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetPreviousPageUri

`func (o *ListRecordingAddOnResultResponse) GetPreviousPageUri() string`

GetPreviousPageUri returns the PreviousPageUri field if non-nil, zero value otherwise.

### GetPreviousPageUriOk

`func (o *ListRecordingAddOnResultResponse) GetPreviousPageUriOk() (*string, bool)`

GetPreviousPageUriOk returns a tuple with the PreviousPageUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousPageUri

`func (o *ListRecordingAddOnResultResponse) SetPreviousPageUri(v string)`

SetPreviousPageUri sets PreviousPageUri field to given value.

### HasPreviousPageUri

`func (o *ListRecordingAddOnResultResponse) HasPreviousPageUri() bool`

HasPreviousPageUri returns a boolean if a field has been set.

### GetStart

`func (o *ListRecordingAddOnResultResponse) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ListRecordingAddOnResultResponse) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ListRecordingAddOnResultResponse) SetStart(v int32)`

SetStart sets Start field to given value.

### HasStart

`func (o *ListRecordingAddOnResultResponse) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetUri

`func (o *ListRecordingAddOnResultResponse) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ListRecordingAddOnResultResponse) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ListRecordingAddOnResultResponse) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *ListRecordingAddOnResultResponse) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


