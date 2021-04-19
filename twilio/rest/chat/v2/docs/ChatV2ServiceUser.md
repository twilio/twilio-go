# ChatV2ServiceUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**Attributes** | Pointer to **NullableString** | The JSON string that stores application-specific data | [optional] 
**DateCreated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was last updated | [optional] 
**FriendlyName** | Pointer to **NullableString** | The string that you assigned to describe the resource | [optional] 
**Identity** | Pointer to **NullableString** | The string that identifies the resource&#39;s User | [optional] 
**IsNotifiable** | Pointer to **NullableBool** | Whether the User has a potentially valid Push Notification registration for the Service instance | [optional] 
**IsOnline** | Pointer to **NullableBool** | Whether the User is actively connected to the Service instance and online | [optional] 
**JoinedChannelsCount** | Pointer to **NullableInt32** | The number of Channels the User is a Member of | [optional] 
**Links** | Pointer to **map[string]interface{}** | The absolute URLs of the Channel and Binding resources related to the user | [optional] 
**RoleSid** | Pointer to **NullableString** | The SID of the Role assigned to the user | [optional] 
**ServiceSid** | Pointer to **NullableString** | The SID of the Service that the resource is associated with | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the User resource | [optional] 

## Methods

### NewChatV2ServiceUser

`func NewChatV2ServiceUser() *ChatV2ServiceUser`

NewChatV2ServiceUser instantiates a new ChatV2ServiceUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatV2ServiceUserWithDefaults

`func NewChatV2ServiceUserWithDefaults() *ChatV2ServiceUser`

NewChatV2ServiceUserWithDefaults instantiates a new ChatV2ServiceUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *ChatV2ServiceUser) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *ChatV2ServiceUser) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *ChatV2ServiceUser) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *ChatV2ServiceUser) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *ChatV2ServiceUser) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *ChatV2ServiceUser) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetAttributes

`func (o *ChatV2ServiceUser) GetAttributes() string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ChatV2ServiceUser) GetAttributesOk() (*string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ChatV2ServiceUser) SetAttributes(v string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ChatV2ServiceUser) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *ChatV2ServiceUser) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *ChatV2ServiceUser) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil
### GetDateCreated

`func (o *ChatV2ServiceUser) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ChatV2ServiceUser) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ChatV2ServiceUser) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ChatV2ServiceUser) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *ChatV2ServiceUser) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *ChatV2ServiceUser) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *ChatV2ServiceUser) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *ChatV2ServiceUser) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *ChatV2ServiceUser) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *ChatV2ServiceUser) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *ChatV2ServiceUser) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *ChatV2ServiceUser) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetFriendlyName

`func (o *ChatV2ServiceUser) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *ChatV2ServiceUser) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *ChatV2ServiceUser) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *ChatV2ServiceUser) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *ChatV2ServiceUser) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *ChatV2ServiceUser) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetIdentity

`func (o *ChatV2ServiceUser) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *ChatV2ServiceUser) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *ChatV2ServiceUser) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *ChatV2ServiceUser) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### SetIdentityNil

`func (o *ChatV2ServiceUser) SetIdentityNil(b bool)`

 SetIdentityNil sets the value for Identity to be an explicit nil

### UnsetIdentity
`func (o *ChatV2ServiceUser) UnsetIdentity()`

UnsetIdentity ensures that no value is present for Identity, not even an explicit nil
### GetIsNotifiable

`func (o *ChatV2ServiceUser) GetIsNotifiable() bool`

GetIsNotifiable returns the IsNotifiable field if non-nil, zero value otherwise.

### GetIsNotifiableOk

`func (o *ChatV2ServiceUser) GetIsNotifiableOk() (*bool, bool)`

GetIsNotifiableOk returns a tuple with the IsNotifiable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNotifiable

`func (o *ChatV2ServiceUser) SetIsNotifiable(v bool)`

SetIsNotifiable sets IsNotifiable field to given value.

### HasIsNotifiable

`func (o *ChatV2ServiceUser) HasIsNotifiable() bool`

HasIsNotifiable returns a boolean if a field has been set.

### SetIsNotifiableNil

`func (o *ChatV2ServiceUser) SetIsNotifiableNil(b bool)`

 SetIsNotifiableNil sets the value for IsNotifiable to be an explicit nil

### UnsetIsNotifiable
`func (o *ChatV2ServiceUser) UnsetIsNotifiable()`

UnsetIsNotifiable ensures that no value is present for IsNotifiable, not even an explicit nil
### GetIsOnline

`func (o *ChatV2ServiceUser) GetIsOnline() bool`

GetIsOnline returns the IsOnline field if non-nil, zero value otherwise.

### GetIsOnlineOk

`func (o *ChatV2ServiceUser) GetIsOnlineOk() (*bool, bool)`

