# ProxyV1ServiceSessionParticipant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**DateCreated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was created | [optional] 
**DateDeleted** | Pointer to **NullableTime** | The ISO 8601 date the Participant was removed | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was last updated | [optional] 
**FriendlyName** | Pointer to **NullableString** | The string that you assigned to describe the participant | [optional] 
**Identifier** | Pointer to **NullableString** | The phone number or channel identifier of the Participant | [optional] 
**Links** | Pointer to **map[string]interface{}** | The URLs to resources related the participant | [optional] 
**ProxyIdentifier** | Pointer to **NullableString** | The phone number or short code of the participant&#39;s partner | [optional] 
**ProxyIdentifierSid** | Pointer to **NullableString** | The SID of the Proxy Identifier assigned to the Participant | [optional] 
**ServiceSid** | Pointer to **NullableString** | The SID of the resource&#39;s parent Service | [optional] 
**SessionSid** | Pointer to **NullableString** | The SID of the resource&#39;s parent Session | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the Participant resource | [optional] 

## Methods

### NewProxyV1ServiceSessionParticipant

`func NewProxyV1ServiceSessionParticipant() *ProxyV1ServiceSessionParticipant`

NewProxyV1ServiceSessionParticipant instantiates a new ProxyV1ServiceSessionParticipant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProxyV1ServiceSessionParticipantWithDefaults

`func NewProxyV1ServiceSessionParticipantWithDefaults() *ProxyV1ServiceSessionParticipant`

NewProxyV1ServiceSessionParticipantWithDefaults instantiates a new ProxyV1ServiceSessionParticipant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *ProxyV1ServiceSessionParticipant) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *ProxyV1ServiceSessionParticipant) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *ProxyV1ServiceSessionParticipant) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *ProxyV1ServiceSessionParticipant) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *ProxyV1ServiceSessionParticipant) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *ProxyV1ServiceSessionParticipant) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetDateCreated

`func (o *ProxyV1ServiceSessionParticipant) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ProxyV1ServiceSessionParticipant) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ProxyV1ServiceSessionParticipant) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ProxyV1ServiceSessionParticipant) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *ProxyV1ServiceSessionParticipant) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *ProxyV1ServiceSessionParticipant) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateDeleted

`func (o *ProxyV1ServiceSessionParticipant) GetDateDeleted() time.Time`

GetDateDeleted returns the DateDeleted field if non-nil, zero value otherwise.

### GetDateDeletedOk

`func (o *ProxyV1ServiceSessionParticipant) GetDateDeletedOk() (*time.Time, bool)`

GetDateDeletedOk returns a tuple with the DateDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateDeleted

`func (o *ProxyV1ServiceSessionParticipant) SetDateDeleted(v time.Time)`

SetDateDeleted sets DateDeleted field to given value.

### HasDateDeleted

`func (o *ProxyV1ServiceSessionParticipant) HasDateDeleted() bool`

HasDateDeleted returns a boolean if a field has been set.

### SetDateDeletedNil

`func (o *ProxyV1ServiceSessionParticipant) SetDateDeletedNil(b bool)`

 SetDateDeletedNil sets the value for DateDeleted to be an explicit nil

### UnsetDateDeleted
`func (o *ProxyV1ServiceSessionParticipant) UnsetDateDeleted()`

UnsetDateDeleted ensures that no value is present for DateDeleted, not even an explicit nil
### GetDateUpdated

`func (o *ProxyV1ServiceSessionParticipant) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *ProxyV1ServiceSessionParticipant) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *ProxyV1ServiceSessionParticipant) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *ProxyV1ServiceSessionParticipant) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *ProxyV1ServiceSessionParticipant) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *ProxyV1ServiceSessionParticipant) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetFriendlyName

`func (o *ProxyV1ServiceSessionParticipant) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *ProxyV1ServiceSessionParticipant) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *ProxyV1ServiceSessionParticipant) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *ProxyV1ServiceSessionParticipant) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *ProxyV1ServiceSessionParticipant) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *ProxyV1ServiceSessionParticipant) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetIdentifier

`func (o *ProxyV1ServiceSessionParticipant) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *ProxyV1ServiceSessionParticipant) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *ProxyV1ServiceSessionParticipant) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *ProxyV1ServiceSessionParticipant) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### SetIdentifierNil

`func (o *ProxyV1ServiceSessionParticipant) SetIdentifierNil(b bool)`

 SetIdentifierNil sets the value for Identifier to be an explicit nil

### UnsetIdentifier
`func (o *ProxyV1ServiceSessionParticipant) UnsetIdentifier()`

UnsetIdentifier ensures that no value is present for Identifier, not even an explicit nil
### GetLinks

