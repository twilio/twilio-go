# ConversationsV1ServiceServiceUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**Attributes** | Pointer to **NullableString** | The JSON Object string that stores application-specific data | [optional] 
**ChatServiceSid** | Pointer to **NullableString** | The SID of the Conversation Service that the resource is associated with | [optional] 
**DateCreated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was last updated | [optional] 
**FriendlyName** | Pointer to **NullableString** | The string that you assigned to describe the resource | [optional] 
**Identity** | Pointer to **NullableString** | The string that identifies the resource&#39;s User | [optional] 
**IsNotifiable** | Pointer to **NullableBool** | Whether the User has a potentially valid Push Notification registration for this Conversations Service | [optional] 
**IsOnline** | Pointer to **NullableBool** | Whether the User is actively connected to this Conversations Service and online | [optional] 
**RoleSid** | Pointer to **NullableString** | The SID of a service-level Role assigned to the user | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**Url** | Pointer to **NullableString** | An absolute URL for this user. | [optional] 

## Methods

### NewConversationsV1ServiceServiceUser

`func NewConversationsV1ServiceServiceUser() *ConversationsV1ServiceServiceUser`

NewConversationsV1ServiceServiceUser instantiates a new ConversationsV1ServiceServiceUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConversationsV1ServiceServiceUserWithDefaults

`func NewConversationsV1ServiceServiceUserWithDefaults() *ConversationsV1ServiceServiceUser`

NewConversationsV1ServiceServiceUserWithDefaults instantiates a new ConversationsV1ServiceServiceUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *ConversationsV1ServiceServiceUser) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *ConversationsV1ServiceServiceUser) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *ConversationsV1ServiceServiceUser) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *ConversationsV1ServiceServiceUser) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *ConversationsV1ServiceServiceUser) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *ConversationsV1ServiceServiceUser) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetAttributes

`func (o *ConversationsV1ServiceServiceUser) GetAttributes() string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ConversationsV1ServiceServiceUser) GetAttributesOk() (*string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ConversationsV1ServiceServiceUser) SetAttributes(v string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ConversationsV1ServiceServiceUser) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *ConversationsV1ServiceServiceUser) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *ConversationsV1ServiceServiceUser) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil
### GetChatServiceSid

`func (o *ConversationsV1ServiceServiceUser) GetChatServiceSid() string`

GetChatServiceSid returns the ChatServiceSid field if non-nil, zero value otherwise.

### GetChatServiceSidOk

`func (o *ConversationsV1ServiceServiceUser) GetChatServiceSidOk() (*string, bool)`

GetChatServiceSidOk returns a tuple with the ChatServiceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatServiceSid

`func (o *ConversationsV1ServiceServiceUser) SetChatServiceSid(v string)`

SetChatServiceSid sets ChatServiceSid field to given value.

### HasChatServiceSid

`func (o *ConversationsV1ServiceServiceUser) HasChatServiceSid() bool`

HasChatServiceSid returns a boolean if a field has been set.

### SetChatServiceSidNil

`func (o *ConversationsV1ServiceServiceUser) SetChatServiceSidNil(b bool)`

 SetChatServiceSidNil sets the value for ChatServiceSid to be an explicit nil

### UnsetChatServiceSid
`func (o *ConversationsV1ServiceServiceUser) UnsetChatServiceSid()`

UnsetChatServiceSid ensures that no value is present for ChatServiceSid, not even an explicit nil
### GetDateCreated

`func (o *ConversationsV1ServiceServiceUser) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ConversationsV1ServiceServiceUser) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ConversationsV1ServiceServiceUser) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ConversationsV1ServiceServiceUser) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *ConversationsV1ServiceServiceUser) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *ConversationsV1ServiceServiceUser) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *ConversationsV1ServiceServiceUser) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *ConversationsV1ServiceServiceUser) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *ConversationsV1ServiceServiceUser) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *ConversationsV1ServiceServiceUser) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *ConversationsV1ServiceServiceUser) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *ConversationsV1ServiceServiceUser) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetFriendlyName

`func (o *ConversationsV1ServiceServiceUser) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *ConversationsV1ServiceServiceUser) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *ConversationsV1ServiceServiceUser) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *ConversationsV1ServiceServiceUser) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *ConversationsV1ServiceServiceUser) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *ConversationsV1ServiceServiceUser) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetIdentity

`func (o *ConversationsV1ServiceServiceUser) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *ConversationsV1ServiceServiceUser) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *ConversationsV1ServiceServiceUser) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *ConversationsV1ServiceServiceUser) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### SetIdentityNil

