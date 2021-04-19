# SupersimV1Fleet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**CommandsEnabled** | Pointer to **NullableBool** | Defines whether SIMs in the Fleet are capable of sending and receiving machine-to-machine SMS via Commands | [optional] 
**CommandsMethod** | Pointer to **NullableString** | A string representing the HTTP method to use when making a request to &#x60;commands_url&#x60; | [optional] 
**CommandsUrl** | Pointer to **NullableString** | The URL that will receive a webhook when a Super SIM in the Fleet is used to send an SMS from your device to the Commands number | [optional] 
**DataEnabled** | Pointer to **NullableBool** | Defines whether SIMs in the Fleet are capable of using data connectivity | [optional] 
**DataLimit** | Pointer to **NullableInt32** | The total data usage (download and upload combined) in Megabytes that each Sim resource assigned to the Fleet resource can consume | [optional] 
**DataMetering** | Pointer to **NullableString** | The model by which a SIM is metered and billed | [optional] 
**DateCreated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was last updated | [optional] 
**NetworkAccessProfileSid** | Pointer to **NullableString** | The SID of the Network Access Profile of the Fleet | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**SmsCommandsEnabled** | Pointer to **NullableBool** | Defines whether SIMs in the Fleet are capable of sending and receiving machine-to-machine SMS via Commands | [optional] 
**SmsCommandsMethod** | Pointer to **NullableString** | A string representing the HTTP method to use when making a request to &#x60;sms_commands_url&#x60; | [optional] 
**SmsCommandsUrl** | Pointer to **NullableString** | The URL that will receive a webhook when a Super SIM in the Fleet is used to send an SMS from your device to the SMS Commands number | [optional] 
**UniqueName** | Pointer to **NullableString** | An application-defined string that uniquely identifies the resource | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the Fleet resource | [optional] 

## Methods

### NewSupersimV1Fleet

`func NewSupersimV1Fleet() *SupersimV1Fleet`

NewSupersimV1Fleet instantiates a new SupersimV1Fleet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupersimV1FleetWithDefaults

`func NewSupersimV1FleetWithDefaults() *SupersimV1Fleet`

NewSupersimV1FleetWithDefaults instantiates a new SupersimV1Fleet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *SupersimV1Fleet) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *SupersimV1Fleet) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *SupersimV1Fleet) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *SupersimV1Fleet) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *SupersimV1Fleet) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *SupersimV1Fleet) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetCommandsEnabled

`func (o *SupersimV1Fleet) GetCommandsEnabled() bool`

GetCommandsEnabled returns the CommandsEnabled field if non-nil, zero value otherwise.

### GetCommandsEnabledOk

`func (o *SupersimV1Fleet) GetCommandsEnabledOk() (*bool, bool)`

GetCommandsEnabledOk returns a tuple with the CommandsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandsEnabled

`func (o *SupersimV1Fleet) SetCommandsEnabled(v bool)`

SetCommandsEnabled sets CommandsEnabled field to given value.

### HasCommandsEnabled

`func (o *SupersimV1Fleet) HasCommandsEnabled() bool`

HasCommandsEnabled returns a boolean if a field has been set.

### SetCommandsEnabledNil

`func (o *SupersimV1Fleet) SetCommandsEnabledNil(b bool)`

 SetCommandsEnabledNil sets the value for CommandsEnabled to be an explicit nil

### UnsetCommandsEnabled
`func (o *SupersimV1Fleet) UnsetCommandsEnabled()`

UnsetCommandsEnabled ensures that no value is present for CommandsEnabled, not even an explicit nil
### GetCommandsMethod

`func (o *SupersimV1Fleet) GetCommandsMethod() string`

GetCommandsMethod returns the CommandsMethod field if non-nil, zero value otherwise.

### GetCommandsMethodOk

`func (o *SupersimV1Fleet) GetCommandsMethodOk() (*string, bool)`

GetCommandsMethodOk returns a tuple with the CommandsMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandsMethod

