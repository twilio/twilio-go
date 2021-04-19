# ListWorkerReservationResponse

## Properties

Name | Type | Description
------------ | ------------- | -------------
**Meta** | Pointer to [**ListWorkspaceResponseMeta**](ListWorkspaceResponse_meta.md) |  | [optional] 
**Reservations** | Pointer to [**[]TaskrouterV1WorkspaceWorkerWorkerReservation**](TaskrouterV1WorkspaceWorkerWorkerReservation.md) |  | [optional] 

## Methods

### NewListWorkerReservationResponse

`func NewListWorkerReservationResponse() *ListWorkerReservationResponse`

NewListWorkerReservationResponse instantiates a new ListWorkerReservationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListWorkerReservationResponseWithDefaults

`func NewListWorkerReservationResponseWithDefaults() *ListWorkerReservationResponse`

NewListWorkerReservationResponseWithDefaults instantiates a new ListWorkerReservationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListWorkerReservationResponse) GetMeta() ListWorkspaceResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListWorkerReservationResponse) GetMetaOk() (*ListWorkspaceResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListWorkerReservationResponse) SetMeta(v ListWorkspaceResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListWorkerReservationResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetReservations

`func (o *ListWorkerReservationResponse) GetReservations() []TaskrouterV1WorkspaceWorkerWorkerReservation`

GetReservations returns the Reservations field if non-nil, zero value otherwise.

### GetReservationsOk

`func (o *ListWorkerReservationResponse) GetReservationsOk() (*[]TaskrouterV1WorkspaceWorkerWorkerReservation, bool)`

GetReservationsOk returns a tuple with the Reservations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservations

`func (o *ListWorkerReservationResponse) SetReservations(v []TaskrouterV1WorkspaceWorkerWorkerReservation)`

SetReservations sets Reservations field to given value.

### HasReservations

`func (o *ListWorkerReservationResponse) HasReservations() bool`

HasReservations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