`func (o *ProxyV1ServiceSessionParticipant) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ProxyV1ServiceSessionParticipant) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ProxyV1ServiceSessionParticipant) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ProxyV1ServiceSessionParticipant) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *ProxyV1ServiceSessionParticipant) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *ProxyV1ServiceSessionParticipant) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetProxyIdentifier

`func (o *ProxyV1ServiceSessionParticipant) GetProxyIdentifier() string`

GetProxyIdentifier returns the ProxyIdentifier field if non-nil, zero value otherwise.

### GetProxyIdentifierOk

`func (o *ProxyV1ServiceSessionParticipant) GetProxyIdentifierOk() (*string, bool)`

GetProxyIdentifierOk returns a tuple with the ProxyIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyIdentifier

`func (o *ProxyV1ServiceSessionParticipant) SetProxyIdentifier(v string)`

SetProxyIdentifier sets ProxyIdentifier field to given value.

### HasProxyIdentifier

`func (o *ProxyV1ServiceSessionParticipant) HasProxyIdentifier() bool`

HasProxyIdentifier returns a boolean if a field has been set.

### SetProxyIdentifierNil

`func (o *ProxyV1ServiceSessionParticipant) SetProxyIdentifierNil(b bool)`

 SetProxyIdentifierNil sets the value for ProxyIdentifier to be an explicit nil

### UnsetProxyIdentifier
`func (o *ProxyV1ServiceSessionParticipant) UnsetProxyIdentifier()`

UnsetProxyIdentifier ensures that no value is present for ProxyIdentifier, not even an explicit nil
### GetProxyIdentifierSid

`func (o *ProxyV1ServiceSessionParticipant) GetProxyIdentifierSid() string`

GetProxyIdentifierSid returns the ProxyIdentifierSid field if non-nil, zero value otherwise.

### GetProxyIdentifierSidOk

`func (o *ProxyV1ServiceSessionParticipant) GetProxyIdentifierSidOk() (*string, bool)`

GetProxyIdentifierSidOk returns a tuple with the ProxyIdentifierSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyIdentifierSid

`func (o *ProxyV1ServiceSessionParticipant) SetProxyIdentifierSid(v string)`

SetProxyIdentifierSid sets ProxyIdentifierSid field to given value.

### HasProxyIdentifierSid

`func (o *ProxyV1ServiceSessionParticipant) HasProxyIdentifierSid() bool`

HasProxyIdentifierSid returns a boolean if a field has been set.

### SetProxyIdentifierSidNil

`func (o *ProxyV1ServiceSessionParticipant) SetProxyIdentifierSidNil(b bool)`

 SetProxyIdentifierSidNil sets the value for ProxyIdentifierSid to be an explicit nil

### UnsetProxyIdentifierSid
`func (o *ProxyV1ServiceSessionParticipant) UnsetProxyIdentifierSid()`

UnsetProxyIdentifierSid ensures that no value is present for ProxyIdentifierSid, not even an explicit nil
### GetServiceSid

`func (o *ProxyV1ServiceSessionParticipant) GetServiceSid() string`

GetServiceSid returns the ServiceSid field if non-nil, zero value otherwise.

### GetServiceSidOk

`func (o *ProxyV1ServiceSessionParticipant) GetServiceSidOk() (*string, bool)`

GetServiceSidOk returns a tuple with the ServiceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSid

`func (o *ProxyV1ServiceSessionParticipant) SetServiceSid(v string)`

SetServiceSid sets ServiceSid field to given value.

### HasServiceSid

`func (o *ProxyV1ServiceSessionParticipant) HasServiceSid() bool`

HasServiceSid returns a boolean if a field has been set.

### SetServiceSidNil

`func (o *ProxyV1ServiceSessionParticipant) SetServiceSidNil(b bool)`

 SetServiceSidNil sets the value for ServiceSid to be an explicit nil

### UnsetServiceSid
`func (o *ProxyV1ServiceSessionParticipant) UnsetServiceSid()`

UnsetServiceSid ensures that no value is present for ServiceSid, not even an explicit nil
### GetSessionSid

`func (o *ProxyV1ServiceSessionParticipant) GetSessionSid() string`

GetSessionSid returns the SessionSid field if non-nil, zero value otherwise.

### GetSessionSidOk

`func (o *ProxyV1ServiceSessionParticipant) GetSessionSidOk() (*string, bool)`

GetSessionSidOk returns a tuple with the SessionSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionSid

`func (o *ProxyV1ServiceSessionParticipant) SetSessionSid(v string)`

SetSessionSid sets SessionSid field to given value.

### HasSessionSid

`func (o *ProxyV1ServiceSessionParticipant) HasSessionSid() bool`

HasSessionSid returns a boolean if a field has been set.

### SetSessionSidNil

`func (o *ProxyV1ServiceSessionParticipant) SetSessionSidNil(b bool)`

 SetSessionSidNil sets the value for SessionSid to be an explicit nil

### UnsetSessionSid
`func (o *ProxyV1ServiceSessionParticipant) UnsetSessionSid()`

UnsetSessionSid ensures that no value is present for SessionSid, not even an explicit nil
### GetSid

`func (o *ProxyV1ServiceSessionParticipant) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *ProxyV1ServiceSessionParticipant) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *ProxyV1ServiceSessionParticipant) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *ProxyV1ServiceSessionParticipant) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *ProxyV1ServiceSessionParticipant) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *ProxyV1ServiceSessionParticipant) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetUrl

`func (o *ProxyV1ServiceSessionParticipant) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ProxyV1ServiceSessionParticipant) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ProxyV1ServiceSessionParticipant) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ProxyV1ServiceSessionParticipant) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *ProxyV1ServiceSessionParticipant) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *ProxyV1ServiceSessionParticipant) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