`func (o *SupersimV1Fleet) SetCommandsMethod(v string)`

SetCommandsMethod sets CommandsMethod field to given value.

### HasCommandsMethod

`func (o *SupersimV1Fleet) HasCommandsMethod() bool`

HasCommandsMethod returns a boolean if a field has been set.

### SetCommandsMethodNil

`func (o *SupersimV1Fleet) SetCommandsMethodNil(b bool)`

 SetCommandsMethodNil sets the value for CommandsMethod to be an explicit nil

### UnsetCommandsMethod
`func (o *SupersimV1Fleet) UnsetCommandsMethod()`

UnsetCommandsMethod ensures that no value is present for CommandsMethod, not even an explicit nil
### GetCommandsUrl

`func (o *SupersimV1Fleet) GetCommandsUrl() string`

GetCommandsUrl returns the CommandsUrl field if non-nil, zero value otherwise.

### GetCommandsUrlOk

`func (o *SupersimV1Fleet) GetCommandsUrlOk() (*string, bool)`

GetCommandsUrlOk returns a tuple with the CommandsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandsUrl

`func (o *SupersimV1Fleet) SetCommandsUrl(v string)`

SetCommandsUrl sets CommandsUrl field to given value.

### HasCommandsUrl

`func (o *SupersimV1Fleet) HasCommandsUrl() bool`

HasCommandsUrl returns a boolean if a field has been set.

### SetCommandsUrlNil

`func (o *SupersimV1Fleet) SetCommandsUrlNil(b bool)`

 SetCommandsUrlNil sets the value for CommandsUrl to be an explicit nil

### UnsetCommandsUrl
`func (o *SupersimV1Fleet) UnsetCommandsUrl()`

UnsetCommandsUrl ensures that no value is present for CommandsUrl, not even an explicit nil
### GetDataEnabled

`func (o *SupersimV1Fleet) GetDataEnabled() bool`

GetDataEnabled returns the DataEnabled field if non-nil, zero value otherwise.

### GetDataEnabledOk

`func (o *SupersimV1Fleet) GetDataEnabledOk() (*bool, bool)`

GetDataEnabledOk returns a tuple with the DataEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataEnabled

`func (o *SupersimV1Fleet) SetDataEnabled(v bool)`

SetDataEnabled sets DataEnabled field to given value.

### HasDataEnabled

`func (o *SupersimV1Fleet) HasDataEnabled() bool`

HasDataEnabled returns a boolean if a field has been set.

### SetDataEnabledNil

`func (o *SupersimV1Fleet) SetDataEnabledNil(b bool)`

 SetDataEnabledNil sets the value for DataEnabled to be an explicit nil

### UnsetDataEnabled
`func (o *SupersimV1Fleet) UnsetDataEnabled()`

UnsetDataEnabled ensures that no value is present for DataEnabled, not even an explicit nil
### GetDataLimit

`func (o *SupersimV1Fleet) GetDataLimit() int32`

GetDataLimit returns the DataLimit field if non-nil, zero value otherwise.

### GetDataLimitOk

`func (o *SupersimV1Fleet) GetDataLimitOk() (*int32, bool)`

GetDataLimitOk returns a tuple with the DataLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLimit

`func (o *SupersimV1Fleet) SetDataLimit(v int32)`

SetDataLimit sets DataLimit field to given value.

### HasDataLimit

`func (o *SupersimV1Fleet) HasDataLimit() bool`

HasDataLimit returns a boolean if a field has been set.

### SetDataLimitNil

`func (o *SupersimV1Fleet) SetDataLimitNil(b bool)`

 SetDataLimitNil sets the value for DataLimit to be an explicit nil

### UnsetDataLimit
`func (o *SupersimV1Fleet) UnsetDataLimit()`

UnsetDataLimit ensures that no value is present for DataLimit, not even an explicit nil
### GetDataMetering

`func (o *SupersimV1Fleet) GetDataMetering() string`

GetDataMetering returns the DataMetering field if non-nil, zero value otherwise.

