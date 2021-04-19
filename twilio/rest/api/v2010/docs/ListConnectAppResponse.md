# ListConnectAppResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectApps** | Pointer to [**[]ApiV2010AccountConnectApp**](ApiV2010AccountConnectApp.md) |  | [optional] 
**End** | Pointer to **int32** |  | [optional] 
**FirstPageUri** | Pointer to **string** |  | [optional] 
**NextPageUri** | Pointer to **string** |  | [optional] 
**Page** | Pointer to **int32** |  | [optional] 
**PageSize** | Pointer to **int32** |  | [optional] 
**PreviousPageUri** | Pointer to **string** |  | [optional] 
**Start** | Pointer to **int32** |  | [optional] 
**Uri** | Pointer to **string** |  | [optional] 

## Methods

### NewListConnectAppResponse

`func NewListConnectAppResponse() *ListConnectAppResponse`

NewListConnectAppResponse instantiates a new ListConnectAppResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListConnectAppResponseWithDefaults

`func NewListConnectAppResponseWithDefaults() *ListConnectAppResponse`

NewListConnectAppResponseWithDefaults instantiates a new ListConnectAppResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectApps

`func (o *ListConnectAppResponse) GetConnectApps() []ApiV2010AccountConnectApp`

GetConnectApps returns the ConnectApps field if non-nil, zero value otherwise.

### GetConnectAppsOk

`func (o *ListConnectAppResponse) GetConnectAppsOk() (*[]ApiV2010AccountConnectApp, bool)`

GetConnectAppsOk returns a tuple with the ConnectApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectApps

`func (o *ListConnectAppResponse) SetConnectApps(v []ApiV2010AccountConnectApp)`

SetConnectApps sets ConnectApps field to given value.

### HasConnectApps

`func (o *ListConnectAppResponse) HasConnectApps() bool`

HasConnectApps returns a boolean if a field has been set.

### GetEnd

`func (o *ListConnectAppResponse) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *ListConnectAppResponse) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *ListConnectAppResponse) SetEnd(v int32)`

SetEnd sets End field to given value.

### HasEnd

`func (o *ListConnectAppResponse) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetFirstPageUri

`func (o *ListConnectAppResponse) GetFirstPageUri() string`

GetFirstPageUri returns the FirstPageUri field if non-nil, zero value otherwise.

### GetFirstPageUriOk

`func (o *ListConnectAppResponse) GetFirstPageUriOk() (*string, bool)`

GetFirstPageUriOk returns a tuple with the FirstPageUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstPageUri

`func (o *ListConnectAppResponse) SetFirstPageUri(v string)`

SetFirstPageUri sets FirstPageUri field to given value.

### HasFirstPageUri

`func (o *ListConnectAppResponse) HasFirstPageUri() bool`

HasFirstPageUri returns a boolean if a field has been set.

### GetNextPageUri

`func (o *ListConnectAppResponse) GetNextPageUri() string`

GetNextPageUri returns the NextPageUri field if non-nil, zero value otherwise.

### GetNextPageUriOk

`func (o *ListConnectAppResponse) GetNextPageUriOk() (*string, bool)`

GetNextPageUriOk returns a tuple with the NextPageUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageUri

`func (o *ListConnectAppResponse) SetNextPageUri(v string)`

SetNextPageUri sets NextPageUri field to given value.

### HasNextPageUri

`func (o *ListConnectAppResponse) HasNextPageUri() bool`

HasNextPageUri returns a boolean if a field has been set.

### GetPage

`func (o *ListConnectAppResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ListConnectAppResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ListConnectAppResponse) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *ListConnectAppResponse) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageSize

`func (o *ListConnectAppResponse) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ListConnectAppResponse) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ListConnectAppResponse) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ListConnectAppResponse) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetPreviousPageUri

`func (o *ListConnectAppResponse) GetPreviousPageUri() string`

GetPreviousPageUri returns the PreviousPageUri field if non-nil, zero value otherwise.

### GetPreviousPageUriOk

`func (o *ListConnectAppResponse) GetPreviousPageUriOk() (*string, bool)`

GetPreviousPageUriOk returns a tuple with the PreviousPageUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousPageUri

`func (o *ListConnectAppResponse) SetPreviousPageUri(v string)`

SetPreviousPageUri sets PreviousPageUri field to given value.

### HasPreviousPageUri

`func (o *ListConnectAppResponse) HasPreviousPageUri() bool`

HasPreviousPageUri returns a boolean if a field has been set.

### GetStart

`func (o *ListConnectAppResponse) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ListConnectAppResponse) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ListConnectAppResponse) SetStart(v int32)`

SetStart sets Start field to given value.

### HasStart

`func (o *ListConnectAppResponse) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetUri

`func (o *ListConnectAppResponse) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ListConnectAppResponse) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ListConnectAppResponse) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *ListConnectAppResponse) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


