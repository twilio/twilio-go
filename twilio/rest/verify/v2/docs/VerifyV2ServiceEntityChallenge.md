# VerifyV2ServiceEntityChallenge

## Properties

Name | Type | Description
------------ | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | Account Sid. | [optional] 
**DateCreated** | Pointer to **NullableTime** | The date this Challenge was created | [optional] 
**DateResponded** | Pointer to **NullableTime** | The date this Challenge was responded | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The date this Challenge was updated | [optional] 
**Details** | Pointer to **map[string]interface{}** | Details about the Challenge. | [optional] 
**EntitySid** | Pointer to **NullableString** | Entity Sid. | [optional] 
**ExpirationDate** | Pointer to **NullableTime** | The date-time when this Challenge expires | [optional] 
**FactorSid** | Pointer to **NullableString** | Factor Sid. | [optional] 
**FactorType** | Pointer to **NullableString** | The Factor Type of this Challenge | [optional] 
**HiddenDetails** | Pointer to **map[string]interface{}** | Hidden details about the Challenge | [optional] 
**Identity** | Pointer to **NullableString** | Unique external identifier of the Entity | [optional] 
**RespondedReason** | Pointer to **NullableString** | The Reason of this Challenge &#x60;status&#x60; | [optional] 
**ServiceSid** | Pointer to **NullableString** | Service Sid. | [optional] 
**Sid** | Pointer to **NullableString** | A string that uniquely identifies this Challenge. | [optional] 
**Status** | Pointer to **NullableString** | The Status of this Challenge | [optional] 
**Url** | Pointer to **NullableString** | The URL of this resource. | [optional] 

## Methods

### NewVerifyV2ServiceEntityChallenge

`func NewVerifyV2ServiceEntityChallenge() *VerifyV2ServiceEntityChallenge`

NewVerifyV2ServiceEntityChallenge instantiates a new VerifyV2ServiceEntityChallenge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifyV2ServiceEntityChallengeWithDefaults

`func NewVerifyV2ServiceEntityChallengeWithDefaults() *VerifyV2ServiceEntityChallenge`

NewVerifyV2ServiceEntityChallengeWithDefaults instantiates a new VerifyV2ServiceEntityChallenge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *VerifyV2ServiceEntityChallenge) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *VerifyV2ServiceEntityChallenge) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *VerifyV2ServiceEntityChallenge) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *VerifyV2ServiceEntityChallenge) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *VerifyV2ServiceEntityChallenge) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *VerifyV2ServiceEntityChallenge) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetDateCreated

`func (o *VerifyV2ServiceEntityChallenge) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *VerifyV2ServiceEntityChallenge) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *VerifyV2ServiceEntityChallenge) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *VerifyV2ServiceEntityChallenge) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *VerifyV2ServiceEntityChallenge) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *VerifyV2ServiceEntityChallenge) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateResponded

`func (o *VerifyV2ServiceEntityChallenge) GetDateResponded() time.Time`

GetDateResponded returns the DateResponded field if non-nil, zero value otherwise.

### GetDateRespondedOk

`func (o *VerifyV2ServiceEntityChallenge) GetDateRespondedOk() (*time.Time, bool)`

GetDateRespondedOk returns a tuple with the DateResponded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateResponded

`func (o *VerifyV2ServiceEntityChallenge) SetDateResponded(v time.Time)`

SetDateResponded sets DateResponded field to given value.

### HasDateResponded

`func (o *VerifyV2ServiceEntityChallenge) HasDateResponded() bool`

HasDateResponded returns a boolean if a field has been set.

### SetDateRespondedNil

`func (o *VerifyV2ServiceEntityChallenge) SetDateRespondedNil(b bool)`

 SetDateRespondedNil sets the value for DateResponded to be an explicit nil

### UnsetDateResponded
`func (o *VerifyV2ServiceEntityChallenge) UnsetDateResponded()`

UnsetDateResponded ensures that no value is present for DateResponded, not even an explicit nil
### GetDateUpdated

`func (o *VerifyV2ServiceEntityChallenge) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *VerifyV2ServiceEntityChallenge) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *VerifyV2ServiceEntityChallenge) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *VerifyV2ServiceEntityChallenge) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *VerifyV2ServiceEntityChallenge) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *VerifyV2ServiceEntityChallenge) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetDetails

`func (o *VerifyV2ServiceEntityChallenge) GetDetails() map[string]interface{}`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *VerifyV2ServiceEntityChallenge) GetDetailsOk() (*map[string]interface{}, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *VerifyV2ServiceEntityChallenge) SetDetails(v map[string]interface{})`

SetDetails sets Details field to given value.

### HasDetails

`func (o *VerifyV2ServiceEntityChallenge) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### SetDetailsNil

