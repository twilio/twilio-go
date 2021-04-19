# SupersimV1SmsCommand

## Properties

Name | Type | Description
------------ | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**DateCreated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was last updated | [optional] 
**Direction** | Pointer to **NullableString** | The direction of the SMS Command | [optional] 
**Payload** | Pointer to **NullableString** | The message body of the SMS Command sent to or from the SIM | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**SimSid** | Pointer to **NullableString** | The SID of the SIM that this SMS Command was sent to or from | [optional] 
**Status** | Pointer to **NullableString** | The status of the SMS Command | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the SMS Command resource | [optional] 

## Methods

### NewSupersimV1SmsCommand

`func NewSupersimV1SmsCommand() *SupersimV1SmsCommand`

NewSupersimV1SmsCommand instantiates a new SupersimV1SmsCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupersimV1SmsCommandWithDefaults

`func NewSupersimV1SmsCommandWithDefaults() *SupersimV1SmsCommand`

NewSupersimV1SmsCommandWithDefaults instantiates a new SupersimV1SmsCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *SupersimV1SmsCommand) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *SupersimV1SmsCommand) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *SupersimV1SmsCommand) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *SupersimV1SmsCommand) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *SupersimV1SmsCommand) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *SupersimV1SmsCommand) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetDateCreated

`func (o *SupersimV1SmsCommand) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *SupersimV1SmsCommand) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *SupersimV1SmsCommand) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *SupersimV1SmsCommand) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *SupersimV1SmsCommand) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *SupersimV1SmsCommand) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *SupersimV1SmsCommand) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *SupersimV1SmsCommand) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *SupersimV1SmsCommand) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *SupersimV1SmsCommand) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *SupersimV1SmsCommand) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *SupersimV1SmsCommand) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetDirection

`func (o *SupersimV1SmsCommand) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *SupersimV1SmsCommand) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *SupersimV1SmsCommand) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *SupersimV1SmsCommand) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### SetDirectionNil

`func (o *SupersimV1SmsCommand) SetDirectionNil(b bool)`

 SetDirectionNil sets the value for Direction to be an explicit nil

### UnsetDirection
`func (o *SupersimV1SmsCommand) UnsetDirection()`

UnsetDirection ensures that no value is present for Direction, not even an explicit nil
### GetPayload

`func (o *SupersimV1SmsCommand) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *SupersimV1SmsCommand) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *SupersimV1SmsCommand) SetPayload(v string)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *SupersimV1SmsCommand) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### SetPayloadNil

`func (o *SupersimV1SmsCommand) SetPayloadNil(b bool)`

 SetPayloadNil sets the value for Payload to be an explicit nil

### UnsetPayload
`func (o *SupersimV1SmsCommand) UnsetPayload()`

UnsetPayload ensures that no value is present for Payload, not even an explicit nil
### GetSid

`func (o *SupersimV1SmsCommand) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *SupersimV1SmsCommand) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *SupersimV1SmsCommand) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *SupersimV1SmsCommand) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *SupersimV1SmsCommand) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *SupersimV1SmsCommand) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetSimSid

`func (o *SupersimV1SmsCommand) GetSimSid() string`

GetSimSid returns the SimSid field if non-nil, zero value otherwise.

### GetSimSidOk

`func (o *SupersimV1SmsCommand) GetSimSidOk() (*string, bool)`

GetSimSidOk returns a tuple with the SimSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimSid

`func (o *SupersimV1SmsCommand) SetSimSid(v string)`

SetSimSid sets SimSid field to given value.

### HasSimSid

`func (o *SupersimV1SmsCommand) HasSimSid() bool`

HasSimSid returns a boolean if a field has been set.

### SetSimSidNil

`func (o *SupersimV1SmsCommand) SetSimSidNil(b bool)`

 SetSimSidNil sets the value for SimSid to be an explicit nil

### UnsetSimSid
`func (o *SupersimV1SmsCommand) UnsetSimSid()`

UnsetSimSid ensures that no value is present for SimSid, not even an explicit nil
### GetStatus

`func (o *SupersimV1SmsCommand) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SupersimV1SmsCommand) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SupersimV1SmsCommand) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SupersimV1SmsCommand) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *SupersimV1SmsCommand) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *SupersimV1SmsCommand) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetUrl

`func (o *SupersimV1SmsCommand) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SupersimV1SmsCommand) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SupersimV1SmsCommand) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SupersimV1SmsCommand) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *SupersimV1SmsCommand) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *SupersimV1SmsCommand) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


