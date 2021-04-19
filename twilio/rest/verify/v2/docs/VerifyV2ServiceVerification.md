# VerifyV2ServiceVerification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**Amount** | Pointer to **NullableString** | The amount of the associated PSD2 compliant transaction. | [optional] 
**Channel** | Pointer to **NullableString** | The verification method used. | [optional] 
**DateCreated** | Pointer to **NullableTime** | The RFC 2822 date and time in GMT when the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The RFC 2822 date and time in GMT when the resource was last updated | [optional] 
**Lookup** | Pointer to **map[string]interface{}** | Information about the phone number being verified | [optional] 
**Payee** | Pointer to **NullableString** | The payee of the associated PSD2 compliant transaction | [optional] 
**SendCodeAttempts** | Pointer to **[]map[string]interface{}** | An array of verification attempt objects. | [optional] 
**ServiceSid** | Pointer to **NullableString** | The SID of the Service that the resource is associated with | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**Status** | Pointer to **NullableString** | The status of the verification resource | [optional] 
**To** | Pointer to **NullableString** | The phone number or email being verified | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the Verification resource | [optional] 
**Valid** | Pointer to **NullableBool** | Whether the verification was successful | [optional] 

## Methods

### NewVerifyV2ServiceVerification

`func NewVerifyV2ServiceVerification() *VerifyV2ServiceVerification`

NewVerifyV2ServiceVerification instantiates a new VerifyV2ServiceVerification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifyV2ServiceVerificationWithDefaults

`func NewVerifyV2ServiceVerificationWithDefaults() *VerifyV2ServiceVerification`

NewVerifyV2ServiceVerificationWithDefaults instantiates a new VerifyV2ServiceVerification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *VerifyV2ServiceVerification) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *VerifyV2ServiceVerification) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *VerifyV2ServiceVerification) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *VerifyV2ServiceVerification) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *VerifyV2ServiceVerification) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *VerifyV2ServiceVerification) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetAmount

`func (o *VerifyV2ServiceVerification) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *VerifyV2ServiceVerification) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *VerifyV2ServiceVerification) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *VerifyV2ServiceVerification) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### SetAmountNil

`func (o *VerifyV2ServiceVerification) SetAmountNil(b bool)`

 SetAmountNil sets the value for Amount to be an explicit nil

### UnsetAmount
`func (o *VerifyV2ServiceVerification) UnsetAmount()`

UnsetAmount ensures that no value is present for Amount, not even an explicit nil
### GetChannel

`func (o *VerifyV2ServiceVerification) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *VerifyV2ServiceVerification) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *VerifyV2ServiceVerification) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *VerifyV2ServiceVerification) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### SetChannelNil

`func (o *VerifyV2ServiceVerification) SetChannelNil(b bool)`

 SetChannelNil sets the value for Channel to be an explicit nil

### UnsetChannel
`func (o *VerifyV2ServiceVerification) UnsetChannel()`

UnsetChannel ensures that no value is present for Channel, not even an explicit nil
### GetDateCreated