`func (o *VerifyV2ServiceEntityChallenge) SetDetailsNil(b bool)`

 SetDetailsNil sets the value for Details to be an explicit nil

### UnsetDetails
`func (o *VerifyV2ServiceEntityChallenge) UnsetDetails()`

UnsetDetails ensures that no value is present for Details, not even an explicit nil
### GetEntitySid

`func (o *VerifyV2ServiceEntityChallenge) GetEntitySid() string`

GetEntitySid returns the EntitySid field if non-nil, zero value otherwise.

### GetEntitySidOk

`func (o *VerifyV2ServiceEntityChallenge) GetEntitySidOk() (*string, bool)`

GetEntitySidOk returns a tuple with the EntitySid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitySid

`func (o *VerifyV2ServiceEntityChallenge) SetEntitySid(v string)`

SetEntitySid sets EntitySid field to given value.

### HasEntitySid

`func (o *VerifyV2ServiceEntityChallenge) HasEntitySid() bool`

HasEntitySid returns a boolean if a field has been set.

### SetEntitySidNil

`func (o *VerifyV2ServiceEntityChallenge) SetEntitySidNil(b bool)`

 SetEntitySidNil sets the value for EntitySid to be an explicit nil

### UnsetEntitySid
`func (o *VerifyV2ServiceEntityChallenge) UnsetEntitySid()`

UnsetEntitySid ensures that no value is present for EntitySid, not even an explicit nil
### GetExpirationDate

`func (o *VerifyV2ServiceEntityChallenge) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *VerifyV2ServiceEntityChallenge) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *VerifyV2ServiceEntityChallenge) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *VerifyV2ServiceEntityChallenge) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### SetExpirationDateNil

`func (o *VerifyV2ServiceEntityChallenge) SetExpirationDateNil(b bool)`

 SetExpirationDateNil sets the value for ExpirationDate to be an explicit nil

### UnsetExpirationDate
`func (o *VerifyV2ServiceEntityChallenge) UnsetExpirationDate()`

UnsetExpirationDate ensures that no value is present for ExpirationDate, not even an explicit nil
### GetFactorSid

`func (o *VerifyV2ServiceEntityChallenge) GetFactorSid() string`

GetFactorSid returns the FactorSid field if non-nil, zero value otherwise.

### GetFactorSidOk

`func (o *VerifyV2ServiceEntityChallenge) GetFactorSidOk() (*string, bool)`

GetFactorSidOk returns a tuple with the FactorSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactorSid

`func (o *VerifyV2ServiceEntityChallenge) SetFactorSid(v string)`

SetFactorSid sets FactorSid field to given value.

### HasFactorSid

`func (o *VerifyV2ServiceEntityChallenge) HasFactorSid() bool`

HasFactorSid returns a boolean if a field has been set.

### SetFactorSidNil

`func (o *VerifyV2ServiceEntityChallenge) SetFactorSidNil(b bool)`

 SetFactorSidNil sets the value for FactorSid to be an explicit nil

### UnsetFactorSid
`func (o *VerifyV2ServiceEntityChallenge) UnsetFactorSid()`

UnsetFactorSid ensures that no value is present for FactorSid, not even an explicit nil
### GetFactorType

`func (o *VerifyV2ServiceEntityChallenge) GetFactorType() string`

GetFactorType returns the FactorType field if non-nil, zero value otherwise.

### GetFactorTypeOk

`func (o *VerifyV2ServiceEntityChallenge) GetFactorTypeOk() (*string, bool)`

GetFactorTypeOk returns a tuple with the FactorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactorType

`func (o *VerifyV2ServiceEntityChallenge) SetFactorType(v string)`

SetFactorType sets FactorType field to given value.

### HasFactorType

`func (o *VerifyV2ServiceEntityChallenge) HasFactorType() bool`

HasFactorType returns a boolean if a field has been set.

### SetFactorTypeNil

`func (o *VerifyV2ServiceEntityChallenge) SetFactorTypeNil(b bool)`

 SetFactorTypeNil sets the value for FactorType to be an explicit nil

### UnsetFactorType
`func (o *VerifyV2ServiceEntityChallenge) UnsetFactorType()`

UnsetFactorType ensures that no value is present for FactorType, not even an explicit nil
### GetHiddenDetails

`func (o *VerifyV2ServiceEntityChallenge) GetHiddenDetails() map[string]interface{}`

GetHiddenDetails returns the HiddenDetails field if non-nil, zero value otherwise.

### GetHiddenDetailsOk

`func (o *VerifyV2ServiceEntityChallenge) GetHiddenDetailsOk() (*map[string]interface{}, bool)`

GetHiddenDetailsOk returns a tuple with the HiddenDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiddenDetails

`func (o *VerifyV2ServiceEntityChallenge) SetHiddenDetails(v map[string]interface{})`

SetHiddenDetails sets HiddenDetails field to given value.

### HasHiddenDetails

`func (o *VerifyV2ServiceEntityChallenge) HasHiddenDetails() bool`

HasHiddenDetails returns a boolean if a field has been set.

### SetHiddenDetailsNil

`func (o *VerifyV2ServiceEntityChallenge) SetHiddenDetailsNil(b bool)`

 SetHiddenDetailsNil sets the value for HiddenDetails to be an explicit nil

### UnsetHiddenDetails
`func (o *VerifyV2ServiceEntityChallenge) UnsetHiddenDetails()`

UnsetHiddenDetails ensures that no value is present for HiddenDetails, not even an explicit nil
### GetIdentity

`func (o *VerifyV2ServiceEntityChallenge) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *VerifyV2ServiceEntityChallenge) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *VerifyV2ServiceEntityChallenge) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *VerifyV2ServiceEntityChallenge) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### SetIdentityNil

