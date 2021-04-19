# NumbersV2RegulatoryComplianceRegulation

## Properties

Name | Type | Description
------------ | ------------- | -------------
**EndUserType** | Pointer to **NullableString** | The type of End User of the Regulation resource | [optional] 
**FriendlyName** | Pointer to **NullableString** | A human-readable description of the Regulation resource | [optional] 
**IsoCountry** | Pointer to **NullableString** | The ISO country code of the phone number&#39;s country | [optional] 
**NumberType** | Pointer to **NullableString** | The type of phone number restricted by the regulatory requirement | [optional] 
**Requirements** | Pointer to **map[string]interface{}** | The sid of a regulation object that dictates requirements | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the Regulation resource | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the Regulation resource | [optional] 

## Methods

### NewNumbersV2RegulatoryComplianceRegulation

`func NewNumbersV2RegulatoryComplianceRegulation() *NumbersV2RegulatoryComplianceRegulation`

NewNumbersV2RegulatoryComplianceRegulation instantiates a new NumbersV2RegulatoryComplianceRegulation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNumbersV2RegulatoryComplianceRegulationWithDefaults

`func NewNumbersV2RegulatoryComplianceRegulationWithDefaults() *NumbersV2RegulatoryComplianceRegulation`

NewNumbersV2RegulatoryComplianceRegulationWithDefaults instantiates a new NumbersV2RegulatoryComplianceRegulation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndUserType

`func (o *NumbersV2RegulatoryComplianceRegulation) GetEndUserType() string`

GetEndUserType returns the EndUserType field if non-nil, zero value otherwise.

### GetEndUserTypeOk

`func (o *NumbersV2RegulatoryComplianceRegulation) GetEndUserTypeOk() (*string, bool)`

GetEndUserTypeOk returns a tuple with the EndUserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndUserType

`func (o *NumbersV2RegulatoryComplianceRegulation) SetEndUserType(v string)`

SetEndUserType sets EndUserType field to given value.

### HasEndUserType

`func (o *NumbersV2RegulatoryComplianceRegulation) HasEndUserType() bool`

HasEndUserType returns a boolean if a field has been set.

### SetEndUserTypeNil

`func (o *NumbersV2RegulatoryComplianceRegulation) SetEndUserTypeNil(b bool)`

 SetEndUserTypeNil sets the value for EndUserType to be an explicit nil

### UnsetEndUserType
`func (o *NumbersV2RegulatoryComplianceRegulation) UnsetEndUserType()`

UnsetEndUserType ensures that no value is present for EndUserType, not even an explicit nil
### GetFriendlyName

`func (o *NumbersV2RegulatoryComplianceRegulation) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *NumbersV2RegulatoryComplianceRegulation) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *NumbersV2RegulatoryComplianceRegulation) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *NumbersV2RegulatoryComplianceRegulation) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *NumbersV2RegulatoryComplianceRegulation) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *NumbersV2RegulatoryComplianceRegulation) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetIsoCountry

`func (o *NumbersV2RegulatoryComplianceRegulation) GetIsoCountry() string`

GetIsoCountry returns the IsoCountry field if non-nil, zero value otherwise.

### GetIsoCountryOk

`func (o *NumbersV2RegulatoryComplianceRegulation) GetIsoCountryOk() (*string, bool)`

GetIsoCountryOk returns a tuple with the IsoCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoCountry

`func (o *NumbersV2RegulatoryComplianceRegulation) SetIsoCountry(v string)`

SetIsoCountry sets IsoCountry field to given value.

### HasIsoCountry

`func (o *NumbersV2RegulatoryComplianceRegulation) HasIsoCountry() bool`

HasIsoCountry returns a boolean if a field has been set.

### SetIsoCountryNil

`func (o *NumbersV2RegulatoryComplianceRegulation) SetIsoCountryNil(b bool)`

 SetIsoCountryNil sets the value for IsoCountry to be an explicit nil

### UnsetIsoCountry
`func (o *NumbersV2RegulatoryComplianceRegulation) UnsetIsoCountry()`

UnsetIsoCountry ensures that no value is present for IsoCountry, not even an explicit nil
### GetNumberType

`func (o *NumbersV2RegulatoryComplianceRegulation) GetNumberType() string`

GetNumberType returns the NumberType field if non-nil, zero value otherwise.

### GetNumberTypeOk

`func (o *NumbersV2RegulatoryComplianceRegulation) GetNumberTypeOk() (*string, bool)`

GetNumberTypeOk returns a tuple with the NumberType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberType

`func (o *NumbersV2RegulatoryComplianceRegulation) SetNumberType(v string)`

SetNumberType sets NumberType field to given value.

### HasNumberType

`func (o *NumbersV2RegulatoryComplianceRegulation) HasNumberType() bool`

HasNumberType returns a boolean if a field has been set.

### SetNumberTypeNil

`func (o *NumbersV2RegulatoryComplianceRegulation) SetNumberTypeNil(b bool)`

 SetNumberTypeNil sets the value for NumberType to be an explicit nil

### UnsetNumberType
`func (o *NumbersV2RegulatoryComplianceRegulation) UnsetNumberType()`

UnsetNumberType ensures that no value is present for NumberType, not even an explicit nil
### GetRequirements

`func (o *NumbersV2RegulatoryComplianceRegulation) GetRequirements() map[string]interface{}`

GetRequirements returns the Requirements field if non-nil, zero value otherwise.

### GetRequirementsOk

`func (o *NumbersV2RegulatoryComplianceRegulation) GetRequirementsOk() (*map[string]interface{}, bool)`

GetRequirementsOk returns a tuple with the Requirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirements

`func (o *NumbersV2RegulatoryComplianceRegulation) SetRequirements(v map[string]interface{})`

SetRequirements sets Requirements field to given value.

### HasRequirements

`func (o *NumbersV2RegulatoryComplianceRegulation) HasRequirements() bool`

HasRequirements returns a boolean if a field has been set.

### SetRequirementsNil

`func (o *NumbersV2RegulatoryComplianceRegulation) SetRequirementsNil(b bool)`

 SetRequirementsNil sets the value for Requirements to be an explicit nil

### UnsetRequirements
`func (o *NumbersV2RegulatoryComplianceRegulation) UnsetRequirements()`

UnsetRequirements ensures that no value is present for Requirements, not even an explicit nil
### GetSid

`func (o *NumbersV2RegulatoryComplianceRegulation) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *NumbersV2RegulatoryComplianceRegulation) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *NumbersV2RegulatoryComplianceRegulation) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *NumbersV2RegulatoryComplianceRegulation) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *NumbersV2RegulatoryComplianceRegulation) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *NumbersV2RegulatoryComplianceRegulation) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetUrl

`func (o *NumbersV2RegulatoryComplianceRegulation) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NumbersV2RegulatoryComplianceRegulation) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NumbersV2RegulatoryComplianceRegulation) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *NumbersV2RegulatoryComplianceRegulation) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *NumbersV2RegulatoryComplianceRegulation) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *NumbersV2RegulatoryComplianceRegulation) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


