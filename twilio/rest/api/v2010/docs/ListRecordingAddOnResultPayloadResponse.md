# ListRecordingAddOnResultPayloadResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | Pointer to **int32** |  | [optional] 
**FirstPageUri** | Pointer to **string** |  | [optional] 
**NextPageUri** | Pointer to **string** |  | [optional] 
**Page** | Pointer to **int32** |  | [optional] 
**PageSize** | Pointer to **int32** |  | [optional] 
**Payloads** | Pointer to [**[]ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload**](ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload.md) |  | [optional] 
**PreviousPageUri** | Pointer to **string** |  | [optional] 
**Start** | Pointer to **int32** |  | [optional] 
**Uri** | Pointer to **string** |  | [optional] 

## Methods

### NewListRecordingAddOnResultPayloadResponse

`func NewListRecordingAddOnResultPayloadResponse() *ListRecordingAddOnResultPayloadResponse`

NewListRecordingAddOnResultPayloadResponse instantiates a new ListRecordingAddOnResultPayloadResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListRecordingAddOnResultPayloadResponseWithDefaults

`func NewListRecordingAddOnResultPayloadResponseWithDefaults() *ListRecordingAddOnResultPayloadResponse`

NewListRecordingAddOnResultPayloadResponseWithDefaults instantiates a new ListRecordingAddOnResultPayloadResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *ListRecordingAddOnResultPayloadResponse) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *ListRecordingAddOnResultPayloadResponse) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *ListRecordingAddOnResultPayloadResponse) SetEnd(v int32)`

SetEnd sets End field to given value.

### HasEnd

`func (o *ListRecordingAddOnResultPayloadResponse) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetFirstPageUri

`func (o *ListRecordingAddOnResultPayloadResponse) GetFirstPageUri() string`

GetFirstPageUri returns the FirstPageUri field if non-nil, zero value otherwise.

### GetFirstPageUriOk

`func (o *ListRecordingAddOnResultPayloadResponse) GetFirstPageUriOk() (*string, bool)`

GetFirstPageUriOk returns a tuple with the FirstPageUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstPageUri

`func (o *ListRecordingAddOnResultPayloadResponse) SetFirstPageUri(v string)`

SetFirstPageUri sets FirstPageUri field to given value.

### HasFirstPageUri

`func (o *ListRecordingAddOnResultPayloadResponse) HasFirstPageUri() bool`

HasFirstPageUri returns a boolean if a field has been set.

### GetNextPageUri

`func (o *ListRecordingAddOnResultPayloadResponse) GetNextPageUri() string`

GetNextPageUri returns the NextPageUri field if non-nil, zero value otherwise.

### GetNextPageUriOk

`func (o *ListRecordingAddOnResultPayloadResponse) GetNextPageUriOk() (*string, bool)`

GetNextPageUriOk returns a tuple with the NextPageUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageUri

`func (o *ListRecordingAddOnResultPayloadResponse) SetNextPageUri(v string)`

SetNextPageUri sets NextPageUri field to given value.

### HasNextPageUri

`func (o *ListRecordingAddOnResultPayloadResponse) HasNextPageUri() bool`

HasNextPageUri returns a boolean if a field has been set.

### GetPage

`func (o *ListRecordingAddOnResultPayloadResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ListRecordingAddOnResultPayloadResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ListRecordingAddOnResultPayloadResponse) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *ListRecordingAddOnResultPayloadResponse) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageSize

`func (o *ListRecordingAddOnResultPayloadResponse) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ListRecordingAddOnResultPayloadResponse) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ListRecordingAddOnResultPayloadResponse) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ListRecordingAddOnResultPayloadResponse) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetPayloads

`func (o *ListRecordingAddOnResultPayloadResponse) GetPayloads() []ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload`

GetPayloads returns the Payloads field if non-nil, zero value otherwise.

### GetPayloadsOk

`func (o *ListRecordingAddOnResultPayloadResponse) GetPayloadsOk() (*[]ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload, bool)`

GetPayloadsOk returns a tuple with the Payloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloads

`func (o *ListRecordingAddOnResultPayloadResponse) SetPayloads(v []ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload)`

SetPayloads sets Payloads field to given value.

### HasPayloads

`func (o *ListRecordingAddOnResultPayloadResponse) HasPayloads() bool`

HasPayloads returns a boolean if a field has been set.

### GetPreviousPageUri

`func (o *ListRecordingAddOnResultPayloadResponse) GetPreviousPageUri() string`

GetPreviousPageUri returns the PreviousPageUri field if non-nil, zero value otherwise.

### GetPreviousPageUriOk

`func (o *ListRecordingAddOnResultPayloadResponse) GetPreviousPageUriOk() (*string, bool)`

GetPreviousPageUriOk returns a tuple with the PreviousPageUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousPageUri

`func (o *ListRecordingAddOnResultPayloadResponse) SetPreviousPageUri(v string)`

SetPreviousPageUri sets PreviousPageUri field to given value.

### HasPreviousPageUri

`func (o *ListRecordingAddOnResultPayloadResponse) HasPreviousPageUri() bool`

HasPreviousPageUri returns a boolean if a field has been set.

### GetStart

`func (o *ListRecordingAddOnResultPayloadResponse) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ListRecordingAddOnResultPayloadResponse) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ListRecordingAddOnResultPayloadResponse) SetStart(v int32)`

SetStart sets Start field to given value.

### HasStart

`func (o *ListRecordingAddOnResultPayloadResponse) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetUri

`func (o *ListRecordingAddOnResultPayloadResponse) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ListRecordingAddOnResultPayloadResponse) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ListRecordingAddOnResultPayloadResponse) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *ListRecordingAddOnResultPayloadResponse) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


