# SyncV1ServiceSyncListSyncListItem

## Properties

Name | Type | Description
------------ | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**CreatedBy** | Pointer to **NullableString** | The identity of the List Item&#39;s creator | [optional] 
**Data** | Pointer to **map[string]interface{}** | An arbitrary, schema-less object that the List Item stores | [optional] 
**DateCreated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was created | [optional] 
**DateExpires** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the List Item expires | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was last updated | [optional] 
**Index** | Pointer to **NullableInt32** | The automatically generated index of the List Item | [optional] 
**ListSid** | Pointer to **NullableString** | The SID of the Sync List that contains the List Item | [optional] 
**Revision** | Pointer to **NullableString** | The current revision of the item, represented as a string | [optional] 
**ServiceSid** | Pointer to **NullableString** | The SID of the Sync Service that the resource is associated with | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the List Item resource | [optional] 

## Methods

### NewSyncV1ServiceSyncListSyncListItem

`func NewSyncV1ServiceSyncListSyncListItem() *SyncV1ServiceSyncListSyncListItem`

NewSyncV1ServiceSyncListSyncListItem instantiates a new SyncV1ServiceSyncListSyncListItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncV1ServiceSyncListSyncListItemWithDefaults

`func NewSyncV1ServiceSyncListSyncListItemWithDefaults() *SyncV1ServiceSyncListSyncListItem`

NewSyncV1ServiceSyncListSyncListItemWithDefaults instantiates a new SyncV1ServiceSyncListSyncListItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *SyncV1ServiceSyncListSyncListItem) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *SyncV1ServiceSyncListSyncListItem) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *SyncV1ServiceSyncListSyncListItem) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *SyncV1ServiceSyncListSyncListItem) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *SyncV1ServiceSyncListSyncListItem) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *SyncV1ServiceSyncListSyncListItem) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetCreatedBy

`func (o *SyncV1ServiceSyncListSyncListItem) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *SyncV1ServiceSyncListSyncListItem) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *SyncV1ServiceSyncListSyncListItem) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *SyncV1ServiceSyncListSyncListItem) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *SyncV1ServiceSyncListSyncListItem) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *SyncV1ServiceSyncListSyncListItem) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetData

`func (o *SyncV1ServiceSyncListSyncListItem) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SyncV1ServiceSyncListSyncListItem) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SyncV1ServiceSyncListSyncListItem) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *SyncV1ServiceSyncListSyncListItem) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *SyncV1ServiceSyncListSyncListItem) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *SyncV1ServiceSyncListSyncListItem) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetDateCreated

`func (o *SyncV1ServiceSyncListSyncListItem) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *SyncV1ServiceSyncListSyncListItem) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *SyncV1ServiceSyncListSyncListItem) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *SyncV1ServiceSyncListSyncListItem) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *SyncV1ServiceSyncListSyncListItem) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *SyncV1ServiceSyncListSyncListItem) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateExpires

`func (o *SyncV1ServiceSyncListSyncListItem) GetDateExpires() time.Time`

GetDateExpires returns the DateExpires field if non-nil, zero value otherwise.

### GetDateExpiresOk

`func (o *SyncV1ServiceSyncListSyncListItem) GetDateExpiresOk() (*time.Time, bool)`

GetDateExpiresOk returns a tuple with the DateExpires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateExpires

`func (o *SyncV1ServiceSyncListSyncListItem) SetDateExpires(v time.Time)`

SetDateExpires sets DateExpires field to given value.

### HasDateExpires

`func (o *SyncV1ServiceSyncListSyncListItem) HasDateExpires() bool`

HasDateExpires returns a boolean if a field has been set.

### SetDateExpiresNil

`func (o *SyncV1ServiceSyncListSyncListItem) SetDateExpiresNil(b bool)`

 SetDateExpiresNil sets the value for DateExpires to be an explicit nil

### UnsetDateExpires
`func (o *SyncV1ServiceSyncListSyncListItem) UnsetDateExpires()`

UnsetDateExpires ensures that no value is present for DateExpires, not even an explicit nil
### GetDateUpdated

`func (o *SyncV1ServiceSyncListSyncListItem) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *SyncV1ServiceSyncListSyncListItem) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *SyncV1ServiceSyncListSyncListItem) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *SyncV1ServiceSyncListSyncListItem) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *SyncV1ServiceSyncListSyncListItem) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *SyncV1ServiceSyncListSyncListItem) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetIndex

`func (o *SyncV1ServiceSyncListSyncListItem) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *SyncV1ServiceSyncListSyncListItem) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *SyncV1ServiceSyncListSyncListItem) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *SyncV1ServiceSyncListSyncListItem) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### SetIndexNil

`func (o *SyncV1ServiceSyncListSyncListItem) SetIndexNil(b bool)`

 SetIndexNil sets the value for Index to be an explicit nil

### UnsetIndex
`func (o *SyncV1ServiceSyncListSyncListItem) UnsetIndex()`

UnsetIndex ensures that no value is present for Index, not even an explicit nil
### GetListSid

`func (o *SyncV1ServiceSyncListSyncListItem) GetListSid() string`

GetListSid returns the ListSid field if non-nil, zero value otherwise.

### GetListSidOk

`func (o *SyncV1ServiceSyncListSyncListItem) GetListSidOk() (*string, bool)`

GetListSidOk returns a tuple with the ListSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListSid

`func (o *SyncV1ServiceSyncListSyncListItem) SetListSid(v string)`

SetListSid sets ListSid field to given value.

### HasListSid

`func (o *SyncV1ServiceSyncListSyncListItem) HasListSid() bool`

HasListSid returns a boolean if a field has been set.

### SetListSidNil

`func (o *SyncV1ServiceSyncListSyncListItem) SetListSidNil(b bool)`

 SetListSidNil sets the value for ListSid to be an explicit nil

### UnsetListSid
`func (o *SyncV1ServiceSyncListSyncListItem) UnsetListSid()`

UnsetListSid ensures that no value is present for ListSid, not even an explicit nil
### GetRevision

`func (o *SyncV1ServiceSyncListSyncListItem) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *SyncV1ServiceSyncListSyncListItem) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *SyncV1ServiceSyncListSyncListItem) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *SyncV1ServiceSyncListSyncListItem) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### SetRevisionNil

`func (o *SyncV1ServiceSyncListSyncListItem) SetRevisionNil(b bool)`

 SetRevisionNil sets the value for Revision to be an explicit nil

### UnsetRevision
`func (o *SyncV1ServiceSyncListSyncListItem) UnsetRevision()`

UnsetRevision ensures that no value is present for Revision, not even an explicit nil
### GetServiceSid

`func (o *SyncV1ServiceSyncListSyncListItem) GetServiceSid() string`

GetServiceSid returns the ServiceSid field if non-nil, zero value otherwise.

### GetServiceSidOk

`func (o *SyncV1ServiceSyncListSyncListItem) GetServiceSidOk() (*string, bool)`

GetServiceSidOk returns a tuple with the ServiceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSid

`func (o *SyncV1ServiceSyncListSyncListItem) SetServiceSid(v string)`

SetServiceSid sets ServiceSid field to given value.

### HasServiceSid

`func (o *SyncV1ServiceSyncListSyncListItem) HasServiceSid() bool`

HasServiceSid returns a boolean if a field has been set.

### SetServiceSidNil

`func (o *SyncV1ServiceSyncListSyncListItem) SetServiceSidNil(b bool)`

 SetServiceSidNil sets the value for ServiceSid to be an explicit nil

### UnsetServiceSid
`func (o *SyncV1ServiceSyncListSyncListItem) UnsetServiceSid()`

UnsetServiceSid ensures that no value is present for ServiceSid, not even an explicit nil
### GetUrl

`func (o *SyncV1ServiceSyncListSyncListItem) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SyncV1ServiceSyncListSyncListItem) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SyncV1ServiceSyncListSyncListItem) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SyncV1ServiceSyncListSyncListItem) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *SyncV1ServiceSyncListSyncListItem) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *SyncV1ServiceSyncListSyncListItem) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


