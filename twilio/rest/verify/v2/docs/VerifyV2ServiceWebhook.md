# VerifyV2ServiceWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**DateCreated** | Pointer to **NullableTime** | The RFC 2822 date and time in GMT when the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The RFC 2822 date and time in GMT when the resource was last updated | [optional] 
**EventTypes** | Pointer to **[]string** | The array of events that this Webhook is subscribed to. | [optional] 
**FriendlyName** | Pointer to **NullableString** | The string that you assigned to describe the webhook | [optional] 
**ServiceSid** | Pointer to **NullableString** | Service Sid. | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**Status** | Pointer to **NullableString** | The webhook status | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the Webhook resource | [optional] 
**WebhookMethod** | Pointer to **NullableString** | The method used when calling the webhook&#39;s URL. | [optional] 
**WebhookUrl** | Pointer to **NullableString** | The URL associated with this Webhook. | [optional] 

## Methods

### NewVerifyV2ServiceWebhook

`func NewVerifyV2ServiceWebhook() *VerifyV2ServiceWebhook`

NewVerifyV2ServiceWebhook instantiates a new VerifyV2ServiceWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifyV2ServiceWebhookWithDefaults

`func NewVerifyV2ServiceWebhookWithDefaults() *VerifyV2ServiceWebhook`

NewVerifyV2ServiceWebhookWithDefaults instantiates a new VerifyV2ServiceWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *VerifyV2ServiceWebhook) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *VerifyV2ServiceWebhook) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *VerifyV2ServiceWebhook) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *VerifyV2ServiceWebhook) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *VerifyV2ServiceWebhook) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *VerifyV2ServiceWebhook) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetDateCreated

`func (o *VerifyV2ServiceWebhook) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *VerifyV2ServiceWebhook) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *VerifyV2ServiceWebhook) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *VerifyV2ServiceWebhook) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *VerifyV2ServiceWebhook) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *VerifyV2ServiceWebhook) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *VerifyV2ServiceWebhook) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *VerifyV2ServiceWebhook) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *VerifyV2ServiceWebhook) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *VerifyV2ServiceWebhook) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *VerifyV2ServiceWebhook) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *VerifyV2ServiceWebhook) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetEventTypes

`func (o *VerifyV2ServiceWebhook) GetEventTypes() []string`

GetEventTypes returns the EventTypes field if non-nil, zero value otherwise.

### GetEventTypesOk

`func (o *VerifyV2ServiceWebhook) GetEventTypesOk() (*[]string, bool)`

GetEventTypesOk returns a tuple with the EventTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypes

`func (o *VerifyV2ServiceWebhook) SetEventTypes(v []string)`

SetEventTypes sets EventTypes field to given value.

### HasEventTypes

`func (o *VerifyV2ServiceWebhook) HasEventTypes() bool`

HasEventTypes returns a boolean if a field has been set.

### SetEventTypesNil

`func (o *VerifyV2ServiceWebhook) SetEventTypesNil(b bool)`

 SetEventTypesNil sets the value for EventTypes to be an explicit nil

### UnsetEventTypes
`func (o *VerifyV2ServiceWebhook) UnsetEventTypes()`

UnsetEventTypes ensures that no value is present for EventTypes, not even an explicit nil
### GetFriendlyName

`func (o *VerifyV2ServiceWebhook) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *VerifyV2ServiceWebhook) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *VerifyV2ServiceWebhook) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *VerifyV2ServiceWebhook) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *VerifyV2ServiceWebhook) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *VerifyV2ServiceWebhook) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetServiceSid

`func (o *VerifyV2ServiceWebhook) GetServiceSid() string`

GetServiceSid returns the ServiceSid field if non-nil, zero value otherwise.

### GetServiceSidOk

`func (o *VerifyV2ServiceWebhook) GetServiceSidOk() (*string, bool)`

GetServiceSidOk returns a tuple with the ServiceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSid

`func (o *VerifyV2ServiceWebhook) SetServiceSid(v string)`

SetServiceSid sets ServiceSid field to given value.

### HasServiceSid

`func (o *VerifyV2ServiceWebhook) HasServiceSid() bool`

HasServiceSid returns a boolean if a field has been set.

### SetServiceSidNil

`func (o *VerifyV2ServiceWebhook) SetServiceSidNil(b bool)`

 SetServiceSidNil sets the value for ServiceSid to be an explicit nil

### UnsetServiceSid
`func (o *VerifyV2ServiceWebhook) UnsetServiceSid()`

UnsetServiceSid ensures that no value is present for ServiceSid, not even an explicit nil
### GetSid

`func (o *VerifyV2ServiceWebhook) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *VerifyV2ServiceWebhook) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *VerifyV2ServiceWebhook) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *VerifyV2ServiceWebhook) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *VerifyV2ServiceWebhook) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *VerifyV2ServiceWebhook) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetStatus

`func (o *VerifyV2ServiceWebhook) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VerifyV2ServiceWebhook) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VerifyV2ServiceWebhook) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VerifyV2ServiceWebhook) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *VerifyV2ServiceWebhook) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *VerifyV2ServiceWebhook) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetUrl

`func (o *VerifyV2ServiceWebhook) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *VerifyV2ServiceWebhook) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *VerifyV2ServiceWebhook) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *VerifyV2ServiceWebhook) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *VerifyV2ServiceWebhook) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *VerifyV2ServiceWebhook) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetWebhookMethod

`func (o *VerifyV2ServiceWebhook) GetWebhookMethod() string`

GetWebhookMethod returns the WebhookMethod field if non-nil, zero value otherwise.

### GetWebhookMethodOk

`func (o *VerifyV2ServiceWebhook) GetWebhookMethodOk() (*string, bool)`

GetWebhookMethodOk returns a tuple with the WebhookMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookMethod

`func (o *VerifyV2ServiceWebhook) SetWebhookMethod(v string)`

SetWebhookMethod sets WebhookMethod field to given value.

### HasWebhookMethod

`func (o *VerifyV2ServiceWebhook) HasWebhookMethod() bool`

HasWebhookMethod returns a boolean if a field has been set.

### SetWebhookMethodNil

`func (o *VerifyV2ServiceWebhook) SetWebhookMethodNil(b bool)`

 SetWebhookMethodNil sets the value for WebhookMethod to be an explicit nil

### UnsetWebhookMethod
`func (o *VerifyV2ServiceWebhook) UnsetWebhookMethod()`

UnsetWebhookMethod ensures that no value is present for WebhookMethod, not even an explicit nil
### GetWebhookUrl

`func (o *VerifyV2ServiceWebhook) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *VerifyV2ServiceWebhook) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *VerifyV2ServiceWebhook) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *VerifyV2ServiceWebhook) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### SetWebhookUrlNil

`func (o *VerifyV2ServiceWebhook) SetWebhookUrlNil(b bool)`

 SetWebhookUrlNil sets the value for WebhookUrl to be an explicit nil

### UnsetWebhookUrl
`func (o *VerifyV2ServiceWebhook) UnsetWebhookUrl()`

UnsetWebhookUrl ensures that no value is present for WebhookUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


