# AutopilotV1AssistantModelBuild

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**AssistantSid** | Pointer to **NullableString** | The SID of the Assistant that is the parent of the resource | [optional] 
**BuildDuration** | Pointer to **NullableInt32** | The time in seconds it took to build the model | [optional] 
**DateCreated** | Pointer to **NullableTime** | The RFC 2822 date and time in GMT when the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The RFC 2822 date and time in GMT when the resource was last updated | [optional] 
**ErrorCode** | Pointer to **NullableInt32** | More information about why the model build failed, if &#x60;status&#x60; is &#x60;failed&#x60; | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**Status** | Pointer to **NullableString** | The status of the model build process | [optional] 
**UniqueName** | Pointer to **NullableString** | An application-defined string that uniquely identifies the resource | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the ModelBuild resource | [optional] 

## Methods

### NewAutopilotV1AssistantModelBuild

`func NewAutopilotV1AssistantModelBuild() *AutopilotV1AssistantModelBuild`

NewAutopilotV1AssistantModelBuild instantiates a new AutopilotV1AssistantModelBuild object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutopilotV1AssistantModelBuildWithDefaults

`func NewAutopilotV1AssistantModelBuildWithDefaults() *AutopilotV1AssistantModelBuild`

NewAutopilotV1AssistantModelBuildWithDefaults instantiates a new AutopilotV1AssistantModelBuild object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *AutopilotV1AssistantModelBuild) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *AutopilotV1AssistantModelBuild) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *AutopilotV1AssistantModelBuild) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *AutopilotV1AssistantModelBuild) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *AutopilotV1AssistantModelBuild) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *AutopilotV1AssistantModelBuild) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetAssistantSid

`func (o *AutopilotV1AssistantModelBuild) GetAssistantSid() string`

GetAssistantSid returns the AssistantSid field if non-nil, zero value otherwise.

### GetAssistantSidOk

`func (o *AutopilotV1AssistantModelBuild) GetAssistantSidOk() (*string, bool)`

GetAssistantSidOk returns a tuple with the AssistantSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssistantSid

`func (o *AutopilotV1AssistantModelBuild) SetAssistantSid(v string)`

SetAssistantSid sets AssistantSid field to given value.

### HasAssistantSid

`func (o *AutopilotV1AssistantModelBuild) HasAssistantSid() bool`

HasAssistantSid returns a boolean if a field has been set.

### SetAssistantSidNil

`func (o *AutopilotV1AssistantModelBuild) SetAssistantSidNil(b bool)`

 SetAssistantSidNil sets the value for AssistantSid to be an explicit nil

### UnsetAssistantSid
`func (o *AutopilotV1AssistantModelBuild) UnsetAssistantSid()`

UnsetAssistantSid ensures that no value is present for AssistantSid, not even an explicit nil
### GetBuildDuration

`func (o *AutopilotV1AssistantModelBuild) GetBuildDuration() int32`

GetBuildDuration returns the BuildDuration field if non-nil, zero value otherwise.

### GetBuildDurationOk

`func (o *AutopilotV1AssistantModelBuild) GetBuildDurationOk() (*int32, bool)`

GetBuildDurationOk returns a tuple with the BuildDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildDuration

`func (o *AutopilotV1AssistantModelBuild) SetBuildDuration(v int32)`

SetBuildDuration sets BuildDuration field to given value.

### HasBuildDuration

`func (o *AutopilotV1AssistantModelBuild) HasBuildDuration() bool`

HasBuildDuration returns a boolean if a field has been set.

### SetBuildDurationNil

`func (o *AutopilotV1AssistantModelBuild) SetBuildDurationNil(b bool)`

 SetBuildDurationNil sets the value for BuildDuration to be an explicit nil

### UnsetBuildDuration
`func (o *AutopilotV1AssistantModelBuild) UnsetBuildDuration()`

UnsetBuildDuration ensures that no value is present for BuildDuration, not even an explicit nil
### GetDateCreated

`func (o *AutopilotV1AssistantModelBuild) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *AutopilotV1AssistantModelBuild) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *AutopilotV1AssistantModelBuild) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *AutopilotV1AssistantModelBuild) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *AutopilotV1AssistantModelBuild) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *AutopilotV1AssistantModelBuild) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *AutopilotV1AssistantModelBuild) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *AutopilotV1AssistantModelBuild) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *AutopilotV1AssistantModelBuild) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *AutopilotV1AssistantModelBuild) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *AutopilotV1AssistantModelBuild) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *AutopilotV1AssistantModelBuild) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetErrorCode

`func (o *AutopilotV1AssistantModelBuild) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *AutopilotV1AssistantModelBuild) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *AutopilotV1AssistantModelBuild) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *AutopilotV1AssistantModelBuild) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### SetErrorCodeNil

`func (o *AutopilotV1AssistantModelBuild) SetErrorCodeNil(b bool)`

 SetErrorCodeNil sets the value for ErrorCode to be an explicit nil

### UnsetErrorCode
`func (o *AutopilotV1AssistantModelBuild) UnsetErrorCode()`

UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
### GetSid

`func (o *AutopilotV1AssistantModelBuild) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *AutopilotV1AssistantModelBuild) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *AutopilotV1AssistantModelBuild) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *AutopilotV1AssistantModelBuild) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *AutopilotV1AssistantModelBuild) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *AutopilotV1AssistantModelBuild) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetStatus

`func (o *AutopilotV1AssistantModelBuild) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AutopilotV1AssistantModelBuild) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AutopilotV1AssistantModelBuild) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AutopilotV1AssistantModelBuild) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *AutopilotV1AssistantModelBuild) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *AutopilotV1AssistantModelBuild) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetUniqueName

`func (o *AutopilotV1AssistantModelBuild) GetUniqueName() string`

GetUniqueName returns the UniqueName field if non-nil, zero value otherwise.

### GetUniqueNameOk

`func (o *AutopilotV1AssistantModelBuild) GetUniqueNameOk() (*string, bool)`

GetUniqueNameOk returns a tuple with the UniqueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueName

`func (o *AutopilotV1AssistantModelBuild) SetUniqueName(v string)`

SetUniqueName sets UniqueName field to given value.

### HasUniqueName

`func (o *AutopilotV1AssistantModelBuild) HasUniqueName() bool`

HasUniqueName returns a boolean if a field has been set.

### SetUniqueNameNil

`func (o *AutopilotV1AssistantModelBuild) SetUniqueNameNil(b bool)`

 SetUniqueNameNil sets the value for UniqueName to be an explicit nil

### UnsetUniqueName
`func (o *AutopilotV1AssistantModelBuild) UnsetUniqueName()`

UnsetUniqueName ensures that no value is present for UniqueName, not even an explicit nil
### GetUrl

`func (o *AutopilotV1AssistantModelBuild) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AutopilotV1AssistantModelBuild) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AutopilotV1AssistantModelBuild) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AutopilotV1AssistantModelBuild) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *AutopilotV1AssistantModelBuild) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *AutopilotV1AssistantModelBuild) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


