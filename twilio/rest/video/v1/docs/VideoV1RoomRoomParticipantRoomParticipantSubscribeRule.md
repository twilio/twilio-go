# VideoV1RoomRoomParticipantRoomParticipantSubscribeRule

## Properties

Name | Type | Description
------------ | ------------- | -------------
**DateCreated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was last updated | [optional] 
**ParticipantSid** | Pointer to **NullableString** | The SID of the Participant resource for the Subscribe Rules | [optional] 
**RoomSid** | Pointer to **NullableString** | The SID of the Room resource for the Subscribe Rules | [optional] 
**Rules** | Pointer to **[]map[string]interface{}** | A collection of Subscribe Rules that describe how to include or exclude matching tracks | [optional] 

## Methods

### NewVideoV1RoomRoomParticipantRoomParticipantSubscribeRule

`func NewVideoV1RoomRoomParticipantRoomParticipantSubscribeRule() *VideoV1RoomRoomParticipantRoomParticipantSubscribeRule`

NewVideoV1RoomRoomParticipantRoomParticipantSubscribeRule instantiates a new VideoV1RoomRoomParticipantRoomParticipantSubscribeRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoV1RoomRoomParticipantRoomParticipantSubscribeRuleWithDefaults

`func NewVideoV1RoomRoomParticipantRoomParticipantSubscribeRuleWithDefaults() *VideoV1RoomRoomParticipantRoomParticipantSubscribeRule`

NewVideoV1RoomRoomParticipantRoomParticipantSubscribeRuleWithDefaults instantiates a new VideoV1RoomRoomParticipantRoomParticipantSubscribeRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateCreated

`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribeRule) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribeRule) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribeRule) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribeRule) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribeRule) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribeRule) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribeRule) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribeRule) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribeRule) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribeRule) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribeRule) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribeRule) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetParticipantSid

`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribeRule) GetParticipantSid() string`

GetParticipantSid returns the ParticipantSid field if non-nil, zero value otherwise.

### GetParticipantSidOk

`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribeRule) GetParticipantSidOk() (*string, bool)`

GetParticipantSidOk returns a tuple with the ParticipantSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipantSid

`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribeRule) SetParticipantSid(v string)`

SetParticipantSid sets ParticipantSid field to given value.

### HasParticipantSid

`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribeRule) HasParticipantSid() bool`

HasParticipantSid returns a boolean if a field has been set.

### SetParticipantSidNil

`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribeRule) SetParticipantSidNil(b bool)`

 SetParticipantSidNil sets the value for ParticipantSid to be an explicit nil

### UnsetParticipantSid
`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribeRule) UnsetParticipantSid()`

UnsetParticipantSid ensures that no value is present for ParticipantSid, not even an explicit nil
### GetRoomSid

`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribeRule) GetRoomSid() string`

GetRoomSid returns the RoomSid field if non-nil, zero value otherwise.

### GetRoomSidOk

`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribeRule) GetRoomSidOk() (*string, bool)`

GetRoomSidOk returns a tuple with the RoomSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoomSid

`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribeRule) SetRoomSid(v string)`

SetRoomSid sets RoomSid field to given value.

### HasRoomSid

`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribeRule) HasRoomSid() bool`

HasRoomSid returns a boolean if a field has been set.

### SetRoomSidNil

`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribeRule) SetRoomSidNil(b bool)`

 SetRoomSidNil sets the value for RoomSid to be an explicit nil

### UnsetRoomSid
`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribeRule) UnsetRoomSid()`

UnsetRoomSid ensures that no value is present for RoomSid, not even an explicit nil
### GetRules

`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribeRule) GetRules() []map[string]interface{}`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribeRule) GetRulesOk() (*[]map[string]interface{}, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribeRule) SetRules(v []map[string]interface{})`

SetRules sets Rules field to given value.

### HasRules

`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribeRule) HasRules() bool`

HasRules returns a boolean if a field has been set.

### SetRulesNil

`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribeRule) SetRulesNil(b bool)`

 SetRulesNil sets the value for Rules to be an explicit nil

### UnsetRules
`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribeRule) UnsetRules()`

UnsetRules ensures that no value is present for Rules, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