`func (o *VerifyV2ServiceEntityChallenge) SetIdentityNil(b bool)`

 SetIdentityNil sets the value for Identity to be an explicit nil

### UnsetIdentity
`func (o *VerifyV2ServiceEntityChallenge) UnsetIdentity()`

UnsetIdentity ensures that no value is present for Identity, not even an explicit nil
### GetRespondedReason

`func (o *VerifyV2ServiceEntityChallenge) GetRespondedReason() string`

GetRespondedReason returns the RespondedReason field if non-nil, zero value otherwise.

### GetRespondedReasonOk

`func (o *VerifyV2ServiceEntityChallenge) GetRespondedReasonOk() (*string, bool)`

GetRespondedReasonOk returns a tuple with the RespondedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRespondedReason

`func (o *VerifyV2ServiceEntityChallenge) SetRespondedReason(v string)`

SetRespondedReason sets RespondedReason field to given value.

### HasRespondedReason

`func (o *VerifyV2ServiceEntityChallenge) HasRespondedReason() bool`

HasRespondedReason returns a boolean if a field has been set.

### SetRespondedReasonNil

`func (o *VerifyV2ServiceEntityChallenge) SetRespondedReasonNil(b bool)`

 SetRespondedReasonNil sets the value for RespondedReason to be an explicit nil

### UnsetRespondedReason
`func (o *VerifyV2ServiceEntityChallenge) UnsetRespondedReason()`

UnsetRespondedReason ensures that no value is present for RespondedReason, not even an explicit nil
### GetServiceSid

`func (o *VerifyV2ServiceEntityChallenge) GetServiceSid() string`

GetServiceSid returns the ServiceSid field if non-nil, zero value otherwise.

### GetServiceSidOk

`func (o *VerifyV2ServiceEntityChallenge) GetServiceSidOk() (*string, bool)`

GetServiceSidOk returns a tuple with the ServiceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSid

`func (o *VerifyV2ServiceEntityChallenge) SetServiceSid(v string)`

SetServiceSid sets ServiceSid field to given value.

### HasServiceSid

`func (o *VerifyV2ServiceEntityChallenge) HasServiceSid() bool`

HasServiceSid returns a boolean if a field has been set.

### SetServiceSidNil

`func (o *VerifyV2ServiceEntityChallenge) SetServiceSidNil(b bool)`

 SetServiceSidNil sets the value for ServiceSid to be an explicit nil

### UnsetServiceSid
`func (o *VerifyV2ServiceEntityChallenge) UnsetServiceSid()`

UnsetServiceSid ensures that no value is present for ServiceSid, not even an explicit nil
### GetSid

`func (o *VerifyV2ServiceEntityChallenge) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *VerifyV2ServiceEntityChallenge) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *VerifyV2ServiceEntityChallenge) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *VerifyV2ServiceEntityChallenge) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *VerifyV2ServiceEntityChallenge) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *VerifyV2ServiceEntityChallenge) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetStatus

`func (o *VerifyV2ServiceEntityChallenge) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VerifyV2ServiceEntityChallenge) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VerifyV2ServiceEntityChallenge) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VerifyV2ServiceEntityChallenge) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *VerifyV2ServiceEntityChallenge) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *VerifyV2ServiceEntityChallenge) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetUrl

`func (o *VerifyV2ServiceEntityChallenge) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *VerifyV2ServiceEntityChallenge) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *VerifyV2ServiceEntityChallenge) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *VerifyV2ServiceEntityChallenge) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *VerifyV2ServiceEntityChallenge) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *VerifyV2ServiceEntityChallenge) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


