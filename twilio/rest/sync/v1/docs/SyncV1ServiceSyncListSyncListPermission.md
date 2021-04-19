# SyncV1ServiceSyncListSyncListPermission

## Properties

Name | Type | Description
------------ | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**Identity** | Pointer to **NullableString** | The identity of the user to whom the Sync List Permission applies | [optional] 
**ListSid** | Pointer to **NullableString** | The SID of the Sync List to which the Permission applies | [optional] 
**Manage** | Pointer to **NullableBool** | Manage access | [optional] 
**Read** | Pointer to **NullableBool** | Read access | [optional] 
**ServiceSid** | Pointer to **NullableString** | The SID of the Sync Service that the resource is associated with | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the Sync List Permission resource | [optional] 
**Write** | Pointer to **NullableBool** | Write access | [optional] 

## Methods

### NewSyncV1ServiceSyncListSyncListPermission

`func NewSyncV1ServiceSyncListSyncListPermission() *SyncV1ServiceSyncListSyncListPermission`

NewSyncV1ServiceSyncListSyncListPermission instantiates a new SyncV1ServiceSyncListSyncListPermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncV1ServiceSyncListSyncListPermissionWithDefaults

`func NewSyncV1ServiceSyncListSyncListPermissionWithDefaults() *SyncV1ServiceSyncListSyncListPermission`

NewSyncV1ServiceSyncListSyncListPermissionWithDefaults instantiates a new SyncV1ServiceSyncListSyncListPermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *SyncV1ServiceSyncListSyncListPermission) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *SyncV1ServiceSyncListSyncListPermission) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *SyncV1ServiceSyncListSyncListPermission) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *SyncV1ServiceSyncListSyncListPermission) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *SyncV1ServiceSyncListSyncListPermission) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *SyncV1ServiceSyncListSyncListPermission) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetIdentity

`func (o *SyncV1ServiceSyncListSyncListPermission) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *SyncV1ServiceSyncListSyncListPermission) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *SyncV1ServiceSyncListSyncListPermission) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *SyncV1ServiceSyncListSyncListPermission) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### SetIdentityNil

`func (o *SyncV1ServiceSyncListSyncListPermission) SetIdentityNil(b bool)`

 SetIdentityNil sets the value for Identity to be an explicit nil

### UnsetIdentity
`func (o *SyncV1ServiceSyncListSyncListPermission) UnsetIdentity()`

UnsetIdentity ensures that no value is present for Identity, not even an explicit nil
### GetListSid

`func (o *SyncV1ServiceSyncListSyncListPermission) GetListSid() string`

GetListSid returns the ListSid field if non-nil, zero value otherwise.

### GetListSidOk

`func (o *SyncV1ServiceSyncListSyncListPermission) GetListSidOk() (*string, bool)`

GetListSidOk returns a tuple with the ListSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListSid

`func (o *SyncV1ServiceSyncListSyncListPermission) SetListSid(v string)`

SetListSid sets ListSid field to given value.

### HasListSid

`func (o *SyncV1ServiceSyncListSyncListPermission) HasListSid() bool`

HasListSid returns a boolean if a field has been set.

### SetListSidNil

`func (o *SyncV1ServiceSyncListSyncListPermission) SetListSidNil(b bool)`

 SetListSidNil sets the value for ListSid to be an explicit nil

### UnsetListSid
`func (o *SyncV1ServiceSyncListSyncListPermission) UnsetListSid()`

UnsetListSid ensures that no value is present for ListSid, not even an explicit nil
### GetManage

`func (o *SyncV1ServiceSyncListSyncListPermission) GetManage() bool`

GetManage returns the Manage field if non-nil, zero value otherwise.

### GetManageOk

`func (o *SyncV1ServiceSyncListSyncListPermission) GetManageOk() (*bool, bool)`

GetManageOk returns a tuple with the Manage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManage

`func (o *SyncV1ServiceSyncListSyncListPermission) SetManage(v bool)`

SetManage sets Manage field to given value.

### HasManage

`func (o *SyncV1ServiceSyncListSyncListPermission) HasManage() bool`

HasManage returns a boolean if a field has been set.

### SetManageNil

`func (o *SyncV1ServiceSyncListSyncListPermission) SetManageNil(b bool)`

 SetManageNil sets the value for Manage to be an explicit nil

### UnsetManage
`func (o *SyncV1ServiceSyncListSyncListPermission) UnsetManage()`

UnsetManage ensures that no value is present for Manage, not even an explicit nil
### GetRead

`func (o *SyncV1ServiceSyncListSyncListPermission) GetRead() bool`

GetRead returns the Read field if non-nil, zero value otherwise.

### GetReadOk

`func (o *SyncV1ServiceSyncListSyncListPermission) GetReadOk() (*bool, bool)`

GetReadOk returns a tuple with the Read field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRead

`func (o *SyncV1ServiceSyncListSyncListPermission) SetRead(v bool)`

SetRead sets Read field to given value.

### HasRead

`func (o *SyncV1ServiceSyncListSyncListPermission) HasRead() bool`

HasRead returns a boolean if a field has been set.

### SetReadNil

`func (o *SyncV1ServiceSyncListSyncListPermission) SetReadNil(b bool)`

 SetReadNil sets the value for Read to be an explicit nil

### UnsetRead
`func (o *SyncV1ServiceSyncListSyncListPermission) UnsetRead()`

UnsetRead ensures that no value is present for Read, not even an explicit nil
### GetServiceSid

`func (o *SyncV1ServiceSyncListSyncListPermission) GetServiceSid() string`

GetServiceSid returns the ServiceSid field if non-nil, zero value otherwise.

### GetServiceSidOk

`func (o *SyncV1ServiceSyncListSyncListPermission) GetServiceSidOk() (*string, bool)`

GetServiceSidOk returns a tuple with the ServiceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSid

`func (o *SyncV1ServiceSyncListSyncListPermission) SetServiceSid(v string)`

SetServiceSid sets ServiceSid field to given value.

### HasServiceSid

`func (o *SyncV1ServiceSyncListSyncListPermission) HasServiceSid() bool`

HasServiceSid returns a boolean if a field has been set.

### SetServiceSidNil

`func (o *SyncV1ServiceSyncListSyncListPermission) SetServiceSidNil(b bool)`

 SetServiceSidNil sets the value for ServiceSid to be an explicit nil

### UnsetServiceSid
`func (o *SyncV1ServiceSyncListSyncListPermission) UnsetServiceSid()`

UnsetServiceSid ensures that no value is present for ServiceSid, not even an explicit nil
### GetUrl

`func (o *SyncV1ServiceSyncListSyncListPermission) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SyncV1ServiceSyncListSyncListPermission) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SyncV1ServiceSyncListSyncListPermission) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SyncV1ServiceSyncListSyncListPermission) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *SyncV1ServiceSyncListSyncListPermission) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *SyncV1ServiceSyncListSyncListPermission) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetWrite

`func (o *SyncV1ServiceSyncListSyncListPermission) GetWrite() bool`

GetWrite returns the Write field if non-nil, zero value otherwise.

### GetWriteOk

`func (o *SyncV1ServiceSyncListSyncListPermission) GetWriteOk() (*bool, bool)`

GetWriteOk returns a tuple with the Write field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWrite

`func (o *SyncV1ServiceSyncListSyncListPermission) SetWrite(v bool)`

SetWrite sets Write field to given value.

### HasWrite

`func (o *SyncV1ServiceSyncListSyncListPermission) HasWrite() bool`

HasWrite returns a boolean if a field has been set.

### SetWriteNil

`func (o *SyncV1ServiceSyncListSyncListPermission) SetWriteNil(b bool)`

 SetWriteNil sets the value for Write to be an explicit nil

### UnsetWrite
`func (o *SyncV1ServiceSyncListSyncListPermission) UnsetWrite()`

UnsetWrite ensures that no value is present for Write, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


