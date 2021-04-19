# WirelessV1SimDataSession

## Properties

Name | Type | Description
------------ | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**CellId** | Pointer to **NullableString** | The unique ID of the cellular tower that the device was attached to at the moment when the Data Session was last updated | [optional] 
**CellLocationEstimate** | Pointer to **map[string]interface{}** | An object with the estimated location where the device&#39;s Data Session took place | [optional] 
**End** | Pointer to **NullableTime** | The date that the record ended, given as GMT in ISO 8601 format | [optional] 
**Imei** | Pointer to **NullableString** | The unique ID of the device using the SIM to connect | [optional] 
**LastUpdated** | Pointer to **NullableTime** | The date that the resource was last updated, given as GMT in ISO 8601 format | [optional] 
**OperatorCountry** | Pointer to **NullableString** | The three letter country code representing where the device&#39;s Data Session took place | [optional] 
**OperatorMcc** | Pointer to **NullableString** | The &#39;mobile country code&#39; is the unique ID of the home country where the Data Session took place | [optional] 
**OperatorMnc** | Pointer to **NullableString** | The &#39;mobile network code&#39; is the unique ID specific to the mobile operator network where the Data Session took place | [optional] 
**OperatorName** | Pointer to **NullableString** | The friendly name of the mobile operator network that the SIM-connected device is attached to | [optional] 
**PacketsDownloaded** | Pointer to **NullableInt32** | The number of packets downloaded by the device between the start time and when the Data Session was last updated | [optional] 
**PacketsUploaded** | Pointer to **NullableInt32** | The number of packets uploaded by the device between the start time and when the Data Session was last updated | [optional] 
**RadioLink** | Pointer to **NullableString** | The generation of wireless technology that the device was using | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**SimSid** | Pointer to **NullableString** | The SID of the Sim resource that the Data Session is for | [optional] 
**Start** | Pointer to **NullableTime** | The date that the Data Session started, given as GMT in ISO 8601 format | [optional] 

## Methods

### NewWirelessV1SimDataSession

`func NewWirelessV1SimDataSession() *WirelessV1SimDataSession`

NewWirelessV1SimDataSession instantiates a new WirelessV1SimDataSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWirelessV1SimDataSessionWithDefaults

`func NewWirelessV1SimDataSessionWithDefaults() *WirelessV1SimDataSession`

NewWirelessV1SimDataSessionWithDefaults instantiates a new WirelessV1SimDataSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *WirelessV1SimDataSession) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *WirelessV1SimDataSession) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *WirelessV1SimDataSession) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *WirelessV1SimDataSession) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *WirelessV1SimDataSession) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *WirelessV1SimDataSession) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetCellId

`func (o *WirelessV1SimDataSession) GetCellId() string`

GetCellId returns the CellId field if non-nil, zero value otherwise.

### GetCellIdOk

`func (o *WirelessV1SimDataSession) GetCellIdOk() (*string, bool)`

GetCellIdOk returns a tuple with the CellId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCellId

`func (o *WirelessV1SimDataSession) SetCellId(v string)`

SetCellId sets CellId field to given value.

### HasCellId

`func (o *WirelessV1SimDataSession) HasCellId() bool`

HasCellId returns a boolean if a field has been set.

### SetCellIdNil

`func (o *WirelessV1SimDataSession) SetCellIdNil(b bool)`

 SetCellIdNil sets the value for CellId to be an explicit nil

### UnsetCellId
`func (o *WirelessV1SimDataSession) UnsetCellId()`

UnsetCellId ensures that no value is present for CellId, not even an explicit nil
### GetCellLocationEstimate

`func (o *WirelessV1SimDataSession) GetCellLocationEstimate() map[string]interface{}`

GetCellLocationEstimate returns the CellLocationEstimate field if non-nil, zero value otherwise.

### GetCellLocationEstimateOk

`func (o *WirelessV1SimDataSession) GetCellLocationEstimateOk() (*map[string]interface{}, bool)`

GetCellLocationEstimateOk returns a tuple with the CellLocationEstimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCellLocationEstimate

