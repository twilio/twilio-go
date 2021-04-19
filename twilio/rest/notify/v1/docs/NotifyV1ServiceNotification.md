# NotifyV1ServiceNotification

## Properties

Name | Type | Description
------------ | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**Action** | Pointer to **NullableString** | The actions to display for the notification | [optional] 
**Alexa** | Pointer to **map[string]interface{}** | Deprecated | [optional] 
**Apn** | Pointer to **map[string]interface{}** | The APNS-specific payload that overrides corresponding attributes in a generic payload for APNS Bindings | [optional] 
**Body** | Pointer to **NullableString** | The notification body text | [optional] 
**Data** | Pointer to **map[string]interface{}** | The custom key-value pairs of the notification&#39;s payload | [optional] 
**DateCreated** | Pointer to **NullableTime** | The RFC 2822 date and time in GMT when the resource was created | [optional] 
**FacebookMessenger** | Pointer to **map[string]interface{}** | Deprecated | [optional] 
**Fcm** | Pointer to **map[string]interface{}** | The FCM-specific payload that overrides corresponding attributes in generic payload for FCM Bindings | [optional] 
**Gcm** | Pointer to **map[string]interface{}** | The GCM-specific payload that overrides corresponding attributes in generic payload for GCM Bindings | [optional] 
**Identities** | Pointer to **[]string** | The list of identity values of the Users to notify | [optional] 
**Priority** | Pointer to **NullableString** | The priority of the notification | [optional] 
**Segments** | Pointer to **[]string** | The list of Segments to notify | [optional] 
**ServiceSid** | Pointer to **NullableString** | The SID of the Service that the resource is associated with | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**Sms** | Pointer to **map[string]interface{}** | The SMS-specific payload that overrides corresponding attributes in generic payload for SMS Bindings | [optional] 
**Sound** | Pointer to **NullableString** | The name of the sound to be played for the notification | [optional] 
**Tags** | Pointer to **[]string** | The tags that select the Bindings to notify | [optional] 
**Title** | Pointer to **NullableString** | The notification title | [optional] 
**Ttl** | Pointer to **NullableInt32** | How long, in seconds, the notification is valid | [optional] 

## Methods

### NewNotifyV1ServiceNotification

`func NewNotifyV1ServiceNotification() *NotifyV1ServiceNotification`

NewNotifyV1ServiceNotification instantiates a new NotifyV1ServiceNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotifyV1ServiceNotificationWithDefaults

`func NewNotifyV1ServiceNotificationWithDefaults() *NotifyV1ServiceNotification`

NewNotifyV1ServiceNotificationWithDefaults instantiates a new NotifyV1ServiceNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *NotifyV1ServiceNotification) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *NotifyV1ServiceNotification) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *NotifyV1ServiceNotification) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *NotifyV1ServiceNotification) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *NotifyV1ServiceNotification) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *NotifyV1ServiceNotification) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetAction

`func (o *NotifyV1ServiceNotification) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *NotifyV1ServiceNotification) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *NotifyV1ServiceNotification) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *NotifyV1ServiceNotification) HasAction() bool`

HasAction returns a boolean if a field has been set.

### SetActionNil

`func (o *NotifyV1ServiceNotification) SetActionNil(b bool)`

 SetActionNil sets the value for Action to be an explicit nil

### UnsetAction
`func (o *NotifyV1ServiceNotification) UnsetAction()`

UnsetAction ensures that no value is present for Action, not even an explicit nil
### GetAlexa

`func (o *NotifyV1ServiceNotification) GetAlexa() map[string]interface{}`

GetAlexa returns the Alexa field if non-nil, zero value otherwise.

### GetAlexaOk

`func (o *NotifyV1ServiceNotification) GetAlexaOk() (*map[string]interface{}, bool)`

GetAlexaOk returns a tuple with the Alexa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlexa

`func (o *NotifyV1ServiceNotification) SetAlexa(v map[string]interface{})`

SetAlexa sets Alexa field to given value.

### HasAlexa

`func (o *NotifyV1ServiceNotification) HasAlexa() bool`