### GetDataMeteringOk

`func (o *SupersimV1Fleet) GetDataMeteringOk() (*string, bool)`

GetDataMeteringOk returns a tuple with the DataMetering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataMetering

`func (o *SupersimV1Fleet) SetDataMetering(v string)`

SetDataMetering sets DataMetering field to given value.

### HasDataMetering

`func (o *SupersimV1Fleet) HasDataMetering() bool`

HasDataMetering returns a boolean if a field has been set.

### SetDataMeteringNil

`func (o *SupersimV1Fleet) SetDataMeteringNil(b bool)`

 SetDataMeteringNil sets the value for DataMetering to be an explicit nil

### UnsetDataMetering
`func (o *SupersimV1Fleet) UnsetDataMetering()`

UnsetDataMetering ensures that no value is present for DataMetering, not even an explicit nil
### GetDateCreated

`func (o *SupersimV1Fleet) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *SupersimV1Fleet) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *SupersimV1Fleet) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *SupersimV1Fleet) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *SupersimV1Fleet) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *SupersimV1Fleet) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *SupersimV1Fleet) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *SupersimV1Fleet) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *SupersimV1Fleet) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *SupersimV1Fleet) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *SupersimV1Fleet) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *SupersimV1Fleet) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetNetworkAccessProfileSid

`func (o *SupersimV1Fleet) GetNetworkAccessProfileSid() string`

GetNetworkAccessProfileSid returns the NetworkAccessProfileSid field if non-nil, zero value otherwise.

### GetNetworkAccessProfileSidOk

`func (o *SupersimV1Fleet) GetNetworkAccessProfileSidOk() (*string, bool)`

GetNetworkAccessProfileSidOk returns a tuple with the NetworkAccessProfileSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkAccessProfileSid

`func (o *SupersimV1Fleet) SetNetworkAccessProfileSid(v string)`

SetNetworkAccessProfileSid sets NetworkAccessProfileSid field to given value.

### HasNetworkAccessProfileSid

`func (o *SupersimV1Fleet) HasNetworkAccessProfileSid() bool`

HasNetworkAccessProfileSid returns a boolean if a field has been set.

### SetNetworkAccessProfileSidNil

`func (o *SupersimV1Fleet) SetNetworkAccessProfileSidNil(b bool)`

 SetNetworkAccessProfileSidNil sets the value for NetworkAccessProfileSid to be an explicit nil

### UnsetNetworkAccessProfileSid
`func (o *SupersimV1Fleet) UnsetNetworkAccessProfileSid()`

UnsetNetworkAccessProfileSid ensures that no value is present for NetworkAccessProfileSid, not even an explicit nil
### GetSid

`func (o *SupersimV1Fleet) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *SupersimV1Fleet) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *SupersimV1Fleet) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *SupersimV1Fleet) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *SupersimV1Fleet) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *SupersimV1Fleet) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetSmsCommandsEnabled

`func (o *SupersimV1Fleet) GetSmsCommandsEnabled() bool`

GetSmsCommandsEnabled returns the SmsCommandsEnabled field if non-nil, zero value otherwise.

### GetSmsCommandsEnabledOk

`func (o *SupersimV1Fleet) GetSmsCommandsEnabledOk() (*bool, bool)`

GetSmsCommandsEnabledOk returns a tuple with the SmsCommandsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsCommandsEnabled

`func (o *SupersimV1Fleet) SetSmsCommandsEnabled(v bool)`

SetSmsCommandsEnabled sets SmsCommandsEnabled field to given value.

### HasSmsCommandsEnabled

`func (o *SupersimV1Fleet) HasSmsCommandsEnabled() bool`

HasSmsCommandsEnabled returns a boolean if a field has been set.

### SetSmsCommandsEnabledNil

`func (o *SupersimV1Fleet) SetSmsCommandsEnabledNil(b bool)`

 SetSmsCommandsEnabledNil sets the value for SmsCommandsEnabled to be an explicit nil

