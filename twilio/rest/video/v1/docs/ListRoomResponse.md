# ListRoomResponse

## Properties

Name | Type | Description
------------ | ------------- | -------------
**Meta** | Pointer to [**ListCompositionHookResponseMeta**](ListCompositionHookResponse_meta.md) |  | [optional] 
**Rooms** | Pointer to [**[]VideoV1Room**](VideoV1Room.md) |  | [optional] 

## Methods

### NewListRoomResponse

`func NewListRoomResponse() *ListRoomResponse`

NewListRoomResponse instantiates a new ListRoomResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListRoomResponseWithDefaults

`func NewListRoomResponseWithDefaults() *ListRoomResponse`

NewListRoomResponseWithDefaults instantiates a new ListRoomResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListRoomResponse) GetMeta() ListCompositionHookResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListRoomResponse) GetMetaOk() (*ListCompositionHookResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListRoomResponse) SetMeta(v ListCompositionHookResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListRoomResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetRooms

`func (o *ListRoomResponse) GetRooms() []VideoV1Room`

GetRooms returns the Rooms field if non-nil, zero value otherwise.

### GetRoomsOk

`func (o *ListRoomResponse) GetRoomsOk() (*[]VideoV1Room, bool)`

GetRoomsOk returns a tuple with the Rooms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRooms

`func (o *ListRoomResponse) SetRooms(v []VideoV1Room)`

SetRooms sets Rooms field to given value.

### HasRooms

`func (o *ListRoomResponse) HasRooms() bool`

HasRooms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


