# SupersimV1Sim

## Properties

Name | Type | Description
------------ | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that the Super SIM belongs to | [optional] 
**DateCreated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was last updated | [optional] 
**FleetSid** | Pointer to **NullableString** | The unique ID of the Fleet configured for this SIM | [optional] 
**Iccid** | Pointer to **NullableString** | The ICCID associated with the SIM | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**Status** | Pointer to **NullableString** | The status of the Super SIM | [optional] 
**UniqueName** | Pointer to **NullableString** | An application-defined string that uniquely identifies the resource | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the Sim Resource | [optional] 

## Methods

### NewSupersimV1Sim

`func NewSupersimV1Sim() *SupersimV1Sim`

NewSupersimV1Sim instantiates a new SupersimV1Sim object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupersimV1SimWithDefaults

`func NewSupersimV1SimWithDefaults() *SupersimV1Sim`

NewSupersimV1SimWithDefaults instantiates a new SupersimV1Sim object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *SupersimV1Sim) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *SupersimV1Sim) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *SupersimV1Sim) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *SupersimV1Sim) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *SupersimV1Sim) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *SupersimV1Sim) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetDateCreated

`func (o *SupersimV1Sim) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *SupersimV1Sim) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *SupersimV1Sim) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *SupersimV1Sim) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *SupersimV1Sim) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *SupersimV1Sim) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *SupersimV1Sim) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *SupersimV1Sim) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *SupersimV1Sim) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *SupersimV1Sim) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *SupersimV1Sim) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *SupersimV1Sim) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetFleetSid

`func (o *SupersimV1Sim) GetFleetSid() string`

GetFleetSid returns the FleetSid field if non-nil, zero value otherwise.

### GetFleetSidOk

`func (o *SupersimV1Sim) GetFleetSidOk() (*string, bool)`

GetFleetSidOk returns a tuple with the FleetSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetSid

`func (o *SupersimV1Sim) SetFleetSid(v string)`

SetFleetSid sets FleetSid field to given value.

### HasFleetSid

`func (o *SupersimV1Sim) HasFleetSid() bool`

HasFleetSid returns a boolean if a field has been set.

### SetFleetSidNil

`func (o *SupersimV1Sim) SetFleetSidNil(b bool)`

 SetFleetSidNil sets the value for FleetSid to be an explicit nil

### UnsetFleetSid
`func (o *SupersimV1Sim) UnsetFleetSid()`

UnsetFleetSid ensures that no value is present for FleetSid, not even an explicit nil
### GetIccid

`func (o *SupersimV1Sim) GetIccid() string`

GetIccid returns the Iccid field if non-nil, zero value otherwise.

### GetIccidOk

`func (o *SupersimV1Sim) GetIccidOk() (*string, bool)`

GetIccidOk returns a tuple with the Iccid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIccid

`func (o *SupersimV1Sim) SetIccid(v string)`

SetIccid sets Iccid field to given value.

### HasIccid

`func (o *SupersimV1Sim) HasIccid() bool`

HasIccid returns a boolean if a field has been set.

### SetIccidNil

`func (o *SupersimV1Sim) SetIccidNil(b bool)`

 SetIccidNil sets the value for Iccid to be an explicit nil

### UnsetIccid
`func (o *SupersimV1Sim) UnsetIccid()`

UnsetIccid ensures that no value is present for Iccid, not even an explicit nil
### GetSid

`func (o *SupersimV1Sim) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *SupersimV1Sim) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *SupersimV1Sim) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *SupersimV1Sim) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *SupersimV1Sim) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *SupersimV1Sim) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetStatus

`func (o *SupersimV1Sim) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SupersimV1Sim) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SupersimV1Sim) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SupersimV1Sim) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *SupersimV1Sim) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *SupersimV1Sim) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetUniqueName

`func (o *SupersimV1Sim) GetUniqueName() string`

GetUniqueName returns the UniqueName field if non-nil, zero value otherwise.

### GetUniqueNameOk

`func (o *SupersimV1Sim) GetUniqueNameOk() (*string, bool)`

GetUniqueNameOk returns a tuple with the UniqueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueName

`func (o *SupersimV1Sim) SetUniqueName(v string)`

SetUniqueName sets UniqueName field to given value.

### HasUniqueName

`func (o *SupersimV1Sim) HasUniqueName() bool`

HasUniqueName returns a boolean if a field has been set.

### SetUniqueNameNil

`func (o *SupersimV1Sim) SetUniqueNameNil(b bool)`

 SetUniqueNameNil sets the value for UniqueName to be an explicit nil

### UnsetUniqueName
`func (o *SupersimV1Sim) UnsetUniqueName()`

UnsetUniqueName ensures that no value is present for UniqueName, not even an explicit nil
### GetUrl

`func (o *SupersimV1Sim) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SupersimV1Sim) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SupersimV1Sim) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SupersimV1Sim) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *SupersimV1Sim) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *SupersimV1Sim) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