HasAlexa returns a boolean if a field has been set.

### SetAlexaNil

`func (o *NotifyV1ServiceNotification) SetAlexaNil(b bool)`

 SetAlexaNil sets the value for Alexa to be an explicit nil

### UnsetAlexa
`func (o *NotifyV1ServiceNotification) UnsetAlexa()`

UnsetAlexa ensures that no value is present for Alexa, not even an explicit nil
### GetApn

`func (o *NotifyV1ServiceNotification) GetApn() map[string]interface{}`

GetApn returns the Apn field if non-nil, zero value otherwise.

### GetApnOk

`func (o *NotifyV1ServiceNotification) GetApnOk() (*map[string]interface{}, bool)`

GetApnOk returns a tuple with the Apn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApn

`func (o *NotifyV1ServiceNotification) SetApn(v map[string]interface{})`

SetApn sets Apn field to given value.

### HasApn

`func (o *NotifyV1ServiceNotification) HasApn() bool`

HasApn returns a boolean if a field has been set.

### SetApnNil

`func (o *NotifyV1ServiceNotification) SetApnNil(b bool)`

 SetApnNil sets the value for Apn to be an explicit nil

### UnsetApn
`func (o *NotifyV1ServiceNotification) UnsetApn()`

UnsetApn ensures that no value is present for Apn, not even an explicit nil
### GetBody

`func (o *NotifyV1ServiceNotification) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *NotifyV1ServiceNotification) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *NotifyV1ServiceNotification) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *NotifyV1ServiceNotification) HasBody() bool`

HasBody returns a boolean if a field has been set.

### SetBodyNil

`func (o *NotifyV1ServiceNotification) SetBodyNil(b bool)`

 SetBodyNil sets the value for Body to be an explicit nil

### UnsetBody
`func (o *NotifyV1ServiceNotification) UnsetBody()`

UnsetBody ensures that no value is present for Body, not even an explicit nil
### GetData

`func (o *NotifyV1ServiceNotification) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *NotifyV1ServiceNotification) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *NotifyV1ServiceNotification) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *NotifyV1ServiceNotification) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *NotifyV1ServiceNotification) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *NotifyV1ServiceNotification) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetDateCreated

`func (o *NotifyV1ServiceNotification) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *NotifyV1ServiceNotification) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *NotifyV1ServiceNotification) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *NotifyV1ServiceNotification) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *NotifyV1ServiceNotification) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *NotifyV1ServiceNotification) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetFacebookMessenger

`func (o *NotifyV1ServiceNotification) GetFacebookMessenger() map[string]interface{}`

GetFacebookMessenger returns the FacebookMessenger field if non-nil, zero value otherwise.

### GetFacebookMessengerOk

`func (o *NotifyV1ServiceNotification) GetFacebookMessengerOk() (*map[string]interface{}, bool)`

GetFacebookMessengerOk returns a tuple with the FacebookMessenger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacebookMessenger

`func (o *NotifyV1ServiceNotification) SetFacebookMessenger(v map[string]interface{})`

SetFacebookMessenger sets FacebookMessenger field to given value.

### HasFacebookMessenger

`func (o *NotifyV1ServiceNotification) HasFacebookMessenger() bool`

HasFacebookMessenger returns a boolean if a field has been set.

### SetFacebookMessengerNil

`func (o *NotifyV1ServiceNotification) SetFacebookMessengerNil(b bool)`

 SetFacebookMessengerNil sets the value for FacebookMessenger to be an explicit nil

### UnsetFacebookMessenger
`func (o *NotifyV1ServiceNotification) UnsetFacebookMessenger()`

UnsetFacebookMessenger ensures that no value is present for FacebookMessenger, not even an explicit nil
### GetFcm

`func (o *NotifyV1ServiceNotification) GetFcm() map[string]interface{}`

GetFcm returns the Fcm field if non-nil, zero value otherwise.

### GetFcmOk

`func (o *NotifyV1ServiceNotification) GetFcmOk() (*map[string]interface{}, bool)`

