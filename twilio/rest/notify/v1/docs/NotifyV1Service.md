# NotifyV1Service

## Properties

Name | Type | Description
------------ | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**AlexaSkillId** | Pointer to **NullableString** | Deprecated | [optional] 
**ApnCredentialSid** | Pointer to **NullableString** | The SID of the Credential to use for APN Bindings | [optional] 
**DateCreated** | Pointer to **NullableTime** | The RFC 2822 date and time in GMT when the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The RFC 2822 date and time in GMT when the resource was last updated | [optional] 
**DefaultAlexaNotificationProtocolVersion** | Pointer to **NullableString** | Deprecated | [optional] 
**DefaultApnNotificationProtocolVersion** | Pointer to **NullableString** | The protocol version to use for sending APNS notifications | [optional] 
**DefaultFcmNotificationProtocolVersion** | Pointer to **NullableString** | The protocol version to use for sending FCM notifications | [optional] 
**DefaultGcmNotificationProtocolVersion** | Pointer to **NullableString** | The protocol version to use for sending GCM notifications | [optional] 
**DeliveryCallbackEnabled** | Pointer to **NullableBool** | Enable delivery callbacks | [optional] 
**DeliveryCallbackUrl** | Pointer to **NullableString** | Webhook URL | [optional] 
**FacebookMessengerPageId** | Pointer to **NullableString** | Deprecated | [optional] 
**FcmCredentialSid** | Pointer to **NullableString** | The SID of the Credential to use for FCM Bindings | [optional] 
**FriendlyName** | Pointer to **NullableString** | The string that you assigned to describe the resource | [optional] 
**GcmCredentialSid** | Pointer to **NullableString** | The SID of the Credential to use for GCM Bindings | [optional] 
**Links** | Pointer to **map[string]interface{}** | The URLs of the resources related to the service | [optional] 
**LogEnabled** | Pointer to **NullableBool** | Whether to log notifications | [optional] 
**MessagingServiceSid** | Pointer to **NullableString** | The SID of the Messaging Service to use for SMS Bindings | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the Service resource | [optional] 

## Methods

### NewNotifyV1Service

`func NewNotifyV1Service() *NotifyV1Service`

NewNotifyV1Service instantiates a new NotifyV1Service object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotifyV1ServiceWithDefaults

`func NewNotifyV1ServiceWithDefaults() *NotifyV1Service`

NewNotifyV1ServiceWithDefaults instantiates a new NotifyV1Service object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *NotifyV1Service) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *NotifyV1Service) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *NotifyV1Service) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *NotifyV1Service) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *NotifyV1Service) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *NotifyV1Service) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetAlexaSkillId

`func (o *NotifyV1Service) GetAlexaSkillId() string`

GetAlexaSkillId returns the AlexaSkillId field if non-nil, zero value otherwise.

### GetAlexaSkillIdOk

`func (o *NotifyV1Service) GetAlexaSkillIdOk() (*string, bool)`

GetAlexaSkillIdOk returns a tuple with the AlexaSkillId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlexaSkillId

`func (o *NotifyV1Service) SetAlexaSkillId(v string)`

SetAlexaSkillId sets AlexaSkillId field to given value.

### HasAlexaSkillId

`func (o *NotifyV1Service) HasAlexaSkillId() bool`

HasAlexaSkillId returns a boolean if a field has been set.

### SetAlexaSkillIdNil

`func (o *NotifyV1Service) SetAlexaSkillIdNil(b bool)`

 SetAlexaSkillIdNil sets the value for AlexaSkillId to be an explicit nil

### UnsetAlexaSkillId
`func (o *NotifyV1Service) UnsetAlexaSkillId()`

UnsetAlexaSkillId ensures that no value is present for AlexaSkillId, not even an explicit nil
### GetApnCredentialSid

`func (o *NotifyV1Service) GetApnCredentialSid() string`

GetApnCredentialSid returns the ApnCredentialSid field if non-nil, zero value otherwise.

### GetApnCredentialSidOk

`func (o *NotifyV1Service) GetApnCredentialSidOk() (*string, bool)`

GetApnCredentialSidOk returns a tuple with the ApnCredentialSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApnCredentialSid

`func (o *NotifyV1Service) SetApnCredentialSid(v string)`

SetApnCredentialSid sets ApnCredentialSid field to given value.

