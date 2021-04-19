# SyncV1ServiceSyncList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**CreatedBy** | Pointer to **NullableString** | The identity of the Sync List&#39;s creator | [optional] 
**DateCreated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was created | [optional] 
**DateExpires** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the Sync List expires | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was last updated | [optional] 
**Links** | Pointer to **map[string]interface{}** | The URLs of the Sync List&#39;s nested resources | [optional] 
**Revision** | Pointer to **NullableString** | The current revision of the Sync List, represented as a string | [optional] 
**ServiceSid** | Pointer to **NullableString** | The SID of the Sync Service that the resource is associated with | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**UniqueName** | Pointer to **NullableString** | An application-defined string that uniquely identifies the resource | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the Sync List resource | [optional] 

## Methods

### NewSyncV1ServiceSyncList

`func NewSyncV1ServiceSyncList() *SyncV1ServiceSyncList`

NewSyncV1ServiceSyncList instantiates a new SyncV1ServiceSyncList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncV1ServiceSyncListWithDefaults

`func NewSyncV1ServiceSyncListWithDefaults() *SyncV1ServiceSyncList`

NewSyncV1ServiceSyncListWithDefaults instantiates a new SyncV1ServiceSyncList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *SyncV1ServiceSyncList) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *SyncV1ServiceSyncList) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *SyncV1ServiceSyncList) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *SyncV1ServiceSyncList) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *SyncV1ServiceSyncList) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *SyncV1ServiceSyncList) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetCreatedBy

`func (o *SyncV1ServiceSyncList) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *SyncV1ServiceSyncList) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *SyncV1ServiceSyncList) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *SyncV1ServiceSyncList) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *SyncV1ServiceSyncList) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *SyncV1ServiceSyncList) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetDateCreated

`func (o *SyncV1ServiceSyncList) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *SyncV1ServiceSyncList) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *SyncV1ServiceSyncList) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *SyncV1ServiceSyncList) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *SyncV1ServiceSyncList) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *SyncV1ServiceSyncList) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateExpires

`func (o *SyncV1ServiceSyncList) GetDateExpires() time.Time`

GetDateExpires returns the DateExpires field if non-nil, zero value otherwise.

### GetDateExpiresOk

`func (o *SyncV1ServiceSyncList) GetDateExpiresOk() (*time.Time, bool)`

GetDateExpiresOk returns a tuple with the DateExpires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateExpires

`func (o *SyncV1ServiceSyncList) SetDateExpires(v time.Time)`

SetDateExpires sets DateExpires field to given value.

### HasDateExpires

`func (o *SyncV1ServiceSyncList) HasDateExpires() bool`

HasDateExpires returns a boolean if a field has been set.

### SetDateExpiresNil

`func (o *SyncV1ServiceSyncList) SetDateExpiresNil(b bool)`

 SetDateExpiresNil sets the value for DateExpires to be an explicit nil

### UnsetDateExpires
`func (o *SyncV1ServiceSyncList) UnsetDateExpires()`

UnsetDateExpires ensures that no value is present for DateExpires, not even an explicit nil
### GetDateUpdated

`func (o *SyncV1ServiceSyncList) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *SyncV1ServiceSyncList) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *SyncV1ServiceSyncList) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *SyncV1ServiceSyncList) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *SyncV1ServiceSyncList) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *SyncV1ServiceSyncList) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetLinks

`func (o *SyncV1ServiceSyncList) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SyncV1ServiceSyncList) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SyncV1ServiceSyncList) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *SyncV1ServiceSyncList) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *SyncV1ServiceSyncList) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *SyncV1ServiceSyncList) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetRevision

`func (o *SyncV1ServiceSyncList) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *SyncV1ServiceSyncList) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *SyncV1ServiceSyncList) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *SyncV1ServiceSyncList) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### SetRevisionNil

`func (o *SyncV1ServiceSyncList) SetRevisionNil(b bool)`

 SetRevisionNil sets the value for Revision to be an explicit nil

### UnsetRevision
`func (o *SyncV1ServiceSyncList) UnsetRevision()`

UnsetRevision ensures that no value is present for Revision, not even an explicit nil
### GetServiceSid

`func (o *SyncV1ServiceSyncList) GetServiceSid() string`

GetServiceSid returns the ServiceSid field if non-nil, zero value otherwise.

### GetServiceSidOk

`func (o *SyncV1ServiceSyncList) GetServiceSidOk() (*string, bool)`

GetServiceSidOk returns a tuple with the ServiceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSid

`func (o *SyncV1ServiceSyncList) SetServiceSid(v string)`

SetServiceSid sets ServiceSid field to given value.

### HasServiceSid

`func (o *SyncV1ServiceSyncList) HasServiceSid() bool`

HasServiceSid returns a boolean if a field has been set.

### SetServiceSidNil

`func (o *SyncV1ServiceSyncList) SetServiceSidNil(b bool)`

 SetServiceSidNil sets the value for ServiceSid to be an explicit nil

### UnsetServiceSid
`func (o *SyncV1ServiceSyncList) UnsetServiceSid()`

UnsetServiceSid ensures that no value is present for ServiceSid, not even an explicit nil
### GetSid

`func (o *SyncV1ServiceSyncList) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *SyncV1ServiceSyncList) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *SyncV1ServiceSyncList) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *SyncV1ServiceSyncList) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *SyncV1ServiceSyncList) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *SyncV1ServiceSyncList) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetUniqueName

`func (o *SyncV1ServiceSyncList) GetUniqueName() string`

GetUniqueName returns the UniqueName field if non-nil, zero value otherwise.

### GetUniqueNameOk

`func (o *SyncV1ServiceSyncList) GetUniqueNameOk() (*string, bool)`

GetUniqueNameOk returns a tuple with the UniqueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueName

`func (o *SyncV1ServiceSyncList) SetUniqueName(v string)`

SetUniqueName sets UniqueName field to given value.

### HasUniqueName

`func (o *SyncV1ServiceSyncList) HasUniqueName() bool`

HasUniqueName returns a boolean if a field has been set.

### SetUniqueNameNil

`func (o *SyncV1ServiceSyncList) SetUniqueNameNil(b bool)`

 SetUniqueNameNil sets the value for UniqueName to be an explicit nil

### UnsetUniqueName
`func (o *SyncV1ServiceSyncList) UnsetUniqueName()`

UnsetUniqueName ensures that no value is present for UniqueName, not even an explicit nil
### GetUrl

`func (o *SyncV1ServiceSyncList) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SyncV1ServiceSyncList) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SyncV1ServiceSyncList) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SyncV1ServiceSyncList) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *SyncV1ServiceSyncList) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *SyncV1ServiceSyncList) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


