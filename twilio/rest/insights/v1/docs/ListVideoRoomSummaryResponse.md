# ListVideoRoomSummaryResponse

## Properties

Name | Type | Description
------------ | ------------- | -------------
**Meta** | Pointer to [**ListVideoRoomSummaryResponseMeta**](ListVideoRoomSummaryResponse_meta.md) |  | [optional] 
**Rooms** | Pointer to [**[]InsightsV1VideoRoomSummary**](InsightsV1VideoRoomSummary.md) |  | [optional] 

## Methods

### NewListVideoRoomSummaryResponse

`func NewListVideoRoomSummaryResponse() *ListVideoRoomSummaryResponse`

NewListVideoRoomSummaryResponse instantiates a new ListVideoRoomSummaryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListVideoRoomSummaryResponseWithDefaults

`func NewListVideoRoomSummaryResponseWithDefaults() *ListVideoRoomSummaryResponse`

NewListVideoRoomSummaryResponseWithDefaults instantiates a new ListVideoRoomSummaryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListVideoRoomSummaryResponse) GetMeta() ListVideoRoomSummaryResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListVideoRoomSummaryResponse) GetMetaOk() (*ListVideoRoomSummaryResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListVideoRoomSummaryResponse) SetMeta(v ListVideoRoomSummaryResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListVideoRoomSummaryResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetRooms

`func (o *ListVideoRoomSummaryResponse) GetRooms() []InsightsV1VideoRoomSummary`

GetRooms returns the Rooms field if non-nil, zero value otherwise.

### GetRoomsOk

`func (o *ListVideoRoomSummaryResponse) GetRoomsOk() (*[]InsightsV1VideoRoomSummary, bool)`

GetRoomsOk returns a tuple with the Rooms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRooms

`func (o *ListVideoRoomSummaryResponse) SetRooms(v []InsightsV1VideoRoomSummary)`

SetRooms sets Rooms field to given value.

### HasRooms

`func (o *ListVideoRoomSummaryResponse) HasRooms() bool`

HasRooms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