GetIsOnlineOk returns a tuple with the IsOnline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOnline

`func (o *ChatV2ServiceUser) SetIsOnline(v bool)`

SetIsOnline sets IsOnline field to given value.

### HasIsOnline

`func (o *ChatV2ServiceUser) HasIsOnline() bool`

HasIsOnline returns a boolean if a field has been set.

### SetIsOnlineNil

`func (o *ChatV2ServiceUser) SetIsOnlineNil(b bool)`

 SetIsOnlineNil sets the value for IsOnline to be an explicit nil

### UnsetIsOnline
`func (o *ChatV2ServiceUser) UnsetIsOnline()`

UnsetIsOnline ensures that no value is present for IsOnline, not even an explicit nil
### GetJoinedChannelsCount

`func (o *ChatV2ServiceUser) GetJoinedChannelsCount() int32`

GetJoinedChannelsCount returns the JoinedChannelsCount field if non-nil, zero value otherwise.

### GetJoinedChannelsCountOk

`func (o *ChatV2ServiceUser) GetJoinedChannelsCountOk() (*int32, bool)`

GetJoinedChannelsCountOk returns a tuple with the JoinedChannelsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinedChannelsCount

`func (o *ChatV2ServiceUser) SetJoinedChannelsCount(v int32)`

SetJoinedChannelsCount sets JoinedChannelsCount field to given value.

### HasJoinedChannelsCount

`func (o *ChatV2ServiceUser) HasJoinedChannelsCount() bool`

HasJoinedChannelsCount returns a boolean if a field has been set.

### SetJoinedChannelsCountNil

`func (o *ChatV2ServiceUser) SetJoinedChannelsCountNil(b bool)`

 SetJoinedChannelsCountNil sets the value for JoinedChannelsCount to be an explicit nil

### UnsetJoinedChannelsCount
`func (o *ChatV2ServiceUser) UnsetJoinedChannelsCount()`

UnsetJoinedChannelsCount ensures that no value is present for JoinedChannelsCount, not even an explicit nil
### GetLinks

`func (o *ChatV2ServiceUser) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ChatV2ServiceUser) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ChatV2ServiceUser) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ChatV2ServiceUser) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *ChatV2ServiceUser) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *ChatV2ServiceUser) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetRoleSid

`func (o *ChatV2ServiceUser) GetRoleSid() string`

GetRoleSid returns the RoleSid field if non-nil, zero value otherwise.

### GetRoleSidOk

`func (o *ChatV2ServiceUser) GetRoleSidOk() (*string, bool)`

GetRoleSidOk returns a tuple with the RoleSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleSid

`func (o *ChatV2ServiceUser) SetRoleSid(v string)`

SetRoleSid sets RoleSid field to given value.

### HasRoleSid

`func (o *ChatV2ServiceUser) HasRoleSid() bool`

HasRoleSid returns a boolean if a field has been set.

### SetRoleSidNil

`func (o *ChatV2ServiceUser) SetRoleSidNil(b bool)`

 SetRoleSidNil sets the value for RoleSid to be an explicit nil

### UnsetRoleSid
`func (o *ChatV2ServiceUser) UnsetRoleSid()`

UnsetRoleSid ensures that no value is present for RoleSid, not even an explicit nil
### GetServiceSid

`func (o *ChatV2ServiceUser) GetServiceSid() string`

GetServiceSid returns the ServiceSid field if non-nil, zero value otherwise.

### GetServiceSidOk

`func (o *ChatV2ServiceUser) GetServiceSidOk() (*string, bool)`

GetServiceSidOk returns a tuple with the ServiceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSid

`func (o *ChatV2ServiceUser) SetServiceSid(v string)`

SetServiceSid sets ServiceSid field to given value.

### HasServiceSid

`func (o *ChatV2ServiceUser) HasServiceSid() bool`

HasServiceSid returns a boolean if a field has been set.

### SetServiceSidNil

`func (o *ChatV2ServiceUser) SetServiceSidNil(b bool)`

 SetServiceSidNil sets the value for ServiceSid to be an explicit nil

### UnsetServiceSid
`func (o *ChatV2ServiceUser) UnsetServiceSid()`

UnsetServiceSid ensures that no value is present for ServiceSid, not even an explicit nil
### GetSid

`func (o *ChatV2ServiceUser) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *ChatV2ServiceUser) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *ChatV2ServiceUser) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *ChatV2ServiceUser) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *ChatV2ServiceUser) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *ChatV2ServiceUser) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetUrl

`func (o *ChatV2ServiceUser) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ChatV2ServiceUser) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ChatV2ServiceUser) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ChatV2ServiceUser) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *ChatV2ServiceUser) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *ChatV2ServiceUser) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


