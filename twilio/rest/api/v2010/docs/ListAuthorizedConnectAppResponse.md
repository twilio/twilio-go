# ListAuthorizedConnectAppResponse

## Properties

Name | Type | Description
------------ | ------------- | -------------
**AuthorizedConnectApps** | Pointer to [**[]ApiV2010AccountAuthorizedConnectApp**](ApiV2010AccountAuthorizedConnectApp.md) |  | [optional] 
**End** | Pointer to **int32** |  | [optional] 
**FirstPageUri** | Pointer to **string** |  | [optional] 
**NextPageUri** | Pointer to **string** |  | [optional] 
**Page** | Pointer to **int32** |  | [optional] 
**PageSize** | Pointer to **int32** |  | [optional] 
**PreviousPageUri** | Pointer to **string** |  | [optional] 
**Start** | Pointer to **int32** |  | [optional] 
**Uri** | Pointer to **string** |  | [optional] 

## Methods

### NewListAuthorizedConnectAppResponse

`func NewListAuthorizedConnectAppResponse() *ListAuthorizedConnectAppResponse`

NewListAuthorizedConnectAppResponse instantiates a new ListAuthorizedConnectAppResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAuthorizedConnectAppResponseWithDefaults

`func NewListAuthorizedConnectAppResponseWithDefaults() *ListAuthorizedConnectAppResponse`

NewListAuthorizedConnectAppResponseWithDefaults instantiates a new ListAuthorizedConnectAppResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorizedConnectApps

`func (o *ListAuthorizedConnectAppResponse) GetAuthorizedConnectApps() []ApiV2010AccountAuthorizedConnectApp`

GetAuthorizedConnectApps returns the AuthorizedConnectApps field if non-nil, zero value otherwise.

### GetAuthorizedConnectAppsOk

`func (o *ListAuthorizedConnectAppResponse) GetAuthorizedConnectAppsOk() (*[]ApiV2010AccountAuthorizedConnectApp, bool)`

GetAuthorizedConnectAppsOk returns a tuple with the AuthorizedConnectApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedConnectApps

`func (o *ListAuthorizedConnectAppResponse) SetAuthorizedConnectApps(v []ApiV2010AccountAuthorizedConnectApp)`

SetAuthorizedConnectApps sets AuthorizedConnectApps field to given value.

### HasAuthorizedConnectApps

`func (o *ListAuthorizedConnectAppResponse) HasAuthorizedConnectApps() bool`

HasAuthorizedConnectApps returns a boolean if a field has been set.

### GetEnd

`func (o *ListAuthorizedConnectAppResponse) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *ListAuthorizedConnectAppResponse) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *ListAuthorizedConnectAppResponse) SetEnd(v int32)`

SetEnd sets End field to given value.

### HasEnd

`func (o *ListAuthorizedConnectAppResponse) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetFirstPageUri

`func (o *ListAuthorizedConnectAppResponse) GetFirstPageUri() string`

GetFirstPageUri returns the FirstPageUri field if non-nil, zero value otherwise.

### GetFirstPageUriOk

`func (o *ListAuthorizedConnectAppResponse) GetFirstPageUriOk() (*string, bool)`

GetFirstPageUriOk returns a tuple with the FirstPageUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstPageUri

`func (o *ListAuthorizedConnectAppResponse) SetFirstPageUri(v string)`

SetFirstPageUri sets FirstPageUri field to given value.

### HasFirstPageUri

`func (o *ListAuthorizedConnectAppResponse) HasFirstPageUri() bool`

HasFirstPageUri returns a boolean if a field has been set.

### GetNextPageUri

`func (o *ListAuthorizedConnectAppResponse) GetNextPageUri() string`

GetNextPageUri returns the NextPageUri field if non-nil, zero value otherwise.

### GetNextPageUriOk

`func (o *ListAuthorizedConnectAppResponse) GetNextPageUriOk() (*string, bool)`

GetNextPageUriOk returns a tuple with the NextPageUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageUri

`func (o *ListAuthorizedConnectAppResponse) SetNextPageUri(v string)`

SetNextPageUri sets NextPageUri field to given value.

### HasNextPageUri

`func (o *ListAuthorizedConnectAppResponse) HasNextPageUri() bool`

HasNextPageUri returns a boolean if a field has been set.

### GetPage

`func (o *ListAuthorizedConnectAppResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ListAuthorizedConnectAppResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ListAuthorizedConnectAppResponse) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *ListAuthorizedConnectAppResponse) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageSize

`func (o *ListAuthorizedConnectAppResponse) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ListAuthorizedConnectAppResponse) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ListAuthorizedConnectAppResponse) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ListAuthorizedConnectAppResponse) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetPreviousPageUri

`func (o *ListAuthorizedConnectAppResponse) GetPreviousPageUri() string`

GetPreviousPageUri returns the PreviousPageUri field if non-nil, zero value otherwise.

### GetPreviousPageUriOk

`func (o *ListAuthorizedConnectAppResponse) GetPreviousPageUriOk() (*string, bool)`

GetPreviousPageUriOk returns a tuple with the PreviousPageUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousPageUri

`func (o *ListAuthorizedConnectAppResponse) SetPreviousPageUri(v string)`

SetPreviousPageUri sets PreviousPageUri field to given value.

### HasPreviousPageUri

`func (o *ListAuthorizedConnectAppResponse) HasPreviousPageUri() bool`

HasPreviousPageUri returns a boolean if a field has been set.

### GetStart

`func (o *ListAuthorizedConnectAppResponse) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ListAuthorizedConnectAppResponse) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ListAuthorizedConnectAppResponse) SetStart(v int32)`

SetStart sets Start field to given value.

### HasStart

`func (o *ListAuthorizedConnectAppResponse) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetUri

`func (o *ListAuthorizedConnectAppResponse) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ListAuthorizedConnectAppResponse) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ListAuthorizedConnectAppResponse) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *ListAuthorizedConnectAppResponse) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


