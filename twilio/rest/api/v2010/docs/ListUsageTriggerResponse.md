# ListUsageTriggerResponse

## Properties

Name | Type | Description
------------ | ------------- | -------------
**End** | Pointer to **int32** |  | [optional] 
**FirstPageUri** | Pointer to **string** |  | [optional] 
**NextPageUri** | Pointer to **string** |  | [optional] 
**Page** | Pointer to **int32** |  | [optional] 
**PageSize** | Pointer to **int32** |  | [optional] 
**PreviousPageUri** | Pointer to **string** |  | [optional] 
**Start** | Pointer to **int32** |  | [optional] 
**Uri** | Pointer to **string** |  | [optional] 
**UsageTriggers** | Pointer to [**[]ApiV2010AccountUsageUsageTrigger**](ApiV2010AccountUsageUsageTrigger.md) |  | [optional] 

## Methods

### NewListUsageTriggerResponse

`func NewListUsageTriggerResponse() *ListUsageTriggerResponse`

NewListUsageTriggerResponse instantiates a new ListUsageTriggerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListUsageTriggerResponseWithDefaults

`func NewListUsageTriggerResponseWithDefaults() *ListUsageTriggerResponse`

NewListUsageTriggerResponseWithDefaults instantiates a new ListUsageTriggerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *ListUsageTriggerResponse) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *ListUsageTriggerResponse) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *ListUsageTriggerResponse) SetEnd(v int32)`

SetEnd sets End field to given value.

### HasEnd

`func (o *ListUsageTriggerResponse) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetFirstPageUri

`func (o *ListUsageTriggerResponse) GetFirstPageUri() string`

GetFirstPageUri returns the FirstPageUri field if non-nil, zero value otherwise.

### GetFirstPageUriOk

`func (o *ListUsageTriggerResponse) GetFirstPageUriOk() (*string, bool)`

GetFirstPageUriOk returns a tuple with the FirstPageUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstPageUri

`func (o *ListUsageTriggerResponse) SetFirstPageUri(v string)`

SetFirstPageUri sets FirstPageUri field to given value.

### HasFirstPageUri

`func (o *ListUsageTriggerResponse) HasFirstPageUri() bool`

HasFirstPageUri returns a boolean if a field has been set.

### GetNextPageUri

`func (o *ListUsageTriggerResponse) GetNextPageUri() string`

GetNextPageUri returns the NextPageUri field if non-nil, zero value otherwise.

### GetNextPageUriOk

`func (o *ListUsageTriggerResponse) GetNextPageUriOk() (*string, bool)`

GetNextPageUriOk returns a tuple with the NextPageUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageUri

`func (o *ListUsageTriggerResponse) SetNextPageUri(v string)`

SetNextPageUri sets NextPageUri field to given value.

### HasNextPageUri

`func (o *ListUsageTriggerResponse) HasNextPageUri() bool`

HasNextPageUri returns a boolean if a field has been set.

### GetPage

`func (o *ListUsageTriggerResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ListUsageTriggerResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ListUsageTriggerResponse) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *ListUsageTriggerResponse) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageSize

`func (o *ListUsageTriggerResponse) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ListUsageTriggerResponse) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ListUsageTriggerResponse) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ListUsageTriggerResponse) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetPreviousPageUri

`func (o *ListUsageTriggerResponse) GetPreviousPageUri() string`

GetPreviousPageUri returns the PreviousPageUri field if non-nil, zero value otherwise.

### GetPreviousPageUriOk

`func (o *ListUsageTriggerResponse) GetPreviousPageUriOk() (*string, bool)`

GetPreviousPageUriOk returns a tuple with the PreviousPageUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousPageUri

`func (o *ListUsageTriggerResponse) SetPreviousPageUri(v string)`

SetPreviousPageUri sets PreviousPageUri field to given value.

### HasPreviousPageUri

`func (o *ListUsageTriggerResponse) HasPreviousPageUri() bool`

HasPreviousPageUri returns a boolean if a field has been set.

### GetStart

`func (o *ListUsageTriggerResponse) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ListUsageTriggerResponse) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ListUsageTriggerResponse) SetStart(v int32)`

SetStart sets Start field to given value.

### HasStart

`func (o *ListUsageTriggerResponse) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetUri

`func (o *ListUsageTriggerResponse) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ListUsageTriggerResponse) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ListUsageTriggerResponse) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *ListUsageTriggerResponse) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetUsageTriggers

`func (o *ListUsageTriggerResponse) GetUsageTriggers() []ApiV2010AccountUsageUsageTrigger`

GetUsageTriggers returns the UsageTriggers field if non-nil, zero value otherwise.

### GetUsageTriggersOk

`func (o *ListUsageTriggerResponse) GetUsageTriggersOk() (*[]ApiV2010AccountUsageUsageTrigger, bool)`

GetUsageTriggersOk returns a tuple with the UsageTriggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageTriggers

`func (o *ListUsageTriggerResponse) SetUsageTriggers(v []ApiV2010AccountUsageUsageTrigger)`

SetUsageTriggers sets UsageTriggers field to given value.

### HasUsageTriggers

`func (o *ListUsageTriggerResponse) HasUsageTriggers() bool`

HasUsageTriggers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


