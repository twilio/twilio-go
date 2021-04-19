# WirelessV1SimUsageRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**Commands** | Pointer to **map[string]interface{}** | An object that describes the SIM&#39;s usage of Commands during the specified period | [optional] 
**Data** | Pointer to **map[string]interface{}** | An object that describes the SIM&#39;s data usage during the specified period | [optional] 
**Period** | Pointer to **map[string]interface{}** | The time period for which the usage is reported | [optional] 
**SimSid** | Pointer to **NullableString** | The SID of the Sim resource that this Usage Record is for | [optional] 

## Methods

### NewWirelessV1SimUsageRecord

`func NewWirelessV1SimUsageRecord() *WirelessV1SimUsageRecord`

NewWirelessV1SimUsageRecord instantiates a new WirelessV1SimUsageRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWirelessV1SimUsageRecordWithDefaults

`func NewWirelessV1SimUsageRecordWithDefaults() *WirelessV1SimUsageRecord`

NewWirelessV1SimUsageRecordWithDefaults instantiates a new WirelessV1SimUsageRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *WirelessV1SimUsageRecord) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *WirelessV1SimUsageRecord) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *WirelessV1SimUsageRecord) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *WirelessV1SimUsageRecord) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *WirelessV1SimUsageRecord) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *WirelessV1SimUsageRecord) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetCommands

`func (o *WirelessV1SimUsageRecord) GetCommands() map[string]interface{}`

GetCommands returns the Commands field if non-nil, zero value otherwise.

### GetCommandsOk

`func (o *WirelessV1SimUsageRecord) GetCommandsOk() (*map[string]interface{}, bool)`

GetCommandsOk returns a tuple with the Commands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommands

`func (o *WirelessV1SimUsageRecord) SetCommands(v map[string]interface{})`

SetCommands sets Commands field to given value.

### HasCommands

`func (o *WirelessV1SimUsageRecord) HasCommands() bool`

HasCommands returns a boolean if a field has been set.

### SetCommandsNil

`func (o *WirelessV1SimUsageRecord) SetCommandsNil(b bool)`

 SetCommandsNil sets the value for Commands to be an explicit nil

### UnsetCommands
`func (o *WirelessV1SimUsageRecord) UnsetCommands()`

UnsetCommands ensures that no value is present for Commands, not even an explicit nil
### GetData

`func (o *WirelessV1SimUsageRecord) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WirelessV1SimUsageRecord) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WirelessV1SimUsageRecord) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *WirelessV1SimUsageRecord) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *WirelessV1SimUsageRecord) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *WirelessV1SimUsageRecord) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetPeriod

`func (o *WirelessV1SimUsageRecord) GetPeriod() map[string]interface{}`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *WirelessV1SimUsageRecord) GetPeriodOk() (*map[string]interface{}, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *WirelessV1SimUsageRecord) SetPeriod(v map[string]interface{})`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *WirelessV1SimUsageRecord) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### SetPeriodNil

`func (o *WirelessV1SimUsageRecord) SetPeriodNil(b bool)`

 SetPeriodNil sets the value for Period to be an explicit nil

### UnsetPeriod
`func (o *WirelessV1SimUsageRecord) UnsetPeriod()`

UnsetPeriod ensures that no value is present for Period, not even an explicit nil
### GetSimSid

`func (o *WirelessV1SimUsageRecord) GetSimSid() string`

GetSimSid returns the SimSid field if non-nil, zero value otherwise.

### GetSimSidOk

`func (o *WirelessV1SimUsageRecord) GetSimSidOk() (*string, bool)`

GetSimSidOk returns a tuple with the SimSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimSid

`func (o *WirelessV1SimUsageRecord) SetSimSid(v string)`

SetSimSid sets SimSid field to given value.

### HasSimSid

`func (o *WirelessV1SimUsageRecord) HasSimSid() bool`

HasSimSid returns a boolean if a field has been set.

### SetSimSidNil

`func (o *WirelessV1SimUsageRecord) SetSimSidNil(b bool)`

 SetSimSidNil sets the value for SimSid to be an explicit nil

### UnsetSimSid
`func (o *WirelessV1SimUsageRecord) UnsetSimSid()`

UnsetSimSid ensures that no value is present for SimSid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