GetFcmOk returns a tuple with the Fcm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcm

`func (o *NotifyV1ServiceNotification) SetFcm(v map[string]interface{})`

SetFcm sets Fcm field to given value.

### HasFcm

`func (o *NotifyV1ServiceNotification) HasFcm() bool`

HasFcm returns a boolean if a field has been set.

### SetFcmNil

`func (o *NotifyV1ServiceNotification) SetFcmNil(b bool)`

 SetFcmNil sets the value for Fcm to be an explicit nil

### UnsetFcm
`func (o *NotifyV1ServiceNotification) UnsetFcm()`

UnsetFcm ensures that no value is present for Fcm, not even an explicit nil
### GetGcm

`func (o *NotifyV1ServiceNotification) GetGcm() map[string]interface{}`

GetGcm returns the Gcm field if non-nil, zero value otherwise.

### GetGcmOk

`func (o *NotifyV1ServiceNotification) GetGcmOk() (*map[string]interface{}, bool)`

GetGcmOk returns a tuple with the Gcm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcm

`func (o *NotifyV1ServiceNotification) SetGcm(v map[string]interface{})`

SetGcm sets Gcm field to given value.

### HasGcm

`func (o *NotifyV1ServiceNotification) HasGcm() bool`

HasGcm returns a boolean if a field has been set.

### SetGcmNil

`func (o *NotifyV1ServiceNotification) SetGcmNil(b bool)`

 SetGcmNil sets the value for Gcm to be an explicit nil

### UnsetGcm
`func (o *NotifyV1ServiceNotification) UnsetGcm()`

UnsetGcm ensures that no value is present for Gcm, not even an explicit nil
### GetIdentities

`func (o *NotifyV1ServiceNotification) GetIdentities() []string`

GetIdentities returns the Identities field if non-nil, zero value otherwise.

### GetIdentitiesOk

`func (o *NotifyV1ServiceNotification) GetIdentitiesOk() (*[]string, bool)`

GetIdentitiesOk returns a tuple with the Identities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentities

`func (o *NotifyV1ServiceNotification) SetIdentities(v []string)`

SetIdentities sets Identities field to given value.

### HasIdentities

`func (o *NotifyV1ServiceNotification) HasIdentities() bool`

HasIdentities returns a boolean if a field has been set.

### SetIdentitiesNil

`func (o *NotifyV1ServiceNotification) SetIdentitiesNil(b bool)`

 SetIdentitiesNil sets the value for Identities to be an explicit nil

### UnsetIdentities
`func (o *NotifyV1ServiceNotification) UnsetIdentities()`

UnsetIdentities ensures that no value is present for Identities, not even an explicit nil
### GetPriority

`func (o *NotifyV1ServiceNotification) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *NotifyV1ServiceNotification) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *NotifyV1ServiceNotification) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *NotifyV1ServiceNotification) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *NotifyV1ServiceNotification) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *NotifyV1ServiceNotification) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil
### GetSegments

`func (o *NotifyV1ServiceNotification) GetSegments() []string`

GetSegments returns the Segments field if non-nil, zero value otherwise.

### GetSegmentsOk

`func (o *NotifyV1ServiceNotification) GetSegmentsOk() (*[]string, bool)`

GetSegmentsOk returns a tuple with the Segments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegments

`func (o *NotifyV1ServiceNotification) SetSegments(v []string)`

SetSegments sets Segments field to given value.

### HasSegments

`func (o *NotifyV1ServiceNotification) HasSegments() bool`

HasSegments returns a boolean if a field has been set.

### SetSegmentsNil

`func (o *NotifyV1ServiceNotification) SetSegmentsNil(b bool)`

 SetSegmentsNil sets the value for Segments to be an explicit nil

### UnsetSegments
`func (o *NotifyV1ServiceNotification) UnsetSegments()`

UnsetSegments ensures that no value is present for Segments, not even an explicit nil
### GetServiceSid

`func (o *NotifyV1ServiceNotification) GetServiceSid() string`

GetServiceSid returns the ServiceSid field if non-nil, zero value otherwise.

### GetServiceSidOk

`func (o *NotifyV1ServiceNotification) GetServiceSidOk() (*string, bool)`

GetServiceSidOk returns a tuple with the ServiceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSid

`func (o *NotifyV1ServiceNotification) SetServiceSid(v string)`

SetServiceSid sets ServiceSid field to given value.

### HasServiceSid

`func (o *NotifyV1ServiceNotification) HasServiceSid() bool`

HasServiceSid returns a boolean if a field has been set.

### SetServiceSidNil

`func (o *NotifyV1ServiceNotification) SetServiceSidNil(b bool)`

 SetServiceSidNil sets the value for ServiceSid to be an explicit nil

### UnsetServiceSid
`func (o *NotifyV1ServiceNotification) UnsetServiceSid()`

UnsetServiceSid ensures that no value is present for ServiceSid, not even an explicit nil
### GetSid

`func (o *NotifyV1ServiceNotification) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *NotifyV1ServiceNotification) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *NotifyV1ServiceNotification) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *NotifyV1ServiceNotification) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *NotifyV1ServiceNotification) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *NotifyV1ServiceNotification) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetSms

