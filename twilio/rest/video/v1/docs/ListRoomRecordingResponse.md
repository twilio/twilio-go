# ListRoomRecordingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**ListCompositionHookResponseMeta**](ListCompositionHookResponse_meta.md) |  | [optional] 
**Recordings** | Pointer to [**[]VideoV1RoomRoomRecording**](VideoV1RoomRoomRecording.md) |  | [optional] 

## Methods

### NewListRoomRecordingResponse

`func NewListRoomRecordingResponse() *ListRoomRecordingResponse`

NewListRoomRecordingResponse instantiates a new ListRoomRecordingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListRoomRecordingResponseWithDefaults

`func NewListRoomRecordingResponseWithDefaults() *ListRoomRecordingResponse`

NewListRoomRecordingResponseWithDefaults instantiates a new ListRoomRecordingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListRoomRecordingResponse) GetMeta() ListCompositionHookResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListRoomRecordingResponse) GetMetaOk() (*ListCompositionHookResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListRoomRecordingResponse) SetMeta(v ListCompositionHookResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListRoomRecordingResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetRecordings

`func (o *ListRoomRecordingResponse) GetRecordings() []VideoV1RoomRoomRecording`

GetRecordings returns the Recordings field if non-nil, zero value otherwise.

### GetRecordingsOk

`func (o *ListRoomRecordingResponse) GetRecordingsOk() (*[]VideoV1RoomRoomRecording, bool)`

GetRecordingsOk returns a tuple with the Recordings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordings

`func (o *ListRoomRecordingResponse) SetRecordings(v []VideoV1RoomRoomRecording)`

SetRecordings sets Recordings field to given value.

### HasRecordings

`func (o *ListRoomRecordingResponse) HasRecordings() bool`

HasRecordings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


