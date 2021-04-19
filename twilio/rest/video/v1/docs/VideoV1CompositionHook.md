# VideoV1CompositionHook

## Properties

Name | Type | Description
------------ | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**AudioSources** | Pointer to **[]string** | The array of track names to include in the compositions created by the composition hook | [optional] 
**AudioSourcesExcluded** | Pointer to **[]string** | The array of track names to exclude from the compositions created by the composition hook | [optional] 
**DateCreated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was last updated | [optional] 
**Enabled** | Pointer to **NullableBool** | Whether the CompositionHook is active | [optional] 
**Format** | Pointer to **NullableString** | The container format of the media files used by the compositions created by the composition hook | [optional] 
**FriendlyName** | Pointer to **NullableString** | The string that you assigned to describe the resource | [optional] 
**Resolution** | Pointer to **NullableString** | The dimensions of the video image in pixels expressed as columns (width) and rows (height) | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**StatusCallback** | Pointer to **NullableString** | The URL to send status information to your application | [optional] 
**StatusCallbackMethod** | Pointer to **NullableString** | The HTTP method we should use to call status_callback | [optional] 
**Trim** | Pointer to **NullableBool** | Whether intervals with no media are clipped | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the resource | [optional] 
**VideoLayout** | Pointer to **map[string]interface{}** | A JSON object that describes the video layout of the Composition | [optional] 

## Methods

### NewVideoV1CompositionHook

`func NewVideoV1CompositionHook() *VideoV1CompositionHook`

NewVideoV1CompositionHook instantiates a new VideoV1CompositionHook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoV1CompositionHookWithDefaults

`func NewVideoV1CompositionHookWithDefaults() *VideoV1CompositionHook`

NewVideoV1CompositionHookWithDefaults instantiates a new VideoV1CompositionHook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *VideoV1CompositionHook) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *VideoV1CompositionHook) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *VideoV1CompositionHook) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *VideoV1CompositionHook) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *VideoV1CompositionHook) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *VideoV1CompositionHook) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetAudioSources

`func (o *VideoV1CompositionHook) GetAudioSources() []string`

GetAudioSources returns the AudioSources field if non-nil, zero value otherwise.

### GetAudioSourcesOk

`func (o *VideoV1CompositionHook) GetAudioSourcesOk() (*[]string, bool)`

GetAudioSourcesOk returns a tuple with the AudioSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioSources

`func (o *VideoV1CompositionHook) SetAudioSources(v []string)`

SetAudioSources sets AudioSources field to given value.

### HasAudioSources

`func (o *VideoV1CompositionHook) HasAudioSources() bool`

HasAudioSources returns a boolean if a field has been set.

### SetAudioSourcesNil

`func (o *VideoV1CompositionHook) SetAudioSourcesNil(b bool)`

 SetAudioSourcesNil sets the value for AudioSources to be an explicit nil

### UnsetAudioSources
`func (o *VideoV1CompositionHook) UnsetAudioSources()`

UnsetAudioSources ensures that no value is present for AudioSources, not even an explicit nil
### GetAudioSourcesExcluded

`func (o *VideoV1CompositionHook) GetAudioSourcesExcluded() []string`

GetAudioSourcesExcluded returns the AudioSourcesExcluded field if non-nil, zero value otherwise.

### GetAudioSourcesExcludedOk

`func (o *VideoV1CompositionHook) GetAudioSourcesExcludedOk() (*[]string, bool)`

GetAudioSourcesExcludedOk returns a tuple with the AudioSourcesExcluded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioSourcesExcluded

`func (o *VideoV1CompositionHook) SetAudioSourcesExcluded(v []string)`

SetAudioSourcesExcluded sets AudioSourcesExcluded field to given value.

### HasAudioSourcesExcluded

`func (o *VideoV1CompositionHook) HasAudioSourcesExcluded() bool`

HasAudioSourcesExcluded returns a boolean if a field has been set.

### SetAudioSourcesExcludedNil

`func (o *VideoV1CompositionHook) SetAudioSourcesExcludedNil(b bool)`

 SetAudioSourcesExcludedNil sets the value for AudioSourcesExcluded to be an explicit nil

### UnsetAudioSourcesExcluded
`func (o *VideoV1CompositionHook) UnsetAudioSourcesExcluded()`

UnsetAudioSourcesExcluded ensures that no value is present for AudioSourcesExcluded, not even an explicit nil
### GetDateCreated

`func (o *VideoV1CompositionHook) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *VideoV1CompositionHook) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *VideoV1CompositionHook) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *VideoV1CompositionHook) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *VideoV1CompositionHook) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *VideoV1CompositionHook) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *VideoV1CompositionHook) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *VideoV1CompositionHook) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *VideoV1CompositionHook) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *VideoV1CompositionHook) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *VideoV1CompositionHook) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *VideoV1CompositionHook) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetEnabled

`func (o *VideoV1CompositionHook) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *VideoV1CompositionHook) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *VideoV1CompositionHook) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *VideoV1CompositionHook) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *VideoV1CompositionHook) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *VideoV1CompositionHook) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
### GetFormat

`func (o *VideoV1CompositionHook) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *VideoV1CompositionHook) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *VideoV1CompositionHook) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *VideoV1CompositionHook) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### SetFormatNil

`func (o *VideoV1CompositionHook) SetFormatNil(b bool)`

 SetFormatNil sets the value for Format to be an explicit nil

### UnsetFormat
`func (o *VideoV1CompositionHook) UnsetFormat()`

UnsetFormat ensures that no value is present for Format, not even an explicit nil
### GetFriendlyName

`func (o *VideoV1CompositionHook) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *VideoV1CompositionHook) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *VideoV1CompositionHook) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *VideoV1CompositionHook) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *VideoV1CompositionHook) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *VideoV1CompositionHook) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetResolution

`func (o *VideoV1CompositionHook) GetResolution() string`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *VideoV1CompositionHook) GetResolutionOk() (*string, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *VideoV1CompositionHook) SetResolution(v string)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *VideoV1CompositionHook) HasResolution() bool`

HasResolution returns a boolean if a field has been set.

### SetResolutionNil

`func (o *VideoV1CompositionHook) SetResolutionNil(b bool)`

 SetResolutionNil sets the value for Resolution to be an explicit nil

### UnsetResolution
`func (o *VideoV1CompositionHook) UnsetResolution()`

UnsetResolution ensures that no value is present for Resolution, not even an explicit nil
### GetSid

`func (o *VideoV1CompositionHook) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *VideoV1CompositionHook) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *VideoV1CompositionHook) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *VideoV1CompositionHook) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *VideoV1CompositionHook) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *VideoV1CompositionHook) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetStatusCallback

`func (o *VideoV1CompositionHook) GetStatusCallback() string`

GetStatusCallback returns the StatusCallback field if non-nil, zero value otherwise.

### GetStatusCallbackOk

`func (o *VideoV1CompositionHook) GetStatusCallbackOk() (*string, bool)`

GetStatusCallbackOk returns a tuple with the StatusCallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCallback

`func (o *VideoV1CompositionHook) SetStatusCallback(v string)`

SetStatusCallback sets StatusCallback field to given value.

### HasStatusCallback

`func (o *VideoV1CompositionHook) HasStatusCallback() bool`

HasStatusCallback returns a boolean if a field has been set.

### SetStatusCallbackNil

`func (o *VideoV1CompositionHook) SetStatusCallbackNil(b bool)`

 SetStatusCallbackNil sets the value for StatusCallback to be an explicit nil

### UnsetStatusCallback
`func (o *VideoV1CompositionHook) UnsetStatusCallback()`

UnsetStatusCallback ensures that no value is present for StatusCallback, not even an explicit nil
### GetStatusCallbackMethod

`func (o *VideoV1CompositionHook) GetStatusCallbackMethod() string`

GetStatusCallbackMethod returns the StatusCallbackMethod field if non-nil, zero value otherwise.

### GetStatusCallbackMethodOk

`func (o *VideoV1CompositionHook) GetStatusCallbackMethodOk() (*string, bool)`

GetStatusCallbackMethodOk returns a tuple with the StatusCallbackMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCallbackMethod

`func (o *VideoV1CompositionHook) SetStatusCallbackMethod(v string)`

SetStatusCallbackMethod sets StatusCallbackMethod field to given value.

### HasStatusCallbackMethod

`func (o *VideoV1CompositionHook) HasStatusCallbackMethod() bool`

HasStatusCallbackMethod returns a boolean if a field has been set.

### SetStatusCallbackMethodNil

`func (o *VideoV1CompositionHook) SetStatusCallbackMethodNil(b bool)`

 SetStatusCallbackMethodNil sets the value for StatusCallbackMethod to be an explicit nil

### UnsetStatusCallbackMethod
`func (o *VideoV1CompositionHook) UnsetStatusCallbackMethod()`

UnsetStatusCallbackMethod ensures that no value is present for StatusCallbackMethod, not even an explicit nil
### GetTrim

`func (o *VideoV1CompositionHook) GetTrim() bool`

GetTrim returns the Trim field if non-nil, zero value otherwise.

### GetTrimOk

`func (o *VideoV1CompositionHook) GetTrimOk() (*bool, bool)`

GetTrimOk returns a tuple with the Trim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrim

`func (o *VideoV1CompositionHook) SetTrim(v bool)`

SetTrim sets Trim field to given value.

### HasTrim

`func (o *VideoV1CompositionHook) HasTrim() bool`

HasTrim returns a boolean if a field has been set.

### SetTrimNil

`func (o *VideoV1CompositionHook) SetTrimNil(b bool)`

 SetTrimNil sets the value for Trim to be an explicit nil

### UnsetTrim
`func (o *VideoV1CompositionHook) UnsetTrim()`

UnsetTrim ensures that no value is present for Trim, not even an explicit nil
### GetUrl

`func (o *VideoV1CompositionHook) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *VideoV1CompositionHook) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *VideoV1CompositionHook) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *VideoV1CompositionHook) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *VideoV1CompositionHook) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *VideoV1CompositionHook) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetVideoLayout

`func (o *VideoV1CompositionHook) GetVideoLayout() map[string]interface{}`

GetVideoLayout returns the VideoLayout field if non-nil, zero value otherwise.

### GetVideoLayoutOk

`func (o *VideoV1CompositionHook) GetVideoLayoutOk() (*map[string]interface{}, bool)`

GetVideoLayoutOk returns a tuple with the VideoLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoLayout

`func (o *VideoV1CompositionHook) SetVideoLayout(v map[string]interface{})`

SetVideoLayout sets VideoLayout field to given value.

### HasVideoLayout

`func (o *VideoV1CompositionHook) HasVideoLayout() bool`

HasVideoLayout returns a boolean if a field has been set.

### SetVideoLayoutNil

`func (o *VideoV1CompositionHook) SetVideoLayoutNil(b bool)`

 SetVideoLayoutNil sets the value for VideoLayout to be an explicit nil

### UnsetVideoLayout
`func (o *VideoV1CompositionHook) UnsetVideoLayout()`

UnsetVideoLayout ensures that no value is present for VideoLayout, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