`func (o *WirelessV1SimDataSession) SetCellLocationEstimate(v map[string]interface{})`

SetCellLocationEstimate sets CellLocationEstimate field to given value.

### HasCellLocationEstimate

`func (o *WirelessV1SimDataSession) HasCellLocationEstimate() bool`

HasCellLocationEstimate returns a boolean if a field has been set.

### SetCellLocationEstimateNil

`func (o *WirelessV1SimDataSession) SetCellLocationEstimateNil(b bool)`

 SetCellLocationEstimateNil sets the value for CellLocationEstimate to be an explicit nil

### UnsetCellLocationEstimate
`func (o *WirelessV1SimDataSession) UnsetCellLocationEstimate()`

UnsetCellLocationEstimate ensures that no value is present for CellLocationEstimate, not even an explicit nil
### GetEnd

`func (o *WirelessV1SimDataSession) GetEnd() time.Time`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *WirelessV1SimDataSession) GetEndOk() (*time.Time, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *WirelessV1SimDataSession) SetEnd(v time.Time)`

SetEnd sets End field to given value.

### HasEnd

`func (o *WirelessV1SimDataSession) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### SetEndNil

`func (o *WirelessV1SimDataSession) SetEndNil(b bool)`

 SetEndNil sets the value for End to be an explicit nil

### UnsetEnd
`func (o *WirelessV1SimDataSession) UnsetEnd()`

UnsetEnd ensures that no value is present for End, not even an explicit nil
### GetImei

`func (o *WirelessV1SimDataSession) GetImei() string`

GetImei returns the Imei field if non-nil, zero value otherwise.

### GetImeiOk

`func (o *WirelessV1SimDataSession) GetImeiOk() (*string, bool)`

GetImeiOk returns a tuple with the Imei field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImei

`func (o *WirelessV1SimDataSession) SetImei(v string)`

SetImei sets Imei field to given value.

### HasImei

`func (o *WirelessV1SimDataSession) HasImei() bool`

HasImei returns a boolean if a field has been set.

### SetImeiNil

`func (o *WirelessV1SimDataSession) SetImeiNil(b bool)`

 SetImeiNil sets the value for Imei to be an explicit nil

### UnsetImei
`func (o *WirelessV1SimDataSession) UnsetImei()`

UnsetImei ensures that no value is present for Imei, not even an explicit nil
### GetLastUpdated

`func (o *WirelessV1SimDataSession) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *WirelessV1SimDataSession) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *WirelessV1SimDataSession) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *WirelessV1SimDataSession) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### SetLastUpdatedNil

`func (o *WirelessV1SimDataSession) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *WirelessV1SimDataSession) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetOperatorCountry

`func (o *WirelessV1SimDataSession) GetOperatorCountry() string`

GetOperatorCountry returns the OperatorCountry field if non-nil, zero value otherwise.

### GetOperatorCountryOk

`func (o *WirelessV1SimDataSession) GetOperatorCountryOk() (*string, bool)`

GetOperatorCountryOk returns a tuple with the OperatorCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorCountry

`func (o *WirelessV1SimDataSession) SetOperatorCountry(v string)`

SetOperatorCountry sets OperatorCountry field to given value.

### HasOperatorCountry

`func (o *WirelessV1SimDataSession) HasOperatorCountry() bool`

HasOperatorCountry returns a boolean if a field has been set.

### SetOperatorCountryNil

`func (o *WirelessV1SimDataSession) SetOperatorCountryNil(b bool)`

 SetOperatorCountryNil sets the value for OperatorCountry to be an explicit nil

### UnsetOperatorCountry
`func (o *WirelessV1SimDataSession) UnsetOperatorCountry()`

UnsetOperatorCountry ensures that no value is present for OperatorCountry, not even an explicit nil
### GetOperatorMcc

`func (o *WirelessV1SimDataSession) GetOperatorMcc() string`

GetOperatorMcc returns the OperatorMcc field if non-nil, zero value otherwise.

### GetOperatorMccOk

`func (o *WirelessV1SimDataSession) GetOperatorMccOk() (*string, bool)`

GetOperatorMccOk returns a tuple with the OperatorMcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorMcc

