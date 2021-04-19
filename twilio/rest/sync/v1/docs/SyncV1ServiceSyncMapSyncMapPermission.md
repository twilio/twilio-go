# SyncV1ServiceSyncMapSyncMapPermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**Identity** | Pointer to **NullableString** | The identity of the user to whom the Sync Document Permission applies | [optional] 
**Manage** | Pointer to **NullableBool** | Manage access | [optional] 
**MapSid** | Pointer to **NullableString** | Sync Map SID | [optional] 
**Read** | Pointer to **NullableBool** | Read access | [optional] 
**ServiceSid** | Pointer to **NullableString** | The SID of the Sync Service that the resource is associated with | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the Sync Map Permission resource | [optional] 
**Write** | Pointer to **NullableBool** | Write access | [optional] 

## Methods

### NewSyncV1ServiceSyncMapSyncMapPermission

`func NewSyncV1ServiceSyncMapSyncMapPermission() *SyncV1ServiceSyncMapSyncMapPermission`

NewSyncV1ServiceSyncMapSyncMapPermission instantiates a new SyncV1ServiceSyncMapSyncMapPermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncV1ServiceSyncMapSyncMapPermissionWithDefaults

`func NewSyncV1ServiceSyncMapSyncMapPermissionWithDefaults() *SyncV1ServiceSyncMapSyncMapPermission`

NewSyncV1ServiceSyncMapSyncMapPermissionWithDefaults instantiates a new SyncV1ServiceSyncMapSyncMapPermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *SyncV1ServiceSyncMapSyncMapPermission) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *SyncV1ServiceSyncMapSyncMapPermission) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *SyncV1ServiceSyncMapSyncMapPermission) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *SyncV1ServiceSyncMapSyncMapPermission) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *SyncV1ServiceSyncMapSyncMapPermission) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *SyncV1ServiceSyncMapSyncMapPermission) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetIdentity

