# ListUsageRecordDailyResponse

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
**UsageRecords** | Pointer to [**[]ApiV2010AccountUsageUsageRecordUsageRecordDaily**](ApiV2010AccountUsageUsageRecordUsageRecordDaily.md) |  | [optional] 

## Methods

### NewListUsageRecordDailyResponse

`func NewListUsageRecordDailyResponse() *ListUsageRecordDailyResponse`

NewListUsageRecordDailyResponse instantiates a new ListUsageRecordDailyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListUsageRecordDailyResponseWithDefaults

`func NewListUsageRecordDailyResponseWithDefaults() *ListUsageRecordDailyResponse`

NewListUsageRecordDailyResponseWithDefaults instantiates a new ListUsageRecordDailyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *ListUsageRecordDailyResponse) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *ListUsageRecordDailyResponse) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *ListUsageRecordDailyResponse) SetEnd(v int32)`

SetEnd sets End field to given value.

### HasEnd

`func (o *ListUsageRecordDailyResponse) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetFirstPageUri

`func (o *ListUsageRecordDailyResponse) GetFirstPageUri() string`

GetFirstPageUri returns the FirstPageUri field if non-nil, zero value otherwise.

### GetFirstPageUriOk

`func (o *ListUsageRecordDailyResponse) GetFirstPageUriOk() (*string, bool)`

GetFirstPageUriOk returns a tuple with the FirstPageUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstPageUri

`func (o *ListUsageRecordDailyResponse) SetFirstPageUri(v string)`

SetFirstPageUri sets FirstPageUri field to given value.

### HasFirstPageUri

`func (o *ListUsageRecordDailyResponse) HasFirstPageUri() bool`

HasFirstPageUri returns a boolean if a field has been set.

### GetNextPageUri

`func (o *ListUsageRecordDailyResponse) GetNextPageUri() string`

GetNextPageUri returns the NextPageUri field if non-nil, zero value otherwise.

### GetNextPageUriOk

`func (o *ListUsageRecordDailyResponse) GetNextPageUriOk() (*string, bool)`

GetNextPageUriOk returns a tuple with the NextPageUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageUri

`func (o *ListUsageRecordDailyResponse) SetNextPageUri(v string)`

SetNextPageUri sets NextPageUri field to given value.

### HasNextPageUri

`func (o *ListUsageRecordDailyResponse) HasNextPageUri() bool`

HasNextPageUri returns a boolean if a field has been set.

### GetPage

`func (o *ListUsageRecordDailyResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ListUsageRecordDailyResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ListUsageRecordDailyResponse) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *ListUsageRecordDailyResponse) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageSize

`func (o *ListUsageRecordDailyResponse) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ListUsageRecordDailyResponse) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ListUsageRecordDailyResponse) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ListUsageRecordDailyResponse) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetPreviousPageUri

`func (o *ListUsageRecordDailyResponse) GetPreviousPageUri() string`

GetPreviousPageUri returns the PreviousPageUri field if non-nil, zero value otherwise.

### GetPreviousPageUriOk

`func (o *ListUsageRecordDailyResponse) GetPreviousPageUriOk() (*string, bool)`

GetPreviousPageUriOk returns a tuple with the PreviousPageUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousPageUri

`func (o *ListUsageRecordDailyResponse) SetPreviousPageUri(v string)`

SetPreviousPageUri sets PreviousPageUri field to given value.

### HasPreviousPageUri

`func (o *ListUsageRecordDailyResponse) HasPreviousPageUri() bool`

HasPreviousPageUri returns a boolean if a field has been set.

### GetStart

`func (o *ListUsageRecordDailyResponse) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ListUsageRecordDailyResponse) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ListUsageRecordDailyResponse) SetStart(v int32)`

SetStart sets Start field to given value.

### HasStart

`func (o *ListUsageRecordDailyResponse) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetUri

`func (o *ListUsageRecordDailyResponse) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ListUsageRecordDailyResponse) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ListUsageRecordDailyResponse) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *ListUsageRecordDailyResponse) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetUsageRecords

`func (o *ListUsageRecordDailyResponse) GetUsageRecords() []ApiV2010AccountUsageUsageRecordUsageRecordDaily`

GetUsageRecords returns the UsageRecords field if non-nil, zero value otherwise.

### GetUsageRecordsOk

`func (o *ListUsageRecordDailyResponse) GetUsageRecordsOk() (*[]ApiV2010AccountUsageUsageRecordUsageRecordDaily, bool)`

GetUsageRecordsOk returns a tuple with the UsageRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageRecords

`func (o *ListUsageRecordDailyResponse) SetUsageRecords(v []ApiV2010AccountUsageUsageRecordUsageRecordDaily)`

SetUsageRecords sets UsageRecords field to given value.

### HasUsageRecords

`func (o *ListUsageRecordDailyResponse) HasUsageRecords() bool`

HasUsageRecords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


