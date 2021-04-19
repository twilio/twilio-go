# WirelessV1AccountUsageRecord

## Properties

Name | Type | Description
------------ | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**Commands** | Pointer to **map[string]interface{}** | An object that describes the aggregated Commands usage for all SIMs during the specified period | [optional] 
**Data** | Pointer to **map[string]interface{}** | An object that describes the aggregated Data usage for all SIMs over the period | [optional] 
**Period** | Pointer to **map[string]interface{}** | The time period for which usage is reported | [optional] 

## Methods

### NewWirelessV1AccountUsageRecord

`func NewWirelessV1AccountUsageRecord() *WirelessV1AccountUsageRecord`

NewWirelessV1AccountUsageRecord instantiates a new WirelessV1AccountUsageRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWirelessV1AccountUsageRecordWithDefaults

`func NewWirelessV1AccountUsageRecordWithDefaults() *WirelessV1AccountUsageRecord`

NewWirelessV1AccountUsageRecordWithDefaults instantiates a new WirelessV1AccountUsageRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *WirelessV1AccountUsageRecord) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *WirelessV1AccountUsageRecord) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *WirelessV1AccountUsageRecord) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *WirelessV1AccountUsageRecord) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *WirelessV1AccountUsageRecord) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *WirelessV1AccountUsageRecord) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetCommands

`func (o *WirelessV1AccountUsageRecord) GetCommands() map[string]interface{}`

GetCommands returns the Commands field if non-nil, zero value otherwise.

### GetCommandsOk

`func (o *WirelessV1AccountUsageRecord) GetCommandsOk() (*map[string]interface{}, bool)`

GetCommandsOk returns a tuple with the Commands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommands

`func (o *WirelessV1AccountUsageRecord) SetCommands(v map[string]interface{})`

SetCommands sets Commands field to given value.

### HasCommands

`func (o *WirelessV1AccountUsageRecord) HasCommands() bool`

HasCommands returns a boolean if a field has been set.

### SetCommandsNil

`func (o *WirelessV1AccountUsageRecord) SetCommandsNil(b bool)`

 SetCommandsNil sets the value for Commands to be an explicit nil

### UnsetCommands
`func (o *WirelessV1AccountUsageRecord) UnsetCommands()`

UnsetCommands ensures that no value is present for Commands, not even an explicit nil
### GetData

`func (o *WirelessV1AccountUsageRecord) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WirelessV1AccountUsageRecord) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WirelessV1AccountUsageRecord) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *WirelessV1AccountUsageRecord) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *WirelessV1AccountUsageRecord) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *WirelessV1AccountUsageRecord) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetPeriod

`func (o *WirelessV1AccountUsageRecord) GetPeriod() map[string]interface{}`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *WirelessV1AccountUsageRecord) GetPeriodOk() (*map[string]interface{}, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *WirelessV1AccountUsageRecord) SetPeriod(v map[string]interface{})`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *WirelessV1AccountUsageRecord) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### SetPeriodNil

`func (o *WirelessV1AccountUsageRecord) SetPeriodNil(b bool)`

 SetPeriodNil sets the value for Period to be an explicit nil

### UnsetPeriod
`func (o *WirelessV1AccountUsageRecord) UnsetPeriod()`

UnsetPeriod ensures that no value is present for Period, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