### HasApnCredentialSid

`func (o *NotifyV1Service) HasApnCredentialSid() bool`

HasApnCredentialSid returns a boolean if a field has been set.

### SetApnCredentialSidNil

`func (o *NotifyV1Service) SetApnCredentialSidNil(b bool)`

 SetApnCredentialSidNil sets the value for ApnCredentialSid to be an explicit nil

### UnsetApnCredentialSid
`func (o *NotifyV1Service) UnsetApnCredentialSid()`

UnsetApnCredentialSid ensures that no value is present for ApnCredentialSid, not even an explicit nil
### GetDateCreated

`func (o *NotifyV1Service) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *NotifyV1Service) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *NotifyV1Service) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *NotifyV1Service) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *NotifyV1Service) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *NotifyV1Service) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *NotifyV1Service) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *NotifyV1Service) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *NotifyV1Service) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *NotifyV1Service) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *NotifyV1Service) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *NotifyV1Service) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetDefaultAlexaNotificationProtocolVersion

`func (o *NotifyV1Service) GetDefaultAlexaNotificationProtocolVersion() string`

GetDefaultAlexaNotificationProtocolVersion returns the DefaultAlexaNotificationProtocolVersion field if non-nil, zero value otherwise.

### GetDefaultAlexaNotificationProtocolVersionOk

`func (o *NotifyV1Service) GetDefaultAlexaNotificationProtocolVersionOk() (*string, bool)`

GetDefaultAlexaNotificationProtocolVersionOk returns a tuple with the DefaultAlexaNotificationProtocolVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAlexaNotificationProtocolVersion

`func (o *NotifyV1Service) SetDefaultAlexaNotificationProtocolVersion(v string)`

SetDefaultAlexaNotificationProtocolVersion sets DefaultAlexaNotificationProtocolVersion field to given value.

### HasDefaultAlexaNotificationProtocolVersion

`func (o *NotifyV1Service) HasDefaultAlexaNotificationProtocolVersion() bool`

HasDefaultAlexaNotificationProtocolVersion returns a boolean if a field has been set.

### SetDefaultAlexaNotificationProtocolVersionNil

`func (o *NotifyV1Service) SetDefaultAlexaNotificationProtocolVersionNil(b bool)`

 SetDefaultAlexaNotificationProtocolVersionNil sets the value for DefaultAlexaNotificationProtocolVersion to be an explicit nil

### UnsetDefaultAlexaNotificationProtocolVersion
`func (o *NotifyV1Service) UnsetDefaultAlexaNotificationProtocolVersion()`

UnsetDefaultAlexaNotificationProtocolVersion ensures that no value is present for DefaultAlexaNotificationProtocolVersion, not even an explicit nil
### GetDefaultApnNotificationProtocolVersion

`func (o *NotifyV1Service) GetDefaultApnNotificationProtocolVersion() string`

GetDefaultApnNotificationProtocolVersion returns the DefaultApnNotificationProtocolVersion field if non-nil, zero value otherwise.

### GetDefaultApnNotificationProtocolVersionOk

`func (o *NotifyV1Service) GetDefaultApnNotificationProtocolVersionOk() (*string, bool)`

GetDefaultApnNotificationProtocolVersionOk returns a tuple with the DefaultApnNotificationProtocolVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultApnNotificationProtocolVersion

`func (o *NotifyV1Service) SetDefaultApnNotificationProtocolVersion(v string)`

SetDefaultApnNotificationProtocolVersion sets DefaultApnNotificationProtocolVersion field to given value.

### HasDefaultApnNotificationProtocolVersion

`func (o *NotifyV1Service) HasDefaultApnNotificationProtocolVersion() bool`

HasDefaultApnNotificationProtocolVersion returns a boolean if a field has been set.

### SetDefaultApnNotificationProtocolVersionNil

`func (o *NotifyV1Service) SetDefaultApnNotificationProtocolVersionNil(b bool)`

 SetDefaultApnNotificationProtocolVersionNil sets the value for DefaultApnNotificationProtocolVersion to be an explicit nil

### UnsetDefaultApnNotificationProtocolVersion
`func (o *NotifyV1Service) UnsetDefaultApnNotificationProtocolVersion()`

UnsetDefaultApnNotificationProtocolVersion ensures that no value is present for DefaultApnNotificationProtocolVersion, not even an explicit nil
### GetDefaultFcmNotificationProtocolVersion

`func (o *NotifyV1Service) GetDefaultFcmNotificationProtocolVersion() string`

GetDefaultFcmNotificationProtocolVersion returns the DefaultFcmNotificationProtocolVersion field if non-nil, zero value otherwise.

### GetDefaultFcmNotificationProtocolVersionOk

`func (o *NotifyV1Service) GetDefaultFcmNotificationProtocolVersionOk() (*string, bool)`

GetDefaultFcmNotificationProtocolVersionOk returns a tuple with the DefaultFcmNotificationProtocolVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFcmNotificationProtocolVersion

`func (o *NotifyV1Service) SetDefaultFcmNotificationProtocolVersion(v string)`

SetDefaultFcmNotificationProtocolVersion sets DefaultFcmNotificationProtocolVersion field to given value.

### HasDefaultFcmNotificationProtocolVersion

`func (o *NotifyV1Service) HasDefaultFcmNotificationProtocolVersion() bool`

HasDefaultFcmNotificationProtocolVersion returns a boolean if a field has been set.

### SetDefaultFcmNotificationProtocolVersionNil

`func (o *NotifyV1Service) SetDefaultFcmNotificationProtocolVersionNil(b bool)`

 SetDefaultFcmNotificationProtocolVersionNil sets the value for DefaultFcmNotificationProtocolVersion to be an explicit nil

### UnsetDefaultFcmNotificationProtocolVersion
`func (o *NotifyV1Service) UnsetDefaultFcmNotificationProtocolVersion()`

UnsetDefaultFcmNotificationProtocolVersion ensures that no value is present for DefaultFcmNotificationProtocolVersion, not even an explicit nil
### GetDefaultGcmNotificationProtocolVersion

`func (o *NotifyV1Service) GetDefaultGcmNotificationProtocolVersion() string`

GetDefaultGcmNotificationProtocolVersion returns the DefaultGcmNotificationProtocolVersion field if non-nil, zero value otherwise.

### GetDefaultGcmNotificationProtocolVersionOk

`func (o *NotifyV1Service) GetDefaultGcmNotificationProtocolVersionOk() (*string, bool)`

GetDefaultGcmNotificationProtocolVersionOk returns a tuple with the DefaultGcmNotificationProtocolVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultGcmNotificationProtocolVersion

`func (o *NotifyV1Service) SetDefaultGcmNotificationProtocolVersion(v string)`

SetDefaultGcmNotificationProtocolVersion sets DefaultGcmNotificationProtocolVersion field to given value.

### HasDefaultGcmNotificationProtocolVersion

`func (o *NotifyV1Service) HasDefaultGcmNotificationProtocolVersion() bool`

HasDefaultGcmNotificationProtocolVersion returns a boolean if a field has been set.

### SetDefaultGcmNotificationProtocolVersionNil

`func (o *NotifyV1Service) SetDefaultGcmNotificationProtocolVersionNil(b bool)`

 SetDefaultGcmNotificationProtocolVersionNil sets the value for DefaultGcmNotificationProtocolVersion to be an explicit nil

### UnsetDefaultGcmNotificationProtocolVersion
`func (o *NotifyV1Service) UnsetDefaultGcmNotificationProtocolVersion()`

UnsetDefaultGcmNotificationProtocolVersion ensures that no value is present for DefaultGcmNotificationProtocolVersion, not even an explicit nil
### GetDeliveryCallbackEnabled

`func (o *NotifyV1Service) GetDeliveryCallbackEnabled() bool`

GetDeliveryCallbackEnabled returns the DeliveryCallbackEnabled field if non-nil, zero value otherwise.

### GetDeliveryCallbackEnabledOk

`func (o *NotifyV1Service) GetDeliveryCallbackEnabledOk() (*bool, bool)`

GetDeliveryCallbackEnabledOk returns a tuple with the DeliveryCallbackEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryCallbackEnabled

`func (o *NotifyV1Service) SetDeliveryCallbackEnabled(v bool)`

SetDeliveryCallbackEnabled sets DeliveryCallbackEnabled field to given value.

### HasDeliveryCallbackEnabled

`func (o *NotifyV1Service) HasDeliveryCallbackEnabled() bool`

HasDeliveryCallbackEnabled returns a boolean if a field has been set.

### SetDeliveryCallbackEnabledNil

`func (o *NotifyV1Service) SetDeliveryCallbackEnabledNil(b bool)`

 SetDeliveryCallbackEnabledNil sets the value for DeliveryCallbackEnabled to be an explicit nil

### UnsetDeliveryCallbackEnabled
`func (o *NotifyV1Service) UnsetDeliveryCallbackEnabled()`

UnsetDeliveryCallbackEnabled ensures that no value is present for DeliveryCallbackEnabled, not even an explicit nil
### GetDeliveryCallbackUrl

`func (o *NotifyV1Service) GetDeliveryCallbackUrl() string`

GetDeliveryCallbackUrl returns the DeliveryCallbackUrl field if non-nil, zero value otherwise.

### GetDeliveryCallbackUrlOk

`func (o *NotifyV1Service) GetDeliveryCallbackUrlOk() (*string, bool)`

GetDeliveryCallbackUrlOk returns a tuple with the DeliveryCallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryCallbackUrl

`func (o *NotifyV1Service) SetDeliveryCallbackUrl(v string)`

SetDeliveryCallbackUrl sets DeliveryCallbackUrl field to given value.

### HasDeliveryCallbackUrl

`func (o *NotifyV1Service) HasDeliveryCallbackUrl() bool`

HasDeliveryCallbackUrl returns a boolean if a field has been set.

### SetDeliveryCallbackUrlNil

`func (o *NotifyV1Service) SetDeliveryCallbackUrlNil(b bool)`

 SetDeliveryCallbackUrlNil sets the value for DeliveryCallbackUrl to be an explicit nil

### UnsetDeliveryCallbackUrl
`func (o *NotifyV1Service) UnsetDeliveryCallbackUrl()`

UnsetDeliveryCallbackUrl ensures that no value is present for DeliveryCallbackUrl, not even an explicit nil
### GetFacebookMessengerPageId

`func (o *NotifyV1Service) GetFacebookMessengerPageId() string`

GetFacebookMessengerPageId returns the FacebookMessengerPageId field if non-nil, zero value otherwise.

### GetFacebookMessengerPageIdOk

`func (o *NotifyV1Service) GetFacebookMessengerPageIdOk() (*string, bool)`

GetFacebookMessengerPageIdOk returns a tuple with the FacebookMessengerPageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacebookMessengerPageId

`func (o *NotifyV1Service) SetFacebookMessengerPageId(v string)`

SetFacebookMessengerPageId sets FacebookMessengerPageId field to given value.

### HasFacebookMessengerPageId

`func (o *NotifyV1Service) HasFacebookMessengerPageId() bool`

HasFacebookMessengerPageId returns a boolean if a field has been set.

### SetFacebookMessengerPageIdNil

`func (o *NotifyV1Service) SetFacebookMessengerPageIdNil(b bool)`

 SetFacebookMessengerPageIdNil sets the value for FacebookMessengerPageId to be an explicit nil

### UnsetFacebookMessengerPageId
`func (o *NotifyV1Service) UnsetFacebookMessengerPageId()`

UnsetFacebookMessengerPageId ensures that no value is present for FacebookMessengerPageId, not even an explicit nil
### GetFcmCredentialSid

`func (o *NotifyV1Service) GetFcmCredentialSid() string`

GetFcmCredentialSid returns the FcmCredentialSid field if non-nil, zero value otherwise.

### GetFcmCredentialSidOk

`func (o *NotifyV1Service) GetFcmCredentialSidOk() (*string, bool)`

GetFcmCredentialSidOk returns a tuple with the FcmCredentialSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcmCredentialSid

`func (o *NotifyV1Service) SetFcmCredentialSid(v string)`

SetFcmCredentialSid sets FcmCredentialSid field to given value.

### HasFcmCredentialSid

`func (o *NotifyV1Service) HasFcmCredentialSid() bool`

HasFcmCredentialSid returns a boolean if a field has been set.

### SetFcmCredentialSidNil

`func (o *NotifyV1Service) SetFcmCredentialSidNil(b bool)`

 SetFcmCredentialSidNil sets the value for FcmCredentialSid to be an explicit nil

### UnsetFcmCredentialSid
`func (o *NotifyV1Service) UnsetFcmCredentialSid()`

UnsetFcmCredentialSid ensures that no value is present for FcmCredentialSid, not even an explicit nil
### GetFriendlyName

`func (o *NotifyV1Service) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *NotifyV1Service) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *NotifyV1Service) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *NotifyV1Service) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *NotifyV1Service) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *NotifyV1Service) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetGcmCredentialSid

`func (o *NotifyV1Service) GetGcmCredentialSid() string`

GetGcmCredentialSid returns the GcmCredentialSid field if non-nil, zero value otherwise.

### GetGcmCredentialSidOk

`func (o *NotifyV1Service) GetGcmCredentialSidOk() (*string, bool)`

GetGcmCredentialSidOk returns a tuple with the GcmCredentialSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcmCredentialSid

`func (o *NotifyV1Service) SetGcmCredentialSid(v string)`

SetGcmCredentialSid sets GcmCredentialSid field to given value.

### HasGcmCredentialSid

`func (o *NotifyV1Service) HasGcmCredentialSid() bool`

HasGcmCredentialSid returns a boolean if a field has been set.

### SetGcmCredentialSidNil

`func (o *NotifyV1Service) SetGcmCredentialSidNil(b bool)`

 SetGcmCredentialSidNil sets the value for GcmCredentialSid to be an explicit nil

### UnsetGcmCredentialSid
`func (o *NotifyV1Service) UnsetGcmCredentialSid()`

UnsetGcmCredentialSid ensures that no value is present for GcmCredentialSid, not even an explicit nil
### GetLinks

`func (o *NotifyV1Service) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *NotifyV1Service) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *NotifyV1Service) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *NotifyV1Service) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *NotifyV1Service) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *NotifyV1Service) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetLogEnabled

`func (o *NotifyV1Service) GetLogEnabled() bool`

GetLogEnabled returns the LogEnabled field if non-nil, zero value otherwise.

### GetLogEnabledOk

`func (o *NotifyV1Service) GetLogEnabledOk() (*bool, bool)`

GetLogEnabledOk returns a tuple with the LogEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogEnabled

`func (o *NotifyV1Service) SetLogEnabled(v bool)`

SetLogEnabled sets LogEnabled field to given value.

### HasLogEnabled

`func (o *NotifyV1Service) HasLogEnabled() bool`

HasLogEnabled returns a boolean if a field has been set.

### SetLogEnabledNil

`func (o *NotifyV1Service) SetLogEnabledNil(b bool)`

 SetLogEnabledNil sets the value for LogEnabled to be an explicit nil

### UnsetLogEnabled
`func (o *NotifyV1Service) UnsetLogEnabled()`

UnsetLogEnabled ensures that no value is present for LogEnabled, not even an explicit nil
### GetMessagingServiceSid

`func (o *NotifyV1Service) GetMessagingServiceSid() string`

GetMessagingServiceSid returns the MessagingServiceSid field if non-nil, zero value otherwise.

### GetMessagingServiceSidOk

`func (o *NotifyV1Service) GetMessagingServiceSidOk() (*string, bool)`

GetMessagingServiceSidOk returns a tuple with the MessagingServiceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagingServiceSid

`func (o *NotifyV1Service) SetMessagingServiceSid(v string)`

SetMessagingServiceSid sets MessagingServiceSid field to given value.

### HasMessagingServiceSid

`func (o *NotifyV1Service) HasMessagingServiceSid() bool`

HasMessagingServiceSid returns a boolean if a field has been set.

### SetMessagingServiceSidNil

`func (o *NotifyV1Service) SetMessagingServiceSidNil(b bool)`

 SetMessagingServiceSidNil sets the value for MessagingServiceSid to be an explicit nil

### UnsetMessagingServiceSid
`func (o *NotifyV1Service) UnsetMessagingServiceSid()`

UnsetMessagingServiceSid ensures that no value is present for MessagingServiceSid, not even an explicit nil
### GetSid

`func (o *NotifyV1Service) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *NotifyV1Service) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *NotifyV1Service) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *NotifyV1Service) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *NotifyV1Service) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *NotifyV1Service) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetUrl

`func (o *NotifyV1Service) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NotifyV1Service) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NotifyV1Service) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *NotifyV1Service) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *NotifyV1Service) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *NotifyV1Service) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


