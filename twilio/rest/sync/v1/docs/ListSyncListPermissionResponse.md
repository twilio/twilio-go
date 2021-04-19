# ListSyncListPermissionResponse

## Properties

Name | Type | Description
------------ | ------------- | -------------
**Meta** | Pointer to [**ListServiceResponseMeta**](ListServiceResponse_meta.md) |  | [optional] 
**Permissions** | Pointer to [**[]SyncV1ServiceSyncListSyncListPermission**](SyncV1ServiceSyncListSyncListPermission.md) |  | [optional] 

## Methods

### NewListSyncListPermissionResponse

`func NewListSyncListPermissionResponse() *ListSyncListPermissionResponse`

NewListSyncListPermissionResponse instantiates a new ListSyncListPermissionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSyncListPermissionResponseWithDefaults

`func NewListSyncListPermissionResponseWithDefaults() *ListSyncListPermissionResponse`

NewListSyncListPermissionResponseWithDefaults instantiates a new ListSyncListPermissionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListSyncListPermissionResponse) GetMeta() ListServiceResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListSyncListPermissionResponse) GetMetaOk() (*ListServiceResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListSyncListPermissionResponse) SetMeta(v ListServiceResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListSyncListPermissionResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetPermissions

`func (o *ListSyncListPermissionResponse) GetPermissions() []SyncV1ServiceSyncListSyncListPermission`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *ListSyncListPermissionResponse) GetPermissionsOk() (*[]SyncV1ServiceSyncListSyncListPermission, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *ListSyncListPermissionResponse) SetPermissions(v []SyncV1ServiceSyncListSyncListPermission)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *ListSyncListPermissionResponse) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