`func (o *VerifyV2ServiceVerification) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *VerifyV2ServiceVerification) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *VerifyV2ServiceVerification) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *VerifyV2ServiceVerification) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *VerifyV2ServiceVerification) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *VerifyV2ServiceVerification) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *VerifyV2ServiceVerification) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *VerifyV2ServiceVerification) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *VerifyV2ServiceVerification) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *VerifyV2ServiceVerification) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *VerifyV2ServiceVerification) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *VerifyV2ServiceVerification) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetLookup

`func (o *VerifyV2ServiceVerification) GetLookup() map[string]interface{}`

GetLookup returns the Lookup field if non-nil, zero value otherwise.

### GetLookupOk

`func (o *VerifyV2ServiceVerification) GetLookupOk() (*map[string]interface{}, bool)`

GetLookupOk returns a tuple with the Lookup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLookup

`func (o *VerifyV2ServiceVerification) SetLookup(v map[string]interface{})`

SetLookup sets Lookup field to given value.

### HasLookup

`func (o *VerifyV2ServiceVerification) HasLookup() bool`

HasLookup returns a boolean if a field has been set.

### SetLookupNil

`func (o *VerifyV2ServiceVerification) SetLookupNil(b bool)`

 SetLookupNil sets the value for Lookup to be an explicit nil

### UnsetLookup
`func (o *VerifyV2ServiceVerification) UnsetLookup()`

UnsetLookup ensures that no value is present for Lookup, not even an explicit nil
### GetPayee

`func (o *VerifyV2ServiceVerification) GetPayee() string`

GetPayee returns the Payee field if non-nil, zero value otherwise.

### GetPayeeOk

`func (o *VerifyV2ServiceVerification) GetPayeeOk() (*string, bool)`

GetPayeeOk returns a tuple with the Payee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayee

`func (o *VerifyV2ServiceVerification) SetPayee(v string)`

SetPayee sets Payee field to given value.

### HasPayee

`func (o *VerifyV2ServiceVerification) HasPayee() bool`

HasPayee returns a boolean if a field has been set.

### SetPayeeNil

`func (o *VerifyV2ServiceVerification) SetPayeeNil(b bool)`

 SetPayeeNil sets the value for Payee to be an explicit nil

### UnsetPayee
`func (o *VerifyV2ServiceVerification) UnsetPayee()`

UnsetPayee ensures that no value is present for Payee, not even an explicit nil
### GetSendCodeAttempts

`func (o *VerifyV2ServiceVerification) GetSendCodeAttempts() []map[string]interface{}`

GetSendCodeAttempts returns the SendCodeAttempts field if non-nil, zero value otherwise.

### GetSendCodeAttemptsOk

`func (o *VerifyV2ServiceVerification) GetSendCodeAttemptsOk() (*[]map[string]interface{}, bool)`

GetSendCodeAttemptsOk returns a tuple with the SendCodeAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendCodeAttempts

`func (o *VerifyV2ServiceVerification) SetSendCodeAttempts(v []map[string]interface{})`

SetSendCodeAttempts sets SendCodeAttempts field to given value.

### HasSendCodeAttempts

`func (o *VerifyV2ServiceVerification) HasSendCodeAttempts() bool`

HasSendCodeAttempts returns a boolean if a field has been set.

### SetSendCodeAttemptsNil

`func (o *VerifyV2ServiceVerification) SetSendCodeAttemptsNil(b bool)`

 SetSendCodeAttemptsNil sets the value for SendCodeAttempts to be an explicit nil

### UnsetSendCodeAttempts
`func (o *VerifyV2ServiceVerification) UnsetSendCodeAttempts()`

UnsetSendCodeAttempts ensures that no value is present for SendCodeAttempts, not even an explicit nil
### GetServiceSid

`func (o *VerifyV2ServiceVerification) GetServiceSid() string`

GetServiceSid returns the ServiceSid field if non-nil, zero value otherwise.

### GetServiceSidOk

`func (o *VerifyV2ServiceVerification) GetServiceSidOk() (*string, bool)`

GetServiceSidOk returns a tuple with the ServiceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSid

`func (o *VerifyV2ServiceVerification) SetServiceSid(v string)`

SetServiceSid sets ServiceSid field to given value.

### HasServiceSid

`func (o *VerifyV2ServiceVerification) HasServiceSid() bool`

HasServiceSid returns a boolean if a field has been set.

### SetServiceSidNil

`func (o *VerifyV2ServiceVerification) SetServiceSidNil(b bool)`

 SetServiceSidNil sets the value for ServiceSid to be an explicit nil

### UnsetServiceSid
`func (o *VerifyV2ServiceVerification) UnsetServiceSid()`

UnsetServiceSid ensures that no value is present for ServiceSid, not even an explicit nil
### GetSid

`func (o *VerifyV2ServiceVerification) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *VerifyV2ServiceVerification) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *VerifyV2ServiceVerification) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *VerifyV2ServiceVerification) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *VerifyV2ServiceVerification) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *VerifyV2ServiceVerification) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetStatus

`func (o *VerifyV2ServiceVerification) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VerifyV2ServiceVerification) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VerifyV2ServiceVerification) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VerifyV2ServiceVerification) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *VerifyV2ServiceVerification) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *VerifyV2ServiceVerification) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetTo

`func (o *VerifyV2ServiceVerification) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *VerifyV2ServiceVerification) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *VerifyV2ServiceVerification) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *VerifyV2ServiceVerification) HasTo() bool`

HasTo returns a boolean if a field has been set.

### SetToNil

`func (o *VerifyV2ServiceVerification) SetToNil(b bool)`

 SetToNil sets the value for To to be an explicit nil

### UnsetTo
`func (o *VerifyV2ServiceVerification) UnsetTo()`

UnsetTo ensures that no value is present for To, not even an explicit nil
### GetUrl

`func (o *VerifyV2ServiceVerification) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *VerifyV2ServiceVerification) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *VerifyV2ServiceVerification) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *VerifyV2ServiceVerification) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *VerifyV2ServiceVerification) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *VerifyV2ServiceVerification) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetValid

`func (o *VerifyV2ServiceVerification) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *VerifyV2ServiceVerification) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *VerifyV2ServiceVerification) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *VerifyV2ServiceVerification) HasValid() bool`

HasValid returns a boolean if a field has been set.

### SetValidNil

`func (o *VerifyV2ServiceVerification) SetValidNil(b bool)`

 SetValidNil sets the value for Valid to be an explicit nil

### UnsetValid
`func (o *VerifyV2ServiceVerification) UnsetValid()`

UnsetValid ensures that no value is present for Valid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


