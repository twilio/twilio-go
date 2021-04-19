# ListVideoParticipantSummaryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**ListVideoRoomSummaryResponseMeta**](ListVideoRoomSummaryResponse_meta.md) |  | [optional] 
**Participants** | Pointer to [**[]InsightsV1VideoRoomSummaryVideoParticipantSummary**](InsightsV1VideoRoomSummaryVideoParticipantSummary.md) |  | [optional] 

## Methods

### NewListVideoParticipantSummaryResponse

`func NewListVideoParticipantSummaryResponse() *ListVideoParticipantSummaryResponse`

NewListVideoParticipantSummaryResponse instantiates a new ListVideoParticipantSummaryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListVideoParticipantSummaryResponseWithDefaults

`func NewListVideoParticipantSummaryResponseWithDefaults() *ListVideoParticipantSummaryResponse`

NewListVideoParticipantSummaryResponseWithDefaults instantiates a new ListVideoParticipantSummaryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListVideoParticipantSummaryResponse) GetMeta() ListVideoRoomSummaryResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListVideoParticipantSummaryResponse) GetMetaOk() (*ListVideoRoomSummaryResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListVideoParticipantSummaryResponse) SetMeta(v ListVideoRoomSummaryResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListVideoParticipantSummaryResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetParticipants

`func (o *ListVideoParticipantSummaryResponse) GetParticipants() []InsightsV1VideoRoomSummaryVideoParticipantSummary`

GetParticipants returns the Participants field if non-nil, zero value otherwise.

### GetParticipantsOk

`func (o *ListVideoParticipantSummaryResponse) GetParticipantsOk() (*[]InsightsV1VideoRoomSummaryVideoParticipantSummary, bool)`

GetParticipantsOk returns a tuple with the Participants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipants

`func (o *ListVideoParticipantSummaryResponse) SetParticipants(v []InsightsV1VideoRoomSummaryVideoParticipantSummary)`

SetParticipants sets Participants field to given value.

### HasParticipants

`func (o *ListVideoParticipantSummaryResponse) HasParticipants() bool`

HasParticipants returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


