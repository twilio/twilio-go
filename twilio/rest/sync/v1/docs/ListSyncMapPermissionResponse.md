# ListSyncMapPermissionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**ListServiceResponseMeta**](ListServiceResponse_meta.md) |  | [optional] 
**Permissions** | Pointer to [**[]SyncV1ServiceSyncMapSyncMapPermission**](SyncV1ServiceSyncMapSyncMapPermission.md) |  | [optional] 

## Methods

### NewListSyncMapPermissionResponse

`func NewListSyncMapPermissionResponse() *ListSyncMapPermissionResponse`

NewListSyncMapPermissionResponse instantiates a new ListSyncMapPermissionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSyncMapPermissionResponseWithDefaults

`func NewListSyncMapPermissionResponseWithDefaults() *ListSyncMapPermissionResponse`

NewListSyncMapPermissionResponseWithDefaults instantiates a new ListSyncMapPermissionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListSyncMapPermissionResponse) GetMeta() ListServiceResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListSyncMapPermissionResponse) GetMetaOk() (*ListServiceResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListSyncMapPermissionResponse) SetMeta(v ListServiceResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListSyncMapPermissionResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetPermissions

`func (o *ListSyncMapPermissionResponse) GetPermissions() []SyncV1ServiceSyncMapSyncMapPermission`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *ListSyncMapPermissionResponse) GetPermissionsOk() (*[]SyncV1ServiceSyncMapSyncMapPermission, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *ListSyncMapPermissionResponse) SetPermissions(v []SyncV1ServiceSyncMapSyncMapPermission)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *ListSyncMapPermissionResponse) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


