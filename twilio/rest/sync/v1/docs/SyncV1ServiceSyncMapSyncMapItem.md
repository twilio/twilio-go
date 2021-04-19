# SyncV1ServiceSyncMapSyncMapItem

## Properties

Name | Type | Description
------------ | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**CreatedBy** | Pointer to **NullableString** | The identity of the Map Item&#39;s creator | [optional] 
**Data** | Pointer to **map[string]interface{}** | An arbitrary, schema-less object that the Map Item stores | [optional] 
**DateCreated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was created | [optional] 
**DateExpires** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the Map Item expires | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was last updated | [optional] 
**Key** | Pointer to **NullableString** | The unique, user-defined key for the Map Item | [optional] 
**MapSid** | Pointer to **NullableString** | The SID of the Sync Map that contains the Map Item | [optional] 
**Revision** | Pointer to **NullableString** | The current revision of the Map Item, represented as a string | [optional] 
**ServiceSid** | Pointer to **NullableString** | The SID of the Sync Service that the resource is associated with | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the Map Item resource | [optional] 

## Methods

### NewSyncV1ServiceSyncMapSyncMapItem

`func NewSyncV1ServiceSyncMapSyncMapItem() *SyncV1ServiceSyncMapSyncMapItem`

NewSyncV1ServiceSyncMapSyncMapItem instantiates a new SyncV1ServiceSyncMapSyncMapItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncV1ServiceSyncMapSyncMapItemWithDefaults

`func NewSyncV1ServiceSyncMapSyncMapItemWithDefaults() *SyncV1ServiceSyncMapSyncMapItem`

NewSyncV1ServiceSyncMapSyncMapItemWithDefaults instantiates a new SyncV1ServiceSyncMapSyncMapItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *SyncV1ServiceSyncMapSyncMapItem) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *SyncV1ServiceSyncMapSyncMapItem) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *SyncV1ServiceSyncMapSyncMapItem) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *SyncV1ServiceSyncMapSyncMapItem) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *SyncV1ServiceSyncMapSyncMapItem) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *SyncV1ServiceSyncMapSyncMapItem) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetCreatedBy

`func (o *SyncV1ServiceSyncMapSyncMapItem) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *SyncV1ServiceSyncMapSyncMapItem) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *SyncV1ServiceSyncMapSyncMapItem) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *SyncV1ServiceSyncMapSyncMapItem) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *SyncV1ServiceSyncMapSyncMapItem) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *SyncV1ServiceSyncMapSyncMapItem) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetData

`func (o *SyncV1ServiceSyncMapSyncMapItem) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SyncV1ServiceSyncMapSyncMapItem) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SyncV1ServiceSyncMapSyncMapItem) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *SyncV1ServiceSyncMapSyncMapItem) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *SyncV1ServiceSyncMapSyncMapItem) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *SyncV1ServiceSyncMapSyncMapItem) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetDateCreated

`func (o *SyncV1ServiceSyncMapSyncMapItem) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *SyncV1ServiceSyncMapSyncMapItem) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *SyncV1ServiceSyncMapSyncMapItem) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *SyncV1ServiceSyncMapSyncMapItem) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *SyncV1ServiceSyncMapSyncMapItem) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *SyncV1ServiceSyncMapSyncMapItem) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateExpires

`func (o *SyncV1ServiceSyncMapSyncMapItem) GetDateExpires() time.Time`

GetDateExpires returns the DateExpires field if non-nil, zero value otherwise.

### GetDateExpiresOk

`func (o *SyncV1ServiceSyncMapSyncMapItem) GetDateExpiresOk() (*time.Time, bool)`

GetDateExpiresOk returns a tuple with the DateExpires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateExpires

`func (o *SyncV1ServiceSyncMapSyncMapItem) SetDateExpires(v time.Time)`

SetDateExpires sets DateExpires field to given value.

### HasDateExpires

`func (o *SyncV1ServiceSyncMapSyncMapItem) HasDateExpires() bool`

HasDateExpires returns a boolean if a field has been set.

### SetDateExpiresNil

`func (o *SyncV1ServiceSyncMapSyncMapItem) SetDateExpiresNil(b bool)`

 SetDateExpiresNil sets the value for DateExpires to be an explicit nil

### UnsetDateExpires
`func (o *SyncV1ServiceSyncMapSyncMapItem) UnsetDateExpires()`

UnsetDateExpires ensures that no value is present for DateExpires, not even an explicit nil
### GetDateUpdated

`func (o *SyncV1ServiceSyncMapSyncMapItem) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *SyncV1ServiceSyncMapSyncMapItem) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *SyncV1ServiceSyncMapSyncMapItem) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *SyncV1ServiceSyncMapSyncMapItem) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *SyncV1ServiceSyncMapSyncMapItem) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *SyncV1ServiceSyncMapSyncMapItem) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetKey

`func (o *SyncV1ServiceSyncMapSyncMapItem) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SyncV1ServiceSyncMapSyncMapItem) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SyncV1ServiceSyncMapSyncMapItem) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *SyncV1ServiceSyncMapSyncMapItem) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *SyncV1ServiceSyncMapSyncMapItem) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *SyncV1ServiceSyncMapSyncMapItem) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetMapSid

`func (o *SyncV1ServiceSyncMapSyncMapItem) GetMapSid() string`

GetMapSid returns the MapSid field if non-nil, zero value otherwise.

### GetMapSidOk

`func (o *SyncV1ServiceSyncMapSyncMapItem) GetMapSidOk() (*string, bool)`

GetMapSidOk returns a tuple with the MapSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapSid

`func (o *SyncV1ServiceSyncMapSyncMapItem) SetMapSid(v string)`

SetMapSid sets MapSid field to given value.

### HasMapSid

`func (o *SyncV1ServiceSyncMapSyncMapItem) HasMapSid() bool`

HasMapSid returns a boolean if a field has been set.

### SetMapSidNil

`func (o *SyncV1ServiceSyncMapSyncMapItem) SetMapSidNil(b bool)`

 SetMapSidNil sets the value for MapSid to be an explicit nil

### UnsetMapSid
`func (o *SyncV1ServiceSyncMapSyncMapItem) UnsetMapSid()`

UnsetMapSid ensures that no value is present for MapSid, not even an explicit nil
### GetRevision

`func (o *SyncV1ServiceSyncMapSyncMapItem) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *SyncV1ServiceSyncMapSyncMapItem) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *SyncV1ServiceSyncMapSyncMapItem) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *SyncV1ServiceSyncMapSyncMapItem) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### SetRevisionNil

`func (o *SyncV1ServiceSyncMapSyncMapItem) SetRevisionNil(b bool)`

 SetRevisionNil sets the value for Revision to be an explicit nil

### UnsetRevision
`func (o *SyncV1ServiceSyncMapSyncMapItem) UnsetRevision()`

UnsetRevision ensures that no value is present for Revision, not even an explicit nil
### GetServiceSid

`func (o *SyncV1ServiceSyncMapSyncMapItem) GetServiceSid() string`

GetServiceSid returns the ServiceSid field if non-nil, zero value otherwise.

### GetServiceSidOk

`func (o *SyncV1ServiceSyncMapSyncMapItem) GetServiceSidOk() (*string, bool)`

GetServiceSidOk returns a tuple with the ServiceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSid

`func (o *SyncV1ServiceSyncMapSyncMapItem) SetServiceSid(v string)`

SetServiceSid sets ServiceSid field to given value.

### HasServiceSid

`func (o *SyncV1ServiceSyncMapSyncMapItem) HasServiceSid() bool`

HasServiceSid returns a boolean if a field has been set.

### SetServiceSidNil

`func (o *SyncV1ServiceSyncMapSyncMapItem) SetServiceSidNil(b bool)`

 SetServiceSidNil sets the value for ServiceSid to be an explicit nil

### UnsetServiceSid
`func (o *SyncV1ServiceSyncMapSyncMapItem) UnsetServiceSid()`

UnsetServiceSid ensures that no value is present for ServiceSid, not even an explicit nil
### GetUrl

`func (o *SyncV1ServiceSyncMapSyncMapItem) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SyncV1ServiceSyncMapSyncMapItem) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SyncV1ServiceSyncMapSyncMapItem) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SyncV1ServiceSyncMapSyncMapItem) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *SyncV1ServiceSyncMapSyncMapItem) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *SyncV1ServiceSyncMapSyncMapItem) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


