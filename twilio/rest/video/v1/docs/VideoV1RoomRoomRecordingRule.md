# VideoV1RoomRoomRecordingRule

## Properties

Name | Type | Description
------------ | ------------- | -------------
**DateCreated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was last updated | [optional] 
**RoomSid** | Pointer to **NullableString** | The SID of the Room resource for the Recording Rules | [optional] 
**Rules** | Pointer to [**[]VideoV1RoomRoomRecordingRuleRules**](VideoV1RoomRoomRecordingRuleRules.md) | A collection of recording Rules that describe how to include or exclude matching tracks for recording | [optional] 

## Methods

### NewVideoV1RoomRoomRecordingRule

`func NewVideoV1RoomRoomRecordingRule() *VideoV1RoomRoomRecordingRule`

NewVideoV1RoomRoomRecordingRule instantiates a new VideoV1RoomRoomRecordingRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoV1RoomRoomRecordingRuleWithDefaults

`func NewVideoV1RoomRoomRecordingRuleWithDefaults() *VideoV1RoomRoomRecordingRule`

NewVideoV1RoomRoomRecordingRuleWithDefaults instantiates a new VideoV1RoomRoomRecordingRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateCreated

`func (o *VideoV1RoomRoomRecordingRule) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *VideoV1RoomRoomRecordingRule) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *VideoV1RoomRoomRecordingRule) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *VideoV1RoomRoomRecordingRule) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *VideoV1RoomRoomRecordingRule) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *VideoV1RoomRoomRecordingRule) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *VideoV1RoomRoomRecordingRule) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *VideoV1RoomRoomRecordingRule) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *VideoV1RoomRoomRecordingRule) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *VideoV1RoomRoomRecordingRule) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *VideoV1RoomRoomRecordingRule) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *VideoV1RoomRoomRecordingRule) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetRoomSid

`func (o *VideoV1RoomRoomRecordingRule) GetRoomSid() string`

GetRoomSid returns the RoomSid field if non-nil, zero value otherwise.

### GetRoomSidOk

`func (o *VideoV1RoomRoomRecordingRule) GetRoomSidOk() (*string, bool)`

GetRoomSidOk returns a tuple with the RoomSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoomSid

`func (o *VideoV1RoomRoomRecordingRule) SetRoomSid(v string)`

SetRoomSid sets RoomSid field to given value.

### HasRoomSid

`func (o *VideoV1RoomRoomRecordingRule) HasRoomSid() bool`

HasRoomSid returns a boolean if a field has been set.

### SetRoomSidNil

`func (o *VideoV1RoomRoomRecordingRule) SetRoomSidNil(b bool)`

 SetRoomSidNil sets the value for RoomSid to be an explicit nil

### UnsetRoomSid
`func (o *VideoV1RoomRoomRecordingRule) UnsetRoomSid()`

UnsetRoomSid ensures that no value is present for RoomSid, not even an explicit nil
### GetRules

`func (o *VideoV1RoomRoomRecordingRule) GetRules() []VideoV1RoomRoomRecordingRuleRules`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *VideoV1RoomRoomRecordingRule) GetRulesOk() (*[]VideoV1RoomRoomRecordingRuleRules, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *VideoV1RoomRoomRecordingRule) SetRules(v []VideoV1RoomRoomRecordingRuleRules)`

SetRules sets Rules field to given value.

### HasRules

`func (o *VideoV1RoomRoomRecordingRule) HasRules() bool`

HasRules returns a boolean if a field has been set.

### SetRulesNil

`func (o *VideoV1RoomRoomRecordingRule) SetRulesNil(b bool)`

 SetRulesNil sets the value for Rules to be an explicit nil

### UnsetRules
`func (o *VideoV1RoomRoomRecordingRule) UnsetRules()`

UnsetRules ensures that no value is present for Rules, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


