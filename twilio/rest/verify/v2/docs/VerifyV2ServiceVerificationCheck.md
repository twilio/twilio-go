# VerifyV2ServiceVerificationCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**Amount** | Pointer to **NullableString** | The amount of the associated PSD2 compliant transaction. | [optional] 
**Channel** | Pointer to **NullableString** | The verification method to use | [optional] 
**DateCreated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the Verification Check resource was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the Verification Check resource was last updated | [optional] 
**Payee** | Pointer to **NullableString** | The payee of the associated PSD2 compliant transaction | [optional] 
**ServiceSid** | Pointer to **NullableString** | The SID of the Service that the resource is associated with | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**Status** | Pointer to **NullableString** | The status of the verification resource | [optional] 
**To** | Pointer to **NullableString** | The phone number or email being verified | [optional] 
**Valid** | Pointer to **NullableBool** | Whether the verification was successful | [optional] 

## Methods

### NewVerifyV2ServiceVerificationCheck

`func NewVerifyV2ServiceVerificationCheck() *VerifyV2ServiceVerificationCheck`

NewVerifyV2ServiceVerificationCheck instantiates a new VerifyV2ServiceVerificationCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifyV2ServiceVerificationCheckWithDefaults

`func NewVerifyV2ServiceVerificationCheckWithDefaults() *VerifyV2ServiceVerificationCheck`

NewVerifyV2ServiceVerificationCheckWithDefaults instantiates a new VerifyV2ServiceVerificationCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *VerifyV2ServiceVerificationCheck) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *VerifyV2ServiceVerificationCheck) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *VerifyV2ServiceVerificationCheck) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *VerifyV2ServiceVerificationCheck) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *VerifyV2ServiceVerificationCheck) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *VerifyV2ServiceVerificationCheck) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetAmount

`func (o *VerifyV2ServiceVerificationCheck) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *VerifyV2ServiceVerificationCheck) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *VerifyV2ServiceVerificationCheck) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *VerifyV2ServiceVerificationCheck) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### SetAmountNil

`func (o *VerifyV2ServiceVerificationCheck) SetAmountNil(b bool)`

 SetAmountNil sets the value for Amount to be an explicit nil

### UnsetAmount
`func (o *VerifyV2ServiceVerificationCheck) UnsetAmount()`

UnsetAmount ensures that no value is present for Amount, not even an explicit nil
### GetChannel

`func (o *VerifyV2ServiceVerificationCheck) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *VerifyV2ServiceVerificationCheck) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *VerifyV2ServiceVerificationCheck) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *VerifyV2ServiceVerificationCheck) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### SetChannelNil

`func (o *VerifyV2ServiceVerificationCheck) SetChannelNil(b bool)`

 SetChannelNil sets the value for Channel to be an explicit nil

### UnsetChannel
`func (o *VerifyV2ServiceVerificationCheck) UnsetChannel()`

UnsetChannel ensures that no value is present for Channel, not even an explicit nil
### GetDateCreated

`func (o *VerifyV2ServiceVerificationCheck) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *VerifyV2ServiceVerificationCheck) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *VerifyV2ServiceVerificationCheck) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *VerifyV2ServiceVerificationCheck) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *VerifyV2ServiceVerificationCheck) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *VerifyV2ServiceVerificationCheck) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *VerifyV2ServiceVerificationCheck) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *VerifyV2ServiceVerificationCheck) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *VerifyV2ServiceVerificationCheck) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *VerifyV2ServiceVerificationCheck) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *VerifyV2ServiceVerificationCheck) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *VerifyV2ServiceVerificationCheck) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetPayee

`func (o *VerifyV2ServiceVerificationCheck) GetPayee() string`

GetPayee returns the Payee field if non-nil, zero value otherwise.

### GetPayeeOk

`func (o *VerifyV2ServiceVerificationCheck) GetPayeeOk() (*string, bool)`

GetPayeeOk returns a tuple with the Payee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayee

`func (o *VerifyV2ServiceVerificationCheck) SetPayee(v string)`

SetPayee sets Payee field to given value.

### HasPayee

`func (o *VerifyV2ServiceVerificationCheck) HasPayee() bool`

HasPayee returns a boolean if a field has been set.

### SetPayeeNil

`func (o *VerifyV2ServiceVerificationCheck) SetPayeeNil(b bool)`

 SetPayeeNil sets the value for Payee to be an explicit nil

### UnsetPayee
`func (o *VerifyV2ServiceVerificationCheck) UnsetPayee()`

UnsetPayee ensures that no value is present for Payee, not even an explicit nil
### GetServiceSid

`func (o *VerifyV2ServiceVerificationCheck) GetServiceSid() string`

GetServiceSid returns the ServiceSid field if non-nil, zero value otherwise.

### GetServiceSidOk

`func (o *VerifyV2ServiceVerificationCheck) GetServiceSidOk() (*string, bool)`

GetServiceSidOk returns a tuple with the ServiceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSid

`func (o *VerifyV2ServiceVerificationCheck) SetServiceSid(v string)`

SetServiceSid sets ServiceSid field to given value.

### HasServiceSid

`func (o *VerifyV2ServiceVerificationCheck) HasServiceSid() bool`

HasServiceSid returns a boolean if a field has been set.

### SetServiceSidNil

`func (o *VerifyV2ServiceVerificationCheck) SetServiceSidNil(b bool)`

 SetServiceSidNil sets the value for ServiceSid to be an explicit nil

### UnsetServiceSid
`func (o *VerifyV2ServiceVerificationCheck) UnsetServiceSid()`

UnsetServiceSid ensures that no value is present for ServiceSid, not even an explicit nil
### GetSid

`func (o *VerifyV2ServiceVerificationCheck) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *VerifyV2ServiceVerificationCheck) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *VerifyV2ServiceVerificationCheck) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *VerifyV2ServiceVerificationCheck) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *VerifyV2ServiceVerificationCheck) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *VerifyV2ServiceVerificationCheck) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetStatus

`func (o *VerifyV2ServiceVerificationCheck) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VerifyV2ServiceVerificationCheck) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VerifyV2ServiceVerificationCheck) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VerifyV2ServiceVerificationCheck) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *VerifyV2ServiceVerificationCheck) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *VerifyV2ServiceVerificationCheck) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetTo

`func (o *VerifyV2ServiceVerificationCheck) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *VerifyV2ServiceVerificationCheck) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *VerifyV2ServiceVerificationCheck) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *VerifyV2ServiceVerificationCheck) HasTo() bool`

HasTo returns a boolean if a field has been set.

### SetToNil

`func (o *VerifyV2ServiceVerificationCheck) SetToNil(b bool)`

 SetToNil sets the value for To to be an explicit nil

### UnsetTo
`func (o *VerifyV2ServiceVerificationCheck) UnsetTo()`

UnsetTo ensures that no value is present for To, not even an explicit nil
### GetValid

`func (o *VerifyV2ServiceVerificationCheck) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *VerifyV2ServiceVerificationCheck) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *VerifyV2ServiceVerificationCheck) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *VerifyV2ServiceVerificationCheck) HasValid() bool`

HasValid returns a boolean if a field has been set.

### SetValidNil

`func (o *VerifyV2ServiceVerificationCheck) SetValidNil(b bool)`

 SetValidNil sets the value for Valid to be an explicit nil

### UnsetValid
`func (o *VerifyV2ServiceVerificationCheck) UnsetValid()`

UnsetValid ensures that no value is present for Valid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