`func (o *NotifyV1ServiceNotification) GetSms() map[string]interface{}`

GetSms returns the Sms field if non-nil, zero value otherwise.

### GetSmsOk

`func (o *NotifyV1ServiceNotification) GetSmsOk() (*map[string]interface{}, bool)`

GetSmsOk returns a tuple with the Sms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSms

`func (o *NotifyV1ServiceNotification) SetSms(v map[string]interface{})`

SetSms sets Sms field to given value.

### HasSms

`func (o *NotifyV1ServiceNotification) HasSms() bool`

HasSms returns a boolean if a field has been set.

### SetSmsNil

`func (o *NotifyV1ServiceNotification) SetSmsNil(b bool)`

 SetSmsNil sets the value for Sms to be an explicit nil

### UnsetSms
`func (o *NotifyV1ServiceNotification) UnsetSms()`

UnsetSms ensures that no value is present for Sms, not even an explicit nil
### GetSound

`func (o *NotifyV1ServiceNotification) GetSound() string`

GetSound returns the Sound field if non-nil, zero value otherwise.

### GetSoundOk

`func (o *NotifyV1ServiceNotification) GetSoundOk() (*string, bool)`

GetSoundOk returns a tuple with the Sound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSound

`func (o *NotifyV1ServiceNotification) SetSound(v string)`

SetSound sets Sound field to given value.

### HasSound

`func (o *NotifyV1ServiceNotification) HasSound() bool`

HasSound returns a boolean if a field has been set.

### SetSoundNil

`func (o *NotifyV1ServiceNotification) SetSoundNil(b bool)`

 SetSoundNil sets the value for Sound to be an explicit nil

### UnsetSound
`func (o *NotifyV1ServiceNotification) UnsetSound()`

UnsetSound ensures that no value is present for Sound, not even an explicit nil
### GetTags

`func (o *NotifyV1ServiceNotification) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *NotifyV1ServiceNotification) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *NotifyV1ServiceNotification) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *NotifyV1ServiceNotification) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *NotifyV1ServiceNotification) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *NotifyV1ServiceNotification) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetTitle

`func (o *NotifyV1ServiceNotification) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *NotifyV1ServiceNotification) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *NotifyV1ServiceNotification) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *NotifyV1ServiceNotification) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *NotifyV1ServiceNotification) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *NotifyV1ServiceNotification) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetTtl

`func (o *NotifyV1ServiceNotification) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *NotifyV1ServiceNotification) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *NotifyV1ServiceNotification) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *NotifyV1ServiceNotification) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### SetTtlNil

`func (o *NotifyV1ServiceNotification) SetTtlNil(b bool)`

 SetTtlNil sets the value for Ttl to be an explicit nil

### UnsetTtl
`func (o *NotifyV1ServiceNotification) UnsetTtl()`

UnsetTtl ensures that no value is present for Ttl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


