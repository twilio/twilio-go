# SupersimV1UsageRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that incurred the usage. | [optional] 
**DataDownload** | Pointer to **NullableInt32** | Total data downloaded in bytes, aggregated by the query parameters. | [optional] 
**DataTotal** | Pointer to **NullableInt32** | Total of data_upload and data_download. | [optional] 
**DataUpload** | Pointer to **NullableInt32** | Total data uploaded in bytes, aggregated by the query parameters. | [optional] 
**FleetSid** | Pointer to **NullableString** | SID of the Fleet resource on which the usage occurred. | [optional] 
**IsoCountry** | Pointer to **NullableString** | Alpha-2 ISO Country Code of the country the usage occurred in. | [optional] 
**NetworkSid** | Pointer to **NullableString** | SID of the Network resource on which the usage occurred. | [optional] 
**Period** | Pointer to **map[string]interface{}** | The time period for which the usage is reported. | [optional] 
**SimSid** | Pointer to **NullableString** | SID of a Sim resource to which the UsageRecord belongs. | [optional] 

## Methods

### NewSupersimV1UsageRecord

`func NewSupersimV1UsageRecord() *SupersimV1UsageRecord`

NewSupersimV1UsageRecord instantiates a new SupersimV1UsageRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupersimV1UsageRecordWithDefaults

`func NewSupersimV1UsageRecordWithDefaults() *SupersimV1UsageRecord`

NewSupersimV1UsageRecordWithDefaults instantiates a new SupersimV1UsageRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *SupersimV1UsageRecord) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *SupersimV1UsageRecord) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *SupersimV1UsageRecord) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *SupersimV1UsageRecord) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *SupersimV1UsageRecord) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *SupersimV1UsageRecord) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetDataDownload

`func (o *SupersimV1UsageRecord) GetDataDownload() int32`

GetDataDownload returns the DataDownload field if non-nil, zero value otherwise.

### GetDataDownloadOk

`func (o *SupersimV1UsageRecord) GetDataDownloadOk() (*int32, bool)`

GetDataDownloadOk returns a tuple with the DataDownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataDownload

`func (o *SupersimV1UsageRecord) SetDataDownload(v int32)`

SetDataDownload sets DataDownload field to given value.

### HasDataDownload

`func (o *SupersimV1UsageRecord) HasDataDownload() bool`

HasDataDownload returns a boolean if a field has been set.

### SetDataDownloadNil

`func (o *SupersimV1UsageRecord) SetDataDownloadNil(b bool)`

 SetDataDownloadNil sets the value for DataDownload to be an explicit nil

### UnsetDataDownload
`func (o *SupersimV1UsageRecord) UnsetDataDownload()`

UnsetDataDownload ensures that no value is present for DataDownload, not even an explicit nil
### GetDataTotal

`func (o *SupersimV1UsageRecord) GetDataTotal() int32`

GetDataTotal returns the DataTotal field if non-nil, zero value otherwise.

### GetDataTotalOk

`func (o *SupersimV1UsageRecord) GetDataTotalOk() (*int32, bool)`

GetDataTotalOk returns a tuple with the DataTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTotal

`func (o *SupersimV1UsageRecord) SetDataTotal(v int32)`

SetDataTotal sets DataTotal field to given value.

### HasDataTotal

`func (o *SupersimV1UsageRecord) HasDataTotal() bool`

HasDataTotal returns a boolean if a field has been set.

### SetDataTotalNil

`func (o *SupersimV1UsageRecord) SetDataTotalNil(b bool)`

 SetDataTotalNil sets the value for DataTotal to be an explicit nil

### UnsetDataTotal
`func (o *SupersimV1UsageRecord) UnsetDataTotal()`

UnsetDataTotal ensures that no value is present for DataTotal, not even an explicit nil
### GetDataUpload

`func (o *SupersimV1UsageRecord) GetDataUpload() int32`

GetDataUpload returns the DataUpload field if non-nil, zero value otherwise.

### GetDataUploadOk

`func (o *SupersimV1UsageRecord) GetDataUploadOk() (*int32, bool)`

GetDataUploadOk returns a tuple with the DataUpload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataUpload

`func (o *SupersimV1UsageRecord) SetDataUpload(v int32)`

SetDataUpload sets DataUpload field to given value.

### HasDataUpload

`func (o *SupersimV1UsageRecord) HasDataUpload() bool`

HasDataUpload returns a boolean if a field has been set.

### SetDataUploadNil

`func (o *SupersimV1UsageRecord) SetDataUploadNil(b bool)`

 SetDataUploadNil sets the value for DataUpload to be an explicit nil

### UnsetDataUpload
`func (o *SupersimV1UsageRecord) UnsetDataUpload()`

