# ListTranscriptionResponse

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
**Transcriptions** | Pointer to [**[]ApiV2010AccountTranscription**](ApiV2010AccountTranscription.md) |  | [optional] 
**Uri** | Pointer to **string** |  | [optional] 

## Methods

### NewListTranscriptionResponse

`func NewListTranscriptionResponse() *ListTranscriptionResponse`

NewListTranscriptionResponse instantiates a new ListTranscriptionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTranscriptionResponseWithDefaults

`func NewListTranscriptionResponseWithDefaults() *ListTranscriptionResponse`

NewListTranscriptionResponseWithDefaults instantiates a new ListTranscriptionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *ListTranscriptionResponse) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *ListTranscriptionResponse) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *ListTranscriptionResponse) SetEnd(v int32)`

SetEnd sets End field to given value.

### HasEnd

`func (o *ListTranscriptionResponse) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetFirstPageUri

`func (o *ListTranscriptionResponse) GetFirstPageUri() string`

GetFirstPageUri returns the FirstPageUri field if non-nil, zero value otherwise.

### GetFirstPageUriOk

`func (o *ListTranscriptionResponse) GetFirstPageUriOk() (*string, bool)`

GetFirstPageUriOk returns a tuple with the FirstPageUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstPageUri

`func (o *ListTranscriptionResponse) SetFirstPageUri(v string)`

SetFirstPageUri sets FirstPageUri field to given value.

### HasFirstPageUri

`func (o *ListTranscriptionResponse) HasFirstPageUri() bool`

HasFirstPageUri returns a boolean if a field has been set.

### GetNextPageUri

`func (o *ListTranscriptionResponse) GetNextPageUri() string`

GetNextPageUri returns the NextPageUri field if non-nil, zero value otherwise.

### GetNextPageUriOk

`func (o *ListTranscriptionResponse) GetNextPageUriOk() (*string, bool)`

GetNextPageUriOk returns a tuple with the NextPageUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageUri

`func (o *ListTranscriptionResponse) SetNextPageUri(v string)`

SetNextPageUri sets NextPageUri field to given value.

### HasNextPageUri

`func (o *ListTranscriptionResponse) HasNextPageUri() bool`

HasNextPageUri returns a boolean if a field has been set.

### GetPage

`func (o *ListTranscriptionResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ListTranscriptionResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ListTranscriptionResponse) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *ListTranscriptionResponse) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageSize

`func (o *ListTranscriptionResponse) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ListTranscriptionResponse) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ListTranscriptionResponse) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ListTranscriptionResponse) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetPreviousPageUri

`func (o *ListTranscriptionResponse) GetPreviousPageUri() string`

GetPreviousPageUri returns the PreviousPageUri field if non-nil, zero value otherwise.

### GetPreviousPageUriOk

`func (o *ListTranscriptionResponse) GetPreviousPageUriOk() (*string, bool)`

GetPreviousPageUriOk returns a tuple with the PreviousPageUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousPageUri

`func (o *ListTranscriptionResponse) SetPreviousPageUri(v string)`

SetPreviousPageUri sets PreviousPageUri field to given value.

### HasPreviousPageUri

`func (o *ListTranscriptionResponse) HasPreviousPageUri() bool`

HasPreviousPageUri returns a boolean if a field has been set.

### GetStart

`func (o *ListTranscriptionResponse) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ListTranscriptionResponse) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ListTranscriptionResponse) SetStart(v int32)`

SetStart sets Start field to given value.

### HasStart

`func (o *ListTranscriptionResponse) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetTranscriptions

`func (o *ListTranscriptionResponse) GetTranscriptions() []ApiV2010AccountTranscription`

GetTranscriptions returns the Transcriptions field if non-nil, zero value otherwise.

### GetTranscriptionsOk

`func (o *ListTranscriptionResponse) GetTranscriptionsOk() (*[]ApiV2010AccountTranscription, bool)`

GetTranscriptionsOk returns a tuple with the Transcriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscriptions

`func (o *ListTranscriptionResponse) SetTranscriptions(v []ApiV2010AccountTranscription)`

SetTranscriptions sets Transcriptions field to given value.

### HasTranscriptions

`func (o *ListTranscriptionResponse) HasTranscriptions() bool`

HasTranscriptions returns a boolean if a field has been set.

### GetUri

`func (o *ListTranscriptionResponse) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ListTranscriptionResponse) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ListTranscriptionResponse) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *ListTranscriptionResponse) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


