# ListUsageRecordTodayResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | Pointer to **int32** |  | [optional] 
**FirstPageUri** | Pointer to **string** |  | [optional] 
**NextPageUri** | Pointer to **string** |  | [optional] 
**Page** | Pointer to **int32** |  | [optional] 
**PageSize** | Pointer to **int32** |  | [optional] 
**PreviousPageUri** | Pointer to **string** |  | [optional] 
**Start** | Pointer to **int32** |  | [optional] 
**Uri** | Pointer to **string** |  | [optional] 
**UsageRecords** | Pointer to [**[]ApiV2010AccountUsageUsageRecordUsageRecordToday**](ApiV2010AccountUsageUsageRecordUsageRecordToday.md) |  | [optional] 

## Methods

### NewListUsageRecordTodayResponse

`func NewListUsageRecordTodayResponse() *ListUsageRecordTodayResponse`

NewListUsageRecordTodayResponse instantiates a new ListUsageRecordTodayResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListUsageRecordTodayResponseWithDefaults

`func NewListUsageRecordTodayResponseWithDefaults() *ListUsageRecordTodayResponse`

NewListUsageRecordTodayResponseWithDefaults instantiates a new ListUsageRecordTodayResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *ListUsageRecordTodayResponse) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *ListUsageRecordTodayResponse) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *ListUsageRecordTodayResponse) SetEnd(v int32)`

SetEnd sets End field to given value.

### HasEnd

`func (o *ListUsageRecordTodayResponse) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetFirstPageUri

`func (o *ListUsageRecordTodayResponse) GetFirstPageUri() string`

GetFirstPageUri returns the FirstPageUri field if non-nil, zero value otherwise.

### GetFirstPageUriOk

`func (o *ListUsageRecordTodayResponse) GetFirstPageUriOk() (*string, bool)`

GetFirstPageUriOk returns a tuple with the FirstPageUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstPageUri

`func (o *ListUsageRecordTodayResponse) SetFirstPageUri(v string)`

SetFirstPageUri sets FirstPageUri field to given value.

### HasFirstPageUri

`func (o *ListUsageRecordTodayResponse) HasFirstPageUri() bool`

HasFirstPageUri returns a boolean if a field has been set.

### GetNextPageUri

`func (o *ListUsageRecordTodayResponse) GetNextPageUri() string`

GetNextPageUri returns the NextPageUri field if non-nil, zero value otherwise.

### GetNextPageUriOk

`func (o *ListUsageRecordTodayResponse) GetNextPageUriOk() (*string, bool)`

GetNextPageUriOk returns a tuple with the NextPageUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageUri

`func (o *ListUsageRecordTodayResponse) SetNextPageUri(v string)`

SetNextPageUri sets NextPageUri field to given value.

### HasNextPageUri

`func (o *ListUsageRecordTodayResponse) HasNextPageUri() bool`

HasNextPageUri returns a boolean if a field has been set.

### GetPage

`func (o *ListUsageRecordTodayResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ListUsageRecordTodayResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ListUsageRecordTodayResponse) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *ListUsageRecordTodayResponse) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageSize

`func (o *ListUsageRecordTodayResponse) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ListUsageRecordTodayResponse) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ListUsageRecordTodayResponse) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ListUsageRecordTodayResponse) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetPreviousPageUri

`func (o *ListUsageRecordTodayResponse) GetPreviousPageUri() string`

GetPreviousPageUri returns the PreviousPageUri field if non-nil, zero value otherwise.

### GetPreviousPageUriOk

`func (o *ListUsageRecordTodayResponse) GetPreviousPageUriOk() (*string, bool)`

GetPreviousPageUriOk returns a tuple with the PreviousPageUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousPageUri

`func (o *ListUsageRecordTodayResponse) SetPreviousPageUri(v string)`

SetPreviousPageUri sets PreviousPageUri field to given value.

### HasPreviousPageUri

`func (o *ListUsageRecordTodayResponse) HasPreviousPageUri() bool`

HasPreviousPageUri returns a boolean if a field has been set.

### GetStart

`func (o *ListUsageRecordTodayResponse) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ListUsageRecordTodayResponse) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ListUsageRecordTodayResponse) SetStart(v int32)`

SetStart sets Start field to given value.

### HasStart

`func (o *ListUsageRecordTodayResponse) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetUri

`func (o *ListUsageRecordTodayResponse) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ListUsageRecordTodayResponse) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ListUsageRecordTodayResponse) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *ListUsageRecordTodayResponse) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetUsageRecords

`func (o *ListUsageRecordTodayResponse) GetUsageRecords() []ApiV2010AccountUsageUsageRecordUsageRecordToday`

GetUsageRecords returns the UsageRecords field if non-nil, zero value otherwise.

### GetUsageRecordsOk

`func (o *ListUsageRecordTodayResponse) GetUsageRecordsOk() (*[]ApiV2010AccountUsageUsageRecordUsageRecordToday, bool)`

GetUsageRecordsOk returns a tuple with the UsageRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageRecords

`func (o *ListUsageRecordTodayResponse) SetUsageRecords(v []ApiV2010AccountUsageUsageRecordUsageRecordToday)`

SetUsageRecords sets UsageRecords field to given value.

### HasUsageRecords

`func (o *ListUsageRecordTodayResponse) HasUsageRecords() bool`

HasUsageRecords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


