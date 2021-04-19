# ListSigningKeyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | Pointer to **int32** |  | [optional] 
**FirstPageUri** | Pointer to **string** |  | [optional] 
**NextPageUri** | Pointer to **string** |  | [optional] 
**Page** | Pointer to **int32** |  | [optional] 
**PageSize** | Pointer to **int32** |  | [optional] 
**PreviousPageUri** | Pointer to **string** |  | [optional] 
**SigningKeys** | Pointer to [**[]ApiV2010AccountSigningKey**](ApiV2010AccountSigningKey.md) |  | [optional] 
**Start** | Pointer to **int32** |  | [optional] 
**Uri** | Pointer to **string** |  | [optional] 

## Methods

### NewListSigningKeyResponse

`func NewListSigningKeyResponse() *ListSigningKeyResponse`

NewListSigningKeyResponse instantiates a new ListSigningKeyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSigningKeyResponseWithDefaults

`func NewListSigningKeyResponseWithDefaults() *ListSigningKeyResponse`

NewListSigningKeyResponseWithDefaults instantiates a new ListSigningKeyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *ListSigningKeyResponse) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *ListSigningKeyResponse) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *ListSigningKeyResponse) SetEnd(v int32)`

SetEnd sets End field to given value.

### HasEnd

`func (o *ListSigningKeyResponse) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetFirstPageUri

`func (o *ListSigningKeyResponse) GetFirstPageUri() string`

GetFirstPageUri returns the FirstPageUri field if non-nil, zero value otherwise.

### GetFirstPageUriOk

`func (o *ListSigningKeyResponse) GetFirstPageUriOk() (*string, bool)`

GetFirstPageUriOk returns a tuple with the FirstPageUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstPageUri

`func (o *ListSigningKeyResponse) SetFirstPageUri(v string)`

SetFirstPageUri sets FirstPageUri field to given value.

### HasFirstPageUri

`func (o *ListSigningKeyResponse) HasFirstPageUri() bool`

HasFirstPageUri returns a boolean if a field has been set.

### GetNextPageUri

`func (o *ListSigningKeyResponse) GetNextPageUri() string`

GetNextPageUri returns the NextPageUri field if non-nil, zero value otherwise.

### GetNextPageUriOk

`func (o *ListSigningKeyResponse) GetNextPageUriOk() (*string, bool)`

GetNextPageUriOk returns a tuple with the NextPageUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageUri

`func (o *ListSigningKeyResponse) SetNextPageUri(v string)`

SetNextPageUri sets NextPageUri field to given value.

### HasNextPageUri

`func (o *ListSigningKeyResponse) HasNextPageUri() bool`

HasNextPageUri returns a boolean if a field has been set.

### GetPage

`func (o *ListSigningKeyResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ListSigningKeyResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ListSigningKeyResponse) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *ListSigningKeyResponse) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageSize

`func (o *ListSigningKeyResponse) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ListSigningKeyResponse) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ListSigningKeyResponse) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ListSigningKeyResponse) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetPreviousPageUri

`func (o *ListSigningKeyResponse) GetPreviousPageUri() string`

GetPreviousPageUri returns the PreviousPageUri field if non-nil, zero value otherwise.

### GetPreviousPageUriOk

`func (o *ListSigningKeyResponse) GetPreviousPageUriOk() (*string, bool)`

GetPreviousPageUriOk returns a tuple with the PreviousPageUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousPageUri

`func (o *ListSigningKeyResponse) SetPreviousPageUri(v string)`

SetPreviousPageUri sets PreviousPageUri field to given value.

### HasPreviousPageUri

`func (o *ListSigningKeyResponse) HasPreviousPageUri() bool`

HasPreviousPageUri returns a boolean if a field has been set.

### GetSigningKeys

`func (o *ListSigningKeyResponse) GetSigningKeys() []ApiV2010AccountSigningKey`

GetSigningKeys returns the SigningKeys field if non-nil, zero value otherwise.

### GetSigningKeysOk

`func (o *ListSigningKeyResponse) GetSigningKeysOk() (*[]ApiV2010AccountSigningKey, bool)`

GetSigningKeysOk returns a tuple with the SigningKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningKeys

`func (o *ListSigningKeyResponse) SetSigningKeys(v []ApiV2010AccountSigningKey)`

SetSigningKeys sets SigningKeys field to given value.

### HasSigningKeys

`func (o *ListSigningKeyResponse) HasSigningKeys() bool`

HasSigningKeys returns a boolean if a field has been set.

### GetStart

`func (o *ListSigningKeyResponse) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ListSigningKeyResponse) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ListSigningKeyResponse) SetStart(v int32)`

SetStart sets Start field to given value.

### HasStart

`func (o *ListSigningKeyResponse) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetUri

`func (o *ListSigningKeyResponse) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ListSigningKeyResponse) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ListSigningKeyResponse) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *ListSigningKeyResponse) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


