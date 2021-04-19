# AutopilotV1RestoreAssistant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**CallbackEvents** | Pointer to **NullableString** | Reserved | [optional] 
**CallbackUrl** | Pointer to **NullableString** | Reserved | [optional] 
**DateCreated** | Pointer to **NullableTime** | The RFC 2822 date and time in GMT when the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The RFC 2822 date and time in GMT when the resource was last updated | [optional] 
**DevelopmentStage** | Pointer to **NullableString** | A string describing the state of the assistant. | [optional] 
**FriendlyName** | Pointer to **NullableString** | The string that you assigned to describe the resource | [optional] 
**LatestModelBuildSid** | Pointer to **NullableString** | Reserved | [optional] 
**LogQueries** | Pointer to **NullableBool** | Whether queries should be logged and kept after training | [optional] 
**NeedsModelBuild** | Pointer to **NullableBool** | Whether model needs to be rebuilt | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**UniqueName** | Pointer to **NullableString** | An application-defined string that uniquely identifies the resource | [optional] 

## Methods

### NewAutopilotV1RestoreAssistant

`func NewAutopilotV1RestoreAssistant() *AutopilotV1RestoreAssistant`

NewAutopilotV1RestoreAssistant instantiates a new AutopilotV1RestoreAssistant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutopilotV1RestoreAssistantWithDefaults

`func NewAutopilotV1RestoreAssistantWithDefaults() *AutopilotV1RestoreAssistant`

NewAutopilotV1RestoreAssistantWithDefaults instantiates a new AutopilotV1RestoreAssistant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *AutopilotV1RestoreAssistant) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *AutopilotV1RestoreAssistant) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *AutopilotV1RestoreAssistant) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *AutopilotV1RestoreAssistant) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *AutopilotV1RestoreAssistant) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *AutopilotV1RestoreAssistant) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetCallbackEvents

`func (o *AutopilotV1RestoreAssistant) GetCallbackEvents() string`

GetCallbackEvents returns the CallbackEvents field if non-nil, zero value otherwise.

### GetCallbackEventsOk

`func (o *AutopilotV1RestoreAssistant) GetCallbackEventsOk() (*string, bool)`

GetCallbackEventsOk returns a tuple with the CallbackEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackEvents

`func (o *AutopilotV1RestoreAssistant) SetCallbackEvents(v string)`

SetCallbackEvents sets CallbackEvents field to given value.

### HasCallbackEvents

`func (o *AutopilotV1RestoreAssistant) HasCallbackEvents() bool`

HasCallbackEvents returns a boolean if a field has been set.

### SetCallbackEventsNil

`func (o *AutopilotV1RestoreAssistant) SetCallbackEventsNil(b bool)`

 SetCallbackEventsNil sets the value for CallbackEvents to be an explicit nil

### UnsetCallbackEvents
`func (o *AutopilotV1RestoreAssistant) UnsetCallbackEvents()`

UnsetCallbackEvents ensures that no value is present for CallbackEvents, not even an explicit nil
### GetCallbackUrl

`func (o *AutopilotV1RestoreAssistant) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *AutopilotV1RestoreAssistant) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *AutopilotV1RestoreAssistant) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.

### HasCallbackUrl

`func (o *AutopilotV1RestoreAssistant) HasCallbackUrl() bool`

HasCallbackUrl returns a boolean if a field has been set.

### SetCallbackUrlNil

`func (o *AutopilotV1RestoreAssistant) SetCallbackUrlNil(b bool)`

 SetCallbackUrlNil sets the value for CallbackUrl to be an explicit nil

### UnsetCallbackUrl
`func (o *AutopilotV1RestoreAssistant) UnsetCallbackUrl()`

UnsetCallbackUrl ensures that no value is present for CallbackUrl, not even an explicit nil
### GetDateCreated

`func (o *AutopilotV1RestoreAssistant) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *AutopilotV1RestoreAssistant) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *AutopilotV1RestoreAssistant) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *AutopilotV1RestoreAssistant) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *AutopilotV1RestoreAssistant) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *AutopilotV1RestoreAssistant) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *AutopilotV1RestoreAssistant) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *AutopilotV1RestoreAssistant) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *AutopilotV1RestoreAssistant) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *AutopilotV1RestoreAssistant) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *AutopilotV1RestoreAssistant) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *AutopilotV1RestoreAssistant) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetDevelopmentStage

`func (o *AutopilotV1RestoreAssistant) GetDevelopmentStage() string`

GetDevelopmentStage returns the DevelopmentStage field if non-nil, zero value otherwise.

### GetDevelopmentStageOk

`func (o *AutopilotV1RestoreAssistant) GetDevelopmentStageOk() (*string, bool)`

GetDevelopmentStageOk returns a tuple with the DevelopmentStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevelopmentStage

`func (o *AutopilotV1RestoreAssistant) SetDevelopmentStage(v string)`

SetDevelopmentStage sets DevelopmentStage field to given value.

### HasDevelopmentStage

`func (o *AutopilotV1RestoreAssistant) HasDevelopmentStage() bool`

HasDevelopmentStage returns a boolean if a field has been set.

### SetDevelopmentStageNil

`func (o *AutopilotV1RestoreAssistant) SetDevelopmentStageNil(b bool)`

 SetDevelopmentStageNil sets the value for DevelopmentStage to be an explicit nil

### UnsetDevelopmentStage
`func (o *AutopilotV1RestoreAssistant) UnsetDevelopmentStage()`

