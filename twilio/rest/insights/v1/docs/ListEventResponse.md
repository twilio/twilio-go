# ListEventResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | Pointer to [**[]InsightsV1CallEvent**](InsightsV1CallEvent.md) |  | [optional] 
**Meta** | Pointer to [**ListVideoRoomSummaryResponseMeta**](ListVideoRoomSummaryResponse_meta.md) |  | [optional] 

## Methods

### NewListEventResponse

`func NewListEventResponse() *ListEventResponse`

NewListEventResponse instantiates a new ListEventResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListEventResponseWithDefaults

`func NewListEventResponseWithDefaults() *ListEventResponse`

NewListEventResponseWithDefaults instantiates a new ListEventResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *ListEventResponse) GetEvents() []InsightsV1CallEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *ListEventResponse) GetEventsOk() (*[]InsightsV1CallEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *ListEventResponse) SetEvents(v []InsightsV1CallEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *ListEventResponse) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetMeta

`func (o *ListEventResponse) GetMeta() ListVideoRoomSummaryResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListEventResponse) GetMetaOk() (*ListVideoRoomSummaryResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListEventResponse) SetMeta(v ListVideoRoomSummaryResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListEventResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