UnsetDataUpload ensures that no value is present for DataUpload, not even an explicit nil
### GetFleetSid

`func (o *SupersimV1UsageRecord) GetFleetSid() string`

GetFleetSid returns the FleetSid field if non-nil, zero value otherwise.

### GetFleetSidOk

`func (o *SupersimV1UsageRecord) GetFleetSidOk() (*string, bool)`

GetFleetSidOk returns a tuple with the FleetSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetSid

`func (o *SupersimV1UsageRecord) SetFleetSid(v string)`

SetFleetSid sets FleetSid field to given value.

### HasFleetSid

`func (o *SupersimV1UsageRecord) HasFleetSid() bool`

HasFleetSid returns a boolean if a field has been set.

### SetFleetSidNil

`func (o *SupersimV1UsageRecord) SetFleetSidNil(b bool)`

 SetFleetSidNil sets the value for FleetSid to be an explicit nil

### UnsetFleetSid
`func (o *SupersimV1UsageRecord) UnsetFleetSid()`

UnsetFleetSid ensures that no value is present for FleetSid, not even an explicit nil
### GetIsoCountry

`func (o *SupersimV1UsageRecord) GetIsoCountry() string`

GetIsoCountry returns the IsoCountry field if non-nil, zero value otherwise.

### GetIsoCountryOk

`func (o *SupersimV1UsageRecord) GetIsoCountryOk() (*string, bool)`

GetIsoCountryOk returns a tuple with the IsoCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoCountry

`func (o *SupersimV1UsageRecord) SetIsoCountry(v string)`

SetIsoCountry sets IsoCountry field to given value.

### HasIsoCountry

`func (o *SupersimV1UsageRecord) HasIsoCountry() bool`

HasIsoCountry returns a boolean if a field has been set.

### SetIsoCountryNil

`func (o *SupersimV1UsageRecord) SetIsoCountryNil(b bool)`

 SetIsoCountryNil sets the value for IsoCountry to be an explicit nil

### UnsetIsoCountry
`func (o *SupersimV1UsageRecord) UnsetIsoCountry()`

UnsetIsoCountry ensures that no value is present for IsoCountry, not even an explicit nil
### GetNetworkSid

`func (o *SupersimV1UsageRecord) GetNetworkSid() string`

GetNetworkSid returns the NetworkSid field if non-nil, zero value otherwise.

### GetNetworkSidOk

`func (o *SupersimV1UsageRecord) GetNetworkSidOk() (*string, bool)`

GetNetworkSidOk returns a tuple with the NetworkSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkSid

`func (o *SupersimV1UsageRecord) SetNetworkSid(v string)`

SetNetworkSid sets NetworkSid field to given value.

### HasNetworkSid

`func (o *SupersimV1UsageRecord) HasNetworkSid() bool`

HasNetworkSid returns a boolean if a field has been set.

### SetNetworkSidNil

`func (o *SupersimV1UsageRecord) SetNetworkSidNil(b bool)`

 SetNetworkSidNil sets the value for NetworkSid to be an explicit nil

### UnsetNetworkSid
`func (o *SupersimV1UsageRecord) UnsetNetworkSid()`

UnsetNetworkSid ensures that no value is present for NetworkSid, not even an explicit nil
### GetPeriod

`func (o *SupersimV1UsageRecord) GetPeriod() map[string]interface{}`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *SupersimV1UsageRecord) GetPeriodOk() (*map[string]interface{}, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *SupersimV1UsageRecord) SetPeriod(v map[string]interface{})`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *SupersimV1UsageRecord) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### SetPeriodNil

`func (o *SupersimV1UsageRecord) SetPeriodNil(b bool)`

 SetPeriodNil sets the value for Period to be an explicit nil

### UnsetPeriod
`func (o *SupersimV1UsageRecord) UnsetPeriod()`

UnsetPeriod ensures that no value is present for Period, not even an explicit nil
### GetSimSid

`func (o *SupersimV1UsageRecord) GetSimSid() string`

GetSimSid returns the SimSid field if non-nil, zero value otherwise.

### GetSimSidOk

`func (o *SupersimV1UsageRecord) GetSimSidOk() (*string, bool)`

GetSimSidOk returns a tuple with the SimSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimSid

`func (o *SupersimV1UsageRecord) SetSimSid(v string)`

SetSimSid sets SimSid field to given value.

### HasSimSid

`func (o *SupersimV1UsageRecord) HasSimSid() bool`

HasSimSid returns a boolean if a field has been set.

### SetSimSidNil

`func (o *SupersimV1UsageRecord) SetSimSidNil(b bool)`

 SetSimSidNil sets the value for SimSid to be an explicit nil

### UnsetSimSid
`func (o *SupersimV1UsageRecord) UnsetSimSid()`

UnsetSimSid ensures that no value is present for SimSid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