`func (o *WirelessV1SimDataSession) SetOperatorMcc(v string)`

SetOperatorMcc sets OperatorMcc field to given value.

### HasOperatorMcc

`func (o *WirelessV1SimDataSession) HasOperatorMcc() bool`

HasOperatorMcc returns a boolean if a field has been set.

### SetOperatorMccNil

`func (o *WirelessV1SimDataSession) SetOperatorMccNil(b bool)`

 SetOperatorMccNil sets the value for OperatorMcc to be an explicit nil

### UnsetOperatorMcc
`func (o *WirelessV1SimDataSession) UnsetOperatorMcc()`

UnsetOperatorMcc ensures that no value is present for OperatorMcc, not even an explicit nil
### GetOperatorMnc

`func (o *WirelessV1SimDataSession) GetOperatorMnc() string`

GetOperatorMnc returns the OperatorMnc field if non-nil, zero value otherwise.

### GetOperatorMncOk

`func (o *WirelessV1SimDataSession) GetOperatorMncOk() (*string, bool)`

GetOperatorMncOk returns a tuple with the OperatorMnc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorMnc

`func (o *WirelessV1SimDataSession) SetOperatorMnc(v string)`

SetOperatorMnc sets OperatorMnc field to given value.

### HasOperatorMnc

`func (o *WirelessV1SimDataSession) HasOperatorMnc() bool`

HasOperatorMnc returns a boolean if a field has been set.

### SetOperatorMncNil

`func (o *WirelessV1SimDataSession) SetOperatorMncNil(b bool)`

 SetOperatorMncNil sets the value for OperatorMnc to be an explicit nil

### UnsetOperatorMnc
`func (o *WirelessV1SimDataSession) UnsetOperatorMnc()`

UnsetOperatorMnc ensures that no value is present for OperatorMnc, not even an explicit nil
### GetOperatorName

`func (o *WirelessV1SimDataSession) GetOperatorName() string`

GetOperatorName returns the OperatorName field if non-nil, zero value otherwise.

### GetOperatorNameOk

`func (o *WirelessV1SimDataSession) GetOperatorNameOk() (*string, bool)`

GetOperatorNameOk returns a tuple with the OperatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorName

`func (o *WirelessV1SimDataSession) SetOperatorName(v string)`

SetOperatorName sets OperatorName field to given value.

### HasOperatorName

`func (o *WirelessV1SimDataSession) HasOperatorName() bool`

HasOperatorName returns a boolean if a field has been set.

### SetOperatorNameNil

`func (o *WirelessV1SimDataSession) SetOperatorNameNil(b bool)`

 SetOperatorNameNil sets the value for OperatorName to be an explicit nil

### UnsetOperatorName
`func (o *WirelessV1SimDataSession) UnsetOperatorName()`

UnsetOperatorName ensures that no value is present for OperatorName, not even an explicit nil
### GetPacketsDownloaded

`func (o *WirelessV1SimDataSession) GetPacketsDownloaded() int32`

GetPacketsDownloaded returns the PacketsDownloaded field if non-nil, zero value otherwise.

### GetPacketsDownloadedOk

`func (o *WirelessV1SimDataSession) GetPacketsDownloadedOk() (*int32, bool)`

GetPacketsDownloadedOk returns a tuple with the PacketsDownloaded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPacketsDownloaded

`func (o *WirelessV1SimDataSession) SetPacketsDownloaded(v int32)`

SetPacketsDownloaded sets PacketsDownloaded field to given value.

### HasPacketsDownloaded

`func (o *WirelessV1SimDataSession) HasPacketsDownloaded() bool`

HasPacketsDownloaded returns a boolean if a field has been set.

### SetPacketsDownloadedNil

`func (o *WirelessV1SimDataSession) SetPacketsDownloadedNil(b bool)`

 SetPacketsDownloadedNil sets the value for PacketsDownloaded to be an explicit nil

### UnsetPacketsDownloaded
`func (o *WirelessV1SimDataSession) UnsetPacketsDownloaded()`

UnsetPacketsDownloaded ensures that no value is present for PacketsDownloaded, not even an explicit nil
### GetPacketsUploaded