`func (o *SyncV1ServiceSyncMapSyncMapPermission) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *SyncV1ServiceSyncMapSyncMapPermission) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *SyncV1ServiceSyncMapSyncMapPermission) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *SyncV1ServiceSyncMapSyncMapPermission) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### SetIdentityNil

`func (o *SyncV1ServiceSyncMapSyncMapPermission) SetIdentityNil(b bool)`

 SetIdentityNil sets the value for Identity to be an explicit nil

### UnsetIdentity
`func (o *SyncV1ServiceSyncMapSyncMapPermission) UnsetIdentity()`

UnsetIdentity ensures that no value is present for Identity, not even an explicit nil
### GetManage

`func (o *SyncV1ServiceSyncMapSyncMapPermission) GetManage() bool`

GetManage returns the Manage field if non-nil, zero value otherwise.

### GetManageOk

`func (o *SyncV1ServiceSyncMapSyncMapPermission) GetManageOk() (*bool, bool)`

GetManageOk returns a tuple with the Manage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManage

`func (o *SyncV1ServiceSyncMapSyncMapPermission) SetManage(v bool)`

SetManage sets Manage field to given value.

### HasManage

`func (o *SyncV1ServiceSyncMapSyncMapPermission) HasManage() bool`

HasManage returns a boolean if a field has been set.

### SetManageNil

`func (o *SyncV1ServiceSyncMapSyncMapPermission) SetManageNil(b bool)`

 SetManageNil sets the value for Manage to be an explicit nil

### UnsetManage
`func (o *SyncV1ServiceSyncMapSyncMapPermission) UnsetManage()`

UnsetManage ensures that no value is present for Manage, not even an explicit nil
### GetMapSid

`func (o *SyncV1ServiceSyncMapSyncMapPermission) GetMapSid() string`

GetMapSid returns the MapSid field if non-nil, zero value otherwise.

### GetMapSidOk

`func (o *SyncV1ServiceSyncMapSyncMapPermission) GetMapSidOk() (*string, bool)`

GetMapSidOk returns a tuple with the MapSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapSid

`func (o *SyncV1ServiceSyncMapSyncMapPermission) SetMapSid(v string)`

SetMapSid sets MapSid field to given value.

### HasMapSid

`func (o *SyncV1ServiceSyncMapSyncMapPermission) HasMapSid() bool`

HasMapSid returns a boolean if a field has been set.

### SetMapSidNil

`func (o *SyncV1ServiceSyncMapSyncMapPermission) SetMapSidNil(b bool)`

 SetMapSidNil sets the value for MapSid to be an explicit nil

### UnsetMapSid
`func (o *SyncV1ServiceSyncMapSyncMapPermission) UnsetMapSid()`

UnsetMapSid ensures that no value is present for MapSid, not even an explicit nil
### GetRead

`func (o *SyncV1ServiceSyncMapSyncMapPermission) GetRead() bool`

GetRead returns the Read field if non-nil, zero value otherwise.

### GetReadOk

`func (o *SyncV1ServiceSyncMapSyncMapPermission) GetReadOk() (*bool, bool)`

GetReadOk returns a tuple with the Read field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRead

`func (o *SyncV1ServiceSyncMapSyncMapPermission) SetRead(v bool)`

SetRead sets Read field to given value.

### HasRead

`func (o *SyncV1ServiceSyncMapSyncMapPermission) HasRead() bool`

HasRead returns a boolean if a field has been set.

### SetReadNil

`func (o *SyncV1ServiceSyncMapSyncMapPermission) SetReadNil(b bool)`

 SetReadNil sets the value for Read to be an explicit nil

### UnsetRead
`func (o *SyncV1ServiceSyncMapSyncMapPermission) UnsetRead()`

UnsetRead ensures that no value is present for Read, not even an explicit nil
### GetServiceSid

`func (o *SyncV1ServiceSyncMapSyncMapPermission) GetServiceSid() string`

GetServiceSid returns the ServiceSid field if non-nil, zero value otherwise.

### GetServiceSidOk

`func (o *SyncV1ServiceSyncMapSyncMapPermission) GetServiceSidOk() (*string, bool)`

GetServiceSidOk returns a tuple with the ServiceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSid

`func (o *SyncV1ServiceSyncMapSyncMapPermission) SetServiceSid(v string)`

SetServiceSid sets ServiceSid field to given value.

### HasServiceSid

`func (o *SyncV1ServiceSyncMapSyncMapPermission) HasServiceSid() bool`

HasServiceSid returns a boolean if a field has been set.

### SetServiceSidNil

`func (o *SyncV1ServiceSyncMapSyncMapPermission) SetServiceSidNil(b bool)`

 SetServiceSidNil sets the value for ServiceSid to be an explicit nil

### UnsetServiceSid
`func (o *SyncV1ServiceSyncMapSyncMapPermission) UnsetServiceSid()`

UnsetServiceSid ensures that no value is present for ServiceSid, not even an explicit nil
### GetUrl

`func (o *SyncV1ServiceSyncMapSyncMapPermission) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SyncV1ServiceSyncMapSyncMapPermission) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SyncV1ServiceSyncMapSyncMapPermission) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SyncV1ServiceSyncMapSyncMapPermission) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *SyncV1ServiceSyncMapSyncMapPermission) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *SyncV1ServiceSyncMapSyncMapPermission) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetWrite

`func (o *SyncV1ServiceSyncMapSyncMapPermission) GetWrite() bool`

GetWrite returns the Write field if non-nil, zero value otherwise.

### GetWriteOk

`func (o *SyncV1ServiceSyncMapSyncMapPermission) GetWriteOk() (*bool, bool)`

GetWriteOk returns a tuple with the Write field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWrite

`func (o *SyncV1ServiceSyncMapSyncMapPermission) SetWrite(v bool)`

SetWrite sets Write field to given value.

### HasWrite

`func (o *SyncV1ServiceSyncMapSyncMapPermission) HasWrite() bool`

HasWrite returns a boolean if a field has been set.

### SetWriteNil

`func (o *SyncV1ServiceSyncMapSyncMapPermission) SetWriteNil(b bool)`

 SetWriteNil sets the value for Write to be an explicit nil

### UnsetWrite
`func (o *SyncV1ServiceSyncMapSyncMapPermission) UnsetWrite()`

UnsetWrite ensures that no value is present for Write, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


