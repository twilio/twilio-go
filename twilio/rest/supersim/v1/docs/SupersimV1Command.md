# SupersimV1Command

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**Command** | Pointer to **NullableString** | The message body of the command sent to or from the SIM | [optional] 
**DateCreated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was last updated | [optional] 
**Direction** | Pointer to **NullableString** | The direction of the Command | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**SimSid** | Pointer to **NullableString** | The SID of the SIM that this Command was sent to or from | [optional] 
**Status** | Pointer to **NullableString** | The status of the Command | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the Command resource | [optional] 

## Methods

### NewSupersimV1Command

`func NewSupersimV1Command() *SupersimV1Command`

NewSupersimV1Command instantiates a new SupersimV1Command object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupersimV1CommandWithDefaults

`func NewSupersimV1CommandWithDefaults() *SupersimV1Command`

NewSupersimV1CommandWithDefaults instantiates a new SupersimV1Command object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *SupersimV1Command) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *SupersimV1Command) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *SupersimV1Command) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *SupersimV1Command) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *SupersimV1Command) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *SupersimV1Command) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetCommand

`func (o *SupersimV1Command) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *SupersimV1Command) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *SupersimV1Command) SetCommand(v string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *SupersimV1Command) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### SetCommandNil

`func (o *SupersimV1Command) SetCommandNil(b bool)`

 SetCommandNil sets the value for Command to be an explicit nil

### UnsetCommand
`func (o *SupersimV1Command) UnsetCommand()`

UnsetCommand ensures that no value is present for Command, not even an explicit nil
### GetDateCreated

`func (o *SupersimV1Command) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *SupersimV1Command) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *SupersimV1Command) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *SupersimV1Command) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *SupersimV1Command) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *SupersimV1Command) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *SupersimV1Command) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *SupersimV1Command) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *SupersimV1Command) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *SupersimV1Command) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *SupersimV1Command) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *SupersimV1Command) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetDirection

`func (o *SupersimV1Command) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *SupersimV1Command) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *SupersimV1Command) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *SupersimV1Command) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### SetDirectionNil

`func (o *SupersimV1Command) SetDirectionNil(b bool)`

 SetDirectionNil sets the value for Direction to be an explicit nil

### UnsetDirection
`func (o *SupersimV1Command) UnsetDirection()`

UnsetDirection ensures that no value is present for Direction, not even an explicit nil
### GetSid

`func (o *SupersimV1Command) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *SupersimV1Command) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *SupersimV1Command) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *SupersimV1Command) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *SupersimV1Command) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *SupersimV1Command) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetSimSid

`func (o *SupersimV1Command) GetSimSid() string`

GetSimSid returns the SimSid field if non-nil, zero value otherwise.

### GetSimSidOk

`func (o *SupersimV1Command) GetSimSidOk() (*string, bool)`

GetSimSidOk returns a tuple with the SimSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimSid

`func (o *SupersimV1Command) SetSimSid(v string)`

SetSimSid sets SimSid field to given value.

### HasSimSid

`func (o *SupersimV1Command) HasSimSid() bool`

HasSimSid returns a boolean if a field has been set.

### SetSimSidNil

`func (o *SupersimV1Command) SetSimSidNil(b bool)`

 SetSimSidNil sets the value for SimSid to be an explicit nil

### UnsetSimSid
`func (o *SupersimV1Command) UnsetSimSid()`

UnsetSimSid ensures that no value is present for SimSid, not even an explicit nil
### GetStatus

`func (o *SupersimV1Command) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SupersimV1Command) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SupersimV1Command) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SupersimV1Command) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *SupersimV1Command) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *SupersimV1Command) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetUrl

`func (o *SupersimV1Command) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SupersimV1Command) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SupersimV1Command) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SupersimV1Command) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *SupersimV1Command) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *SupersimV1Command) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