`func (o *WirelessV1SimDataSession) GetPacketsUploaded() int32`

GetPacketsUploaded returns the PacketsUploaded field if non-nil, zero value otherwise.

### GetPacketsUploadedOk

`func (o *WirelessV1SimDataSession) GetPacketsUploadedOk() (*int32, bool)`

GetPacketsUploadedOk returns a tuple with the PacketsUploaded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPacketsUploaded

`func (o *WirelessV1SimDataSession) SetPacketsUploaded(v int32)`

SetPacketsUploaded sets PacketsUploaded field to given value.

### HasPacketsUploaded

`func (o *WirelessV1SimDataSession) HasPacketsUploaded() bool`

HasPacketsUploaded returns a boolean if a field has been set.

### SetPacketsUploadedNil

`func (o *WirelessV1SimDataSession) SetPacketsUploadedNil(b bool)`

 SetPacketsUploadedNil sets the value for PacketsUploaded to be an explicit nil

### UnsetPacketsUploaded
`func (o *WirelessV1SimDataSession) UnsetPacketsUploaded()`

UnsetPacketsUploaded ensures that no value is present for PacketsUploaded, not even an explicit nil
### GetRadioLink

`func (o *WirelessV1SimDataSession) GetRadioLink() string`

GetRadioLink returns the RadioLink field if non-nil, zero value otherwise.

### GetRadioLinkOk

`func (o *WirelessV1SimDataSession) GetRadioLinkOk() (*string, bool)`

GetRadioLinkOk returns a tuple with the RadioLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadioLink

`func (o *WirelessV1SimDataSession) SetRadioLink(v string)`

SetRadioLink sets RadioLink field to given value.

### HasRadioLink

`func (o *WirelessV1SimDataSession) HasRadioLink() bool`

HasRadioLink returns a boolean if a field has been set.

### SetRadioLinkNil

`func (o *WirelessV1SimDataSession) SetRadioLinkNil(b bool)`

 SetRadioLinkNil sets the value for RadioLink to be an explicit nil

### UnsetRadioLink
`func (o *WirelessV1SimDataSession) UnsetRadioLink()`

UnsetRadioLink ensures that no value is present for RadioLink, not even an explicit nil
### GetSid

`func (o *WirelessV1SimDataSession) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *WirelessV1SimDataSession) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *WirelessV1SimDataSession) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *WirelessV1SimDataSession) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *WirelessV1SimDataSession) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *WirelessV1SimDataSession) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetSimSid

`func (o *WirelessV1SimDataSession) GetSimSid() string`

GetSimSid returns the SimSid field if non-nil, zero value otherwise.

### GetSimSidOk

`func (o *WirelessV1SimDataSession) GetSimSidOk() (*string, bool)`

GetSimSidOk returns a tuple with the SimSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimSid

`func (o *WirelessV1SimDataSession) SetSimSid(v string)`

SetSimSid sets SimSid field to given value.

### HasSimSid

`func (o *WirelessV1SimDataSession) HasSimSid() bool`

HasSimSid returns a boolean if a field has been set.

### SetSimSidNil

`func (o *WirelessV1SimDataSession) SetSimSidNil(b bool)`

 SetSimSidNil sets the value for SimSid to be an explicit nil

### UnsetSimSid
`func (o *WirelessV1SimDataSession) UnsetSimSid()`

UnsetSimSid ensures that no value is present for SimSid, not even an explicit nil
### GetStart

`func (o *WirelessV1SimDataSession) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *WirelessV1SimDataSession) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *WirelessV1SimDataSession) SetStart(v time.Time)`

SetStart sets Start field to given value.

### HasStart

`func (o *WirelessV1SimDataSession) HasStart() bool`

HasStart returns a boolean if a field has been set.

### SetStartNil

`func (o *WirelessV1SimDataSession) SetStartNil(b bool)`

 SetStartNil sets the value for Start to be an explicit nil

### UnsetStart
`func (o *WirelessV1SimDataSession) UnsetStart()`

UnsetStart ensures that no value is present for Start, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