`func (o *ConversationsV1ServiceServiceUser) SetIdentityNil(b bool)`

 SetIdentityNil sets the value for Identity to be an explicit nil

### UnsetIdentity
`func (o *ConversationsV1ServiceServiceUser) UnsetIdentity()`

UnsetIdentity ensures that no value is present for Identity, not even an explicit nil
### GetIsNotifiable

`func (o *ConversationsV1ServiceServiceUser) GetIsNotifiable() bool`

GetIsNotifiable returns the IsNotifiable field if non-nil, zero value otherwise.

### GetIsNotifiableOk

`func (o *ConversationsV1ServiceServiceUser) GetIsNotifiableOk() (*bool, bool)`

GetIsNotifiableOk returns a tuple with the IsNotifiable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNotifiable

`func (o *ConversationsV1ServiceServiceUser) SetIsNotifiable(v bool)`

SetIsNotifiable sets IsNotifiable field to given value.

### HasIsNotifiable

`func (o *ConversationsV1ServiceServiceUser) HasIsNotifiable() bool`

HasIsNotifiable returns a boolean if a field has been set.

### SetIsNotifiableNil

`func (o *ConversationsV1ServiceServiceUser) SetIsNotifiableNil(b bool)`

 SetIsNotifiableNil sets the value for IsNotifiable to be an explicit nil

### UnsetIsNotifiable
`func (o *ConversationsV1ServiceServiceUser) UnsetIsNotifiable()`

UnsetIsNotifiable ensures that no value is present for IsNotifiable, not even an explicit nil
### GetIsOnline

`func (o *ConversationsV1ServiceServiceUser) GetIsOnline() bool`

GetIsOnline returns the IsOnline field if non-nil, zero value otherwise.

### GetIsOnlineOk

`func (o *ConversationsV1ServiceServiceUser) GetIsOnlineOk() (*bool, bool)`

GetIsOnlineOk returns a tuple with the IsOnline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOnline

`func (o *ConversationsV1ServiceServiceUser) SetIsOnline(v bool)`

SetIsOnline sets IsOnline field to given value.

### HasIsOnline

`func (o *ConversationsV1ServiceServiceUser) HasIsOnline() bool`

HasIsOnline returns a boolean if a field has been set.

### SetIsOnlineNil

`func (o *ConversationsV1ServiceServiceUser) SetIsOnlineNil(b bool)`

 SetIsOnlineNil sets the value for IsOnline to be an explicit nil

### UnsetIsOnline
`func (o *ConversationsV1ServiceServiceUser) UnsetIsOnline()`

UnsetIsOnline ensures that no value is present for IsOnline, not even an explicit nil
### GetRoleSid

`func (o *ConversationsV1ServiceServiceUser) GetRoleSid() string`

GetRoleSid returns the RoleSid field if non-nil, zero value otherwise.

### GetRoleSidOk

`func (o *ConversationsV1ServiceServiceUser) GetRoleSidOk() (*string, bool)`

GetRoleSidOk returns a tuple with the RoleSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleSid

`func (o *ConversationsV1ServiceServiceUser) SetRoleSid(v string)`

SetRoleSid sets RoleSid field to given value.

### HasRoleSid

`func (o *ConversationsV1ServiceServiceUser) HasRoleSid() bool`

HasRoleSid returns a boolean if a field has been set.

### SetRoleSidNil

`func (o *ConversationsV1ServiceServiceUser) SetRoleSidNil(b bool)`

 SetRoleSidNil sets the value for RoleSid to be an explicit nil

### UnsetRoleSid
`func (o *ConversationsV1ServiceServiceUser) UnsetRoleSid()`

UnsetRoleSid ensures that no value is present for RoleSid, not even an explicit nil
### GetSid

`func (o *ConversationsV1ServiceServiceUser) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *ConversationsV1ServiceServiceUser) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *ConversationsV1ServiceServiceUser) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *ConversationsV1ServiceServiceUser) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *ConversationsV1ServiceServiceUser) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *ConversationsV1ServiceServiceUser) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetUrl

`func (o *ConversationsV1ServiceServiceUser) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ConversationsV1ServiceServiceUser) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ConversationsV1ServiceServiceUser) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ConversationsV1ServiceServiceUser) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *ConversationsV1ServiceServiceUser) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *ConversationsV1ServiceServiceUser) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


