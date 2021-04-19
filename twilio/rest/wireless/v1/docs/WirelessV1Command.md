# WirelessV1Command

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**Command** | Pointer to **NullableString** | The message being sent to or from the SIM | [optional] 
**CommandMode** | Pointer to **NullableString** | The mode used to send the SMS message | [optional] 
**DateCreated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was last updated format | [optional] 
**DeliveryReceiptRequested** | Pointer to **NullableBool** | Whether to request a delivery receipt | [optional] 
**Direction** | Pointer to **NullableString** | The direction of the Command | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**SimSid** | Pointer to **NullableString** | The SID of the Sim resource that the Command was sent to or from | [optional] 
**Status** | Pointer to **NullableString** | The status of the Command | [optional] 
**Transport** | Pointer to **NullableString** | The type of transport used | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the resource | [optional] 

## Methods

### NewWirelessV1Command

`func NewWirelessV1Command() *WirelessV1Command`

NewWirelessV1Command instantiates a new WirelessV1Command object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWirelessV1CommandWithDefaults

`func NewWirelessV1CommandWithDefaults() *WirelessV1Command`

NewWirelessV1CommandWithDefaults instantiates a new WirelessV1Command object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *WirelessV1Command) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *WirelessV1Command) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *WirelessV1Command) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *WirelessV1Command) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *WirelessV1Command) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *WirelessV1Command) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetCommand

`func (o *WirelessV1Command) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *WirelessV1Command) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *WirelessV1Command) SetCommand(v string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *WirelessV1Command) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### SetCommandNil

`func (o *WirelessV1Command) SetCommandNil(b bool)`

 SetCommandNil sets the value for Command to be an explicit nil

### UnsetCommand
`func (o *WirelessV1Command) UnsetCommand()`

UnsetCommand ensures that no value is present for Command, not even an explicit nil
### GetCommandMode

`func (o *WirelessV1Command) GetCommandMode() string`

GetCommandMode returns the CommandMode field if non-nil, zero value otherwise.

### GetCommandModeOk

`func (o *WirelessV1Command) GetCommandModeOk() (*string, bool)`

GetCommandModeOk returns a tuple with the CommandMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandMode

`func (o *WirelessV1Command) SetCommandMode(v string)`

SetCommandMode sets CommandMode field to given value.

### HasCommandMode

`func (o *WirelessV1Command) HasCommandMode() bool`

HasCommandMode returns a boolean if a field has been set.

### SetCommandModeNil

`func (o *WirelessV1Command) SetCommandModeNil(b bool)`

 SetCommandModeNil sets the value for CommandMode to be an explicit nil

### UnsetCommandMode
`func (o *WirelessV1Command) UnsetCommandMode()`

UnsetCommandMode ensures that no value is present for CommandMode, not even an explicit nil
### GetDateCreated

`func (o *WirelessV1Command) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *WirelessV1Command) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *WirelessV1Command) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *WirelessV1Command) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *WirelessV1Command) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *WirelessV1Command) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *WirelessV1Command) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *WirelessV1Command) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *WirelessV1Command) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *WirelessV1Command) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *WirelessV1Command) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *WirelessV1Command) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetDeliveryReceiptRequested

`func (o *WirelessV1Command) GetDeliveryReceiptRequested() bool`

GetDeliveryReceiptRequested returns the DeliveryReceiptRequested field if non-nil, zero value otherwise.

### GetDeliveryReceiptRequestedOk

`func (o *WirelessV1Command) GetDeliveryReceiptRequestedOk() (*bool, bool)`

GetDeliveryReceiptRequestedOk returns a tuple with the DeliveryReceiptRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryReceiptRequested

`func (o *WirelessV1Command) SetDeliveryReceiptRequested(v bool)`

SetDeliveryReceiptRequested sets DeliveryReceiptRequested field to given value.

### HasDeliveryReceiptRequested

`func (o *WirelessV1Command) HasDeliveryReceiptRequested() bool`

HasDeliveryReceiptRequested returns a boolean if a field has been set.

### SetDeliveryReceiptRequestedNil

`func (o *WirelessV1Command) SetDeliveryReceiptRequestedNil(b bool)`

 SetDeliveryReceiptRequestedNil sets the value for DeliveryReceiptRequested to be an explicit nil

### UnsetDeliveryReceiptRequested
`func (o *WirelessV1Command) UnsetDeliveryReceiptRequested()`

UnsetDeliveryReceiptRequested ensures that no value is present for DeliveryReceiptRequested, not even an explicit nil
### GetDirection

`func (o *WirelessV1Command) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *WirelessV1Command) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *WirelessV1Command) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *WirelessV1Command) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### SetDirectionNil

`func (o *WirelessV1Command) SetDirectionNil(b bool)`

 SetDirectionNil sets the value for Direction to be an explicit nil

### UnsetDirection
`func (o *WirelessV1Command) UnsetDirection()`

UnsetDirection ensures that no value is present for Direction, not even an explicit nil
### GetSid

`func (o *WirelessV1Command) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *WirelessV1Command) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *WirelessV1Command) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *WirelessV1Command) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *WirelessV1Command) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *WirelessV1Command) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetSimSid

`func (o *WirelessV1Command) GetSimSid() string`

GetSimSid returns the SimSid field if non-nil, zero value otherwise.

### GetSimSidOk

`func (o *WirelessV1Command) GetSimSidOk() (*string, bool)`

GetSimSidOk returns a tuple with the SimSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimSid

`func (o *WirelessV1Command) SetSimSid(v string)`

SetSimSid sets SimSid field to given value.

### HasSimSid

`func (o *WirelessV1Command) HasSimSid() bool`

HasSimSid returns a boolean if a field has been set.

### SetSimSidNil

`func (o *WirelessV1Command) SetSimSidNil(b bool)`

 SetSimSidNil sets the value for SimSid to be an explicit nil

### UnsetSimSid
`func (o *WirelessV1Command) UnsetSimSid()`

UnsetSimSid ensures that no value is present for SimSid, not even an explicit nil
### GetStatus

`func (o *WirelessV1Command) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WirelessV1Command) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WirelessV1Command) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WirelessV1Command) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *WirelessV1Command) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *WirelessV1Command) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetTransport

`func (o *WirelessV1Command) GetTransport() string`

GetTransport returns the Transport field if non-nil, zero value otherwise.

### GetTransportOk

`func (o *WirelessV1Command) GetTransportOk() (*string, bool)`

GetTransportOk returns a tuple with the Transport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransport

`func (o *WirelessV1Command) SetTransport(v string)`

SetTransport sets Transport field to given value.

### HasTransport

`func (o *WirelessV1Command) HasTransport() bool`

HasTransport returns a boolean if a field has been set.

### SetTransportNil

`func (o *WirelessV1Command) SetTransportNil(b bool)`

 SetTransportNil sets the value for Transport to be an explicit nil

### UnsetTransport
`func (o *WirelessV1Command) UnsetTransport()`

UnsetTransport ensures that no value is present for Transport, not even an explicit nil
### GetUrl

`func (o *WirelessV1Command) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WirelessV1Command) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WirelessV1Command) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *WirelessV1Command) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *WirelessV1Command) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *WirelessV1Command) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