UnsetDevelopmentStage ensures that no value is present for DevelopmentStage, not even an explicit nil
### GetFriendlyName

`func (o *AutopilotV1RestoreAssistant) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *AutopilotV1RestoreAssistant) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *AutopilotV1RestoreAssistant) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *AutopilotV1RestoreAssistant) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *AutopilotV1RestoreAssistant) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *AutopilotV1RestoreAssistant) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetLatestModelBuildSid

`func (o *AutopilotV1RestoreAssistant) GetLatestModelBuildSid() string`

GetLatestModelBuildSid returns the LatestModelBuildSid field if non-nil, zero value otherwise.

### GetLatestModelBuildSidOk

`func (o *AutopilotV1RestoreAssistant) GetLatestModelBuildSidOk() (*string, bool)`

GetLatestModelBuildSidOk returns a tuple with the LatestModelBuildSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestModelBuildSid

`func (o *AutopilotV1RestoreAssistant) SetLatestModelBuildSid(v string)`

SetLatestModelBuildSid sets LatestModelBuildSid field to given value.

### HasLatestModelBuildSid

`func (o *AutopilotV1RestoreAssistant) HasLatestModelBuildSid() bool`

HasLatestModelBuildSid returns a boolean if a field has been set.

### SetLatestModelBuildSidNil

`func (o *AutopilotV1RestoreAssistant) SetLatestModelBuildSidNil(b bool)`

 SetLatestModelBuildSidNil sets the value for LatestModelBuildSid to be an explicit nil

### UnsetLatestModelBuildSid
`func (o *AutopilotV1RestoreAssistant) UnsetLatestModelBuildSid()`

UnsetLatestModelBuildSid ensures that no value is present for LatestModelBuildSid, not even an explicit nil
### GetLogQueries

`func (o *AutopilotV1RestoreAssistant) GetLogQueries() bool`

GetLogQueries returns the LogQueries field if non-nil, zero value otherwise.

### GetLogQueriesOk

`func (o *AutopilotV1RestoreAssistant) GetLogQueriesOk() (*bool, bool)`

GetLogQueriesOk returns a tuple with the LogQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogQueries

`func (o *AutopilotV1RestoreAssistant) SetLogQueries(v bool)`

SetLogQueries sets LogQueries field to given value.

### HasLogQueries

`func (o *AutopilotV1RestoreAssistant) HasLogQueries() bool`

HasLogQueries returns a boolean if a field has been set.

### SetLogQueriesNil

`func (o *AutopilotV1RestoreAssistant) SetLogQueriesNil(b bool)`

 SetLogQueriesNil sets the value for LogQueries to be an explicit nil

### UnsetLogQueries
`func (o *AutopilotV1RestoreAssistant) UnsetLogQueries()`

UnsetLogQueries ensures that no value is present for LogQueries, not even an explicit nil
### GetNeedsModelBuild

`func (o *AutopilotV1RestoreAssistant) GetNeedsModelBuild() bool`

GetNeedsModelBuild returns the NeedsModelBuild field if non-nil, zero value otherwise.

### GetNeedsModelBuildOk

`func (o *AutopilotV1RestoreAssistant) GetNeedsModelBuildOk() (*bool, bool)`

GetNeedsModelBuildOk returns a tuple with the NeedsModelBuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedsModelBuild

`func (o *AutopilotV1RestoreAssistant) SetNeedsModelBuild(v bool)`

SetNeedsModelBuild sets NeedsModelBuild field to given value.

### HasNeedsModelBuild

`func (o *AutopilotV1RestoreAssistant) HasNeedsModelBuild() bool`

HasNeedsModelBuild returns a boolean if a field has been set.

### SetNeedsModelBuildNil

`func (o *AutopilotV1RestoreAssistant) SetNeedsModelBuildNil(b bool)`

 SetNeedsModelBuildNil sets the value for NeedsModelBuild to be an explicit nil

### UnsetNeedsModelBuild
`func (o *AutopilotV1RestoreAssistant) UnsetNeedsModelBuild()`

UnsetNeedsModelBuild ensures that no value is present for NeedsModelBuild, not even an explicit nil
### GetSid

`func (o *AutopilotV1RestoreAssistant) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *AutopilotV1RestoreAssistant) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *AutopilotV1RestoreAssistant) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *AutopilotV1RestoreAssistant) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *AutopilotV1RestoreAssistant) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *AutopilotV1RestoreAssistant) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetUniqueName

`func (o *AutopilotV1RestoreAssistant) GetUniqueName() string`

GetUniqueName returns the UniqueName field if non-nil, zero value otherwise.

### GetUniqueNameOk

`func (o *AutopilotV1RestoreAssistant) GetUniqueNameOk() (*string, bool)`

GetUniqueNameOk returns a tuple with the UniqueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueName

`func (o *AutopilotV1RestoreAssistant) SetUniqueName(v string)`

SetUniqueName sets UniqueName field to given value.

### HasUniqueName

`func (o *AutopilotV1RestoreAssistant) HasUniqueName() bool`

HasUniqueName returns a boolean if a field has been set.

### SetUniqueNameNil

`func (o *AutopilotV1RestoreAssistant) SetUniqueNameNil(b bool)`

 SetUniqueNameNil sets the value for UniqueName to be an explicit nil

### UnsetUniqueName
`func (o *AutopilotV1RestoreAssistant) UnsetUniqueName()`

UnsetUniqueName ensures that no value is present for UniqueName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


