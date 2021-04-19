# StudioV1Flow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**DateCreated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was last updated | [optional] 
**FriendlyName** | Pointer to **NullableString** | The string that you assigned to describe the Flow | [optional] 
**Links** | Pointer to **map[string]interface{}** | Nested resource URLs | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**Status** | Pointer to **NullableString** | The status of the Flow | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the resource | [optional] 
**Version** | Pointer to **NullableInt32** | The latest version number of the Flow&#39;s definition | [optional] 

## Methods

### NewStudioV1Flow

`func NewStudioV1Flow() *StudioV1Flow`

NewStudioV1Flow instantiates a new StudioV1Flow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStudioV1FlowWithDefaults

`func NewStudioV1FlowWithDefaults() *StudioV1Flow`

NewStudioV1FlowWithDefaults instantiates a new StudioV1Flow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *StudioV1Flow) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *StudioV1Flow) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *StudioV1Flow) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *StudioV1Flow) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *StudioV1Flow) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *StudioV1Flow) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetDateCreated

`func (o *StudioV1Flow) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *StudioV1Flow) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *StudioV1Flow) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *StudioV1Flow) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *StudioV1Flow) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *StudioV1Flow) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *StudioV1Flow) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *StudioV1Flow) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *StudioV1Flow) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *StudioV1Flow) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *StudioV1Flow) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *StudioV1Flow) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetFriendlyName

`func (o *StudioV1Flow) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *StudioV1Flow) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *StudioV1Flow) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *StudioV1Flow) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *StudioV1Flow) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *StudioV1Flow) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetLinks

`func (o *StudioV1Flow) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *StudioV1Flow) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *StudioV1Flow) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *StudioV1Flow) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *StudioV1Flow) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *StudioV1Flow) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetSid

`func (o *StudioV1Flow) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *StudioV1Flow) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *StudioV1Flow) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *StudioV1Flow) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *StudioV1Flow) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *StudioV1Flow) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetStatus

`func (o *StudioV1Flow) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StudioV1Flow) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StudioV1Flow) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StudioV1Flow) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *StudioV1Flow) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *StudioV1Flow) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetUrl

`func (o *StudioV1Flow) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *StudioV1Flow) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *StudioV1Flow) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *StudioV1Flow) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *StudioV1Flow) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *StudioV1Flow) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetVersion

`func (o *StudioV1Flow) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *StudioV1Flow) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *StudioV1Flow) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *StudioV1Flow) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *StudioV1Flow) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *StudioV1Flow) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


