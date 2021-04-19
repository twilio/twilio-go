# ListTaskReservationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**ListWorkspaceResponseMeta**](ListWorkspaceResponse_meta.md) |  | [optional] 
**Reservations** | Pointer to [**[]TaskrouterV1WorkspaceTaskTaskReservation**](TaskrouterV1WorkspaceTaskTaskReservation.md) |  | [optional] 

## Methods

### NewListTaskReservationResponse

`func NewListTaskReservationResponse() *ListTaskReservationResponse`

NewListTaskReservationResponse instantiates a new ListTaskReservationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTaskReservationResponseWithDefaults

`func NewListTaskReservationResponseWithDefaults() *ListTaskReservationResponse`

NewListTaskReservationResponseWithDefaults instantiates a new ListTaskReservationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListTaskReservationResponse) GetMeta() ListWorkspaceResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListTaskReservationResponse) GetMetaOk() (*ListWorkspaceResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListTaskReservationResponse) SetMeta(v ListWorkspaceResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListTaskReservationResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetReservations

`func (o *ListTaskReservationResponse) GetReservations() []TaskrouterV1WorkspaceTaskTaskReservation`

GetReservations returns the Reservations field if non-nil, zero value otherwise.

### GetReservationsOk

`func (o *ListTaskReservationResponse) GetReservationsOk() (*[]TaskrouterV1WorkspaceTaskTaskReservation, bool)`

GetReservationsOk returns a tuple with the Reservations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservations

`func (o *ListTaskReservationResponse) SetReservations(v []TaskrouterV1WorkspaceTaskTaskReservation)`

SetReservations sets Reservations field to given value.

### HasReservations

`func (o *ListTaskReservationResponse) HasReservations() bool`

HasReservations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


