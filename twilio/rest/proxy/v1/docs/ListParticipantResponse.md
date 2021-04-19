# ListParticipantResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**ListServiceResponseMeta**](ListServiceResponse_meta.md) |  | [optional] 
**Participants** | Pointer to [**[]ProxyV1ServiceSessionParticipant**](ProxyV1ServiceSessionParticipant.md) |  | [optional] 

## Methods

### NewListParticipantResponse

`func NewListParticipantResponse() *ListParticipantResponse`

NewListParticipantResponse instantiates a new ListParticipantResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListParticipantResponseWithDefaults

`func NewListParticipantResponseWithDefaults() *ListParticipantResponse`

NewListParticipantResponseWithDefaults instantiates a new ListParticipantResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListParticipantResponse) GetMeta() ListServiceResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListParticipantResponse) GetMetaOk() (*ListServiceResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListParticipantResponse) SetMeta(v ListServiceResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListParticipantResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetParticipants

`func (o *ListParticipantResponse) GetParticipants() []ProxyV1ServiceSessionParticipant`

GetParticipants returns the Participants field if non-nil, zero value otherwise.

### GetParticipantsOk

`func (o *ListParticipantResponse) GetParticipantsOk() (*[]ProxyV1ServiceSessionParticipant, bool)`

GetParticipantsOk returns a tuple with the Participants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipants

`func (o *ListParticipantResponse) SetParticipants(v []ProxyV1ServiceSessionParticipant)`

SetParticipants sets Participants field to given value.

### HasParticipants

`func (o *ListParticipantResponse) HasParticipants() bool`

HasParticipants returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