### UnsetSmsCommandsEnabled
`func (o *SupersimV1Fleet) UnsetSmsCommandsEnabled()`

UnsetSmsCommandsEnabled ensures that no value is present for SmsCommandsEnabled, not even an explicit nil
### GetSmsCommandsMethod

`func (o *SupersimV1Fleet) GetSmsCommandsMethod() string`

GetSmsCommandsMethod returns the SmsCommandsMethod field if non-nil, zero value otherwise.

### GetSmsCommandsMethodOk

`func (o *SupersimV1Fleet) GetSmsCommandsMethodOk() (*string, bool)`

GetSmsCommandsMethodOk returns a tuple with the SmsCommandsMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsCommandsMethod

`func (o *SupersimV1Fleet) SetSmsCommandsMethod(v string)`

SetSmsCommandsMethod sets SmsCommandsMethod field to given value.

### HasSmsCommandsMethod

`func (o *SupersimV1Fleet) HasSmsCommandsMethod() bool`

HasSmsCommandsMethod returns a boolean if a field has been set.

### SetSmsCommandsMethodNil

`func (o *SupersimV1Fleet) SetSmsCommandsMethodNil(b bool)`

 SetSmsCommandsMethodNil sets the value for SmsCommandsMethod to be an explicit nil

### UnsetSmsCommandsMethod
`func (o *SupersimV1Fleet) UnsetSmsCommandsMethod()`

UnsetSmsCommandsMethod ensures that no value is present for SmsCommandsMethod, not even an explicit nil
### GetSmsCommandsUrl

`func (o *SupersimV1Fleet) GetSmsCommandsUrl() string`

GetSmsCommandsUrl returns the SmsCommandsUrl field if non-nil, zero value otherwise.

### GetSmsCommandsUrlOk

`func (o *SupersimV1Fleet) GetSmsCommandsUrlOk() (*string, bool)`

GetSmsCommandsUrlOk returns a tuple with the SmsCommandsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsCommandsUrl

`func (o *SupersimV1Fleet) SetSmsCommandsUrl(v string)`

SetSmsCommandsUrl sets SmsCommandsUrl field to given value.

### HasSmsCommandsUrl

`func (o *SupersimV1Fleet) HasSmsCommandsUrl() bool`

HasSmsCommandsUrl returns a boolean if a field has been set.

### SetSmsCommandsUrlNil

`func (o *SupersimV1Fleet) SetSmsCommandsUrlNil(b bool)`

 SetSmsCommandsUrlNil sets the value for SmsCommandsUrl to be an explicit nil

### UnsetSmsCommandsUrl
`func (o *SupersimV1Fleet) UnsetSmsCommandsUrl()`

UnsetSmsCommandsUrl ensures that no value is present for SmsCommandsUrl, not even an explicit nil
### GetUniqueName

`func (o *SupersimV1Fleet) GetUniqueName() string`

GetUniqueName returns the UniqueName field if non-nil, zero value otherwise.

### GetUniqueNameOk

`func (o *SupersimV1Fleet) GetUniqueNameOk() (*string, bool)`

GetUniqueNameOk returns a tuple with the UniqueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueName

`func (o *SupersimV1Fleet) SetUniqueName(v string)`

SetUniqueName sets UniqueName field to given value.

### HasUniqueName

`func (o *SupersimV1Fleet) HasUniqueName() bool`

HasUniqueName returns a boolean if a field has been set.

### SetUniqueNameNil

`func (o *SupersimV1Fleet) SetUniqueNameNil(b bool)`

 SetUniqueNameNil sets the value for UniqueName to be an explicit nil

### UnsetUniqueName
`func (o *SupersimV1Fleet) UnsetUniqueName()`

UnsetUniqueName ensures that no value is present for UniqueName, not even an explicit nil
### GetUrl

`func (o *SupersimV1Fleet) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SupersimV1Fleet) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SupersimV1Fleet) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SupersimV1Fleet) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *SupersimV1Fleet) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *SupersimV1Fleet) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


